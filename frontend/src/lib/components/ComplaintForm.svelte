<script lang="ts">
  import { defaultRegencies } from '$lib/stores/appState';
  import { submitComplaint, uploadEvidence } from '$lib/api/client';
  import type { Complaint, EmailNotice, Regency, ViolationType } from '$lib/types';
  import { isValidEmail, maskNIK, maskName, maskPhone } from '$lib/utils/masking';
  import FileUploader from '$lib/components/ui/FileUploader.svelte';
  import StatusBadge from '$lib/components/ui/StatusBadge.svelte';

  export let initialViolationType: ViolationType = 'GANGGUAN_TIBUM';

  let regencies: Regency[] = defaultRegencies;

  // Form Fields
  let nik = '';
  let name = '';
  let phone = '';
  let email = '';
  let isAnonymous = false;
  let violationType: ViolationType = initialViolationType;
  let regencyId = '7301'; // Default: Kota Makassar
  let districtName = '';
  let villageName = '';
  let addressDetail = '';
  // Poin 2: Koordinat default tanpa popup izin GPS otomatis
  let latitude = -5.147665;
  let longitude = 119.432732;
  let description = '';

  // Evidence Upload State
  let evidenceFile: File | null = null;
  let evidencePreview = '';
  let evidenceFileName = '';

  // Submission State
  let isLoading = false;
  let submitSuccess = false;
  let errorMessage = '';
  let createdTicket: Complaint | null = null;
  let emailNotice: EmailNotice | null = null;

  const violationOptions: { id: ViolationType; label: string; desc: string }[] = [
    {
      id: 'GANGGUAN_TIBUM',
      label: 'Gangguan Ketenteraman & Ketertiban Umum',
      desc: 'Kebisingan berlebih, tawuran warga, balap liar, kerumunan tanpa izin, atau aktivitas meresahkan di ruang publik.'
    },
    {
      id: 'PELANGGARAN_PERDA',
      label: 'Pelanggaran Peraturan Daerah (Perda)',
      desc: 'Bangunan liar di sempadan sungai/jalan, tempat usaha tanpa izin resmi, reklame ilegal, perusakan fasilitas umum.'
    },
    {
      id: 'PENGAMBILDALIHAN_ASET',
      label: 'Penyerobotan / Penguasaan Aset Pemprov',
      desc: 'Lahan, gedung, atau fasilitas milik Pemprov Sulsel yang dipagari, dihuni sepihak, atau dialihfungsikan tanpa hak.'
    },
    {
      id: 'PERLINDUNGAN_MASYARAKAT',
      label: 'Layanan Linmas & Bencana Sosial',
      desc: 'Permohonan bantuan pengamanan pemukiman, evakuasi darurat ketertiban, dan mediasi konflik lingkungan.'
    }
  ];

  function handleFileChange(e: CustomEvent<{ url: string; file: File; name: string }>) {
    evidencePreview = e.detail.url;
    evidenceFile = e.detail.file;
    evidenceFileName = e.detail.name;
  }

  function handleFileClear() {
    evidencePreview = '';
    evidenceFile = null;
    evidenceFileName = '';
  }

  async function handleSubmit() {
    errorMessage = '';

    if (!nik || nik.trim().length < 16) {
      errorMessage = 'NIK wajib diisi 16 digit sesuai identitas resmi KTP.';
      return;
    }
    if (!name.trim()) {
      errorMessage = 'Nama lengkap pelapor wajib diisi.';
      return;
    }
    if (!phone.trim()) {
      errorMessage = 'Nomor telepon / WhatsApp aktif wajib diisi.';
      return;
    }
    if (!isValidEmail(email)) {
      errorMessage = 'Masukkan alamat email aktif yang valid (contoh: nama@domain.com) untuk menerima resi aduan.';
      return;
    }
    if (!addressDetail.trim() || !description.trim()) {
      errorMessage = 'Alamat detail tempat kejadian dan uraian kronologi wajib diisi.';
      return;
    }

    isLoading = true;
    try {
      const selectedRegency = regencies.find(r => r.id === regencyId);
      const selectedViolation = violationOptions.find(v => v.id === violationType);

      let finalEvidenceUrl = evidencePreview;
      if (evidenceFile) {
        const uploaded = await uploadEvidence(evidenceFile);
        if (uploaded) finalEvidenceUrl = uploaded;
      }

      const payload: Partial<Complaint> = {
        is_anonymous: isAnonymous,
        reporter_nik: nik.trim(),
        reporter_name: name.trim(),
        reporter_phone: phone.trim(),
        reporter_email: email.trim(),
        masked_nik: maskNIK(nik),
        masked_name: isAnonymous ? 'Warga (Anonim Terlindungi)' : maskName(name),
        masked_phone: maskPhone(phone),
        violation_type: selectedViolation?.label || violationType,
        regency_id: regencyId,
        regency_name: selectedRegency?.name || 'Kota Makassar',
        district_name: districtName.trim(),
        village_name: villageName.trim(),
        address_detail: addressDetail.trim(),
        latitude,
        longitude,
        description: description.trim(),
        evidence_image_url: finalEvidenceUrl || '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg'
      };

      const res = await submitComplaint(payload);
      createdTicket = res.data ?? ({
        id: Date.now(),
        ticket_code: res.ticket_code,
        ...payload,
        status: 'MENUNGGU_VERIFIKASI',
        created_at: new Date().toISOString()
      } as Complaint);
      emailNotice = res.email;
      submitSuccess = true;

      if (typeof window !== 'undefined') {
        window.scrollTo({ top: 100, behavior: 'smooth' });
      }
    } catch (err: any) {
      errorMessage = err?.message || 'Gagal mengirim pengaduan. Silakan coba beberapa saat lagi.';
    } finally {
      isLoading = false;
    }
  }

  function resetForm() {
    nik = '';
    name = '';
    phone = '';
    email = '';
    isAnonymous = false;
    regencyId = '7301';
    districtName = '';
    villageName = '';
    addressDetail = '';
    description = '';
    evidenceFile = null;
    evidencePreview = '';
    evidenceFileName = '';
    submitSuccess = false;
    createdTicket = null;
    emailNotice = null;
    errorMessage = '';
  }
</script>

<div class="bg-white border border-[#eee9db] rounded-2xl shadow-sm p-6 md:p-8 max-w-4xl mx-auto">
  {#if submitSuccess && createdTicket}
    <!-- SUCCESS RECEIPT / TICKET VIEW -->
    <div class="space-y-6 text-center py-4" role="status" aria-live="polite">
      <div class="w-16 h-16 bg-emerald-100 text-emerald-700 rounded-full flex items-center justify-center mx-auto text-2xl font-bold">
        &#10003;
      </div>

      <div>
        <span class="bg-emerald-50 text-emerald-800 text-[11px] font-bold uppercase tracking-widest px-3 py-1 rounded-full border border-emerald-200">
          Laporan Berhasil Diterbitkan
        </span>
        <h3 class="text-2xl font-extrabold text-[#1a120c] mt-2">
          Bukti Registrasi Resi SIAP TERTIB
        </h3>
        <p class="text-xs text-stone-500 max-w-lg mx-auto mt-1">
          Laporan Anda telah masuk ke Sistem Operasional Terpadu Satpol PP Prov. Sulsel dan siap diverifikasi oleh petugas piket.
        </p>
      </div>

      <!-- Resi Card -->
      <div class="bg-[#faf7f2] border-2 border-[#e8dbc9] rounded-2xl p-6 text-left max-w-xl mx-auto space-y-4 shadow-xs">
        <div class="flex flex-wrap justify-between items-start gap-2 border-b border-[#eee9db] pb-3">
          <div>
            <div class="text-[10px] text-stone-500 uppercase font-bold tracking-wider">Kode Resi Pengaduan Resmi</div>
            <div class="text-xl font-mono font-extrabold text-[#F08300] tracking-wide mt-0.5">
              {createdTicket.ticket_code}
            </div>
          </div>
          <StatusBadge status={createdTicket.status} size="md" />
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 gap-3 text-xs">
          <div>
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Identitas Pelapor</div>
            <div class="font-bold text-[#1a120c]">
              {createdTicket.is_anonymous ? 'Warga (Anonim Terlindungi)' : createdTicket.reporter_name}
            </div>
            {#if createdTicket.is_anonymous}
              <div class="text-[10px] text-emerald-700 font-bold mt-0.5">Mode Anonim &amp; Masking Aktif</div>
            {/if}
          </div>

          <div>
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Waktu Masuk</div>
            <div class="font-bold text-[#1a120c]">
              {new Date(createdTicket.created_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' })} WITA
            </div>
          </div>

          <div class="sm:col-span-2">
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Email Penerima Notifikasi</div>
            <div class="font-bold text-[#1a120c] font-mono">{createdTicket.reporter_email || email}</div>
          </div>

          <div class="sm:col-span-2">
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Klasifikasi Aduan</div>
            <div class="font-bold text-[#1a120c]">{createdTicket.violation_type}</div>
          </div>

          <div class="sm:col-span-2">
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Wilayah Tempat Kejadian</div>
            <div class="font-bold text-[#1a120c]">{createdTicket.regency_name}</div>
            <div class="text-stone-600 text-[11px] mt-0.5">{createdTicket.address_detail}</div>
          </div>
        </div>

        <!-- Email Notification Status / Preview -->
        {#if emailNotice}
          <div class="bg-white border border-[#e8dbc9] rounded-xl p-3.5 text-[11px] space-y-1">
            <div class="flex items-center justify-between gap-2">
              <span class="font-bold text-[#3d261a] uppercase tracking-wider">
                {emailNotice.mode === 'smtp' ? 'Notifikasi Email Terkirim' : 'Simulasi Notifikasi Email'}
              </span>
              <span class={`text-[10px] font-bold px-2 py-0.5 rounded-full ${emailNotice.sent ? 'bg-emerald-100 text-emerald-800' : 'bg-amber-100 text-amber-800'}`}>
                {emailNotice.sent ? 'SMTP OK' : 'Preview Lokal'}
              </span>
            </div>
            <div class="text-stone-600">Kepada: <strong class="font-mono text-[#1a120c]">{emailNotice.to}</strong></div>
            <div class="text-stone-600">Subjek: <span class="text-[#1a120c] font-medium">{emailNotice.subject}</span></div>
            <p class="text-stone-500 italic">{emailNotice.message}</p>
          </div>
        {/if}

        <div class="bg-amber-50 border-l-4 border-[#F08300] p-3 text-[11px] text-[#3d261a] rounded-r-lg">
          Simpan <strong>Kode Resi Tiket</strong> di atas untuk memantau perkembangan penanganan melalui menu <strong>Lacak Aduan</strong>.
        </div>
      </div>

      <div class="flex flex-wrap justify-center gap-3 pt-2">
        <button
          type="button"
          on:click={resetForm}
          class="bg-[#3d261a] hover:bg-[#271810] text-white text-xs font-bold uppercase px-6 py-2.5 rounded-xl transition"
        >
          Kirim Laporan Baru
        </button>
        <a
          href={`/lacak?ticket=${encodeURIComponent(createdTicket.ticket_code)}`}
          class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-6 py-2.5 rounded-xl transition shadow-xs"
        >
          Lacak Tiket Sekarang &rarr;
        </a>
      </div>
    </div>
  {:else}
    <!-- INPUT FORM -->
    <form on:submit|preventDefault={handleSubmit} class="space-y-6" aria-labelledby="siap-tertib-form-heading">
      <div class="border-b border-[#eee9db] pb-4">
        <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Aplikasi SIAP TERTIB</div>
        <h3 id="siap-tertib-form-heading" class="text-xl font-extrabold text-[#1a120c] tracking-tight">Formulir Laporan Pelanggaran Trantibum &amp; Perda</h3>
        <p class="text-xs text-stone-500 mt-1">
          Layanan pelaporan terintegrasi Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan. Identitas pelapor dilindungi pada tampilan publik.
        </p>
      </div>

      {#if errorMessage}
        <div role="alert" class="bg-rose-50 border border-rose-300 text-rose-800 text-xs p-4 rounded-xl font-semibold flex items-center gap-2">
          <svg class="w-4 h-4 shrink-0 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span>{errorMessage}</span>
        </div>
      {/if}

      <!-- 1. Identitas Pelapor -->
      <fieldset class="space-y-4">
        <legend class="font-extrabold text-sm uppercase text-[#1a120c] tracking-wide border-b border-stone-100 pb-2 w-full">
          1. Identitas Pelapor (Validasi KTP &amp; Email)
        </legend>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label for="nik-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">NIK (16 Digit KTP) *</label>
            <input
              id="nik-input"
              name="reporter_nik"
              type="text"
              inputmode="numeric"
              bind:value={nik}
              maxlength="16"
              minlength="16"
              pattern="\d{16}"
              required
              autocorrect="off"
              autocapitalize="none"
              spellcheck="false"
              placeholder="7371xxxxxxxxxxxx"
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
          </div>

          <div>
            <label for="name-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">Nama Lengkap *</label>
            <input
              id="name-input"
              name="reporter_name"
              type="text"
              autocomplete="name"
              bind:value={name}
              required
              placeholder="Sesuai identitas KTP"
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
          </div>

          <div>
            <label for="phone-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">No. WhatsApp / Telepon *</label>
            <input
              id="phone-input"
              name="reporter_phone"
              type="tel"
              autocomplete="tel"
              bind:value={phone}
              required
              placeholder="08xxxxxxxxxx"
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
          </div>

          <div>
            <label for="email-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-0.5">Email Pelapor (Notifikasi Resi) *</label>
            <span id="email-hint" class="block text-[10px] text-stone-500 mb-1">Kode tiket dan pembaruan progres akan dikirim ke alamat ini (format: nama@domain.com).</span>
            <input
              id="email-input"
              name="reporter_email"
              type="email"
              autocomplete="email"
              autocorrect="off"
              autocapitalize="none"
              spellcheck="false"
              aria-describedby="email-hint"
              bind:value={email}
              required
              placeholder="pelapor@email.com"
              class="peer w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
            <div class="hidden peer-[:user-invalid]:block text-[11px] text-rose-700 font-semibold mt-1">
              Masukkan alamat email valid dalam format nama@domain.com.
            </div>
          </div>
        </div>

        <!-- Opsi Anonim Masking Identitas -->
        <div class="bg-[#faf7f2] border border-[#e8dbc9] p-3.5 rounded-xl flex items-start gap-3">
          <input
            type="checkbox"
            id="anon-check"
            bind:checked={isAnonymous}
            class="mt-1 h-4 w-4 text-[#F08300] focus:ring-[#F08300] border-stone-300 rounded accent-[#F08300]"
          />
          <label for="anon-check" class="text-xs text-[#1a120c] cursor-pointer">
            <span class="font-bold block">Sembunyikan Nama Saya Sepenuhnya (Mode Anonim)</span>
            <span class="text-stone-500 text-[11px] block mt-0.5">
              Seluruh pelacakan tiket publik sudah menyamarkan identitas pelapor otomatis. Centang opsi ini jika ingin label tampil sebagai &ldquo;Warga (Anonim Terlindungi)&rdquo;.
            </span>
          </label>
        </div>
      </fieldset>

      <!-- 2. Klasifikasi Pelanggaran & Lokasi -->
      <fieldset class="space-y-4 pt-4 border-t border-stone-100">
        <legend class="font-extrabold text-sm uppercase text-[#1a120c] tracking-wide border-b border-stone-100 pb-2 w-full">
          2. Klasifikasi Masalah &amp; Titik Lokasi
        </legend>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
          {#each violationOptions as opt}
            <label class={`border rounded-xl p-3.5 cursor-pointer transition flex items-start gap-3 ${violationType === opt.id ? 'border-[#F08300] bg-[#faf7f2] ring-1 ring-[#F08300]' : 'border-stone-200 hover:border-stone-300 bg-white'}`}>
              <input
                type="radio"
                name="violation_group"
                value={opt.id}
                bind:group={violationType}
                class="mt-1 text-[#F08300] focus:ring-[#F08300] accent-[#F08300]"
              />
              <div class="text-xs">
                <div class="font-bold text-[#1a120c] leading-snug">{opt.label}</div>
                <div class="text-[11px] text-stone-500 mt-0.5 leading-relaxed">{opt.desc}</div>
              </div>
            </label>
          {/each}
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label for="regency-select" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">Kabupaten / Kota *</label>
            <select
              id="regency-select"
              name="regency_id"
              bind:value={regencyId}
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs bg-white text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            >
              {#each regencies as reg}
                <option value={reg.id}>{reg.name}</option>
              {/each}
            </select>
          </div>

          <div>
            <label for="district-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">Kecamatan</label>
            <input
              id="district-input"
              name="district_name"
              type="text"
              bind:value={districtName}
              placeholder="Contoh: Panakkukang"
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
          </div>

          <div>
            <label for="village-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">Kelurahan / Desa</label>
            <input
              id="village-input"
              name="village_name"
              type="text"
              bind:value={villageName}
              placeholder="Contoh: Panaikang"
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
          </div>
        </div>

        <div>
          <label for="address-textarea" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">
            Alamat Detail Tempat Kejadian *
          </label>
          <textarea
            id="address-textarea"
            name="address_detail"
            autocomplete="street-address"
            bind:value={addressDetail}
            rows="2"
            required
            placeholder="Sebutkan nama jalan, patokan gedung, nomor bangunan, atau persimpangan tempat peristiwa terjadi..."
            class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
          ></textarea>
        </div>
      </fieldset>

      <!-- 3. Bukti Foto & Uraian (Reusable FileUploader) -->
      <fieldset class="space-y-4 pt-4 border-t border-stone-100">
        <legend class="font-extrabold text-sm uppercase text-[#1a120c] tracking-wide border-b border-stone-100 pb-2 w-full">
          3. Bukti Gambar &amp; Uraian Kejadian
        </legend>

        <FileUploader
          value={evidencePreview}
          fileName={evidenceFileName}
          label="Foto Bukti Kejadian (JPG, PNG, WEBP)"
          description="Seret &amp; lepas foto situasi lapangan di sini, atau klik untuk memilih berkas (Maks. 5MB)"
          on:change={handleFileChange}
          on:clear={handleFileClear}
        />

        <div>
          <label for="desc-textarea" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">
            Kronologi &amp; Uraian Pengaduan *
          </label>
          <textarea
            id="desc-textarea"
            name="description"
            bind:value={description}
            rows="4"
            required
            placeholder="Jelaskan secara ringkas situasi, perkiraan jumlah pihak terlibat, waktu kejadian, serta potensi bahaya terhadap ketertiban umum..."
            class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
          ></textarea>
        </div>
      </fieldset>

      <div class="pt-4 border-t border-stone-100 flex flex-wrap items-center justify-between gap-3">
        <span class="text-[11px] text-stone-500">* Wajib diisi secara lengkap untuk keabsahan laporan</span>
        <button
          type="submit"
          disabled={isLoading}
          class="bg-[#3d261a] hover:bg-[#271810] disabled:opacity-50 text-white font-bold text-xs uppercase px-8 py-3 rounded-xl transition shadow-md"
        >
          {#if isLoading}
            Mengirim Laporan...
          {:else}
            Terbitkan Laporan &rarr;
          {/if}
        </button>
      </div>
    </form>
  {/if}
</div>
