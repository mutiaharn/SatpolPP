<script lang="ts">
  import { articlesStore } from '$lib/stores/appState';
  import type { Article } from '$lib/types';

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

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Warta Praja Wibawa</div>
    <h1 class="text-3xl font-extrabold text-[#1a120c] tracking-tight mt-1">Kabar &amp; Informasi Terkini</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3 rounded-full"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Publikasi resmi seputar giat operasional, penegakan Peraturan Daerah, pembinaan ketertiban umum, dan pelindungan masyarakat di 24 Kabupaten/Kota se-Sulawesi Selatan.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-8">
  <!-- Search & Filter Controls -->
  <div class="flex flex-col md:flex-row justify-between items-stretch md:items-center gap-4 bg-white p-5 border border-[#eee9db] rounded-xl shadow-xs">
    <div class="flex flex-wrap gap-2">
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

    <!-- Reactive On-Chain Search Input (Poin 11) -->
    <search class="relative w-full md:w-80">
      <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-stone-400">
        <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </div>
      <input
        type="search"
        bind:value={searchQuery}
        placeholder="Cari judul berita atau topik (on-chain)..."
        class="w-full pl-9 pr-8 py-2 text-xs bg-white text-[#1a120c] border border-stone-300 rounded-xl placeholder-stone-400 focus:outline-none focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] transition"
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
    </search>
  </div>

  <!-- Articles Grid -->
  {#if paginatedArticles.length > 0}
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      {#each paginatedArticles as article}
        <article class="bg-white border border-[#eee9db] rounded-2xl overflow-hidden flex flex-col justify-between h-full hover:shadow-lg transition">
          <div>
            <div class="relative h-48 overflow-hidden bg-stone-100">
              <img 
                src={article.thumbnail} 
                alt={article.title} 
                class="w-full h-full object-cover hover:scale-105 transition duration-300"
                loading="lazy"
                onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
              />
              <span class="absolute top-3 left-3 bg-[#1a120c]/85 text-[#F08300] font-bold text-[10px] uppercase px-2.5 py-1 tracking-wider border-l-2 border-[#F08300] rounded-lg">
                {article.category}
              </span>
            </div>

            <div class="p-5">
              <div class="flex items-center gap-2 text-[11px] text-[#6f7775] mb-2 font-semibold">
                <span>{formatDate(article.published_at)}</span>
                <span>&bull;</span>
                <span class="truncate">{article.author_name}</span>
              </div>

              <h2 class="text-sm font-bold text-[#1a120c] line-clamp-2 min-h-[2.5rem] leading-snug hover:text-[#F08300] transition">
                <a href={`/berita/${article.slug}`}>{article.title}</a>
              </h2>

              <p class="text-xs text-[#6f7775] mt-2 line-clamp-3 min-h-[3.5rem] leading-relaxed">
                {article.summary}
              </p>
            </div>
          </div>

          <div class="p-5 pt-0 border-t border-stone-100 flex items-center justify-between text-xs mt-3">
            <span class="text-[11px] text-[#6f7775]">{article.view_count} pembaca</span>
            <a 
              href={`/berita/${article.slug}`} 
              class="text-xs font-bold uppercase text-[#F08300] hover:text-[#d97400] flex items-center gap-1"
            >
              Baca Lengkap &rarr;
            </a>
          </div>
        </article>
      {/each}
    </div>

    <!-- Pagination Controls (Poin 10) -->
    <div class="p-4 bg-white border border-[#eee9db] rounded-xl flex flex-col sm:flex-row justify-between items-center gap-3 text-xs text-[#6f7775]">
      <div>
        Menampilkan <strong class="text-[#1a120c]">{startIndex + 1}</strong> - <strong class="text-[#1a120c]">{endIndex}</strong> dari <strong class="text-[#1a120c]">{totalArticles}</strong> warta berita
      </div>

      {#if totalPages > 1}
        <div class="flex items-center gap-1.5">
          <button
            type="button"
            on:click={() => goToPage(currentPage - 1)}
            disabled={currentPage === 1}
            class="px-3 py-1.5 rounded-lg border border-stone-200 bg-white text-stone-700 hover:bg-stone-50 disabled:opacity-40 disabled:cursor-not-allowed transition font-semibold"
          >
            &larr; Sebelumnya
          </button>

          <div class="flex items-center gap-1">
            {#each Array.from({ length: totalPages }, (_, i) => i + 1) as p}
              <button
                type="button"
                on:click={() => goToPage(p)}
                class={`w-8 h-8 rounded-lg font-bold text-xs transition ${
                  currentPage === p
                    ? 'bg-[#3d261a] text-white shadow-xs'
                    : 'bg-stone-50 border border-stone-200 text-stone-700 hover:bg-stone-100'
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
            class="px-3 py-1.5 rounded-lg border border-stone-200 bg-white text-stone-700 hover:bg-stone-50 disabled:opacity-40 disabled:cursor-not-allowed transition font-semibold"
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
