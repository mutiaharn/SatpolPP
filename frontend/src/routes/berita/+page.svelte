<script lang="ts">
  import { articlesStore } from '$lib/stores/appState';
  import type { Article } from '$lib/types';

  let selectedCategory = '';
  let searchQuery = '';

  const categories = [
    { id: '', label: 'Semua Berita' },
    { id: 'Kegiatan', label: 'Kegiatan & Patroli' },
    { id: 'Penertiban Aset', label: 'Penertiban Aset' },
    { id: 'Pengumuman', label: 'Pengumuman Resmi' }
  ];

  $: filteredArticles = $articlesStore.filter(a => {
    const matchCat = !selectedCategory || a.category.toLowerCase() === selectedCategory.toLowerCase();
    const matchSearch = !searchQuery || 
      a.title.toLowerCase().includes(searchQuery.toLowerCase()) || 
      a.summary.toLowerCase().includes(searchQuery.toLowerCase());
    return matchCat && matchSearch;
  });

  function formatDate(iso: string) {
    try {
      const d = new Date(iso);
      return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' });
    } catch {
      return iso;
    }
  }
</script>

<svelte:head>
  <title>Warta Berita &amp; Kegiatan Resmi | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Warta Praja Wibawa</div>
    <h1 class="text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Kabar &amp; Informasi Terkini</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Publikasi resmi seputar giat operasional, penegakan Peraturan Daerah, pembinaan ketertiban umum, dan pelindungan masyarakat di 24 Kabupaten/Kota se-Sulawesi Selatan.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-8">
  <!-- Search & Filter Controls -->
  <div class="flex flex-col md:flex-row justify-between items-stretch md:items-center gap-4 bg-white p-5 border border-[#eee9db] rounded-2xl shadow-xs">
    <div class="flex flex-wrap gap-2">
      {#each categories as cat}
        <button
          type="button"
          class={`px-4 py-2 text-xs font-bold uppercase transition rounded-xl ${selectedCategory === cat.id ? 'bg-[#F08300] text-white shadow-xs' : 'bg-gray-50 border border-gray-200 text-[#101116] hover:bg-gray-100'}`}
          on:click={() => selectedCategory = cat.id}
        >
          {cat.label}
        </button>
      {/each}
    </div>

    <div class="w-full md:w-72">
      <input
        type="text"
        bind:value={searchQuery}
        placeholder="Cari judul berita atau topik..."
        class="w-full border border-gray-300 rounded-xl px-3.5 py-2 text-xs text-[#101116] focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] focus:outline-none"
      />
    </div>
  </div>

  <!-- Articles Grid -->
  {#if filteredArticles.length > 0}
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
      {#each filteredArticles as article}
        <article class="bg-white border border-[#eee9db] rounded-2xl overflow-hidden flex flex-col justify-between hover:shadow-lg transition">
          <div>
            <div class="relative h-48 overflow-hidden bg-gray-100">
              <img 
                src={article.thumbnail} 
                alt={article.title} 
                class="w-full h-full object-cover"
                loading="lazy"
                onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
              />
              <span class="absolute top-3 left-3 bg-[#101116]/85 text-[#F08300] font-bold text-[10px] uppercase px-2.5 py-1 tracking-wider border-l-2 border-[#F08300] rounded-sm">
                {article.category}
              </span>
            </div>

            <div class="p-5">
              <div class="flex items-center gap-2 text-[11px] text-[#6f7775] mb-2 font-semibold">
                <span>{formatDate(article.published_at)}</span>
                <span>&bull;</span>
                <span>{article.author_name}</span>
              </div>

              <h2 class="text-sm font-bold text-[#101116] line-clamp-2 leading-snug hover:text-[#F08300] transition">
                <a href={`/berita/${article.slug}`}>{article.title}</a>
              </h2>

              <p class="text-xs text-[#6f7775] mt-2 line-clamp-3 leading-relaxed">
                {article.summary}
              </p>
            </div>
          </div>

          <div class="p-5 pt-0 border-t border-gray-50 flex items-center justify-between text-xs mt-3">
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
  {:else}
    <div class="bg-white border border-[#eee9db] rounded-2xl p-12 text-center text-gray-400 text-xs">
      Tidak ada artikel berita yang sesuai dengan kriteria pencarian Anda.
    </div>
  {/if}
</div>
