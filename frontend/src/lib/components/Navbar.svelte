<script lang="ts">
  let mobileMenuOpen = false;
  let profilOpen = false;
  let bidangOpen = false;
</script>

<header class="w-full sticky top-0 z-50 bg-white/95 backdrop-blur-md shadow-xs border-b border-[#eee9db]">
  <!-- Top Bar: Official Contacts (Cokelat Tua Praja #1a120c, Kontras Tinggi & Tanpa Tombol Login) -->
  <div class="bg-[#1a120c] text-stone-200 text-xs py-2 px-4 border-b border-[#3d261a]">
    <div class="max-w-7xl mx-auto flex flex-wrap justify-between items-center gap-2">
      <div class="flex items-center gap-4 sm:gap-6 text-[11px] sm:text-xs">
        <a href="tel:0411-453105" class="hover:text-[#F08300] transition flex items-center gap-1.5">
          <svg class="w-3.5 h-3.5 text-[#F08300]" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" /></svg>
          <span>Telepon: (0411) 453105</span>
        </a>
        <a href="mailto:satpol@sulselprov.go.id" class="hidden sm:inline hover:text-[#F08300] transition">Email: satpol@sulselprov.go.id</a>
        <a href="https://ppidsatpolppsulsel.my.id/" target="_blank" rel="noreferrer" class="hidden md:inline hover:text-[#F08300] transition">Layanan PPID</a>
      </div>
      <div class="flex items-center gap-2 text-[11px] text-stone-300">
        <span class="inline-block w-2 h-2 rounded-full bg-emerald-400"></span>
        <span>Pemerintah Provinsi Sulawesi Selatan</span>
      </div>
    </div>
  </div>

  <!-- Main Navigation Bar -->
  <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
    <!-- Poin 4: Official Logo saja (Teks instansi HTML di samping dihilangkan) -->
    <a href="/" class="flex items-center group py-0.5" title="Beranda Satpol PP Sulsel">
      <img 
        src="/uploads/foto_website/1686928245logo_black.png" 
        alt="Logo Satpol PP Prov. Sulsel" 
        class="h-11 sm:h-12 w-auto object-contain transition duration-200 group-hover:opacity-90" 
      />
    </a>

    <!-- Desktop Navigation Menu Links -->
    <nav class="hidden lg:flex items-center gap-7 text-[13px] font-bold text-[#101116] uppercase tracking-wider">
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
          <div class="absolute top-full left-0 w-60 bg-white border border-[#eee9db] shadow-xl rounded-xl py-2 z-50 text-xs normal-case">
            <a href="/profil/sejarah" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Sejarah Singkat</a>
            <a href="/profil/visi-misi" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Visi &amp; Misi</a>
            <a href="/profil/struktur" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Struktur Organisasi</a>
            <a href="/profil/tupoksi" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Tugas Pokok &amp; Fungsi</a>
            <a href="/profil/pejabat" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Profil Pejabat</a>
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
          <div class="absolute top-full left-0 w-64 bg-white border border-[#eee9db] shadow-xl rounded-xl py-2 z-50 text-xs normal-case">
            <a href="/bidang/sekretariat" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Sekretariat</a>
            <a href="/bidang/perda" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Bidang Penegakan Perda</a>
            <a href="/bidang/trantibum" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Bidang Trantibum</a>
            <a href="/bidang/binmas" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Bidang Pembinaan Masyarakat</a>
            <a href="/bidang/linmas-damkar" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold text-[#101116]">Bidang Linmas &amp; Damkar</a>
          </div>
        {/if}
      </div>

      <a href="/lacak" class="hover:text-[#F08300] transition">Lacak Aduan</a>
      <a href="/kontak" class="hover:text-[#F08300] transition">Kontak</a>
    </nav>

    <!-- Call-to-Action: SIAP TERTIB -->
    <div class="hidden lg:flex items-center gap-3">
      <a 
        href="/lapor" 
        class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-5 py-2.5 rounded-lg transition shadow-xs hover:shadow-md"
      >
        Lapor SIAP TERTIB
      </a>
    </div>

    <!-- Mobile Menu Button (Hamburger) -->
    <button 
      type="button"
      class="lg:hidden p-2 rounded-lg text-[#101116] hover:text-[#F08300] hover:bg-[#f9f4e8] transition focus:outline-none"
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

  <!-- Poin 6: Fullscreen Mobile Drawer (Cokelat Tua Praja #241710, menggantikan biru tua) -->
  {#if mobileMenuOpen}
    <div 
      class="fixed inset-0 z-[100] bg-[#241710]/95 backdrop-blur-md text-white flex flex-col justify-between overflow-y-auto lg:hidden animate-in fade-in duration-200"
    >
      <!-- Mobile Bar Top with Close Button -->
      <div class="flex items-center justify-between px-6 py-4 border-b border-white/10 bg-black/30">
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
          <div class="pl-3 space-y-2.5 text-sm text-stone-200">
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
          <div class="pl-3 space-y-2.5 text-sm text-stone-200">
            <a href="/bidang/sekretariat" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Sekretariat</a>
            <a href="/bidang/perda" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Penegakan Perda</a>
            <a href="/bidang/trantibum" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Trantibum</a>
            <a href="/bidang/binmas" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Binmas</a>
            <a href="/bidang/linmas-damkar" class="block hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Linmas &amp; Damkar</a>
          </div>
        </div>

        <!-- Menu Lainnya -->
        <div class="border-t border-white/10 pt-4 space-y-3 text-base font-semibold">
          <a href="/lacak" class="block hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Lacak Tiket Aduan</a>
          <a href="/kontak" class="block hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Kontak Resmi</a>
        </div>
      </div>

      <!-- Mobile Footer Action (Tanpa Tombol Login Sesuai Poin 22) -->
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
