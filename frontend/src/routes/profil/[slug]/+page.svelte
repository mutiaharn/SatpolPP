<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { Building2, Users, ChevronRight, Shield, Award, Landmark } from 'lucide-svelte';
  import { organizationStore, siteSettingsStore } from '$lib/stores/appState';
  import { fetchOrganization } from '$lib/api/client';

  onMount(() => {
    fetchOrganization();
  });

  interface ProfileItem {
    title: string;
    subtitle: string;
    updated: string;
    contentHtml: string;
  }

  $: orgNodes = [...$organizationStore].sort((a, b) =>
    a.level !== b.level ? a.level - b.level : a.sort_order - b.sort_order
  );
  $: rootNodes = orgNodes.filter(n => !n.parent_code || n.level === 1);
  $: unitNodes = orgNodes.filter(n => n.level >= 2);

  const profileData: Record<string, ProfileItem> = {
    'sejarah': {
      title: 'Sejarah Singkat Satpol PP',
      subtitle: 'Napak tilas panjang pengabdian Korps Praja Wibawa sejak 1950 dalam merajut ketenteraman dan mengayomi masyarakat di Sulawesi Selatan.',
      updated: '18 Juni 2023',
      contentHtml: `
        <div class="space-y-4 text-slate-700 leading-relaxed text-sm">
          <p class="leading-relaxed"><strong>Satuan Polisi Pamong Praja</strong> dibentuk berawal di Yogyakarta pada tanggal <strong>3 Maret 1950</strong>, dengan semboyan kehormatan <strong>PRAJA WIBAWA</strong>, untuk mewadahi penyelenggaraan sebagian kewenangan pemerintah daerah dalam memelihara ketertiban.</p>
          <p class="leading-relaxed">Sebelum proklamasi kemerdekaan, fungsi penjagaan keamanan ketertiban telah dilaksanakan pemerintah. Pasca proklamasi di tengah dinamika awal berdirinya NKRI, dibentuklah Detasemen Polisi sebagai Penjaga Keamanan Kapanewon di Yogyakarta berdasarkan Surat Perintah Jawatan Praja Daerah Istimewa Yogyakarta.</p>
          <p class="leading-relaxed">Seiring perkembangan tata kelola ketatanegaraan Indonesia, keberadaan Polisi Pamong Praja dikukuhkan melalui berbagai regulasi nasional hingga UU No. 23 Tahun 2014 tentang Pemerintahan Daerah dan PP No. 16 Tahun 2018. Di Sulawesi Selatan, Satpol PP Prov. Sulsel memegang peranan strategis sebagai komando koordinasi penegakan ketertiban umum lintas 24 Kabupaten dan Kota.</p>
        </div>
      `
    },
    'visi-misi': {
      title: 'Visi dan Misi Satpol PP Sulsel',
      subtitle: 'Tekad bersama mewujudkan aparatur penegak Perda yang profesional, berwibawa, humanis, serta berintegritas tinggi dalam melayani warga.',
      updated: '31 Oktober 2023',
      contentHtml: `
        <div class="bg-amber-50/70 border border-amber-200/60 border-l-4 border-l-[#F08300] p-6 mb-6 rounded-2xl shadow-xs">
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider mb-1.5">
            Visi Utama Pelayanan Daerah
          </div>
          <blockquote class="text-base md:text-lg font-bold text-slate-900 italic leading-snug">
            &ldquo;Terwujudnya Polisi Pamong Praja yang Profesional dan Berwibawa dalam pelaksanaan tugas, menjadi pengayom dan pelayan masyarakat, serta Penegak Perda yang tangguh dan mumpuni di Sulawesi Selatan.&rdquo;
          </blockquote>
        </div>

        <h3 class="text-xs font-bold uppercase text-slate-900 tracking-wider mb-4">Misi Pembangunan Strategis</h3>
        <ol class="list-decimal pl-5 space-y-3 text-xs md:text-sm text-slate-700 leading-relaxed">
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
      subtitle: 'Susunan kepemimpinan, pejabat struktural, dan simpul unit kerja terpadu Satpol PP Provinsi Sulawesi Selatan.',
      updated: 'Terhubung CMS Organisasi',
      contentHtml: ''
    },
    'tupoksi': {
      title: 'Tugas Pokok & Fungsi',
      subtitle: 'Mandat pengabdian aparatur dalam menegakkan peraturan daerah, menyelenggarakan ketertiban umum, dan melindungi keselamatan warga.',
      updated: '18 Juni 2023',
      contentHtml: `
        <p class="leading-relaxed mb-6 text-slate-700 text-sm">Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan merupakan Perangkat Daerah yang mempunyai Tugas Pokok dan Fungsi serta Kewenangan sebagaimana diatur dalam Pasal 255, 256 dan 257 Undang-Undang Nomor 23 Tahun 2014 tentang Pemerintahan Daerah jo. Peraturan Pemerintah Nomor 16 Tahun 2018.</p>

        <div class="space-y-4">
          <div class="bg-slate-50/90 border border-slate-200/80 p-5 rounded-2xl">
            <h4 class="font-bold text-xs uppercase text-slate-900 mb-2 tracking-wider">
              Tugas Pokok
            </h4>
            <p class="text-xs sm:text-sm text-slate-600 leading-relaxed">
              Menyelenggarakan urusan pemerintahan bidang ketenteraman dan ketertiban umum serta pelindungan masyarakat, menegakkan Peraturan Daerah (Perda) dan Peraturan Kepala Daerah (Perkada).
            </p>
          </div>

          <div class="bg-slate-50/90 border border-slate-200/80 p-5 rounded-2xl">
            <h4 class="font-bold text-xs uppercase text-slate-900 mb-2 tracking-wider">
              Fungsi Pelaksanaan
            </h4>
            <ol class="list-decimal pl-5 space-y-2.5 text-xs sm:text-sm text-slate-600 leading-relaxed">
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
      subtitle: 'Mengenal jajaran pimpinan dan pejabat administrator yang mengawal tugas ketertiban di 24 Kabupaten dan Kota se-Sulawesi Selatan.',
      updated: 'Terhubung CMS Organisasi',
      contentHtml: ''
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

  function getParentTitle(code: string | null) {
    if (!code) return null;
    return orgNodes.find(n => n.org_code === code)?.position_title || code;
  }
</script>

<svelte:head>
  <title>{current.title} | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#246bfd]/8 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <!-- Breadcrumb Pills -->
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-slate-400">Profil Instansi</span>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">{current.title}</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">{current.title}</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      {current.subtitle}
    </p>
  </div>
</div>

<!-- Main Content Area -->
<div class="py-10 px-4 max-w-6xl mx-auto relative">
  <!-- Subtle light blue gradient splashes giving soft depth & living glass refraction -->
  <div class="absolute top-12 left-10 w-96 h-96 rounded-full bg-sky-400/15 blur-3xl pointer-events-none -z-10"></div>
  <div class="absolute top-1/2 -right-10 w-96 h-96 rounded-full bg-blue-500/12 blur-3xl pointer-events-none -z-10"></div>
  <div class="absolute bottom-10 left-1/3 w-80 h-80 rounded-full bg-sky-300/15 blur-3xl pointer-events-none -z-10"></div>

  <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
    <!-- Sidebar Navigation -->
    <div class="lg:col-span-1 space-y-5">
      <div class="ds-glass border border-white/90 p-4 rounded-3xl shadow-glass ds-specular-top">
        <h3 class="text-xs font-bold text-slate-900 uppercase tracking-wider border-b border-white/80 pb-3 mb-3">
          Menu Profil
        </h3>
        <nav class="space-y-2">
          {#each menuItems as item}
            <a
              href={`/profil/${item.slug}`}
              class={`block px-4 py-2.5 text-xs font-bold transition duration-200 rounded-xl ds-press ${
                slug === item.slug
                  ? 'ds-pill-active'
                  : 'ds-pill-inactive'
              }`}
            >
              {item.label}
            </a>
          {/each}
        </nav>
      </div>

      <!-- Quick Info Card -->
      <div class="ds-glass-dark text-white p-5 rounded-3xl border border-white/15 shadow-glass ds-specular-dark">
        <div class="text-[10px] text-[#F08300] font-bold uppercase tracking-wider mb-1">Pusat Informasi Resmi</div>
        <div class="font-bold text-xs text-white">Satpol PP Prov. Sulsel</div>
        <div class="text-[11px] text-slate-300 mt-2.5 space-y-1.5 font-light">
          <div>Telp: <strong class="text-white font-medium">{$siteSettingsStore.call_center}</strong></div>
          <div class="leading-relaxed">Alamat: <strong class="text-white font-medium">{$siteSettingsStore.office_address}</strong></div>
        </div>
      </div>
    </div>

    <!-- Main Content Column -->
    <div class="lg:col-span-3 ds-glass-strong border border-white/90 p-6 md:p-8 space-y-6 rounded-3xl shadow-glass ds-specular-top">
      <div class="border-b border-white/80 pb-4">
        <div class="flex flex-wrap items-center justify-between gap-2">
          <span class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Informasi Kelembagaan</span>
          <span class="text-[11px] text-slate-500 font-medium bg-slate-100 px-2.5 py-0.5 rounded-full border border-slate-200">
            Status: {current.updated}
          </span>
        </div>
        <h2 class="text-2xl font-extrabold text-slate-900 tracking-tight mt-2">{current.title}</h2>
        <p class="text-xs sm:text-sm text-slate-500 mt-1">{current.subtitle}</p>
      </div>

      {#if slug === 'struktur'}
        <div class="space-y-6 text-xs md:text-sm text-slate-700">
          <p class="leading-relaxed">
            Berdasarkan Peraturan Gubernur Sulawesi Selatan tentang Kedudukan, Susunan Organisasi, Tugas dan Fungsi serta Tata Kerja Satuan Polisi Pamong Praja, berikut susunan unit kerja terstruktur:
          </p>

          <!-- Pimpinan Tertinggi / Level 1 (Simpul Utama Struktur Organisasi) -->
          <div class="space-y-3">
            {#each rootNodes as root}
              <div class="ds-glass border border-slate-200/80 p-5 sm:p-6 rounded-2xl shadow-glass hover:border-[#246bfd] transition relative overflow-hidden group">
                <!-- Subtle light blue gradient splash in card corner -->
                <div class="absolute -top-12 -right-12 w-48 h-48 rounded-full bg-sky-400/12 blur-2xl pointer-events-none"></div>

                <div class="relative z-10">
                  <div class="flex flex-wrap items-center justify-between gap-2 mb-2">
                    <div class="flex items-center gap-2">
                      <span class="px-2.5 py-0.5 rounded-lg bg-[#0b172a] text-[#F08300] font-mono text-[10px] font-bold">
                        {root.org_code}
                      </span>
                      {#if root.rank}
                        <span class="text-[10px] font-bold text-slate-600 bg-slate-100 px-2 py-0.5 rounded-lg border border-slate-200">
                          {root.rank}
                        </span>
                      {/if}
                    </div>
                    <span class="text-[10px] font-bold text-blue-600 bg-blue-50/90 px-2.5 py-0.5 rounded-full border border-blue-200/70">
                      Pimpinan Perangkat Daerah
                    </span>
                  </div>

                  <h3 class="font-bold text-base sm:text-lg text-slate-900 group-hover:text-blue-600 transition-colors">
                    {root.position_title}
                  </h3>
                  <div class="text-xs font-bold text-[#F08300] mt-0.5">{root.official_name}</div>

                  {#if root.description}
                    <p class="text-xs text-slate-600 mt-2.5 leading-relaxed font-light border-t border-slate-100 pt-2.5">
                      {root.description}
                    </p>
                  {/if}
                </div>
              </div>
            {/each}
          </div>

          <!-- Tree Hierarchy Connector -->
          <div class="flex flex-col items-center justify-center py-1" aria-hidden="true">
            <div class="w-0.5 h-3.5 bg-gradient-to-b from-[#246bfd]/50 to-[#246bfd]/20"></div>
            <div class="px-3 py-1 rounded-full bg-blue-50/90 border border-blue-200/60 text-[10px] font-semibold text-blue-700 tracking-wider uppercase shadow-xs">
              Unit Kerja & Bidang Operasional
            </div>
            <div class="w-0.5 h-3.5 bg-gradient-to-b from-[#246bfd]/20 to-[#246bfd]/50"></div>
          </div>

          <!-- Unit Bidang & Sub Bagian / Level >= 2 -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            {#each unitNodes as unit}
              <div class="ds-glass border border-slate-200/80 p-5 rounded-2xl flex flex-col justify-between space-y-3 hover:border-[#246bfd] hover:shadow-glass-hover transition">
                <div>
                  <div class="flex items-center justify-between gap-2 mb-2">
                    <span class="px-2.5 py-0.5 rounded-lg bg-[#0b172a] text-[#F08300] font-mono text-[10px] font-bold">
                      {unit.org_code}
                    </span>
                    {#if unit.rank}
                      <span class="text-[10px] font-bold text-slate-600 bg-slate-100 px-2 py-0.5 rounded-lg border border-slate-200">
                        {unit.rank}
                      </span>
                    {/if}
                  </div>
                  <h4 class="font-bold text-sm text-slate-900">{unit.position_title}</h4>
                  <div class="text-xs font-bold text-[#F08300] mt-0.5">{unit.official_name}</div>
                  {#if unit.parent_code}
                    <div class="text-[11px] text-slate-500 mt-1">
                      Atasan: {getParentTitle(unit.parent_code)}
                    </div>
                  {/if}
                  {#if unit.sub_units && unit.sub_units.length > 0}
                    <ul class="text-xs text-slate-600 mt-3 space-y-1 list-disc pl-4">
                      {#each unit.sub_units as sub}
                        <li>{sub}</li>
                      {/each}
                    </ul>
                  {/if}
                </div>
              </div>
            {/each}
          </div>
        </div>
      {:else if slug === 'pejabat'}
        <div class="space-y-6">
          {#each rootNodes as leader}
            <div class="ds-glass border border-slate-200/80 p-6 flex flex-col md:flex-row gap-6 items-center rounded-2xl shadow-glass">
              {#if leader.photo_url}
                <img
                  src={leader.photo_url}
                  alt={leader.official_name}
                  class="w-32 h-40 object-cover border-2 border-[#F08300] bg-white flex-shrink-0 rounded-2xl shadow-xs"
                />
              {:else}
                <div class="w-28 h-28 rounded-2xl bg-[#0b172a] text-[#F08300] flex items-center justify-center shrink-0">
                  <Users class="w-10 h-10" />
                </div>
              {/if}
              <div class="space-y-2 text-center md:text-left">
                <div class="flex flex-wrap items-center justify-center md:justify-start gap-2">
                  <span class="bg-[#0b172a] text-[#F08300] font-mono text-[10px] font-bold px-2.5 py-1 rounded-lg">
                    {leader.org_code}
                  </span>
                  <span class="bg-[#F08300] text-white text-[10px] font-bold uppercase px-3 py-1 tracking-wider rounded-lg">
                    {leader.rank || 'Pimpinan'}
                  </span>
                </div>
                <h3 class="text-lg font-bold text-slate-900">{leader.official_name}</h3>
                <p class="text-xs font-bold text-slate-600">{leader.position_title}</p>
                {#if leader.description}
                  <div class="text-xs text-slate-600 pt-2 border-t border-slate-200 leading-relaxed font-light">
                    {leader.description}
                  </div>
                {/if}
              </div>
            </div>
          {/each}

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            {#each unitNodes as unit}
              <div class="ds-glass border border-slate-200/80 p-5 rounded-2xl shadow-xs flex gap-4 items-start hover:border-[#246bfd] transition">
                {#if unit.photo_url}
                  <img
                    src={unit.photo_url}
                    alt={unit.official_name}
                    class="w-16 h-16 rounded-xl object-cover border border-slate-200 shrink-0"
                  />
                {:else}
                  <div class="w-14 h-14 rounded-xl bg-slate-100 border border-slate-200 text-slate-600 flex items-center justify-center shrink-0">
                    <Users class="w-6 h-6" />
                  </div>
                {/if}
                <div class="min-w-0">
                  <div class="flex flex-wrap items-center gap-1.5 mb-1">
                    <span class="font-mono text-[10px] font-bold bg-[#0b172a] text-[#F08300] px-2 py-0.5 rounded-md">
                      {unit.org_code}
                    </span>
                    {#if unit.rank}
                      <span class="text-[10px] font-bold text-slate-500">{unit.rank}</span>
                    {/if}
                  </div>
                  <div class="text-xs font-bold text-[#F08300]">{unit.position_title}</div>
                  <div class="font-bold text-sm text-slate-900 mt-0.5">{unit.official_name}</div>
                  {#if unit.nip}
                    <div class="text-[11px] font-mono text-slate-500">NIP {unit.nip}</div>
                  {/if}
                  {#if unit.description}
                    <p class="text-xs text-slate-600 mt-2 leading-relaxed font-light">{unit.description}</p>
                  {/if}
                </div>
              </div>
            {/each}
          </div>
        </div>
      {:else}
        <div class="prose max-w-none text-xs md:text-sm text-slate-700 leading-relaxed">
          {@html current.contentHtml}
        </div>
      {/if}
    </div>
  </div>
</div>
