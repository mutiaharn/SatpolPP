<script lang="ts">
  let mobileMenuOpen = false;
  let profilOpen = false;
  let bidangOpen = false;
</script>

<header class="w-full sticky top-0 z-50 bg-white shadow-sm border-b border-[#eee9db]">
  <!-- Top Bar: Official Contacts -->
  <div class="bg-[#101116] text-[#e5e7eb] text-xs py-2 px-4 border-b border-black/10">
    <div class="max-w-7xl mx-auto flex flex-wrap justify-between items-center gap-2">
      <div class="flex items-center gap-6">
        <a href="tel:0411-453105" class="hover:text-[#F08300] transition">Telepon: (0411) 453105</a>
        <a href="mailto:satpol@sulselprov.go.id" class="hidden sm:inline hover:text-[#F08300] transition">Email: satpol@sulselprov.go.id</a>
        <a href="https://ppidsatpolppsulsel.my.id/" target="_blank" rel="noreferrer" class="hidden md:inline hover:text-[#F08300] transition">Layanan PPID</a>
      </div>
      <div class="flex items-center gap-4 text-[11px] text-[#9ca3af]">
        <span>Pemerintah Provinsi Sulawesi Selatan</span>
        <span>|</span>
        <a href="/admin/login" class="text-white hover:text-[#F08300] font-semibold transition">Akses Petugas</a>
        <span>|</span>
        <a href="/superadmin/dashboard" class="text-[#F08300] hover:underline font-semibold transition">Superadmin</a>
      </div>
    </div>
  </div>

  <!-- Main Navigation Bar -->
  <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
    <!-- Official Logo -->
    <a href="/" class="flex items-center gap-3">
      <img src="/uploads/foto_website/1686928245logo_black.png" alt="Logo Satpol PP Prov. Sulsel" class="h-12 w-auto object-contain" />
      <div class="border-l border-gray-300 pl-3 hidden sm:block">
        <div class="text-xs font-extrabold uppercase tracking-tight text-[#101116] leading-none">Satuan Polisi Pamong Praja</div>
        <div class="text-[10px] text-[#6f7775] font-semibold uppercase tracking-wider mt-0.5">Provinsi Sulawesi Selatan</div>
      </div>
    </a>

    <!-- Desktop Navigation Menu Links -->
    <nav class="hidden lg:flex items-center gap-6 text-[13px] font-bold text-[#101116] uppercase tracking-wider">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>

      <!-- Dropdown Profil -->
      <div 
        class="relative group"
        role="menu"
        tabindex="0"
        on:mouseenter={() => profilOpen = true}
        on:mouseleave={() => profilOpen = false}
      >
        <button class="flex items-center gap-1 hover:text-[#F08300] py-2 transition uppercase" aria-haspopup="true" aria-expanded={profilOpen}>
          <span>Profil</span>
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
        </button>
        {#if profilOpen}
          <div class="absolute top-full left-0 w-56 bg-white border border-[#eee9db] shadow-lg py-2 z-50 text-xs normal-case">
            <a href="/profil/sejarah" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Sejarah Singkat</a>
            <a href="/profil/visi-misi" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Visi &amp; Misi</a>
            <a href="/profil/struktur" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Struktur Organisasi</a>
            <a href="/profil/tupoksi" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Tugas Pokok &amp; Fungsi</a>
            <a href="/profil/pejabat" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Profil Pejabat</a>
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
        <button class="flex items-center gap-1 hover:text-[#F08300] py-2 transition uppercase" aria-haspopup="true" aria-expanded={bidangOpen}>
          <span>Bidang</span>
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" /></svg>
        </button>
        {#if bidangOpen}
          <div class="absolute top-full left-0 w-64 bg-white border border-[#eee9db] shadow-lg py-2 z-50 text-xs normal-case">
            <a href="/bidang/sekretariat" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Sekretariat</a>
            <a href="/bidang/perda" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Bidang Penegakan Perda</a>
            <a href="/bidang/trantibum" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Bidang Trantibum</a>
            <a href="/bidang/binmas" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Bidang Pembinaan Masyarakat</a>
            <a href="/bidang/linmas-damkar" class="block px-4 py-2 hover:bg-[#f9f4e8] hover:text-[#F08300] font-semibold">Bidang Linmas &amp; Damkar</a>
          </div>
        {/if}
      </div>

      <a href="/berita" class="hover:text-[#F08300] transition">Berita</a>
      <a href="/dokumen" class="hover:text-[#F08300] transition">Dokumen</a>
      <a href="/galeri" class="hover:text-[#F08300] transition">Galeri</a>
      <a href="/lacak" class="hover:text-[#F08300] transition">Lacak Aduan</a>
      <a href="/kontak" class="hover:text-[#F08300] transition">Kontak</a>
    </nav>

    <!-- Call-to-Action: SIAP TERTIB -->
    <div class="hidden lg:flex items-center gap-3">
      <a 
        href="/lapor" 
        class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-5 py-2.5 transition shadow-sm"
      >
        Lapor SIAP TERTIB
      </a>
    </div>

    <!-- Mobile Menu Button -->
    <button 
      class="lg:hidden p-2 text-[#101116] hover:text-[#F08300]"
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

  <!-- Mobile Drawer -->
  {#if mobileMenuOpen}
    <div class="lg:hidden bg-[#2b2f3f] text-white px-5 py-4 space-y-3 text-xs uppercase font-bold tracking-wider border-t border-[#101116]">
      <a href="/" class="block py-2 hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Beranda</a>
      
      <div class="border-t border-gray-700 pt-2">
        <div class="text-[10px] text-[#F08300] uppercase font-bold mb-1">Profil Satpol PP</div>
        <div class="pl-3 space-y-1 normal-case text-gray-300">
          <a href="/profil/sejarah" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Sejarah Singkat</a>
          <a href="/profil/visi-misi" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Visi &amp; Misi</a>
          <a href="/profil/struktur" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Struktur Organisasi</a>
          <a href="/profil/tupoksi" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Tugas Pokok &amp; Fungsi</a>
          <a href="/profil/pejabat" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Profil Pejabat</a>
        </div>
      </div>

      <div class="border-t border-gray-700 pt-2">
        <div class="text-[10px] text-[#F08300] uppercase font-bold mb-1">Bidang Kerja</div>
        <div class="pl-3 space-y-1 normal-case text-gray-300">
          <a href="/bidang/sekretariat" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Sekretariat</a>
          <a href="/bidang/perda" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Penegakan Perda</a>
          <a href="/bidang/trantibum" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Trantibum</a>
          <a href="/bidang/binmas" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Binmas</a>
          <a href="/bidang/linmas-damkar" class="block py-1 hover:text-white" on:click={() => mobileMenuOpen = false}>Bidang Linmas &amp; Damkar</a>
        </div>
      </div>

      <div class="border-t border-gray-700 pt-2 space-y-2">
        <a href="/berita" class="block py-1 hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Berita &amp; Informasi</a>
        <a href="/dokumen" class="block py-1 hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Dokumen Publik</a>
        <a href="/galeri" class="block py-1 hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Galeri Foto</a>
        <a href="/lacak" class="block py-1 hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Lacak Tiket Aduan</a>
        <a href="/kontak" class="block py-1 hover:text-[#F08300]" on:click={() => mobileMenuOpen = false}>Kontak Resmi</a>
      </div>

      <div class="pt-3 border-t border-gray-700 space-y-2">
        <a href="/lapor" class="block bg-[#F08300] text-center text-white font-bold py-2.5 transition" on:click={() => mobileMenuOpen = false}>
          Lapor SIAP TERTIB
        </a>
        <a href="/admin/login" class="block bg-gray-800 text-center text-white font-bold py-2 border border-gray-600 transition" on:click={() => mobileMenuOpen = false}>
          Login Petugas Satpol
        </a>
      </div>
    </div>
  {/if}
</header>
