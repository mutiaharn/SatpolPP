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

<!-- Native <dialog> opened via showModal(); Esc & top-layer focus trap ditangani browser -->
<dialog
  bind:this={dialogEl}
  aria-labelledby={title ? titleId : undefined}
  aria-label={!title ? subtitle || 'Dialog' : undefined}
  on:close={handleNativeClose}
  on:pointerdown={handlePointerDown}
  on:pointerup={handlePointerUp}
  class={`m-auto p-0 bg-white rounded-2xl shadow-2xl border border-[#eee9db] w-[calc(100%-1.5rem)] ${sizeClasses[size]} max-h-[90vh] overflow-hidden open:flex flex-col`}
>
  <!-- Header -->
  <div class="px-5 py-4 border-b border-[#eee9db] bg-[#faf7f2] flex items-center justify-between shrink-0 gap-3">
    <div class="min-w-0">
      {#if subtitle}
        <span class="text-[10px] font-bold uppercase tracking-wider text-[#F08300] block">
          {subtitle}
        </span>
      {/if}
      {#if title}
        <h3 id={titleId} class="text-sm sm:text-base font-bold text-[#1a120c] leading-tight truncate">
          {title}
        </h3>
      {/if}
      <slot name="header" />
    </div>

    {#if showCloseButton}
      <button
        type="button"
        on:click={handleClose}
        class="p-1.5 rounded-lg text-stone-500 hover:text-[#1a120c] hover:bg-stone-200/60 transition shrink-0"
        aria-label="Tutup modal"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    {/if}
  </div>

  <!-- Body / Content -->
  <div class="p-5 sm:p-6 overflow-y-auto flex-1 text-xs text-stone-700 leading-relaxed">
    <slot />
  </div>

  <!-- Footer -->
  {#if $$slots.footer}
    <div class="px-5 py-3.5 border-t border-[#eee9db] bg-[#faf7f2]/80 flex flex-wrap items-center justify-end gap-2 shrink-0">
      <slot name="footer" {handleClose} />
    </div>
  {/if}
</dialog>
