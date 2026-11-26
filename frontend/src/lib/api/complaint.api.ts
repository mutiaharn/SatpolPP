import { get } from 'svelte/store';
import type { Complaint, ComplaintStatus, EmailNotice } from '$lib/types';
import { complaintsStore, logAudit } from '$lib/stores/appState';
import { STATUS_LABELS } from '$lib/config/constants';
import { maskComplaintIdentity } from '$lib/utils/masking';
import { API_BASE, authHeaders, tryFetch } from './http';

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

export async function uploadEvidence(file: File): Promise<string | null> {
  try {
    const form = new FormData();
    form.append('foto', file);
    const res = await fetch(`${API_BASE}/upload/evidence`, { method: 'POST', body: form });
    if (res.ok) {
      const data = await res.json();
      return data.image_url || null;
    }
  } catch {
    // offline fallback
  }
  return null;
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

export async function submitComplaint(payload: Partial<Complaint>): Promise<SubmitComplaintResult> {
  const remote = await tryFetch<SubmitComplaintResult>(`${API_BASE}/siaptertib/lapor`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload)
  });

  if (remote?.success) {
    const stored = remote.data ?? buildLocalComplaint(payload, remote.ticket_code);
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
      subject: buildEmailSubject(ticket, STATUS_LABELS[update.status]),
      message: reporterEmail ? 'Backend belum berjalan, notifikasi email disimulasikan.' : 'Pelapor tidak mencantumkan email.'
    }
  };
}

export async function unmaskComplaintIdentity(ticket: string, reason: string) {
  return tryFetch<{
    success: boolean;
    data?: {
      ticket_code: string;
      reporter_name: string;
      reporter_nik: string;
      reporter_phone: string;
      reporter_email: string;
    };
  }>(`${API_BASE}/admin/complaints/${encodeURIComponent(ticket)}/unmask`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify({ reason })
  });
}
