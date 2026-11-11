<script lang="ts">
  import { onMount } from 'svelte';
  import { page } from '$app/stores';
  import { Building2, Users } from 'lucide-svelte';
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
        <div class="bg-[#f9f4e8] border-l-4 border-[#F08300] p-6 mb-6 rounded-2xl shadow-xs">
          <div class="text-xs font-bold text-[#F08300] uppercase tracking-wider mb-1">Visi Utama</div>
          <blockquote class="text-base md:text-lg font-extrabold text-[#1a120c] italic leading-snug">
            &ldquo;Terwujudnya Polisi Pamong Praja yang Profesional dan Berwibawa dalam pelaksanaan tugas, menjadi pengayom dan pelayan masyarakat, serta Penegak Perda yang tangguh dan mumpuni di Sulawesi Selatan.&rdquo;
          </blockquote>
        </div>

        <h3 class="text-sm font-bold uppercase text-[#1a120c] tracking-wider mb-3">Misi Pembangunan</h3>
        <ol class="list-decimal pl-5 space-y-3 text-xs md:text-sm text-[#3d261a] leading-relaxed">
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
      subtitle: 'Bagan Susunan Organisasi dan Tata Kerja Dinamis (SOTK) Satpol PP Prov. Sulsel',
      updated: 'Terhubung CMS Organisasi',
      contentHtml: ''
    },
    'tupoksi': {
      title: 'Tugas Pokok dan Fungsi',
      subtitle: 'Landasan Yuridis dan Mandat Penyelenggaraan Ketertiban Berdasarkan PP No. 16/2018',
      updated: '18 Juni 2023',
      contentHtml: `
        <p class="leading-relaxed mb-4">Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan merupakan Perangkat Daerah yang mempunyai Tugas Pokok dan Fungsi serta Kewenangan sebagaimana diatur dalam Pasal 255, 256 dan 257 Undang-Undang Nomor 23 Tahun 2014 tentang Pemerintahan Daerah jo. Peraturan Pemerintah Nomor 16 Tahun 2018.</p>

        <div class="space-y-4">
          <div class="bg-[#f9f4e8]/75 border border-[#eee9db] p-5 rounded-xl">
            <h4 class="font-bold text-xs uppercase text-[#1a120c] mb-2 tracking-wider">Tugas Pokok</h4>
            <p class="text-xs text-[#3d261a] leading-relaxed">
              Menyelenggarakan urusan pemerintahan bidang ketenteraman dan ketertiban umum serta pelindungan masyarakat, menegakkan Peraturan Daerah (Perda) dan Peraturan Kepala Daerah (Perkada).
            </p>
          </div>

          <div class="bg-[#f9f4e8]/75 border border-[#eee9db] p-5 rounded-xl">
            <h4 class="font-bold text-xs uppercase text-[#1a120c] mb-2 tracking-wider">Fungsi Pelaksanaan</h4>
            <ol class="list-decimal pl-5 space-y-2 text-xs text-[#3d261a] leading-relaxed">
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

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-6 px-4">
  <div class="max-w-6xl mx-auto flex flex-wrap items-center gap-2 text-[11px] text-[#3d261a]/75">
    <a href="/" class="hover:text-[#F08300]">Beranda</a>
    <span>/</span>
    <span>Profil Instansi</span>
    <span>/</span>
    <span class="text-[#1a120c] font-bold">{current.title}</span>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto">
  <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
    <!-- Sidebar Navigation -->
    <div class="lg:col-span-1 space-y-4">
      <div class="bg-white border border-[#eee9db] shadow-xs p-4 rounded-2xl">
        <h3 class="text-xs font-bold text-[#1a120c] uppercase tracking-wider border-b border-[#eee9db] pb-2 mb-3">
          Menu Profil
        </h3>
        <nav class="space-y-1">
          {#each menuItems as item}
            <a
              href={`/profil/${item.slug}`}
              class={`block px-3.5 py-2.5 text-xs font-bold transition rounded-xl ${slug === item.slug ? 'bg-[#F08300] text-white shadow-xs' : 'text-[#3d261a] hover:bg-[#f9f4e8]'}`}
            >
              {item.label}
            </a>
          {/each}
        </nav>
      </div>

      <div class="bg-[#3d261a] text-white p-5 rounded-2xl border-t-4 border-[#F08300] shadow-xs">
        <div class="text-[10px] text-[#F08300] font-bold uppercase tracking-wider mb-1">Pusat Informasi</div>
        <div class="font-bold text-xs">Satpol PP Prov. Sulsel</div>
        <div class="text-[11px] text-[#f9f4e8]/80 mt-2 space-y-1">
          <div>Telp: <strong>{$siteSettingsStore.call_center}</strong></div>
          <div>Alamat: <strong>{$siteSettingsStore.office_address}</strong></div>
        </div>
      </div>
    </div>

    <!-- Main Content Column -->
    <div class="lg:col-span-3 bg-white border border-[#eee9db] shadow-xs p-6 md:p-8 space-y-6 rounded-2xl">
      <div class="border-b border-[#eee9db] pb-4">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Profil Resmi Satpol PP Prov. Sulsel</div>
        <h1 class="text-2xl md:text-3xl font-black text-[#1a120c] tracking-tight mt-1">{current.title}</h1>
        <p class="text-xs text-[#3d261a]/75 mt-1">{current.subtitle}</p>
        <div class="text-[10px] text-[#3d261a]/60 mt-2">Status data: {current.updated}</div>
      </div>

      {#if slug === 'struktur'}
        <div class="space-y-6 text-xs md:text-sm text-[#3d261a]">
          <p class="leading-relaxed">
            Berdasarkan Peraturan Gubernur Sulawesi Selatan tentang Kedudukan, Susunan Organisasi, Tugas dan Fungsi serta Tata Kerja Satuan Polisi Pamong Praja, berikut susunan unit kerja terstruktur (kode ID khusus dapat diperbarui melalui Panel Admin/Superadmin):
          </p>

          <!-- Pimpinan Tertinggi / Level 1 -->
          <div class="space-y-3">
            {#each rootNodes as root}
              <div class="bg-[#1a120c] text-white rounded-2xl p-6 border-b-4 border-[#F08300] flex flex-col sm:flex-row items-center gap-5 shadow-sm">
                {#if root.photo_url}
                  <img
                    src={root.photo_url}
                    alt={root.official_name}
                    class="w-20 h-20 rounded-2xl object-cover border-2 border-[#F08300] bg-[#241710] shrink-0"
                  />
                {:else}
                  <div class="w-16 h-16 rounded-2xl bg-[#241710] border border-[#3d261a] flex items-center justify-center text-[#F08300] shrink-0">
                    <Building2 class="w-7 h-7" />
                  </div>
                {/if}
                <div class="text-center sm:text-left min-w-0 flex-1">
                  <div class="flex flex-wrap items-center justify-center sm:justify-start gap-2 mb-1.5">
                    <span class="px-2.5 py-0.5 rounded-lg bg-[#241710] border border-[#3d261a] text-[#F08300] font-mono text-[11px] font-black">
                      {root.org_code}
                    </span>
                    {#if root.rank}
                      <span class="px-2.5 py-0.5 rounded-lg bg-[#F08300]/20 text-[#F08300] text-[11px] font-bold">
                        {root.rank}
                      </span>
                    {/if}
                  </div>
                  <h2 class="text-base md:text-lg font-black text-white">{root.position_title}</h2>
                  <div class="text-sm font-bold text-[#F08300] mt-0.5">{root.official_name}</div>
                  {#if root.description}
                    <p class="text-xs text-[#f9f4e8]/75 mt-2 leading-relaxed">{root.description}</p>
                  {/if}
                </div>
              </div>
            {/each}
          </div>

          <!-- Unit Bidang & Sub Bagian / Level >= 2 -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            {#each unitNodes as unit}
              <div class="border border-[#eee9db] bg-[#f9f4e8]/55 p-5 rounded-2xl flex flex-col justify-between space-y-3 hover:border-[#F08300] transition">
                <div>
                  <div class="flex items-center justify-between gap-2 mb-2">
                    <span class="px-2.5 py-0.5 rounded-lg bg-[#1a120c] text-[#F08300] font-mono text-[10px] font-black">
                      {unit.org_code}
                    </span>
                    {#if unit.rank}
                      <span class="text-[10px] font-bold text-[#3d261a]/75 bg-white px-2 py-0.5 rounded-lg border border-[#eee9db]">
                        {unit.rank}
                      </span>
                    {/if}
                  </div>
                  <h3 class="font-black text-sm text-[#1a120c]">{unit.position_title}</h3>
                  <div class="text-xs font-bold text-[#F08300] mt-0.5">{unit.official_name}</div>
                  {#if unit.parent_code}
                    <div class="text-[11px] text-[#3d261a]/70 mt-1">
                      Atasan: {getParentTitle(unit.parent_code)}
                    </div>
                  {/if}
                  {#if unit.sub_units && unit.sub_units.length > 0}
                    <ul class="text-xs text-[#3d261a] mt-3 space-y-1 list-disc pl-4">
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
            <div class="border border-[#eee9db] p-6 bg-[#f9f4e8]/50 flex flex-col md:flex-row gap-6 items-center rounded-2xl shadow-xs">
              {#if leader.photo_url}
                <img
                  src={leader.photo_url}
                  alt={leader.official_name}
                  class="w-32 h-40 object-cover border-2 border-[#F08300] bg-white flex-shrink-0 rounded-2xl shadow-xs"
                />
              {:else}
                <div class="w-28 h-28 rounded-2xl bg-[#1a120c] text-[#F08300] flex items-center justify-center shrink-0">
                  <Users class="w-10 h-10" />
                </div>
              {/if}
              <div class="space-y-2 text-center md:text-left">
                <div class="flex flex-wrap items-center justify-center md:justify-start gap-2">
                  <span class="bg-[#1a120c] text-[#F08300] font-mono text-[10px] font-bold px-2.5 py-1 rounded-lg">
                    {leader.org_code}
                  </span>
                  <span class="bg-[#F08300] text-white text-[10px] font-bold uppercase px-3 py-1 tracking-wider rounded-lg">
                    {leader.rank || 'Pimpinan'}
                  </span>
                </div>
                <h2 class="text-lg font-black text-[#1a120c]">{leader.official_name}</h2>
                <p class="text-xs font-bold text-[#3d261a]/80">{leader.position_title}</p>
                {#if leader.description}
                  <div class="text-xs text-[#3d261a] pt-2 border-t border-[#eee9db] leading-relaxed">
                    {leader.description}
                  </div>
                {/if}
              </div>
            </div>
          {/each}

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            {#each unitNodes as unit}
              <div class="border border-[#eee9db] p-5 bg-white rounded-2xl shadow-xs flex gap-4 items-start">
                {#if unit.photo_url}
                  <img
                    src={unit.photo_url}
                    alt={unit.official_name}
                    class="w-16 h-16 rounded-xl object-cover border border-[#eee9db] shrink-0"
                  />
                {:else}
                  <div class="w-14 h-14 rounded-xl bg-[#f9f4e8] border border-[#eee9db] text-[#3d261a] flex items-center justify-center shrink-0">
                    <Users class="w-6 h-6" />
                  </div>
                {/if}
                <div class="min-w-0">
                  <div class="flex flex-wrap items-center gap-1.5 mb-1">
                    <span class="font-mono text-[10px] font-black bg-[#1a120c] text-[#F08300] px-2 py-0.5 rounded-md">
                      {unit.org_code}
                    </span>
                    {#if unit.rank}
                      <span class="text-[10px] font-bold text-[#3d261a]/75">{unit.rank}</span>
                    {/if}
                  </div>
                  <div class="text-xs font-black text-[#F08300]">{unit.position_title}</div>
                  <div class="font-bold text-sm text-[#1a120c] mt-0.5">{unit.official_name}</div>
                  {#if unit.nip}
                    <div class="text-[11px] font-mono text-[#3d261a]/70">NIP {unit.nip}</div>
                  {/if}
                  {#if unit.description}
                    <p class="text-xs text-[#3d261a]/80 mt-2 leading-relaxed">{unit.description}</p>
                  {/if}
                </div>
              </div>
            {/each}
          </div>
        </div>
      {:else}
        <div class="prose max-w-none text-xs md:text-sm text-[#3d261a] leading-relaxed">
          {@html current.contentHtml}
        </div>
      {/if}
    </div>
  </div>
</div>
