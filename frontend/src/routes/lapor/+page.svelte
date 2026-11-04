<script lang="ts">
  import ComplaintForm from '$lib/components/ComplaintForm.svelte';
</script>

<svelte:head>
  <title>Lapor Pelanggaran SIAP TERTIB | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-4xl mx-auto">
    <div class="flex items-center gap-2 text-xs text-[#6f7775] mb-3">
      <a href="/" class="hover:text-[#F08300] font-semibold">&larr; Kembali ke Beranda Portal</a>
    </div>
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Kanal Partisipasi Warga Sulsel</div>
    <h1 class="text-3xl font-extrabold text-[#101116] tracking-tight mt-1">SIAP TERTIB &bull; Pengaduan Ketertiban Umum</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl leading-relaxed">
      Sistem Informasi Pengaduan Pelanggaran Peraturan Daerah, Ketenteraman, Ketertiban Umum, dan Pelindungan Masyarakat Provinsi Sulawesi Selatan. Setiap laporan yang masuk akan ditindaklanjuti sesuai Standar Operasional Prosedur (SOP).
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-4xl mx-auto space-y-8">
  <!-- SOP Flow Banner -->
  <div class="bg-white border border-[#eee9db] p-5 shadow-sm">
    <div class="text-xs font-bold text-[#101116] uppercase tracking-wider mb-3">Alur &amp; Prosedur Penanganan Laporan:</div>
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-3 text-xs">
      <div class="border-l-2 border-[#F08300] pl-3 py-1 bg-gray-50/50">
        <div class="font-bold text-[#101116]">1. Verifikasi</div>
        <div class="text-[11px] text-[#6f7775] mt-0.5">Pemeriksaan bukti foto &amp; validitas aduan warga.</div>
      </div>
      <div class="border-l-2 border-[#2b2f3f] pl-3 py-1 bg-gray-50/50">
        <div class="font-bold text-[#101116]">2. Disposisi</div>
        <div class="text-[11px] text-[#6f7775] mt-0.5">Penugasan ke Bidang Trantibum atau Perda.</div>
      </div>
      <div class="border-l-2 border-[#F08300] pl-3 py-1 bg-gray-50/50">
        <div class="font-bold text-[#101116]">3. Penindakan</div>
        <div class="text-[11px] text-[#6f7775] mt-0.5">Giat patroli lapangan &amp; surat peringatan resmi.</div>
      </div>
      <div class="border-l-2 border-emerald-600 pl-3 py-1 bg-gray-50/50">
        <div class="font-bold text-emerald-800">4. Berita Acara</div>
        <div class="text-[11px] text-[#6f7775] mt-0.5">Penyusunan BAP penertiban &amp; penutupan tiket.</div>
      </div>
    </div>
  </div>

  <ComplaintForm />

  <div class="text-center pt-2">
    <a href="/lacak" class="text-xs font-bold uppercase text-[#F08300] hover:text-[#d97400] underline">
      Sudah memiliki nomor resi aduan? Klik di sini untuk melacak status laporan Anda &rarr;
    </a>
  </div>
</div>
