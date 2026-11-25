<script lang="ts">
  import { Search } from 'lucide-svelte';
  import CommandPalette from '$lib/components/ui/CommandPalette.svelte';
  import { Button } from '$lib/components/ui';

  let mobileMenuOpen = false;
  let searchOpen = false;
</script>

<header class="w-full sticky top-0 z-50 bg-white/85 backdrop-blur-2xl border-b border-white/80 shadow-[0_8px_32px_rgba(15,23,42,0.06)] ds-specular-top transition-all duration-300">
  <!-- Top Bar: Official Contacts (Unified Seamless Glass Sub-row) -->
  <div class="border-b border-slate-200/50 text-xs py-1.5 px-4 hidden sm:block">
    <div class="max-w-7xl mx-auto flex flex-wrap justify-between items-center gap-2">
      <div class="flex items-center gap-4 sm:gap-6 text-[11px] sm:text-xs">
        <a href="tel:0411-453105" class="hover:text-[#F08300] transition flex items-center gap-1.5 text-slate-600 font-medium">
          <svg class="w-3.5 h-3.5 text-[#F08300]" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" /></svg>
          <span>Telepon: (0411) 453105</span>
        </a>
        <a href="mailto:satpol@sulselprov.go.id" class="hidden sm:inline hover:text-[#F08300] transition text-slate-600">Email: satpol@sulselprov.go.id</a>
      </div>
      <div class="flex items-center gap-2 text-[11px] text-slate-500">
        <span class="font-semibold text-slate-600">Pemerintah Provinsi Sulawesi Selatan</span>
      </div>
    </div>
  </div>

  <!-- Main Navigation Bar -->
  <div>
    <div class="max-w-7xl mx-auto px-4 py-2.5 flex items-center justify-between gap-4">
      <!-- Official Logo -->
      <a href="/" class="flex items-center group py-0.5 ds-press" title="Beranda Satpol PP Sulsel">
        <img 
          src="/uploads/foto_website/1686928245logo_black.png" 
          alt="Logo Satpol PP Prov. Sulsel" 
          class="h-10 sm:h-11 w-auto object-contain transition duration-200 group-hover:scale-[1.02]" 
        />
      </a>

      <!-- Desktop Navigation Menu Links (Minimalist Apple HIG Layout) -->
      <nav class="hidden lg:flex items-center gap-1.5 xl:gap-2.5 text-[12px] xl:text-[13px] font-bold text-slate-700 uppercase tracking-wider shrink-0">
        <a href="/" class="px-4 py-2 rounded-xl min-h-[38px] flex items-center hover:bg-slate-900/5 hover:text-[#F08300] transition duration-200 ds-press">Beranda</a>
        <a href="/lacak" class="px-4 py-2 rounded-xl min-h-[38px] flex items-center hover:bg-slate-900/5 hover:text-[#F08300] transition duration-200 whitespace-nowrap ds-press">Lacak Aduan</a>
        <a href="/kontak" class="px-4 py-2 rounded-xl min-h-[38px] flex items-center hover:bg-slate-900/5 hover:text-[#F08300] transition duration-200 ds-press">Kontak</a>
      </nav>

      <!-- Search Engine Trigger & CTA Button -->
      <div class="hidden lg:flex items-center gap-2.5 shrink-0">
        <!-- Search Command Trigger (Apple Frosted Glass Capsule) -->
        <button
          type="button"
          on:click={() => searchOpen = true}
          class="inline-flex items-center gap-2 px-3.5 py-2 text-xs text-slate-600 bg-white/70 hover:bg-white/95 rounded-xl border border-white/80 min-h-[40px] transition duration-200 shadow-xs hover:shadow-sm focus:outline-none focus:ring-2 focus:ring-[#F08300]/40 shrink-0 ds-press backdrop-blur-md"
          title="Buka Pencarian Cepat (⌘K / Ctrl+K)"
        >
          <Search class="w-3.5 h-3.5 text-slate-400" />
          <span class="hidden xl:inline">Cari...</span>
          <kbd class="bg-white/90 px-1.5 py-0.5 rounded text-[10px] font-mono text-slate-500 border border-slate-200 shadow-2xs">⌘K</kbd>
        </button>

        <!-- Call-to-Action: SIAP TERTIB (Luminous Amber Glass Button) -->
        <Button
          href="/lapor"
          variant="gold"
          size="sm"
          className="uppercase tracking-wider px-4 xl:px-5 py-2.5 min-h-[40px] font-bold"
        >
          <span>Lapor SIAP TERTIB</span>
        </Button>
      </div>

      <!-- Mobile Search & Menu Controls -->
      <div class="flex items-center gap-1.5 lg:hidden">
        <button
          type="button"
          on:click={() => searchOpen = true}
          class="p-2.5 min-w-[44px] min-h-[44px] flex items-center justify-center rounded-xl text-slate-700 hover:text-[#F08300] hover:bg-white/60 transition focus:outline-none ds-press"
          aria-label="Cari"
        >
          <Search class="w-5 h-5" />
        </button>

        <button 
          type="button" 
          class="p-2.5 min-w-[44px] min-h-[44px] flex items-center justify-center rounded-xl text-slate-800 hover:text-[#F08300] hover:bg-white/60 transition focus:outline-none ds-press"
          on:click={() => mobileMenuOpen = !mobileMenuOpen}
          aria-label="Buka Menu"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            {#if mobileMenuOpen}
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            {:else}
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
            {/if}
          </svg>
        </button>
      </div>
    </div>
  </div>

  <!-- Fullscreen Mobile Frosted Drawer -->
  {#if mobileMenuOpen}
    <div 
      class="fixed inset-0 z-[100] bg-[#070e1c]/88 backdrop-blur-2xl text-white flex flex-col justify-between overflow-y-auto lg:hidden animate-in fade-in duration-200"
    >
      <!-- Mobile Bar Top with Close Button -->
      <div class="flex items-center justify-between px-6 py-4 border-b border-white/10 bg-white/5 backdrop-blur-xl">
        <a href="/" on:click={() => mobileMenuOpen = false} class="ds-press">
          <img 
            src="/uploads/foto_website/1686928245logo_white.png" 
            alt="Logo Satpol PP" 
            class="h-10 w-auto object-contain" 
          />
        </a>
        <button 
          type="button" 
          class="p-2 rounded-xl bg-white/10 hover:bg-white/20 text-white transition focus:outline-none ds-press"
          on:click={() => mobileMenuOpen = false}
          aria-label="Tutup Menu"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Mobile Navigation Links (Minimalist Clean Drawer) -->
      <div class="px-6 py-6 space-y-4 flex-1 overflow-y-auto">
        <a 
          href="/" 
          class="block text-lg font-bold uppercase tracking-wider text-white hover:text-[#F08300] py-2 transition border-b border-white/10"
          on:click={() => mobileMenuOpen = false}
        >
          Beranda
        </a>
        <a 
          href="/lacak" 
          class="block text-lg font-bold uppercase tracking-wider text-white hover:text-[#F08300] py-2 transition border-b border-white/10"
          on:click={() => mobileMenuOpen = false}
        >
          Lacak Tiket Aduan
        </a>
        <a 
          href="/kontak" 
          class="block text-lg font-bold uppercase tracking-wider text-white hover:text-[#F08300] py-2 transition border-b border-white/10"
          on:click={() => mobileMenuOpen = false}
        >
          Kontak Resmi
        </a>

        <!-- Profil & Bidang Quick Access (Symmetric Frosted Cards) -->
        <div class="pt-2 grid grid-cols-2 gap-2.5 text-xs">
          <div class="ds-glass-dark border border-white/10 rounded-xl p-3 space-y-1.5">
            <span class="text-[#F08300] font-bold uppercase tracking-wider text-[11px] block">Profil</span>
            <div class="space-y-1 text-slate-300">
              <a href="/profil/sejarah" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Sejarah</a>
              <a href="/profil/visi-misi" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Visi &amp; Misi</a>
              <a href="/profil/struktur" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Struktur</a>
              <a href="/profil/tupoksi" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Tupoksi</a>
              <a href="/profil/pejabat" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Pejabat</a>
            </div>
          </div>
          <div class="ds-glass-dark border border-white/10 rounded-xl p-3 space-y-1.5">
            <span class="text-[#F08300] font-bold uppercase tracking-wider text-[11px] block">Bidang</span>
            <div class="space-y-1 text-slate-300">
              <a href="/bidang/sekretariat" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Sekretariat</a>
              <a href="/bidang/perda" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Perda</a>
              <a href="/bidang/trantibum" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Trantibum</a>
              <a href="/bidang/binmas" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Binmas</a>
              <a href="/bidang/linmas-damkar" class="block hover:text-white transition py-0.5" on:click={() => mobileMenuOpen = false}>Linmas</a>
            </div>
          </div>
        </div>
      </div>

      <!-- Mobile Footer Action -->
      <div class="p-6 border-t border-white/10 bg-black/40 backdrop-blur-xl">
        <Button
          href="/lapor"
          variant="gold"
          size="lg"
          fullWidth
          className="uppercase tracking-wider font-bold"
          on:click={() => mobileMenuOpen = false}
        >
          <span>Lapor SIAP TERTIB</span>
        </Button>
      </div>
    </div>
  {/if}
</header>

<CommandPalette bind:open={searchOpen} />
