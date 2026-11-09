<script lang="ts">
  import { page } from '$app/stores';
  import { documentsStore } from '$lib/stores/appState';
  import type { PublicDocument } from '$lib/types';

  let selectedCategory = '';
  let searchQuery = '';

  const categories = [
    { id: '', label: 'Semua Dokumen' },
    { id: 'renja', label: 'Rencana Kerja (Renja)' },
    { id: 'renstra', label: 'Rencana Strategis (Renstra)' },
    { id: 'dpa', label: 'Dokumen Pelaksanaan Anggaran (DPA)' },
    { id: 'lakip', label: 'LAKIP / LKjIP' },
    { id: 'keuangan', label: 'Laporan Keuangan' },
    { id: 'perda', label: 'Peraturan Daerah (Perda)' },
    { id: 'sop', label: 'Standar Operasional (SOP)' }
  ];

  $: queryCat = $page.url.searchParams.get('category');
  $: if (queryCat !== null && queryCat !== selectedCategory) {
    selectedCategory = queryCat;
  }

  $: filteredDocs = $documentsStore.filter(d => {
    const matchCat = !selectedCategory || d.category.toLowerCase() === selectedCategory.toLowerCase();
    if (!matchCat) return false;
    if (!searchQuery) return true;
    const q = searchQuery.toLowerCase();
    return d.title.toLowerCase().includes(q) || (d.doc_number && d.doc_number.toLowerCase().includes(q));
  });
</script>

<svelte:head>
  <title>Dokumen Publik &amp; Transparansi Kinerja | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Transparansi Penyelenggaraan Negara</div>
    <h1 class="text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Dokumen &amp; Regulasi Publik Resmi</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Akses berkas perencanaan, penganggaran tahunan, laporan akuntabilitas kinerja (LAKIP/LKjIP), serta peraturan daerah ketertiban umum Satpol PP Prov. Sulsel.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-6">
  <!-- Filter & Search Bar -->
  <div class="flex flex-col md:flex-row justify-between items-stretch md:items-center gap-4 bg-white p-5 border border-[#eee9db] rounded-2xl shadow-xs">
    <div class="flex flex-wrap gap-2">
      {#each categories as cat}
        <button 
          type="button"
          class={`px-3.5 py-2 text-xs font-bold uppercase transition rounded-xl ${selectedCategory === cat.id ? 'bg-[#F08300] text-white shadow-xs' : 'bg-gray-50 border border-gray-200 text-[#101116] hover:bg-gray-100'}`}
          on:click={() => selectedCategory = cat.id}
        >
          {cat.label}
        </button>
      {/each}
    </div>

    <div class="w-full md:w-64">
      <input 
        type="text" 
        bind:value={searchQuery}
        placeholder="Cari nama atau nomor dokumen..."
        class="w-full border border-gray-300 rounded-xl px-3.5 py-2 text-xs text-[#101116] focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] focus:outline-none"
      />
    </div>
  </div>

  <!-- Documents Table -->
  <div class="bg-white border border-[#eee9db] rounded-2xl shadow-xs overflow-hidden">
    <div class="overflow-x-auto">
      <table class="w-full text-left text-xs">
        <thead class="bg-gray-50 border-b border-gray-200 text-[#6f7775] uppercase text-[10px] tracking-wider">
          <tr>
            <th class="p-4">Nama Dokumen Resmi</th>
            <th class="p-4">Nomor Registrasi</th>
            <th class="p-4">Kategori Berkas</th>
            <th class="p-4 text-center">Tahun</th>
            <th class="p-4 text-right">Aksi</th>
          </tr>
        </thead>
        <tbody class="divide-y divide-gray-100">
          {#each filteredDocs as doc}
            <tr class="hover:bg-gray-50/70 transition">
              <td class="p-4 font-bold text-[#101116] max-w-md">{doc.title}</td>
              <td class="p-4 text-[#6f7775] font-mono text-xs">{doc.doc_number || '-'}</td>
              <td class="p-4">
                <span class="bg-[#f9f4e8] text-[#101116] font-bold text-[10px] uppercase px-2.5 py-1 rounded-md border border-[#eee9db]">
                  {doc.category}
                </span>
              </td>
              <td class="p-4 text-center text-[#101116] font-semibold">{doc.year}</td>
              <td class="p-4 text-right">
                <a 
                  href={doc.file_url} 
                  target="_blank" 
                  rel="noreferrer"
                  download={doc.title}
                  class="inline-block bg-[#3d261a] hover:bg-[#271810] text-white font-bold text-xs uppercase px-4 py-2 rounded-lg transition"
                >
                  Unduh ({doc.file_size || 'Berkas'})
                </a>
              </td>
            </tr>
          {:else}
            <tr>
              <td colspan="5" class="p-8 text-center text-gray-400 text-xs">
                Tidak ditemukan dokumen resmi yang sesuai dengan filter atau kata kunci.
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>
</div>
