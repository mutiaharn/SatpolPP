<script lang="ts">
  import '../app.css';
  import { page } from '$app/stores';
  import Navbar from '$lib/components/Navbar.svelte';
  import Footer from '$lib/components/Footer.svelte';
  import { ScrollToTop } from '$lib/components/ui';

  // Sembunyikan Navbar dan Footer publik saat berada di halaman admin/superadmin
  $: isAdminRoute = $page.url.pathname.startsWith('/admin') || $page.url.pathname.startsWith('/superadmin');
  $: isLacakRoute = $page.url.pathname === '/lacak';
  $: isLoginRoute = $page.url.pathname === '/admin/login';
</script>

<div class="min-h-screen flex flex-col bg-gradient-to-b from-[#f8fafc] via-[#f1f5f9]/70 to-[#f8fafc] text-[#0f172a] selection:bg-[#F08300] selection:text-white relative">
  <!-- Ambient Background Gradient Mesh Blobs for Living Glassmorphic Refraction -->
  <div class="pointer-events-none fixed inset-0 -z-10 overflow-hidden select-none" aria-hidden="true">
    <!-- Praja Gold/Amber Ambient Orb (Top-Right) -->
    <div class="absolute top-1/4 -right-24 w-[650px] h-[650px] rounded-full bg-amber-500/8 blur-[120px] transform-gpu"></div>
    <!-- Soft Warm Ambient Orb (Bottom-Left) -->
    <div class="absolute -bottom-40 -left-20 w-[650px] h-[650px] rounded-full bg-amber-500/6 blur-[120px] transform-gpu"></div>
  </div>

  {#if !isAdminRoute}
    <Navbar />
  {/if}

  <main class={`flex-1 ${isLacakRoute ? 'flex flex-col' : ''}`}>
    <slot />
  </main>

  {#if !isAdminRoute}
    <Footer compact={isLacakRoute} />
  {/if}

  {#if !isLoginRoute}
    <ScrollToTop />
  {/if}
</div>
