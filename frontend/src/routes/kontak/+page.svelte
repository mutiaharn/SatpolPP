<script lang="ts">
  let senderName = '';
  let senderEmail = '';
  let senderPhone = '';
  let subject = '';
  let message = '';
  let isSending = false;
  let sentSuccess = false;

  function handleSubmit() {
    if (!senderName || !senderEmail || !message) {
      alert('Nama, email, dan pesan wajib diisi.');
      return;
    }
    isSending = true;
    setTimeout(() => {
      isSending = false;
      sentSuccess = true;
      senderName = '';
      senderEmail = '';
      senderPhone = '';
      subject = '';
      message = '';
      setTimeout(() => sentSuccess = false, 5000);
    }, 600);
  }
</script>

<svelte:head>
  <title>Kontak &amp; Informasi Layanan | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="bg-[#f9f4e8] border-b border-[#eee9db] py-10 px-4">
  <div class="max-w-6xl mx-auto">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-widest">Pusat Informasi &amp; Koordinasi</div>
    <h1 class="text-3xl font-extrabold text-[#101116] tracking-tight mt-1">Kontak &amp; Alamat Resmi</h1>
    <div class="w-12 h-1 bg-[#F08300] mt-3"></div>
    <p class="text-xs text-[#6f7775] mt-2 max-w-2xl">
      Hubungi markas komando Satpol PP Provinsi Sulawesi Selatan untuk keperluan dinas, koordinasi kedinasan, atau tindak lanjut aduan ketertiban umum.
    </p>
  </div>
</div>

<div class="py-10 px-4 max-w-6xl mx-auto space-y-10">
  <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
    <!-- Contact Info Cards -->
    <div class="space-y-4">
      <div class="bg-white border border-[#eee9db] p-6 shadow-sm space-y-3">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Markas Komando</div>
        <h3 class="text-base font-extrabold text-[#101116]">Kantor Satpol PP Prov. Sulsel</h3>
        <p class="text-xs text-[#6f7775] leading-relaxed">
          Gedung Gabungan Dinas Pemprov Sulsel, Jl. Urip Sumoharjo No. 269, Kel. Karampuang, Kec. Panakkukang, Kota Makassar, Sulawesi Selatan 90231
        </p>
      </div>

      <div class="bg-white border border-[#eee9db] p-6 shadow-sm space-y-3">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Telepon &amp; Surel</div>
        <div class="space-y-2 text-xs text-[#2b2f3f]">
          <div>
            <div class="text-[10px] text-[#6f7775] uppercase font-bold">Telepon Kantor / Sekretariat:</div>
            <div class="font-bold text-[#101116] mt-0.5">0411-453105</div>
          </div>
          <div>
            <div class="text-[10px] text-[#6f7775] uppercase font-bold">Email Resmi:</div>
            <div class="font-bold text-[#101116] mt-0.5">satpol@sulselprov.go.id</div>
          </div>
          <div>
            <div class="text-[10px] text-[#6f7775] uppercase font-bold">Hotline Darurat / Cepat (SIAP TERTIB):</div>
            <div class="font-bold text-[#F08300] mt-0.5">112 / 0811-4141-112</div>
          </div>
        </div>
      </div>

      <div class="bg-[#2b2f3f] text-white p-6 shadow-sm border-l-4 border-[#F08300] space-y-3">
        <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Jam Operasional Pelayanan</div>
        <div class="space-y-2 text-xs text-gray-300">
          <div>
            <div class="text-white font-bold">Administrasi Kantor &amp; Surat:</div>
            <div class="text-[11px] text-gray-400">Senin &ndash; Jumat: 07.30 &ndash; 16.00 WITA</div>
          </div>
          <div class="pt-2 border-t border-gray-700">
            <div class="text-white font-bold">Posko Siaga Patroli &amp; Tindak Cepat:</div>
            <div class="text-[11px] text-emerald-400 font-semibold">24 Jam Non-Stop (Setiap Hari)</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Contact Form & Map Column -->
    <div class="lg:col-span-2 space-y-6">
      <div class="bg-white border border-[#eee9db] p-6 md:p-8 shadow-sm space-y-6">
        <div>
          <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Aspirasi &amp; Pertanyaan</div>
          <h2 class="text-lg font-bold text-[#101116] tracking-tight mt-1">Formulir Kirim Pesan / Koordinasi Dinas</h2>
          <p class="text-xs text-[#6f7775] mt-1">
            Gunakan formulir ini untuk korespondensi formal atau permohonan informasi. Untuk pelaporan pelanggaran Perda lapangan, disarankan menggunakan menu <a href="/lapor" class="text-[#F08300] font-bold hover:underline">SIAP TERTIB</a> agar mendapatkan nomor tiket registrasi.
          </p>
        </div>

        {#if sentSuccess}
          <div class="bg-emerald-50 border border-emerald-300 text-emerald-800 text-xs p-4 font-semibold">
            Pesan Anda telah berhasil dikirimkan ke sekretariat Satpol PP Prov. Sulsel. Petugas kami akan membalas melalui email resmi.
          </div>
        {/if}

        <form on:submit|preventDefault={handleSubmit} class="space-y-4">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label for="contact-name" class="block text-xs font-bold uppercase text-[#101116] mb-1">Nama Lengkap</label>
              <input id="contact-name" type="text" bind:value={senderName} required class="w-full border border-gray-300 p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" />
            </div>
            <div>
              <label for="contact-email" class="block text-xs font-bold uppercase text-[#101116] mb-1">Email Anda</label>
              <input id="contact-email" type="email" bind:value={senderEmail} required class="w-full border border-gray-300 p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" />
            </div>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
            <div>
              <label for="contact-phone" class="block text-xs font-bold uppercase text-[#101116] mb-1">Nomor Telepon / WhatsApp</label>
              <input id="contact-phone" type="tel" bind:value={senderPhone} class="w-full border border-gray-300 p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" />
            </div>
            <div>
              <label for="contact-subject" class="block text-xs font-bold uppercase text-[#101116] mb-1">Perihal Pesan</label>
              <input id="contact-subject" type="text" bind:value={subject} placeholder="Contoh: Koordinasi Penertiban Wilayah" class="w-full border border-gray-300 p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" />
            </div>
          </div>

          <div>
            <label for="contact-message" class="block text-xs font-bold uppercase text-[#101116] mb-1">Isi Pesan / Keterangan Lengkap</label>
            <textarea id="contact-message" bind:value={message} rows="4" required class="w-full border border-gray-300 p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none"></textarea>
          </div>

          <div class="flex justify-end pt-2">
            <button
              type="submit"
              disabled={isSending}
              class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-6 py-2.5 transition disabled:opacity-50"
            >
              {#if isSending}Mengirimkan...{:else}Kirim Pesan Resmi{/if}
            </button>
          </div>
        </form>
      </div>

      <!-- Map Embed Card -->
      <div class="bg-white border border-[#eee9db] shadow-sm overflow-hidden">
        <div class="p-4 border-b border-gray-100 flex justify-between items-center text-xs">
          <span class="font-bold text-[#101116] uppercase">Lokasi Markas Komando (Peta Koordinat)</span>
          <span class="text-[#6f7775] font-mono text-[11px]">-5.1432&deg; S, 119.4521&deg; E</span>
        </div>
        <div class="h-64 bg-gray-100 relative">
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
