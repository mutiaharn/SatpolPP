<script lang="ts">
  import { galleryStore } from '$lib/stores/appState';
  import type { GalleryItem } from '$lib/types';

  let selectedCategory = '';
  let activePreview: GalleryItem | null = null;

  const categories = [
    { id: '', label: 'Semua Dokumentasi' },
    { id: 'Penertiban', label: 'Operasi & Penertiban' },
    { id: 'Apel & Upacara', label: 'Apel & Peringatan Hari Besar' },
    { id: 'Koordinasi', label: 'Koordinasi & Sosialisasi' }
  ];

  $: filteredItems = selectedCategory 
    ? $galleryStore.filter(item => item.category === selectedCategory)
    : $galleryStore;
</script>

<svelte:head>
  <title>Galeri Dokumentasi Kegiatan | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Dokumentasi Multimedia</div>
    <h1 class="text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Galeri Kegiatan &amp; Penugasan</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Rekam jejak visual pengabdian personel Satpol PP Provinsi Sulawesi Selatan dalam mengawal penegakan Peraturan Daerah, menjaga ketenteraman masyarakat, serta apel kesiapsiagaan korps Praja Wibawa.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-8">
  <!-- Filter Tabs -->
  <div class="flex flex-wrap gap-2">
    {#each categories as cat}
      <button 
        type="button"
        class={`px-4 py-2 text-xs font-bold uppercase transition ${selectedCategory === cat.id ? 'bg-[#F08300] text-white' : 'bg-white border border-gray-300 text-[#101116] hover:bg-gray-100'}`}
        on:click={() => selectedCategory = cat.id}
      >
        {cat.label}
      </button>
    {/each}
  </div>

  <!-- Gallery Grid -->
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
    {#each filteredItems as item}
      <button 
        type="button"
        class="bg-white border border-[#eee9db] rounded-xl overflow-hidden group text-left focus:outline-none focus:ring-2 focus:ring-[#F08300] shadow-sm hover:shadow-md transition"
        on:click={() => activePreview = item}
      >
        <div class="relative h-56 bg-gray-100 overflow-hidden">
          <img 
            src={item.image_url} 
            alt={item.title} 
            class="w-full h-full object-cover group-hover:scale-105 transition duration-300"
            loading="lazy"
          />
          <span class="absolute top-3 left-3 bg-[#101116]/85 text-[#F08300] font-bold text-[10px] uppercase px-2.5 py-1 tracking-wider border-l-2 border-[#F08300] rounded-sm">
            {item.category}
          </span>
        </div>
        <div class="p-4 space-y-1">
          <div class="text-[10px] text-[#6f7775]">{item.date}</div>
          <h3 class="text-xs font-bold text-[#101116] group-hover:text-[#F08300] transition line-clamp-2 leading-snug">
            {item.title}
          </h3>
        </div>
      </button>
    {/each}
  </div>
</div>

<!-- Modal Lightbox -->
{#if activePreview}
  <div 
    class="fixed inset-0 z-50 bg-black/85 flex items-center justify-center p-4 backdrop-blur-sm"
    role="dialog"
    aria-modal="true"
  >
    <div 
      class="bg-white max-w-3xl w-full border border-gray-800 rounded-2xl overflow-hidden shadow-2xl relative"
    >
      <div class="bg-[#101116] text-white p-4 flex justify-between items-center border-b border-[#F08300]">
        <div>
          <span class="text-[10px] font-bold uppercase text-[#F08300] tracking-wider">{activePreview.category} &bull; {activePreview.date}</span>
          <h4 class="text-sm font-bold text-white mt-0.5">{activePreview.title}</h4>
        </div>
        <button 
          on:click={() => activePreview = null} 
          class="text-gray-400 hover:text-white font-bold text-lg px-2 rounded-md"
          aria-label="Tutup preview"
        >
          &times;
        </button>
      </div>
      <div class="bg-black max-h-[70vh] flex items-center justify-center">
        <img 
          src={activePreview.image_url} 
          alt={activePreview.title} 
          class="max-h-[70vh] w-auto max-w-full object-contain"
        />
      </div>
      <div class="p-4 bg-[#f9f4e8] border-t border-[#eee9db] flex justify-between items-center text-xs">
        <span class="text-[#6f7775]">Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan</span>
        <button 
          on:click={() => activePreview = null}
          class="bg-[#3d261a] text-white px-4 py-1.5 font-bold uppercase text-xs rounded-lg hover:bg-[#271810] transition"
        >
          Tutup
        </button>
      </div>
    </div>
  </div>
{/if}
