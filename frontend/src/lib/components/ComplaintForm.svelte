<script lang="ts">
  import { onMount } from 'svelte';
  import { defaultRegencies, complaintsStore } from '$lib/stores/appState';
  import type { Regency, ViolationType, Complaint } from '$lib/types';
  import { maskNIK, maskName, maskPhone } from '$lib/utils/masking';

  export let initialViolationType: ViolationType = 'GANGGUAN_TIBUM';

  let regencies: Regency[] = defaultRegencies;

  // Form Fields
  let nik = '';
  let name = '';
  let phone = '';
  let isAnonymous = false;
  let violationType: ViolationType = initialViolationType;
  let regencyId = 1; // Default to Makassar
  let districtName = '';
  let villageName = '';
  let addressDetail = '';
  // Default Sulsel coordinates (Poin 2: koordinat default tanpa popup izin GPS otomatis)
  let latitude = -5.147665;
  let longitude = 119.432732;
  let description = '';

  // Evidence Image State
  let evidenceImageFile: File | null = null;
  let evidenceImagePreview: string = '';
  let imageError = '';

  // Submission State
  let isSubmitting = false;
  let isLoading = false;
  let submitSuccess = false;
  let errorMessage = '';
  let createdTicket: Complaint | null = null;

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

  function handleImageUpload(e: Event) {
    const input = e.target as HTMLInputElement;
    imageError = '';

    if (input.files && input.files[0]) {
      const file = input.files[0];

      // Validate image format strictly
      const validTypes = ['image/jpeg', 'image/png', 'image/webp'];
      if (!validTypes.includes(file.type)) {
        imageError = 'Format file tidak didukung. Harap unggah foto bukti dalam format JPG, PNG, atau WEBP.';
        input.value = '';
        return;
      }

      // Max 5MB
      if (file.size > 5 * 1024 * 1024) {
        imageError = 'Ukuran foto terlalu besar. Maksimal ukuran berkas foto adalah 5 MB.';
        input.value = '';
        return;
      }

      evidenceImageFile = file;

      // Create preview
      const reader = new FileReader();
      reader.onload = (event) => {
        if (event.target?.result) {
          evidenceImagePreview = event.target.result as string;
        }
      };
      reader.readAsDataURL(file);
    }
  }

  function removeImage() {
    evidenceImageFile = null;
    evidenceImagePreview = '';
    imageError = '';
  }

  async function handleSubmit() {
    errorMessage = '';

    if (!nik || nik.length < 16) {
      errorMessage = 'NIK wajib diisi 16 digit sesuai identitas resmi KTP.';
      return;
    }

    if (!name.trim()) {
      errorMessage = 'Nama lengkap pelapor wajib diisi.';
      return;
    }

    if (!phone.trim()) {
      errorMessage = 'Nomor telepon / WhatsApp aktif wajib diisi untuk verifikasi laporan.';
      return;
    }

    if (!addressDetail.trim() || !description.trim()) {
      errorMessage = 'Alamat detail tempat kejadian dan uraian kronologi wajib diisi.';
      return;
    }

    isLoading = true;
    try {
      const selectedRegency = regencies.find(r => r.id === regencyId);

      // Perform masking if citizen chose anonymous mode
      const masked_nik = isAnonymous ? maskNIK(nik) : nik;
      const masked_name = isAnonymous ? maskName(name) : name;
      const masked_phone = isAnonymous ? maskPhone(phone) : phone;

      // Generate Ticket Code: TTB-YYYYMMDD-XXXX
      const datePart = new Date().toISOString().slice(0, 10).replace(/-/g, '');
      const randomPart = Math.floor(1000 + Math.random() * 9000);
      const generatedTicketCode = `TTB-${datePart}-${randomPart}`;

      const newComplaint: Complaint = {
        id: Date.now(),
        ticket_code: generatedTicketCode,
        reporter_nik: nik,
        reporter_name: name,
        reporter_phone: phone,
        is_anonymous: isAnonymous,
        masked_nik,
        masked_name,
        masked_phone,
        violation_type: violationType,
        regency_id: regencyId,
        regency_name: selectedRegency?.name || 'Provinsi Sulawesi Selatan',
        district_name: districtName,
        village_name: villageName,
        address_detail: addressDetail,
        latitude,
        longitude,
        description,
        evidence_image_url: evidenceImagePreview || '/uploads/bukti/default_evidence.jpg',
        status: 'MENUNGGU_VERIFIKASI',
        disposed_to: 'Bidang Ketenteraman dan Ketertiban Umum',
        created_at: new Date().toISOString()
      };

      // Push to app store
      complaintsStore.update(items => [newComplaint, ...items]);

      createdTicket = newComplaint;
      submitSuccess = true;

      // Scroll to result view
      if (typeof window !== 'undefined') {
        window.scrollTo({ top: 100, behavior: 'smooth' });
      }
    } catch (err: any) {
      errorMessage = err.message || 'Gagal mengirim pengaduan. Silakan periksa koneksi internet Anda.';
    } finally {
      isLoading = false;
    }
  }

  function resetForm() {
    nik = '';
    name = '';
    phone = '';
    isAnonymous = false;
    districtName = '';
    villageName = '';
    addressDetail = '';
    description = '';
    evidenceImageFile = null;
    evidenceImagePreview = '';
    submitSuccess = false;
    createdTicket = null;
    errorMessage = '';
  }
</script>

<div class="bg-white border border-[#eee9db] rounded-2xl shadow-sm p-6 md:p-8 max-w-4xl mx-auto">
  {#if submitSuccess && createdTicket}
    <!-- SUCCESS RECEIPT / TICKET VIEW -->
    <div class="space-y-6 text-center py-4">
      <div class="w-16 h-16 bg-emerald-100 text-emerald-600 rounded-full flex items-center justify-center mx-auto text-2xl font-bold">
        ✓
      </div>

      <div>
        <span class="bg-emerald-50 text-emerald-800 text-[11px] font-bold uppercase tracking-widest px-3 py-1 rounded-full border border-emerald-200">
          Laporan Berhasil Diterbitkan
        </span>
        <h3 class="text-2xl font-extrabold text-[#1a120c] mt-2">
          Bukti Registrasi Resi SIAP TERTIB
        </h3>
        <p class="text-xs text-[#6f7775] max-w-lg mx-auto mt-1">
          Laporan Anda telah masuk ke Sistem Operasional Terpadu Satpol PP Prov. Sulsel dan siap diverifikasi oleh petugas piket.
        </p>
      </div>

      <!-- Resi Card -->
      <div class="bg-[#faf7f2] border-2 border-[#e8dbc9] rounded-2xl p-6 text-left max-w-xl mx-auto space-y-4 shadow-xs">
        <div class="flex justify-between items-start border-b border-[#eee9db] pb-3">
          <div>
            <div class="text-[10px] text-stone-500 uppercase font-bold tracking-wider">Kode Resi Pengaduan Resmi</div>
            <div class="text-xl font-mono font-extrabold text-[#F08300] tracking-wide mt-0.5">
              {createdTicket.ticket_code}
            </div>
          </div>
          <span class="bg-amber-100 text-amber-900 border border-amber-300 text-[10px] font-bold uppercase px-2.5 py-1 rounded-md">
            Menunggu Verifikasi
          </span>
        </div>

        <div class="grid grid-cols-2 gap-3 text-xs">
          <div>
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Identitas Pelapor</div>
            <div class="font-bold text-[#1a120c]">
              {createdTicket.is_anonymous ? createdTicket.masked_name : createdTicket.reporter_name}
            </div>
            {#if createdTicket.is_anonymous}
              <div class="text-[10px] text-emerald-600 font-bold mt-0.5">🛡️ Mode Anonim Terlindungi</div>
            {/if}
          </div>

          <div>
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Waktu Masuk</div>
            <div class="font-bold text-[#1a120c]">
              {new Date(createdTicket.created_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric', hour: '2-digit', minute: '2-digit' })} WITA
            </div>
          </div>

          <div class="col-span-2">
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Klasifikasi Aduan</div>
            <div class="font-bold text-[#1a120c]">{createdTicket.violation_type}</div>
          </div>

          <div class="col-span-2">
            <div class="text-[10px] text-stone-500 uppercase font-semibold">Wilayah Tempat Kejadian</div>
            <div class="font-bold text-[#1a120c]">{createdTicket.regency_name}</div>
            <div class="text-stone-600 text-[11px] mt-0.5">{createdTicket.address_detail}</div>
          </div>
        </div>

        <div class="bg-amber-50 border-l-4 border-[#F08300] p-3 text-[11px] text-[#3d261a] rounded-r-lg">
          Simpan <strong>Kode Resi Tiket</strong> di atas untuk memantau status tindak lanjut petugas di lapangan melalui menu <strong>Lacak Aduan</strong>.
        </div>
      </div>

      <div class="flex justify-center gap-3 pt-2">
        <button
          type="button"
          on:click={resetForm}
          class="bg-[#3d261a] hover:bg-[#271810] text-white text-xs font-bold uppercase px-6 py-2.5 rounded-xl transition"
        >
          Kirim Laporan Baru
        </button>
        <a
          href="/lacak"
          class="bg-stone-100 hover:bg-stone-200 text-stone-800 text-xs font-bold uppercase px-6 py-2.5 rounded-xl transition"
        >
          Buka Pelacak Tiket
        </a>
      </div>
    </div>
  {:else}
    <!-- INPUT FORM -->
    <form on:submit|preventDefault={handleSubmit} class="space-y-6">
      <div class="border-b border-[#eee9db] pb-4">
        <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Aplikasi SIAP TERTIB</div>
        <h3 class="text-xl font-extrabold text-[#1a120c] tracking-tight">Formulir Laporan Pelanggaran Trantibum &amp; Perda</h3>
        <p class="text-xs text-[#6f7775] mt-1">
          Layanan pelaporan terintegrasi Satuan Polisi Pamong Praja Provinsi Sulawesi Selatan. Identitas Anda dijamin kerahasiaannya sesuai ketentuan undang-undang.
        </p>
      </div>

      {#if errorMessage}
        <div class="bg-rose-50 border border-rose-300 text-rose-800 text-xs p-4 rounded-xl font-semibold flex items-center gap-2">
          <svg class="w-4 h-4 shrink-0 text-rose-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span>{errorMessage}</span>
        </div>
      {/if}

      <!-- 1. Identitas Pelapor -->
      <div class="space-y-4">
        <div class="border-b border-stone-100 pb-2">
          <h4 class="font-extrabold text-sm uppercase text-[#1a120c] tracking-wide">1. Identitas Pelapor (Validasi KTP)</h4>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label for="nik-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">NIK (16 Digit KTP) *</label>
            <input 
              id="nik-input" 
              type="text" 
              bind:value={nik} 
              maxlength="16" 
              required
              placeholder="7371xxxxxxxxxxxx" 
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" 
            />
          </div>

          <div>
            <label for="name-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">Nama Lengkap *</label>
            <input 
              id="name-input" 
              type="text" 
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
              type="tel" 
              bind:value={phone} 
              required
              placeholder="08xxxxxxxxxx" 
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" 
            />
          </div>
        </div>

        <!-- Opsi Anonim Masking Identitas -->
        <div class="bg-[#faf7f2] border border-[#e8dbc9] p-3.5 rounded-xl flex items-start gap-3">
          <input 
            type="checkbox" 
            id="anon-check" 
            bind:checked={isAnonymous} 
            class="mt-1 h-4 w-4 text-[#F08300] focus:ring-[#F08300] border-stone-300 rounded" 
          />
          <label for="anon-check" class="text-xs text-[#1a120c] cursor-pointer">
            <span class="font-bold block">Lindungi Identitas Saya (Mode Anonim / Masking Otomatis)</span>
            <span class="text-[#6f7775] text-[11px] block mt-0.5">
              Nama dan NIK Anda akan disamarkan (*misal: M****** F*****) pada tampilan riwayat laporan publik. Petugas verifikator tetap dapat memverifikasi keabsahan KTP secara aman.
            </span>
          </label>
        </div>
      </div>

      <!-- 2. Klasifikasi Pelanggaran & Lokasi -->
      <div class="space-y-4 pt-4 border-t border-stone-100">
        <div class="border-b border-stone-100 pb-2">
          <h4 class="font-extrabold text-sm uppercase text-[#1a120c] tracking-wide">2. Klasifikasi Masalah &amp; Titik Lokasi</h4>
        </div>

        <!-- Violation Radio Options -->
        <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
          {#each violationOptions as opt}
            <label class={`border rounded-xl p-3.5 cursor-pointer transition flex items-start gap-3 ${violationType === opt.id ? 'border-[#F08300] bg-[#faf7f2] ring-1 ring-[#F08300]' : 'border-stone-200 hover:border-stone-300 bg-white'}`}>
              <input 
                type="radio" 
                name="violation_group" 
                value={opt.id} 
                bind:group={violationType} 
                class="mt-1 text-[#F08300] focus:ring-[#F08300]" 
              />
              <div class="text-xs">
                <div class="font-bold text-[#1a120c] leading-snug">{opt.label}</div>
                <div class="text-[11px] text-[#6f7775] mt-0.5 leading-relaxed">{opt.desc}</div>
              </div>
            </label>
          {/each}
        </div>

        <!-- Wilayah Dropdown Manual (Poin 2: Terstruktur manual tanpa deteksi GPS otomatis) -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label for="regency-select" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">Kabupaten / Kota *</label>
            <select 
              id="regency-select" 
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
              type="text" 
              bind:value={villageName} 
              placeholder="Contoh: Panaikang" 
              class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" 
            />
          </div>
        </div>

        <!-- Alamat Detail Tempat Kejadian (Manual Input) -->
        <div>
          <label for="address-textarea" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">
            Alamat Detail Tempat Kejadian *
          </label>
          <textarea 
            id="address-textarea" 
            bind:value={addressDetail} 
            rows="2" 
            required
            placeholder="Sebutkan nama jalan, patokan gedung, nomor bangunan, atau persimpangan tempat peristiwa terjadi..." 
            class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
          ></textarea>
        </div>
      </div>

      <!-- 3. Bukti Foto Kejadian & Uraian -->
      <div class="space-y-4 pt-4 border-t border-stone-100">
        <div class="border-b border-stone-100 pb-2">
          <h4 class="font-extrabold text-sm uppercase text-[#1a120c] tracking-wide">3. Bukti Gambar &amp; Uraian Kejadian</h4>
        </div>

        <!-- Image Upload Area -->
        <div>
          <label for="evidence-photo-input" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">
            Foto Bukti Kejadian (Format Gambar: JPG, PNG, WEBP)
          </label>
          
          <div class="border-2 border-dashed border-stone-300 rounded-xl p-5 text-center bg-stone-50/50 hover:bg-stone-50 transition">
            {#if evidenceImagePreview}
              <div class="space-y-3">
                <img 
                  src={evidenceImagePreview} 
                  alt="Bukti Aduan" 
                  class="max-h-56 mx-auto rounded-xl object-contain border border-stone-300 shadow-xs" 
                />
                <button 
                  type="button" 
                  on:click={removeImage} 
                  class="text-xs text-rose-600 font-bold hover:underline"
                >
                  ✕ Hapus &amp; Ganti Foto Bukti
                </button>
              </div>
            {:else}
              <input 
                id="evidence-photo-input" 
                type="file" 
                accept="image/*" 
                on:change={handleImageUpload} 
                class="hidden" 
              />
              <label for="evidence-photo-input" class="cursor-pointer block space-y-2">
                <div class="w-10 h-10 rounded-full bg-stone-100 text-[#8c6b45] flex items-center justify-center mx-auto">
                  📷
                </div>
                <div class="text-xs font-bold text-[#F08300]">Klik untuk unggah foto bukti lapangan</div>
                <p class="text-[11px] text-[#6f7775]">Foto situasi pelanggaran mempermudah percepatan tindak lanjut petugas di lapangan (Maks. 5MB)</p>
              </label>
            {/if}
          </div>

          {#if imageError}
            <div class="text-xs text-rose-600 font-semibold mt-1">{imageError}</div>
          {/if}
        </div>

        <div>
          <label for="desc-textarea" class="block font-bold uppercase text-[11px] text-[#1a120c] mb-1">
            Kronologi &amp; Uraian Pengaduan *
          </label>
          <textarea 
            id="desc-textarea" 
            bind:value={description} 
            rows="4" 
            required
            placeholder="Jelaskan secara ringkas situasi, perkiraan jumlah pihak terlibat, waktu kejadian, serta potensi bahaya terhadap ketertiban umum..." 
            class="w-full border border-stone-300 rounded-lg p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
          ></textarea>
        </div>
      </div>

      <!-- Submit Action Button -->
      <div class="pt-4 border-t border-stone-100 flex items-center justify-between">
        <span class="text-[11px] text-[#6f7775]">* Wajib diisi secara lengkap untuk keabsahan laporan</span>
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
