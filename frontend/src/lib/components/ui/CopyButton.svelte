<script lang="ts">
  import { onDestroy } from 'svelte';

  export let text: string;
  export let label: string | undefined = undefined;
  export let copiedLabel: string = 'Tersalin!';
  export let size: 'xs' | 'sm' | 'md' = 'sm';
  export let variant: 'primary' | 'glass' | 'ghost' | 'soft' = 'glass';
  export let className: string = '';
  export let ariaLabel: string = 'Salin ke papan klip';

  let copied = false;
  let copyTimeout: ReturnType<typeof setTimeout> | null = null;

  async function handleCopy() {
    if (!text || typeof window === 'undefined') return;
    try {
      if (navigator?.clipboard?.writeText) {
        await navigator.clipboard.writeText(text);
      } else {
        const textArea = document.createElement('textarea');
        textArea.value = text;
        textArea.style.position = 'fixed';
        textArea.style.opacity = '0';
        document.body.appendChild(textArea);
        textArea.focus();
        textArea.select();
        document.execCommand('copy');
        document.body.removeChild(textArea);
      }
      copied = true;
      if (copyTimeout) clearTimeout(copyTimeout);
      copyTimeout = setTimeout(() => {
        copied = false;
      }, 2000);
    } catch (err) {
      console.error('Failed to copy to clipboard', err);
    }
  }

  onDestroy(() => {
    if (copyTimeout) clearTimeout(copyTimeout);
  });

  const sizeClasses: Record<typeof size, string> = {
    xs: 'px-2 py-1 text-[11px] gap-1 rounded-lg',
    sm: 'px-2.5 py-1.5 text-xs gap-1.5 rounded-xl',
    md: 'px-3.5 py-2 text-sm gap-2 rounded-xl'
  };

  const variantClasses: Record<typeof variant, string> = {
    primary:
      'bg-gradient-to-r from-[#246bfd] to-[#1d4ed8] text-white shadow-sm border border-white/20 hover:brightness-105',
    glass:
      'bg-white/80 text-slate-700 border border-slate-200/90 shadow-sm hover:bg-white hover:text-slate-900 backdrop-blur-md',
    ghost:
      'text-slate-500 hover:text-slate-900 hover:bg-slate-100/80 border border-transparent',
    soft:
      'bg-blue-50/80 text-blue-700 border border-blue-200/60 hover:bg-blue-100'
  };
</script>

<button
  type="button"
  aria-label={copied ? copiedLabel : ariaLabel}
  title={copied ? copiedLabel : ariaLabel}
  on:click={handleCopy}
  class="inline-flex items-center justify-center font-semibold transition-all duration-150 ease-out focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 focus-visible:ring-offset-2 active:scale-95 {sizeClasses[size]} {copied ? 'bg-emerald-50 text-emerald-700 border border-emerald-300' : variantClasses[variant]} {className}"
>
  {#if copied}
    <!-- Checkmark Icon -->
    <svg class="w-3.5 h-3.5 text-emerald-600 transition-transform duration-200 scale-110" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 13l4 4L19 7" />
    </svg>
    <span>{copiedLabel}</span>
  {:else}
    <!-- Copy Icon -->
    <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
    </svg>
    {#if label}
      <span>{label}</span>
    {/if}
  {/if}
</button>
