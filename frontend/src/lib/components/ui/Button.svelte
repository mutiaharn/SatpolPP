<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let variant: 'primary' | 'secondary' | 'glass' | 'ghost' | 'danger' | 'gold' | 'soft' = 'primary';
  export let size: 'xs' | 'sm' | 'md' | 'lg' = 'md';
  export let type: 'button' | 'submit' | 'reset' = 'button';
  export let disabled: boolean = false;
  export let loading: boolean = false;
  export let fullWidth: boolean = false;
  export let href: string | undefined = undefined;
  export let target: string | undefined = undefined;
  export let rel: string | undefined = undefined;
  export let className: string = '';

  const dispatch = createEventDispatcher<{ click: MouseEvent }>();

  const baseClasses =
    'inline-flex items-center justify-center font-bold select-none whitespace-nowrap shrink-0 transition-all duration-150 ease-out focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-[#246bfd] focus-visible:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed disabled:pointer-events-none active:scale-[0.975]';

  const sizeClasses: Record<typeof size, string> = {
    xs: 'px-3 py-1 text-xs gap-1.5 rounded-lg min-h-[30px]',
    sm: 'px-3.5 py-1.5 text-xs gap-1.5 rounded-xl min-h-[36px]',
    md: 'px-5 py-2.5 text-sm gap-2 rounded-xl min-h-[44px]',
    lg: 'px-6 py-3 text-base gap-2.5 rounded-xl min-h-[48px]'
  };

  const variantClasses: Record<typeof variant, string> = {
    primary:
      'bg-gradient-to-r from-[#246bfd] to-[#1d4ed8] text-white shadow-[0_4px_14px_rgba(36,107,253,0.35),inset_0_1px_0_rgba(255,255,255,0.35)] border border-blue-400/40 hover:from-[#1d4ed8] hover:to-[#1e40af] hover:shadow-[0_6px_20px_rgba(36,107,253,0.45)] active:brightness-95',
    secondary:
      'bg-white/90 text-slate-800 border border-slate-200/90 shadow-xs hover:bg-white hover:border-slate-300 hover:shadow-sm active:bg-slate-50 backdrop-blur-md',
    glass:
      'bg-white/70 text-slate-800 border border-white/80 shadow-[0_2px_8px_rgba(0,0,0,0.04),inset_0_1px_0_rgba(255,255,255,0.8)] hover:bg-white/95 hover:border-white active:bg-white/60 backdrop-blur-md',
    ghost:
      'text-slate-600 hover:text-slate-900 hover:bg-slate-100/80 active:bg-slate-200/60 border border-transparent',
    danger:
      'bg-rose-50 text-rose-700 border border-rose-200 hover:bg-rose-100 hover:border-rose-300 active:bg-rose-200/80 shadow-xs',
    gold:
      'bg-gradient-to-r from-[#f08300] to-[#e07000] text-white shadow-[0_4px_14px_rgba(240,131,0,0.35),inset_0_1px_0_rgba(255,255,255,0.35)] border border-amber-400/40 hover:from-[#d97400] hover:to-[#c66500] hover:shadow-[0_6px_20px_rgba(240,131,0,0.45)] active:brightness-95',
    soft:
      'bg-blue-50/80 text-blue-700 border border-blue-200/60 hover:bg-blue-100 hover:border-blue-300 active:bg-blue-200/60'
  };

  function handleClick(e: MouseEvent) {
    if (disabled || loading) {
      e.preventDefault();
      return;
    }
    dispatch('click', e);
  }
</script>

{#if href}
  <a
    {href}
    {target}
    rel={target === '_blank' ? (rel || 'noopener noreferrer') : rel}
    class="{baseClasses} {sizeClasses[size]} {variantClasses[variant]} {fullWidth ? 'w-full' : ''} {className}"
    on:click={handleClick}
    {...$$restProps}
  >
    {#if loading}
      <svg class="animate-spin -ml-0.5 h-4 w-4 shrink-0" fill="none" viewBox="0 0 24 24" aria-hidden="true">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="3" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
      </svg>
    {/if}
    <slot name="iconLeading" />
    <slot />
    <slot name="iconTrailing" />
  </a>
{:else}
  <button
    {type}
    disabled={disabled || loading}
    class="{baseClasses} {sizeClasses[size]} {variantClasses[variant]} {fullWidth ? 'w-full' : ''} {className}"
    on:click={handleClick}
    {...$$restProps}
  >
    {#if loading}
      <svg class="animate-spin -ml-0.5 h-4 w-4 shrink-0" fill="none" viewBox="0 0 24 24" aria-hidden="true">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="3" />
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
      </svg>
    {/if}
    <slot name="iconLeading" />
    <slot />
    <slot name="iconTrailing" />
  </button>
{/if}
