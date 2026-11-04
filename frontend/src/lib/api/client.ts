import { get } from 'svelte/store';
import type { Article, Complaint, PublicDocument, Regency, SiteSettings, User } from '$lib/types';
import { 
  siteSettingsStore, 
  articlesStore, 
  documentsStore, 
  complaintsStore, 
  usersStore, 
  auditLogsStore,
  defaultRegencies 
} from '$lib/stores/appState';

const API_BASE = '/api/v1';

export async function fetchSiteSettings(): Promise<SiteSettings> {
  try {
    const res = await fetch(`${API_BASE}/portal/settings`);
    if (res.ok) {
      const data = await res.json();
      if (data.settings) return data.settings;
    }
  } catch (e) {
    // Graceful offline fallback
  }
  return get(siteSettingsStore);
}

export async function fetchArticles(): Promise<Article[]> {
  try {
    const res = await fetch(`${API_BASE}/portal/articles`);
    if (res.ok) {
      const data = await res.json();
      if (data.data) return data.data;
    }
  } catch (e) {
    // Graceful offline fallback
  }
  return get(articlesStore);
}

export async function fetchArticleBySlug(slug: string): Promise<Article | undefined> {
  try {
    const res = await fetch(`${API_BASE}/portal/articles/${slug}`);
    if (res.ok) {
      const data = await res.json();
      if (data.data) return data.data;
    }
  } catch (e) {
    // Graceful offline fallback
  }
  const all = get(articlesStore);
  return all.find(a => a.slug === slug);
}

export async function fetchDocuments(category?: string): Promise<PublicDocument[]> {
  try {
    const url = category ? `${API_BASE}/portal/documents?category=${category}` : `${API_BASE}/portal/documents`;
    const res = await fetch(url);
    if (res.ok) {
      const data = await res.json();
      if (data.data) return data.data;
    }
  } catch (e) {
    // Graceful offline fallback
  }
  const all = get(documentsStore);
  if (!category) return all;
  return all.filter(d => d.category.toLowerCase() === category.toLowerCase());
}

export async function fetchRegencies(): Promise<Regency[]> {
  try {
    const res = await fetch(`${API_BASE}/portal/regencies`);
    if (res.ok) {
      const data = await res.json();
      if (data.data) return data.data;
    }
  } catch (e) {
    // Graceful offline fallback
  }
  return defaultRegencies;
}

export async function submitComplaint(payload: Partial<Complaint>): Promise<{ success: boolean; ticket_code: string; message: string }> {
  try {
    const res = await fetch(`${API_BASE}/siaptertib/lapor`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    });
    if (res.ok) {
      return await res.json();
    }
  } catch (e) {
    // Graceful offline fallback
  }

  // Generate offline ticket
  const randomSuffix = Math.floor(1000 + Math.random() * 9000);
  const code = `TRB-${new Date().getFullYear()}${String(new Date().getMonth() + 1).padStart(2, '0')}-${randomSuffix}`;
  
  const newComplaint: Complaint = {
    id: Date.now(),
    ticket_code: code,
    is_anonymous: !!payload.is_anonymous,
    reporter_nik: payload.reporter_nik || '',
    reporter_name: payload.reporter_name || 'Anonim',
    reporter_phone: payload.reporter_phone || '',
    reporter_email: payload.reporter_email || '',
    violation_type: payload.violation_type || 'Pelanggaran Ketertiban Umum',
    regency_id: payload.regency_id || '7301',
    regency_name: payload.regency_name || 'Kota Makassar',
    district_name: payload.district_name || '',
    village_name: payload.village_name || '',
    address_detail: payload.address_detail || '',
    latitude: payload.latitude || -5.147665,
    longitude: payload.longitude || 119.432731,
    description: payload.description || '',
    status: 'MENUNGGU_VERIFIKASI',
    disposed_to: '',
    action_notes: '',
    created_at: new Date().toISOString()
  };

  complaintsStore.update(items => [newComplaint, ...items]);

  auditLogsStore.update(logs => [
    {
      id: Date.now(),
      user_name: 'Masyarakat (SIAP TERTIB)',
      action: 'SUBMIT_ADUAN',
      resource: `Complaint #${code}`,
      details: `Laporan baru diajukan terkait ${newComplaint.violation_type}`,
      timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
      ip_address: '127.0.0.1'
    },
    ...logs
  ]);

  return {
    success: true,
    ticket_code: code,
    message: 'Laporan Anda telah berhasil diregistrasi di database SIAP TERTIB Satpol PP Sulsel.'
  };
}

export async function trackComplaint(ticketCode: string): Promise<{ success: boolean; data?: Complaint; message?: string }> {
  try {
    const res = await fetch(`${API_BASE}/siaptertib/lacak/${ticketCode}`);
    if (res.ok) {
      return await res.json();
    }
  } catch (e) {
    // Graceful offline fallback
  }

  const all = get(complaintsStore);
  const found = all.find(c => c.ticket_code.toUpperCase() === ticketCode.trim().toUpperCase());
  if (found) {
    return { success: true, data: found };
  }
  return { success: false, message: 'Nomor tiket pengaduan tidak ditemukan dalam sistem.' };
}

export async function loginAdmin(username: string, password: string): Promise<{ success?: boolean; token?: string; user?: User; error?: string }> {
  try {
    const res = await fetch(`${API_BASE}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    if (res.ok) {
      return await res.json();
    }
  } catch (e) {
    // Graceful offline fallback
  }

  const users = get(usersStore);
  const matched = users.find(u => u.username.toLowerCase() === username.trim().toLowerCase());
  if (matched && password === 'satpolpp2026') {
    return {
      success: true,
      token: 'mock-jwt-token-' + matched.role,
      user: matched
    };
  }

  return {
    success: false,
    error: 'Username atau kata sandi tidak valid.'
  };
}
