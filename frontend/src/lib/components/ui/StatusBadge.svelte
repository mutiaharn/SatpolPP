<script lang="ts" context="module">
  import type { ComplaintStatus } from '$lib/types';

  export const STATUS_META: Record<ComplaintStatus, { label: string; className: string }> = {
    MENUNGGU_VERIFIKASI: { label: 'Menunggu Verifikasi', className: 'bg-amber-50 text-amber-800 border-amber-200' },
    DISPOSISI_BIDANG: { label: 'Disposisi Bidang', className: 'bg-[#f5ede4] text-[#3d261a] border-[#e8dbc9]' },
    PENINDAKAN_LAPANGAN: { label: 'Penindakan Lapangan', className: 'bg-orange-50 text-[#b45f00] border-orange-200' },
    SELESAI: { label: 'Selesai', className: 'bg-emerald-50 text-emerald-800 border-emerald-200' },
    DITOLAK: { label: 'Ditolak', className: 'bg-rose-50 text-rose-800 border-rose-200' }
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
