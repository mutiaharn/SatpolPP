import { writable } from 'svelte/store';
import type { Article, Complaint, GalleryItem, OrganizationNode, PublicDocument, Regency, SiteSettings, User } from '$lib/types';

export interface AuditEntry {
  id: number;
  user_name: string;
  action: string;
  resource: string;
  details: string;
  timestamp: string;
  ip_address: string;
}

// 24 Master Regencies of South Sulawesi
export const defaultRegencies: Regency[] = [
  { id: '7301', name: 'Kota Makassar' },
  { id: '7302', name: 'Kota Parepare' },
  { id: '7303', name: 'Kota Palopo' },
  { id: '7304', name: 'Kabupaten Gowa' },
  { id: '7305', name: 'Kabupaten Maros' },
  { id: '7306', name: 'Kabupaten Pangkajene dan Kepulauan' },
  { id: '7307', name: 'Kabupaten Barru' },
  { id: '7308', name: 'Kabupaten Bone' },
  { id: '7309', name: 'Kabupaten Soppeng' },
  { id: '7310', name: 'Kabupaten Wajo' },
  { id: '7311', name: 'Kabupaten Sinjai' },
  { id: '7312', name: 'Kabupaten Bulukumba' },
  { id: '7313', name: 'Kabupaten Bantaeng' },
  { id: '7314', name: 'Kabupaten Jeneponto' },
  { id: '7315', name: 'Kabupaten Takalar' },
  { id: '7316', name: 'Kabupaten Kepulauan Selayar' },
  { id: '7317', name: 'Kabupaten Enrekang' },
  { id: '7318', name: 'Kabupaten Pinrang' },
  { id: '7319', name: 'Kabupaten Sidenreng Rappang' },
  { id: '7320', name: 'Kabupaten Tana Toraja' },
  { id: '7321', name: 'Kabupaten Toraja Utara' },
  { id: '7322', name: 'Kabupaten Luwu' },
  { id: '7323', name: 'Kabupaten Luwu Utara' },
  { id: '7324', name: 'Kabupaten Luwu Timur' },
];

export const defaultSettings: SiteSettings = {
  app_name: 'Satuan Polisi Pamong Praja Prov. Sulsel & SIAP TERTIB',
  agency_tagline: 'Praja Wibawa: Menegakkan Perda, Menyelenggarakan Ketenteraman dan Ketertiban Umum, serta Pelindungan Masyarakat.',
  office_address: 'Jl. Urip Sumoharjo No. 269, Kota Makassar, Sulawesi Selatan',
  call_center: '0411-453105',
  emergency_hotline: '112 / 0811-4141-112',
  email_contact: 'satpol@sulselprov.go.id',
  kasatpol_name: 'Andi Arwin Azis S.STP., M.M',
  kasatpol_quote: 'Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan senantiasa mengedepankan pendekatan humanis, persuasif, dan profesional dalam menegakkan Peraturan Daerah dan menjaga ketertiban umum demi kesejahteraan masyarakat.'
};

export const defaultArticles: Article[] = [
  {
    id: 1,
    title: 'Perbaikan Rekonstruksi Ruas Paleteang-Malimpung Batas Enrekang, Satpol PP Amankan Aset',
    slug: 'perbaikan-rekonstruksi-ruas-paleteang-malimpung',
    summary: 'Satpol PP Prov Sulsel bersinergi dalam pengamanan aset pembangunan rekonstruksi ruas jalan provinsi di Pinrang agar pengerjaan infrastruktur berjalan lancar.',
    content: '<p>Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan melaksanakan pendampingan dan penertiban pengamanan aset daerah pada proyek rekonstruksi ruas jalan Paleteang-Malimpung...</p>',
    thumbnail: '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg',
    category: 'Kegiatan',
    author_name: 'Humas Satpol PP',
    view_count: 320,
    published_at: '2026-01-12T08:08:43Z'
  },
  {
    id: 2,
    title: 'Rapat Internal Komandan Kompi dan Komandan Pleton Operasional Satpol PP Sulsel',
    slug: 'rapat-internal-danki-dan-danton-ops',
    summary: 'Konsolidasi jajaran pimpinan kompi dan pleton operasional guna meningkatkan efektivitas patroli penegakan Perda dan respon cepat aduan warga di 24 kabupaten/kota.',
    content: '<p>Kepala Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan memimpin rapat koordinasi berkala bersama seluruh Danki dan Danton operasional di markas komando...</p>',
    thumbnail: '/uploads/foto_berita/1758537465AF0666BD-2416-4D0A-9101-E1BBB12A8FA9.png',
    category: 'Kegiatan',
    author_name: 'Bidang Trantibum',
    view_count: 185,
    published_at: '2026-01-10T14:30:00Z'
  },
  {
    id: 3,
    title: 'Apel Kesiapsiagaan Patroli Ketertiban Umum Pasca Penertiban Aset Milik Daerah Pemprov',
    slug: 'apel-persiapan-patroli-pasca-penertiban-aset',
    summary: 'Pemeriksaan kesiapan personel dan armada patroli Satpol PP dalam mengawal aset tanah dan bangunan milik Pemerintah Provinsi Sulawesi Selatan.',
    content: '<p>Guna menjaga ketenteraman dan ketertiban umum pasca dilaksanakannya penertiban aset daerah di beberapa titik strategis, Satpol PP Sulsel menggelar apel siaga...</p>',
    thumbnail: '/uploads/foto_berita/175853755659BEED7D-9D9F-4A4E-A69E-DCC8EE0A2A50.png',
    category: 'Penertiban Aset',
    author_name: 'Bidang Penegakan Perda',
    view_count: 245,
    published_at: '2026-01-08T07:15:00Z'
  },
  {
    id: 4,
    title: 'Koordinasi Penertiban Ruang Milik Jalan Poros Makassar-Maros Bersama Balai Jalan',
    slug: 'koordinasi-penertiban-ruang-milik-jalan-makassar-maros',
    summary: 'Rapat koordinasi teknis penertiban pemanfaatan bahu jalan dan trotoar sepanjang koridor Trans-Sulawesi guna kelancaran arus logistik dan keselamatan pengguna jalan.',
    content: '<p>Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan bersama Balai Pelaksanaan Jalan Nasional (BPJN) menggelar rapat koordinasi intensif terkait penertiban ruang milik jalan (Rumija) di sepanjang jalur poros Makassar Menuju Maros...</p>',
    thumbnail: '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg',
    category: 'Kegiatan',
    author_name: 'Bidang Trantibum',
    view_count: 410,
    published_at: '2026-01-05T09:00:00Z'
  },
  {
    id: 5,
    title: 'Sosialisasi Peraturan Daerah Ketertiban Umum Bagi Generasi Muda dan Pelajar',
    slug: 'sosialisasi-peraturan-daerah-trantibum-pelajar',
    summary: 'Edukasi preventif mengenai pentingnya ketenteraman masyarakat dan kepatuhan terhadap norma hukum daerah kepada siswa SMA/SMK se-Sulsel.',
    content: '<p>Bidang Pembinaan Masyarakat Satpol PP Sulsel mengintensifkan program edukasi dan sosialisasi Perda Nomor 2 Tahun 2021 tentang Trantibum ke berbagai institusi pendidikan...</p>',
    thumbnail: '/uploads/foto_berita/1758537465AF0666BD-2416-4D0A-9101-E1BBB12A8FA9.png',
    category: 'Pengumuman',
    author_name: 'Bidang Binmas',
    view_count: 198,
    published_at: '2026-01-03T11:20:00Z'
  },
  {
    id: 6,
    title: 'Pengamanan Aset Tanah Pemprov Sulsel di Wilayah Somba Opu Kabupaten Gowa',
    slug: 'pengamanan-aset-tanah-pemprov-gowa',
    summary: 'Tindakan preventif pemasangan plang pengamanan aset milik Pemerintah Provinsi Sulawesi Selatan untuk mencegah penyerobotan lahan secara ilegal.',
    content: '<p>Tim Terpadu Satpol PP Prov Sulsel didampingi PPNS Penegakan Perda melaksanakan giat pemasangan plang penguasaan aset daerah di wilayah Kecamatan Somba Opu...</p>',
    thumbnail: '/uploads/foto_berita/175853755659BEED7D-9D9F-4A4E-A69E-DCC8EE0A2A50.png',
    category: 'Penertiban Aset',
    author_name: 'Bidang Penegakan Perda',
    view_count: 530,
    published_at: '2025-12-28T10:00:00Z'
  }
];

export const defaultDocuments: PublicDocument[] = [
  {
    id: 1,
    title: 'Rencana Kerja (Renja) Satpol PP Prov. Sulsel Tahun Anggaran 2026',
    doc_number: '050/112/SATPOL/2026',
    category: 'renja',
    year: 2026,
    file_url: '/uploads/dokumen/renja_2026.pdf',
    file_size: '2.4 MB'
  },
  {
    id: 2,
    title: 'Rencana Strategis (Renstra) Satpol PP Prov. Sulsel Periode 2024-2026',
    doc_number: '050/085/RENSTRA/SATPOL',
    category: 'renstra',
    year: 2024,
    file_url: '/uploads/dokumen/renstra_2024_2026.pdf',
    file_size: '4.8 MB'
  },
  {
    id: 3,
    title: 'Dokumen Pelaksanaan Anggaran (DPA) Satpol PP Prov. Sulsel TA 2025',
    doc_number: '900/DPA-SKPD/SATPOL/2025',
    category: 'dpa',
    year: 2025,
    file_url: '/uploads/dokumen/dpa_2025.pdf',
    file_size: '3.1 MB'
  },
  {
    id: 4,
    title: 'Laporan Akuntabilitas Kinerja Instansi Pemerintah (LAKIP) Tahun 2025',
    doc_number: '050/450/LAKIP/2025',
    category: 'lakip',
    year: 2025,
    file_url: '/uploads/dokumen/lakip_2025.pdf',
    file_size: '3.9 MB'
  },
  {
    id: 5,
    title: 'Laporan Keuangan Satpol PP Prov. Sulsel Audited TA 2024',
    doc_number: '900/LK-SATPOL/2024',
    category: 'keuangan',
    year: 2024,
    file_url: '/uploads/dokumen/laporan_keuangan_2024.pdf',
    file_size: '2.8 MB'
  },
  {
    id: 6,
    title: 'Peraturan Daerah Prov. Sulsel No. 2 Tahun 2021 tentang Penyelenggaraan Ketenteraman dan Ketertiban Umum',
    doc_number: 'Perda No. 2 Tahun 2021',
    category: 'perda',
    year: 2021,
    file_url: '/uploads/dokumen/perda_trantibum_2021.pdf',
    file_size: '1.8 MB'
  },
  {
    id: 7,
    title: 'Standar Operasional Prosedur (SOP) Penertiban dan Penegakan Peraturan Daerah',
    doc_number: 'SOP/01/SATPOL/2022',
    category: 'sop',
    year: 2022,
    file_url: '/uploads/dokumen/sop_penertiban_perda.pdf',
    file_size: '1.2 MB'
  }
];

// Realistic complaints illustrating the Satpol PP Sulsel SOP life-cycle
export const defaultComplaints: Complaint[] = [
  {
    id: 1,
    ticket_code: 'TRB-202611-SAMPLE',
    is_anonymous: false,
    reporter_nik: '7371101234567890',
    reporter_name: 'Ahmad Fauzan Pratama',
    reporter_phone: '081234567890',
    reporter_email: 'ahmad.fauzan@gmail.com',
    violation_type: 'Pedagang Kaki Lima (PKL) Menutup Fasilitas Umum',
    regency_id: '7301',
    regency_name: 'Kota Makassar',
    district_name: 'Panakkukang',
    village_name: 'Panaikang',
    address_detail: 'Depan Halte Bus Kampus UMI, Jl. Urip Sumoharjo',
    latitude: -5.1432,
    longitude: 119.4521,
    description: 'Terdapat sejumlah lapak PKL liar yang memakan seluruh badan trotoar pejalan kaki dan menyebabkan kemacetan parah di jam sibuk.',
    status: 'PENINDAKAN_LAPANGAN',
    disposed_to: 'Bidang Ketenteraman dan Ketertiban Umum (Pleton 2)',
    action_notes: 'Pleton 2 Patroli Trantibum telah menerbitkan Surat Peringatan (SP 1) dan saat ini melakukan pengawasan berkala di lokasi.',
    created_at: '2026-11-04T08:15:00Z'
  },
  {
    id: 2,
    ticket_code: 'TRB-202611-8821',
    is_anonymous: true,
    reporter_nik: '',
    reporter_name: 'Warga (Anonim)',
    reporter_phone: '085299887766',
    reporter_email: '',
    violation_type: 'Penertiban Reklame dan Spanduk Liar Tanpa Izin',
    regency_id: '7305',
    regency_name: 'Kabupaten Maros',
    district_name: 'Turikale',
    village_name: 'Pettuadae',
    address_detail: 'Sepanjang median jalan poros Maros-Pangkep KM 28',
    latitude: -5.0041,
    longitude: 119.5732,
    description: 'Pemasangan baliho komersial tanpa izin resmi yang dipaku di pohon perindang jalan raya, membahayakan pengendara dan merusak lingkungan.',
    status: 'DISPOSISI_BIDANG',
    disposed_to: 'Bidang Penegakan Peraturan Daerah',
    action_notes: 'Laporan telah diverifikasi keabsahannya dan didisposisikan kepada PPNS Penegakan Perda untuk koordinasi dengan Bapenda wilayah.',
    created_at: '2026-11-04T10:30:00Z'
  },
  {
    id: 3,
    ticket_code: 'TRB-202611-4412',
    is_anonymous: false,
    reporter_nik: '7306029876543210',
    reporter_name: 'Ibu Hj. Rosdiana',
    reporter_phone: '081344556677',
    reporter_email: 'rosdiana.sulsel@yahoo.com',
    violation_type: 'Penyerobotan dan Pemanfaatan Aset Daerah Pemprov',
    regency_id: '7304',
    regency_name: 'Kabupaten Gowa',
    district_name: 'Somba Opu',
    village_name: 'Sungguminasa',
    address_detail: 'Lahan eks-Balai Benih Dinas Pertanian Sulsel, Jl. Malino',
    latitude: -5.2014,
    longitude: 119.4589,
    description: 'Diduga terjadi pembangunan bangunan semi permanen ilegal di atas lahan bersertifikat milik Pemerintah Provinsi Sulawesi Selatan.',
    status: 'MENUNGGU_VERIFIKASI',
    disposed_to: '',
    action_notes: '',
    created_at: '2026-11-04T14:45:00Z'
  },
  {
    id: 4,
    ticket_code: 'TRB-202611-1092',
    is_anonymous: false,
    reporter_nik: '7371051122334455',
    reporter_name: 'Budi Santoso S.E',
    reporter_phone: '0811442233',
    reporter_email: 'budi.santoso@gmail.com',
    violation_type: 'Pelanggaran Ketertiban Tempat Hiburan dan Miras',
    regency_id: '7301',
    regency_name: 'Kota Makassar',
    district_name: 'Ujung Pandang',
    village_name: 'Losari',
    address_detail: 'Kawasan Pantai Losari, dekat anjungan Mandar',
    latitude: -5.1456,
    longitude: 119.4068,
    description: 'Aktivitas penjualan minuman keras tanpa izin pada jam operasional terlarang yang meresahkan pengunjung keluarga di fasilitas publik.',
    status: 'SELESAI',
    disposed_to: 'Bidang Penegakan Perda & Trantibum',
    action_notes: 'Operasi penertiban gabungan telah dilaksanakan. Dilakukan penyitaan 24 botol barang bukti dan pemilik menandatangani Berita Acara Pemeriksaan (BAP).',
    created_at: '2026-11-02T22:10:00Z'
  }
];

export const defaultUsers: User[] = [
  {
    id: 1,
    name: 'Tim Pengembang Diskominfo-SP Sulsel',
    username: 'superadmin',
    email: 'diskominfo@sulselprov.go.id',
    role: 'SUPERADMIN_DISKOMINFO',
    bidang: 'Aplikasi & Informatika (Diskominfo)',
    is_active: true
  },
  {
    id: 2,
    name: 'Andi Arwin Azis S.STP., M.M',
    username: 'kasatpol',
    email: 'kasatpol@sulselprov.go.id',
    role: 'ADMIN_SATPOL',
    bidang: 'Pimpinan Kasatpol PP',
    is_active: true
  },
  {
    id: 3,
    name: 'Petugas Piket Bidang Trantibum',
    username: 'stafftrantibum',
    email: 'trantibum@sulselprov.go.id',
    role: 'STAFF_SATPOL',
    bidang: 'Bidang Ketenteraman dan Ketertiban Umum',
    is_active: true
  },
  {
    id: 4,
    name: 'Penyidik PNS Bidang Penegakan Perda',
    username: 'staffperda',
    email: 'penegakanperda@sulselprov.go.id',
    role: 'STAFF_SATPOL',
    bidang: 'Bidang Penegakan Peraturan Daerah',
    is_active: true
  }
];

export const defaultAuditLogs: AuditEntry[] = [
  {
    id: 1,
    user_name: 'stafftrantibum',
    action: 'DISPOSISI_ADUAN',
    resource: 'Complaint #TRB-202611-SAMPLE',
    details: 'Mendisposisikan laporan ke Pleton 2 Patroli Trantibum',
    timestamp: '2026-11-04 09:30:12',
    ip_address: '103.151.191.73'
  },
  {
    id: 2,
    user_name: 'superadmin',
    action: 'UPDATE_CMS_SETTINGS',
    resource: 'SiteSettings',
    details: 'Memperbarui nomor Call Center dan sambutan pimpinan',
    timestamp: '2026-11-04 11:15:45',
    ip_address: '192.168.10.45'
  },
  {
    id: 3,
    user_name: 'staffperda',
    action: 'TUNTASKAN_ADUAN',
    resource: 'Complaint #TRB-202611-1092',
    details: 'Mengunggah Berita Acara Penertiban dan menutup tiket aduan',
    timestamp: '2026-11-03 14:20:00',
    ip_address: '103.151.191.73'
  }
];

// Helper to load or initialize from localStorage
function getStored<T>(key: string, fallback: T): T {
  if (typeof window !== 'undefined') {
    const raw = localStorage.getItem(key);
    if (raw) {
      try {
        return JSON.parse(raw);
      } catch (e) {
        console.error(e);
      }
    }
  }
  return fallback;
}

function saveStored<T>(key: string, value: T) {
  if (typeof window !== 'undefined') {
    localStorage.setItem(key, JSON.stringify(value));
  }
}

export const defaultGalleryItems: GalleryItem[] = [
  {
    id: 1,
    title: 'Apel Gelar Pasukan Pengamanan Wilayah Satpol PP Sulsel',
    category: 'Apel & Upacara',
    image_url: '/uploads/foto_galeri/1687082076352158666_3553275628273619_3927930681691724075_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 2,
    title: 'Pemeriksaan Barisan dan Pasukan Patroli Motor Trantibum',
    category: 'Apel & Upacara',
    image_url: '/uploads/foto_galeri/1687082083351435884_922155618893266_5055471206785735616_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 3,
    title: 'Operasi Penertiban Reklame Liar dan Fasilitas Umum',
    category: 'Penertiban',
    image_url: '/uploads/foto_galeri/1687082094352119793_1422358638552857_2632695974006539857_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 4,
    title: 'Pengamanan Aset Tanah dan Bangunan Milik Pemprov Sulsel',
    category: 'Penertiban',
    image_url: '/uploads/foto_galeri/1687082124353350967_992919668385650_5927602443603722172_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 5,
    title: 'Sosialisasi Penegakan Peraturan Daerah kepada Pelaku Usaha',
    category: 'Koordinasi',
    image_url: '/uploads/foto_galeri/1687082135353549157_140269625733302_2400965521521597112_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 6,
    title: 'Rapat Koordinasi Teknis Lintas Satpol PP Kab/Kota se-Sulsel',
    category: 'Koordinasi',
    image_url: '/uploads/foto_galeri/1687082148354520636_294019849664728_1122190031379920679_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 7,
    title: 'Patroli Terpadu Pengawasan Protokol Ketertiban Umum',
    category: 'Penertiban',
    image_url: '/uploads/foto_galeri/1687082157352404995_214459758138830_8621228371327837008_n.jpg',
    date: '18 Juni 2023'
  },
  {
    id: 8,
    title: 'HUT Satpol PP ke-74 dan Satlinmas ke-62 Tingkat Provinsi',
    category: 'Apel & Upacara',
    image_url: '/uploads/foto_galeri/1703125359Beige and Brown Modern Grunge Catalogue Cover Instagram Post.jpg',
    date: '21 Desember 2023'
  },
  {
    id: 9,
    title: 'Pemberian Piagam Penghargaan Personel Teladan Praja Wibawa',
    category: 'Apel & Upacara',
    image_url: '/uploads/foto_galeri/1703125614Snapinsta.app_409563710_18118554457326712_4118981162754069338_n_1024.jpg',
    date: '21 Desember 2023'
  },
  {
    id: 10,
    title: 'Giat Pendampingan Proyek Strategis Pemerintah Provinsi',
    category: 'Koordinasi',
    image_url: '/uploads/foto_galeri/1703126701WhatsApp Image 2023-12-21 at 10.42.28_3677c577.jpg',
    date: '21 Desember 2023'
  }
];

// Stores
export const siteSettingsStore = writable<SiteSettings>(getStored('satpol_settings', defaultSettings));
siteSettingsStore.subscribe(val => saveStored('satpol_settings', val));

export const articlesStore = writable<Article[]>(getStored('satpol_articles', defaultArticles));
articlesStore.subscribe(val => saveStored('satpol_articles', val));

export const documentsStore = writable<PublicDocument[]>(getStored('satpol_documents', defaultDocuments));
documentsStore.subscribe(val => saveStored('satpol_documents', val));

export const complaintsStore = writable<Complaint[]>(getStored('satpol_complaints', defaultComplaints));
complaintsStore.subscribe(val => saveStored('satpol_complaints', val));

export const galleryStore = writable<GalleryItem[]>(getStored('satpol_gallery', defaultGalleryItems));
galleryStore.subscribe(val => saveStored('satpol_gallery', val));

export const usersStore = writable<User[]>(getStored('satpol_users', defaultUsers));
usersStore.subscribe(val => saveStored('satpol_users', val));

export const auditLogsStore = writable<AuditEntry[]>(getStored('satpol_audit_logs', defaultAuditLogs));
auditLogsStore.subscribe(val => saveStored('satpol_audit_logs', val));


// Struktur Organisasi Dinamis (SOTK) Satpol PP Prov. Sulsel
export const defaultOrganization: OrganizationNode[] = [
  {
    id: 1,
    org_code: 'ORG-KST-01',
    parent_code: null,
    position_title: 'Kepala Satuan Polisi Pamong Praja',
    official_name: 'Andi Arwin Azis S.STP., M.M',
    rank: 'Eselon IIa',
    description: 'Memimpin perumusan kebijakan, pengawasan operasional trantibum di 24 kabupaten/kota, serta koordinasi ketenteraman wilayah bersama Forkopimda Sulawesi Selatan.',
    sub_units: [],
    photo_url: '/uploads/foto_pimpinan/pimpinan_cutout_hd.png',
    level: 1,
    sort_order: 1
  },
  {
    id: 2,
    org_code: 'ORG-SEK-01',
    parent_code: 'ORG-KST-01',
    position_title: 'Sekretariat',
    official_name: 'Sekretaris Satpol PP Sulsel',
    rank: 'Eselon IIIa',
    description: 'Mengoordinasikan urusan perencanaan program, pengelolaan keuangan anggaran, dan tata usaha kepegawaian.',
    sub_units: ['Sub Bagian Umum, Kepegawaian dan Hukum', 'Sub Bagian Program dan Pelaporan', 'Sub Bagian Keuangan dan Pengelolaan Aset'],
    level: 2,
    sort_order: 1
  },
  {
    id: 3,
    org_code: 'ORG-BID-01',
    parent_code: 'ORG-KST-01',
    position_title: 'Bidang Penegakan Peraturan Daerah',
    official_name: 'Kabid Penegakan Perda',
    rank: 'Eselon IIIb',
    description: 'Memimpin penindakan yustisial dan non-yustisial serta penyidikan pelanggaran peraturan daerah oleh PPNS.',
    sub_units: ['Seksi Pengawasan dan Pembinaan Penegakan Perda', 'Seksi Penindakan, Penyidikan dan PPNS Daerah'],
    level: 2,
    sort_order: 2
  },
  {
    id: 4,
    org_code: 'ORG-BID-02',
    parent_code: 'ORG-KST-01',
    position_title: 'Bidang Ketenteraman dan Ketertiban Umum',
    official_name: 'Kabid Trantibum',
    rank: 'Eselon IIIb',
    description: 'Mengendalikan operasi patroli ketertiban umum, penertiban fasilitas publik, dan pengawalan pejabat negara.',
    sub_units: ['Seksi Operasi dan Pengendalian Ketertiban', 'Seksi Pengamanan, Pengawalan dan Patroli Wilayah'],
    level: 2,
    sort_order: 3
  },
  {
    id: 5,
    org_code: 'ORG-BID-03',
    parent_code: 'ORG-KST-01',
    position_title: 'Bidang Pelindungan Masyarakat dan Damkar',
    official_name: 'Kabid Linmas & Damkar',
    rank: 'Eselon IIIb',
    description: 'Menyelenggarakan pembinaan teknis anggota Satlinmas dan koordinasi penanganan kedaruratan kebakaran daerah.',
    sub_units: ['Seksi Satuan Pelindungan Masyarakat (Satlinmas)', 'Seksi Kesiapsiagaan, Kebakaran dan Bencana'],
    level: 2,
    sort_order: 4
  }
];

export const organizationStore = writable<OrganizationNode[]>(getStored('satpol_organization', defaultOrganization));
organizationStore.subscribe(val => saveStored('satpol_organization', val));

/** Helper pencatatan audit trail agar tidak duplikasi kode di setiap halaman */
export function logAudit(userName: string, action: string, resource: string, details: string) {
  auditLogsStore.update(logs => [
    {
      id: Date.now() + Math.floor(Math.random() * 1000),
      user_name: userName || 'sistem',
      action,
      resource,
      details,
      timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
      ip_address: '127.0.0.1'
    },
    ...logs
  ]);
}
