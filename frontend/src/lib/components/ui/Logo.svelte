<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let variant: 'white' | 'black' = 'white';
  export let size: 'sm' | 'md' | 'lg' | 'xl' = 'md';
  export let href: string | undefined = undefined;
  export let className: string = '';
  export let alt: string = 'Logo Satpol PP Prov. Sulsel';

  const dispatch = createEventDispatcher<{ click: MouseEvent }>();

  const sizeClasses: Record<typeof size, string> = {
    sm: 'h-8 sm:h-9',
    md: 'h-10 sm:h-11',
    lg: 'h-12 sm:h-14',
    xl: 'h-16 sm:h-20'
  };

  const srcMap: Record<typeof variant, string> = {
    white: '/uploads/foto_website/1686928245logo_white.png',
    black: '/uploads/foto_website/1686928245logo_black.png'
  };

  function handleClick(e: MouseEvent) {
    dispatch('click', e);
  }
</script>

{#if href}
  <a
    {href}
    class="inline-flex items-center group py-0.5 select-none ds-press shrink-0 {className}"
    title={alt}
    on:click={handleClick}
    {...$$restProps}
  >
    <img
      src={srcMap[variant]}
      {alt}
      class="{sizeClasses[size]} w-auto object-contain transition-transform duration-200 group-hover:scale-[1.02] filter drop-shadow-xs"
      loading="eager"
    />
  </a>
{:else}
  <img
    src={srcMap[variant]}
    {alt}
    class="{sizeClasses[size]} w-auto object-contain select-none shrink-0 {className}"
    loading="eager"
  />
{/if}
