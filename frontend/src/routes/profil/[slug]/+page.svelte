<script lang="ts">
  import { page } from '$app/stores';

  interface ProfileItem {
    title: string;
    subtitle: string;
    updated: string;
    contentHtml: string;
  }

  const profileData: Record<string, ProfileItem> = {
    'sejarah': {
      title: 'Sejarah Singkat Satpol PP',
      subtitle: 'Napak Tilas Pembentukan Korps Praja Wibawa Penjaga Ketenteraman Masyarakat',
      updated: '18 Juni 2023',
      contentHtml: `
        <p class="leading-relaxed"><strong>Satuan Polisi Pamong Praja</strong> dibentuk berawal di Yogyakarta pada tanggal <strong>3 Maret 1950</strong>, dengan semboyan kehormatan <strong>PRAJA WIBAWA</strong>, untuk mewadahi penyelenggaraan sebagian kewenangan pemerintah daerah dalam memelihara ketertiban.</p>
        <p class="leading-relaxed">Sebelum proklamasi kemerdekaan, fungsi penjagaan keamanan ketertiban telah dilaksanakan pemerintah. Pasca proklamasi di tengah dinamika awal berdirinya NKRI, dibentuklah Detasemen Polisi sebagai Penjaga Keamanan Kapanewon di Yogyakarta berdasarkan Surat Perintah Jawatan Praja Daerah Istimewa Yogyakarta.</p>
        <p class="leading-relaxed">Seiring perkembangan tata kelola ketatanegaraan Indonesia, keberadaan Polisi Pamong Praja dikukuhkan melalui berbagai regulasi nasional hingga UU No. 23 Tahun 2014 tentang Pemerintahan Daerah dan PP No. 16 Tahun 2018. Di Sulawesi Selatan, Satpol PP Prov. Sulsel memegang peranan strategis sebagai komando koordinasi penegakan ketertiban umum lintas 24 Kabupaten dan Kota.</p>
      `
    },
    'visi-misi': {
      title: 'Visi dan Misi Satpol PP Sulsel',
      subtitle: 'Komitmen Menuju Satpol PP yang Profesional, Berwibawa, Humanis, dan Berintegritas',
      updated: '31 Oktober 2023',
      contentHtml: `
        <div class="bg-[#f9f4e8] border-l-4 border-[#F08300] p-6 mb-6">
          <div class="text-xs font-bold text-[#F08300] uppercase tracking-wider mb-1">Visi Utama</div>
          <blockquote class="text-base md:text-lg font-extrabold text-[#101116] italic leading-snug">
            &ldquo;Terwujudnya Polisi Pamong Praja yang Profesional dan Berwibawa dalam pelaksanaan tugas, menjadi pengayom dan pelayan masyarakat, serta Penegak Perda yang tangguh dan mumpuni di Sulawesi Selatan.&rdquo;
          </blockquote>
        </div>

        <h3 class="text-sm font-bold uppercase text-[#101116] tracking-wider mb-3">Misi Pembangunan</h3>
        <ol class="list-decimal pl-5 space-y-3 text-xs md:text-sm text-[#2b2f3f] leading-relaxed">
          <li><strong>Meningkatkan profesionalisme dan kapasitas sumber daya aparatur</strong> Satuan Polisi Pamong Praja yang berintegritas tinggi serta berdisiplin tinggi.</li>
          <li><strong>Menegakkan Peraturan Daerah dan Peraturan Kepala Daerah</strong> secara tegas, berkeadilan, terukur, dan mengedepankan pendekatan humanis-persuasif.</li>
          <li><strong>Mewujudkan situasi ketenteraman dan ketertiban umum</strong> yang kondusif di seluruh wilayah Provinsi Sulawesi Selatan guna mendukung stabilitas pembangunan dan investasi daerah.</li>
          <li><strong>Mengoptimalkan pelindungan masyarakat (Linmas)</strong> serta kesiapsiagaan penanggulangan gangguan ketertiban dan bencana bersama unsur stakeholder terkait.</li>
          <li><strong>Meningkatkan sinergi kelembagaan dan koordinasi operasional</strong> dengan Satpol PP Kabupaten/Kota se-Sulawesi Selatan dan jajaran Forkopimda.</li>
        </ol>
      `
    },
    'struktur': {
      title: 'Struktur Organisasi',
      subtitle: 'Bagan Susunan Organisasi dan Tata Kerja Satpol PP Provinsi Sulawesi Selatan',
      updated: '18 Juni 2023',
      contentHtml: `
        <p class="leading-relaxed mb-6">Berdasarkan Peraturan Gubernur Sulawesi Selatan tentang Kedudukan, Susunan Organisasi, Tugas dan Fungsi serta Tata Kerja Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan, struktur organisasi dipimpin oleh seorang Kepala Satuan setingkat Eselon IIa.</p>

        <div class="border border-[#eee9db] p-6 bg-white space-y-4">
          <div class="text-center pb-4 border-b border-gray-100">
            <div class="font-extrabold text-sm text-[#101116] uppercase">Kepala Satuan Polisi Pamong Praja</div>
            <div class="text-xs text-[#F08300] font-bold">Andi Arwin Azis S.STP., M.M</div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 pt-2">
            <div class="border border-gray-200 p-3 bg-gray-50/50">
              <div class="font-bold text-xs text-[#101116] uppercase">Sekretariat</div>
              <ul class="text-[11px] text-[#6f7775] mt-1 space-y-0.5 list-disc pl-4">
                <li>Sub Bagian Umum, Kepegawaian dan Hukum</li>
                <li>Sub Bagian Program dan Pelaporan</li>
                <li>Sub Bagian Keuangan dan Pengelolaan Aset</li>
              </ul>
            </div>

            <div class="border border-gray-200 p-3 bg-gray-50/50">
              <div class="font-bold text-xs text-[#101116] uppercase">Bidang Penegakan Peraturan Daerah</div>
              <ul class="text-[11px] text-[#6f7775] mt-1 space-y-0.5 list-disc pl-4">
                <li>Seksi Pengawasan dan Pembinaan Penegakan Perda</li>
                <li>Seksi Penindakan, Penyidikan dan PPNS Daerah</li>
              </ul>
            </div>

            <div class="border border-gray-200 p-3 bg-gray-50/50">
              <div class="font-bold text-xs text-[#101116] uppercase">Bidang Ketenteraman &amp; Ketertiban Umum</div>
              <ul class="text-[11px] text-[#6f7775] mt-1 space-y-0.5 list-disc pl-4">
                <li>Seksi Operasi dan Pengendalian Ketertiban</li>
                <li>Seksi Pengamanan, Pengawalan dan Patroli Wilayah</li>
              </ul>
            </div>

            <div class="border border-gray-200 p-3 bg-gray-50/50">
              <div class="font-bold text-xs text-[#101116] uppercase">Bidang Pelindungan Masyarakat &amp; Damkar</div>
              <ul class="text-[11px] text-[#6f7775] mt-1 space-y-0.5 list-disc pl-4">
                <li>Seksi Satuan Pelindungan Masyarakat (Satlinmas)</li>
                <li>Seksi Kesiapsiagaan, Kebakaran dan Bencana</li>
              </ul>
            </div>
          </div>
        </div>
      `
    },
    'tupoksi': {
      title: 'Tugas Pokok dan Fungsi',
      subtitle: 'Landasan Yuridis dan Mandat Penyelenggaraan Ketertiban Berdasarkan PP No. 16/2018',
      updated: '18 Juni 2023',
      contentHtml: `
        <p class="leading-relaxed mb-4">Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan merupakan Perangkat Daerah yang mempunyai Tugas Pokok dan Fungsi serta Kewenangan sebagaimana diatur dalam Pasal 255, 256 dan 257 Undang-Undang Nomor 23 Tahun 2014 tentang Pemerintahan Daerah jo. Peraturan Pemerintah Nomor 16 Tahun 2018.</p>

        <div class="space-y-4">
          <div class="bg-gray-50 border border-gray-200 p-5">
            <h4 class="font-bold text-xs uppercase text-[#101116] mb-2 tracking-wider">Tugas Pokok</h4>
            <p class="text-xs text-[#2b2f3f] leading-relaxed">
              Menyelenggarakan urusan pemerintahan bidang ketenteraman dan ketertiban umum serta pelindungan masyarakat, menegakkan Peraturan Daerah (Perda) dan Peraturan Kepala Daerah (Perkada).
            </p>
          </div>

          <div class="bg-gray-50 border border-gray-200 p-5">
            <h4 class="font-bold text-xs uppercase text-[#101116] mb-2 tracking-wider">Fungsi Pelaksanaan</h4>
            <ol class="list-decimal pl-5 space-y-2 text-xs text-[#2b2f3f] leading-relaxed">
              <li>Penyusunan program dan pelaksanaan penegakan Perda dan Perkada, penyelenggaraan Trantibum serta Linmas di tingkat provinsi.</li>
              <li>Pelaksanaan kebijakan penegakan Perda dan Perkada di wilayah Provinsi Sulawesi Selatan.</li>
              <li>Pelaksanaan koordinasi pemeliharaan dan penyelenggaraan ketenteraman dan ketertiban umum bersama instansi kepolisian, TNI, kejaksaan, dan Satpol PP Kab/Kota.</li>
              <li>Pengawasan terhadap masyarakat, aparatur sipil negara, atau badan hukum agar mematuhi dan menaati Perda dan Peraturan Gubernur.</li>
              <li>Pemberdayaan dan pembinaan teknis Satuan Pelindungan Masyarakat (Satlinmas).</li>
            </ol>
          </div>
        </div>
      `
    },
    'pejabat': {
      title: 'Profil Pejabat Struktural',
      subtitle: 'Pimpinan dan Pejabat Administrator di Lingkungan Satpol PP Prov. Sulsel',
      updated: '18 Juni 2023',
      contentHtml: `
        <div class="space-y-6">
          <div class="border border-[#eee9db] p-6 bg-white flex flex-col md:flex-row gap-6 items-center">
            <img 
              src="/uploads/foto_pimpinan/kasatpol.jpg" 
              alt="Kepala Satpol PP Sulsel" 
              class="w-32 h-40 object-cover border-2 border-[#F08300] bg-gray-100 flex-shrink-0"
              onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
            />
            <div class="space-y-2 text-center md:text-left">
              <span class="bg-[#101116] text-[#F08300] text-[10px] font-bold uppercase px-2.5 py-1 tracking-wider">
                Kepala Satuan (Eselon IIa)
              </span>
              <h3 class="text-lg font-extrabold text-[#101116]">Andi Arwin Azis S.STP., M.M</h3>
              <p class="text-xs text-[#6f7775]">Kepala Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan</p>
              <div class="text-xs text-[#2b2f3f] pt-2 border-t border-gray-100">
                Memimpin perumusan kebijakan, pengawasan operasional trantibum di 24 kabupaten/kota, serta koordinasi ketenteraman wilayah bersama Forkopimda Sulawesi Selatan.
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="border border-gray-200 p-4 bg-gray-50">
              <div class="text-[10px] font-bold text-[#F08300] uppercase">Sekretaris Satuan</div>
              <div class="font-bold text-xs text-[#101116] mt-0.5">Sekretaris Satpol PP Sulsel</div>
              <p class="text-[11px] text-[#6f7775] mt-1">Mengoordinasikan urusan perencanaan program, pengelolaan keuangan anggaran, dan tata usaha kepegawaian.</p>
            </div>
            <div class="border border-gray-200 p-4 bg-gray-50">
              <div class="text-[10px] font-bold text-[#F08300] uppercase">Kepala Bidang Penegakan Perda</div>
              <div class="font-bold text-xs text-[#101116] mt-0.5">Kabid Penegakan Perda</div>
              <p class="text-[11px] text-[#6f7775] mt-1">Memimpin penindakan yustisial dan non-yustisial serta penyidikan pelanggaran peraturan daerah oleh PPNS.</p>
            </div>
            <div class="border border-gray-200 p-4 bg-gray-50">
              <div class="text-[10px] font-bold text-[#F08300] uppercase">Kepala Bidang Trantibum</div>
              <div class="font-bold text-xs text-[#101116] mt-0.5">Kabid Trantibum</div>
              <p class="text-[11px] text-[#6f7775] mt-1">Mengendalikan operasi patroli ketertiban umum, penertiban fasilitas publik, dan pengawalan pejabat negara.</p>
            </div>
            <div class="border border-gray-200 p-4 bg-gray-50">
              <div class="text-[10px] font-bold text-[#F08300] uppercase">Kepala Bidang Linmas &amp; Damkar</div>
              <div class="font-bold text-xs text-[#101116] mt-0.5">Kabid Linmas &amp; Damkar</div>
              <p class="text-[11px] text-[#6f7775] mt-1">Menyelenggarakan pembinaan teknis anggota Satlinmas dan koordinasi penanganan kedaruratan kebakaran daerah.</p>
            </div>
          </div>
        </div>
      `
    }
  };

  $: slug = $page.params.slug || 'sejarah';
  $: current = profileData[slug] || profileData['sejarah'];

  const menuItems = [
    { slug: 'sejarah', label: 'Sejarah Singkat' },
    { slug: 'visi-misi', label: 'Visi dan Misi' },
    { slug: 'struktur', label: 'Struktur Organisasi' },
    { slug: 'tupoksi', label: 'Tugas Pokok & Fungsi' },
    { slug: 'pejabat', label: 'Profil Pejabat' }
  ];
</script>

<svelte:head>
  <title>{current.title} | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-6 px-4">
  <div class="max-w-6xl mx-auto flex flex-wrap items-center gap-2 text-[11px] text-[#6f7775]">
    <a href="/" class="hover:text-[#F08300]">Beranda</a>
    <span>/</span>
    <span class="hover:text-[#F08300]">Profil Instansi</span>
    <span>/</span>
    <span class="text-[#101116] font-semibold">{current.title}</span>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto">
  <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
    <!-- Sidebar Navigation -->
    <div class="lg:col-span-1 space-y-4">
      <div class="bg-white border border-[#eee9db] shadow-sm p-4">
        <h3 class="text-xs font-bold text-[#101116] uppercase tracking-wider border-b border-gray-100 pb-2 mb-3">
          Menu Profil
        </h3>
        <nav class="space-y-1">
          {#each menuItems as item}
            <a 
              href={`/profil/${item.slug}`}
              class={`block px-3 py-2 text-xs font-bold transition ${slug === item.slug ? 'bg-[#F08300] text-white' : 'text-[#2b2f3f] hover:bg-gray-100'}`}
            >
              {item.label}
            </a>
          {/each}
        </nav>
      </div>

      <!-- Quick Hotline Card -->
      <div class="bg-[#2b2f3f] text-white p-5 border-t-4 border-[#F08300]">
        <div class="text-[10px] text-[#F08300] font-bold uppercase tracking-wider mb-1">Pusat Informasi</div>
        <div class="font-bold text-xs">Satpol PP Prov. Sulsel</div>
        <div class="text-[11px] text-gray-300 mt-2 space-y-1">
          <div>Telp: <strong>0411-453105</strong></div>
          <div>Alamat: <strong>Jl. Urip Sumoharjo No. 269 Makassar</strong></div>
        </div>
      </div>
    </div>

    <!-- Main Content Column -->
    <div class="lg:col-span-3 bg-white border border-[#eee9db] shadow-sm p-6 md:p-8 space-y-6">
      <div class="border-b border-gray-100 pb-4">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Profil Resmi Satpol PP Prov. Sulsel</div>
        <h1 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">{current.title}</h1>
        <p class="text-xs text-[#6f7775] mt-1">{current.subtitle}</p>
        <div class="text-[10px] text-gray-400 mt-2">Pembaruan data terakhir: {current.updated}</div>
      </div>

      <div class="prose max-w-none text-xs md:text-sm text-[#2b2f3f] leading-relaxed">
        {@html current.contentHtml}
      </div>
    </div>
  </div>
</div>
