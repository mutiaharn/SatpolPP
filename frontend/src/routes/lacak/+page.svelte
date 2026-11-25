<script lang="ts">
  import { page } from '$app/stores';
  import TicketTracker from '$lib/components/TicketTracker.svelte';
  import { Search, ChevronRight } from 'lucide-svelte';

  $: ticketQuery = $page.url.searchParams.get('ticket') || '';
</script>

<svelte:head>
  <title>Lacak Status Tiket Pengaduan | SIAP TERTIB Sulsel</title>
</svelte:head>

<!-- Sovereign Glass Tracking Console (Full Viewport, Zero Scroll) -->
<div class="flex-1 flex flex-col justify-center items-center px-4 py-8 sm:py-12 bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white relative overflow-hidden ds-specular-dark">
  <!-- Ambient Orbs -->
  <div class="absolute top-0 right-10 w-96 h-96 rounded-full bg-[#F08300]/5 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-72 h-72 rounded-full bg-[#F08300]/10 blur-3xl pointer-events-none"></div>

  <div class="max-w-2xl w-full mx-auto relative z-10 text-center space-y-4 my-auto">
    <!-- Breadcrumb Pills -->
    <div class="flex flex-wrap items-center justify-center gap-2 text-xs text-slate-300">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-400" />
      <span class="text-slate-400">Layanan Warga</span>
      <ChevronRight class="w-3.5 h-3.5 text-slate-400" />
      <span class="text-[#F08300] font-semibold">Pelacakan Tiket</span>
    </div>

    <!-- Capsule Badge -->
    <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-amber-500/15 text-[#F08300] font-bold text-[11px] uppercase tracking-wider border border-amber-500/30">
      <Search class="w-3.5 h-3.5 text-[#F08300]" />
      <span>Transparansi Penanganan Aduan</span>
    </div>

    <!-- Heading & Humanized Subtitle -->
    <h1 class="text-2xl sm:text-3xl md:text-4xl font-extrabold text-white tracking-tight">
      Pelacakan Tiket Pengaduan
    </h1>
    <p class="text-xs sm:text-sm text-slate-300 max-w-lg mx-auto font-light leading-relaxed">
      Masukkan nomor resi tiket pengaduan untuk memantau status verifikasi dan tindak lanjut penanganan lapangan secara waktu-nyata.
    </p>

    <!-- Search Form Console -->
    <div class="pt-2 text-left">
      <TicketTracker initialTicket={ticketQuery} showHeader={false} />
    </div>
  </div>
</div>
