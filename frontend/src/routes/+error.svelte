<script lang="ts">
  import { page } from '$app/stores';
</script>

<svelte:head>
  <title>Halaman Tidak Ditemukan | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="py-16 px-4 max-w-4xl mx-auto text-center space-y-6">
  <div class="inline-block bg-[#101116] text-[#F08300] font-mono font-bold text-xs uppercase px-4 py-1.5 border-b-2 border-[#F08300] tracking-widest">
    Status Kode: {$page.status}
  </div>

  <h1 class="text-3xl md:text-5xl font-extrabold text-[#101116] tracking-tight">
    {#if $page.status === 404}
      Halaman Tidak Ditemukan
    {:else}
      Terjadi Kendala Pada Server
    {/if}
  </h1>

  <div class="w-16 h-1 bg-[#F08300] mx-auto"></div>

  <p class="text-xs md:text-sm text-[#6f7775] max-w-lg mx-auto leading-relaxed">
    {#if $page.status === 404}
      Halaman atau dokumen yang Anda tuju tidak dapat ditemukan pada server resmi Satpol PP Provinsi Sulawesi Selatan. Kemungkinan tautan telah diperbarui atau alamat URL salah ketik.
    {:else}
      {$page.error?.message || 'Sistem sedang mengalami pemeliharaan rutin. Silakan coba beberapa saat lagi.'}
    {/if}
  </p>

  <div class="pt-4 flex flex-wrap justify-center gap-3">
    <a 
      href="/" 
      class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-6 py-2.5 transition"
    >
      Kembali ke Beranda
    </a>
    <a 
      href="/berita" 
      class="bg-[#3d261a] hover:bg-[#271810] text-white text-xs font-bold uppercase px-6 py-2.5 rounded-lg transition"
    >
      Warta Berita
    </a>
    <a 
      href="/lapor" 
      class="bg-white border border-gray-300 hover:bg-gray-100 text-[#101116] text-xs font-bold uppercase px-6 py-2.5 transition"
    >
      Lapor SIAP TERTIB
    </a>
  </div>
</div>
