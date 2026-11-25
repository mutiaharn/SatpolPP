<script lang="ts">
  export let variant: 'default' | 'subtle' | 'strong' | 'interactive' = 'default';
  export let padding: 'none' | 'sm' | 'md' | 'lg' | 'xl' = 'md';
  export let rounded: 'xl' | '2xl' | '3xl' = '2xl';
  export let specular: boolean = true;
  export let className: string = '';

  const paddingClasses: Record<typeof padding, string> = {
    none: 'p-0',
    sm: 'p-3 sm:p-4',
    md: 'p-5 sm:p-6',
    lg: 'p-6 sm:p-8',
    xl: 'p-8 sm:p-10'
  };

  const roundedClasses: Record<typeof rounded, string> = {
    xl: 'rounded-xl',
    '2xl': 'rounded-2xl',
    '3xl': 'rounded-3xl'
  };

  const variantClasses: Record<typeof variant, string> = {
    default:
      'bg-white/75 backdrop-blur-xl border border-white/90 shadow-sm shadow-slate-900/5',
    subtle:
      'bg-white/50 backdrop-blur-md border border-white/60 shadow-sm shadow-slate-900/5',
    strong:
      'bg-white/90 backdrop-blur-2xl border border-white/95 shadow-md shadow-slate-900/5',
    interactive:
      'bg-white/80 backdrop-blur-xl border border-white/90 shadow-sm shadow-slate-900/5 hover:shadow-lg hover:shadow-slate-900/8 hover:-translate-y-0.5 hover:bg-white/95 transition-all duration-200 cursor-pointer'
  };
</script>

<div
  class="relative overflow-hidden {roundedClasses[rounded]} {variantClasses[variant]} {paddingClasses[padding]} {specular ? 'ds-specular-top' : ''} {className}"
  {...$$restProps}
>
  <slot />
</div>
