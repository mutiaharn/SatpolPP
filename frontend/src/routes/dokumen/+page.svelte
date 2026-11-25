<script lang="ts">
  import { page } from '$app/stores';
  import { documentsStore } from '$lib/stores/appState';
  import type { PublicDocument } from '$lib/types';
  import DataTable from '$lib/components/ui/DataTable.svelte';
  import Modal from '$lib/components/ui/Modal.svelte';
  import { Button, Badge } from '$lib/components/ui';
  import { ChevronRight, FileText } from 'lucide-svelte';

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

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#246bfd]/8 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <!-- Breadcrumb Pills -->
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-slate-400">Informasi Publik</span>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">Dokumen &amp; Regulasi</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">Dokumen &amp; Regulasi Resmi</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      Akses terbuka terhadap rencana kerja, transparansi anggaran, laporan kinerja instansi, serta peraturan daerah untuk mewujudkan tata kelola pemerintahan yang bersih dan akuntabel.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-6">
  <!-- Category Filter Pills (Apple HIG Glass Capsules) -->
  <div class="flex flex-wrap gap-2 ds-glass p-4 sm:p-5 border border-white/90 rounded-3xl shadow-glass ds-specular-top">
    {#each categories as cat}
      <button 
        type="button"
        class={`px-4 py-2.5 min-h-[40px] inline-flex items-center justify-center text-xs font-bold uppercase tracking-wider transition duration-200 rounded-xl ds-press ${selectedCategory === cat.id ? 'ds-pill-active' : 'ds-pill-inactive'}`}
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
      <tr class="hover:bg-slate-50/80 transition">
        <td class="p-4 font-bold text-slate-900 max-w-md">
          <button
            type="button"
            on:click={() => openDocDetail(item)}
            class="text-left hover:text-[#F08300] transition flex items-start gap-2.5"
          >
            <svg class="w-4 h-4 text-slate-400 mt-0.5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 002-2V9.414a1 1 0 00-.293-.707l-5.414-5.414A1 1 0 0012.586 3H7a2 2 0 00-2 2v14a2 2 0 002 2z" />
            </svg>
            <span>{item.title}</span>
          </button>
        </td>
        <td class="p-4 text-slate-500 font-mono text-xs">{item.doc_number || '-'}</td>
        <td class="p-4">
          <Badge variant="default" size="sm">
            {item.category}
          </Badge>
        </td>
        <td class="p-4 text-center text-slate-900 font-semibold">{item.year}</td>
        <td class="p-4 text-right">
          <div class="inline-flex items-center gap-2 justify-end">
            <!-- Modal Detail Trigger -->
            <Button
              variant="secondary"
              size="xs"
              className="uppercase tracking-wider font-bold"
              on:click={() => openDocDetail(item)}
            >
              Detail
            </Button>

            <!-- Direct Download Button -->
            <Button 
              href={item.file_url} 
              target="_blank" 
              rel="noreferrer"
              download={item.title}
              variant="primary"
              size="xs"
              className="uppercase tracking-wider font-bold inline-flex items-center gap-1 shadow-xs"
            >
              <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
              </svg>
              <span>Unduh</span>
            </Button>
          </div>
        </td>
      </tr>
    </svelte:fragment>
  </DataTable>
</div>

<!-- Modal Pratinjau Dokumen Publik -->
<Modal
  bind:open={showPreviewModal}
  title={selectedDoc?.title || 'Rincian Dokumen Resmi'}
  subtitle={`DOKUMEN RESMI SATPOL PP SULSEL • TAHUN ${selectedDoc?.year || 2026}`}
  size="lg"
>
  {#if selectedDoc}
    <div class="space-y-5">
      <!-- Metadata Grid -->
      <div class="grid grid-cols-1 sm:grid-cols-2 gap-3.5 bg-slate-50/80 p-4 rounded-2xl border border-slate-200/80">
        <div>
          <span class="text-[10px] text-slate-400 uppercase font-semibold block">Nomor Registrasi</span>
          <span class="font-mono text-xs font-bold text-slate-900">{selectedDoc.doc_number || 'Tidak tertera'}</span>
        </div>
        <div>
          <span class="text-[10px] text-slate-400 uppercase font-semibold block">Kategori Berkas</span>
          <span class="text-xs font-bold text-[#F08300]">{getCategoryLabel(selectedDoc.category)}</span>
        </div>
        <div>
          <span class="text-[10px] text-slate-400 uppercase font-semibold block">Tahun Penerbitan</span>
          <span class="text-xs font-bold text-slate-900">{selectedDoc.year}</span>
        </div>
        <div>
          <span class="text-[10px] text-slate-400 uppercase font-semibold block">Estimasi Ukuran File</span>
          <span class="text-xs font-bold text-slate-900">{selectedDoc.file_size || '1.8 MB'}</span>
        </div>
      </div>

      <!-- Description / Disclaimer -->
      <div class="border-l-4 border-[#F08300] bg-amber-50/70 p-4 rounded-r-2xl border border-amber-200/50 text-slate-700 leading-relaxed text-xs">
        <p class="font-semibold text-slate-900 mb-1">Keterbukaan Informasi Publik (UU No. 14 Tahun 2008)</p>
        <p class="font-light">
          Dokumen ini dipublikasikan secara resmi oleh Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan sebagai bentuk akuntabilitas kinerja, perencanaan strategis, dan transparansi regulasi bagi seluruh masyarakat.
        </p>
      </div>

      <!-- Action Box -->
      <div class="p-4 bg-white border border-slate-200 rounded-2xl flex items-center justify-between gap-3 shadow-2xs">
        <div class="flex items-center gap-3">
          <div class="w-10 h-10 rounded-xl bg-red-50 text-red-600 flex items-center justify-center font-bold text-xs border border-red-200">
            PDF
          </div>
          <div>
            <div class="font-bold text-xs text-slate-900">{selectedDoc.title}.pdf</div>
            <div class="text-[11px] text-slate-500 font-light">Berkas digital resmi terverifikasi</div>
          </div>
        </div>

        <Button
          href={selectedDoc.file_url}
          target="_blank"
          rel="noreferrer"
          download={selectedDoc.title}
          variant="gold"
          size="md"
          className="uppercase tracking-wider font-bold inline-flex items-center gap-1.5 shadow-xs"
        >
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
          </svg>
          <span>Unduh Berkas</span>
        </Button>
      </div>
    </div>
  {/if}

  <svelte:fragment slot="footer" let:handleClose>
    <Button
      variant="secondary"
      size="sm"
      className="font-semibold"
      on:click={handleClose}
    >
      Tutup
    </Button>
  </svelte:fragment>
</Modal>
