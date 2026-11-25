<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let ariaLabel: string;
  export let variant: 'primary' | 'secondary' | 'glass' | 'ghost' | 'danger' = 'ghost';
  export let size: 'xs' | 'sm' | 'md' | 'lg' = 'md';
  export let type: 'button' | 'submit' = 'button';
  export let disabled: boolean = false;
  export let loading: boolean = false;
  export let title: string | undefined = undefined;
  export let className: string = '';

  const dispatch = createEventDispatcher<{ click: MouseEvent }>();

  const sizeClasses: Record<typeof size, string> = {
    xs: 'w-7 h-7 rounded-lg p-1',
    sm: 'w-8 h-8 rounded-xl p-1.5',
    md: 'w-10 h-10 rounded-xl p-2',
    lg: 'w-11 h-11 rounded-2xl p-2.5'
  };

  const variantClasses: Record<typeof variant, string> = {
    primary:
      'bg-gradient-to-r from-[#f08300] to-[#d97400] text-white shadow-sm shadow-amber-500/25 border border-amber-400/30 hover:brightness-105 active:brightness-95',
    secondary:
      'bg-white/85 text-slate-700 border border-slate-200/90 shadow-sm hover:bg-white hover:text-slate-900 hover:border-slate-300',
    glass:
      'bg-white/60 text-slate-700 border border-white/80 shadow-sm hover:bg-white/90 hover:text-slate-900 backdrop-blur-md',
    ghost:
      'text-slate-500 hover:text-slate-900 hover:bg-slate-100/80 active:bg-slate-200/60 border border-transparent',
    danger:
      'bg-rose-50 text-rose-600 border border-rose-200/80 hover:bg-rose-100 hover:text-rose-700 hover:border-rose-300'
  };

  function handleClick(e: MouseEvent) {
    if (disabled || loading) {
      e.preventDefault();
      return;
    }
    dispatch('click', e);
  }
</script>

<button
  {type}
  aria-label={ariaLabel}
  title={title || ariaLabel}
  disabled={disabled || loading}
  class="inline-flex items-center justify-center shrink-0 transition-all duration-150 ease-out focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2 disabled:opacity-40 disabled:cursor-not-allowed active:scale-95 {sizeClasses[size]} {variantClasses[variant]} {className}"
  on:click={handleClick}
  {...$$restProps}
>
  {#if loading}
    <svg class="animate-spin w-4 h-4" fill="none" viewBox="0 0 24 24" aria-hidden="true">
      <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="3" />
      <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
    </svg>
  {:else}
    <slot />
  {/if}
</button>
