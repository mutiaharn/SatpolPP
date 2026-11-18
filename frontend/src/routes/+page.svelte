<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import { articlesStore, galleryStore } from '$lib/stores/appState';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';
  import { 
    ShieldCheck, 
    ArrowRight, 
    FileText, 
    CheckCircle2, 
    Search, 
    Sparkles, 
    Clock, 
    Send, 
    ChevronRight,
    Users,
    Shield,
    FileCheck2
  } from 'lucide-svelte';

  let currentSlide = 0;
  let showLightbox = false;
  let selectedGalleryIndex = 0;
  let quickTrackCode = '';

  $: previewGallery = $galleryStore.slice(0, 4);

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
    const timer = setInterval(() => {
      currentSlide = (currentSlide + 1) % slides.length;
    }, 6000);
    return () => clearInterval(timer);
  });

  function scrollToTarget(id: string) {
    const el = document.getElementById(id);
    if (!el) return;
    const header = document.querySelector('header');
    const headerHeight = header ? Math.round(header.getBoundingClientRect().height) : 96;
    const elPos = el.getBoundingClientRect().top + window.pageYOffset;
    window.scrollTo({
      top: Math.max(0, Math.round(elPos - headerHeight)),
      behavior: 'smooth'
    });
  }

  function openGalleryModal(index: number) {
    selectedGalleryIndex = index;
    showLightbox = true;
  }

  function handleQuickTrack() {
    if (quickTrackCode.trim()) {
      goto(`/lacak?ticket=${encodeURIComponent(quickTrackCode.trim())}`);
    } else {
      goto('/lacak');
    }
  }
</script>

<!-- 1. Hero Banner Slider (Layar Penuh dengan Glassmorphism & Sovereign Accent) -->
<section class="relative bg-[#070e1c] text-white min-h-[calc(100vh-96px)] flex items-center overflow-hidden">
  {#each slides as slide, idx}
    <div
      class={`absolute inset-0 bg-cover bg-center transition-opacity duration-1000 ${currentSlide === idx ? 'opacity-100 z-10' : 'opacity-0 z-0'}`}
      style={`background-image: linear-gradient(135deg, rgba(7, 14, 28, 0.82) 0%, rgba(11, 23, 42, 0.90) 100%), url('${slide.image}')`}
    ></div>
  {/each}

  <!-- Ambient Light Orbs -->
  <div class="absolute top-1/4 -left-20 w-96 h-96 rounded-full bg-[#246bfd]/15 blur-3xl -z-0 pointer-events-none"></div>
  <div class="absolute bottom-1/4 -right-20 w-96 h-96 rounded-full bg-[#f08300]/12 blur-3xl -z-0 pointer-events-none"></div>

  <div class="max-w-7xl mx-auto px-4 py-16 md:py-24 relative z-20 w-full">
    <div class="max-w-3xl space-y-6">
      <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-white/15 backdrop-blur-md shadow-xs">
        <span class="w-2 h-2 rounded-full bg-[#F08300] animate-pulse"></span>
        <span>{slides[currentSlide].tag}</span>
        <span class="text-white/40">&bull;</span>
        <span class="text-slate-300 font-medium lowercase">prov. sulawesi selatan</span>
      </div>

      <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-extrabold tracking-tight leading-[1.12]">
        {slides[currentSlide].title}
      </h1>

      <p class="text-sm md:text-base lg:text-lg text-slate-300 leading-relaxed max-w-2xl font-light">
        {slides[currentSlide].desc}
      </p>

      <div class="pt-2 flex flex-wrap items-center gap-4">
        <a
          href="/lapor"
          class="inline-flex items-center gap-2 bg-gradient-to-r from-[#F08300] to-[#e07000] hover:from-[#d97400] hover:to-[#c66500] text-white font-bold text-xs uppercase px-7 py-4 rounded-xl transition shadow-glass-gold hover:shadow-xl hover:-translate-y-0.5 transform"
        >
          <span>Lapor SIAP TERTIB</span>
          <ArrowRight class="w-4 h-4" />
        </a>
        <a 
          href="/dokumen" 
          class="inline-flex items-center gap-2 bg-white/10 hover:bg-white/20 text-white font-bold text-xs uppercase px-7 py-4 rounded-xl border border-white/25 transition backdrop-blur-md hover:-translate-y-0.5 transform"
        >
          <FileText class="w-4 h-4 text-slate-300" />
          <span>Dokumen Publik</span>
        </a>
      </div>
    </div>
  </div>

  <!-- Scroll Down Indicator -->
  <a
    href="#berita-terbaru"
    on:click={(e) => { e.preventDefault(); scrollToTarget('berita-terbaru'); }}
    class="hidden md:flex absolute bottom-8 left-1/2 -translate-x-1/2 z-30 flex-col items-center gap-1.5 text-slate-400 hover:text-[#F08300] transition group focus:outline-none cursor-pointer"
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
        type="button"
        class={`h-2 transition-all rounded-full ${currentSlide === i ? 'w-8 bg-[#F08300]' : 'w-2 bg-white/30 hover:bg-white/60'}`}
        on:click={() => currentSlide = i}
        aria-label={`Slide ${i + 1}`}
      ></button>
    {/each}
  </div>
</section>

<!-- 2. Berita Terbaru (Glassmorphic Cards, Squircles & Balanced Heights) -->
<section id="berita-terbaru" class="scroll-mt-24 pt-12 md:pt-16 pb-16 px-4 max-w-7xl mx-auto">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-10 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-blue-50 text-[#246bfd] font-bold text-xs uppercase tracking-widest border border-blue-200/60 mb-2">
        <Sparkles class="w-3.5 h-3.5 text-[#246bfd]" />
        <span>Publikasi &amp; Kegiatan</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Berita &amp; Informasi Terkini
      </h2>
      <p class="text-xs md:text-sm text-slate-500 mt-1 max-w-xl">
        Informasi pelaksanaan penertiban Perda, ketenteraman masyarakat, dan kegiatan jajaran Satpol PP Sulsel.
      </p>
    </div>

    <a
      href="/berita"
      class="inline-flex items-center gap-1.5 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2.5 px-4 rounded-xl bg-orange-50/70 border border-orange-200/50 hover:bg-orange-100/60 transition group shrink-0"
    >
      <span>Lihat Semua Berita</span>
      <ChevronRight class="w-4 h-4 text-[#F08300] group-hover:translate-x-1 transition-transform" />
    </a>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-8 items-stretch">
    {#each $articlesStore.slice(0, 3) as item}
      <article class="ds-glass rounded-2xl overflow-hidden flex flex-col justify-between h-full hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 group border border-slate-200/80">
        <div>
          <div class="relative h-52 bg-slate-100 overflow-hidden">
            <img
              src={item.thumbnail}
              alt={item.title}
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
              onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
            />
            <div class="absolute top-3 right-3 bg-[#070e1c]/80 backdrop-blur-md text-white text-[11px] font-bold px-3 py-1 rounded-full border border-white/20 shadow-xs">
              {new Date(item.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })}
            </div>
            <div class="absolute bottom-3 left-3 bg-[#F08300] text-white text-[10px] font-bold uppercase tracking-wider px-2.5 py-1 rounded-md shadow-xs">
              {item.category}
            </div>
          </div>

          <div class="p-6 space-y-3">
            <h3 class="font-bold text-base text-[#071936] leading-snug line-clamp-2 min-h-[2.8rem] group-hover:text-[#246bfd] transition">
              <a href={`/berita/${item.slug}`}>{item.title}</a>
            </h3>
            <p class="text-xs text-slate-500 leading-relaxed line-clamp-3 min-h-[3.2rem]">
              {item.summary}
            </p>
          </div>
        </div>

        <div class="px-6 pb-6 pt-3 border-t border-slate-100/80 mt-auto flex items-center justify-between">
          <a href={`/berita/${item.slug}`} class="text-xs font-bold text-[#246bfd] uppercase tracking-wider hover:underline flex items-center gap-1.5">
            <span>Baca Selengkapnya</span>
            <ArrowRight class="w-3.5 h-3.5" />
          </a>
          <span class="text-[11px] text-slate-400 flex items-center gap-1">
            <Clock class="w-3.5 h-3.5" />
            <span>2 mnt baca</span>
          </span>
        </div>
      </article>
    {/each}
  </div>
</section>

<!-- 3. Profil Pimpinan Satpol PP Sulsel (Revisi 2.4: Sambutan Pimpinan Menonjol & Tidak Monoton) -->
<section class="py-16 md:py-20 px-4 bg-gradient-to-br from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white relative overflow-hidden ds-specular-dark">
  <!-- Sovereign Radiant Halo -->
  <div class="absolute top-1/2 left-1/4 -translate-y-1/2 -translate-x-1/2 w-[500px] h-[500px] rounded-full bg-[#246bfd]/15 blur-3xl pointer-events-none"></div>
  <div class="absolute top-1/2 right-10 -translate-y-1/2 w-[400px] h-[400px] rounded-full bg-[#F08300]/12 blur-3xl pointer-events-none"></div>

  <div class="max-w-7xl mx-auto grid grid-cols-1 lg:grid-cols-12 gap-8 lg:gap-14 items-center relative z-10">
    <!-- Pimpinan Cutout with Ambient Glass Frame -->
    <div class="lg:col-span-5 flex justify-center items-end relative">
      <div class="relative flex flex-col items-center group w-full max-w-md">
        <!-- Glow Backplate Arch -->
        <div class="absolute bottom-0 left-1/2 -translate-x-1/2 w-[310px] sm:w-[370px] h-[490px] sm:h-[570px] rounded-t-full bg-gradient-to-t from-white/10 via-[#246bfd]/15 to-transparent border-t border-x border-white/20 -z-10 shadow-2xl backdrop-blur-xs"></div>
        <div class="absolute -bottom-2 left-1/2 -translate-x-1/2 w-4/5 h-8 bg-black/50 rounded-[100%] blur-md -z-10"></div>

        <div class="relative z-10 transition-transform duration-500 ease-out group-hover:-translate-y-2 group-hover:scale-[1.01]">
          <img
            src="/uploads/foto_pimpinan/pimpinan_cutout_hd.png"
            alt="Andi Arwin Azis S.STP., M.M"
            class="w-auto h-[460px] sm:h-[540px] md:h-[600px] lg:h-[640px] object-contain drop-shadow-[0_25px_35px_rgba(0,0,0,0.55)] select-none pointer-events-none"
          />

          <!-- Floating Sovereign Identity Pill -->
          <div class="absolute bottom-5 left-1/2 -translate-x-1/2 w-[92%] sm:w-max max-w-md ds-glass-dark px-6 py-3.5 rounded-2xl border border-white/20 shadow-2xl text-center transform transition-all duration-300 group-hover:scale-105">
            <div class="flex items-center justify-center gap-1.5 text-[10px] font-bold uppercase tracking-wider text-[#F08300] mb-0.5">
              <ShieldCheck class="w-3.5 h-3.5 text-[#F08300]" />
              <span>Pejabat Pimpinan Tinggi Pratama</span>
            </div>
            <div class="font-extrabold text-sm sm:text-base tracking-wide text-white uppercase whitespace-nowrap">
              ANDI ARWIN AZIS S.STP., M.M
            </div>
            <div class="text-[11px] text-slate-300 font-medium tracking-wide mt-0.5 whitespace-nowrap">
              Kepala Satuan Polisi Pamong Praja Prov. Sulsel
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Amanat & Visi Strategis -->
    <div class="lg:col-span-7 space-y-6">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-white/10 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-white/15 backdrop-blur-md">
        <span class="w-2 h-2 rounded-full bg-[#F08300]"></span>
        <span>Amanat &amp; Komitmen Pimpinan</span>
      </div>

      <h2 class="text-2xl sm:text-3xl md:text-4xl font-extrabold text-white tracking-tight leading-tight">
        "Menegakkan Perda Secara Humanis, Persuasif, dan Berkeadilan"
      </h2>

      <p class="text-sm md:text-base text-slate-300 leading-relaxed font-light">
        Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan senantiasa mengedepankan pendekatan humanis, persuasif, dan profesional dalam menegakkan Peraturan Daerah dan menjaga ketertiban umum demi kesejahteraan masyarakat Sulawesi Selatan.
      </p>

      <!-- 4 Strategic Mission Cards (Tamalate Glass Style) -->
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-3.5 pt-2">
        <div class="bg-white/5 hover:bg-white/10 border border-white/10 rounded-xl p-4 transition backdrop-blur-xs">
          <div class="flex items-center gap-2.5 text-xs font-bold text-slate-200">
            <span class="w-2 h-2 rounded-full bg-[#246bfd]"></span>
            <span>Trantibum &amp; Ketenteraman</span>
          </div>
          <p class="text-[11px] text-slate-400 mt-1.5 leading-relaxed">
            Penyelenggaraan ketenteraman dan ketertiban umum di seluruh 24 kab/kota Sulawesi Selatan.
          </p>
        </div>

        <div class="bg-white/5 hover:bg-white/10 border border-white/10 rounded-xl p-4 transition backdrop-blur-xs">
          <div class="flex items-center gap-2.5 text-xs font-bold text-slate-200">
            <span class="w-2 h-2 rounded-full bg-[#F08300]"></span>
            <span>Penegakan Perda &amp; Perkada</span>
          </div>
          <p class="text-[11px] text-slate-400 mt-1.5 leading-relaxed">
            Pengawasan dan penindakan yuridis terhadap pelanggaran peraturan daerah secara profesional.
          </p>
        </div>

        <div class="bg-white/5 hover:bg-white/10 border border-white/10 rounded-xl p-4 transition backdrop-blur-xs">
          <div class="flex items-center gap-2.5 text-xs font-bold text-slate-200">
            <span class="w-2 h-2 rounded-full bg-emerald-400"></span>
            <span>Pengawalan Aset Daerah</span>
          </div>
          <p class="text-[11px] text-slate-400 mt-1.5 leading-relaxed">
            Penertiban dan pengamanan aset barang milik daerah demi pemanfaatan yang tertib dan berkelanjutan.
          </p>
        </div>

        <div class="bg-white/5 hover:bg-white/10 border border-white/10 rounded-xl p-4 transition backdrop-blur-xs">
          <div class="flex items-center gap-2.5 text-xs font-bold text-slate-200">
            <span class="w-2 h-2 rounded-full bg-purple-400"></span>
            <span>Pembinaan Satlinmas</span>
          </div>
          <p class="text-[11px] text-slate-400 mt-1.5 leading-relaxed">
            Bimbingan masyarakat dan koordinasi satuan perlindungan masyarakat di pelosok daerah.
          </p>
        </div>
      </div>

      <div class="pt-2 flex flex-wrap gap-4">
        <a 
          href="/profil/struktur" 
          class="inline-flex items-center gap-2 bg-[#246bfd] hover:bg-[#1a56db] text-white text-xs font-bold uppercase px-6 py-3.5 rounded-xl transition shadow-glass-azure hover:-translate-y-0.5"
        >
          <span>Struktur Organisasi</span>
          <ArrowRight class="w-4 h-4" />
        </a>
        <a 
          href="/profil/visi-misi" 
          class="inline-flex items-center gap-2 bg-white/10 hover:bg-white/15 text-white text-xs font-bold uppercase px-6 py-3.5 rounded-xl border border-white/20 transition"
        >
          <span>Visi &amp; Misi Instansi</span>
        </a>
      </div>
    </div>
  </div>
</section>

<!-- 4. Layanan SIAP TERTIB (Revisi 2.3: Call to Action Kreatif & Tidak Monoton) -->
<section id="layanan-siaptertib" class="scroll-mt-24 py-16 md:py-20 px-4 bg-[#f8fafc] border-b border-slate-200/80">
  <div class="max-w-7xl mx-auto">
    <!-- Living Workflow Hub Container -->
    <div class="ds-glass-strong rounded-3xl p-8 md:p-12 border border-slate-200 shadow-2xl relative overflow-hidden">
      <!-- Glow ambient in the card -->
      <div class="absolute -top-24 -right-24 w-80 h-80 rounded-full bg-[#f08300]/10 blur-3xl pointer-events-none"></div>
      <div class="absolute -bottom-24 -left-24 w-80 h-80 rounded-full bg-[#246bfd]/10 blur-3xl pointer-events-none"></div>

      <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 items-center relative z-10">
        <!-- Left: Interactive CTA & Quick Ticket Tracker -->
        <div class="lg:col-span-7 space-y-5">
          <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full bg-orange-100 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-orange-200">
            <span class="w-2 h-2 rounded-full bg-[#F08300] animate-pulse"></span>
            <span>Saluran Pengaduan Terpadu</span>
          </div>

          <h2 class="text-2xl sm:text-3xl md:text-4xl font-extrabold text-[#071936] tracking-tight leading-tight">
            SIAP TERTIB &bull; Ruang Partisipasi Warga Sulsel
          </h2>

          <p class="text-xs md:text-sm text-slate-600 leading-relaxed max-w-xl">
            Laporkan pelanggaran Peraturan Daerah, gangguan ketertiban umum, atau penyerobotan aset daerah secara resmi. Identitas pelapor dienkripsi otomatis, lengkap dengan resi pelacakan tiket transparan dan tembusan email.
          </p>

          <!-- Interactive Quick Ticket Input (Lacak Instan dari Beranda) -->
          <div class="bg-white p-2 rounded-2xl border border-slate-200 shadow-sm max-w-lg">
            <form on:submit|preventDefault={handleQuickTrack} class="flex items-center gap-2">
              <div class="pl-3 text-slate-400">
                <Search class="w-4 h-4" />
              </div>
              <input
                type="text"
                bind:value={quickTrackCode}
                placeholder="Masukkan Nomor Tiket (misal: ST-2026-0001)..."
                class="w-full text-xs sm:text-sm bg-transparent border-0 focus:outline-none focus:ring-0 text-slate-800 placeholder-slate-400"
              />
              <button
                type="submit"
                class="bg-[#070e1c] hover:bg-slate-800 text-white font-bold text-xs px-4 py-2.5 rounded-xl shrink-0 transition"
              >
                Cek Tiket
              </button>
            </form>
          </div>

          <!-- Direct Action Buttons -->
          <div class="flex flex-wrap items-center gap-3 pt-2">
            <a
              href="/lapor"
              class="inline-flex items-center gap-2 bg-gradient-to-r from-[#F08300] to-[#e07000] hover:from-[#d97400] hover:to-[#c66500] text-white font-bold text-xs uppercase px-6 py-3.5 rounded-xl transition shadow-glass-gold hover:-translate-y-0.5"
            >
              <Send class="w-4 h-4" />
              <span>Buat Laporan Baru</span>
            </a>
            <a
              href="/lacak"
              class="inline-flex items-center gap-2 bg-slate-100 hover:bg-slate-200/80 text-slate-700 font-bold text-xs uppercase px-6 py-3.5 rounded-xl border border-slate-200 transition"
            >
              <span>Panduan Pelaporan</span>
            </a>
          </div>
        </div>

        <!-- Right: Connected 4-Step Interactive Timeline -->
        <div class="lg:col-span-5 bg-gradient-to-b from-[#070e1c] to-[#0f223d] text-white rounded-2xl p-6 sm:p-7 shadow-xl border border-white/10 relative">
          <div class="flex items-center justify-between pb-4 border-b border-white/10 mb-5">
            <div class="flex items-center gap-2">
              <span class="w-2.5 h-2.5 rounded-full bg-emerald-400 animate-pulse"></span>
              <span class="text-xs font-bold uppercase tracking-wider text-slate-200">Alur Standar Tindak Lanjut</span>
            </div>
            <span class="text-[11px] font-mono text-[#F08300] font-semibold">SOP 24/7</span>
          </div>

          <div class="space-y-4 relative">
            <!-- Step 1 -->
            <div class="flex items-start gap-3.5">
              <div class="w-7 h-7 rounded-full bg-[#246bfd] text-white flex items-center justify-center font-bold text-xs shrink-0 shadow-xs mt-0.5">
                1
              </div>
              <div class="flex-1">
                <div class="text-xs font-bold text-white flex items-center justify-between">
                  <span>Verifikasi &amp; Masking</span>
                  <span class="text-[10px] text-slate-400 font-normal">&le; 2 Jam</span>
                </div>
                <p class="text-[11px] text-slate-300 mt-0.5">Pemeriksaan bukti foto dan enkripsi data pribadi pelapor.</p>
              </div>
            </div>

            <!-- Step 2 -->
            <div class="flex items-start gap-3.5">
              <div class="w-7 h-7 rounded-full bg-[#F08300] text-white flex items-center justify-center font-bold text-xs shrink-0 shadow-xs mt-0.5">
                2
              </div>
              <div class="flex-1">
                <div class="text-xs font-bold text-white flex items-center justify-between">
                  <span>Disposisi Cepat</span>
                  <span class="text-[10px] text-slate-400 font-normal">&le; 4 Jam</span>
                </div>
                <p class="text-[11px] text-slate-300 mt-0.5">Penerusan ke bidang teknis (Trantibum, Perda, atau Linmas).</p>
              </div>
            </div>

            <!-- Step 3 -->
            <div class="flex items-start gap-3.5">
              <div class="w-7 h-7 rounded-full bg-amber-500 text-white flex items-center justify-center font-bold text-xs shrink-0 shadow-xs mt-0.5">
                3
              </div>
              <div class="flex-1">
                <div class="text-xs font-bold text-white flex items-center justify-between">
                  <span>Patroli &amp; Penindakan</span>
                  <span class="text-[10px] text-slate-400 font-normal">SLA H+1</span>
                </div>
                <p class="text-[11px] text-slate-300 mt-0.5">Operasi lapangan persuasif &amp; penerbitan surat peringatan resmi.</p>
              </div>
            </div>

            <!-- Step 4 -->
            <div class="flex items-start gap-3.5">
              <div class="w-7 h-7 rounded-full bg-emerald-500 text-white flex items-center justify-center font-bold text-xs shrink-0 shadow-xs mt-0.5">
                4
              </div>
              <div class="flex-1">
                <div class="text-xs font-bold text-white flex items-center justify-between">
                  <span>Penyelesaian &amp; BAP</span>
                  <span class="text-[10px] text-emerald-400 font-semibold">Tuntas</span>
                </div>
                <p class="text-[11px] text-slate-300 mt-0.5">Dokumentasi hasil dan update status transparan ke email warga.</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- 5. Dokumen Publik (Sovereign Glass Cards) -->
<section id="dokumen-publik" class="py-16 md:py-20 px-4 max-w-7xl mx-auto">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-10 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-amber-50 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-amber-200/60 mb-2">
        <FileCheck2 class="w-3.5 h-3.5 text-[#F08300]" />
        <span>Keterbukaan Informasi Publik</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Dokumen &amp; Regulasi Daerah
      </h2>
      <p class="text-xs md:text-sm text-slate-500 mt-1 max-w-xl">
        Akses terbuka ke dokumen perencanaan anggaran, laporan akuntabilitas kinerja, serta regulasi Perda Sulsel.
      </p>
    </div>

    <a
      href="/dokumen"
      class="inline-flex items-center gap-1.5 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2.5 px-4 rounded-xl bg-orange-50/70 border border-orange-200/50 hover:bg-orange-100/60 transition group shrink-0"
    >
      <span>Lihat Arsip Dokumen</span>
      <ChevronRight class="w-4 h-4 text-[#F08300] group-hover:translate-x-1 transition-transform" />
    </a>
  </div>

  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 items-stretch">
    <!-- Renja -->
    <div class="ds-glass rounded-2xl p-6 hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 flex flex-col justify-between h-full border border-slate-200/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#246bfd] uppercase tracking-wider bg-blue-50 px-2.5 py-1 rounded-md border border-blue-100">Perencanaan</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#246bfd] transition" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#246bfd] transition">Renja (Rencana Kerja)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2">Dokumen perencanaan tahunan program dan target kegiatan Satpol PP Prov. Sulsel.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100 flex items-center justify-between">
        <a href="/dokumen?category=renja" class="text-xs font-bold text-[#246bfd] uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400">PDF Resmi</span>
      </div>
    </div>

    <!-- Renstra -->
    <div class="ds-glass rounded-2xl p-6 hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 flex flex-col justify-between h-full border border-slate-200/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#246bfd] uppercase tracking-wider bg-blue-50 px-2.5 py-1 rounded-md border border-blue-100">Strategis</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#246bfd] transition" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#246bfd] transition">Renstra (Rencana Strategis)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2">Rencana strategis jangka menengah pencapaian visi, misi, dan sasaran instansi.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100 flex items-center justify-between">
        <a href="/dokumen?category=renstra" class="text-xs font-bold text-[#246bfd] uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400">PDF Resmi</span>
      </div>
    </div>

    <!-- DPA -->
    <div class="ds-glass rounded-2xl p-6 hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 flex flex-col justify-between h-full border border-slate-200/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#F08300] uppercase tracking-wider bg-orange-50 px-2.5 py-1 rounded-md border border-orange-100">Penganggaran</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition">DPA (Pelaksanaan Anggaran)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2">Alokasi anggaran belanja dan rincian pembiayaan kegiatan operasional.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100 flex items-center justify-between">
        <a href="/dokumen?category=dpa" class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400">PDF Resmi</span>
      </div>
    </div>

    <!-- LAKIP -->
    <div class="ds-glass rounded-2xl p-6 hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 flex flex-col justify-between h-full border border-slate-200/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-emerald-600 uppercase tracking-wider bg-emerald-50 px-2.5 py-1 rounded-md border border-emerald-100">Kinerja</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-emerald-600 transition" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-emerald-600 transition">LAKIP / LKjIP</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2">Laporan akuntabilitas kinerja instansi pemerintah tahunan.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100 flex items-center justify-between">
        <a href="/dokumen?category=lakip" class="text-xs font-bold text-emerald-600 uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400">PDF Resmi</span>
      </div>
    </div>

    <!-- Laporan Keuangan -->
    <div class="ds-glass rounded-2xl p-6 hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 flex flex-col justify-between h-full border border-slate-200/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-purple-600 uppercase tracking-wider bg-purple-50 px-2.5 py-1 rounded-md border border-purple-100">Pertanggungjawaban</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-purple-600 transition" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-purple-600 transition">Laporan Keuangan</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2">Laporan pertanggungjawaban realisasi keuangan dan aset tahunan.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100 flex items-center justify-between">
        <a href="/dokumen?category=keuangan" class="text-xs font-bold text-purple-600 uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400">PDF Resmi</span>
      </div>
    </div>

    <!-- Perda & SOP -->
    <div class="ds-glass rounded-2xl p-6 hover:shadow-glass-hover hover:-translate-y-1 transition duration-300 flex flex-col justify-between h-full border border-slate-200/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#F08300] uppercase tracking-wider bg-orange-50 px-2.5 py-1 rounded-md border border-orange-100">Regulasi</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition">Perda &amp; SOP Satpol PP</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2">Peraturan Daerah Provinsi Sulawesi Selatan dan Standar Operasional Prosedur.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100 flex items-center justify-between">
        <a href="/dokumen?category=perda" class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400">PDF Resmi</span>
      </div>
    </div>
  </div>
</section>

<!-- 6. Galeri Kegiatan (Sovereign Glass Grid) -->
<section id="galeri-kegiatan" class="pt-8 pb-16 px-4 max-w-7xl mx-auto">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-8 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-blue-50 text-[#246bfd] font-bold text-xs uppercase tracking-widest border border-blue-200/60 mb-2">
        <Users class="w-3.5 h-3.5 text-[#246bfd]" />
        <span>Dokumentasi Lapangan</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Galeri Operasional &amp; Kegiatan
      </h2>
    </div>

    <a
      href="/galeri"
      class="inline-flex items-center gap-1.5 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2.5 px-4 rounded-xl bg-orange-50/70 border border-orange-200/50 hover:bg-orange-100/60 transition group shrink-0"
    >
      <span>Lihat Semua Foto</span>
      <ChevronRight class="w-4 h-4 text-[#F08300] group-hover:translate-x-1 transition-transform" />
    </a>
  </div>

  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    {#each previewGallery as gal, idx}
      <button
        type="button"
        on:click={() => openGalleryModal(idx)}
        class="relative h-44 md:h-48 bg-slate-100 overflow-hidden rounded-2xl shadow-xs group text-left focus:outline-none focus:ring-2 focus:ring-[#F08300] border border-slate-200"
        aria-label={`Perbesar foto ${gal.title}`}
      >
        <img src={gal.image_url} alt={gal.title} class="w-full h-full object-cover group-hover:scale-110 transition duration-500" />
        <div class="absolute inset-0 bg-gradient-to-t from-[#070e1c]/85 via-[#070e1c]/20 to-transparent opacity-0 group-hover:opacity-100 transition duration-300 flex items-end p-4">
          <span class="text-xs font-bold text-white line-clamp-2">{gal.title}</span>
        </div>
      </button>
    {/each}
  </div>
</section>

<Lightbox
  bind:open={showLightbox}
  items={previewGallery}
  currentIndex={selectedGalleryIndex}
  on:change={(e) => selectedGalleryIndex = e.detail}
  on:close={() => showLightbox = false}
/>
