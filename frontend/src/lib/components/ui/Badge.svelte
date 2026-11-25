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
    primary: 'bg-blue-50 text-blue-700 border-blue-200/80',
    success: 'bg-emerald-50 text-emerald-700 border-emerald-200/80',
    warning: 'bg-amber-50 text-amber-800 border-amber-200/80',
    danger: 'bg-rose-50 text-rose-700 border-rose-200/80',
    gold: 'bg-orange-50 text-[#b45f00] border-orange-200/80',
    indigo: 'bg-indigo-50 text-indigo-700 border-indigo-200/80'
  };

  const dotClasses: Record<typeof variant, string> = {
    neutral: 'bg-slate-400',
    primary: 'bg-blue-500',
    success: 'bg-emerald-500',
    warning: 'bg-amber-500',
    danger: 'bg-rose-500',
    gold: 'bg-[#f08300]',
    indigo: 'bg-indigo-500'
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
