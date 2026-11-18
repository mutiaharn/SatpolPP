<script lang="ts">
  import { page } from '$app/stores';
</script>

<svelte:head>
  <title>Halaman Tidak Ditemukan | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="py-20 px-4 max-w-3xl mx-auto text-center space-y-6">
  <div class="inline-flex items-center gap-2 bg-[#0b172a] text-[#F08300] font-mono font-bold text-xs uppercase px-4 py-1.5 rounded-full border border-white/10 tracking-widest shadow-xs">
    Status Kode: {$page.status}
  </div>

  <h1 class="text-3xl md:text-5xl font-extrabold text-slate-900 tracking-tight">
    {#if $page.status === 404}
      Halaman Tidak Ditemukan
    {:else}
      Terjadi Kendala Pada Server
    {/if}
  </h1>

  <div class="w-16 h-1 bg-[#F08300] mx-auto rounded-full"></div>

  <p class="text-xs md:text-sm text-slate-500 max-w-lg mx-auto leading-relaxed font-light">
    {#if $page.status === 404}
      Halaman atau dokumen yang Anda tuju tidak dapat ditemukan pada server resmi Satpol PP Provinsi Sulawesi Selatan. Kemungkinan tautan telah diperbarui atau alamat URL salah ketik.
    {:else}
      {$page.error?.message || 'Sistem sedang mengalami pemeliharaan rutin. Silakan coba beberapa saat lagi.'}
    {/if}
  </p>

  <div class="pt-6 flex flex-wrap justify-center gap-3">
    <a 
      href="/" 
      class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-6 py-3 rounded-xl transition shadow-glass-gold"
    >
      Kembali ke Beranda
    </a>
    <a 
      href="/berita" 
      class="bg-[#0b172a] hover:bg-[#1e3a5f] text-white text-xs font-bold uppercase px-6 py-3 rounded-xl transition shadow-xs"
    >
      Warta Berita
    </a>
    <a 
      href="/lapor" 
      class="bg-white border border-slate-200 hover:bg-slate-100 text-slate-800 text-xs font-bold uppercase px-6 py-3 rounded-xl transition shadow-2xs"
    >
      Lapor SIAP TERTIB
    </a>
  </div>
</div>
