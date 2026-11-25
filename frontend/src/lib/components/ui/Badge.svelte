<script lang="ts">
  export let variant: 'neutral' | 'primary' | 'success' | 'warning' | 'danger' | 'gold' | 'indigo' = 'neutral';
  export let size: 'sm' | 'md' = 'sm';
  export let dot: boolean = false;
  export let className: string = '';

  const sizeClasses: Record<typeof size, string> = {
    sm: 'text-[10px] px-2 py-0.5 rounded-md gap-1 font-semibold uppercase tracking-wider',
    md: 'text-xs px-2.5 py-1 rounded-lg gap-1.5 font-medium'
  };

  const variantClasses: Record<typeof variant, string> = {
    neutral: 'bg-slate-100 text-slate-700 border-slate-200/80',
    primary: 'bg-amber-500/10 text-[#d97400] border-amber-500/25',
    success: 'bg-slate-900 text-white border-slate-800',
    warning: 'bg-amber-500/10 text-[#d97400] border-amber-500/25',
    danger: 'bg-slate-100 text-slate-800 border-slate-300',
    gold: 'bg-amber-500/10 text-[#d97400] border-amber-500/25',
    indigo: 'bg-slate-100 text-slate-700 border-slate-200/80'
  };

  const dotClasses: Record<typeof variant, string> = {
    neutral: 'bg-slate-400',
    primary: 'bg-[#f08300]',
    success: 'bg-white',
    warning: 'bg-[#f08300]',
    danger: 'bg-slate-600',
    gold: 'bg-[#f08300]',
    indigo: 'bg-slate-400'
  };
</script>

<span
  class="inline-flex items-center border whitespace-nowrap select-none {sizeClasses[size]} {variantClasses[variant]} {className}"
  {...$$restProps}
>
  {#if dot}
    <span class="w-1.5 h-1.5 rounded-full {dotClasses[variant]}" aria-hidden="true"></span>
  {/if}
  <slot />
</span>
