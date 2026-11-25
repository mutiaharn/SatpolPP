<script lang="ts">
  import { page } from '$app/stores';
  import { ChevronRight, Shield, CheckCircle2, ArrowRight } from 'lucide-svelte';

  interface BidangInfo {
    title: string;
    lead: string;
    tasks: string[];
    sections: { name: string; desc: string }[];
  }

  const bidangData: Record<string, BidangInfo> = {
    'sekretariat': {
      title: 'Sekretariat Satpol PP Sulsel',
      lead: 'Pusat koordinasi perencanaan, tata kelola keuangan, kepegawaian aparatur, dan pelayanan administrasi terpadu.',
      tasks: [
        'Pengkoordinasian penyusunan rencana program, anggaran, dan pelaporan akuntabilitas kinerja instansi.',
        'Pengelolaan urusan kepegawaian, pengembangan kompetensi personel, dan tata usaha persuratan.',
        'Pengelolaan keuangan dinas, verifikasi pembukuan, dan penatausahaan aset milik daerah.',
        'Pelaksanaan urusan rumah tangga, keprotokolan, dan pendampingan advokasi hukum.'
      ],
      sections: [
        { name: 'Sub Bagian Umum, Kepegawaian & Hukum', desc: 'Menyelenggarakan urusan persuratan dinas, arsip, formasi jabatan, disiplin aparatur, serta telaah hukum regulasi daerah.' },
        { name: 'Sub Bagian Program & Pelaporan', desc: 'Menyusun dokumen perencanaan strategis (Renstra), rencana kerja tahunan (Renja), serta pelaporan kinerja (LAKIP).' },
        { name: 'Sub Bagian Keuangan & Pengelolaan Aset', desc: 'Melaksanakan penatausahaan belanja langsung/tidak langsung, laporan keuangan audited, dan inventarisasi sarana prasarana dinas.' }
      ]
    },
    'perda': {
      title: 'Bidang Penegakan Peraturan Daerah (Perda)',
      lead: 'Mengawal penegakan Perda dan Peraturan Gubernur secara adil, terukur, dan persuasif demi menciptakan kepatuhan hukum yang berkeadilan.',
      tasks: [
        'Penyusunan rencana teknis operasional penegakan dan pengawasan kepatuhan terhadap Perda dan Perkada.',
        'Pemeriksaan dan penyelidikan awal atas laporan dugaan pelanggaran Perda oleh masyarakat atau badan usaha.',
        'Pelaksanaan tindakan non-yustisial secara persuasif, peringatan tertulis, hingga penyegelan sementara tempat usaha ilegal.',
        'Penyidikan tindak pidana pelanggaran Perda oleh Penyidik Pegawai Negeri Sipil (PPNS) untuk diajukan ke sidang tindak pidana ringan (Tipiring).'
      ],
      sections: [
        { name: 'Seksi Pengawasan & Pembinaan', desc: 'Melakukan monitoring kepatuhan pelaku usaha terhadap perizinan daerah, retribusi, dan tata ruang provinsi.' },
        { name: 'Seksi Penindakan, Penyidikan & PPNS', desc: 'Memeriksa berkas perkara, memanggil pihak terlapor, menyusun Berita Acara Pemeriksaan (BAP), dan mengoordinasikan persidangan yustisi.' }
      ]
    },
    'trantibum': {
      title: 'Bidang Ketenteraman & Ketertiban Umum',
      lead: 'Menyelenggarakan kesiapsiagaan patroli 24 jam, pengamanan fasilitas publik, dan penertiban ruang kota agar aman dan tenteram bagi seluruh warga.',
      tasks: [
        'Pelaksanaan patroli ketertiban rutin dan terpadu pada titik rawan kerumunan dan jalur protokol provinsi.',
        'Penertiban pedagang liar (PKL), spanduk tanpa izin, dan bangunan semi permanen yang mengokupasi ruang publik.',
        'Pengamanan fisik aset tanah, gedung perkantoran, dan fasilitas umum milik Pemerintah Provinsi Sulawesi Selatan.',
        'Pengawalan kehormatan bagi pimpinan daerah (Gubernur, Wakil Gubernur) dan tamu negara (VVIP/VIP).'
      ],
      sections: [
        { name: 'Seksi Operasi & Pengendalian Ketertiban', desc: 'Mengatur pengerahan kompi dan pleton operasional lapangan, pemetaan zona rawan pelanggaran, dan patroli bermotor.' },
        { name: 'Seksi Pengamanan, Pengawalan & Pengawasan', desc: 'Menyelenggarakan protokol pengawalan pimpinan dan pengamanan ketat terhadap obyek vital milik Pemprov Sulsel.' }
      ]
    },
    'binmas': {
      title: 'Bidang Pembinaan Masyarakat & Kapasitas',
      lead: 'Membangun dialog kemitraan, penyuluhan kesadaran hukum warga, dan pembinaan partisipasi publik dalam memelihara ketertiban.',
      tasks: [
        'Sosialisasi materi Peraturan Daerah dan etika ketertiban umum ke lingkungan sekolah, kampus, dan komunitas.',
        'Pelaksanaan pembinaan mental, fisik, dan disiplin (samapta) anggota Satpol PP Sulawesi Selatan.',
        'Penyelenggaraan forum koordinasi dan bimbingan teknis bersama Satpol PP dari 24 Kabupaten/Kota se-Sulsel.',
        'Mendorong partisipasi aktif warga dalam menjaga lingkungan tertib dan aman melalui kanal SIAP TERTIB.'
      ],
      sections: [
        { name: 'Seksi Pembinaan Hubungan Antar Lembaga', desc: 'Membangun jejaring sinergi dengan instansi vertikal, TNI, Polri, ormas, dan tokoh masyarakat.' },
        { name: 'Seksi Peningkatan Kapasitas Aparatur', desc: 'Mengorganisir diklat teknis PPNS, pelatihan bela diri taktis, dan standarisasi SOP penindakan humanis.' }
      ]
    },
    'linmas-damkar': {
      title: 'Bidang Pelindungan Masyarakat & Penyelamatan Kebakaran',
      lead: 'Mengkoordinasikan satuan perlindungan masyarakat (Satlinmas) dan kesiapsiagaan penanggulangan gangguan ketertiban.',
      tasks: [
        'Pendataan, registrasi, dan pembinaan kelembagaan Satlinmas di 24 Kabupaten/Kota se-Sulawesi Selatan.',
        'Pengerahan anggota Satlinmas dalam pengamanan agenda nasional, pemilihan umum (Pemilu/Pilkada), dan penanggulangan bencana.',
        'Koordinasi lintas sektor pencegahan dan penanggulangan kebakaran hutan dan lahan (Karhutla) serta pemukiman.',
        'Pelatihan keterampilan dasar pertolongan pertama dan evakuasi korban bencana bagi relawan linmas.'
      ],
      sections: [
        { name: 'Seksi Pelindungan Masyarakat (Satlinmas)', desc: 'Mengelola database anggota Linmas tingkat desa/kelurahan, seragam kedinasan, dan penugasan perbantuan Kamtibmas.' },
        { name: 'Seksi Kesiapsiagaan Kebakaran & Penyelamatan', desc: 'Mengoordinasikan armada damkar kabupaten/kota, inspeksi alat proteksi kebakaran di gedung pemerintah, dan mitigasi darurat.' }
      ]
    }
  };

  $: slug = $page.params.slug || 'trantibum';
  $: current = bidangData[slug] || bidangData['trantibum'];

  const bidangMenu = [
    { slug: 'sekretariat', label: 'Sekretariat' },
    { slug: 'perda', label: 'Bidang Penegakan Perda' },
    { slug: 'trantibum', label: 'Bidang Trantibum' },
    { slug: 'binmas', label: 'Bidang Pembinaan Masyarakat' },
    { slug: 'linmas-damkar', label: 'Bidang Linmas & Damkar' }
  ];
</script>

<svelte:head>
  <title>{current.title} | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#F08300]/5 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <!-- Breadcrumb Pills -->
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-slate-400">Bidang Kerja</span>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">{current.title}</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">{current.title}</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      {current.lead}
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto">
  <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
    <!-- Sidebar Navigation -->
    <div class="lg:col-span-1 space-y-5">
      <div class="ds-glass border border-white/90 p-4 rounded-3xl shadow-glass ds-specular-top">
        <h3 class="text-xs font-bold text-slate-900 uppercase tracking-wider border-b border-white/80 pb-3 mb-3">
          Struktur Bidang
        </h3>
        <nav class="space-y-2">
          {#each bidangMenu as b}
            <a 
              href={`/bidang/${b.slug}`}
              class={`block px-4 py-2.5 text-xs font-bold transition duration-200 rounded-xl ds-press ${
                slug === b.slug 
                  ? 'ds-pill-active' 
                  : 'ds-pill-inactive'
              }`}
            >
              {b.label}
            </a>
          {/each}
        </nav>
      </div>

      <!-- Quick Action SIAP TERTIB Card -->
      <div class="ds-glass-dark text-white p-5 rounded-3xl border border-white/15 shadow-glass ds-specular-dark">
        <div class="text-[10px] text-[#F08300] font-bold uppercase tracking-wider mb-1">SIAP TERTIB</div>
        <div class="font-bold text-xs text-white">Aduan Pelanggaran Wilayah</div>
        <p class="text-[11px] text-slate-300 mt-2 leading-relaxed font-light">
          Laporan masyarakat yang masuk melalui portal akan didisposisikan langsung ke bidang teknis terkait untuk ditindaklanjuti secara terpadu.
        </p>
        <a 
          href="/lapor" 
          class="mt-3.5 inline-flex items-center justify-center gap-2 w-full bg-gradient-to-r from-[#F08300] to-[#e07000] hover:from-[#d97400] hover:to-[#c66500] text-white font-bold text-xs uppercase py-2.5 rounded-xl transition shadow-glass-gold ds-press"
        >
          <span>Kirim Laporan</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
      </div>
    </div>

    <!-- Main Content Column -->
    <div class="lg:col-span-3 ds-glass-strong border border-white/90 p-6 md:p-8 space-y-6 rounded-3xl shadow-glass ds-specular-top">
      <div class="border-b border-white/80 pb-4">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Unit Kerja Satpol PP Sulsel</div>
        <h2 class="text-2xl font-extrabold text-slate-900 tracking-tight mt-1">{current.title}</h2>
        <p class="text-xs sm:text-sm text-slate-600 mt-2 leading-relaxed font-light">{current.lead}</p>
      </div>

      <!-- Tasks List -->
      <div class="space-y-4">
        <h3 class="text-xs font-bold uppercase tracking-wider text-slate-900">
          Uraian Tugas dan Tanggung Jawab:
        </h3>
        <ul class="space-y-3 text-xs sm:text-sm text-slate-700 leading-relaxed">
          {#each current.tasks as task}
            <li class="flex items-start gap-3 bg-slate-50/70 border border-slate-200/80 p-3.5 rounded-xl">
              <CheckCircle2 class="w-4 h-4 text-[#F08300] mt-0.5 shrink-0" />
              <span>{task}</span>
            </li>
          {/each}
        </ul>
      </div>

      <!-- Sub-sections List -->
      <div class="pt-6 border-t border-slate-200/70 space-y-4">
        <h3 class="text-xs font-bold uppercase tracking-wider text-slate-900">
          Seksi &amp; Sub-Bagian di Bawah Bidang:
        </h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          {#each current.sections as sec}
            <div class="ds-glass border border-slate-200/80 p-5 rounded-2xl hover:border-[#F08300] transition shadow-xs">
              <div class="font-bold text-xs text-slate-900 uppercase">
                {sec.name}
              </div>
              <p class="text-xs text-slate-600 mt-2 leading-relaxed font-light">{sec.desc}</p>
            </div>
          {/each}
        </div>
      </div>
    </div>
  </div>
</div>
