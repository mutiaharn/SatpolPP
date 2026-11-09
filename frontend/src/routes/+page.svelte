<script lang="ts">
  import { onMount } from 'svelte';
  import { articlesStore, galleryStore } from '$lib/stores/appState';
  import ComplaintForm from '$lib/components/ComplaintForm.svelte';
  import TicketTracker from '$lib/components/TicketTracker.svelte';

  let currentSlide = 0;
  let activeTab: 'lapor' | 'lacak' = 'lapor';

  const slides = [
    {
      image: '/uploads/foto_banner/1716426402berita.png',
      tag: 'PRAJA WIBAWA',
      title: 'Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan',
      desc: 'Menegakkan Peraturan Daerah, Menyelenggarakan Ketenteraman dan Ketertiban Umum, serta Pelindungan Masyarakat.'
    },
    {
      image: '/uploads/foto_banner/1747198131welcome.png',
      tag: 'PELAYANAN PUBLIK',
      title: 'SIAP TERTIB: Sistem Informasi Penegakan Perda & Trantibum',
      desc: 'Saluran pengaduan elektronik terpadu masyarakat untuk mewujudkan Sulawesi Selatan yang aman, tertib, dan kondusif.'
    },
    {
      image: '/uploads/foto_banner/174089363110.png',
      tag: 'PENGAWALAN ASET',
      title: 'Penertiban dan Pengamanan Aset Pemerintah Daerah',
      desc: 'Optimalisasi dan penertiban aset barang milik daerah demi pemanfaatan yang tertib dan berkesinambungan.'
    }
  ];

  onMount(() => {
    // Auto rotate banner every 6 seconds
    const timer = setInterval(() => {
      currentSlide = (currentSlide + 1) % slides.length;
    }, 6000);
    return () => clearInterval(timer);
  });

  // Smooth scroll presisi: tepat memposisikan bagian di bawah sticky navbar
  function scrollToTarget(id: string) {
    const el = document.getElementById(id);
    if (!el) return;
    const header = document.querySelector('header');
    const headerHeight = header ? Math.round(header.getBoundingClientRect().height) : 106;
    const elPos = el.getBoundingClientRect().top + window.pageYOffset;
    window.scrollTo({
      top: Math.max(0, Math.round(elPos - headerHeight)),
      behavior: 'smooth'
    });
  }
</script>

<!-- 1. Hero Banner Slider (Poin 6: Dibuat Full 1 Layar Penuh / 100vh saat dibuka) -->
<section class="relative bg-[#1a120c] text-white min-h-[calc(100vh-106px)] flex items-center overflow-hidden">
  <!-- Slide Images -->
  {#each slides as slide, idx}
    <div 
      class={`absolute inset-0 bg-cover bg-center transition-opacity duration-1000 ${currentSlide === idx ? 'opacity-100 z-10' : 'opacity-0 z-0'}`}
      style={`background-image: linear-gradient(rgba(26, 18, 12, 0.74), rgba(26, 18, 12, 0.86)), url('${slide.image}')`}
    ></div>
  {/each}

  <div class="max-w-7xl mx-auto px-4 py-16 md:py-24 relative z-20 w-full">
    <div class="max-w-3xl space-y-5">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-[#F08300]/20 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-[#F08300]/30 backdrop-blur-xs">
        <span class="w-2 h-2 rounded-full bg-[#F08300] animate-pulse"></span>
        {slides[currentSlide].tag}
      </div>
      <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-extrabold tracking-tight leading-[1.15]">
        {slides[currentSlide].title}
      </h1>
      <p class="text-sm md:text-base lg:text-lg text-stone-200 leading-relaxed max-w-2xl">
        {slides[currentSlide].desc}
      </p>
      <div class="pt-3 flex flex-wrap gap-4">
        <a 
          href="#layanan-siaptertib" 
          on:click={(e) => { e.preventDefault(); scrollToTarget('layanan-siaptertib'); }}
          class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-7 py-4 rounded-xl transition shadow-lg hover:shadow-xl hover:-translate-y-0.5 transform cursor-pointer"
        >
          Lapor SIAP TERTIB
        </a>
        <a href="/dokumen" class="bg-white/10 hover:bg-white/20 text-white font-bold text-xs uppercase px-7 py-4 rounded-xl border border-white/25 transition backdrop-blur-md">
          Dokumen Publik
        </a>
      </div>
    </div>
  </div>

  <!-- Scroll Down Indicator (Penunjuk Gulir ke Bawah dengan Scroll Presisi) -->
  <a 
    href="#berita-terbaru" 
    on:click={(e) => { e.preventDefault(); scrollToTarget('berita-terbaru'); }}
    class="hidden md:flex absolute bottom-8 left-1/2 -translate-x-1/2 z-30 flex-col items-center gap-1.5 text-stone-300 hover:text-[#F08300] transition group focus:outline-none cursor-pointer"
    aria-label="Gulir ke konten berita"
  >
    <span class="text-[10px] font-bold uppercase tracking-widest text-[#F08300]">Jelajahi Portal</span>
    <svg class="w-5 h-5 animate-bounce text-[#F08300]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
    </svg>
  </a>

  <!-- Slide Navigation Indicators -->
  <div class="absolute bottom-8 right-8 z-30 flex items-center gap-2">
    {#each slides as _, i}
      <button 
        class={`h-2.5 transition-all rounded-full ${currentSlide === i ? 'w-8 bg-[#F08300]' : 'w-2.5 bg-white/40 hover:bg-white/70'}`}
        on:click={() => currentSlide = i}
        aria-label={`Slide ${i + 1}`}
      ></button>
    {/each}
  </div>
</section>

<!-- 2. Berita Terbaru (Poin 20: Card Berita Fixed & Balanced, pas 1 layar saat scroll Jelajahi) -->
<section id="berita-terbaru" class="scroll-mt-[106px] pt-8 md:pt-10 pb-16 px-4 max-w-7xl mx-auto">
  <div class="text-center max-w-xl mx-auto mb-8 md:mb-10">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Publikasi Resmi</div>
    <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Berita Terbaru</h2>
    <div class="w-12 h-1 bg-[#F08300] mx-auto mt-2.5 rounded-full"></div>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-8 items-stretch">
    {#each $articlesStore.slice(0, 3) as item}
      <article class="bg-white border border-[#eee9db] rounded-xl overflow-hidden flex flex-col justify-between h-full hover:shadow-lg transition">
        <div>
          <div class="relative h-52 bg-gray-100 overflow-hidden">
            <img 
              src={item.thumbnail} 
              alt={item.title} 
              class="w-full h-full object-cover hover:scale-105 transition duration-500" 
              onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
            />
            <div class="absolute top-0 right-0 bg-[#F08300] text-white text-[11px] font-bold px-3 py-1.5 uppercase rounded-bl-lg shadow-xs">
              {new Date(item.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'short' })}
            </div>
          </div>
          <div class="p-6 space-y-3">
            <div class="text-[11px] font-bold uppercase tracking-wider text-[#F08300]">
              {item.category}
            </div>
            <h3 class="font-bold text-base text-[#101116] leading-snug line-clamp-2 min-h-[2.8rem] hover:text-[#F08300] transition">
              <a href={`/berita/${item.slug}`}>{item.title}</a>
            </h3>
            <p class="text-xs text-stone-500 leading-relaxed line-clamp-3 min-h-[3.2rem]">
              {item.summary}
            </p>
          </div>
        </div>
        <div class="px-6 pb-6 pt-3 border-t border-gray-100 mt-auto">
          <a href={`/berita/${item.slug}`} class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline flex items-center gap-1.5">
            <span>Selengkapnya</span>
            <span>&rarr;</span>
          </a>
        </div>
      </article>
    {/each}
  </div>

  <!-- Link Menuju Halaman Berita Lengkap -->
  <div class="mt-8 text-center">
    <a 
      href="/berita" 
      class="inline-flex items-center gap-2 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2 px-4 rounded-lg hover:bg-[#f9f4e8] transition duration-200 group"
    >
      <span>Lihat Berita Lainnya</span>
      <svg class="w-4 h-4 text-[#F08300] transform group-hover:translate-x-1.5 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3" />
      </svg>
    </a>
  </div>
</section>

<!-- 3. Profil Pimpinan Satpol PP Sulsel (Life-size 3D Cutout Pimpinan Berdiri Langsung) -->
<section class="pt-12 pb-16 px-4 bg-[#f9f4e8] border-t border-b border-[#eee9db] overflow-hidden">
  <div class="max-w-7xl mx-auto grid grid-cols-1 lg:grid-cols-12 gap-8 lg:gap-12 items-end">
    <!-- Left: Foto Pimpinan 3D Life-size (Sangat Besar Seolah Berdiri Langsung) -->
    <div class="lg:col-span-5 flex justify-center items-end relative">
      <div class="relative flex flex-col items-center group w-full max-w-md">
        <!-- Grand 3D Portal Arch Backdrop -->
        <div class="absolute bottom-0 left-1/2 -translate-x-1/2 w-[300px] sm:w-[360px] md:w-[400px] h-[480px] sm:h-[560px] md:h-[620px] rounded-t-full bg-gradient-to-b from-amber-500/20 via-white/50 to-transparent border-t-2 border-x-2 border-amber-500/30 -z-10 shadow-inner"></div>
        <div class="absolute top-1/3 left-1/2 -translate-x-1/2 -translate-y-1/2 w-80 h-80 rounded-full bg-amber-500/15 blur-3xl -z-20"></div>

        <!-- Ground Contact Shadow (Bayangan Berdiri Nyata di Lantai) -->
        <div class="absolute -bottom-2 left-1/2 -translate-x-1/2 w-4/5 h-8 bg-black/30 rounded-[100%] blur-md -z-10"></div>

        <!-- 3D Cutout HD Image (Tinggi Hingga 680px - Life-size Impact) -->
        <div class="relative z-10 transition-transform duration-500 ease-out group-hover:-translate-y-3 group-hover:scale-[1.02]">
          <img 
            src="/uploads/foto_pimpinan/pimpinan_cutout_hd.png" 
            alt="Andi Arwin Azis S.STP., M.M" 
            class="w-auto h-[480px] sm:h-[560px] md:h-[620px] lg:h-[680px] object-contain drop-shadow-[0_25px_35px_rgba(42,29,23,0.40)] drop-shadow-[0_10px_15px_rgba(42,29,23,0.22)] select-none pointer-events-none" 
          />

          <!-- Floating 3D Podium Badge -->
          <div class="absolute bottom-4 left-1/2 -translate-x-1/2 w-[94%] sm:w-max max-w-md bg-[#1a120c]/95 backdrop-blur-md text-white px-6 py-3 rounded-xl border border-amber-500/40 shadow-[0_16px_32px_rgba(26,18,12,0.50)] text-center transform transition-transform duration-300 group-hover:scale-105">
            <div class="font-black text-xs sm:text-sm md:text-base tracking-wide text-white uppercase whitespace-nowrap">
              ANDI ARWIN AZIS S.STP., M.M
            </div>
            <div class="text-[10px] sm:text-[11px] text-[#F08300] font-bold uppercase tracking-wider mt-0.5 whitespace-nowrap">
              Kepala Satuan Polisi Pamong Praja Prov. Sulsel
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Right: Sambutan & Tugas Pokok -->
    <div class="lg:col-span-7 space-y-5 pb-4">
      <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-[#F08300]/15 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-[#F08300]/25">
        <span class="w-1.5 h-1.5 rounded-full bg-[#F08300]"></span>
        Pimpinan Instansi
      </div>
      <h2 class="text-2xl sm:text-3xl md:text-4xl font-extrabold text-[#101116] tracking-tight leading-tight">
        Menegakkan Perda Secara Humanis dan Berkeadilan
      </h2>
      <div class="w-16 h-1.5 bg-[#F08300] rounded-full"></div>

      <p class="text-sm md:text-base text-stone-700 leading-relaxed pt-1">
        Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan senantiasa mengedepankan pendekatan humanis, persuasif, dan profesional dalam menegakkan Peraturan Daerah dan menjaga ketertiban umum demi kesejahteraan masyarakat Sulawesi Selatan.
      </p>

      <div class="pt-2 space-y-2.5 text-xs sm:text-sm text-stone-700">
        <div class="flex items-start gap-2.5">
          <span class="text-[#F08300] font-black text-base leading-none">&bull;</span>
          <span>Penyelenggaraan ketenteraman dan ketertiban umum di seluruh wilayah Sulawesi Selatan.</span>
        </div>
        <div class="flex items-start gap-2.5">
          <span class="text-[#F08300] font-black text-base leading-none">&bull;</span>
          <span>Penegakan Peraturan Daerah (Perda) dan Peraturan Kepala Daerah (Perkada).</span>
        </div>
        <div class="flex items-start gap-2.5">
          <span class="text-[#F08300] font-black text-base leading-none">&bull;</span>
          <span>Pengawalan dan pengamanan aset milik Pemerintah Provinsi Sulawesi Selatan.</span>
        </div>
        <div class="flex items-start gap-2.5">
          <span class="text-[#F08300] font-black text-base leading-none">&bull;</span>
          <span>Pembinaan, bimbingan masyarakat, dan koordinasi perlindungan masyarakat (Linmas).</span>
        </div>
      </div>

      <!-- Tombol Cokelat Praja Menggantikan Biru -->
      <div class="pt-3">
        <a href="/profil/pejabat" class="inline-block bg-[#3d261a] hover:bg-[#271810] text-white text-xs font-bold uppercase px-7 py-3.5 rounded-xl transition shadow-md hover:shadow-lg hover:-translate-y-0.5 transform">
          Profil Struktur Organisasi &rarr;
        </a>
      </div>
    </div>
  </div>
</section>

<!-- 4. Dokumen Publik (Full Screen Putih Bersih: Bagian Kuning di Bawah Hanya Terlihat Saat Di-scroll) -->
<section id="dokumen-publik" class="bg-white min-h-[calc(100vh-106px)] flex flex-col justify-center py-12 md:py-16 px-4">
  <div class="max-w-7xl mx-auto w-full">
    <div class="text-center max-w-xl mx-auto mb-10 md:mb-12">
      <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Akuntabilitas Kinerja</div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Dokumen Publik</h2>
      <div class="w-12 h-1 bg-[#F08300] mx-auto mt-3 rounded-full"></div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 items-stretch">
      <div class="bg-white border border-[#eee9db] rounded-xl p-6 hover:border-[#F08300] hover:shadow-md transition flex flex-col justify-between h-full">
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Perencanaan</div>
          <h3 class="font-bold text-base text-[#101116] mt-1">Renja (Rencana Kerja)</h3>
          <p class="text-xs text-stone-500 mt-2 line-clamp-2">Dokumen perencanaan tahunan program dan target kegiatan Satpol PP Prov. Sulsel.</p>
        </div>
        <div class="mt-4 pt-3 border-t border-gray-100">
          <a href="/dokumen?category=renja" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
        </div>
      </div>

      <div class="bg-white border border-[#eee9db] rounded-xl p-6 hover:border-[#F08300] hover:shadow-md transition flex flex-col justify-between h-full">
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Strategis</div>
          <h3 class="font-bold text-base text-[#101116] mt-1">Renstra (Rencana Strategis)</h3>
          <p class="text-xs text-stone-500 mt-2 line-clamp-2">Rencana strategis jangka menengah pencapaian visi, misi, dan sasaran instansi.</p>
        </div>
        <div class="mt-4 pt-3 border-t border-gray-100">
          <a href="/dokumen?category=renstra" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
        </div>
      </div>

      <div class="bg-white border border-[#eee9db] rounded-xl p-6 hover:border-[#F08300] hover:shadow-md transition flex flex-col justify-between h-full">
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Penganggaran</div>
          <h3 class="font-bold text-base text-[#101116] mt-1">DPA (Pelaksanaan Anggaran)</h3>
          <p class="text-xs text-stone-500 mt-2 line-clamp-2">Alokasi anggaran belanja dan rincian pembiayaan kegiatan operasional.</p>
        </div>
        <div class="mt-4 pt-3 border-t border-gray-100">
          <a href="/dokumen?category=dpa" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
        </div>
      </div>

      <div class="bg-white border border-[#eee9db] rounded-xl p-6 hover:border-[#F08300] hover:shadow-md transition flex flex-col justify-between h-full">
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Kinerja</div>
          <h3 class="font-bold text-base text-[#101116] mt-1">LAKIP / LKjIP</h3>
          <p class="text-xs text-stone-500 mt-2 line-clamp-2">Laporan akuntabilitas kinerja instansi pemerintah tahunan.</p>
        </div>
        <div class="mt-4 pt-3 border-t border-gray-100">
          <a href="/dokumen?category=lakip" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
        </div>
      </div>

      <div class="bg-white border border-[#eee9db] rounded-xl p-6 hover:border-[#F08300] hover:shadow-md transition flex flex-col justify-between h-full">
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Pertanggungjawaban</div>
          <h3 class="font-bold text-base text-[#101116] mt-1">Laporan Keuangan</h3>
          <p class="text-xs text-stone-500 mt-2 line-clamp-2">Laporan pertanggungjawaban realisasi keuangan dan aset tahunan.</p>
        </div>
        <div class="mt-4 pt-3 border-t border-gray-100">
          <a href="/dokumen?category=keuangan" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
        </div>
      </div>

      <div class="bg-white border border-[#eee9db] rounded-xl p-6 hover:border-[#F08300] hover:shadow-md transition flex flex-col justify-between h-full">
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Regulasi</div>
          <h3 class="font-bold text-base text-[#101116] mt-1">Perda &amp; SOP Satpol PP</h3>
          <p class="text-xs text-stone-500 mt-2 line-clamp-2">Peraturan Daerah Provinsi Sulawesi Selatan dan Standar Operasional Prosedur.</p>
        </div>
        <div class="mt-4 pt-3 border-t border-gray-100">
          <a href="/dokumen?category=perda" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
        </div>
      </div>
    </div>

    <!-- Link Menuju Halaman Dokumen Lengkap -->
    <div class="mt-8 text-center">
      <a 
        href="/dokumen" 
        class="inline-flex items-center gap-2 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2 px-4 rounded-lg hover:bg-[#f9f4e8] transition duration-200 group"
      >
        <span>Lihat Dokumen Lainnya</span>
        <svg class="w-4 h-4 text-[#F08300] transform group-hover:translate-x-1.5 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3" />
        </svg>
      </a>
    </div>
  </div>
</section>

<!-- 5. Layanan SIAP TERTIB & Pelacakan Aduan (Latar Krem #f9f4e8 Klasik) -->
<section id="layanan-siaptertib" class="scroll-mt-36 py-16 px-4 bg-[#f9f4e8] border-t border-b border-[#eee9db]">
  <div class="max-w-4xl mx-auto">
    <div class="text-center max-w-xl mx-auto mb-8">
      <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Layanan Pengaduan Masyarakat</div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">SIAP TERTIB</h2>
      <div class="w-12 h-1 bg-[#F08300] mx-auto mt-3 rounded-full"></div>
      <p class="text-xs text-stone-500 mt-3">
        Sistem Informasi Penegakan Peraturan Daerah &amp; Ketenteraman Ketertiban Umum.
      </p>
    </div>

    <!-- Tab Selector -->
    <div class="flex justify-center mb-6">
      <div class="inline-flex border border-gray-300 bg-white p-1 rounded-xl shadow-xs">
        <button 
          class={`px-6 py-2.5 text-xs font-bold uppercase tracking-wider rounded-lg transition ${activeTab === 'lapor' ? 'bg-[#F08300] text-white shadow-xs' : 'text-[#101116] hover:bg-[#f9f4e8]'}`}
          on:click={() => activeTab = 'lapor'}
        >
          Formulir Pengaduan
        </button>
        <button 
          class={`px-6 py-2.5 text-xs font-bold uppercase tracking-wider rounded-lg transition ${activeTab === 'lacak' ? 'bg-[#F08300] text-white shadow-xs' : 'text-[#101116] hover:bg-[#f9f4e8]'}`}
          on:click={() => activeTab = 'lacak'}
        >
          Lacak Resi Aduan
        </button>
      </div>
    </div>

    {#if activeTab === 'lapor'}
      <ComplaintForm />
    {:else}
      <TicketTracker />
    {/if}
  </div>
</section>

<!-- 6. Galeri Kegiatan (Matching gallery-one) -->
<section id="galeri-kegiatan" class="pt-7 md:pt-8 pb-7 md:pb-8 px-4 max-w-7xl mx-auto">
  <div class="text-center max-w-xl mx-auto mb-5">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Dokumentasi</div>
    <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Galeri Kegiatan</h2>
    <div class="w-12 h-1 bg-[#F08300] mx-auto mt-2 rounded-full"></div>
  </div>

  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    {#each $galleryStore.slice(0, 4) as gal}
      <div class="h-40 md:h-44 bg-gray-100 overflow-hidden border border-[#eee9db] rounded-xl shadow-xs group">
        <img src={gal.image_url} alt={gal.title} class="w-full h-full object-cover group-hover:scale-105 transition duration-300" />
      </div>
    {/each}
  </div>

  <!-- Link Menuju Halaman Galeri Lengkap -->
  <div class="mt-5 text-center">
    <a 
      href="/galeri" 
      class="inline-flex items-center gap-2 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2 px-4 rounded-lg hover:bg-[#f9f4e8] transition duration-200 group"
    >
      <span>Lihat Galeri Lainnya</span>
      <svg class="w-4 h-4 text-[#F08300] transform group-hover:translate-x-1.5 transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13.5 4.5L21 12m0 0l-7.5 7.5M21 12H3" />
      </svg>
    </a>
  </div>
</section>
