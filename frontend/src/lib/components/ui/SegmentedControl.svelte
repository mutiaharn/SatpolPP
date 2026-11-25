<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export interface SegmentOption {
    value: string;
    label: string;
    badge?: string | number;
    disabled?: boolean;
  }

  export let options: SegmentOption[] = [];
  export let value: string = '';
  export let size: 'sm' | 'md' | 'lg' = 'md';
  export let fullWidth: boolean = true;
  export let className: string = '';

  const dispatch = createEventDispatcher<{ change: string }>();

  function selectOption(val: string) {
    if (value !== val) {
      value = val;
      dispatch('change', val);
    }
  }

  const containerSizes: Record<typeof size, string> = {
    sm: 'p-0.5 rounded-xl text-xs gap-1',
    md: 'p-1 rounded-2xl text-xs sm:text-sm gap-1',
    lg: 'p-1.5 rounded-2xl text-sm gap-1.5'
  };

  const itemSizes: Record<typeof size, string> = {
    sm: 'py-1 px-2.5 rounded-lg',
    md: 'py-2 px-3 rounded-xl',
    lg: 'py-2.5 px-4 rounded-xl'
  };
</script>

<div
  role="tablist"
  class="relative inline-flex items-center bg-slate-100/95 border border-slate-200/70 backdrop-blur-md select-none {containerSizes[size]} {fullWidth ? 'w-full' : ''} {className}"
>
  {#each options as opt}
    {@const active = opt.value === value}
    <button
      type="button"
      role="tab"
      aria-selected={active}
      disabled={opt.disabled}
      on:click={() => selectOption(opt.value)}
      class="relative flex-1 inline-flex items-center justify-center font-semibold transition-all duration-200 ease-out focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 disabled:opacity-40 disabled:cursor-not-allowed active:scale-[0.98] {itemSizes[size]} {active
        ? 'bg-white text-slate-900 shadow-sm shadow-slate-900/10 border border-slate-200/80 ds-specular-top'
        : 'text-slate-500 hover:text-slate-800 hover:bg-white/50 border border-transparent'}"
    >
      <span class="truncate">{opt.label}</span>
      {#if opt.badge !== undefined && opt.badge !== ''}
        <span
          class="ml-1.5 px-1.5 py-0.5 text-[10px] rounded-full font-bold {active
            ? 'bg-blue-50 text-blue-700'
            : 'bg-slate-200/80 text-slate-600'}"
        >
          {opt.badge}
        </span>
      {/if}
    </button>
  {/each}
</div>
