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

<div class={`bg-white border border-[#eee9db] rounded-xl shadow-xs overflow-hidden flex flex-col ${className}`}>
  <!-- Header & Toolbar -->
  <div class="p-4 sm:p-5 border-b border-[#eee9db] bg-gradient-to-r from-[#faf7f2] to-white flex flex-col sm:flex-row justify-between items-stretch sm:items-center gap-3">
    <!-- Title / Description -->
    <div>
      {#if title}
        <h3 class="text-sm sm:text-base font-bold text-[#1a120c] flex items-center gap-2">
          {title}
          <span class="text-[11px] font-semibold text-[#8c6b45] bg-[#f5ede4] px-2 py-0.5 rounded-full border border-[#e8dbc9]">
            {totalItems} total
          </span>
        </h3>
      {/if}
      {#if description}
        <p class="text-[11px] text-[#6f7775] mt-0.5">{description}</p>
      {/if}
      <slot name="header-left" />
    </div>

    <!-- Actions & On-Chain Search -->
    <div class="flex flex-wrap items-center gap-2.5">
      <slot name="header-actions" />

      {#if showSearch}
        <div class="relative w-full sm:w-64 md:w-72">
          <!-- Search Icon -->
          <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-stone-400">
            <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
          <input
            type="text"
            value={searchQuery}
            on:input={handleSearchInput}
            placeholder={searchPlaceholder}
            class="w-full pl-9 pr-8 py-2 text-xs bg-white text-[#1a120c] border border-stone-300 rounded-lg placeholder-stone-400 focus:outline-none focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] transition"
          />
          {#if searchQuery}
            <button
              type="button"
              on:click={clearSearch}
              class="absolute inset-y-0 right-0 pr-2.5 flex items-center text-stone-400 hover:text-stone-600"
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
      <thead class="bg-[#f9f4e8] border-b border-[#eee9db] text-[#3d261a] uppercase text-[10px] tracking-wider font-bold">
        <slot name="head" />
      </thead>
      <tbody class="divide-y divide-gray-100">
        {#if paginatedItems.length > 0}
          {#each paginatedItems as item, idx}
            <slot name="row" {item} index={startIndex + idx} />
          {/each}
        {:else}
          <tr>
            <td colspan="100" class="py-12 px-4 text-center">
              <slot name="empty">
                <div class="flex flex-col items-center justify-center text-stone-400 space-y-2">
                  <div class="w-12 h-12 rounded-full bg-stone-100 flex items-center justify-center text-stone-400">
                    <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.172 16.172a4 4 0 015.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <p class="text-xs font-medium text-stone-600">{emptyMessage}</p>
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
    <div class="p-3.5 sm:p-4 border-t border-[#eee9db] bg-[#faf7f2]/60 flex flex-col sm:flex-row justify-between items-center gap-3 text-xs text-[#6f7775]">
      <!-- Range Info & Page Size Selector -->
      <div class="flex flex-wrap items-center gap-3">
        <span>
          Menampilkan <strong class="text-[#1a120c]">{totalItems > 0 ? startIndex + 1 : 0}</strong> - <strong class="text-[#1a120c]">{endIndex}</strong> dari <strong class="text-[#1a120c]">{totalItems}</strong> data
        </span>
        <div class="flex items-center gap-1.5 pl-2 border-l border-stone-200">
          <label for="page-size-select" class="text-[11px]">Baris:</label>
          <select
            id="page-size-select"
            bind:value={pageSize}
            class="bg-white border border-stone-300 rounded-md text-[11px] px-2 py-1 text-[#1a120c] focus:outline-none focus:border-[#F08300]"
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
            class="p-1.5 rounded-lg border border-stone-200 bg-white text-stone-600 hover:bg-stone-50 disabled:opacity-40 disabled:cursor-not-allowed transition"
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
            class="p-1.5 rounded-lg border border-stone-200 bg-white text-stone-600 hover:bg-stone-50 disabled:opacity-40 disabled:cursor-not-allowed transition"
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
                <span class="px-2 text-stone-400 select-none">&hellip;</span>
              {:else}
                <button
                  type="button"
                  on:click={() => goToPage(Number(p))}
                  class={`min-w-[28px] h-7 px-2 text-xs font-semibold rounded-lg transition ${
                    currentPage === p
                      ? 'bg-[#3d261a] text-white shadow-xs'
                      : 'bg-white border border-stone-200 text-stone-700 hover:bg-stone-100'
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
            class="p-1.5 rounded-lg border border-stone-200 bg-white text-stone-600 hover:bg-stone-50 disabled:opacity-40 disabled:cursor-not-allowed transition"
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
            class="p-1.5 rounded-lg border border-stone-200 bg-white text-stone-600 hover:bg-stone-50 disabled:opacity-40 disabled:cursor-not-allowed transition"
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
