<script lang="ts">
  import ComplaintForm from '$lib/components/ComplaintForm.svelte';
  import { Shield, ArrowLeft, CheckCircle2, ChevronRight } from 'lucide-svelte';
</script>

<svelte:head>
  <title>Lapor Pelanggaran SIAP TERTIB | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#F08300]/5 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <!-- Breadcrumb Pills -->
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-slate-400">Layanan Warga</span>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">Pengaduan SIAP TERTIB</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">SIAP TERTIB &mdash; Pengaduan Masyarakat</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      Saluran resmi bagi masyarakat Sulawesi Selatan untuk melaporkan gangguan ketertiban umum dan aset daerah secara aman, terenkripsi, dan terpantau secara transparan.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-4xl mx-auto space-y-8">
  <!-- SOP Flow Banner (Tamalate Specular Glass Style) -->
  <div class="ds-glass rounded-3xl p-6 border border-white/90 shadow-glass ds-specular-top">
    <div class="text-xs font-bold text-[#071936] uppercase tracking-wider mb-4 flex items-center justify-between">
      <span>Alur &amp; Prosedur Penanganan Laporan:</span>
      <span class="text-[10px] font-mono text-[#d97400] font-semibold bg-amber-500/10 px-3 py-1 rounded-full border border-amber-500/25">SOP Resmi</span>
    </div>
    <div class="grid grid-cols-1 sm:grid-cols-4 gap-3 text-xs">
      <div class="border-l-3 border-[#F08300] pl-3 py-2 ds-glass-subtle rounded-r-xl border border-white/60">
        <div class="font-bold text-[#071936]">1. Verifikasi</div>
        <div class="text-[11px] text-slate-500 mt-0.5 font-light">Pemeriksaan bukti foto &amp; validitas aduan warga.</div>
      </div>
      <div class="border-l-3 border-[#F08300] pl-3 py-2 ds-glass-subtle rounded-r-xl border border-white/60">
        <div class="font-bold text-[#071936]">2. Disposisi</div>
        <div class="text-[11px] text-slate-500 mt-0.5 font-light">Penugasan ke Bidang Trantibum atau Perda.</div>
      </div>
      <div class="border-l-3 border-amber-500 pl-3 py-2 ds-glass-subtle rounded-r-xl border border-white/60">
        <div class="font-bold text-[#071936]">3. Penindakan</div>
        <div class="text-[11px] text-slate-500 mt-0.5 font-light">Giat patroli lapangan &amp; surat peringatan resmi.</div>
      </div>
      <div class="border-l-3 border-emerald-600 pl-3 py-2 ds-glass-subtle rounded-r-xl border border-white/60">
        <div class="font-bold text-emerald-800">4. Berita Acara</div>
        <div class="text-[11px] text-slate-500 mt-0.5 font-light">Penyusunan BAP penertiban &amp; penutupan tiket.</div>
      </div>
    </div>
  </div>

  <ComplaintForm />

  <div class="text-center pt-2">
    <a href="/lacak" class="text-xs font-bold uppercase text-[#F08300] hover:text-[#d97400] underline inline-flex items-center gap-1.5 ds-press">
      <span>Sudah memiliki nomor resi aduan? Klik di sini untuk melacak status laporan Anda &rarr;</span>
    </a>
  </div>
</div>
