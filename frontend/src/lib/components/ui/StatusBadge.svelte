<script lang="ts" context="module">
  import type { ComplaintStatus } from '$lib/types';

  export const STATUS_META: Record<ComplaintStatus, { label: string; className: string }> = {
    MENUNGGU_VERIFIKASI: { label: 'Menunggu Verifikasi', className: 'bg-amber-500/10 text-[#d97400] border-amber-500/25' },
    DISPOSISI_BIDANG: { label: 'Disposisi Bidang', className: 'bg-slate-100 text-slate-800 border-slate-200' },
    PENINDAKAN_LAPANGAN: { label: 'Penindakan Lapangan', className: 'bg-amber-500/15 text-[#b45f00] border-amber-500/30' },
    SELESAI: { label: 'Selesai', className: 'bg-slate-900 text-white border-slate-800' },
    DITOLAK: { label: 'Ditolak', className: 'bg-slate-100 text-slate-700 border-slate-300' }
  };

  export function statusLabel(status: string): string {
    return STATUS_META[status as ComplaintStatus]?.label ?? status;
  }
</script>

<script lang="ts">
  export let status: string;
  export let size: 'sm' | 'md' = 'sm';

  $: meta = STATUS_META[status as ComplaintStatus] ?? { label: status, className: 'bg-stone-100 text-stone-700 border-stone-200' };
</script>

<span
  class={`inline-flex items-center whitespace-nowrap font-bold uppercase tracking-wider border rounded-md ${meta.className} ${size === 'sm' ? 'text-[10px] px-2 py-0.5' : 'text-xs px-3 py-1'}`}
>
  {meta.label}
</span>
