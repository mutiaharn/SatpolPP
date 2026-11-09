<script lang="ts">
  import { page } from '$app/stores';
  import { articlesStore } from '$lib/stores/appState';
  import type { Article } from '$lib/types';

  $: slug = $page.params.slug;
  $: article = $articlesStore.find(a => a.slug === slug);
  $: recentArticles = $articlesStore.filter(a => a.slug !== slug).slice(0, 4);

  function formatDate(iso: string) {
    try {
      const d = new Date(iso);
      return d.toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' });
    } catch {
      return iso;
    }
  }
</script>

<svelte:head>
  <title>{article ? article.title : 'Warta Berita'} | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-6 px-4">
  <div class="max-w-6xl mx-auto flex flex-wrap items-center gap-2 text-[11px] text-[#6f7775]">
    <a href="/" class="hover:text-[#F08300]">Beranda</a>
    <span>/</span>
    <a href="/berita" class="hover:text-[#F08300]">Berita &amp; Informasi</a>
    <span>/</span>
    <span class="text-[#101116] font-semibold truncate max-w-md">{article?.title || 'Detail Berita'}</span>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto">
  {#if article}
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-10">
      <!-- Main Content Column -->
      <div class="lg:col-span-2 space-y-6">
        <div>
          <span class="inline-block bg-[#101116] text-[#F08300] font-bold text-[10px] uppercase px-3 py-1 rounded-sm border-l-2 border-[#F08300] tracking-wider mb-3">
            {article.category}
          </span>
          <h1 class="text-2xl md:text-3xl font-extrabold text-[#101116] tracking-tight leading-tight">
            {article.title}
          </h1>

          <div class="flex flex-wrap items-center gap-4 text-xs text-[#6f7775] mt-4 pt-3 border-t border-gray-100">
            <div>Oleh: <strong class="text-[#101116]">{article.author_name}</strong></div>
            <div>&bull;</div>
            <div>Tanggal: {formatDate(article.published_at)}</div>
            <div>&bull;</div>
            <div>Dilihat: {article.view_count} kali</div>
          </div>
        </div>

        <!-- Featured Media -->
        <div class="border border-[#eee9db] rounded-2xl overflow-hidden bg-gray-50 shadow-xs">
          <img 
            src={article.thumbnail} 
            alt={article.title} 
            class="w-full max-h-[440px] object-cover"
            onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
          />
          <div class="p-3 text-[11px] text-[#6f7775] italic bg-[#f9f4e8] border-t border-[#eee9db]">
            Dokumentasi: Personel Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan
          </div>
        </div>

        <!-- Article Summary & Body Content -->
        <div class="prose max-w-none text-xs md:text-sm text-[#2b2f3f] leading-relaxed space-y-4">
          <p class="font-bold text-sm text-[#101116] leading-relaxed bg-[#fcfbf7] border-l-4 border-[#F08300] p-4 rounded-r-xl">
            {article.summary}
          </p>

          <div class="space-y-4 pt-2">
            {@html article.content}
            <p>
              Pelaksanaan tugas penegakan Peraturan Daerah dan perlindungan masyarakat oleh jajaran Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan selalu didasarkan pada prinsip profesionalisme, akuntabilitas, serta kepatuhan pada regulasi hukum yang berlaku. Sinergi antara pemerintah daerah, aparat penegak hukum, dan partisipasi aktif masyarakat merupakan pilar utama terwujudnya ketertiban umum yang berkesinambungan.
            </p>
            <p>
              Masyarakat diimbau untuk turut menjaga fasilitas publik serta segera melaporkan setiap potensi gangguan ketenteraman umum melalui kanal pengaduan resmi SIAP TERTIB Satpol PP Prov. Sulsel.
            </p>
          </div>
        </div>

        <!-- Back Button & Social Sharing -->
        <div class="pt-6 border-t border-gray-200 flex flex-wrap justify-between items-center gap-3">
          <a 
            href="/berita" 
            class="bg-gray-100 hover:bg-gray-200 text-[#101116] text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition"
          >
            &larr; Kembali ke Indeks Berita
          </a>
          <a 
            href="/lapor" 
            class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition shadow-xs"
          >
            Laporkan Pelanggaran di Wilayah Anda
          </a>
        </div>
      </div>

      <!-- Sidebar Column -->
      <div class="space-y-6">
        <!-- SIAP TERTIB Quick CTA -->
        <div class="bg-[#3d261a] text-white p-6 rounded-xl border-t-4 border-[#F08300] shadow-sm">
          <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider mb-1">
            Kanal Pengaduan Warga
          </div>
          <h3 class="text-base font-bold">Layanan SIAP TERTIB</h3>
          <p class="text-xs text-[#d1d5db] mt-2 leading-relaxed">
            Temukan pelanggaran ketertiban umum, pedagang liar, atau gangguan kenyamanan publik di Sulawesi Selatan? Laporkan langsung secara online.
          </p>
          <a 
            href="/lapor" 
            class="mt-4 block text-center bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase py-2.5 rounded-xl transition shadow-xs"
          >
            Kirim Laporan Pengaduan
          </a>
        </div>

        <!-- Latest Articles List -->
        <div class="bg-white border border-[#eee9db] rounded-2xl shadow-xs p-6">
          <h3 class="text-xs font-bold text-[#101116] uppercase tracking-wider border-b border-gray-100 pb-3 mb-4">
            Berita Terkini Lainnya
          </h3>
          <div class="space-y-4">
            {#each recentArticles as rec}
              <div class="flex gap-3 group">
                <img 
                  src={rec.thumbnail} 
                  alt={rec.title} 
                  class="w-16 h-16 object-cover flex-shrink-0 bg-gray-100 rounded-xl"
                  onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
                />
                <div class="flex-1 min-w-0">
                  <div class="text-[10px] text-[#6f7775]">{formatDate(rec.published_at)}</div>
                  <a 
                    href={`/berita/${rec.slug}`} 
                    class="text-xs font-bold text-[#101116] group-hover:text-[#F08300] line-clamp-2 leading-snug transition mt-0.5"
                  >
                    {rec.title}
                  </a>
                </div>
              </div>
            {/each}
          </div>
        </div>

        <!-- Contact Hotline Info -->
        <div class="bg-[#f9f4e8] border border-[#eee9db] rounded-2xl p-5 text-xs text-[#2b2f3f] space-y-2">
          <div class="font-bold uppercase text-[11px] text-[#101116]">Posko Komando Satpol PP</div>
          <p class="text-[11px] text-[#6f7775]">Jl. Urip Sumoharjo No. 269, Makassar</p>
          <div class="pt-2 border-t border-[#eee9db] space-y-1 text-[11px]">
            <div>Telepon: <strong>0411-453105</strong></div>
            <div>Hotline Darurat: <strong>112 / 0811-4141-112</strong></div>
          </div>
        </div>
      </div>
    </div>
  {:else}
    <div class="bg-white border border-[#eee9db] rounded-2xl p-12 text-center text-xs space-y-3">
      <div class="text-base font-bold text-[#101116]">Artikel Tidak Ditemukan</div>
      <p class="text-[#6f7775]">Artikel yang Anda cari tidak tersedia atau tautan telah dipindahkan.</p>
      <a href="/berita" class="inline-block bg-[#F08300] text-white text-xs font-bold uppercase px-4 py-2 mt-2 rounded-xl">
        Kembali ke Indeks Berita
      </a>
    </div>
  {/if}
</div>
