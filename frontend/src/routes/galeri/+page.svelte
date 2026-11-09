<script lang="ts">
  import { galleryStore } from '$lib/stores/appState';
  import type { GalleryItem } from '$lib/types';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';

  let selectedCategory = '';
  let showLightbox = false;
  let selectedIndex = 0;

  const categories = [
    { id: '', label: 'Semua Dokumentasi' },
    { id: 'Penertiban', label: 'Operasi & Penertiban' },
    { id: 'Apel & Upacara', label: 'Apel & Peringatan Hari Besar' },
    { id: 'Koordinasi', label: 'Koordinasi & Sosialisasi' }
  ];

  $: filteredItems = selectedCategory 
    ? $galleryStore.filter(item => item.category === selectedCategory)
    : $galleryStore;

  function openLightbox(index: number) {
    selectedIndex = index;
    showLightbox = true;
  }
</script>

<svelte:head>
  <title>Galeri Dokumentasi Kegiatan | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Dokumentasi Multimedia</div>
    <h1 class="text-3xl font-extrabold text-[#1a120c] tracking-tight mt-1">Galeri Kegiatan &amp; Penugasan</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Rekam jejak visual pengabdian personel Satpol PP Provinsi Sulawesi Selatan dalam mengawal penegakan Peraturan Daerah, menjaga ketenteraman masyarakat, serta apel kesiapsiagaan korps Praja Wibawa.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-8">
  <!-- Filter Tabs -->
  <div class="flex flex-wrap gap-2 bg-white p-4 sm:p-5 border border-[#eee9db] rounded-xl shadow-xs">
    {#each categories as cat}
      <button 
        type="button"
        class={`px-4 py-2 text-xs font-bold uppercase transition rounded-lg ${selectedCategory === cat.id ? 'bg-[#3d261a] text-white shadow-xs' : 'bg-stone-50 border border-stone-200 text-[#1a120c] hover:bg-stone-100'}`}
        on:click={() => selectedCategory = cat.id}
      >
        {cat.label}
      </button>
    {/each}
  </div>

  <!-- Gallery Grid -->
  {#if filteredItems.length > 0}
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
      {#each filteredItems as item, idx}
        <button 
          type="button"
          class="bg-white border border-[#eee9db] rounded-xl overflow-hidden group text-left focus:outline-none focus:ring-2 focus:ring-[#F08300] shadow-xs hover:shadow-md transition flex flex-col justify-between"
          on:click={() => openLightbox(idx)}
        >
          <div class="relative h-56 bg-stone-100 overflow-hidden w-full">
            <img 
              src={item.image_url} 
              alt={item.title} 
              class="w-full h-full object-cover group-hover:scale-105 transition duration-300"
              loading="lazy"
            />
            <span class="absolute top-3 left-3 bg-[#1a120c]/85 text-[#F08300] font-bold text-[10px] uppercase px-2.5 py-1 tracking-wider border-l-2 border-[#F08300] rounded-sm">
              {item.category}
            </span>

            <!-- Zoom overlay icon indicator -->
            <div class="absolute inset-0 bg-black/20 opacity-0 group-hover:opacity-100 transition flex items-center justify-center">
              <span class="bg-[#3d261a]/90 text-white p-2.5 rounded-full shadow-lg">
                <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0zM10 7v3m0 0v3m0-3h3m-3 0H7" />
                </svg>
              </span>
            </div>
          </div>
          <div class="p-4 space-y-1">
            <div class="text-[10px] text-[#6f7775] font-semibold">{item.date}</div>
            <h3 class="text-xs font-bold text-[#1a120c] group-hover:text-[#F08300] transition line-clamp-2 leading-snug">
              {item.title}
            </h3>
          </div>
        </button>
      {/each}
    </div>
  {:else}
    <div class="bg-white p-12 text-center rounded-xl border border-[#eee9db] text-stone-400">
      <p class="text-xs font-medium">Tidak ada foto dokumentasi pada kategori ini.</p>
    </div>
  {/if}
</div>

<!-- Reusable Lightbox Modal with Next/Prev and Keyboard Navigation (Poin 16) -->
<Lightbox
  bind:open={showLightbox}
  items={filteredItems}
  currentIndex={selectedIndex}
  on:change={(e) => selectedIndex = e.detail}
  on:close={() => showLightbox = false}
/>
