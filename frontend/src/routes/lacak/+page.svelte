<script lang="ts">
  import { page } from '$app/stores';
  import TicketTracker from '$lib/components/TicketTracker.svelte';
  import { Search, ArrowLeft } from 'lucide-svelte';

  $: ticketQuery = $page.url.searchParams.get('ticket') || '';
</script>

<svelte:head>
  <title>Lacak Status Tiket Pengaduan | SIAP TERTIB Sulsel</title>
</svelte:head>

<div class="bg-gradient-to-b from-[#070e1c] to-[#0f223d] text-white py-12 px-4 relative overflow-hidden ds-specular-dark">
  <div class="absolute top-0 right-10 w-96 h-96 rounded-full bg-[#246bfd]/10 blur-3xl pointer-events-none"></div>
  <div class="max-w-4xl mx-auto relative z-10">
    <div class="flex items-center gap-2 text-xs text-slate-300 mb-3">
      <a href="/" class="hover:text-[#F08300] font-semibold inline-flex items-center gap-1.5 transition">
        <ArrowLeft class="w-3.5 h-3.5" />
        <span>Kembali ke Beranda Portal</span>
      </a>
    </div>
    <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-blue-500/20 text-[#246bfd] font-bold text-xs uppercase tracking-widest border border-blue-500/30 mb-2">
      <Search class="w-3.5 h-3.5 text-[#246bfd]" />
      <span>Transparansi Penanganan Aduan</span>
    </div>
    <h1 class="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">Pelacakan Tiket SIAP TERTIB</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-2 max-w-2xl leading-relaxed font-light">
      Pantau tindak lanjut laporan pelanggaran ketertiban umum Anda secara berkala. Masukkan nomor registrasi tiket aduan (contoh: TRB-202611-SAMPLE) untuk melihat kronologi penanganan oleh personel Satpol PP.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-4xl mx-auto space-y-6">
  <div class="ds-glass-strong rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-glass">
    <TicketTracker initialTicket={ticketQuery} />
  </div>

  <div class="text-center pt-4">
    <a href="/lapor" class="text-xs font-bold uppercase text-[#F08300] hover:text-[#d97400] underline inline-flex items-center gap-1">
      <span>Ingin mengajukan pengaduan pelanggaran baru? Klik formulir SIAP TERTIB &rarr;</span>
    </a>
  </div>
</div>
