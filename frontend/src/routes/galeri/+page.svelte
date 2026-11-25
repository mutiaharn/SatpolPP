<script lang="ts">
  import { galleryStore } from '$lib/stores/appState';
  import type { GalleryItem } from '$lib/types';
  import Lightbox from '$lib/components/ui/Lightbox.svelte';
  import { Badge } from '$lib/components/ui';
  import { Camera, ChevronRight, Image as ImageIcon, ZoomIn } from 'lucide-svelte';

  let selectedCategory = '';
  let showLightbox = false;
  let selectedIndex = 0;

  const categories = [
    { id: '', label: 'Semua Dokumentasi' },
    { id: 'Penertiban', label: 'Operasi & Penertiban' },
    { id: 'Apel & Upacara', label: 'Apel & Upacara Korps' },
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

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#F08300]/5 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">Galeri Multimedia</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">Galeri Foto &amp; Kegiatan</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      Rekam jejak visual dedikasi dan pengabdian personel Satpol PP Sulsel dalam mengayomi warga, apel kesiapsiagaan, dan patroli pengamanan wilayah.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-8">
  <!-- Filter Tabs (Apple HIG Glass Capsules) -->
  <div class="flex flex-wrap gap-2 ds-glass p-4 sm:p-5 border border-white/90 rounded-3xl shadow-glass ds-specular-top">
    {#each categories as cat}
      <button 
        type="button"
        class={`px-4 py-2.5 min-h-[40px] inline-flex items-center justify-center text-xs font-bold uppercase tracking-wider transition duration-200 rounded-xl ds-press ${
          selectedCategory === cat.id 
            ? 'ds-pill-active' 
            : 'ds-pill-inactive'
        }`}
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
          class="ds-glass-card border border-white/80 rounded-2xl overflow-hidden group text-left focus:outline-none focus:ring-2 focus:ring-[#F08300] shadow-glass flex flex-col justify-between ds-press"
          on:click={() => openLightbox(idx)}
        >
          <div class="relative h-56 bg-slate-100 overflow-hidden w-full">
            <img 
              src={item.image_url} 
              alt={item.title} 
              class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
              loading="lazy"
            />
            <div class="absolute top-3 left-3">
              <Badge variant="gold" size="sm">
                {item.category}
              </Badge>
            </div>

            <!-- Zoom overlay icon indicator -->
            <div class="absolute inset-0 bg-[#070e1c]/35 opacity-0 group-hover:opacity-100 transition-opacity duration-300 flex items-center justify-center backdrop-blur-xs">
              <span class="ds-glass-dark text-white p-3.5 rounded-full shadow-2xl border border-white/25">
                <ZoomIn class="w-5 h-5 text-[#F08300]" />
              </span>
            </div>
          </div>
          <div class="p-5 space-y-1.5 flex-1 flex flex-col justify-between">
            <div class="text-[10px] text-slate-400 font-semibold uppercase tracking-wider font-mono">{item.date}</div>
            <h3 class="text-xs sm:text-sm font-bold text-slate-900 group-hover:text-[#F08300] transition duration-200 line-clamp-2 leading-snug">
              {item.title}
            </h3>
          </div>
        </button>
      {/each}
    </div>
  {:else}
    <div class="ds-glass border border-white/90 p-12 text-center rounded-3xl text-slate-400 space-y-2 shadow-glass">
      <ImageIcon class="w-8 h-8 text-slate-300 mx-auto mb-1" />
      <p class="text-xs font-medium text-slate-600">Tidak ada foto dokumentasi pada kategori ini.</p>
    </div>
  {/if}
</div>

<!-- Modal Lightbox -->
{#if showLightbox}
  <Lightbox
    items={filteredItems}
    {selectedIndex}
    on:close={() => showLightbox = false}
  />
{/if}
