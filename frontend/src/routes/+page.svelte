<script lang="ts">
  import { onMount } from 'svelte';
  import { articlesStore, galleryStore } from '$lib/stores/appState';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';
  import { 
    ShieldCheck, 
    ArrowRight, 
    FileText, 
    CheckCircle2, 
    Sparkles, 
    Clock, 
    ChevronRight,
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
  <div class="absolute top-1/4 -left-20 w-96 h-96 rounded-full bg-[#246bfd]/18 blur-3xl -z-0 pointer-events-none transform-gpu"></div>
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
        <a
          href="/lapor"
          class="inline-flex items-center gap-2 bg-gradient-to-r from-[#F08300] to-[#e07000] hover:from-[#d97400] hover:to-[#c66500] text-white font-bold text-xs uppercase px-7 py-4 rounded-xl transition duration-200 shadow-glass-gold hover:shadow-xl hover:-translate-y-0.5 transform ds-press"
        >
          <span>Lapor SIAP TERTIB</span>
          <ArrowRight class="w-4 h-4" />
        </a>
        <a 
          href="/dokumen" 
          class="inline-flex items-center gap-2 ds-glass-dark hover:bg-white/15 text-white font-bold text-xs uppercase px-7 py-4 rounded-xl border border-white/25 transition duration-200 hover:-translate-y-0.5 transform ds-press"
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
    class="hidden md:flex absolute bottom-8 left-1/2 -translate-x-1/2 z-30 flex-col items-center gap-1.5 text-slate-400 hover:text-[#F08300] transition group focus:outline-none cursor-pointer ds-press"
    aria-label="Gulir ke konten berita"
  >
    <span class="text-[10px] font-bold uppercase tracking-widest text-[#F08300]">Jelajahi Portal</span>
    <svg class="w-5 h-5 animate-bounce text-[#F08300]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14l-7 7m0 0l-7-7m7 7V3" />
    </svg>
  </a>

  <!-- Slide Navigation Indicators (Frosted Glass Dock) -->
  <div class="absolute bottom-8 right-8 z-30 flex items-center gap-2 ds-glass-dark border border-white/20 px-3.5 py-2 rounded-full shadow-lg">
    {#each slides as _, i}
      <button
        type="button"
        class={`h-2 transition-all duration-300 rounded-full ${currentSlide === i ? 'w-8 bg-[#F08300]' : 'w-2 bg-white/40 hover:bg-white/80'}`}
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
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full ds-glass text-[#246bfd] font-bold text-xs uppercase tracking-widest border border-blue-200/80 mb-2.5">
        <Sparkles class="w-3.5 h-3.5 text-[#246bfd]" />
        <span>Publikasi &amp; Kegiatan</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Berita &amp; Informasi Terkini
      </h2>
      <p class="text-xs md:text-sm text-slate-500 mt-1 max-w-xl font-light">
        Informasi pelaksanaan penertiban Perda, ketenteraman masyarakat, dan kegiatan jajaran Satpol PP Sulsel.
      </p>
    </div>

    <a
      href="/berita"
      class="inline-flex items-center gap-2 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2.5 px-4 rounded-xl ds-glass border border-orange-200/80 hover:bg-orange-50/80 transition duration-200 group shrink-0 ds-press"
    >
      <span>Lihat Semua Berita</span>
      <ChevronRight class="w-4 h-4 text-[#F08300] group-hover:translate-x-1 transition-transform" />
    </a>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-8 items-stretch">
    {#each $articlesStore.slice(0, 3) as item}
      <article class="ds-glass-card rounded-2xl overflow-hidden flex flex-col justify-between h-full group">
        <div>
          <div class="relative h-52 bg-slate-100 overflow-hidden">
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

          <div class="p-6 space-y-3">
            <h3 class="font-bold text-base text-[#071936] leading-snug line-clamp-2 min-h-[2.8rem] group-hover:text-[#246bfd] transition duration-200">
              <a href={`/berita/${item.slug}`}>{item.title}</a>
            </h3>
            <p class="text-xs text-slate-500 leading-relaxed line-clamp-3 min-h-[3.2rem] font-light">
              {item.summary}
            </p>
          </div>
        </div>

        <div class="px-6 pb-6 pt-3 border-t border-slate-100/90 mt-auto flex items-center justify-between">
          <a href={`/berita/${item.slug}`} class="text-xs font-bold text-[#246bfd] uppercase tracking-wider hover:underline flex items-center gap-1.5 group-hover:text-[#1d4ed8] transition">
            <span>Baca Selengkapnya</span>
            <ArrowRight class="w-3.5 h-3.5 group-hover:translate-x-0.5 transition-transform" />
          </a>
          <span class="text-[11px] text-slate-400 flex items-center gap-1 font-mono">
            <Clock class="w-3.5 h-3.5" />
            <span>2 mnt</span>
          </span>
        </div>
      </article>
    {/each}
  </div>
</section>

<!-- 3. SIAP TERTIB: editorial call-to-action -->
<section id="siap-tertib" class="relative isolate overflow-hidden px-4 py-10 md:py-14">
  <div class="absolute inset-0 bg-gradient-to-b from-[#eef1f4]/90 via-transparent to-[#eef1f4]/70 pointer-events-none"></div>
  <div class="relative max-w-7xl mx-auto overflow-hidden rounded-[2rem] bg-[#070e1c] text-white shadow-[0_28px_80px_rgba(7,14,28,0.20)]">
    <div class="absolute -right-20 -top-28 h-80 w-80 rounded-full border border-white/10"></div>
    <div class="absolute -right-5 -top-12 h-52 w-52 rounded-full border border-[#F08300]/25"></div>
    <div class="absolute -bottom-28 left-1/3 h-72 w-72 rounded-full bg-slate-400/10 blur-3xl pointer-events-none"></div>
    <div class="absolute right-8 bottom-0 hidden select-none text-[7rem] font-black leading-none tracking-tighter text-white/[0.035] lg:block">SIAP</div>

    <div class="relative grid grid-cols-1 lg:grid-cols-12 gap-9 lg:gap-12 px-7 py-9 sm:px-10 md:px-14 md:py-14">
      <div class="lg:col-span-7 max-w-2xl">
        <div class="inline-flex items-center gap-2 text-[#F6B15C] text-[10px] sm:text-xs font-bold uppercase tracking-[0.2em]">
          <span class="h-px w-8 bg-[#F08300]"></span>
          <span>Ruang Pengaduan Masyarakat</span>
        </div>
        <h2 class="mt-5 text-3xl sm:text-4xl md:text-5xl font-extrabold leading-[1.08] tracking-tight">
          Ada gangguan ketertiban di sekitar Anda?
        </h2>
        <p class="mt-5 max-w-xl text-sm md:text-base leading-relaxed text-slate-300">
          Sampaikan laporan melalui SIAP TERTIB. Simpan kode tiket Anda dan pantau proses tindak lanjutnya secara mandiri.
        </p>
        <div class="mt-8 flex flex-wrap items-center gap-3">
          <a href="/lapor" class="inline-flex items-center gap-2 rounded-xl bg-gradient-to-r from-[#F08300] to-[#d97400] px-6 py-3.5 text-xs font-bold uppercase tracking-wider text-white shadow-[0_10px_28px_rgba(240,131,0,0.25)] transition hover:-translate-y-0.5 hover:shadow-xl ds-press">
            Buat Laporan
            <ArrowRight class="h-4 w-4" />
          </a>
          <a href="/lacak" class="inline-flex items-center gap-2 rounded-xl border border-white/20 bg-white/[0.06] px-6 py-3.5 text-xs font-bold uppercase tracking-wider text-white transition hover:bg-white/[0.12] ds-press">
            Lacak Tiket
            <ChevronRight class="h-4 w-4 text-[#F6B15C]" />
          </a>
        </div>
      </div>

      <div class="lg:col-span-5 flex flex-col justify-center gap-5 border-t border-white/15 pt-7 lg:border-l lg:border-t-0 lg:pl-10 lg:pt-0">
        <div class="flex items-start gap-4">
          <span class="font-mono text-xl font-bold text-[#F6B15C]">01</span>
          <div>
            <h3 class="text-sm font-bold text-white">Kirim laporan</h3>
            <p class="mt-1 text-xs leading-relaxed text-slate-400">Jelaskan kejadian dan lengkapi informasi lokasi.</p>
          </div>
        </div>
        <div class="ml-2 h-5 w-px bg-gradient-to-b from-[#F08300]/70 to-white/10"></div>
        <div class="flex items-start gap-4">
          <span class="font-mono text-xl font-bold text-[#F6B15C]">02</span>
          <div>
            <h3 class="text-sm font-bold text-white">Simpan kode tiket</h3>
            <p class="mt-1 text-xs leading-relaxed text-slate-400">Gunakan kode tiket untuk membuka kembali laporan.</p>
          </div>
        </div>
        <div class="ml-2 h-5 w-px bg-gradient-to-b from-[#F08300]/70 to-white/10"></div>
        <div class="flex items-start gap-4">
          <span class="font-mono text-xl font-bold text-[#F6B15C]">03</span>
          <div>
            <h3 class="text-sm font-bold text-white">Pantau tindak lanjut</h3>
            <p class="mt-1 text-xs leading-relaxed text-slate-400">Periksa perkembangan penanganan melalui halaman pelacakan.</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</section>

<!-- 4. Sambutan & Amanat Pimpinan (Leadership Feature Stage) -->
<section class="relative overflow-hidden py-16 md:py-24 px-4">
  <!-- Ambient Halo Behind Sambutan Section -->
  <div class="absolute top-1/2 -left-20 -translate-y-1/2 w-[520px] h-[520px] rounded-full bg-slate-300/35 blur-3xl pointer-events-none transform-gpu"></div>
  <div class="absolute bottom-0 right-0 w-[480px] h-[480px] rounded-full bg-[#F08300]/12 blur-3xl pointer-events-none transform-gpu"></div>

  <div class="max-w-7xl mx-auto grid grid-cols-1 lg:grid-cols-12 gap-8 lg:gap-12 items-center relative z-10 rounded-[2rem] overflow-hidden bg-[#070e1c] px-5 py-10 sm:px-8 md:px-12 md:py-14 shadow-[0_28px_80px_rgba(7,14,28,0.18)] border border-white/10">
    <!-- Pimpinan Cutout Frame -->
    <div class="lg:col-span-5 flex justify-center items-end relative">
      <div class="relative flex flex-col items-center group w-full max-w-md">
        <!-- Subtle Frosted Halo Arch -->
        <div class="absolute bottom-0 left-1/2 -translate-x-1/2 w-[280px] sm:w-[340px] h-[440px] sm:h-[520px] rounded-t-full bg-gradient-to-t from-slate-400/25 via-[#F08300]/12 to-transparent border-t border-x border-white/20 -z-10 shadow-lg backdrop-blur-md"></div>

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
        <h2 class="text-2xl sm:text-3xl md:text-4xl lg:text-[2.75rem] font-extrabold text-white tracking-tight leading-[1.12] relative z-10">
          Menegakkan Perda Secara Humanis, Persuasif, dan Berkeadilan.
        </h2>
      </div>

      <p class="text-sm md:text-base text-slate-300 leading-relaxed font-light">
        Bagi kami, ketenteraman masyarakat bukan sekadar kepatuhan pada teks aturan, melainkan hadirnya rasa aman, tertib, dan terlindungi di setiap sudut Sulawesi Selatan.
      </p>

      <p class="text-sm md:text-base text-slate-300 leading-relaxed font-light">
        Aparatur Satpol PP hadir di tengah warga bukan untuk ditakuti, melainkan untuk mengayomi, mengedukasi, dan menegakkan wibawa daerah dengan ketulusan hati serta integritas yang teguh.
      </p>

      <!-- CTA Buttons -->
      <div class="pt-2 flex flex-wrap items-center gap-3.5">
        <a 
          href="/profil/visi-misi" 
          class="inline-flex items-center gap-2 bg-gradient-to-r from-[#F08300] to-[#d97400] hover:from-[#d97400] hover:to-[#c66500] text-white text-xs font-bold uppercase px-6 py-3.5 rounded-xl transition duration-200 shadow-[0_10px_28px_rgba(240,131,0,0.22)] hover:-translate-y-0.5 ds-press"
        >
          <span>Visi &amp; Misi Instansi</span>
          <ArrowRight class="w-4 h-4 text-[#F08300]" />
        </a>
        <a 
          href="/profil/struktur" 
          class="inline-flex items-center gap-2 bg-white/[0.06] hover:bg-white/[0.12] text-white text-xs font-bold uppercase px-6 py-3.5 rounded-xl border border-white/20 transition duration-200 ds-press"
        >
          <span>Struktur Kelembagaan</span>
        </a>
      </div>
    </div>
  </div>
</section>

<!-- 5. Dokumen Publik (Sovereign Glass Cards) -->
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

    <a
      href="/dokumen"
      class="inline-flex items-center gap-2 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2.5 px-4 rounded-xl ds-glass border border-orange-200/80 hover:bg-orange-50/80 transition duration-200 group shrink-0 ds-press"
    >
      <span>Lihat Arsip Dokumen</span>
      <ChevronRight class="w-4 h-4 text-[#F08300] group-hover:translate-x-1 transition-transform" />
    </a>
  </div>

  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 items-stretch">
    <!-- Renja -->
    <div class="ds-glass-card rounded-2xl p-6 flex flex-col justify-between h-full border border-white/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#246bfd] uppercase tracking-wider bg-blue-50/90 px-2.5 py-1 rounded-md border border-blue-100">Perencanaan</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#246bfd] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#246bfd] transition duration-200">Renja (Rencana Kerja)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Dokumen perencanaan tahunan program dan target kegiatan Satpol PP Prov. Sulsel.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=renja" class="text-xs font-bold text-[#246bfd] uppercase tracking-wider hover:underline flex items-center gap-1 group-hover:text-[#1d4ed8]">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </div>

    <!-- Renstra -->
    <div class="ds-glass-card rounded-2xl p-6 flex flex-col justify-between h-full border border-white/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#246bfd] uppercase tracking-wider bg-blue-50/90 px-2.5 py-1 rounded-md border border-blue-100">Strategis</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#246bfd] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#246bfd] transition duration-200">Renstra (Rencana Strategis)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Rencana strategis jangka menengah pencapaian visi, misi, dan sasaran instansi.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=renstra" class="text-xs font-bold text-[#246bfd] uppercase tracking-wider hover:underline flex items-center gap-1 group-hover:text-[#1d4ed8]">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </div>

    <!-- DPA -->
    <div class="ds-glass-card rounded-2xl p-6 flex flex-col justify-between h-full border border-white/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#F08300] uppercase tracking-wider bg-orange-50/90 px-2.5 py-1 rounded-md border border-orange-100">Penganggaran</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">DPA (Pelaksanaan Anggaran)</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Alokasi anggaran belanja dan rincian pembiayaan kegiatan operasional.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=dpa" class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </div>

    <!-- LAKIP -->
    <div class="ds-glass-card rounded-2xl p-6 flex flex-col justify-between h-full border border-white/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-emerald-600 uppercase tracking-wider bg-emerald-50/90 px-2.5 py-1 rounded-md border border-emerald-100">Kinerja</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-emerald-600 transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-emerald-600 transition duration-200">LAKIP / LKjIP</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Laporan akuntabilitas kinerja instansi pemerintah tahunan.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=lakip" class="text-xs font-bold text-emerald-600 uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </div>

    <!-- Laporan Keuangan -->
    <div class="ds-glass-card rounded-2xl p-6 flex flex-col justify-between h-full border border-white/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-purple-600 uppercase tracking-wider bg-purple-50/90 px-2.5 py-1 rounded-md border border-purple-100">Pertanggungjawaban</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-purple-600 transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-purple-600 transition duration-200">Laporan Keuangan</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Laporan pertanggungjawaban realisasi keuangan dan aset tahunan.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=keuangan" class="text-xs font-bold text-purple-600 uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </div>

    <!-- Perda & SOP -->
    <div class="ds-glass-card rounded-2xl p-6 flex flex-col justify-between h-full border border-white/80 group">
      <div>
        <div class="flex items-center justify-between mb-3">
          <span class="text-[10px] font-bold text-[#F08300] uppercase tracking-wider bg-orange-50/90 px-2.5 py-1 rounded-md border border-orange-100">Regulasi</span>
          <FileText class="w-4 h-4 text-slate-400 group-hover:text-[#F08300] transition duration-200" />
        </div>
        <h3 class="font-bold text-base text-[#071936] group-hover:text-[#F08300] transition duration-200">Perda &amp; SOP Satpol PP</h3>
        <p class="text-xs text-slate-500 mt-2 line-clamp-2 font-light">Peraturan Daerah Provinsi Sulawesi Selatan dan Standar Operasional Prosedur.</p>
      </div>
      <div class="mt-5 pt-3.5 border-t border-slate-100/90 flex items-center justify-between">
        <a href="/dokumen?category=perda" class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline flex items-center gap-1">
          <span>Akses Dokumen</span>
          <ArrowRight class="w-3.5 h-3.5" />
        </a>
        <span class="text-[11px] text-slate-400 font-mono">PDF Resmi</span>
      </div>
    </div>
  </div>
</section>

<!-- 6. Galeri Kegiatan (Sovereign Glass Grid) -->
<section id="galeri-kegiatan" class="pt-8 pb-16 px-4 max-w-7xl mx-auto">
  <div class="flex flex-col md:flex-row md:items-end justify-between mb-8 gap-4">
    <div>
      <div class="inline-flex items-center gap-2 px-3.5 py-1.5 rounded-full ds-glass text-[#246bfd] font-bold text-xs uppercase tracking-widest border border-blue-200/80 mb-2.5">
        <Users class="w-3.5 h-3.5 text-[#246bfd]" />
        <span>Dokumentasi Lapangan</span>
      </div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#071936] tracking-tight">
        Galeri Operasional &amp; Kegiatan
      </h2>
    </div>

    <a
      href="/galeri"
      class="inline-flex items-center gap-2 text-xs font-bold text-[#F08300] hover:text-[#d97400] uppercase tracking-wider py-2.5 px-4 rounded-xl ds-glass border border-orange-200/80 hover:bg-orange-50/80 transition duration-200 group shrink-0 ds-press"
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
