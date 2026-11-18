<script lang="ts">
  import { createEventDispatcher, onDestroy } from 'svelte';
  import type { GalleryItem } from '$lib/types';

  export let open: boolean = false;
  export let items: GalleryItem[] = [];
  export let currentIndex: number = 0;

  const dispatch = createEventDispatcher<{
    close: void;
    change: number;
  }>();

  let dialogEl: HTMLDialogElement;

  $: currentItem = items[currentIndex] || null;
  $: if (dialogEl) {
    if (open && currentItem && !dialogEl.open) {
      dialogEl.showModal();
    } else if ((!open || !currentItem) && dialogEl.open) {
      dialogEl.close();
    }
  }

  function handleClose() {
    if (dialogEl?.open) dialogEl.close();
    open = false;
    dispatch('close');
  }

  function handleNativeClose() {
    open = false;
    dispatch('close');
  }

  function handlePrev() {
    if (items.length <= 1) return;
    currentIndex = (currentIndex - 1 + items.length) % items.length;
    dispatch('change', currentIndex);
  }

  function handleNext() {
    if (items.length <= 1) return;
    currentIndex = (currentIndex + 1) % items.length;
    dispatch('change', currentIndex);
  }

  function handleKeyDown(e: KeyboardEvent) {
    if (!open) return;
    if (e.key === 'ArrowLeft') handlePrev();
    else if (e.key === 'ArrowRight') handleNext();
  }

  onDestroy(() => {
    if (dialogEl?.open) dialogEl.close();
  });
</script>

<svelte:window on:keydown={handleKeyDown} />

<dialog
  bind:this={dialogEl}
  aria-label={currentItem ? `Penampil Galeri: ${currentItem.title}` : 'Penampil Galeri Foto'}
  on:close={handleNativeClose}
  class="m-0 max-w-none max-h-none w-screen h-screen bg-[#070e1c]/88 backdrop-blur-2xl text-white p-3 sm:p-6 select-none open:flex flex-col justify-between border-0 ds-specular-dark"
>
  {#if currentItem}
    <!-- Top Bar (Frosted Obsidian Glass Pill) -->
    <div class="flex items-center justify-between text-white border-b border-white/10 pb-3 shrink-0">
      <div class="flex items-center gap-3 truncate pr-4">
        <span class="bg-[#F08300] text-white text-[10px] font-bold uppercase tracking-wider px-3 py-1 rounded-full shadow-xs">
          {currentItem.category || 'Dokumentasi'}
        </span>
        <span class="text-xs text-slate-200 font-medium truncate font-light">
          {currentItem.title}
        </span>
      </div>

      <div class="flex items-center gap-3 shrink-0">
        {#if items.length > 1}
          <span class="text-xs font-mono text-slate-400 ds-glass-dark px-2.5 py-1 rounded-full border border-white/10">
            {currentIndex + 1} / {items.length}
          </span>
        {/if}
        <button
          type="button"
          on:click={handleClose}
          class="p-2 text-slate-300 hover:text-white ds-glass-dark hover:bg-white/15 rounded-full border border-white/15 transition ds-press"
          aria-label="Tutup penampil galeri"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </div>

    <!-- Center Stage -->
    <div class="relative flex-1 flex items-center justify-center my-3 overflow-hidden">
      {#if items.length > 1}
        <button
          type="button"
          on:click={handlePrev}
          class="absolute left-2 sm:left-4 z-10 p-3.5 rounded-full ds-glass-dark hover:bg-white/20 text-white/90 hover:text-white backdrop-blur-xl border border-white/25 transition transform hover:scale-105 ds-press shadow-2xl"
          aria-label="Foto Sebelumnya"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
          </svg>
        </button>
      {/if}

      <img
        src={currentItem.image_url}
        alt={currentItem.title}
        class="max-w-full max-h-[74vh] object-contain rounded-2xl shadow-2xl transition duration-200 border border-white/15"
      />

      {#if items.length > 1}
        <button
          type="button"
          on:click={handleNext}
          class="absolute right-2 sm:right-4 z-10 p-3.5 rounded-full ds-glass-dark hover:bg-white/20 text-white/90 hover:text-white backdrop-blur-xl border border-white/25 transition transform hover:scale-105 ds-press shadow-2xl"
          aria-label="Foto Selanjutnya"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
          </svg>
        </button>
      {/if}
    </div>

    <!-- Bottom Caption -->
    <div class="bg-black/40 border-t border-white/10 rounded-xl p-3 sm:p-4 text-center text-xs text-stone-300 max-w-3xl mx-auto shrink-0 space-y-1">
      <h4 class="font-bold text-white text-sm">{currentItem.title}</h4>
      <div class="flex items-center justify-center gap-2 text-[11px] text-stone-400">
        <span>Tanggal: {currentItem.date}</span>
        <span>&bull;</span>
        <span>Satpol PP Provinsi Sulawesi Selatan</span>
      </div>
    </div>
  {/if}
</dialog>
