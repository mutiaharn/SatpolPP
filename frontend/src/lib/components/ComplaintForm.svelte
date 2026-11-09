<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchRegencies, submitComplaint } from '$lib/api/client';
  import { maskNIK, maskName, maskPhone } from '$lib/utils/masking';
  import type { Regency } from '$lib/types';

  let regencies: Regency[] = [];
  let isAnonymous = false;
  
  // Real citizen data (Always collected for verification)
  let nik = '';
  let name = '';
  let phone = '';
  let email = '';

  // Authentic 9 Violation Categories from SIAP TERTIB Sulsel
  const violationTypes = [
    'Jalan (Trotoar, PKL, Rumija, Fasilitas Lalu Lintas)',
    'Sungai, Saluran, Danau dan Perairan Pesisir',
    'Lingkungan (Sampah Liar, Polusi, Ketenteraman Lingkungan)',
    'Pendidikan (Penertiban Pelajar, Ketertiban Lingkungan Sekolah/Kampus)',
    'Usaha Tertentu (Tempat Hiburan Malam, Miras Tanpa Izin, Izin Usaha)',
    'Sosial (Gelandangan, Pengemis, Ketertiban Sosial Masyarakat)',
    'Barang Milik Daerah (Penyerobotan & Pemanfaatan Aset Pemprov)',
    'Keadaan Bencana (Mitigasi Kebakaran, Evakuasi Darurat, Linmas)',
    'Tertib Lainnya sesuai Kewenangan Pemerintah Provinsi'
  ];

  let violationType = violationTypes[0];
  let regencyId = '7301'; // Default Kota Makassar
  let districtName = '';
  let villageName = '';
  let addressDetail = '';
  let latitude = -5.147665;
  let longitude = 119.432731;
  let description = '';

  // Evidence Image Upload (STRICTLY images only: jpeg, png, webp)
  let evidenceImageFile: File | null = null;
  let evidenceImagePreview: string = '';
  let imageError = '';

  let isLoading = false;
  let successResult: { ticket_code: string; message: string; masked_info?: string } | null = null;
  let errorMessage = '';

  onMount(async () => {
    try {
      regencies = await fetchRegencies();
    } catch (e) {
      console.error(e);
    }
  });

  // Handle strictly images only
  function handleImageChange(e: Event) {
    imageError = '';
    const input = e.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];

      // Validate mime type strictly
      const validTypes = ['image/jpeg', 'image/png', 'image/webp', 'image/jpg'];
      if (!validTypes.includes(file.type.toLowerCase())) {
        imageError = 'Format berkas tidak valid! Hanya diperbolehkan berkas foto/gambar (JPG, PNG, atau WEBP).';
        input.value = '';
        return;
      }

      // Check file size (max 5 MB)
      if (file.size > 5 * 1024 * 1024) {
        imageError = 'Ukuran foto terlalu besar! Maksimal ukuran foto adalah 5 MB.';
        input.value = '';
        return;
      }

      evidenceImageFile = file;
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

  function detectGPS() {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (pos) => {
          latitude = pos.coords.latitude;
          longitude = pos.coords.longitude;
        },
        (err) => {
          alert('Deteksi lokasi gagal: ' + err.message);
        }
      );
    } else {
      alert('Browser Anda tidak mendukung deteksi lokasi otomatis.');
    }
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

      const res = await submitComplaint({
        is_anonymous: isAnonymous,
        reporter_nik: nik,
        reporter_name: name,
        reporter_phone: phone,
        reporter_email: email,
        masked_nik,
        masked_name,
        masked_phone,
        violation_type: violationType,
        regency_id: regencyId,
        regency_name: selectedRegency ? selectedRegency.name : 'Sulawesi Selatan',
        district_name: districtName,
        village_name: villageName,
        address_detail: addressDetail,
        latitude,
        longitude,
        description,
        evidence_photos: evidenceImagePreview || ''
      });

      if (res.success) {
        successResult = {
          ticket_code: res.ticket_code,
          message: res.message,
          masked_info: isAnonymous ? `${masked_name} • ${masked_nik}` : undefined
        };
      } else {
        errorMessage = res.message || 'Gagal mengirimkan laporan pengaduan.';
      }
    } catch (err: any) {
      errorMessage = 'Terjadi kesalahan sistem: ' + err.message;
    } finally {
      isLoading = false;
    }
  }
</script>

<div class="bg-white border border-[#eee9db] shadow-xs rounded-2xl overflow-hidden">
  <!-- Form Header matching SIAP TERTIB -->
  <div class="bg-[#12151e] text-white p-6 md:p-8 border-b-4 border-[#F08300]">
    <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider mb-1">
      Layanan Pengaduan Elektronik &bull; SIAP TERTIB
    </div>
    <h3 class="text-xl md:text-2xl font-extrabold tracking-tight">Formulir Laporan Pelanggaran Perda &amp; Trantibum</h3>
    <p class="text-xs text-[#d1d5db] mt-1.5 leading-relaxed">
      Silakan lengkapi formulir di bawah ini dengan informasi akurat. Bukti foto kejadian wajib berupa berkas gambar asli (JPG, PNG, WEBP) untuk diverifikasi oleh personel Satpol PP Prov. Sulsel.
    </p>
  </div>

  {#if successResult}
    <!-- Success Confirmation Receipt Card -->
    <div class="p-8 md:p-10 text-center space-y-5">
      <div class="w-14 h-14 bg-emerald-100 text-emerald-800 rounded-full flex items-center justify-center mx-auto text-xl font-extrabold">
        [OK]
      </div>
      <div>
        <div class="text-xs font-bold text-emerald-800 uppercase tracking-widest">Laporan Berhasil Diregistrasi</div>
        <h4 class="text-xl font-extrabold text-[#101116] mt-1">Tanda Terima Pengaduan Resmi</h4>
        <p class="text-xs text-[#6f7775] max-w-md mx-auto mt-1">{successResult.message}</p>
      </div>
      
      <div class="bg-[#f9f4e8] border border-[#eee9db] p-5 max-w-md mx-auto text-center rounded-2xl shadow-xs space-y-2">
        <div class="text-[11px] text-[#6f7775] font-semibold uppercase tracking-wider">Nomor Resi / Tiket Aduan:</div>
        <div class="text-2xl md:text-3xl font-extrabold text-[#101116] font-mono tracking-wider">{successResult.ticket_code}</div>
        {#if successResult.masked_info}
          <div class="text-[11px] text-[#F08300] font-bold pt-1 border-t border-[#eee9db]">
            Mode Anonim Aktif: {successResult.masked_info}
          </div>
        {/if}
        <div class="text-[10px] text-[#6f7775]">
          Simpan nomor tiket ini untuk memantau status perkembangan penanganan oleh Satpol PP Sulsel.
        </div>
      </div>

      <div class="pt-4 flex flex-wrap justify-center gap-3">
        <a 
          href={`/lacak?ticket=${successResult.ticket_code}`} 
          class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-6 py-3 rounded-xl transition shadow-xs"
        >
          Pantau Status Penanganan &rarr;
        </a>
        <button 
          on:click={() => { successResult = null; evidenceImagePreview = ''; evidenceImageFile = null; }}
          class="bg-gray-100 hover:bg-gray-200 text-[#101116] text-xs font-bold uppercase px-6 py-3 rounded-xl transition"
        >
          Kirim Laporan Lainnya
        </button>
      </div>
    </div>
  {:else}
    <!-- Form Content -->
    <form on:submit|preventDefault={handleSubmit} class="p-6 md:p-8 space-y-6 text-xs">
      {#if errorMessage}
        <div class="bg-rose-50 border border-rose-300 text-rose-900 text-xs p-4 rounded-xl font-semibold">
          {errorMessage}
        </div>
      {/if}

      <!-- 1. Identitas Pelapor & Mode Anonim -->
      <div class="space-y-4">
        <div class="border-b border-gray-100 pb-2 flex justify-between items-center">
          <h4 class="font-extrabold text-sm uppercase text-[#101116] tracking-wide">1. Data Diri Pelapor</h4>
          <span class="text-[11px] text-[#6f7775] font-medium">* Wajib diisi untuk verifikasi hukum</span>
        </div>

        <!-- Mode Anonim Toggle Banner -->
        <div class="bg-[#f9f4e8] border border-[#eee9db] p-4 rounded-xl flex items-start justify-between gap-4">
          <div>
            <label for="anonim-check" class="font-bold text-xs uppercase tracking-wider text-[#101116] cursor-pointer block">
              Sembunyikan Identitas (Lapor Anonim Terverifikasi)
            </label>
            <p class="text-[11px] text-[#6f7775] mt-1 leading-relaxed">
              Jika diaktifkan, data NIK, nama, dan telepon Anda tetap kami catat untuk pencegahan laporan palsu, namun <strong>otomatis dimasking/disamarkan</strong> (contoh: NIK <code>7371**********12</code>) pada tanda terima dan publik.
            </p>
          </div>
          <input 
            id="anonim-check" 
            type="checkbox" 
            bind:checked={isAnonymous} 
            class="h-5 w-5 text-[#F08300] focus:ring-[#F08300] border-gray-300 rounded cursor-pointer mt-0.5 flex-shrink-0"
          />
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label for="nik-input" class="block font-bold uppercase text-[#101116] mb-1">
              Nomor Induk Kependudukan (NIK) *
            </label>
            <input 
              id="nik-input" 
              type="text" 
              maxlength="16" 
              bind:value={nik} 
              placeholder="16 Digit NIK KTP Anda" 
              required
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] focus:outline-none" 
            />
            {#if isAnonymous && nik.length >= 6}
              <div class="text-[10px] text-emerald-700 font-semibold mt-1">
                Akan dimasking menjadi: <strong>{maskNIK(nik)}</strong>
              </div>
            {/if}
          </div>

          <div>
            <label for="name-input" class="block font-bold uppercase text-[#101116] mb-1">
              Nama Lengkap Pelapor *
            </label>
            <input 
              id="name-input" 
              type="text" 
              bind:value={name} 
              placeholder="Nama sesuai identitas resmi" 
              required
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
            />
            {#if isAnonymous && name.trim().length > 2}
              <div class="text-[10px] text-emerald-700 font-semibold mt-1">
                Akan dimasking menjadi: <strong>{maskName(name)}</strong>
              </div>
            {/if}
          </div>

          <div>
            <label for="phone-input" class="block font-bold uppercase text-[#101116] mb-1">
              Nomor Telepon / WhatsApp Aktif *
            </label>
            <input 
              id="phone-input" 
              type="tel" 
              bind:value={phone} 
              placeholder="08xxxxxxxxxx" 
              required
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
            />
          </div>

          <div>
            <label for="email-input" class="block font-bold uppercase text-[#101116] mb-1">
              Alamat Email (Opsional)
            </label>
            <input 
              id="email-input" 
              type="email" 
              bind:value={email} 
              placeholder="email@domain.com" 
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
            />
          </div>
        </div>
      </div>

      <!-- 2. Klasifikasi Pelanggaran & Lokasi Kejadian -->
      <div class="space-y-4 pt-4 border-t border-gray-100">
        <div class="border-b border-gray-100 pb-2">
          <h4 class="font-extrabold text-sm uppercase text-[#101116] tracking-wide">2. Klasifikasi &amp; Lokasi Pelanggaran</h4>
        </div>

        <div>
          <label for="violation-select" class="block font-bold uppercase text-[#101116] mb-1">
            Jenis Pelanggaran / Aduan (9 Kategori Resmi SIAP TERTIB Sulsel) *
          </label>
          <select 
            id="violation-select" 
            bind:value={violationType} 
            class="w-full border border-gray-300 rounded-lg p-2.5 text-xs font-semibold text-[#101116] focus:border-[#F08300] focus:outline-none bg-white"
          >
            {#each violationTypes as vt}
              <option value={vt}>{vt}</option>
            {/each}
          </select>
        </div>

        <!-- Lokasi Bertingkat (Sulsel) -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label for="regency-select" class="block font-bold uppercase text-[#101116] mb-1">Kabupaten / Kota *</label>
            <select 
              id="regency-select" 
              bind:value={regencyId} 
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs focus:border-[#F08300] focus:outline-none bg-white"
            >
              {#each regencies as reg}
                <option value={reg.id}>{reg.name}</option>
              {/each}
            </select>
          </div>

          <div>
            <label for="district-input" class="block font-bold uppercase text-[#101116] mb-1">Kecamatan</label>
            <input 
              id="district-input" 
              type="text" 
              bind:value={districtName} 
              placeholder="Contoh: Panakkukang" 
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
            />
          </div>

          <div>
            <label for="village-input" class="block font-bold uppercase text-[#101116] mb-1">Kelurahan / Desa</label>
            <input 
              id="village-input" 
              type="text" 
              bind:value={villageName} 
              placeholder="Contoh: Panaikang" 
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
            />
          </div>
        </div>

        <!-- Alamat Detail & GPS Coordinates -->
        <div>
          <div class="flex justify-between items-center mb-1">
            <label for="address-textarea" class="font-bold uppercase text-[#101116]">
              Alamat Detail Tempat Kejadian *
            </label>
            <button 
              type="button" 
              on:click={detectGPS} 
              class="text-xs font-bold text-[#F08300] hover:underline"
            >
              Deteksi Koordinat GPS Otomatis
            </button>
          </div>
          <textarea 
            id="address-textarea" 
            bind:value={addressDetail} 
            rows="2" 
            required
            placeholder="Sebutkan nama jalan, patokan gedung, nomor bangunan, atau titik persimpangan..." 
            class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none"
          ></textarea>
          <div class="text-[11px] text-[#6f7775] mt-1 font-mono">
            Titik Koordinat: {latitude.toFixed(6)}, {longitude.toFixed(6)}
          </div>
        </div>
      </div>

      <!-- 3. Bukti Foto Kejadian (HANYA GAMBAR) & Uraian -->
      <div class="space-y-4 pt-4 border-t border-gray-100">
        <div class="border-b border-gray-100 pb-2">
          <h4 class="font-extrabold text-sm uppercase text-[#101116] tracking-wide">3. Bukti Gambar &amp; Uraian Kejadian</h4>
        </div>

        <!-- Strict Image-Only Upload Area -->
        <div>
          <label for="evidence-photo-input" class="block font-bold uppercase text-[#101116] mb-1">
            Foto Bukti Kejadian (Wajib Berkas Gambar: JPG, PNG, WEBP)
          </label>
          
          <div class="border-2 border-dashed border-gray-300 rounded-2xl p-5 text-center bg-gray-50/50 hover:bg-gray-50 transition">
            {#if evidenceImagePreview}
              <div class="space-y-3">
                <img 
                  src={evidenceImagePreview} 
                  alt="Bukti Aduan" 
                  class="max-h-56 mx-auto rounded-xl object-contain border border-gray-300 shadow-xs" 
                />
                <div class="flex justify-center items-center gap-3">
                  <span class="text-[11px] font-semibold text-emerald-700">Foto bukti siap dilampirkan</span>
                  <button 
                    type="button" 
                    on:click={removeImage}
                    class="bg-rose-50 hover:bg-rose-100 text-rose-700 font-bold text-[11px] uppercase px-3 py-1 rounded-lg border border-rose-200 transition"
                  >
                    Hapus / Ganti Foto
                  </button>
                </div>
              </div>
            {:else}
              <div class="space-y-2">
                <div class="text-xs text-gray-500 font-medium">
                  Pilih berkas foto bukti kejadian dari perangkat Anda:
                </div>
                <input 
                  id="evidence-photo-input"
                  type="file" 
                  accept="image/jpeg,image/png,image/webp" 
                  on:change={handleImageChange}
                  class="text-xs text-gray-600 file:mr-3 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-xs file:font-bold file:bg-[#3d261a] file:text-white hover:file:bg-[#271810] cursor-pointer"
                />
                <div class="text-[11px] text-[#6f7775]">
                  Catatan: Hanya menerima berkas foto/gambar. Maksimal ukuran 5 MB.
                </div>
              </div>
            {/if}

            {#if imageError}
              <div class="text-xs text-rose-700 font-semibold mt-2">
                {imageError}
              </div>
            {/if}
          </div>
        </div>

        <!-- Uraian Kronologi Kejadian -->
        <div>
          <label for="description-textarea" class="block font-bold uppercase text-[#101116] mb-1">
            Uraian / Kronologi Lengkap Kejadian *
          </label>
          <textarea 
            id="description-textarea" 
            bind:value={description} 
            rows="4" 
            required
            placeholder="Jelaskan secara rinci situasi pelanggaran, perkiraan waktu kejadian, dan dampak terhadap ketertiban umum..." 
            class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none"
          ></textarea>
        </div>
      </div>

      <!-- Submit Button -->
      <div class="pt-2">
        <button 
          type="submit" 
          disabled={isLoading}
          class="w-full bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase tracking-wider py-4 rounded-xl transition disabled:opacity-50 shadow-md flex items-center justify-center gap-2"
        >
          {#if isLoading}
            <span>Sedang Memproses Laporan...</span>
          {:else}
            <span>Kirim Laporan Pengaduan SIAP TERTIB</span>
          {/if}
        </button>
      </div>
    </form>
  {/if}
</div>
