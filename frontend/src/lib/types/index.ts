export interface User {
  id: number;
  name: string;
  username: string;
  email: string;
  role: 'SUPERADMIN_DISKOMINFO' | 'ADMIN_SATPOL' | 'STAFF_SATPOL' | 'PUBLIC';
  bidang?: string;
  is_active: boolean;
}

export interface Article {
  id: number;
  title: string;
  slug: string;
  summary: string;
  content: string;
  thumbnail: string;
  category: string;
  author_name: string;
  view_count: number;
  published_at: string;
}

export interface PublicDocument {
  id: number;
  title: string;
  doc_number: string;
  category: 'renja' | 'renstra' | 'dpa' | 'lakip' | 'keuangan' | 'rpjmd' | 'perda' | 'sop';
  year: number;
  file_url: string;
  file_size: string;
}

export interface Complaint {
  id: number;
  ticket_code: string;
  is_anonymous: boolean;
  reporter_nik: string;
  reporter_name: string;
  reporter_phone: string;
  reporter_email: string;
  masked_nik?: string;
  masked_name?: string;
  masked_phone?: string;
  violation_type: string;
  regency_id: string;
  regency_name: string;
  district_name: string;
  village_name: string;
  address_detail: string;
  latitude: number;
  longitude: number;
  description: string;
  evidence_photos?: string;
  status: 'MENUNGGU_VERIFIKASI' | 'DISPOSISI_BIDANG' | 'PENINDAKAN_LAPANGAN' | 'SELESAI' | 'DITOLAK';
  disposed_to?: string;
  action_notes?: string;
  created_at: string;
}

export interface Regency {
  id: string;
  name: string;
}

export interface SiteSettings {
  app_name: string;
  agency_tagline: string;
  office_address: string;
  call_center: string;
  emergency_hotline: string;
  email_contact: string;
  kasatpol_name: string;
  kasatpol_quote: string;
}

export interface GalleryItem {
  id: number;
  title: string;
  category: string;
  image_url: string;
  date: string;
}
