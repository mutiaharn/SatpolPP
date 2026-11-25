<script lang="ts">
  import { tick } from 'svelte';
  import { ChevronRight, Phone, Mail, MapPin, Clock, Send, ShieldAlert, CheckCircle2 } from 'lucide-svelte';
  import { Button, Input } from '$lib/components/ui';

  let senderName = '';
  let senderEmail = '';
  let senderPhone = '';
  let subject = '';
  let message = '';
  let isSending = false;
  let sentSuccess = false;

  async function handleSubmit() {
    if (!senderName || !senderEmail || !message) {
      alert('Nama, email, dan pesan wajib diisi.');
      return;
    }
    isSending = true;
    setTimeout(async () => {
      isSending = false;
      sentSuccess = true;
      senderName = '';
      senderEmail = '';
      senderPhone = '';
      subject = '';
      message = '';
      await tick();
      const alertEl = document.getElementById('contact-sent-alert');
      if (alertEl) {
        alertEl.scrollIntoView({ behavior: 'smooth', block: 'nearest' });
      }
      setTimeout(() => sentSuccess = false, 5000);
    }, 600);
  }
</script>

<svelte:head>
  <title>Kontak &amp; Informasi Layanan | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Sovereign Slim Hero Banner -->
<div class="bg-gradient-to-b from-[#070e1c] via-[#0b172a] to-[#0f223d] text-white py-6 sm:py-7 md:py-8 px-4 relative overflow-hidden ds-specular-dark border-b border-white/10">
  <div class="absolute top-0 right-10 w-72 h-72 rounded-full bg-[#F08300]/5 blur-3xl pointer-events-none"></div>
  <div class="absolute bottom-0 left-10 w-56 h-56 rounded-full bg-[#F08300]/8 blur-3xl pointer-events-none"></div>

  <div class="max-w-6xl mx-auto relative z-10">
    <div class="flex flex-wrap items-center gap-1.5 text-xs text-slate-400 mb-2">
      <a href="/" class="hover:text-[#F08300] transition">Beranda</a>
      <ChevronRight class="w-3.5 h-3.5 text-slate-500" />
      <span class="text-[#F08300] font-medium">Kontak &amp; Layanan</span>
    </div>

    <h1 class="text-2xl sm:text-3xl font-extrabold text-white tracking-tight">Kontak &amp; Alamat Resmi</h1>
    <p class="text-xs sm:text-sm text-slate-300 mt-1.5 max-w-2xl font-light leading-relaxed">
      Hubungi posko pelayanan terpadu atau kunjungi Markas Komando Satpol PP Provinsi Sulawesi Selatan untuk konsultasi dan koordinasi langsung.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-10">
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
    <!-- Contact Info Cards -->
    <div class="space-y-5">
      <div class="ds-glass border border-white/90 p-6 shadow-glass space-y-3 rounded-3xl ds-specular-top">
        <div class="flex items-center gap-2 text-[#F08300] font-bold text-xs uppercase tracking-wider">
          <MapPin class="w-4 h-4 text-[#F08300]" />
          <span>Markas Komando</span>
        </div>
        <h3 class="text-base font-bold text-slate-900">Kantor Satpol PP Prov. Sulsel</h3>
        <p class="text-xs sm:text-sm text-slate-600 leading-relaxed font-light">
          Gedung Gabungan Dinas Pemprov Sulsel, Jl. Urip Sumoharjo No. 269, Kel. Karampuang, Kec. Panakkukang, Kota Makassar, Sulawesi Selatan 90231
        </p>
      </div>

      <div class="ds-glass border border-white/90 p-6 shadow-glass space-y-4 rounded-3xl ds-specular-top">
        <div class="flex items-center gap-2 text-[#F08300] font-bold text-xs uppercase tracking-wider">
          <Phone class="w-4 h-4 text-[#F08300]" />
          <span>Telepon &amp; Surel Resmi</span>
        </div>
        <div class="space-y-3 text-xs">
          <div class="ds-glass-subtle p-3 rounded-xl border border-white/80">
            <div class="text-[10px] text-slate-500 uppercase font-bold">Telepon Kantor / Sekretariat:</div>
            <div class="font-bold text-slate-900 text-sm mt-0.5">0411-453105</div>
          </div>
          <div class="ds-glass-subtle p-3 rounded-xl border border-white/80">
            <div class="text-[10px] text-slate-500 uppercase font-bold">Email Resmi Instansi:</div>
            <div class="font-bold text-slate-900 text-sm mt-0.5">satpol@sulselprov.go.id</div>
          </div>
          <div class="ds-glass-gold p-3 rounded-xl border border-orange-300/60">
            <div class="text-[10px] text-[#F08300] uppercase font-bold">Hotline Darurat Cepat (SIAP TERTIB):</div>
            <div class="font-bold text-[#F08300] text-sm mt-0.5">112 / 0811-4141-112</div>
          </div>
        </div>
      </div>

      <div class="ds-glass-dark text-white p-6 rounded-3xl shadow-glass border border-white/15 border-l-4 border-l-[#F08300] space-y-3 ds-specular-dark">
        <div class="flex items-center gap-2 text-[#F08300] font-bold text-xs uppercase tracking-wider">
          <Clock class="w-4 h-4 text-[#F08300]" />
          <span>Jam Operasional Pelayanan</span>
        </div>
        <div class="space-y-2.5 text-xs">
          <div>
            <div class="text-white font-bold">Administrasi Kantor &amp; Persuratan:</div>
            <div class="text-[11px] text-slate-300 mt-0.5 font-light">Senin &ndash; Jumat: 07.30 &ndash; 16.00 WITA</div>
          </div>
          <div class="pt-2 border-t border-white/10">
            <div class="text-white font-bold">Posko Siaga Patroli &amp; Reaksi Cepat:</div>
            <div class="text-[11px] text-emerald-400 font-semibold mt-0.5">
              24 Jam Non-Stop (Setiap Hari Termasuk Libur)
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Contact Form & Map Column -->
    <div class="lg:col-span-2 space-y-6">
      <div class="ds-glass-strong border border-white/90 p-6 md:p-8 shadow-glass space-y-6 rounded-3xl ds-specular-top">
        <div class="border-b border-white/80 pb-4">
          <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Aspirasi &amp; Pertanyaan</div>
          <h2 class="text-xl font-bold text-slate-900 tracking-tight mt-1">Formulir Kirim Pesan / Koordinasi Dinas</h2>
          <p class="text-xs text-slate-600 mt-1.5 leading-relaxed font-light">
            Gunakan formulir ini untuk korespondensi formal atau permohonan informasi. Untuk pelaporan pelanggaran Perda lapangan, disarankan menggunakan menu <a href="/lapor" class="text-[#F08300] font-bold hover:underline">SIAP TERTIB</a> agar mendapatkan nomor tiket registrasi.
          </p>
        </div>

        {#if sentSuccess}
          <div id="contact-sent-alert" class="bg-emerald-50/90 border border-emerald-300 text-emerald-800 text-xs p-4 font-semibold rounded-2xl flex items-center gap-2 backdrop-blur-md">
            <CheckCircle2 class="w-4 h-4 text-emerald-600 shrink-0" />
            <span>Pesan Anda telah berhasil dikirimkan ke sekretariat Satpol PP Prov. Sulsel. Petugas kami akan membalas melalui email resmi.</span>
          </div>
        {/if}

        <form on:submit|preventDefault={handleSubmit} class="space-y-4">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <Input
              id="contact-name"
              label="Nama Lengkap"
              bind:value={senderName}
              required
              placeholder="Nama Anda"
            />
            <Input
              id="contact-email"
              type="email"
              label="Email Anda"
              bind:value={senderEmail}
              required
              placeholder="nama@email.com"
            />
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <Input
              id="contact-phone"
              type="tel"
              label="Nomor Telepon / WhatsApp"
              bind:value={senderPhone}
              placeholder="08xxxxxxxxxx"
            />
            <Input
              id="contact-subject"
              label="Perihal Pesan"
              bind:value={subject}
              placeholder="Contoh: Koordinasi Penertiban Wilayah"
            />
          </div>

          <Input
            id="contact-message"
            multiline
            rows={4}
            label="Isi Pesan / Keterangan Lengkap"
            bind:value={message}
            required
            placeholder="Tuliskan keterangan lengkap atau permohonan koordinasi..."
          />

          <div class="flex justify-end pt-2">
            <Button
              type="submit"
              loading={isSending}
              disabled={isSending}
              variant="gold"
              size="md"
              className="uppercase tracking-wider px-6 font-bold"
            >
              <Send class="w-3.5 h-3.5" />
              <span>{isSending ? 'Mengirimkan...' : 'Kirim Pesan Resmi'}</span>
            </Button>
          </div>
        </form>
      </div>

      <!-- Map Embed Card -->
      <div class="ds-glass border border-white/90 shadow-glass overflow-hidden rounded-3xl ds-specular-top">
        <div class="p-4 border-b border-white/80 bg-white/40 backdrop-blur-md flex justify-between items-center text-xs">
          <span class="font-bold text-slate-900 uppercase flex items-center gap-1.5">
            <MapPin class="w-3.5 h-3.5 text-[#F08300]" />
            <span>Lokasi Markas Komando (Peta Koordinat)</span>
          </span>
          <span class="text-slate-500 font-mono text-[11px]">-5.1432&deg; S, 119.4521&deg; E</span>
        </div>
        <div class="h-64 bg-slate-100 relative">
          <iframe 
            title="Lokasi Kantor Satpol PP Prov Sulsel"
            src="https://maps.google.com/maps?q=-5.1432,119.4521&hl=id&z=15&output=embed"
            class="w-full h-full border-0"
            loading="lazy"
          ></iframe>
        </div>
      </div>
    </div>
  </div>
</div>
