<script lang="ts">
  import { page } from '$app/stores';

  interface BidangInfo {
    title: string;
    lead: string;
    tasks: string[];
    sections: { name: string; desc: string }[];
  }

  const bidangData: Record<string, BidangInfo> = {
    'sekretariat': {
      title: 'Sekretariat Satpol PP Sulsel',
      lead: 'Sekretariat mempunyai tugas memberikan pelayanan teknis dan administratif kepada seluruh unsur di lingkungan Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan.',
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
      lead: 'Bidang Penegakan Peraturan Daerah mempunyai tugas menyelenggarakan penegakan, pengawasan, dan penyidikan atas pelanggaran Peraturan Daerah dan Peraturan Kepala Daerah Provinsi Sulawesi Selatan.',
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
      lead: 'Bidang Ketenteraman dan Ketertiban Umum bertugas merencanakan, mengendalikan, dan melaksanakan operasi patroli, pengamanan aset milik daerah, serta penertiban fasilitas umum di wilayah Provinsi Sulawesi Selatan.',
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
      lead: 'Bidang Pembinaan Masyarakat mempunyai tugas mengedukasi publik, membina kesadaran hukum masyarakat, serta menyelenggarakan bimbingan teknis bagi personel Satpol PP di kabupaten/kota.',
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
      lead: 'Bidang Linmas dan Damkar bertugas mengoordinasikan pembinaan Satuan Pelindungan Masyarakat (Satlinmas) serta kesiapsiagaan pengendalian bahaya kebakaran dan kedaruratan daerah.',
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

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-6 px-4">
  <div class="max-w-6xl mx-auto flex flex-wrap items-center gap-2 text-[11px] text-[#6f7775]">
    <a href="/" class="hover:text-[#F08300]">Beranda</a>
    <span>/</span>
    <span class="hover:text-[#F08300]">Bidang Kerja</span>
    <span>/</span>
    <span class="text-[#101116] font-semibold">{current.title}</span>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto">
  <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
    <!-- Sidebar Navigation -->
    <div class="lg:col-span-1 space-y-4">
      <div class="bg-white border border-[#eee9db] shadow-xs p-4 rounded-2xl">
        <h3 class="text-xs font-bold text-[#101116] uppercase tracking-wider border-b border-gray-100 pb-2 mb-3">
          Bidang Kerja
        </h3>
        <nav class="space-y-1">
          {#each bidangMenu as b}
            <a 
              href={`/bidang/${b.slug}`}
              class={`block px-3.5 py-2.5 text-xs font-bold transition rounded-xl ${slug === b.slug ? 'bg-[#F08300] text-white shadow-xs' : 'text-[#2b2f3f] hover:bg-gray-100'}`}
            >
              {b.label}
            </a>
          {/each}
        </nav>
      </div>

      <!-- Quick Action Card -->
      <div class="bg-[#3d261a] text-white p-5 rounded-2xl border-t-4 border-[#F08300] shadow-xs">
        <div class="text-[10px] text-[#F08300] font-bold uppercase tracking-wider mb-1">SIAP TERTIB</div>
        <div class="font-bold text-xs">Aduan Pelanggaran Wilayah</div>
        <p class="text-[11px] text-gray-300 mt-2 leading-relaxed">
          Laporan masyarakat yang masuk melalui portal akan didisposisikan langsung ke bidang teknis terkait untuk ditindaklanjuti.
        </p>
        <a 
          href="/lapor" 
          class="mt-3 block text-center bg-[#F08300] text-white font-bold text-xs uppercase py-2.5 rounded-xl hover:bg-[#d97400] transition shadow-xs"
        >
          Kirim Laporan
        </a>
      </div>
    </div>

    <!-- Main Content Column -->
    <div class="lg:col-span-3 bg-white border border-[#eee9db] shadow-xs p-6 md:p-8 space-y-6 rounded-2xl">
      <div class="border-b border-gray-100 pb-4">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Unit Kerja Satpol PP Sulsel</div>
        <h1 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">{current.title}</h1>
        <p class="text-xs text-[#2b2f3f] font-semibold mt-2 leading-relaxed">{current.lead}</p>
      </div>

      <!-- Tasks List -->
      <div class="space-y-3">
        <h3 class="text-xs font-bold uppercase tracking-wider text-[#101116]">Uraian Tugas dan Tanggung Jawab:</h3>
        <ul class="space-y-2 text-xs text-[#2b2f3f] leading-relaxed">
          {#each current.tasks as task}
            <li class="flex items-start gap-2.5">
              <span class="inline-block w-1.5 h-1.5 bg-[#F08300] mt-1.5 flex-shrink-0"></span>
              <span>{task}</span>
            </li>
          {/each}
        </ul>
      </div>

      <!-- Sub-sections List -->
      <div class="pt-4 border-t border-gray-100 space-y-3">
        <h3 class="text-xs font-bold uppercase tracking-wider text-[#101116]">Seksi &amp; Sub-Bagian di Bawah Bidang:</h3>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          {#each current.sections as sec}
            <div class="border border-gray-200 p-4 bg-gray-50/50 rounded-xl">
              <div class="font-bold text-xs text-[#101116] uppercase">{sec.name}</div>
              <p class="text-[11px] text-[#6f7775] mt-1 leading-relaxed">{sec.desc}</p>
            </div>
          {/each}
        </div>
      </div>
    </div>
  </div>
</div>
