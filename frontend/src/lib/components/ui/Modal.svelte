<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let open: boolean = false;
  export let title: string = '';
  export let subtitle: string = '';
  export let size: 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' = 'lg';
  export let showCloseButton: boolean = true;
  export let closeOnBackdropClick: boolean = true;
  export let closeOnEscape: boolean = true;

  const dispatch = createEventDispatcher<{
    close: void;
  }>();

  function handleClose() {
    open = false;
    dispatch('close');
  }

  function handleBackdropClick() {
    if (closeOnBackdropClick) {
      handleClose();
    }
  }

  function handleKeyDown(e: KeyboardEvent) {
    if (open && closeOnEscape && e.key === 'Escape') {
      handleClose();
    }
  }

  const sizeClasses = {
    sm: 'max-w-md',
    md: 'max-w-lg',
    lg: 'max-w-2xl',
    xl: 'max-w-4xl',
    '2xl': 'max-w-5xl',
    full: 'max-w-6xl'
  };
</script>

<svelte:window on:keydown={handleKeyDown} />

{#if open}
  <!-- Backdrop container -->
  <div
    class="fixed inset-0 z-50 overflow-y-auto bg-black/60 backdrop-blur-xs flex items-center justify-center p-3 sm:p-5 transition-opacity"
    role="dialog"
    aria-modal="true"
    aria-labelledby="modal-title"
  >
    <!-- Invisible Backdrop click overlay button -->
    <button
      type="button"
      class="fixed inset-0 w-full h-full bg-transparent cursor-default border-0 p-0 focus:outline-none"
      on:click={handleBackdropClick}
      aria-label="Tutup modal latar"
      tabindex="-1"
    ></button>

    <!-- Modal Card (in front of backdrop) -->
    <div
      class={`relative z-10 bg-white rounded-xl shadow-2xl border border-stone-200 w-full ${sizeClasses[size]} overflow-hidden transform transition-all duration-200 flex flex-col max-h-[92vh]`}
    >
      <!-- Header -->
      <div class="px-5 py-4 border-b border-[#eee9db] bg-[#faf7f2] flex items-center justify-between shrink-0">
        <div>
          {#if subtitle}
            <span class="text-[10px] font-bold uppercase tracking-wider text-[#F08300] block">
              {subtitle}
            </span>
          {/if}
          {#if title}
            <h3 id="modal-title" class="text-sm sm:text-base font-bold text-[#1a120c] leading-tight">
              {title}
            </h3>
          {/if}
          <slot name="header" />
        </div>

        {#if showCloseButton}
          <button
            type="button"
            on:click={handleClose}
            class="p-1.5 rounded-lg text-stone-400 hover:text-stone-700 hover:bg-stone-200/60 transition"
            aria-label="Tutup modal"
          >
            <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        {/if}
      </div>

      <!-- Body / Content -->
      <div class="p-5 sm:p-6 overflow-y-auto flex-1 text-xs text-[#2b2f3f] leading-relaxed">
        <slot />
      </div>

      <!-- Footer (if provided) -->
      {#if $$slots.footer}
        <div class="px-5 py-3.5 border-t border-[#eee9db] bg-[#faf7f2]/70 flex flex-wrap items-center justify-end gap-2 shrink-0">
          <slot name="footer" {handleClose} />
        </div>
      {/if}
    </div>
  </div>
{/if}
