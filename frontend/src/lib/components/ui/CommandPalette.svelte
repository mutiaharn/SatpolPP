<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { goto } from '$app/navigation';
  import { articlesStore, documentsStore } from '$lib/stores/appState';
  import { Search, X, FileText, Newspaper, Shield, Building2, Phone, ArrowRight, CornerDownLeft } from 'lucide-svelte';

  export let open = false;

  let query = '';
  let inputEl: HTMLInputElement;
  let selectedIndex = 0;

  interface SearchResult {
    id: string;
    type: 'berita' | 'dokumen' | 'layanan' | 'profil';
    category: string;
    title: string;
    description: string;
    href: string;
  }

  const staticLinks: SearchResult[] = [
    {
      id: 'lapor',
      type: 'layanan',
      category: 'SIAP TERTIB',
      title: 'Buat Laporan Baru (SIAP TERTIB)',
      description: 'Saluran pengaduan pelanggaran Perda & Trantibum dengan perlindungan identitas',
      href: '/lapor'
    },
    {
      id: 'lacak',
      type: 'layanan',
      category: 'SIAP TERTIB',
      title: 'Lacak Tiket Aduan',
      description: 'Periksa status tindak lanjut tiket laporan masyarakat',
      href: '/lacak'
    },
    {
      id: 'struktur',
      type: 'profil',
      category: 'Profil',
      title: 'Struktur Organisasi Satpol PP',
      description: 'Bagan kepemimpinan, bidang, subbagian, dan satuan tugas',
      href: '/profil/struktur'
    },
    {
      id: 'visimisi',
      type: 'profil',
      category: 'Profil',
      title: 'Visi & Misi Satpol PP Sulsel',
      description: 'Arah kebijakan dan target strategis Praja Wibawa',
      href: '/profil/visi-misi'
    },
    {
      id: 'sejarah',
      type: 'profil',
      category: 'Profil',
      title: 'Sejarah Singkat Satpol PP',
      description: 'Tonggak sejarah pengabdian pamong praja di Sulawesi Selatan',
      href: '/profil/sejarah'
    },
    {
      id: 'kontak',
      type: 'layanan',
      category: 'Kontak',
      title: 'Kontak & Call Center',
      description: 'Telepon (0411) 453105, email, dan alamat kantor Jl. Urip Sumoharjo',
      href: '/kontak'
    }
  ];

  $: articleResults = $articlesStore.map(a => ({
    id: `art-${a.id}`,
    type: 'berita' as const,
    category: a.category || 'Berita',
    title: a.title,
    description: [a.summary, a.author_name].filter(Boolean).join(' · '),
    href: `/berita/${a.slug}`
  }));

  $: documentResults = $documentsStore.map(d => ({
    id: `doc-${d.id}`,
    type: 'dokumen' as const,
    category: d.category.toUpperCase(),
    title: d.title,
    description: d.description || `Dokumen publik tahun ${d.year}`,
    href: `/dokumen?category=${d.category}`
  }));

  $: allPool = [...staticLinks, ...articleResults, ...documentResults];

  $: filteredResults = query.trim() === ''
    ? staticLinks
    : allPool.filter(item => {
        const q = query.toLowerCase();
        return (
          item.title.toLowerCase().includes(q) ||
          item.description.toLowerCase().includes(q) ||
          item.category.toLowerCase().includes(q)
        );
      }).slice(0, 10);

  $: if (filteredResults.length && selectedIndex >= filteredResults.length) {
    selectedIndex = 0;
  }

  function handleKeydown(e: KeyboardEvent) {
    if ((e.ctrlKey || e.metaKey) && e.key.toLowerCase() === 'k') {
      e.preventDefault();
      open = !open;
    } else if (e.key === 'Escape' && open) {
      e.preventDefault();
      close();
    } else if (open) {
      if (e.key === 'ArrowDown') {
        e.preventDefault();
        selectedIndex = (selectedIndex + 1) % Math.max(1, filteredResults.length);
      } else if (e.key === 'ArrowUp') {
        e.preventDefault();
        selectedIndex = (selectedIndex - 1 + filteredResults.length) % Math.max(1, filteredResults.length);
      } else if (e.key === 'Enter') {
        e.preventDefault();
        if (filteredResults[selectedIndex]) {
          navigate(filteredResults[selectedIndex].href);
        }
      }
    }
  }

  function close() {
    open = false;
    query = '';
  }

  function navigate(href: string) {
    close();
    goto(href);
  }

  $: if (open) {
    setTimeout(() => inputEl?.focus(), 50);
  }

  onMount(() => {
    window.addEventListener('keydown', handleKeydown);
  });

  onDestroy(() => {
    if (typeof window !== 'undefined') {
      window.removeEventListener('keydown', handleKeydown);
    }
  });
</script>

{#if open}
  <!-- Backdrop -->
  <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
  <div
    class="fixed inset-0 z-50 bg-[#070e1c]/60 backdrop-blur-xl flex items-start justify-center pt-16 sm:pt-24 px-4 transition-opacity animate-in fade-in duration-200"
    role="presentation"
    on:click|self={close}
  >
    <!-- Modal Container (Sovereign Specular Frosted Glass) -->
    <div class="w-full max-w-2xl ds-glass-strong rounded-3xl shadow-2xl border border-white/95 overflow-hidden text-[#0f172a] transform transition-all animate-in zoom-in-95 duration-200 ds-specular-top">
      <!-- Search Input Bar -->
      <div class="relative flex items-center px-4 py-3.5 border-b border-white/80 bg-white/40 backdrop-blur-md">
        <Search class="w-5 h-5 text-slate-400 shrink-0 ml-1" />
        <input
          bind:this={inputEl}
          bind:value={query}
          type="text"
          placeholder="Cari berita, regulasi, dokumen publik, atau layanan..."
          class="w-full px-3 py-1.5 bg-transparent border-0 text-sm md:text-base text-slate-800 placeholder-slate-400 focus:outline-none focus:ring-0 font-light"
        />
        {#if query}
          <button
            type="button"
            on:click={() => query = ''}
            class="p-1 text-slate-400 hover:text-slate-600 rounded-lg hover:bg-slate-900/5 transition ds-press"
            aria-label="Hapus pencarian"
          >
            <X class="w-4 h-4" />
          </button>
        {/if}
        <div class="hidden sm:flex items-center gap-1 text-[11px] font-mono text-slate-400 bg-white/80 px-2 py-0.5 rounded-md border border-white/80 ml-2 shadow-2xs">
          <span>ESC</span>
        </div>
      </div>

      <!-- Result List -->
      <div class="max-h-[60vh] overflow-y-auto p-2 divide-y divide-white/60">
        {#if filteredResults.length > 0}
          <div class="px-3 py-1.5 text-[10px] font-bold uppercase tracking-wider text-slate-400 flex justify-between items-center">
            <span>Hasil Pencarian ({filteredResults.length})</span>
            <span class="text-[9px] text-slate-400 font-normal">Gunakan &uarr;&darr; untuk navigasi, Enter untuk memilih</span>
          </div>

          <div class="space-y-1 pt-1">
            {#each filteredResults as item, idx}
              <button
                type="button"
                on:click={() => navigate(item.href)}
                on:mouseenter={() => selectedIndex = idx}
                class="w-full text-left px-3.5 py-2.5 rounded-xl flex items-start gap-3 transition-colors duration-150 {selectedIndex === idx ? 'bg-white/80 text-[#071936] ring-1 ring-[#246bfd]/30 shadow-xs' : 'hover:bg-white/50 text-slate-700'}"
              >
                <!-- Icon by Type -->
                <div class="mt-0.5 p-2 rounded-lg shrink-0 {item.type === 'berita' ? 'bg-blue-50 text-[#246bfd]' : item.type === 'dokumen' ? 'bg-amber-50 text-[#f08300]' : item.type === 'layanan' ? 'bg-emerald-50 text-emerald-600' : 'bg-purple-50 text-purple-600'}">
                  {#if item.type === 'berita'}
                    <Newspaper class="w-4 h-4" />
                  {:else if item.type === 'dokumen'}
                    <FileText class="w-4 h-4" />
                  {:else if item.type === 'layanan'}
                    <Shield class="w-4 h-4" />
                  {:else}
                    <Building2 class="w-4 h-4" />
                  {/if}
                </div>

                <!-- Info -->
                <div class="flex-1 min-w-0">
                  <div class="flex items-center gap-2">
                    <span class="text-[10px] font-bold uppercase tracking-wider px-1.5 py-0.5 rounded-md {item.type === 'berita' ? 'bg-blue-100/60 text-blue-700' : item.type === 'dokumen' ? 'bg-amber-100/60 text-amber-800' : 'bg-slate-200/70 text-slate-700'}">
                      {item.category}
                    </span>
                    <h4 class="text-xs sm:text-sm font-semibold truncate {selectedIndex === idx ? 'text-[#071936] font-bold' : 'text-slate-800'}">
                      {item.title}
                    </h4>
                  </div>
                  <p class="text-[11px] text-slate-500 line-clamp-1 mt-0.5 font-light">
                    {item.description}
                  </p>
                </div>

                <div class="self-center pl-2 shrink-0 opacity-0 group-hover:opacity-100 transition {selectedIndex === idx ? 'opacity-100 text-[#246bfd]' : 'text-slate-300'}">
                  <CornerDownLeft class="w-3.5 h-3.5" />
                </div>
              </button>
            {/each}
          </div>
        {:else}
          <div class="text-center py-10 px-4">
            <Search class="w-8 h-8 text-slate-300 mx-auto mb-2" />
            <p class="text-sm font-semibold text-slate-700">Tidak ada hasil untuk "{query}"</p>
            <p class="text-xs text-slate-400 mt-1">Coba gunakan kata kunci lain seperti "berita", "lapor", "renja", atau "struktur".</p>
          </div>
        {/if}
      </div>

      <!-- Footer Help Bar -->
      <div class="px-4 py-2.5 bg-white/40 backdrop-blur-md border-t border-white/80 flex items-center justify-between text-[11px] text-slate-400">
        <div class="flex items-center gap-3">
          <span class="inline-flex items-center gap-1"><kbd class="bg-white/80 px-1.5 py-0.5 rounded border border-white/80 text-[10px] font-mono shadow-2xs">&uarr;&darr;</kbd> Navigasi</span>
          <span class="inline-flex items-center gap-1"><kbd class="bg-white/80 px-1.5 py-0.5 rounded border border-white/80 text-[10px] font-mono shadow-2xs">&crarr;</kbd> Pilih</span>
        </div>
        <span class="font-medium text-[#f08300]">Sistem Pencarian Satpol PP Sulsel</span>
      </div>
    </div>
  </div>
{/if}
