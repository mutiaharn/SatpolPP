<script lang="ts">
  import { onMount, tick } from 'svelte';
  import { trackComplaint } from '$lib/api/client';
  import type { Complaint } from '$lib/types';
  import { Button, CopyButton, GlassCard, StatusBadge } from '$lib/components/ui';
  import { Search } from 'lucide-svelte';

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
        await tick();
        scrollToResult();
      } else {
        errorMessage = res.message || 'Nomor tiket tidak ditemukan dalam sistem database.';
      }
    } catch (e: any) {
      errorMessage = 'Gagal menghubungi server backend: ' + e.message;
    } finally {
      isLoading = false;
    }
  }

  function scrollToResult() {
    if (typeof window === 'undefined') return;
    const el = document.getElementById('ticket-result-card');
    if (!el) return;

    const rect = el.getBoundingClientRect();
    const footer = document.querySelector('footer');
    const footerHeight = footer ? Math.round(footer.getBoundingClientRect().height) : 48;
    const cardBottom = rect.bottom;
    const viewportHeight = window.innerHeight;

    // Jika bagian bawah kartu dan footer berada di luar batas bawah layar,
    // gulir secukupnya agar seluruh kartu dan footer terlihat nyaman tanpa perlu scroll manual
    if (cardBottom + footerHeight > viewportHeight) {
      const scrollDelta = (cardBottom + footerHeight + 16) - viewportHeight;
      window.scrollBy({
        top: scrollDelta,
        behavior: 'smooth'
      });
    }
  }

  $: evidenceUrl = complaintData?.evidence_image_url || complaintData?.evidence_photos || '';
</script>

<GlassCard variant="strong" padding="md" rounded="3xl" className="w-full">
  {#if showHeader}
    <div class="max-w-xl mx-auto text-center space-y-1 mb-5">
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
    <form on:submit|preventDefault={handleSearch} class="max-w-xl mx-auto flex gap-2">
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
      <Button
        type="submit"
        variant="gold"
        size="md"
        loading={isLoading}
        className="font-bold whitespace-nowrap px-6 rounded-xl shadow-glass-gold"
      >
        <Search class="w-4 h-4" />
        <span>Cari Tiket</span>
      </Button>
    </form>
  </search>

  {#if errorMessage}
    <div role="alert" class="max-w-md mx-auto bg-rose-50/90 border border-rose-200 text-rose-800 text-xs p-3 rounded-xl text-center font-semibold backdrop-blur-md mt-4">
      {errorMessage}
    </div>
  {/if}

  {#if complaintData}
    <div id="ticket-result-card" class="mt-4 w-full animate-in fade-in zoom-in-95 duration-200">
      <GlassCard variant="default" padding="md" rounded="2xl" className="w-full space-y-3.5 shadow-glass">
        <!-- Header Bar: Resi & Status -->
        <div class="flex flex-wrap items-center justify-between gap-3 border-b border-white/80 pb-3">
          <div>
            <div class="text-[10px] text-slate-500 uppercase tracking-wider font-semibold">Nomor Resi Tiket</div>
            <div class="flex items-center gap-2 mt-0.5">
              <span class="text-xl sm:text-2xl font-bold text-slate-900 font-mono select-all tracking-tight">{complaintData.ticket_code}</span>
              <CopyButton text={complaintData.ticket_code} label="Salin" size="xs" variant="glass" />
            </div>
          </div>
          <StatusBadge status={complaintData.status} size="md" />
        </div>

        <!-- 2-Column Responsive Dashboard Layout (Wider, Compact Height) -->
        <div class="grid grid-cols-1 md:grid-cols-12 gap-4 text-xs text-left">
          <!-- Col Left: Meta & Status Lapangan -->
          <div class="md:col-span-6 space-y-3">
            <div class="grid grid-cols-2 gap-2.5 p-3 rounded-xl ds-glass-subtle border border-white/80">
              <div>
                <span class="text-slate-500 font-light block text-[11px]">Jenis Pelanggaran:</span>
                <div class="font-bold text-slate-900 mt-0.5 leading-snug">{complaintData.violation_type}</div>
              </div>
              <div>
                <span class="text-slate-500 font-light block text-[11px]">Wilayah / Lokasi:</span>
                <div class="font-bold text-slate-900 mt-0.5 leading-snug">
                  {complaintData.regency_name}{complaintData.district_name ? `, Kec. ${complaintData.district_name}` : ''}
                </div>
              </div>
              <div>
                <span class="text-slate-500 font-light block text-[11px]">Identitas Pelapor:</span>
                <div class="font-bold text-slate-900 mt-0.5">
                  {complaintData.masked_name || complaintData.reporter_name}
                </div>
                <div class="text-[10px] font-mono text-slate-500">NIK: {complaintData.masked_nik || complaintData.reporter_nik || '-'}</div>
              </div>
              <div>
                <span class="text-slate-500 font-light block text-[11px]">Tanggal Lapor:</span>
                <div class="font-bold text-slate-900 mt-0.5">{new Date(complaintData.created_at).toLocaleString('id-ID')} WITA</div>
              </div>
            </div>

            {#if complaintData.disposed_to || complaintData.action_notes}
              <div class="bg-white/95 border border-slate-200/90 p-3 rounded-xl text-xs space-y-1 shadow-2xs">
                <div class="font-bold text-slate-900 uppercase tracking-wider text-[10.5px] border-b border-slate-100 pb-1 flex items-center justify-between">
                  <span>Catatan Tindak Lanjut Satpol PP Sulsel:</span>
                  <span class="text-[10px] text-[#F08300] font-semibold lowercase">resmi</span>
                </div>
                {#if complaintData.disposed_to}
                  <div class="text-slate-600">Disposisi Bidang: <strong class="text-slate-900">{complaintData.disposed_to}</strong></div>
                {/if}
                {#if complaintData.action_notes}
                  <div class="text-slate-600">Status Lapangan / Berita Acara: <span class="text-slate-800 font-medium">{complaintData.action_notes}</span></div>
                {/if}
              </div>
            {/if}
          </div>

          <!-- Col Right: Uraian Kejadian & Lampiran Foto -->
          <div class="md:col-span-6 space-y-3 flex flex-col justify-between">
            <div class="space-y-1">
              <span class="text-[11px] text-slate-500 font-semibold uppercase tracking-wider">Uraian Kejadian:</span>
              <div class="text-xs text-slate-800 ds-glass-subtle p-3 rounded-xl border border-white/80 leading-relaxed font-light max-h-36 overflow-y-auto">
                {complaintData.description}
              </div>
            </div>

            {#if evidenceUrl}
              <div class="pt-2 border-t border-slate-100">
                <span class="text-[11px] text-slate-500 font-semibold">Lampiran Bukti Foto:</span>
                <div class="mt-1">
                  <img
                    src={evidenceUrl}
                    alt="Bukti Aduan"
                    class="h-24 max-w-full rounded-xl border border-slate-200 object-cover bg-white shadow-2xs"
                  />
                </div>
              </div>
            {/if}
          </div>
        </div>
      </GlassCard>
    </div>
  {/if}
</GlassCard>
