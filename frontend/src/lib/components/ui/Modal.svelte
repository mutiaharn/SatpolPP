<script lang="ts">
  import { createEventDispatcher, onDestroy } from 'svelte';

  export let open: boolean = false;
  export let title: string = '';
  export let subtitle: string = '';
  export let size: 'sm' | 'md' | 'lg' | 'xl' | '2xl' | 'full' = 'lg';
  export let showCloseButton: boolean = true;
  export let closeOnBackdropClick: boolean = true;

  const dispatch = createEventDispatcher<{ close: void }>();

  let dialogEl: HTMLDialogElement;
  let pressedOutside = false;

  const titleId = `modal-title-${Math.random().toString(36).slice(2, 9)}`;

  const sizeClasses: Record<typeof size, string> = {
    sm: 'max-w-md',
    md: 'max-w-lg',
    lg: 'max-w-2xl',
    xl: 'max-w-4xl',
    '2xl': 'max-w-5xl',
    full: 'max-w-6xl'
  };

  $: if (dialogEl) {
    if (open && !dialogEl.open) {
      dialogEl.showModal();
    } else if (!open && dialogEl.open) {
      dialogEl.close();
    }
  }

  function handleClose() {
    if (dialogEl?.open) dialogEl.close();
    open = false;
    dispatch('close');
  }

  function handleNativeClose() {
    open = false;
    dispatch('close');
  }

  function isOutside(e: PointerEvent): boolean {
    if (!dialogEl) return false;
    const r = dialogEl.getBoundingClientRect();
    return e.clientX < r.left || e.clientX > r.right || e.clientY < r.top || e.clientY > r.bottom;
  }

  function handlePointerDown(e: PointerEvent) {
    pressedOutside = closeOnBackdropClick && isOutside(e);
  }

  function handlePointerUp(e: PointerEvent) {
    if (pressedOutside && closeOnBackdropClick && isOutside(e)) {
      handleClose();
    }
    pressedOutside = false;
  }

  onDestroy(() => {
    if (dialogEl?.open) dialogEl.close();
  });
</script>

<!-- Native <dialog> with Tamalate Specular Glassmorphism -->
<dialog
  bind:this={dialogEl}
  aria-labelledby={title ? titleId : undefined}
  aria-label={!title ? subtitle || 'Dialog' : undefined}
  on:close={handleNativeClose}
  on:pointerdown={handlePointerDown}
  on:pointerup={handlePointerUp}
  class={`m-auto p-0 ds-glass-strong rounded-3xl shadow-2xl border border-white/95 w-[calc(100%-1.5rem)] ${sizeClasses[size]} max-h-[90vh] overflow-hidden open:flex flex-col ds-specular-top`}
>
  <!-- Header (Frosted Glass Header) -->
  <div class="px-6 py-4 border-b border-white/80 bg-white/40 flex items-center justify-between shrink-0 gap-3 backdrop-blur-md">
    <div class="min-w-0">
      {#if subtitle}
        <span class="text-[10px] font-bold uppercase tracking-wider text-[#F08300] block mb-0.5">
          {subtitle}
        </span>
      {/if}
      {#if title}
        <h3 id={titleId} class="text-sm sm:text-base font-bold text-slate-900 leading-tight truncate">
          {title}
        </h3>
      {/if}
      <slot name="header" />
    </div>

    {#if showCloseButton}
      <button
        type="button"
        on:click={handleClose}
        class="p-2 rounded-xl text-slate-400 hover:text-slate-900 hover:bg-slate-900/5 transition shrink-0 ds-press"
        aria-label="Tutup modal"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    {/if}
  </div>

  <!-- Body / Content -->
  <div class="p-6 overflow-y-auto flex-1 text-xs text-slate-700 leading-relaxed font-light">
    <slot />
  </div>

  <!-- Footer (Frosted Glass Footer) -->
  {#if $$slots.footer}
    <div class="px-6 py-4 border-t border-white/80 bg-white/40 flex flex-wrap items-center justify-end gap-2 shrink-0 backdrop-blur-md">
      <slot name="footer" {handleClose} />
    </div>
  {/if}
</dialog>
