<script lang="ts">
  import { page } from '$app/stores';
  import { documentsStore } from '$lib/stores/appState';
  import type { PublicDocument } from '$lib/types';
  import DataTable from '$lib/components/ui/DataTable.svelte';
  import Modal from '$lib/components/ui/Modal.svelte';

  let selectedCategory = '';
  let selectedDoc: PublicDocument | null = null;
  let showPreviewModal = false;

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

  // Filter based on selected category (DataTable handles real-time search on-chain)
  $: categoryDocs = selectedCategory
    ? $documentsStore.filter(d => d.category.toLowerCase() === selectedCategory.toLowerCase())
    : $documentsStore;

  function openDocDetail(doc: PublicDocument) {
    selectedDoc = doc;
    showPreviewModal = true;
  }

  function getCategoryLabel(catId: string): string {
    const found = categories.find(c => c.id === catId);
    return found ? found.label : catId.toUpperCase();
  }
</script>

<svelte:head>
  <title>Dokumen Publik &amp; Transparansi Kinerja | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Transparansi Penyelenggaraan Negara</div>
    <h1 class="text-3xl font-extrabold text-[#1a120c] tracking-tight mt-1">Dokumen &amp; Regulasi Publik Resmi</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3 rounded-full"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Akses berkas perencanaan, penganggaran tahunan, laporan akuntabilitas kinerja (LAKIP/LKjIP), serta peraturan daerah ketertiban umum Satpol PP Prov. Sulsel.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-6">
  <!-- Category Filter Pills -->
  <div class="flex flex-wrap gap-2 bg-white p-4 sm:p-5 border border-[#eee9db] rounded-xl shadow-xs">
    {#each categories as cat}
      <button 
        type="button"
        class={`px-3.5 py-2 text-xs font-bold uppercase transition rounded-lg ${selectedCategory === cat.id ? 'bg-[#3d261a] text-white shadow-xs' : 'bg-stone-50 border border-stone-200 text-[#1a120c] hover:bg-stone-100'}`}
        on:click={() => selectedCategory = cat.id}
      >
        {cat.label}
      </button>
    {/each}
  </div>

  <!-- Reusable DataTable Component -->
  <DataTable
    items={categoryDocs}
    searchFields={['title', 'doc_number', 'category', 'year']}
    searchPlaceholder="Cari nama, nomor regulasi, atau tahun..."
    pageSize={10}
    pageSizeOptions={[5, 10, 25, 50]}
    title="Daftar Dokumen Resmi"
    description="Tabel arsip transparansi publik dan regulasi Satpol PP Sulsel"
    emptyMessage="Tidak ada dokumen yang sesuai dengan kata kunci atau kategori terpilih."
  >
    <svelte:fragment slot="head">
      <tr>
        <th class="p-4">Nama Dokumen Resmi</th>
        <th class="p-4">Nomor Registrasi</th>
        <th class="p-4">Kategori Berkas</th>
        <th class="p-4 text-center">Tahun</th>
        <th class="p-4 text-right">Aksi</th>
      </tr>
    </svelte:fragment>

    <svelte:fragment slot="row" let:item>
      <tr class="hover:bg-stone-50/70 transition">
        <td class="p-4 font-bold text-[#1a120c] max-w-md">
          <button
            type="button"
            on:click={() => openDocDetail(item)}
            class="text-left hover:text-[#F08300] transition flex items-start gap-2"
          >
            <svg class="w-4 h-4 text-stone-400 mt-0.5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
            </svg>
            <span>{item.title}</span>
          </button>
        </td>
        <td class="p-4 text-[#6f7775] font-mono text-xs">{item.doc_number || '-'}</td>
        <td class="p-4">
          <span class="bg-[#f5ede4] text-[#3d261a] font-bold text-[10px] uppercase px-2.5 py-1 rounded-md border border-[#e8dbc9]">
            {item.category}
          </span>
        </td>
        <td class="p-4 text-center text-[#1a120c] font-semibold">{item.year}</td>
        <td class="p-4 text-right">
          <div class="inline-flex items-center gap-1.5 justify-end">
            <!-- Modal Detail Trigger -->
            <button
              type="button"
              on:click={() => openDocDetail(item)}
              class="bg-stone-100 hover:bg-stone-200 text-stone-700 font-bold text-[11px] uppercase px-3 py-1.5 rounded-lg transition"
              title="Pratinjau Rincian"
            >
              Detail
            </button>

            <!-- Direct Download Button -->
            <a 
              href={item.file_url} 
              target="_blank" 
              rel="noreferrer"
              download={item.title}
              class="bg-[#3d261a] hover:bg-[#271810] text-white font-bold text-[11px] uppercase px-3 py-1.5 rounded-lg transition inline-flex items-center gap-1"
            >
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
              </svg>
              <span>Unduh</span>
            </a>
          </div>
        </td>
      </tr>
    </svelte:fragment>
  </DataTable>
</div>

<!-- Modal Pratinjau Dokumen Publik (Poin 15) -->
<Modal
  bind:open={showPreviewModal}
  title={selectedDoc?.title || 'Rincian Dokumen Resmi'}
  subtitle={`DOKUMEN RESMI SATPOL PP SULSEL • TAHUN ${selectedDoc?.year || 2026}`}
  size="lg"
>
  {#if selectedDoc}
    <div class="space-y-5">
      <!-- Metadata Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-3.5 bg-[#faf7f2] p-4 rounded-xl border border-[#eee9db]">
        <div>
          <span class="text-[10px] text-stone-500 uppercase font-semibold block">Nomor Registrasi</span>
          <span class="font-mono text-xs font-bold text-[#1a120c]">{selectedDoc.doc_number || 'Tidak tertera'}</span>
        </div>
        <div>
          <span class="text-[10px] text-stone-500 uppercase font-semibold block">Kategori Berkas</span>
          <span class="text-xs font-bold text-[#F08300]">{getCategoryLabel(selectedDoc.category)}</span>
        </div>
        <div>
          <span class="text-[10px] text-stone-500 uppercase font-semibold block">Tahun Penerbitan</span>
          <span class="text-xs font-bold text-[#1a120c]">{selectedDoc.year}</span>
        </div>
        <div>
          <span class="text-[10px] text-stone-500 uppercase font-semibold block">Estimasi Ukuran File</span>
          <span class="text-xs font-bold text-[#1a120c]">{selectedDoc.file_size || '1.8 MB'}</span>
        </div>
      </div>

      <!-- Description / Disclaimer -->
      <div class="border-l-4 border-[#F08300] bg-stone-50 p-3.5 rounded-r-lg text-stone-700 leading-relaxed text-xs">
        <p class="font-semibold text-[#1a120c] mb-1">Keterbukaan Informasi Publik (UU No. 14 Tahun 2008)</p>
        <p>
          Dokumen ini dipublikasikan secara resmi oleh Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan sebagai bentuk akuntabilitas kinerja, perencanaan strategis, dan transparansi regulasi bagi seluruh masyarakat.
        </p>
      </div>

      <!-- Action Box -->
      <div class="p-4 bg-white border border-stone-200 rounded-xl flex items-center justify-between gap-3">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-lg bg-red-50 text-red-600 flex items-center justify-center font-bold text-xs border border-red-200">
            PDF
          </div>
          <div>
            <div class="font-bold text-xs text-[#1a120c]">{selectedDoc.title}.pdf</div>
            <div class="text-[11px] text-stone-500">Berkas digital resmi terverifikasi</div>
          </div>
        </div>

        <a
          href={selectedDoc.file_url}
          target="_blank"
          rel="noreferrer"
          download={selectedDoc.title}
          class="bg-[#3d261a] hover:bg-[#271810] text-white font-bold text-xs uppercase px-4 py-2.5 rounded-lg transition inline-flex items-center gap-1.5 shadow-xs"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
          </svg>
          <span>Unduh Berkas</span>
        </a>
      </div>
    </div>
  {/if}

  <svelte:fragment slot="footer" let:handleClose>
    <button
      type="button"
      on:click={handleClose}
      class="px-4 py-2 text-xs font-semibold text-stone-600 hover:text-stone-800 bg-stone-100 hover:bg-stone-200 rounded-lg transition"
    >
      Tutup
    </button>
  </svelte:fragment>
</Modal>
