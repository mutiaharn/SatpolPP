<script lang="ts">
  import { onMount } from 'svelte';
  import { articlesStore, galleryStore } from '$lib/stores/appState';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';
  import { Button, GlassCard, Badge } from '$lib/components/ui';
  import { 
    ShieldCheck, 
    ArrowRight, 
    FileText, 
    CheckCircle2, 
    Sparkles, 
    Clock, 
    ChevronRight,
    ChevronDown,
    Users,
    Shield,
    FileCheck2
  } from 'lucide-svelte';

  let currentSlide = 0;
  let showLightbox = false;
  let selectedGalleryIndex = 0;

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
</script>

<!-- 1. Hero Banner Slider (Layar Penuh dengan Glassmorphism & Sovereign Accent) -->
<section class="relative bg-[#070e1c] text-white min-h-[calc(100vh-96px)] flex items-center overflow-hidden">
  {#each slides as slide, idx}
    <div
      class={`absolute inset-0 bg-cover bg-center transition-opacity duration-1000 ${currentSlide === idx ? 'opacity-100 z-10' : 'opacity-0 z-0'}`}
      style={`background-image: linear-gradient(135deg, rgba(7, 14, 28, 0.80) 0%, rgba(11, 23, 42, 0.88) 100%), url('${slide.image}')`}
    ></div>
  {/each}

  <!-- Ambient Light Orbs underneath Hero -->
  <div class="absolute bottom-1/4 -right-20 w-96 h-96 rounded-full bg-[#f08300]/15 blur-3xl -z-0 pointer-events-none transform-gpu"></div>

  <div class="max-w-7xl mx-auto px-4 py-16 md:py-24 relative z-20 w-full">
    <div class="max-w-3xl space-y-6">
      <!-- Floating Glass Tag Pill -->
      <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full ds-glass-dark text-[#F08300] font-bold text-xs uppercase tracking-widest border border-white/20 shadow-glass">
        <span>{slides[currentSlide].tag}</span>
        <span class="text-white/40">/</span>
        <span class="text-slate-300 font-medium lowercase">prov. sulawesi selatan</span>
      </div>

      <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-extrabold tracking-tight leading-[1.14] text-white">
        {slides[currentSlide].title}
      </h1>

      <p class="text-sm md:text-base lg:text-lg text-slate-200 leading-relaxed max-w-2xl font-light">
        {slides[currentSlide].desc}
      </p>

      <div class="pt-2 flex flex-wrap items-center gap-4">
        <Button
          href="/lapor"
          variant="gold"
          size="lg"
          className="uppercase tracking-wider px-7"
        >
          <span>Lapor SIAP TERTIB</span>
          <ArrowRight class="w-4 h-4" />
        </Button>
        <Button 
          href="/dokumen" 
          variant="glass"
          size="lg"
          className="uppercase tracking-wider px-7 border-white/25 text-white hover:bg-white/15"
        >
          <FileText class="w-4 h-4 text-slate-300" />
          <span>Dokumen Publik</span>
        </Button>
      </div>
    </div>
  </div>

  <!-- Scroll Down Indicator (Apple HIG Frosted Liquid Glass Capsule - Zero Dodge Stable Centering) -->
  <button
    type="button"
    on:click={() => scrollToTarget('berita-terbaru')}
    class="hidden md:inline-flex items-center gap-2.5 absolute bottom-8 left-0 right-0 mx-auto w-max z-30 px-5 py-2.5 rounded-full bg-slate-900/75 hover:bg-slate-900/90 active:scale-[0.97] backdrop-blur-xl border border-white/20 hover:border-[#F08300]/50 text-white hover:text-[#F08300] shadow-[0_8px_24px_rgba(0,0,0,0.3)] transition-all duration-200 select-none cursor-pointer focus:outline-none focus-visible:ring-2 focus-visible:ring-[#F08300] focus-visible:ring-offset-2 focus-visible:ring-offset-slate-950 group"
    aria-label="Jelajahi Portal - Gulir ke konten berita"
  >
    <span class="text-xs font-bold tracking-wider uppercase text-white/95 group-hover:text-[#F08300] transition-colors">Jelajahi Portal</span>
    <ChevronDown class="w-4 h-4 text-[#F08300] group-hover:translate-y-0.5 transition-transform duration-200 shrink-0" />
  </button>

  <!-- Slide Navigation Indicators (Frosted Glass Dock - Apple HIG Ergonomic Touch Targets) -->
  <div class="absolute bottom-8 right-8 z-30 flex items-center gap-1 ds-glass-dark border border-white/20 p-1.5 rounded-full shadow-lg">
    {#each slides as _, i}
      <button
        type="button"
        class="h-8 min-w-[32px] px-1.5 flex items-center justify-center rounded-full transition-all focus:outline-none focus-visible:ring-2 focus-visible:ring-[#F08300] ds-press cursor-pointer"
        on:click={() => currentSlide = i}
        aria-label={`Pindah ke slide ${i + 1}`}
      >
        <span class={`h-2 transition-all duration-300 rounded-full ${currentSlide === i ? 'w-7 bg-[#F08300] shadow-[0_0_8px_rgba(240,131,0,0.6)]' : 'w-2 bg-white/40 hover:bg-white/80'}`}></span>
      </button>
    {/each}
  </div>
</section>

<!-- 2. Berita Terbaru (Glassmorphic Cards, Squircles & Balanced Vertical Centering) -->
<section id="berita-terbaru" class="scroll-mt-20 md:scroll-mt-24 py-8 md:py-12 px-4 max-w-7xl mx-auto md:min-h-[calc(100vh-96px)] flex flex-col justify-center">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-6 md:mb-8 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full ds-glass text-[#F08300] font-bold text-[11px] uppercase tracking-widest border border-amber-200/80 mb-2">
        <Sparkles class="w-3.5 h-3.5 text-[#F08300]" />
        <span>Publikasi &amp; Kegiatan</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Berita &amp; Informasi Terkini
      </h2>
      <p class="text-xs md:text-sm text-slate-500 mt-1 max-w-xl font-light">
        Informasi pelaksanaan penertiban Perda, ketenteraman masyarakat, dan kegiatan jajaran Satpol PP Sulsel.
      </p>
    </div>

    <Button
      href="/berita"
      variant="secondary"
      size="sm"
      className="uppercase tracking-wider shrink-0 text-[#F08300] hover:text-[#d97400] border-orange-200/80 hover:bg-orange-50/80"
    >
      <span>Lihat Semua Berita</span>
      <ChevronRight class="w-4 h-4 text-[#F08300]" />
    </Button>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-5 lg:gap-7 items-stretch">
    {#each $articlesStore.slice(0, 3) as item}
      <GlassCard hoverLift padding="none" className="h-full flex flex-col justify-between group overflow-hidden">
        <div>
          <div class="relative h-44 sm:h-48 bg-slate-100 overflow-hidden">
            <img
              src={item.thumbnail}
              alt={item.title}
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
              onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
            />
            <div class="absolute top-3 right-3 ds-glass-dark text-white text-[11px] font-bold px-3 py-1 rounded-full border border-white/25 shadow-xs">
              {new Date(item.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', year: 'numeric' })}
            </div>
            <div class="absolute bottom-3 left-3 bg-[#F08300] text-white text-[10px] font-bold uppercase tracking-wider px-2.5 py-1 rounded-lg shadow-xs">
              {item.category}
            </div>
          </div>

          <div class="p-5 space-y-2">
            <h3 class="font-bold text-[15px] sm:text-base text-[#071936] leading-snug line-clamp-2 min-h-[2.6rem] group-hover:text-[#F08300] transition duration-200">
              <a href={`/berita/${item.slug}`}>{item.title}</a>
            </h3>
            <p class="text-xs text-slate-500 leading-relaxed line-clamp-2 min-h-[2.2rem] font-light">
              {item.summary}
            </p>
          </div>
        </div>

        <div class="px-5 py-3 border-t border-slate-100/90 mt-auto flex items-center justify-between">
          <a href={`/berita/${item.slug}`} class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline flex items-center gap-1.5 group-hover:text-[#d97400] transition">
            <span>Baca Selengkapnya</span>
            <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
          </a>
          <span class="text-[11px] text-slate-400 flex items-center gap-1 font-mono">
            <Clock class="w-3.5 h-3.5" />
            <span>2 mnt</span>
          </span>
        </div>
      </GlassCard>
    {/each}
  </div>
</section>

<!-- 3. Sambutan & Amanat Pimpinan (Seamless Stage) -->
<section class="py-16 md:py-24 px-4 max-w-7xl mx-auto relative">
  <!-- Ambient Halo Behind Sambutan Section -->
  <div class="absolute bottom-0 right-0 w-[480px] h-[480px] rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none transform-gpu"></div>

  <div class="grid grid-cols-1 lg:grid-cols-12 gap-10 lg:gap-14 items-center relative z-10">
    <!-- Pimpinan Cutout Frame -->
    <div class="lg:col-span-5 flex justify-center items-end relative">
      <div class="relative flex flex-col items-center group w-full max-w-md">
        <!-- Subtle Frosted Halo Arch -->
        <div class="absolute bottom-0 left-1/2 -translate-x-1/2 w-[280px] sm:w-[340px] h-[440px] sm:h-[520px] rounded-t-full bg-gradient-to-t from-slate-900/10 via-amber-500/10 to-transparent border-t border-x border-white/70 -z-10 shadow-lg backdrop-blur-md"></div>

        <div class="relative z-10 transition-transform duration-500 ease-out group-hover:-translate-y-1 group-hover:scale-[1.01]">
          <img
            src="/uploads/foto_pimpinan/pimpinan_cutout_hd.png"
            alt="Andi Arwin Azis S.STP., M.M"
            class="w-auto h-[420px] sm:h-[480px] md:h-[530px] lg:h-[560px] object-contain drop-shadow-[0_15px_30px_rgba(15,23,42,0.18)] select-none pointer-events-none"
          />

          <!-- Floating Identity Card -->
          <div class="absolute bottom-4 left-1/2 -translate-x-1/2 w-[92%] sm:w-max max-w-md ds-glass px-5 py-3.5 rounded-2xl border border-white/90 shadow-glass text-center ds-specular-top">
            <div class="text-[10px] font-bold uppercase tracking-widest text-[#F08300] mb-0.5">
              Pejabat Pimpinan Tinggi Pratama
            </div>
            <div class="font-extrabold text-sm sm:text-base tracking-tight text-[#071936] uppercase whitespace-nowrap">
              ANDI ARWIN AZIS S.STP., M.M
            </div>
            <div class="text-[11px] text-slate-500 font-medium tracking-wide mt-0.5 whitespace-nowrap">
              Kepala Satuan Polisi Pamong Praja Prov. Sulsel
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Sambutan Pimpinan (Vocal Point) -->
    <div class="lg:col-span-7 space-y-6">
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full ds-glass text-[#F08300] font-bold text-xs uppercase tracking-widest border border-amber-200/80">
        <span>Pesan &amp; Sambutan Kasatpol PP</span>
      </div>

      <!-- Grand Focal Point Quote -->
      <div class="relative">
        <span class="text-6xl sm:text-7xl font-serif text-[#F08300]/25 select-none leading-none absolute -top-8 -left-3 pointer-events-none">&ldquo;</span>
        <h2 class="text-2xl sm:text-3xl md:text-4xl font-extrabold text-[#071936] tracking-tight leading-tight relative z-10">
          Menegakkan Perda Secara Humanis, Persuasif, dan Berkeadilan.
        </h2>
      </div>

      <p class="text-sm md:text-base text-slate-600 leading-relaxed font-light">
        Bagi kami, ketenteraman masyarakat bukan sekadar kepatuhan pada teks aturan, melainkan hadirnya rasa aman, tertib, dan terlindungi di setiap sudut Sulawesi Selatan.
      </p>

      <p class="text-sm md:text-base text-slate-600 leading-relaxed font-light">
        Aparatur Satpol PP hadir di tengah warga bukan untuk ditakuti, melainkan untuk mengayomi, mengedukasi, dan menegakkan wibawa daerah dengan ketulusan hati serta integritas yang teguh.
      </p>

      <!-- CTA Buttons (Apple HIG Ergonomics) -->
      <div class="pt-2 flex flex-wrap items-center gap-3.5">
        <Button 
          href="/profil/visi-misi" 
          variant="primary"
          size="lg"
          className="uppercase tracking-wider px-6"
        >
          <span>Visi &amp; Misi Instansi</span>
          <ArrowRight class="w-4 h-4 text-[#F08300]" />
        </Button>
        <Button 
          href="/profil/struktur" 
          variant="secondary"
          size="lg"
          className="uppercase tracking-wider px-6"
        >
          <span>Struktur Kelembagaan</span>
        </Button>
      </div>
    </div>
  </div>
</section>

<!-- 4. Dokumen Publik (Sovereign Glass Cards) -->
<section id="dokumen-publik" class="py-16 md:py-20 px-4 max-w-7xl mx-auto">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-10 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full ds-glass text-[#F08300] font-bold text-xs uppercase tracking-widest border border-amber-200/80 mb-2.5">
        <FileCheck2 class="w-3.5 h-3.5 text-[#F08300]" />
        <span>Keterbukaan Informasi Publik</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Dokumen &amp; Regulasi Daerah
      </h2>
      <p class="text-xs md:text-sm text-slate-500 mt-1 max-w-xl font-light">
        Akses terbuka ke dokumen perencanaan anggaran, laporan akuntabilitas kinerja, serta regulasi Perda Sulsel.
      </p>
    </div>

    <Button
      href="/dokumen"
      variant="secondary"
      size="sm"
      className="uppercase tracking-wider shrink-0 text-[#F08300] hover:text-[#d97400] border-orange-200/80 hover:bg-orange-50/80"
    >
      <span>Lihat Arsip Dokumen</span>
      <ChevronRight class="w-4 h-4 text-[#F08300]" />
    </Button>
  </div>

  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 items-stretch">
    <!-- Renja -->
    <GlassCard hoverLift className="flex flex-col justify-between h-full group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <Badge variant="gold" size="sm">Perencanaan</Badge>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">Renja (Rencana Kerja)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Dokumen perencanaan tahunan program dan target kegiatan Satpol PP Prov. Sulsel.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=renja" class="text-xs font-bold text-[#F08300] uppercase tracking-wider flex items-center gap-1 group-hover:text-[#d97400] py-1.5 px-2.5 -ml-2.5 rounded-lg hover:bg-orange-50/70 transition ds-press">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </GlassCard>

    <!-- Renstra -->
    <GlassCard hoverLift className="flex flex-col justify-between h-full group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <Badge variant="neutral" size="sm">Strategis</Badge>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">Renstra (Rencana Strategis)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Rencana strategis jangka menengah pencapaian visi, misi, dan sasaran instansi.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=renstra" class="text-xs font-bold text-[#F08300] uppercase tracking-wider flex items-center gap-1 group-hover:text-[#d97400] py-1.5 px-2.5 -ml-2.5 rounded-lg hover:bg-orange-50/70 transition ds-press">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </GlassCard>

    <!-- DPA -->
    <GlassCard hoverLift className="flex flex-col justify-between h-full group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <Badge variant="gold" size="sm">Penganggaran</Badge>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">DPA (Pelaksanaan Anggaran)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Alokasi anggaran belanja dan rincian pembiayaan kegiatan operasional.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=dpa" class="text-xs font-bold text-[#F08300] uppercase tracking-wider flex items-center gap-1 group-hover:text-[#d97400] py-1.5 px-2.5 -ml-2.5 rounded-lg hover:bg-orange-50/70 transition ds-press">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </GlassCard>

    <!-- LAKIP -->
    <GlassCard hoverLift className="flex flex-col justify-between h-full group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <Badge variant="neutral" size="sm">Kinerja</Badge>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">LAKIP / LKjIP</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Laporan akuntabilitas kinerja instansi pemerintah tahunan.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=lakip" class="text-xs font-bold text-[#F08300] uppercase tracking-wider flex items-center gap-1 group-hover:text-[#d97400] py-1.5 px-2.5 -ml-2.5 rounded-lg hover:bg-orange-50/70 transition ds-press">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </GlassCard>

    <!-- Laporan Keuangan -->
    <GlassCard hoverLift className="flex flex-col justify-between h-full group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <Badge variant="neutral" size="sm">Pertanggungjawaban</Badge>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">Laporan Keuangan</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Laporan pertanggungjawaban realisasi keuangan dan aset tahunan.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=keuangan" class="text-xs font-bold text-[#F08300] uppercase tracking-wider flex items-center gap-1 group-hover:text-[#d97400] py-1.5 px-2.5 -ml-2.5 rounded-lg hover:bg-orange-50/70 transition ds-press">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </GlassCard>

    <!-- Perda & SOP -->
    <GlassCard hoverLift className="flex flex-col justify-between h-full group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <Badge variant="gold" size="sm">Regulasi</Badge>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">Perda &amp; SOP Satpol PP</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Peraturan Daerah Provinsi Sulawesi Selatan dan Standar Operasional Prosedur.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=perda" class="text-xs font-bold text-[#F08300] uppercase tracking-wider flex items-center gap-1 group-hover:text-[#d97400] py-1.5 px-2.5 -ml-2.5 rounded-lg hover:bg-orange-50/70 transition ds-press">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </GlassCard>
  </div>
</section>

<!-- 6. Galeri Kegiatan (Sovereign Glass Grid) -->
<section id="galeri-kegiatan" class="pt-8 pb-16 px-4 max-w-7xl mx-auto">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-8 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full ds-glass text-[#F08300] font-bold text-xs uppercase tracking-widest border border-amber-200/80 mb-2.5">
        <Users class="w-3.5 h-3.5 text-[#F08300]" />
        <span>Dokumentasi Lapangan</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Galeri Operasional &amp; Kegiatan
      </h2>
    </div>

    <Button
      href="/galeri"
      variant="secondary"
      size="sm"
      className="uppercase tracking-wider shrink-0 text-[#F08300] hover:text-[#d97400] border-orange-200/80 hover:bg-orange-50/80"
    >
      <span>Lihat Semua Foto</span>
      <ChevronRight class="w-4 h-4 text-[#F08300]" />
    </Button>
  </div>

  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    {#each previewGallery as gal, idx}
      <button
        type="button"
        on:click={() => openGalleryModal(idx)}
        class="relative h-44 md:h-48 bg-slate-100 overflow-hidden rounded-2xl shadow-glass group text-left focus:outline-none focus:ring-2 focus:ring-[#F08300] border border-white/80 ds-hover-lift ds-press"
        aria-label={`Perbesar foto ${gal.title}`}
      >
        <img src={gal.image_url} alt={gal.title} class="w-full h-full object-cover group-hover:scale-110 transition duration-500" />
        <div class="absolute inset-0 bg-gradient-to-t from-[#070e1c]/88 via-[#070e1c]/30 to-transparent opacity-0 group-hover:opacity-100 transition duration-300 flex items-end p-4 backdrop-blur-xs">
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
