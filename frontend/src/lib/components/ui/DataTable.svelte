<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  // Generic data item type
  type T = $$Generic<Record<string, any>>;

  export let items: T[] = [];
  export let searchFields: string[] = [];
  export let searchPlaceholder: string = 'Ketik kata kunci pencarian...';
  export let searchQuery: string = '';
  export let pageSize: number = 10;
  export let pageSizeOptions: number[] = [5, 10, 25, 50];
  export let title: string = '';
  export let description: string = '';
  export let emptyMessage: string = 'Tidak ada data yang sesuai dengan pencarian.';
  export let showSearch: boolean = true;
  export let showPagination: boolean = true;
  export let className: string = '';

  const dispatch = createEventDispatcher<{
    rowClick: T;
    searchChange: string;
  }>();

  let currentPage = 1;

  // Reactive real-time on-chain search filter
  $: filteredItems = items.filter(item => {
    if (!searchQuery || searchQuery.trim() === '') return true;
    const q = searchQuery.toLowerCase().trim();

    if (searchFields.length > 0) {
      return searchFields.some(field => {
        const val = item[field];
        if (val === null || val === undefined) return false;
        return String(val).toLowerCase().includes(q);
      });
    }

    // Default: search all string/number properties
    return Object.values(item).some(val => {
      if (val === null || val === undefined) return false;
      return String(val).toLowerCase().includes(q);
    });
  });

  // Reset page when search or items change
  $: if (searchQuery !== undefined || items) {
    currentPage = 1;
  }

  // Pagination calculations
  $: totalItems = filteredItems.length;
  $: totalPages = Math.max(1, Math.ceil(totalItems / pageSize));
  $: if (currentPage > totalPages) currentPage = totalPages;

  $: startIndex = (currentPage - 1) * pageSize;
  $: endIndex = Math.min(startIndex + pageSize, totalItems);
  $: paginatedItems = filteredItems.slice(startIndex, endIndex);

  function handleSearchInput(e: Event) {
    const target = e.target as HTMLInputElement;
    searchQuery = target.value;
    dispatch('searchChange', searchQuery);
  }

  function clearSearch() {
    searchQuery = '';
    dispatch('searchChange', '');
  }

  function goToPage(page: number) {
    if (page >= 1 && page <= totalPages) {
      currentPage = page;
    }
  }

  // Generate page numbers with smart ellipsis
  function getVisiblePages(current: number, total: number): (number | string)[] {
    if (total <= 7) {
      return Array.from({ length: total }, (_, i) => i + 1);
    }
    if (current <= 3) {
      return [1, 2, 3, 4, '...', total];
    }
    if (current >= total - 2) {
      return [1, '...', total - 3, total - 2, total - 1, total];
    }
    return [1, '...', current - 1, current, current + 1, '...', total];
  }

  $: visiblePages = getVisiblePages(currentPage, totalPages);
</script>

<div class={`ds-glass-strong border border-white/90 rounded-3xl shadow-glass overflow-hidden flex flex-col ds-specular-top ${className}`}>
  <!-- Header & Toolbar -->
  <div class="p-4 sm:p-5 border-b border-white/80 bg-white/40 backdrop-blur-md flex flex-col sm:flex-row justify-between items-stretch sm:items-center gap-3">
    <!-- Title / Description -->
    <div>
      {#if title}
        <h3 class="text-sm sm:text-base font-bold text-slate-900 flex items-center gap-2">
          {title}
          <span class="text-[11px] font-semibold text-[#F08300] bg-orange-50/90 px-2.5 py-0.5 rounded-full border border-orange-200/80 shadow-2xs">
            {totalItems} total
          </span>
        </h3>
      {/if}
      {#if description}
        <p class="text-[11px] text-slate-500 mt-0.5 font-light">{description}</p>
      {/if}
      <slot name="header-left" />
    </div>

    <!-- Actions & On-Chain Search -->
    <div class="flex flex-wrap items-center gap-2.5">
      <slot name="header-actions" />

      {#if showSearch}
        <div class="relative w-full sm:w-64 md:w-72">
          <!-- Search Icon -->
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            type="text"
            value={searchQuery}
            on:input={handleSearchInput}
            placeholder={searchPlaceholder}
            class="w-full pl-9 pr-8 py-2 text-xs ds-glass-input text-slate-900 rounded-xl placeholder-slate-400 focus:outline-none transition font-light"
          />
          {#if searchQuery}
            <button
              type="button"
              on:click={clearSearch}
              class="absolute inset-y-0 right-0 pr-2.5 flex items-center text-slate-400 hover:text-slate-600 ds-press"
              aria-label="Bersihkan pencarian"
            >
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          {/if}
        </div>
      {/if}
    </div>
  </div>

  <!-- Table Container with Horizontal Scroll -->
  <div class="overflow-x-auto min-h-[160px]">
    <table class="w-full text-left text-xs border-collapse">
      <thead class="bg-slate-100/70 backdrop-blur-md border-b border-white/80 text-slate-700 uppercase text-[10px] tracking-wider font-bold">
        <slot name="head" />
      </thead>
      <tbody class="divide-y divide-white/60">
        {#if paginatedItems.length > 0}
          {#each paginatedItems as item, idx}
            <slot name="row" {item} index={startIndex + idx} />
          {/each}
        {:else}
          <tr>
            <td colspan="100" class="py-12 px-4 text-center">
              <slot name="empty">
                <div class="flex flex-col items-center justify-center text-slate-400 space-y-2">
                  <div class="w-12 h-12 rounded-full bg-slate-100 flex items-center justify-center text-slate-400">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <p class="text-xs font-medium text-slate-600">{emptyMessage}</p>
                  {#if searchQuery}
                    <button
                      type="button"
                      on:click={clearSearch}
                      class="text-[11px] font-bold text-[#F08300] hover:underline"
                    >
                      Reset filter pencarian
                    </button>
                  {/if}
                </div>
              </slot>
            </td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>

  <!-- Pagination Footer -->
  {#if showPagination && totalItems > 0}
    <div class="p-3.5 sm:p-4 border-t border-white/80 bg-white/40 backdrop-blur-md flex flex-col sm:flex-row justify-between items-center gap-3 text-xs text-slate-600">
      <!-- Range Info & Page Size Selector -->
      <div class="flex flex-wrap items-center gap-3">
        <span>
          Menampilkan <strong class="text-slate-900">{totalItems > 0 ? startIndex + 1 : 0}</strong> - <strong class="text-slate-900">{endIndex}</strong> dari <strong class="text-slate-900">{totalItems}</strong> data
        </span>
        <div class="flex items-center gap-1.5 pl-2 border-l border-slate-200/80">
          <label for="page-size-select" class="text-[11px]">Baris:</label>
          <select
            id="page-size-select"
            bind:value={pageSize}
            class="ds-glass-input rounded-lg text-[11px] px-2 py-1 text-slate-900 focus:outline-none"
          >
            {#each pageSizeOptions as opt}
              <option value={opt}>{opt}</option>
            {/each}
          </select>
        </div>
      </div>

      <!-- Page Controls -->
      {#if totalPages > 1}
        <div class="flex items-center gap-1">
          <!-- First -->
          <button
            type="button"
            on:click={() => goToPage(1)}
            disabled={currentPage === 1}
            class="p-1.5 rounded-lg border border-white/80 ds-glass text-slate-600 hover:bg-white/80 disabled:opacity-40 disabled:cursor-not-allowed transition ds-press"
            title="Halaman Pertama"
            aria-label="Halaman Pertama"
          >
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 19l-7-7 7-7m8 14l-7-7 7-7" />
            </svg>
          </button>

          <!-- Prev -->
          <button
            type="button"
            on:click={() => goToPage(currentPage - 1)}
            disabled={currentPage === 1}
            class="p-1.5 rounded-lg border border-white/80 ds-glass text-slate-600 hover:bg-white/80 disabled:opacity-40 disabled:cursor-not-allowed transition ds-press"
            title="Halaman Sebelumnya"
            aria-label="Halaman Sebelumnya"
          >
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
            </svg>
          </button>

          <!-- Page numbers -->
          <div class="flex items-center gap-1 px-1">
            {#each visiblePages as p}
              {#if p === '...'}
                <span class="px-2 text-slate-400 select-none">&hellip;</span>
              {:else}
                <button
                  type="button"
                  on:click={() => goToPage(Number(p))}
                  class={`min-w-[28px] h-7 px-2 text-xs font-bold rounded-lg transition ds-press ${
                    currentPage === p
                      ? 'ds-pill-active'
                      : 'ds-glass border border-white/80 text-slate-700 hover:bg-white/90 hover:text-slate-900'
                  }`}
                >
                  {p}
                </button>
              {/if}
            {/each}
          </div>

          <!-- Next -->
          <button
            type="button"
            on:click={() => goToPage(currentPage + 1)}
            disabled={currentPage === totalPages}
            class="p-1.5 rounded-lg border border-white/80 ds-glass text-slate-600 hover:bg-white/80 disabled:opacity-40 disabled:cursor-not-allowed transition ds-press"
            title="Halaman Selanjutnya"
            aria-label="Halaman Selanjutnya"
          >
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </button>

          <!-- Last -->
          <button
            type="button"
            on:click={() => goToPage(totalPages)}
            disabled={currentPage === totalPages}
            class="p-1.5 rounded-lg border border-white/80 ds-glass text-slate-600 hover:bg-white/80 disabled:opacity-40 disabled:cursor-not-allowed transition ds-press"
            title="Halaman Terakhir"
            aria-label="Halaman Terakhir"
          >
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5l7 7-7 7M5 5l7 7-7 7" />
            </svg>
          </button>
        </div>
      {/if}
    </div>
  {/if}
</div>
