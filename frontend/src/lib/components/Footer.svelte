<script lang="ts">
  export let compact = false;
</script>

<footer class="bg-[#070e1c]/85 text-slate-300 text-xs border-t border-white/10 mt-auto backdrop-blur-2xl ds-specular-dark transition-colors">
  {#if !compact}
    <div class="max-w-7xl mx-auto px-4 sm:px-6 pt-10 pb-8 grid grid-cols-1 md:grid-cols-4 gap-6 md:gap-8">
      <!-- Col 1: Instansi & Logo -->
      <div class="space-y-3.5">
        <img src="/uploads/foto_website/1686928245logo_white.png" alt="Logo Satpol PP Sulsel" class="h-10 sm:h-11 w-auto object-contain" />
        <p class="leading-relaxed text-slate-400 text-[11.5px] line-clamp-3 font-light">
          Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan bertugas menyelenggarakan urusan pemerintahan daerah di bidang ketenteraman, ketertiban umum, dan pelindungan masyarakat.
        </p>
        <div class="text-[#F08300] font-semibold uppercase tracking-wider text-[11px] pt-0.5">
          Semboyan: Praja Wibawa
        </div>
      </div>

      <!-- Col 2: Profil Satpol PP -->
      <div class="space-y-3">
        <div class="text-xs font-bold text-slate-100 uppercase tracking-wider border-b-2 border-[#F08300]/60 pb-1 inline-block">
          Profil
        </div>
        <ul class="space-y-2 pt-0.5 text-slate-300">
          <li><a href="/profil/sejarah" class="hover:text-[#F08300] transition">Sejarah Singkat</a></li>
          <li><a href="/profil/visi-misi" class="hover:text-[#F08300] transition">Visi &amp; Misi</a></li>
          <li><a href="/profil/struktur" class="hover:text-[#F08300] transition">Struktur Organisasi</a></li>
          <li><a href="/profil/tupoksi" class="hover:text-[#F08300] transition">Tugas Pokok &amp; Fungsi</a></li>
          <li><a href="/profil/pejabat" class="hover:text-[#F08300] transition">Profil Pejabat</a></li>
        </ul>
      </div>

      <!-- Col 3: Bidang Kerja -->
      <div class="space-y-3">
        <div class="text-xs font-bold text-slate-100 uppercase tracking-wider border-b-2 border-[#F08300]/60 pb-1 inline-block">
          Bidang
        </div>
        <ul class="space-y-2 pt-0.5 text-slate-300">
          <li><a href="/bidang/sekretariat" class="hover:text-[#F08300] transition">Sekretariat</a></li>
          <li><a href="/bidang/perda" class="hover:text-[#F08300] transition">Bidang Penegakan Perda</a></li>
          <li><a href="/bidang/trantibum" class="hover:text-[#F08300] transition">Bidang Trantibum</a></li>
          <li><a href="/bidang/binmas" class="hover:text-[#F08300] transition">Bidang Pembinaan Masyarakat</a></li>
          <li><a href="/bidang/linmas-damkar" class="hover:text-[#F08300] transition">Bidang Linmas &amp; Damkar</a></li>
        </ul>
      </div>

      <!-- Col 4: Alamat & Kontak Resmi -->
      <div class="space-y-3">
        <div class="text-xs font-bold text-slate-100 uppercase tracking-wider border-b-2 border-[#F08300]/60 pb-1 inline-block">
          Alamat Kantor
        </div>
        <p class="leading-relaxed text-slate-300 text-xs font-light">
          Jl. Urip Sumoharjo No. 269, Kota Makassar, Sulawesi Selatan<br/>
          <span class="text-slate-100 font-medium">Telepon:</span> (0411) 453105<br/>
          <span class="text-slate-100 font-medium">Email:</span> satpol@sulselprov.go.id
        </p>
        <div class="text-[11px] text-slate-400 pt-1.5 border-t border-white/10 font-light">
          Infrastruktur Server didukung oleh Diskominfo-SP Prov. Sulsel.
        </div>
      </div>
    </div>
  {/if}

  <!-- Bottom Bar: Harmonious Obsidian Frosted Strip -->
  <div class="bg-[#050a14]/90 border-t border-white/5 py-3 sm:py-3.5 px-4 text-center">
    <div class="max-w-7xl mx-auto flex flex-col sm:flex-row justify-between items-center gap-2 text-xs text-slate-400 font-medium">
      <p>
        &copy; 2026 Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan. Hak Cipta Dilindungi Undang-Undang.
      </p>
      <div class="flex items-center gap-4 text-[11px] text-slate-400">
        <a href="/dokumen?category=perda" class="hover:text-[#F08300] transition">Regulasi Daerah</a>
        <span>&bull;</span>
        <a href="/kontak" class="hover:text-[#F08300] transition">Kontak Resmi</a>
        <span>&bull;</span>
        <span class="text-[#F08300] font-semibold">Praja Wibawa</span>
      </div>
    </div>
  </div>
</footer>
