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
    'inline-flex items-center justify-center font-semibold select-none transition-all duration-150 ease-out focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2 disabled:opacity-50 disabled:cursor-not-allowed disabled:pointer-events-none active:scale-[0.98]';

  const sizeClasses: Record<typeof size, string> = {
    xs: 'px-2.5 py-1 text-xs gap-1.5 rounded-lg min-h-[30px]',
    sm: 'px-3.5 py-1.5 text-xs gap-1.5 rounded-xl min-h-[36px]',
    md: 'px-4.5 py-2.5 text-sm gap-2 rounded-2xl min-h-[44px]',
    lg: 'px-6 py-3.5 text-base gap-2.5 rounded-2xl min-h-[48px]'
  };

  const variantClasses: Record<typeof variant, string> = {
    primary:
      'bg-gradient-to-r from-[#246bfd] to-[#1d4ed8] text-white shadow-md shadow-blue-500/25 border border-white/25 hover:brightness-105 active:brightness-95 ds-specular-top',
    secondary:
      'bg-white/85 text-slate-800 border border-slate-200/90 shadow-sm hover:bg-white hover:border-slate-300 hover:shadow backdrop-blur-md',
    glass:
      'bg-white/70 text-slate-800 border border-white/80 shadow-sm hover:bg-white/90 hover:border-white backdrop-blur-md ds-specular-top',
    ghost:
      'text-slate-600 hover:text-slate-900 hover:bg-slate-100/80 active:bg-slate-200/60 border border-transparent',
    danger:
      'bg-rose-50 text-rose-700 border border-rose-200 hover:bg-rose-100 hover:border-rose-300 active:bg-rose-200/80',
    gold:
      'bg-gradient-to-r from-[#f08300] to-[#d97706] text-white shadow-md shadow-amber-500/25 border border-white/25 hover:brightness-105 active:brightness-95 ds-specular-top',
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
