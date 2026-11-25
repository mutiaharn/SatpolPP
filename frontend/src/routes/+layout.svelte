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
    <!-- Azure Sovereign Ambient Orb (Top-Left) -->
    <div class="absolute -top-24 -left-24 w-[600px] h-[600px] rounded-full bg-blue-500/18 blur-[110px] transform-gpu"></div>
    <!-- Praja Gold/Amber Ambient Orb (Top-Right) -->
    <div class="absolute top-1/4 -right-24 w-[650px] h-[650px] rounded-full bg-amber-500/16 blur-[120px] transform-gpu"></div>
    <!-- Sky Cyan Ambient Orb (Mid-Center) -->
    <div class="absolute top-2/3 left-1/3 -translate-x-1/2 w-[700px] h-[700px] rounded-full bg-sky-400/14 blur-[130px] transform-gpu"></div>
    <!-- Royal Indigo Ambient Orb (Bottom-Left) -->
    <div class="absolute -bottom-40 -left-20 w-[650px] h-[650px] rounded-full bg-indigo-500/12 blur-[120px] transform-gpu"></div>
    <!-- Soft Amber Ambient Orb (Bottom-Right) -->
    <div class="absolute -bottom-32 right-10 w-[550px] h-[550px] rounded-full bg-orange-400/15 blur-[110px] transform-gpu"></div>
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
