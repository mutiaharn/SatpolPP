import { get } from 'svelte/store';
import type { Article, Complaint, ComplaintStatus, EmailNotice, OrganizationNode, PublicDocument, Regency, SiteSettings, User } from '$lib/types';
import {
  siteSettingsStore,
  articlesStore,
  documentsStore,
  complaintsStore,
  usersStore,
  organizationStore,
  logAudit,
  defaultRegencies
} from '$lib/stores/appState';
import { maskComplaintIdentity } from '$lib/utils/masking';

const API_BASE = '/api/v1';

function authHeaders(): Record<string, string> {
  if (typeof sessionStorage === 'undefined') return {};
  try {
    const raw = sessionStorage.getItem('satpol_auth');
    const token = raw ? JSON.parse(raw).token : '';
    return token ? { Authorization: `Bearer ${token}` } : {};
  } catch {
    return {};
  }
}

async function tryFetch<T>(url: string, init?: RequestInit): Promise<T | null> {
  try {
    const res = await fetch(url, init);
    if (res.ok) return (await res.json()) as T;
  } catch (e) {
    // Graceful offline fallback (backend belum berjalan)
  }
  return null;
}

export async function fetchSiteSettings(): Promise<SiteSettings> {
  const data = await tryFetch<{ settings?: SiteSettings }>(`${API_BASE}/portal/settings`);
  return data?.settings ?? get(siteSettingsStore);
}

export async function fetchArticles(): Promise<Article[]> {
  const data = await tryFetch<{ data?: Article[] }>(`${API_BASE}/portal/articles`);
  return data?.data ?? get(articlesStore);
}

export async function fetchArticleBySlug(slug: string): Promise<Article | undefined> {
  const data = await tryFetch<{ data?: Article }>(`${API_BASE}/portal/articles/${slug}`);
  return data?.data ?? get(articlesStore).find(a => a.slug === slug);
}

export async function fetchDocuments(category?: string): Promise<PublicDocument[]> {
  const url = category ? `${API_BASE}/portal/documents?category=${category}` : `${API_BASE}/portal/documents`;
  const data = await tryFetch<{ data?: PublicDocument[] }>(url);
  if (data?.data) return data.data;
  const all = get(documentsStore);
  if (!category) return all;
  return all.filter(d => d.category.toLowerCase() === category.toLowerCase());
}

export async function fetchRegencies(): Promise<Regency[]> {
  const data = await tryFetch<{ data?: Regency[] }>(`${API_BASE}/portal/regencies`);
  return data?.data ?? defaultRegencies;
}

// ==========================================================
// SIAP TERTIB
// ==========================================================

/** Unggah foto bukti ke backend. Mengembalikan URL server, atau null bila backend offline. */
export async function uploadEvidence(file: File): Promise<string | null> {
  try {
    const form = new FormData();
    form.append('foto', file);
    const res = await fetch(`${API_BASE}/upload/evidence`, { method: 'POST', body: form });
    if (res.ok) {
      const data = await res.json();
      return data.image_url || null;
    }
  } catch (e) {
    // offline
  }
  return null;
}

function buildEmailSubject(ticket: string, statusLabel?: string) {
  return statusLabel
    ? `[SIAP TERTIB] Pembaruan Status Tiket ${ticket}: ${statusLabel}`
    : `[SIAP TERTIB] Laporan Anda Diterima - Tiket ${ticket}`;
}

export interface SubmitComplaintResult {
  success: boolean;
  ticket_code: string;
  message: string;
  email: EmailNotice;
  data?: Complaint;
}

export async function submitComplaint(payload: Partial<Complaint>): Promise<SubmitComplaintResult> {
  const remote = await tryFetch<SubmitComplaintResult>(`${API_BASE}/siaptertib/lapor`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });

  if (remote?.success) {
    const stored = buildLocalComplaint(payload, remote.ticket_code);
    // Simpan juga salinan lokal agar dashboard mode offline tetap sinkron
    complaintsStore.update(items => [stored, ...items.filter(i => i.ticket_code !== stored.ticket_code)]);
    return {
      ...remote,
      email: remote.email ?? {
        sent: false,
        mode: 'simulated',
        to: payload.reporter_email || '',
        subject: buildEmailSubject(remote.ticket_code),
        message: 'Server tidak mengirimkan status email.'
      },
      data: stored
    };
  }

  // ===== Offline fallback =====
  const now = new Date();
  const randomSuffix = Math.floor(1000 + Math.random() * 9000);
  const code = `TRB-${now.getFullYear()}${String(now.getMonth() + 1).padStart(2, '0')}-${randomSuffix}`;
  const newComplaint = buildLocalComplaint(payload, code);

  complaintsStore.update(items => [newComplaint, ...items]);
  logAudit('Masyarakat (SIAP TERTIB)', 'SUBMIT_ADUAN', `Complaint #${code}`, `Laporan baru diajukan terkait ${newComplaint.violation_type}`);

  return {
    success: true,
    ticket_code: code,
    message: 'Laporan Anda telah berhasil diregistrasi di sistem SIAP TERTIB Satpol PP Sulsel.',
    email: {
      sent: false,
      mode: payload.reporter_email ? 'simulated' : 'skipped',
      to: payload.reporter_email || '',
      subject: buildEmailSubject(code),
      message: 'Backend belum berjalan, email ditampilkan sebagai simulasi pratinjau.'
    },
    data: newComplaint
  };
}

function buildLocalComplaint(payload: Partial<Complaint>, code: string): Complaint {
  return {
    id: Date.now(),
    ticket_code: code,
    is_anonymous: !!payload.is_anonymous,
    reporter_nik: payload.reporter_nik || '',
    reporter_name: payload.reporter_name || 'Warga',
    reporter_phone: payload.reporter_phone || '',
    reporter_email: payload.reporter_email || '',
    violation_type: payload.violation_type || 'Gangguan Ketenteraman & Ketertiban Umum',
    regency_id: payload.regency_id || '7301',
    regency_name: payload.regency_name || 'Kota Makassar',
    district_name: payload.district_name || '',
    village_name: payload.village_name || '',
    address_detail: payload.address_detail || '',
    latitude: payload.latitude || -5.147665,
    longitude: payload.longitude || 119.432731,
    description: payload.description || '',
    evidence_image_url: payload.evidence_image_url || '',
    status: 'MENUNGGU_VERIFIKASI',
    disposed_to: '',
    action_notes: '',
    created_at: new Date().toISOString()
  };
}

/** Data publik: identitas pelapor SELALU disamarkan. */
function toPublicComplaint(c: Complaint): Complaint {
  const masked = maskComplaintIdentity(c);
  return {
    ...c,
    reporter_name: masked.name,
    reporter_nik: masked.nik,
    reporter_phone: masked.phone,
    reporter_email: masked.email,
    masked_name: masked.name,
    masked_nik: masked.nik,
    masked_phone: masked.phone,
    masked_email: masked.email,
    evidence_image_url: c.evidence_image_url || c.evidence_photos || ''
  };
}

export async function trackComplaint(ticketCode: string): Promise<{ success: boolean; data?: Complaint; message?: string }> {
  const code = ticketCode.trim().toUpperCase();
  const remote = await tryFetch<{ success: boolean; data?: Complaint }>(`${API_BASE}/siaptertib/lacak/${encodeURIComponent(code)}`);
  if (remote?.success && remote.data) {
    return { success: true, data: toPublicComplaint(remote.data) };
  }

  const found = get(complaintsStore).find(c => c.ticket_code.toUpperCase() === code);
  if (found) return { success: true, data: toPublicComplaint(found) };
  return { success: false, message: 'Nomor tiket pengaduan tidak ditemukan dalam sistem.' };
}

const STATUS_LABEL: Record<ComplaintStatus, string> = {
  MENUNGGU_VERIFIKASI: 'Menunggu Verifikasi',
  DISPOSISI_BIDANG: 'Disposisi Bidang',
  PENINDAKAN_LAPANGAN: 'Penindakan Lapangan',
  SELESAI: 'Selesai',
  DITOLAK: 'Ditolak'
};

export async function updateComplaintStatus(
  ticket: string,
  update: { status: ComplaintStatus; disposed_to: string; action_notes: string; officer_name: string }
): Promise<{ success: boolean; email: EmailNotice }> {
  const remote = await tryFetch<{ success: boolean; email?: EmailNotice }>(`${API_BASE}/admin/complaints/${encodeURIComponent(ticket)}/status`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify(update)
  });

  let reporterEmail = '';
  complaintsStore.update(list =>
    list.map(item => {
      if (item.ticket_code !== ticket) return item;
      reporterEmail = item.reporter_email;
      return { ...item, status: update.status, disposed_to: update.disposed_to, action_notes: update.action_notes };
    })
  );

  return {
    success: true,
    email: remote?.email ?? {
      sent: false,
      mode: reporterEmail ? 'simulated' : 'skipped',
      to: reporterEmail,
      subject: buildEmailSubject(ticket, STATUS_LABEL[update.status]),
      message: reporterEmail ? 'Backend belum berjalan, notifikasi email disimulasikan.' : 'Pelapor tidak mencantumkan email.'
    }
  };
}

// ==========================================================
// AUTH
// ==========================================================

/** Kata sandi akun demo saat backend offline (sama dengan seeder backend Go) */
const DEMO_PASSWORDS: Record<string, string> = {
  superadmin: 'sulselprov2026!',
  kasatpol: 'satpolpp2026!',
  stafftrantibum: 'trantibum2026!'
};
const DEFAULT_PASSWORD = 'satpolpp2026';

export async function loginAdmin(username: string, password: string): Promise<{ success?: boolean; token?: string; user?: User; error?: string }> {
  try {
    const res = await fetch(`${API_BASE}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    if (res.ok) return await res.json();
    if (res.status === 401 || res.status === 403) {
      const data = await res.json().catch(() => ({}));
      return { success: false, error: data.error || 'Username atau kata sandi tidak valid.' };
    }
  } catch (e) {
    // Graceful offline fallback
  }

  const uname = username.trim().toLowerCase();
  const matched = get(usersStore).find(u => u.username.toLowerCase() === uname);
  const validPassword = password === DEFAULT_PASSWORD || (DEMO_PASSWORDS[uname] && password === DEMO_PASSWORDS[uname]);

  if (matched && validPassword) {
    if (!matched.is_active) {
      return { success: false, error: 'Akun ini sedang dinonaktifkan oleh Superadmin.' };
    }
    return { success: true, token: `jwt-mock-${matched.role}-${matched.id}`, user: matched };
  }
  return { success: false, error: 'Username atau kata sandi tidak valid.' };
}

// ==========================================================
// SUPERADMIN: MANAJEMEN PENGGUNA
// ==========================================================

export async function saveUser(user: User, isNew: boolean, password?: string): Promise<User> {
  const url = isNew ? `${API_BASE}/superadmin/users` : `${API_BASE}/superadmin/users/${user.id}`;
  const remote = await tryFetch<{ data?: User }>(url, {
    method: isNew ? 'POST' : 'PUT',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify({ ...user, password })
  });
  const saved = remote?.data ?? user;
  usersStore.update(list => (isNew ? [...list, saved] : list.map(u => (u.id === saved.id ? saved : u))));
  return saved;
}

export async function setUserActive(id: number, isActive: boolean): Promise<void> {
  await tryFetch(`${API_BASE}/superadmin/users/${id}/status`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify({ is_active: isActive })
  });
  usersStore.update(list => list.map(u => (u.id === id ? { ...u, is_active: isActive } : u)));
}

export async function deleteUser(id: number): Promise<void> {
  await tryFetch(`${API_BASE}/superadmin/users/${id}`, { method: 'DELETE', headers: authHeaders() });
  usersStore.update(list => list.filter(u => u.id !== id));
}

// ==========================================================
// STRUKTUR ORGANISASI DINAMIS
// ==========================================================

export async function fetchOrganization(): Promise<OrganizationNode[]> {
  const data = await tryFetch<{ data?: OrganizationNode[] }>(`${API_BASE}/portal/organization`);
  if (data?.data && data.data.length) {
    organizationStore.set(data.data);
    return data.data;
  }
  return get(organizationStore);
}

export async function saveOrganizationNode(node: OrganizationNode, isNew: boolean): Promise<OrganizationNode> {
  const url = isNew ? `${API_BASE}/admin/organization` : `${API_BASE}/admin/organization/${node.id}`;
  const remote = await tryFetch<{ data?: OrganizationNode }>(url, {
    method: isNew ? 'POST' : 'PUT',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify(node)
  });
  const saved = remote?.data ?? node;
  organizationStore.update(list => (isNew ? [...list, saved] : list.map(n => (n.id === saved.id ? saved : n))));
  return saved;
}

export async function deleteOrganizationNode(id: number): Promise<void> {
  await tryFetch(`${API_BASE}/admin/organization/${id}`, { method: 'DELETE', headers: authHeaders() });
  organizationStore.update(list => list.filter(n => n.id !== id));
}
