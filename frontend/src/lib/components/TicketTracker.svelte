<script lang="ts">
  import { onMount } from 'svelte';
  import { trackComplaint } from '$lib/api/client';
  import type { Complaint } from '$lib/types';
  import StatusBadge from '$lib/components/ui/StatusBadge.svelte';

  export let initialTicket = '';
  export let showHeader = false;
  let ticketInput = initialTicket;
  let complaintData: Complaint | null = null;
  let isLoading = false;
  let errorMessage = '';

  onMount(() => {
    if (initialTicket) handleSearch();
  });

  async function handleSearch() {
    if (!ticketInput.trim()) return;
    isLoading = true;
    errorMessage = '';
    complaintData = null;

    try {
      const res = await trackComplaint(ticketInput.trim());
      if (res.success && res.data) {
        complaintData = res.data;
      } else {
        errorMessage = res.message || 'Nomor tiket tidak ditemukan dalam sistem database.';
      }
    } catch (e: any) {
      errorMessage = 'Gagal menghubungi server backend: ' + e.message;
    } finally {
      isLoading = false;
    }
  }

  $: evidenceUrl = complaintData?.evidence_image_url || complaintData?.evidence_photos || '';
</script>

<div class="ds-glass-strong border border-white/90 rounded-3xl p-4 sm:p-6 shadow-glass ds-specular-top">
  {#if showHeader}
    <div class="max-w-xl mx-auto text-center space-y-1 mb-6">
      <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">
        Sistem Pelacakan Pengaduan
      </div>
      <h3 class="text-xl font-bold text-slate-900">Lacak Status Penanganan Aduan</h3>
      <p class="text-xs text-slate-500 font-light">
        Masukkan nomor resi tiket pengaduan yang Anda terima saat melapor di SIAP TERTIB.
      </p>
    </div>
  {/if}

  <!-- Search Form -->
  <search>
    <form on:submit|preventDefault={handleSearch} class="max-w-md mx-auto flex gap-2">
      <label for="ticket-search-input" class="sr-only">Nomor Tiket Pengaduan</label>
      <input
        id="ticket-search-input"
        type="text"
        bind:value={ticketInput}
        autocorrect="off"
        autocapitalize="characters"
        spellcheck="false"
        placeholder="Contoh: TRB-202611-SAMPLE"
        class="flex-1 ds-glass-input rounded-xl px-4 py-2.5 text-xs font-mono font-bold uppercase tracking-wider text-slate-900 focus:outline-none focus:ring-2 focus:ring-[#F08300]/30"
      />
      <button
        type="submit"
        disabled={isLoading}
        class="bg-gradient-to-r from-[#F08300] to-[#e07000] hover:from-[#d97400] hover:to-[#c66500] text-white font-bold text-xs uppercase px-5 py-2.5 rounded-xl transition disabled:opacity-50 shadow-glass-gold ds-press whitespace-nowrap"
      >
        {#if isLoading}Mencari...{:else}Cari Tiket{/if}
      </button>
    </form>
  </search>

  {#if errorMessage}
    <div role="alert" class="max-w-md mx-auto bg-rose-50/90 border border-rose-200 text-rose-800 text-xs p-3.5 rounded-xl text-center font-semibold backdrop-blur-md mt-4">
      {errorMessage}
    </div>
  {/if}

  {#if complaintData}
    <div class="max-w-2xl mx-auto mt-6 ds-glass border border-white/90 rounded-2xl p-6 space-y-4 shadow-glass">
      <div class="flex flex-wrap items-center justify-between gap-3 border-b border-white/80 pb-4">
        <div>
          <div class="text-[10px] text-slate-500 uppercase tracking-wider font-semibold">Nomor Resi Tiket</div>
          <div class="text-xl font-bold text-slate-900 font-mono">{complaintData.ticket_code}</div>
        </div>
        <StatusBadge status={complaintData.status} size="md" />
      </div>

      <!-- Details Grid (Identitas pelapor selalu disamarkan pada tampilan publik) -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-3 text-xs">
        <div>
          <span class="text-slate-500 font-light">Jenis Pelanggaran:</span>
          <div class="font-bold text-slate-900 mt-0.5">{complaintData.violation_type}</div>
        </div>
        <div>
          <span class="text-slate-500 font-light">Wilayah / Lokasi:</span>
          <div class="font-bold text-slate-900 mt-0.5">
            {complaintData.regency_name}{complaintData.district_name ? `, Kec. ${complaintData.district_name}` : ''}
          </div>
        </div>
        <div>
          <span class="text-slate-500 font-light">Identitas Pelapor (Terlindungi):</span>
          <div class="font-bold text-slate-900 mt-0.5">
            {complaintData.masked_name || complaintData.reporter_name}
          </div>
          <div class="text-[10px] font-mono text-slate-500">NIK: {complaintData.masked_nik || complaintData.reporter_nik || '-'}</div>
        </div>
        <div>
          <span class="text-slate-500 font-light">Tanggal Lapor:</span>
          <div class="font-bold text-slate-900 mt-0.5">{new Date(complaintData.created_at).toLocaleString('id-ID')} WITA</div>
        </div>
      </div>

      <div class="border-t border-white/80 pt-3">
        <span class="text-xs text-slate-500 font-medium">Uraian Kejadian:</span>
        <div class="text-xs text-slate-800 ds-glass-subtle p-3.5 rounded-xl border border-white/80 mt-1 leading-relaxed font-light">
          {complaintData.description}
        </div>
      </div>

      {#if evidenceUrl}
        <div class="border-t border-slate-200/80 pt-3">
          <span class="text-xs text-slate-500 font-semibold">Lampiran Foto Bukti Kejadian:</span>
          <div class="mt-2">
            <img
              src={evidenceUrl}
              alt="Bukti Aduan"
              class="max-h-60 rounded-xl border border-slate-200 object-contain bg-white"
            />
          </div>
        </div>
      {/if}

      {#if complaintData.disposed_to || complaintData.action_notes}
        <div class="bg-white border border-slate-200 p-4 rounded-xl text-xs space-y-1.5 shadow-2xs">
          <div class="font-bold text-slate-900 uppercase tracking-wider text-[11px] border-b border-slate-100 pb-1">
            Catatan Tindak Lanjut Satpol PP Sulsel:
          </div>
          {#if complaintData.disposed_to}
            <div class="text-slate-600">Disposisi Bidang: <strong class="text-slate-900">{complaintData.disposed_to}</strong></div>
          {/if}
          {#if complaintData.action_notes}
            <div class="text-slate-600">Status Lapangan / Berita Acara: <span class="text-slate-800">{complaintData.action_notes}</span></div>
          {/if}
        </div>
      {/if}
    </div>
  {/if}
</div>
