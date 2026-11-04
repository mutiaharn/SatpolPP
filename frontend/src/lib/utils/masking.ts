// Utility functions for masking sensitive citizen personal data (NIK, Nama, Phone)
// Used in SIAP TERTIB anonymous reporting and public tracking

export function maskNIK(nik: string): string {
  if (!nik || nik.length < 6) return '****************';
  // Keep first 4 and last 2 digits: 7371**********12
  const first = nik.slice(0, 4);
  const last = nik.slice(-2);
  const stars = '*'.repeat(Math.max(4, nik.length - 6));
  return `${first}${stars}${last}`;
}

export function maskName(name: string): string {
  if (!name || name.trim().length <= 2) return 'Warga (Anonim)';
  const parts = name.trim().split(/\s+/);
  return parts.map(part => {
    if (part.length <= 2) return part[0] + '*';
    return part[0] + '*'.repeat(part.length - 1);
  }).join(' ') + ' (Anonim Terlindungi)';
}

export function maskPhone(phone: string): string {
  if (!phone || phone.length < 7) return '08**********';
  // Keep first 4 and last 3 digits: 0812****789
  const first = phone.slice(0, 4);
  const last = phone.slice(-3);
  const stars = '*'.repeat(Math.max(3, phone.length - 7));
  return `${first}${stars}${last}`;
}
