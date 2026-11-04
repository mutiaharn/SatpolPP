<script lang="ts">
  import { onMount } from 'svelte';
  import { trackComplaint } from '$lib/api/client';
  import type { Complaint } from '$lib/types';

  export let initialTicket = '';
  let ticketInput = initialTicket;
  let complaintData: Complaint | null = null;
  let isLoading = false;
  let errorMessage = '';

  onMount(() => {
    if (initialTicket) {
      handleSearch();
    }
  });

  async function handleSearch() {
    if (!ticketInput.trim()) return;
    isLoading = true;
    errorMessage = '';
    complaintData = null;

    try {
      const res = await trackComplaint(ticketInput.trim());
      if (res.success) {
        complaintData = res.data;
      } else {
        errorMessage = 'Nomor tiket tidak ditemukan dalam sistem database.';
      }
    } catch (e: any) {
      errorMessage = 'Gagal menghubungi server backend: ' + e.message;
    } finally {
      isLoading = false;
    }
  }

  function getStatusStyle(status: string) {
    switch (status) {
      case 'MENUNGGU_VERIFIKASI':
        return { label: 'Menunggu Verifikasi', color: 'bg-amber-100 text-amber-900 border-amber-300' };
      case 'DISPOSISI_BIDANG':
        return { label: 'Disposisi Bidang Terkait', color: 'bg-blue-100 text-blue-900 border-blue-300' };
      case 'PENINDAKAN_LAPANGAN':
        return { label: 'Sedang Penindakan Lapangan', color: 'bg-purple-100 text-purple-900 border-purple-300' };
      case 'SELESAI':
        return { label: 'Laporan Dituntaskan', color: 'bg-emerald-100 text-emerald-900 border-emerald-300' };
      case 'DITOLAK':
        return { label: 'Ditolak', color: 'bg-red-100 text-red-900 border-red-300' };
      default:
        return { label: status, color: 'bg-gray-100 text-gray-800 border-gray-300' };
    }
  }
</script>

<div class="bg-white border border-[#eee9db] rounded-2xl p-6 md:p-8 shadow-xs">
  <div class="max-w-xl mx-auto text-center space-y-1 mb-6">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">
      Sistem Pelacakan Pengaduan
    </div>
    <h3 class="text-xl font-bold text-[#101116]">Lacak Status Penanganan Aduan</h3>
    <p class="text-xs text-[#6f7775]">
      Silakan masukkan nomor resi tiket pengaduan yang Anda dapatkan saat melapor di SIAP TERTIB.
    </p>
  </div>

  <!-- Search Input Form -->
  <form on:submit|preventDefault={handleSearch} class="max-w-md mx-auto flex gap-2 mb-6">
    <input 
      type="text" 
      bind:value={ticketInput} 
      placeholder="Contoh: TRB-202611-SAMPLE" 
      class="flex-1 border border-gray-300 rounded-xl px-3.5 py-2.5 text-xs font-mono font-bold uppercase tracking-wider focus:border-[#F08300] focus:outline-none"
    />
    <button 
      type="submit" 
      disabled={isLoading}
      class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-5 py-2.5 rounded-xl transition disabled:opacity-50 shadow-xs"
    >
      {#if isLoading}Mencari...{:else}Cari Tiket{/if}
    </button>
  </form>

  {#if errorMessage}
    <div class="max-w-md mx-auto bg-red-50 border border-red-200 text-red-800 text-xs p-3.5 rounded-xl text-center font-semibold">
      {errorMessage}
    </div>
  {/if}

  {#if complaintData}
    {@const badge = getStatusStyle(complaintData.status)}
    <div class="max-w-2xl mx-auto mt-6 bg-[#f9f4e8] border border-[#eee9db] rounded-2xl p-6 space-y-4 shadow-xs">
      <div class="flex flex-wrap items-center justify-between gap-3 border-b border-[#eee9db] pb-4">
        <div>
          <div class="text-[10px] text-[#6f7775] uppercase tracking-wider font-semibold">Nomor Resi Tiket</div>
          <div class="text-xl font-bold text-[#101116] font-mono">{complaintData.ticket_code}</div>
        </div>
        <div class={`px-3 py-1 rounded-lg text-xs font-bold border uppercase tracking-wider ${badge.color}`}>
          {badge.label}
        </div>
      </div>

      <!-- Details Table -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-3 text-xs">
        <div>
          <span class="text-[#6f7775]">Jenis Pelanggaran:</span>
          <div class="font-bold text-[#101116] mt-0.5">{complaintData.violation_type}</div>
        </div>
        <div>
          <span class="text-[#6f7775]">Wilayah / Lokasi:</span>
          <div class="font-bold text-[#101116] mt-0.5">{complaintData.regency_name}, Kec. {complaintData.district_name || '-'}</div>
        </div>
        <div>
          <span class="text-[#6f7775]">Identitas Pelapor:</span>
          {#if complaintData.is_anonymous}
            <div class="font-bold text-[#F08300] mt-0.5">
              {complaintData.masked_name || 'Warga (Anonim Terlindungi)'}
            </div>
            <div class="text-[10px] text-[#6f7775]">NIK: {complaintData.masked_nik || '7371**********12'}</div>
          {:else}
            <div class="font-bold text-[#101116] mt-0.5">{complaintData.reporter_name}</div>
          {/if}
        </div>
        <div>
          <span class="text-[#6f7775]">Tanggal Lapor:</span>
          <div class="font-bold text-[#101116] mt-0.5">{new Date(complaintData.created_at).toLocaleString('id-ID')} WITA</div>
        </div>
      </div>

      <div class="border-t border-[#eee9db] pt-3">
        <span class="text-xs text-[#6f7775]">Uraian Kejadian:</span>
        <div class="text-xs text-[#101116] bg-white p-3 rounded-xl border border-[#eee9db] mt-1 leading-relaxed">
          {complaintData.description}
        </div>
      </div>

      {#if complaintData.evidence_photos}
        <div class="border-t border-[#eee9db] pt-3">
          <span class="text-xs text-[#6f7775] font-semibold">Lampiran Foto Bukti Kejadian:</span>
          <div class="mt-2">
            <img 
              src={complaintData.evidence_photos} 
              alt="Bukti Aduan" 
              class="max-h-56 rounded-xl border border-gray-300 object-contain bg-black/5" 
            />
          </div>
        </div>
      {/if}

      {#if complaintData.disposed_to || complaintData.action_notes}
        <div class="bg-white border border-[#eee9db] p-4 rounded-xl text-xs space-y-1.5 shadow-xs">
          <div class="font-bold text-[#101116] uppercase tracking-wider text-[11px] border-b border-gray-100 pb-1">
            Catatan Tindak Lanjut Satpol PP Sulsel:
          </div>
          {#if complaintData.disposed_to}
            <div class="text-[#2b2f3f]">Disposisi Bidang: <strong>{complaintData.disposed_to}</strong></div>
          {/if}
          {#if complaintData.action_notes}
            <div class="text-[#2b2f3f]">Status Lapangan / Berita Acara: {complaintData.action_notes}</div>
          {/if}
        </div>
      {/if}
    </div>
  {/if}
</div>
