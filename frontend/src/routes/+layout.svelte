<script lang="ts">
  import '../app.css';
  import { page } from '$app/stores';
  import Navbar from '$lib/components/Navbar.svelte';
  import Footer from '$lib/components/Footer.svelte';

  // Sembunyikan Navbar dan Footer publik saat berada di halaman admin/superadmin
  $: isAdminRoute = $page.url.pathname.startsWith('/admin') || $page.url.pathname.startsWith('/superadmin');
</script>

<div class="min-h-screen flex flex-col bg-[#f8fafc] text-[#0f172a] selection:bg-[#F08300] selection:text-white ds-ambient-glow">
  {#if !isAdminRoute}
    <Navbar />
  {/if}

  <main class="flex-1">
    <slot />
  </main>

  {#if !isAdminRoute}
    <Footer />
  {/if}
</div>
