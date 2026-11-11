// Utility functions for masking sensitive citizen personal data (NIK, Nama, Phone, Email)
// Used in SIAP TERTIB public tracking (identitas pelapor selalu disamarkan di tampilan publik)
// and in the officer dashboard before "Buka Masking" is requested.

import type { Complaint } from '$lib/types';

export function maskNIK(nik: string): string {
  const clean = (nik || '').trim();
  if (clean.length < 6) return '****************';
  // Keep first 4 and last 2 digits: 7371**********12
  const first = clean.slice(0, 4);
  const last = clean.slice(-2);
  const stars = '*'.repeat(Math.max(4, clean.length - 6));
  return `${first}${stars}${last}`;
}

export function maskName(name: string): string {
  const clean = (name || '').trim();
  if (clean.length <= 1) return 'Warga Pelapor';
  return clean
    .split(/\s+/)
    .map(part => (part.length <= 2 ? part[0] + '*' : part.slice(0, 2) + '*'.repeat(Math.min(part.length - 2, 5))))
    .join(' ');
}

export function maskPhone(phone: string): string {
  const clean = (phone || '').trim();
  if (clean.length < 7) return '08**********';
  // Keep first 4 and last 3 digits: 0812****789
  const first = clean.slice(0, 4);
  const last = clean.slice(-3);
  const stars = '*'.repeat(Math.max(3, clean.length - 7));
  return `${first}${stars}${last}`;
}

export function maskEmail(email: string): string {
  const clean = (email || '').trim();
  const at = clean.indexOf('@');
  if (at < 1) return '-';
  const local = clean.slice(0, at);
  const domain = clean.slice(at + 1);
  const visible = local.slice(0, Math.min(2, local.length));
  return `${visible}${'*'.repeat(Math.max(3, local.length - visible.length))}@${domain}`;
}

export interface MaskedIdentity {
  name: string;
  nik: string;
  phone: string;
  email: string;
}

/** Selalu menghasilkan identitas tersamarkan, apa pun status anonim pelapor. */
export function maskComplaintIdentity(c: Pick<Complaint, 'reporter_name' | 'reporter_nik' | 'reporter_phone' | 'reporter_email' | 'is_anonymous'>): MaskedIdentity {
  return {
    name: c.is_anonymous ? 'Warga (Anonim Terlindungi)' : maskName(c.reporter_name),
    nik: c.reporter_nik ? maskNIK(c.reporter_nik) : '-',
    phone: c.reporter_phone ? maskPhone(c.reporter_phone) : '-',
    email: c.reporter_email ? maskEmail(c.reporter_email) : '-'
  };
}

export function isValidEmail(email: string): boolean {
  return /^[^\s@]+@[^\s@]+\.[^\s@]{2,}$/.test((email || '').trim());
}
