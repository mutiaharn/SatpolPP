<script lang="ts">
  import { articlesStore } from '$lib/stores/appState';
  import type { Article } from '$lib/types';
  import { ChevronRight, Newspaper } from 'lucide-svelte';

  let selectedCategory = '';
  let searchQuery = '';
  let currentPage = 1;
  let pageSize = 6;

  const categories = [
    { id: '', label: 'Semua Berita' },
    { id: 'Kegiatan', label: 'Kegiatan & Patroli' },
    { id: 'Penertiban Aset', label: 'Penertiban Aset' },
    { id: 'Pengumuman', label: 'Pengumuman Resmi' }
  ];

  // Reactive on-chain search filter (Poin 11)
  $: filteredArticles = $articlesStore.filter(a => {
    const matchCat = !selectedCategory || a.category.toLowerCase() === selectedCategory.toLowerCase();
    const matchSearch = !searchQuery || 
      a.title.toLowerCase().includes(searchQuery.toLowerCase()) || 
      a.summary.toLowerCase().includes(searchQuery.toLowerCase());
    return matchCat && matchSearch;
  });

  // Reset to page 1 on filter or search change
  $: if (searchQuery !== undefined || selectedCategory) {
    currentPage = 1;
  }

  // Pagination calculation (Poin 10)
  $: totalArticles = filteredArticles.length;
  $: totalPages = Math.max(1, Math.ceil(totalArticles / pageSize));
  $: if (currentPage > totalPages) currentPage = totalPages;

  $: startIndex = (currentPage - 1) * pageSize;
  $: endIndex = Math.min(startIndex + pageSize, totalArticles);
  $: paginatedArticles = filteredArticles.slice(startIndex, endIndex);

  function formatDate(iso: string) {
    try {
      const d = new Date(iso);
      return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' });
    } catch {
      return iso;
    }
  }

  function goToPage(p: number) {
    if (p >= 1 && p <= totalPages) {
      currentPage = p;
      if (typeof window !== 'undefined') {
        window.scrollTo({ top: 180, behavior: 'smooth' });
      }
    }
  }

  function clearSearch() {
    searchQuery = '';
  }
</script>

<svelte:head>
  <title>Warta Berita &amp; Kegiatan Resmi | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#246bfd]/8 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <!-- Breadcrumb Pills -->
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">Warta &amp; Berita</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">Kabar &amp; Informasi Terkini</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      Publikasi resmi seputar agenda kegiatan, hasil patroli lapangan, serta pengumuman penting Satpol PP Provinsi Sulawesi Selatan untuk masyarakat.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-8">
  <!-- Search & Filter Controls (Tamalate Specular Glass Style) -->
  <div class="flex flex-col md:flex-row justify-between items-stretch md:items-center gap-4 ds-glass p-5 border border-white/90 rounded-3xl shadow-glass ds-specular-top">
    <div class="flex flex-wrap gap-2">
      {#each categories as cat}
        <button
          type="button"
          class={`px-4 py-2 text-xs font-bold uppercase transition duration-200 rounded-xl ds-press ${selectedCategory === cat.id ? 'ds-pill-active' : 'ds-pill-inactive'}`}
          on:click={() => selectedCategory = cat.id}
        >
          {cat.label}
        </button>
      {/each}
    </div>

    <!-- Reactive On-Chain Search Input -->
    <search class="relative w-full md:w-80">
      <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-slate-400">
        <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </div>
      <input
        type="search"
        bind:value={searchQuery}
        placeholder="Cari judul berita atau topik..."
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
    </search>
  </div>

  <!-- Articles Grid -->
  {#if paginatedArticles.length > 0}
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      {#each paginatedArticles as article}
        <article class="ds-glass-card border border-white/80 rounded-2xl overflow-hidden flex flex-col justify-between h-full group">
          <div>
            <div class="relative h-48 overflow-hidden bg-slate-100">
              <img 
                src={article.thumbnail} 
                alt={article.title} 
                class="w-full h-full object-cover group-hover:scale-105 transition duration-500"
                loading="lazy"
                onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
              />
              <span class="absolute top-3 left-3 ds-glass-dark text-[#F08300] font-bold text-[10px] uppercase px-3 py-1 tracking-wider border border-white/20 rounded-lg shadow-xs">
                {article.category}
              </span>
            </div>

            <div class="p-5">
              <div class="flex items-center gap-2 text-[11px] text-slate-400 mb-2 font-medium">
                <span>{formatDate(article.published_at)}</span>
                <span>&bull;</span>
                <span class="truncate">{article.author_name}</span>
              </div>

              <h2 class="text-sm font-bold text-[#071936] line-clamp-2 min-h-[2.5rem] leading-snug group-hover:text-[#246bfd] transition duration-200">
                <a href={`/berita/${article.slug}`}>{article.title}</a>
              </h2>

              <p class="text-xs text-slate-500 mt-2 line-clamp-3 min-h-[3.5rem] leading-relaxed font-light">
                {article.summary}
              </p>
            </div>
          </div>

          <div class="p-5 pt-0 border-t border-slate-100/90 flex items-center justify-between text-xs mt-3">
            <span class="text-[11px] text-slate-400 font-mono">{article.view_count} pembaca</span>
            <a 
              href={`/berita/${article.slug}`} 
              class="text-xs font-bold uppercase text-[#246bfd] group-hover:text-[#1d4ed8] flex items-center gap-1 transition"
            >
              Baca Lengkap &rarr;
            </a>
          </div>
        </article>
      {/each}
    </div>

    <!-- Pagination Controls (Poin 10) -->
    <div class="p-4 ds-glass-strong border border-white/90 rounded-2xl flex flex-col sm:flex-row justify-between items-center gap-3 text-xs text-slate-600 shadow-glass">
      <div>
        Menampilkan <strong class="text-slate-900">{startIndex + 1}</strong> - <strong class="text-slate-900">{endIndex}</strong> dari <strong class="text-slate-900">{totalArticles}</strong> warta berita
      </div>

      {#if totalPages > 1}
        <div class="flex items-center gap-1.5">
          <button
            type="button"
            on:click={() => goToPage(currentPage - 1)}
            disabled={currentPage === 1}
            class="px-3 py-1.5 rounded-xl ds-glass border border-white/80 text-slate-700 hover:bg-white/90 disabled:opacity-40 disabled:cursor-not-allowed transition font-semibold text-xs ds-press"
          >
            &larr; Sebelumnya
          </button>

          <div class="flex items-center gap-1">
            {#each Array.from({ length: totalPages }, (_, i) => i + 1) as p}
              <button
                type="button"
                on:click={() => goToPage(p)}
                class={`w-8 h-8 rounded-xl font-bold text-xs transition ds-press ${
                  currentPage === p
                    ? 'ds-pill-active'
                    : 'ds-pill-inactive'
                }`}
              >
                {p}
              </button>
            {/each}
          </div>

          <button
            type="button"
            on:click={() => goToPage(currentPage + 1)}
            disabled={currentPage === totalPages}
            class="px-3 py-1.5 rounded-xl ds-glass border border-white/80 text-slate-700 hover:bg-white/90 disabled:opacity-40 disabled:cursor-not-allowed transition font-semibold text-xs ds-press"
          >
            Selanjutnya &rarr;
          </button>
        </div>
      {/if}
    </div>
  {:else}
    <div class="bg-white border border-[#eee9db] rounded-xl p-12 text-center text-stone-400 space-y-2">
      <p class="text-xs font-semibold text-stone-600">Tidak ada artikel berita yang sesuai dengan kriteria pencarian Anda.</p>
      {#if searchQuery}
        <button
          type="button"
          on:click={clearSearch}
          class="text-xs font-bold text-[#F08300] hover:underline"
        >
          Reset kata kunci pencarian
        </button>
      {/if}
    </div>
  {/if}
</div>
