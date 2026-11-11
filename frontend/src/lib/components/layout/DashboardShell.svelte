<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { goto } from '$app/navigation';
  import { PanelLeftClose, PanelLeftOpen, Menu, X, LogOut, ExternalLink } from 'lucide-svelte';
  import { clearSession } from '$lib/utils/auth';

  export interface DashboardNavItem {
    id: string;
    label: string;
    icon: any;
    badge?: number | string;
    section?: string;
  }

  export let portalLabel = 'Panel Admin';
  export let portalTag = 'SIAP TERTIB';
  export let navItems: DashboardNavItem[] = [];
  export let activeTab = '';
  export let pageTitle = '';
  export let pageDescription = '';
  export let user: { name?: string; username?: string; role?: string; bidang?: string } | null = null;
  export let collapsed = false;
  export let quickSwitchHref = '';
  export let quickSwitchLabel = '';

  const dispatch = createEventDispatcher<{
    select: string;
    logout: void;
  }>();

  let mobileDialog: HTMLDialogElement;
  let mobileOpen = false;

  $: if (mobileDialog) {
    if (mobileOpen && !mobileDialog.open) {
      mobileDialog.showModal();
    } else if (!mobileOpen && mobileDialog.open) {
      mobileDialog.close();
    }
  }

  function selectTab(id: string) {
    activeTab = id;
    mobileOpen = false;
    dispatch('select', id);
  }

  function handleLogout() {
    clearSession();
    dispatch('logout');
    goto('/admin/login');
  }

  function handleDialogBackdropClick(e: MouseEvent) {
    if (!mobileDialog || e.target !== mobileDialog) return;
    const rect = mobileDialog.getBoundingClientRect();
    const inside =
      e.clientX >= rect.left &&
      e.clientX <= rect.right &&
      e.clientY >= rect.top &&
      e.clientY <= rect.bottom;
    if (!inside) mobileOpen = false;
  }

  function userInitials(name?: string, username?: string) {
    const src = (name || username || 'AD').trim();
    return src.slice(0, 2).toUpperCase();
  }
</script>

<div class="min-h-screen flex bg-[#f9f4e8] text-[#1a120c]">
  <!-- Desktop Sidebar (Sticky, Non-Scrolling with Page, Collapsible) -->
  <aside
    aria-label="Navigasi Dashboard"
    class="hidden lg:flex lg:sticky lg:top-0 lg:h-screen shrink-0 bg-[#1a120c] text-[#f9f4e8] border-r border-[#3d261a] flex-col transition-all duration-200 select-none z-30 {collapsed ? 'w-20' : 'w-72'}"
  >
    <!-- Sidebar Brand & Toggle -->
    <div class="h-16 px-4 border-b border-[#3d261a] flex items-center {collapsed ? 'justify-center' : 'justify-between'} gap-2 shrink-0 bg-[#120b07]">
      {#if !collapsed}
        <a href="/" class="flex items-center gap-3 min-w-0 group" title="Ke Beranda Publik">
          <img
            src="/uploads/foto_website/1686928245logo_white.png"
            alt="Logo Satpol PP Sulsel"
            class="h-10 w-auto max-w-[140px] object-contain shrink-0 rounded-xl bg-white/10 p-1 group-hover:bg-white/15 transition"
          />
          <div class="min-w-0">
            <div class="text-xs font-black tracking-wide text-white truncate">{portalLabel}</div>
            <div class="text-[10px] font-bold uppercase tracking-widest text-[#F08300] truncate">{portalTag}</div>
          </div>
        </a>
      {:else}
        <a href="/" class="flex items-center justify-center" title="{portalLabel} — Ke Beranda">
          <img
            src="/uploads/foto_website/1686928245logo_white.png"
            alt="Logo Satpol PP Sulsel"
            class="w-10 h-10 object-contain rounded-xl bg-white/10 p-1 hover:bg-white/15 transition"
          />
        </a>
      {/if}

      {#if !collapsed}
        <button
          type="button"
          on:click={() => (collapsed = true)}
          aria-label="Tutup sidebar"
          title="Kecilkan sidebar"
          class="p-2 rounded-xl text-[#d6cbb8] hover:text-white hover:bg-[#241710] transition shrink-0"
        >
          <PanelLeftClose class="w-4 h-4" />
        </button>
      {/if}
    </div>

    {#if collapsed}
      <div class="px-3 pt-3 flex justify-center shrink-0">
        <button
          type="button"
          on:click={() => (collapsed = false)}
          aria-label="Buka sidebar"
          title="Perluas sidebar"
          class="w-full py-2 flex items-center justify-center rounded-xl text-[#d6cbb8] hover:text-white hover:bg-[#241710] border border-[#3d261a] transition"
        >
          <PanelLeftOpen class="w-4 h-4" />
        </button>
      </div>
    {/if}

    <!-- Navigation Links -->
    <nav class="flex-1 min-h-0 overflow-y-auto px-3 py-4 space-y-1.5">
      {#each navItems as item, idx}
        {#if item.section && !collapsed && (idx === 0 || navItems[idx - 1].section !== item.section)}
          <div class="px-3 pt-3 pb-1 text-[10px] font-bold uppercase tracking-widest text-[#d6cbb8]/60">
            {item.section}
          </div>
        {/if}

        <button
          type="button"
          on:click={() => selectTab(item.id)}
          title={collapsed ? item.label : undefined}
          aria-current={activeTab === item.id ? 'page' : undefined}
          class="relative w-full flex items-center {collapsed ? 'justify-center px-2' : 'justify-between px-3.5'} py-2.5 rounded-xl text-xs font-bold transition {activeTab === item.id ? 'bg-[#F08300] text-white shadow-sm' : 'text-[#d6cbb8] hover:bg-[#241710] hover:text-white'}"
        >
          <span class="flex items-center gap-3 min-w-0">
            <svelte:component this={item.icon} class="w-4 h-4 shrink-0" />
            {#if !collapsed}
              <span class="truncate">{item.label}</span>
            {/if}
          </span>

          {#if !collapsed && item.badge !== undefined && item.badge !== null && item.badge !== '' && item.badge !== 0}
            <span class="ml-2 px-2 py-0.5 text-[10px] font-black rounded-lg shrink-0 {activeTab === item.id ? 'bg-[#1a120c]/25 text-white' : 'bg-[#F08300]/20 text-[#F08300]'}">
              {item.badge}
            </span>
          {:else if collapsed && item.badge !== undefined && item.badge !== 0 && item.badge !== ''}
            <span class="absolute top-1.5 right-1.5 w-2 h-2 rounded-full bg-[#F08300]" aria-hidden="true"></span>
          {/if}
        </button>
      {/each}

      {#if quickSwitchHref && quickSwitchLabel}
        <div class="pt-3 mt-3 border-t border-[#3d261a]">
          <a
            href={quickSwitchHref}
            title={collapsed ? quickSwitchLabel : undefined}
            class="w-full flex items-center {collapsed ? 'justify-center px-2' : 'gap-3 px-3.5'} py-2.5 rounded-xl text-xs font-bold text-[#F08300] bg-[#241710] hover:bg-[#3d261a] transition"
          >
            <ExternalLink class="w-4 h-4 shrink-0" />
            {#if !collapsed}
              <span class="truncate">{quickSwitchLabel}</span>
            {/if}
          </a>
        </div>
      {/if}
    </nav>

    <!-- User Footer -->
    <div class="p-3 border-t border-[#3d261a] bg-[#120b07] shrink-0 space-y-2">
      {#if user && !collapsed}
        <div class="flex items-center gap-2.5 px-2 py-1.5 rounded-xl bg-[#241710]/70">
          <div class="w-8 h-8 rounded-xl bg-[#F08300] text-white font-black text-xs flex items-center justify-center shrink-0">
            {userInitials(user.name, user.username)}
          </div>
          <div class="min-w-0">
            <div class="text-xs font-bold text-white truncate">{user.name || user.username}</div>
            <div class="text-[10px] text-[#d6cbb8] truncate">{user.bidang || user.role}</div>
          </div>
        </div>
      {/if}

      <div class="flex items-center {collapsed ? 'flex-col gap-1.5' : 'gap-2'}">
        <a
          href="/"
          title="Lihat Portal Publik"
          class="flex-1 flex items-center justify-center gap-1.5 px-3 py-2 rounded-xl text-[11px] font-bold text-[#d6cbb8] bg-[#241710] hover:bg-[#3d261a] hover:text-white transition"
        >
          <ExternalLink class="w-3.5 h-3.5 shrink-0" />
          {#if !collapsed}<span>Portal</span>{/if}
        </a>
        <button
          type="button"
          on:click={handleLogout}
          title="Keluar Sesi"
          class="flex-1 flex items-center justify-center gap-1.5 px-3 py-2 rounded-xl text-[11px] font-bold text-red-200 bg-red-950/50 hover:bg-red-700 hover:text-white transition"
        >
          <LogOut class="w-3.5 h-3.5 shrink-0" />
          {#if !collapsed}<span>Keluar</span>{/if}
        </button>
      </div>
    </div>
  </aside>

  <!-- Mobile Sidebar Drawer (Native Dialog for Scroll Lock & Esc Handling) -->
  <!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
  <dialog
    id="dashboard-mobile-drawer"
    bind:this={mobileDialog}
    closedby="any"
    on:close={() => (mobileOpen = false)}
    on:click={handleDialogBackdropClick}
    aria-label="Menu navigasi dashboard"
    class="m-0 p-0 h-dvh max-h-dvh w-72 max-w-[85vw] bg-[#1a120c] text-[#f9f4e8] border-r border-[#3d261a] backdrop:bg-[#1a120c]/60 lg:hidden"
  >
    <div class="h-full flex flex-col">
      <div class="h-16 px-4 border-b border-[#3d261a] bg-[#120b07] flex items-center justify-between gap-2 shrink-0">
        <div class="flex items-center gap-3 min-w-0">
          <img src="/uploads/foto_website/1686928245logo_white.png" alt="Logo Satpol PP Sulsel" class="w-9 h-9 object-contain rounded-xl bg-white/10 p-1 shrink-0" />
          <div class="min-w-0">
            <div class="text-xs font-black text-white truncate">{portalLabel}</div>
            <div class="text-[10px] font-bold uppercase tracking-widest text-[#F08300] truncate">{portalTag}</div>
          </div>
        </div>
        <!-- svelte-ignore a11y-autofocus -->
        <button
          type="button"
          autofocus
          on:click={() => (mobileOpen = false)}
          aria-label="Tutup menu"
          class="p-2 rounded-xl text-[#d6cbb8] hover:text-white hover:bg-[#241710]"
        >
          <X class="w-4 h-4" />
        </button>
      </div>

      <nav class="flex-1 overflow-y-auto px-3 py-4 space-y-1.5">
        {#each navItems as item}
          <button
            type="button"
            on:click={() => selectTab(item.id)}
            class="w-full flex items-center justify-between px-3.5 py-2.5 rounded-xl text-xs font-bold transition {activeTab === item.id ? 'bg-[#F08300] text-white' : 'text-[#d6cbb8] hover:bg-[#241710] hover:text-white'}"
          >
            <span class="flex items-center gap-3 min-w-0">
              <svelte:component this={item.icon} class="w-4 h-4 shrink-0" />
              <span class="truncate">{item.label}</span>
            </span>
            {#if item.badge !== undefined && item.badge !== 0 && item.badge !== ''}
              <span class="px-2 py-0.5 text-[10px] font-black rounded-lg bg-[#241710] text-[#F08300]">{item.badge}</span>
            {/if}
          </button>
        {/each}
        {#if quickSwitchHref && quickSwitchLabel}
          <a
            href={quickSwitchHref}
            class="mt-3 w-full flex items-center gap-3 px-3.5 py-2.5 rounded-xl text-xs font-bold text-[#F08300] bg-[#241710] hover:bg-[#3d261a]"
          >
            <ExternalLink class="w-4 h-4 shrink-0" />
            <span>{quickSwitchLabel}</span>
          </a>
        {/if}
      </nav>

      <div class="p-3 border-t border-[#3d261a] bg-[#120b07] space-y-2 shrink-0">
        {#if user}
          <div class="px-3 py-2 rounded-xl bg-[#241710] text-xs">
            <div class="font-bold text-white truncate">{user.name || user.username}</div>
            <div class="text-[11px] text-[#d6cbb8] truncate">{user.bidang || user.role}</div>
          </div>
        {/if}
        <button
          type="button"
          on:click={handleLogout}
          class="w-full flex items-center justify-center gap-2 px-3 py-2.5 rounded-xl text-xs font-bold text-red-200 bg-red-950/60 hover:bg-red-700 hover:text-white transition"
        >
          <LogOut class="w-4 h-4" />
          Keluar Sesi
        </button>
      </div>
    </div>
  </dialog>

  <!-- Main Scrollable Content Area -->
  <div class="flex-1 min-w-0 flex flex-col">
    <!-- Sticky Top Header Bar -->
    <header class="sticky top-0 z-20 bg-white/95 backdrop-blur border-b border-[#eee9db] px-4 md:px-8 h-16 flex items-center justify-between gap-4 shrink-0">
      <div class="flex items-center gap-3 min-w-0">
        <button
          type="button"
          on:click={() => (mobileOpen = true)}
          aria-label="Buka menu navigasi"
          aria-expanded={mobileOpen}
          aria-controls="dashboard-mobile-drawer"
          class="lg:hidden p-2 rounded-xl border border-[#eee9db] bg-[#f9f4e8] text-[#1a120c] hover:bg-[#eee9db] transition shrink-0"
        >
          <Menu class="w-5 h-5" />
        </button>

        <button
          type="button"
          on:click={() => (collapsed = !collapsed)}
          aria-label={collapsed ? 'Perluas sidebar' : 'Kecilkan sidebar'}
          title={collapsed ? 'Perluas sidebar' : 'Kecilkan sidebar'}
          class="hidden lg:inline-flex p-2 rounded-xl border border-[#eee9db] bg-[#f9f4e8] text-[#3d261a] hover:bg-[#eee9db] transition shrink-0"
        >
          {#if collapsed}
            <PanelLeftOpen class="w-4 h-4" />
          {:else}
            <PanelLeftClose class="w-4 h-4" />
          {/if}
        </button>

        <div class="min-w-0">
          <h1 class="text-sm md:text-lg font-black text-[#1a120c] truncate">{pageTitle}</h1>
          {#if pageDescription}
            <p class="hidden sm:block text-[11px] text-[#3d261a]/75 truncate">{pageDescription}</p>
          {/if}
        </div>
      </div>

      <div class="flex items-center gap-2.5 shrink-0">
        <slot name="header-actions" />
        {#if user}
          <div class="hidden md:flex items-center gap-2 px-3 py-1.5 rounded-xl bg-[#f9f4e8] border border-[#eee9db]">
            <span class="w-2 h-2 rounded-full bg-emerald-600" aria-hidden="true"></span>
            <span class="text-xs font-bold text-[#1a120c]">{user.username}</span>
          </div>
        {/if}
      </div>
    </header>

    <!-- Page Body -->
    <main class="flex-1 p-4 md:p-8 max-w-7xl w-full mx-auto space-y-6">
      <slot />
    </main>
  </div>
</div>
