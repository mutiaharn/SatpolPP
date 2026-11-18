<script lang="ts">
  import { Search } from 'lucide-svelte';
  import CommandPalette from '$lib/components/ui/CommandPalette.svelte';

  let mobileMenuOpen = false;
  let profilOpen = false;
  let bidangOpen = false;
  let searchOpen = false;
</script>

<header class="w-full sticky top-0 z-50 bg-white/95 backdrop-blur-md border-b border-slate-200/80 transition-all shadow-xs">
  <!-- Top Bar: Official Contacts (Soft Porcelain Slate - Senada, Elegan & Tidak Mencolok) -->
  <div class="bg-slate-50/90 text-slate-500 text-xs py-1.5 px-4 border-b border-slate-200/60">
    <div class="max-w-7xl mx-auto flex flex-wrap justify-between items-center gap-2">
      <div class="flex items-center gap-4 sm:gap-6 text-[11px] sm:text-xs">
        <a href="tel:0411-453105" class="hover:text-[#F08300] transition flex items-center gap-1.5 text-slate-600">
          <svg class="w-3.5 h-3.5 text-[#F08300]" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" /></svg>
          <span>Telepon: (0411) 453105</span>
        </a>
        <a href="mailto:satpol@sulselprov.go.id" class="hidden sm:inline hover:text-[#F08300] transition text-slate-600">Email: satpol@sulselprov.go.id</a>
        <a href="https://ppidsatpolppsulsel.my.id/" target="_blank" rel="noreferrer" class="hidden md:inline hover:text-[#F08300] transition text-slate-600">Layanan PPID</a>
      </div>
      <div class="flex items-center gap-2 text-[11px] text-slate-500">
        <span class="inline-block w-2 h-2 rounded-full bg-emerald-500 animate-pulse"></span>
        <span class="font-medium">Pemerintah Provinsi Sulawesi Selatan</span>
      </div>
    </div>
  </div>

  <!-- Main Navigation Bar -->
  <div class="max-w-7xl mx-auto px-4 py-2.5 flex items-center justify-between gap-4">
    <!-- Official Logo -->
    <a href="/" class="flex items-center group py-0.5" title="Beranda Satpol PP Sulsel">
      <img 
        src="/uploads/foto_website/1686928245logo_black.png" 
        alt="Logo Satpol PP Prov. Sulsel" 
        class="h-10 sm:h-11 w-auto object-contain transition duration-200 group-hover:scale-[1.02]" 
      />
    </a>

    <!-- Desktop Navigation Menu Links -->
    <nav class="hidden lg:flex items-center gap-4 xl:gap-6 text-[12px] xl:text-[13px] font-bold text-slate-700 uppercase tracking-wider shrink-0">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>

      <!-- Dropdown Profil -->
      <div 
        class="relative group"
        role="menu"
        tabindex="0"
        on:mouseenter={() => profilOpen = true}
        on:mouseleave={() => profilOpen = false}
      >
        <button class="flex items-center gap-1 hover:text-[#F08300] py-2 transition uppercase focus:outline-none" aria-haspopup="true" aria-expanded={profilOpen}>
          <span>Profil</span>
          <svg class="w-3.5 h-3.5 transition-transform duration-200 group-hover:rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
        </button>
        {#if profilOpen}
          <div class="absolute top-full left-0 w-64 bg-white/95 backdrop-blur-xl border border-slate-200 shadow-xl rounded-2xl py-2 z-50 text-xs normal-case animate-in fade-in zoom-in-95 duration-150">
            <a href="/profil/sejarah" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Sejarah Singkat</a>
            <a href="/profil/visi-misi" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Visi &amp; Misi</a>
            <a href="/profil/struktur" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Struktur Organisasi</a>
            <a href="/profil/tupoksi" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Tugas Pokok &amp; Fungsi</a>
            <a href="/profil/pejabat" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Profil Pejabat</a>
          </div>
        {/if}
      </div>

      <!-- Dropdown Bidang -->
      <div 
        class="relative group"
        role="menu"
        tabindex="0"
        on:mouseenter={() => bidangOpen = true}
        on:mouseleave={() => bidangOpen = false}
      >
        <button class="flex items-center gap-1 hover:text-[#F08300] py-2 transition uppercase focus:outline-none" aria-haspopup="true" aria-expanded={bidangOpen}>
          <span>Bidang</span>
          <svg class="w-3.5 h-3.5 transition-transform duration-200 group-hover:rotate-180" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
        </button>
        {#if bidangOpen}
          <div class="absolute top-full left-0 w-64 bg-white/95 backdrop-blur-xl border border-slate-200 shadow-xl rounded-2xl py-2 z-50 text-xs normal-case animate-in fade-in zoom-in-95 duration-150">
            <a href="/bidang/sekretariat" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Sekretariat</a>
            <a href="/bidang/perda" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Bidang Penegakan Perda</a>
            <a href="/bidang/trantibum" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Bidang Trantibum</a>
            <a href="/bidang/binmas" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Bidang Pembinaan Masyarakat</a>
            <a href="/bidang/linmas-damkar" class="block px-4 py-2.5 hover:bg-slate-50 hover:text-[#F08300] font-semibold text-slate-800">Bidang Linmas &amp; Damkar</a>
          </div>
        {/if}
      </div>

      <a href="/berita" class="hover:text-[#F08300] transition">Berita</a>
      <a href="/dokumen" class="hover:text-[#F08300] transition">Dokumen</a>
      <a href="/lacak" class="hover:text-[#F08300] transition whitespace-nowrap">Lacak Aduan</a>
      <a href="/kontak" class="hover:text-[#F08300] transition">Kontak</a>
    </nav>

    <!-- Search Engine Trigger & CTA Button -->
    <div class="hidden lg:flex items-center gap-2.5 shrink-0">
      <!-- Search Command Trigger (Understated & Elegan) -->
      <button
        type="button"
        on:click={() => searchOpen = true}
        class="inline-flex items-center gap-1.5 px-2.5 xl:px-3 py-2 text-xs text-slate-500 bg-slate-100 hover:bg-slate-200/80 rounded-xl border border-slate-200 transition focus:outline-none focus:ring-2 focus:ring-[#F08300]/40 shrink-0"
        title="Buka Pencarian Cepat (Ctrl+K)"
      >
        <Search class="w-3.5 h-3.5 text-slate-400" />
        <span class="hidden xl:inline">Cari...</span>
        <kbd class="bg-white px-1.5 py-0.5 rounded text-[10px] font-mono text-slate-400 border border-slate-200 shadow-2xs">Ctrl K</kbd>
      </button>

      <!-- Call-to-Action: SIAP TERTIB (Praja Amber Elegan, Tidak Terlalu Mencolok) -->
      <a 
        href="/lapor" 
        class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-4 xl:px-5 py-2.5 rounded-xl transition shadow-xs hover:shadow-md transform hover:-translate-y-0.5 whitespace-nowrap shrink-0"
      >
        Lapor SIAP TERTIB
      </a>
    </div>

    <!-- Mobile Search & Menu Controls -->
    <div class="flex items-center gap-1 lg:hidden">
      <button
        type="button"
        on:click={() => searchOpen = true}
        class="p-2 rounded-lg text-slate-700 hover:text-[#F08300] hover:bg-slate-100 transition focus:outline-none"
        aria-label="Cari"
      >
        <Search class="w-5 h-5" />
      </button>

      <button 
        type="button"
        class="p-2 rounded-lg text-slate-800 hover:text-[#F08300] hover:bg-slate-100 transition focus:outline-none"
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

  <!-- Fullscreen Mobile Drawer -->
  {#if mobileMenuOpen}
    <div 
      class="fixed inset-0 z-[100] bg-[#0b172a]/95 backdrop-blur-xl text-white flex flex-col justify-between overflow-y-auto lg:hidden animate-in fade-in duration-200"
    >
      <!-- Mobile Bar Top with Close Button -->
      <div class="flex items-center justify-between px-6 py-4 border-b border-white/10 bg-black/20">
        <a href="/" on:click={() => mobileMenuOpen = false}>
          <img 
            src="/uploads/foto_website/1686928245logo_white.png" 
            alt="Logo Satpol PP" 
            class="h-10 w-auto object-contain" 
          />
        </a>
        <button 
          type="button" 
          class="p-2 rounded-lg bg-white/10 hover:bg-white/20 text-white transition focus:outline-none"
          on:click={() => mobileMenuOpen = false}
          aria-label="Tutup Menu"
        >
          <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>

      <!-- Mobile Navigation Links (Full Screen Scrollable) -->
      <div class="px-6 py-6 space-y-6 flex-1 overflow-y-auto">
        <a 
          href="/" 
          class="block text-lg font-bold uppercase tracking-wider text-white hover:text-[#F08300] py-1"
          on:click={() => mobileMenuOpen = false}
        >
          Beranda
        </a>

        <!-- Profil Group -->
        <div class="border-t border-white/10 pt-4">
          <div class="text-xs text-[#F08300] font-bold uppercase tracking-widest mb-2">Profil Satpol PP</div>
          <div class="pl-3 space-y-2.5 text-sm text-slate-300">
            <a href="/profil/sejarah" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Sejarah Singkat</a>
            <a href="/profil/visi-misi" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Visi &amp; Misi</a>
            <a href="/profil/struktur" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Struktur Organisasi</a>
            <a href="/profil/tupoksi" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Tugas Pokok &amp; Fungsi</a>
            <a href="/profil/pejabat" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Profil Pejabat</a>
          </div>
        </div>

        <!-- Bidang Group -->
        <div class="border-t border-white/10 pt-4">
          <div class="text-xs text-[#F08300] font-bold uppercase tracking-widest mb-2">Bidang Kerja</div>
          <div class="pl-3 space-y-2.5 text-sm text-slate-300">
            <a href="/bidang/sekretariat" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Sekretariat</a>
            <a href="/bidang/perda" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Penegakan Perda</a>
            <a href="/bidang/trantibum" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Trantibum</a>
            <a href="/bidang/binmas" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Binmas</a>
            <a href="/bidang/linmas-damkar" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Linmas &amp; Damkar</a>
          </div>
        </div>

        <!-- Menu Lainnya -->
        <div class="border-t border-white/10 pt-4 space-y-3 text-base font-semibold">
          <a href="/berita" class="block hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Berita Publikasi</a>
          <a href="/dokumen" class="block hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Dokumen Publik</a>
          <a href="/lacak" class="block hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Lacak Tiket Aduan</a>
          <a href="/kontak" class="block hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Kontak Resmi</a>
        </div>
      </div>

      <!-- Mobile Footer Action -->
      <div class="p-6 border-t border-white/10 bg-black/30 space-y-3">
        <a 
          href="/lapor" 
          class="block w-full bg-[#F08300] hover:bg-[#d97400] text-center text-white font-bold py-3.5 rounded-xl uppercase tracking-wider text-xs shadow-lg transition"
          on:click={() => mobileMenuOpen = false}
        >
          Lapor SIAP TERTIB
        </a>
      </div>
    </div>
  {/if}
</header>

<CommandPalette bind:open={searchOpen} />
