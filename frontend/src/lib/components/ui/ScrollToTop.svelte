<script lang="ts">
  import { onMount } from 'svelte';
  import { ChevronUp } from 'lucide-svelte';

  export let threshold = 280;
  let visible = false;

  function checkScroll() {
    if (typeof window !== 'undefined') {
      visible = window.scrollY > threshold;
    }
  }

  function scrollToTop() {
    if (typeof window !== 'undefined') {
      window.scrollTo({
        top: 0,
        behavior: 'smooth'
      });
    }
  }

  onMount(() => {
    checkScroll();
    window.addEventListener('scroll', checkScroll, { passive: true });
    return () => {
      window.removeEventListener('scroll', checkScroll);
    };
  });
</script>

<button
  type="button"
  on:click={scrollToTop}
  aria-label="Kembali ke atas halaman"
  title="Kembali ke atas halaman"
  class="fixed bottom-6 right-6 z-40 w-12 h-12 rounded-2xl flex items-center justify-center transition-all duration-300 ease-out select-none focus:outline-none focus-visible:ring-2 focus-visible:ring-[#F08300] focus-visible:ring-offset-2 ds-press {visible ? 'opacity-100 scale-100 pointer-events-auto' : 'opacity-0 scale-75 pointer-events-none'} bg-white/85 hover:bg-white text-slate-700 hover:text-[#F08300] border border-white/90 backdrop-blur-2xl shadow-[0_10px_30px_-4px_rgba(15,23,42,0.18),inset_0_1px_0_rgba(255,255,255,0.60)]"
>
  <ChevronUp class="w-5 h-5 transition-transform duration-200" />
</button>
