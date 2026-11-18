<script lang="ts">
  import { page } from '$app/stores';
  import { articlesStore } from '$lib/stores/appState';
  import type { Article } from '$lib/types';
  import { ChevronRight, Calendar, User, Eye, ArrowLeft, Shield, Phone, AlertCircle } from 'lucide-svelte';

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

<!-- Sovereign Slim Hero / Breadcrumbs -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#246bfd]/8 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>
  <div class="max-w-6xl mx-auto relative z-10">
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <a href="/berita" class="hover:text-[#F08300] transition">Warta Berita</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium truncate max-w-xs sm:max-w-md">{article?.title || 'Detail Berita'}</span>
    </div>
    {#if article}
      <h1 class="text-xl sm:text-2xl md:text-3xl font-extrabold text-white tracking-tight leading-tight max-w-4xl">
        {article.title}
      </h1>
      <div class="flex flex-wrap items-center gap-3 text-xs text-slate-300 mt-2.5 pt-2 border-t border-white/10 font-light">
        <span class="px-2 py-0.5 rounded-md bg-orange-500/15 text-[#F08300] font-bold text-[10px] uppercase tracking-wider border border-orange-500/30">
          {article.category}
        </span>
        <span class="flex items-center gap-1.5">
          <User class="w-3.5 h-3.5 text-slate-400" />
          <strong class="font-medium text-white">{article.author_name}</strong>
        </span>
        <span>&bull;</span>
        <span class="flex items-center gap-1.5">
          <Calendar class="w-3.5 h-3.5 text-slate-400" />
          <span>{formatDate(article.published_at)}</span>
        </span>
        <span>&bull;</span>
        <span class="flex items-center gap-1.5">
          <Eye class="w-3.5 h-3.5 text-slate-400" />
          <span>{article.view_count} kali dibaca</span>
        </span>
      </div>
    {/if}
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto">
  {#if article}
    <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
      <!-- Main Content Column -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Featured Media -->
        <div class="ds-glass border border-slate-200/80 rounded-2xl overflow-hidden shadow-glass">
          <img 
            src={article.thumbnail} 
            alt={article.title} 
            class="w-full max-h-[460px] object-cover"
            onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
          />
          <div class="p-3.5 text-[11px] text-slate-500 italic bg-slate-50/80 border-t border-slate-200/80">
            Dokumentasi: Personel Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan
          </div>
        </div>

        <!-- Article Summary & Body Content -->
        <div class="ds-glass-strong border border-slate-200/80 rounded-2xl p-6 md:p-8 shadow-glass space-y-5">
          <!-- Lead Paragraph Callout -->
          <div class="bg-amber-50/80 border border-amber-200/60 border-l-4 border-l-[#F08300] p-4 sm:p-5 rounded-2xl shadow-2xs">
            <p class="font-semibold text-xs sm:text-sm text-slate-900 leading-relaxed">
              {article.summary}
            </p>
          </div>

          <div class="prose max-w-none text-xs sm:text-sm text-slate-700 leading-relaxed space-y-4">
            {@html article.content}
            <p>
              Pelaksanaan tugas penegakan Peraturan Daerah dan perlindungan masyarakat oleh jajaran Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan selalu didasarkan pada prinsip profesionalisme, akuntabilitas, serta kepatuhan pada regulasi hukum yang berlaku. Sinergi antara pemerintah daerah, aparat penegak hukum, dan partisipasi aktif masyarakat merupakan pilar utama terwujudnya ketertiban umum yang berkesinambungan.
            </p>
            <p>
              Masyarakat diimbau untuk turut menjaga fasilitas publik serta segera melaporkan setiap potensi gangguan ketenteraman umum melalui kanal pengaduan resmi SIAP TERTIB Satpol PP Prov. Sulsel.
            </p>
          </div>

          <!-- Back Button & Actions -->
          <div class="pt-6 border-t border-slate-200/80 flex flex-wrap justify-between items-center gap-3">
            <a 
              href="/berita" 
              class="inline-flex items-center gap-2 bg-slate-100 hover:bg-slate-200 text-slate-800 text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition"
            >
              <ArrowLeft class="w-3.5 h-3.5" />
              <span>Kembali ke Warta</span>
            </a>
            <a 
              href="/lapor" 
              class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition shadow-xs"
            >
              Lapor Pelanggaran Wilayah
            </a>
          </div>
        </div>
      </div>

      <!-- Sidebar Column -->
      <div class="space-y-6">
        <!-- SIAP TERTIB Quick CTA -->
        <div class="ds-glass-dark text-white p-6 rounded-2xl border border-white/10 shadow-glass space-y-3">
          <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">
            Kanal Pengaduan Warga
          </div>
          <h3 class="text-base font-bold text-white">Layanan SIAP TERTIB</h3>
          <p class="text-xs text-slate-300 leading-relaxed font-light">
            Temukan gangguan kenyamanan publik, perusakan fasilitas, atau penyerobotan aset di Sulawesi Selatan? Laporkan secara transparan.
          </p>
          <a 
            href="/lapor" 
            class="block text-center bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase py-2.5 rounded-xl transition shadow-xs mt-2"
          >
            Kirim Laporan Pengaduan
          </a>
        </div>

        <!-- Latest Articles List -->
        <div class="ds-glass border border-slate-200/80 rounded-2xl shadow-glass p-5 space-y-4">
          <h3 class="text-xs font-bold text-slate-900 uppercase tracking-wider border-b border-slate-200/80 pb-3">
            Warta Terkini Lainnya
          </h3>
          <div class="space-y-3.5">
            {#each recentArticles as rec}
              <a href={`/berita/${rec.slug}`} class="flex gap-3 group items-start">
                <img 
                  src={rec.thumbnail} 
                  alt={rec.title} 
                  class="w-16 h-16 object-cover flex-shrink-0 bg-slate-100 rounded-xl border border-slate-200/80 group-hover:scale-105 transition"
                  onerror="this.src='/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'"
                />
                <div class="flex-1 min-w-0">
                  <div class="text-[10px] text-slate-500 font-medium">{formatDate(rec.published_at)}</div>
                  <h4 class="text-xs font-bold text-slate-900 group-hover:text-[#F08300] line-clamp-2 leading-snug transition mt-0.5">
                    {rec.title}
                  </h4>
                </div>
              </a>
            {/each}
          </div>
        </div>

        <!-- Contact Hotline Info -->
        <div class="ds-glass border border-slate-200/80 rounded-2xl p-5 text-xs text-slate-700 shadow-glass space-y-2.5">
          <div class="font-bold uppercase text-[11px] text-slate-900 flex items-center gap-1.5">
            <Phone class="w-3.5 h-3.5 text-[#F08300]" />
            <span>Posko Siaga Satpol PP Sulsel</span>
          </div>
          <p class="text-[11px] text-slate-500 leading-relaxed font-light">Gedung Gabungan Dinas Pemprov Sulsel, Jl. Urip Sumoharjo No. 269, Makassar</p>
          <div class="pt-2 border-t border-slate-200/80 space-y-1 text-[11px]">
            <div>Telepon: <strong class="text-slate-900">0411-453105</strong></div>
            <div>Hotline Darurat: <strong class="text-[#F08300]">112 / 0811-4141-112</strong></div>
          </div>
        </div>
      </div>
    </div>
  {:else}
    <div class="ds-glass border border-slate-200/80 rounded-2xl p-12 text-center text-xs space-y-3 shadow-glass">
      <AlertCircle class="w-8 h-8 text-slate-400 mx-auto mb-2" />
      <div class="text-base font-bold text-slate-900">Artikel Tidak Ditemukan</div>
      <p class="text-slate-500">Artikel yang Anda cari tidak tersedia atau tautan telah dipindahkan.</p>
      <a href="/berita" class="inline-block bg-[#F08300] text-white text-xs font-bold uppercase px-4 py-2 mt-2 rounded-xl">
        Kembali ke Indeks Berita
      </a>
    </div>
  {/if}
</div>
