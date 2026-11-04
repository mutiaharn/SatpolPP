<script lang="ts">
  import { onMount } from 'svelte';
  import { articlesStore, galleryStore } from '$lib/stores/appState';
  import ComplaintForm from '$lib/components/ComplaintForm.svelte';
  import TicketTracker from '$lib/components/TicketTracker.svelte';

  let currentSlide = 0;
  let activeTab: 'lapor' | 'lacak' = 'lapor';

  const slides = [
    {
      image: '/uploads/foto_banner/1716426402berita.png',
      tag: 'PRAJA WIBAWA',
      title: 'Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan',
      desc: 'Menegakkan Peraturan Daerah, Menyelenggarakan Ketenteraman dan Ketertiban Umum, serta Pelindungan Masyarakat.'
    },
    {
      image: '/uploads/foto_banner/1747198131welcome.png',
      tag: 'PELAYANAN PUBLIK',
      title: 'SIAP TERTIB: Sistem Informasi Penegakan Perda & Trantibum',
      desc: 'Saluran pengaduan elektronik terpadu masyarakat untuk mewujudkan Sulawesi Selatan yang aman, tertib, dan kondusif.'
    },
    {
      image: '/uploads/foto_banner/174089363110.png',
      tag: 'PENGAWALAN ASET',
      title: 'Penertiban dan Pengamanan Aset Pemerintah Daerah',
      desc: 'Optimalisasi dan penertiban aset barang milik daerah demi pemanfaatan yang tertib dan berkesinambungan.'
    }
  ];

  onMount(() => {
    // Auto rotate banner every 6 seconds
    const timer = setInterval(() => {
      currentSlide = (currentSlide + 1) % slides.length;
    }, 6000);
    return () => clearInterval(timer);
  });
</script>

<!-- 1. Hero Banner Slider (Matching Main Slider Satpol PP) -->
<section class="relative bg-[#101116] text-white min-h-[480px] md:min-h-[560px] flex items-center overflow-hidden">
  <!-- Slide Images -->
  {#each slides as slide, idx}
    <div 
      class={`absolute inset-0 bg-cover bg-center transition-opacity duration-1000 ${currentSlide === idx ? 'opacity-100 z-10' : 'opacity-0 z-0'}`}
      style={`background-image: linear-gradient(rgba(16, 17, 22, 0.72), rgba(16, 17, 22, 0.78)), url('${slide.image}')`}
    ></div>
  {/each}

  <div class="max-w-7xl mx-auto px-4 py-16 relative z-20 w-full">
    <div class="max-w-2xl space-y-4">
      <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">
        {slides[currentSlide].tag}
      </div>
      <h1 class="text-3xl md:text-5xl font-extrabold tracking-tight leading-tight">
        {slides[currentSlide].title}
      </h1>
      <p class="text-sm md:text-base text-[#e5e7eb] leading-relaxed">
        {slides[currentSlide].desc}
      </p>
      <div class="pt-4 flex flex-wrap gap-3">
        <a href="#layanan-siaptertib" class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-6 py-3.5 rounded transition shadow-sm">
          Lapor SIAP TERTIB
        </a>
        <a href="/dokumen" class="bg-transparent hover:bg-white/10 text-white font-bold text-xs uppercase px-6 py-3.5 rounded border border-white/30 transition">
          Dokumen Publik
        </a>
      </div>
    </div>
  </div>

  <!-- Slide Navigation Indicators -->
  <div class="absolute bottom-6 right-6 z-30 flex items-center gap-2">
    {#each slides as _, i}
      <button 
        class={`h-2.5 transition-all rounded-full ${currentSlide === i ? 'w-8 bg-[#F08300]' : 'w-2.5 bg-white/40'}`}
        on:click={() => currentSlide = i}
        aria-label={`Slide ${i + 1}`}
      ></button>
    {/each}
  </div>
</section>

<!-- 2. Berita Terbaru (Matching news-one) -->
<section class="py-16 px-4 max-w-7xl mx-auto">
  <div class="text-center max-w-xl mx-auto mb-12">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Publikasi Resmi</div>
    <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Berita Terbaru</h2>
    <div class="w-12 h-1 bg-[#F08300] mx-auto mt-3"></div>
  </div>

  <div class="grid grid-cols-1 md:grid-cols-3 gap-8">
    {#each $articlesStore.slice(0, 3) as item}
      <article class="bg-white border border-[#eee9db] rounded-2xl overflow-hidden flex flex-col hover:shadow-lg transition">
        <div class="relative h-48 bg-gray-100 overflow-hidden">
          <img 
            src={item.thumbnail} 
            alt={item.title} 
            class="w-full h-full object-cover hover:scale-105 transition duration-500" 
            onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
          />
          <div class="absolute top-0 right-0 bg-[#F08300] text-white text-[11px] font-bold px-3 py-1.5 uppercase rounded-bl-xl">
            {new Date(item.published_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'short' })}
          </div>
        </div>
        <div class="p-6 flex-1 flex flex-col justify-between space-y-3">
          <div>
            <div class="text-[11px] font-bold uppercase tracking-wider text-[#F08300]">
              {item.category}
            </div>
            <h3 class="font-bold text-base text-[#101116] leading-snug mt-1 hover:text-[#F08300] transition">
              <a href={`/berita/${item.slug}`}>{item.title}</a>
            </h3>
            <p class="text-xs text-[#6f7775] leading-relaxed mt-2 line-clamp-3">
              {item.summary}
            </p>
          </div>
          <div class="pt-4 border-t border-[#eee9db]">
            <a href={`/berita/${item.slug}`} class="text-xs font-bold text-[#F08300] uppercase tracking-wider hover:underline">
              Selengkapnya &rarr;
            </a>
          </div>
        </div>
      </article>
    {/each}
  </div>
</section>

<!-- 3. Profil Pimpinan Satpol PP Sulsel (Matching about-one) -->
<section class="py-16 px-4 bg-[#f9f4e8] border-t border-b border-[#eee9db]">
  <div class="max-w-7xl mx-auto grid grid-cols-1 lg:grid-cols-12 gap-12 items-center">
    <!-- Left: Portrait Photo with Border -->
    <div class="lg:col-span-5 flex justify-center">
      <div class="relative border-4 border-white shadow-xl max-w-sm">
        <img 
          src="/uploads/foto_pimpinan/1703147058Untitled.png" 
          alt="Andi Arwin Azis S.STP., M.M" 
          class="w-full h-auto object-cover" 
        />
        <div class="bg-[#101116] text-white p-4 text-center">
          <div class="font-bold text-sm tracking-tight">ANDI ARWIN AZIS S.STP., M.M</div>
          <div class="text-[10px] text-[#F08300] font-semibold uppercase tracking-wider mt-0.5">Kepala Satuan Polisi Pamong Praja Prov. Sulsel</div>
        </div>
      </div>
    </div>

    <!-- Right: Sambutan & Tugas Pokok -->
    <div class="lg:col-span-7 space-y-4">
      <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Pimpinan Instansi</div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight">
        Menegakkan Perda Secara Humanis dan Berkeadilan
      </h2>
      <div class="w-12 h-1 bg-[#F08300]"></div>

      <p class="text-xs md:text-sm text-[#2b2f3f] leading-relaxed pt-2">
        Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan senantiasa mengedepankan pendekatan humanis, persuasif, dan profesional dalam menegakkan Peraturan Daerah dan menjaga ketertiban umum demi kesejahteraan masyarakat Sulawesi Selatan.
      </p>

      <div class="pt-2 space-y-2 text-xs text-[#2b2f3f]">
        <div class="flex items-start gap-2">
          <span class="text-[#F08300] font-bold">&bull;</span>
          <span>Penyelenggaraan ketenteraman dan ketertiban umum di seluruh wilayah Sulawesi Selatan.</span>
        </div>
        <div class="flex items-start gap-2">
          <span class="text-[#F08300] font-bold">&bull;</span>
          <span>Penegakan Peraturan Daerah (Perda) dan Peraturan Kepala Daerah (Perkada).</span>
        </div>
        <div class="flex items-start gap-2">
          <span class="text-[#F08300] font-bold">&bull;</span>
          <span>Pengawalan dan pengamanan aset milik Pemerintah Provinsi Sulawesi Selatan.</span>
        </div>
        <div class="flex items-start gap-2">
          <span class="text-[#F08300] font-bold">&bull;</span>
          <span>Pembinaan, bimbingan masyarakat, dan koordinasi perlindungan masyarakat (Linmas).</span>
        </div>
      </div>

      <div class="pt-4">
        <a href="/profil/pejabat" class="inline-block bg-[#2b2f3f] hover:bg-[#101116] text-white text-xs font-bold uppercase px-6 py-3 rounded transition">
          Profil Struktur Organisasi &rarr;
        </a>
      </div>
    </div>
  </div>
</section>

<!-- 4. Dokumen Publik (Matching events-one) -->
<section class="py-16 px-4 max-w-7xl mx-auto">
  <div class="text-center max-w-xl mx-auto mb-12">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Akuntabilitas Kinerja</div>
    <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Dokumen Publik</h2>
    <div class="w-12 h-1 bg-[#F08300] mx-auto mt-3"></div>
  </div>

  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
    <div class="bg-white border border-[#eee9db] p-6 hover:border-[#F08300] transition">
      <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Perencanaan</div>
      <h3 class="font-bold text-base text-[#101116] mt-1">Renja (Rencana Kerja)</h3>
      <p class="text-xs text-[#6f7775] mt-2">Dokumen perencanaan tahunan program dan target kegiatan Satpol PP Prov. Sulsel.</p>
      <div class="mt-4 pt-3 border-t border-gray-100">
        <a href="/dokumen?category=renja" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
      </div>
    </div>

    <div class="bg-white border border-[#eee9db] p-6 hover:border-[#F08300] transition">
      <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Strategis</div>
      <h3 class="font-bold text-base text-[#101116] mt-1">Renstra (Rencana Strategis)</h3>
      <p class="text-xs text-[#6f7775] mt-2">Rencana strategis jangka menengah pencapaian visi, misi, dan sasaran instansi.</p>
      <div class="mt-4 pt-3 border-t border-gray-100">
        <a href="/dokumen?category=renstra" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
      </div>
    </div>

    <div class="bg-white border border-[#eee9db] p-6 hover:border-[#F08300] transition">
      <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Penganggaran</div>
      <h3 class="font-bold text-base text-[#101116] mt-1">DPA (Pelaksanaan Anggaran)</h3>
      <p class="text-xs text-[#6f7775] mt-2">Alokasi anggaran belanja dan rincian pembiayaan kegiatan operasional.</p>
      <div class="mt-4 pt-3 border-t border-gray-100">
        <a href="/dokumen?category=dpa" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
      </div>
    </div>

    <div class="bg-white border border-[#eee9db] p-6 hover:border-[#F08300] transition">
      <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Kinerja</div>
      <h3 class="font-bold text-base text-[#101116] mt-1">LAKIP / LKjIP</h3>
      <p class="text-xs text-[#6f7775] mt-2">Laporan akuntabilitas kinerja instansi pemerintah tahunan.</p>
      <div class="mt-4 pt-3 border-t border-gray-100">
        <a href="/dokumen?category=lakip" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
      </div>
    </div>

    <div class="bg-white border border-[#eee9db] p-6 hover:border-[#F08300] transition">
      <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Pertanggungjawaban</div>
      <h3 class="font-bold text-base text-[#101116] mt-1">Laporan Keuangan</h3>
      <p class="text-xs text-[#6f7775] mt-2">Laporan pertanggungjawaban realisasi keuangan dan aset tahunan.</p>
      <div class="mt-4 pt-3 border-t border-gray-100">
        <a href="/dokumen?category=keuangan" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
      </div>
    </div>

    <div class="bg-white border border-[#eee9db] p-6 hover:border-[#F08300] transition">
      <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Regulasi</div>
      <h3 class="font-bold text-base text-[#101116] mt-1">Perda &amp; SOP Satpol PP</h3>
      <p class="text-xs text-[#6f7775] mt-2">Peraturan Daerah Provinsi Sulawesi Selatan dan Standar Operasional Prosedur.</p>
      <div class="mt-4 pt-3 border-t border-gray-100">
        <a href="/dokumen?category=perda" class="text-xs font-bold text-[#F08300] uppercase hover:underline">Lihat Dokumen &rarr;</a>
      </div>
    </div>
  </div>
</section>

<!-- 5. Layanan SIAP TERTIB & Pelacakan Aduan (Fitur Penggabungan Terpadu) -->
<section id="layanan-siaptertib" class="py-16 px-4 bg-[#f9f4e8] border-t border-b border-[#eee9db]">
  <div class="max-w-4xl mx-auto">
    <div class="text-center max-w-xl mx-auto mb-8">
      <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Layanan Pengaduan Masyarakat</div>
      <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">SIAP TERTIB</h2>
      <div class="w-12 h-1 bg-[#F08300] mx-auto mt-3"></div>
      <p class="text-xs text-[#6f7775] mt-3">
        Sistem Informasi Penegakan Peraturan Daerah &amp; Ketenteraman Ketertiban Umum.
      </p>
    </div>

    <!-- Tab Selector -->
    <div class="flex justify-center mb-6">
      <div class="inline-flex border border-gray-300 bg-white p-1">
        <button 
          class={`px-6 py-2 text-xs font-bold uppercase tracking-wider transition ${activeTab === 'lapor' ? 'bg-[#F08300] text-white' : 'text-[#101116] hover:bg-gray-100'}`}
          on:click={() => activeTab = 'lapor'}
        >
          Formulir Pengaduan
        </button>
        <button 
          class={`px-6 py-2 text-xs font-bold uppercase tracking-wider transition ${activeTab === 'lacak' ? 'bg-[#F08300] text-white' : 'text-[#101116] hover:bg-gray-100'}`}
          on:click={() => activeTab = 'lacak'}
        >
          Lacak Resi Aduan
        </button>
      </div>
    </div>

    {#if activeTab === 'lapor'}
      <ComplaintForm />
    {:else}
      <TicketTracker />
    {/if}
  </div>
</section>

<!-- 6. Galeri Kegiatan (Matching gallery-one) -->
<section class="py-16 px-4 max-w-7xl mx-auto">
  <div class="text-center max-w-xl mx-auto mb-12">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Dokumentasi</div>
    <h2 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Galeri Kegiatan</h2>
    <div class="w-12 h-1 bg-[#F08300] mx-auto mt-3"></div>
  </div>

  <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
    {#each $galleryStore.slice(0, 4) as gal}
      <div class="h-44 bg-gray-100 overflow-hidden border border-[#eee9db] rounded-2xl shadow-xs group">
        <img src={gal.image_url} alt={gal.title} class="w-full h-full object-cover group-hover:scale-105 transition duration-300" />
      </div>
    {/each}
  </div>
</section>
