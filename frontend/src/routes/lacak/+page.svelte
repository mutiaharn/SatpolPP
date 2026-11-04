<script lang="ts">
  import { page } from '$app/stores';
  import TicketTracker from '$lib/components/TicketTracker.svelte';

  $: ticketQuery = $page.url.searchParams.get('ticket') || '';
</script>

<svelte:head>
  <title>Lacak Status Tiket Pengaduan | SIAP TERTIB Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-4xl mx-auto">
    <div class="flex items-center gap-2 text-xs text-[#6f7775] mb-3">
      <a href="/" class="hover:text-[#F08300] font-semibold">&larr; Kembali ke Beranda Portal</a>
    </div>
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Transparansi Penanganan Aduan</div>
    <h1 class="text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Pelacakan Tiket SIAP TERTIB</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl leading-relaxed">
      Pantau tindak lanjut laporan pelanggaran ketertiban umum Anda secara berkala. Masukkan nomor registrasi tiket aduan (contoh: TRB-202611-SAMPLE) untuk melihat kronologi penanganan oleh personel Satpol PP.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-4xl mx-auto space-y-6">
  <TicketTracker initialTicket={ticketQuery} />

  <div class="text-center pt-4">
    <a href="/lapor" class="text-xs font-bold uppercase text-[#F08300] hover:text-[#d97400] underline">
      Ingin mengajukan pengaduan pelanggaran baru? Klik formulir SIAP TERTIB &rarr;
    </a>
  </div>
</div>
