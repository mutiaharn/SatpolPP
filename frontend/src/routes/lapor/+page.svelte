<script lang="ts">
  import ComplaintForm from '$lib/components/ComplaintForm.svelte';
  import { Shield, ArrowLeft, CheckCircle2 } from 'lucide-svelte';
</script>

<svelte:head>
  <title>Lapor Pelanggaran SIAP TERTIB | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-gradient-to-b from-[#070e1c] to-[#0f223d] text-white py-12 px-4 relative overflow-hidden ds-specular-dark">
  <div class="absolute top-0 right-10 w-96 h-96 rounded-full bg-[#f08300]/10 blur-3xl pointer-events-none"></div>
  <div class="max-w-4xl mx-auto relative z-10">
    <div class="flex items-center gap-2 text-xs text-slate-300 mb-3">
      <a href="/" class="hover:text-[#F08300] font-semibold inline-flex items-center gap-1.5 transition">
        <ArrowLeft class="w-3.5 h-3.5" />
        <span>Kembali ke Beranda Portal</span>
      </a>
    </div>
    <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-orange-500/20 text-[#F08300] font-bold text-xs uppercase tracking-widest border border-orange-500/30 mb-2">
      <Shield class="w-3.5 h-3.5 text-[#F08300]" />
      <span>Kanal Partisipasi Warga Sulsel</span>
    </div>
    <h1 class="text-3xl sm:text-4xl font-extrabold text-white tracking-tight">SIAP TERTIB &bull; Pengaduan Ketertiban Umum</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-2 max-w-2xl leading-relaxed font-light">
      Sistem Informasi Pengaduan Pelanggaran Peraturan Daerah, Ketenteraman, Ketertiban Umum, dan Pelindungan Masyarakat Provinsi Sulawesi Selatan. Setiap laporan yang masuk akan ditindaklanjuti sesuai Standar Operasional Prosedur (SOP).
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-4xl mx-auto space-y-8">
  <!-- SOP Flow Banner (Tamalate Glass Style) -->
  <div class="ds-glass rounded-2xl p-6 border border-slate-200/80 shadow-glass">
    <div class="text-xs font-bold text-[#071936] uppercase tracking-wider mb-4 flex items-center justify-between">
      <span>Alur &amp; Prosedur Penanganan Laporan:</span>
      <span class="text-[10px] font-mono text-emerald-600 font-semibold bg-emerald-50 px-2.5 py-0.5 rounded-full border border-emerald-200">SOP Resmi</span>
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-3 text-xs">
      <div class="border-l-3 border-[#246bfd] pl-3 py-1.5 bg-blue-50/50 rounded-r-xl">
        <div class="font-bold text-[#071936]">1. Verifikasi</div>
        <div class="text-[11px] text-slate-500 mt-0.5">Pemeriksaan bukti foto &amp; validitas aduan warga.</div>
      </div>
      <div class="border-l-3 border-[#F08300] pl-3 py-1.5 bg-orange-50/50 rounded-r-xl">
        <div class="font-bold text-[#071936]">2. Disposisi</div>
        <div class="text-[11px] text-slate-500 mt-0.5">Penugasan ke Bidang Trantibum atau Perda.</div>
      </div>
      <div class="border-l-3 border-amber-500 pl-3 py-1.5 bg-amber-50/50 rounded-r-xl">
        <div class="font-bold text-[#071936]">3. Penindakan</div>
        <div class="text-[11px] text-slate-500 mt-0.5">Giat patroli lapangan &amp; surat peringatan resmi.</div>
      </div>
      <div class="border-l-3 border-emerald-600 pl-3 py-1.5 bg-emerald-50/50 rounded-r-xl">
        <div class="font-bold text-emerald-800">4. Berita Acara</div>
        <div class="text-[11px] text-slate-500 mt-0.5">Penyusunan BAP penertiban &amp; penutupan tiket.</div>
      </div>
    </div>
  </div>

  <div class="ds-glass-strong rounded-3xl p-6 sm:p-8 border border-slate-200 shadow-glass-hover">
    <ComplaintForm />
  </div>

  <div class="text-center pt-2">
    <a href="/lacak" class="text-xs font-bold uppercase text-[#F08300] hover:text-[#d97400] underline inline-flex items-center gap-1.5">
      <span>Sudah memiliki nomor resi aduan? Klik di sini untuk melacak status laporan Anda &rarr;</span>
    </a>
  </div>
</div>
