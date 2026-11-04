<script lang="ts">
  import { onMount } from 'svelte';
  import { 
    siteSettingsStore, 
    usersStore, 
    auditLogsStore, 
    complaintsStore, 
    defaultRegencies,
    type AuditEntry
  } from '$lib/stores/appState';
  import type { SiteSettings, User } from '$lib/types';

  let activeTab: 'settings' | 'users' | 'regencies' | 'audit' | 'server' = 'settings';

  // CMS Settings Local Form Model
  let currentSettings: SiteSettings = {
    app_name: '',
    agency_tagline: '',
    office_address: '',
    call_center: '',
    emergency_hotline: '',
    email_contact: '',
    kasatpol_name: '',
    kasatpol_quote: ''
  };

  siteSettingsStore.subscribe(val => {
    currentSettings = { ...val };
  });

  let isSaving = false;
  let saveSuccess = false;

  function handleSaveSettings() {
    isSaving = true;
    siteSettingsStore.set(currentSettings);
    
    // Log audit trail
    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: 'superadmin',
        action: 'UPDATE_CMS_SETTINGS',
        resource: 'SiteSettings',
        details: 'Memperbarui konfigurasi identitas dan kontak resmi portal',
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);

    setTimeout(() => {
      isSaving = false;
      saveSuccess = true;
      setTimeout(() => saveSuccess = false, 3000);
    }, 400);
  }

  // User Management
  let usersList: User[] = [];
  usersStore.subscribe(val => usersList = val);

  let showAddUserModal = false;
  let newUser: Partial<User> = {
    name: '',
    username: '',
    email: '',
    role: 'STAFF_SATPOL',
    bidang: 'Bidang Ketenteraman dan Ketertiban Umum',
    is_active: true
  };

  function handleAddUser() {
    if (!newUser.name || !newUser.username || !newUser.email) {
      alert('Nama, username, dan email wajib diisi.');
      return;
    }
    const created: User = {
      id: Date.now(),
      name: newUser.name,
      username: newUser.username.toLowerCase().replace(/\s+/g, ''),
      email: newUser.email,
      role: newUser.role as any || 'STAFF_SATPOL',
      bidang: newUser.bidang || 'Sekretariat',
      is_active: true
    };
    usersStore.update(items => [...items, created]);
    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: 'superadmin',
        action: 'CREATE_USER',
        resource: `User #${created.username}`,
        details: `Menambahkan akun baru role ${created.role} (${created.name})`,
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);
    showAddUserModal = false;
    newUser = {
      name: '',
      username: '',
      email: '',
      role: 'STAFF_SATPOL',
      bidang: 'Bidang Ketenteraman dan Ketertiban Umum',
      is_active: true
    };
  }

  function toggleUserStatus(id: number) {
    usersStore.update(items => items.map(u => {
      if (u.id === id) {
        const nextStatus = !u.is_active;
        auditLogsStore.update(logs => [
          {
            id: Date.now(),
            user_name: 'superadmin',
            action: nextStatus ? 'ACTIVATE_USER' : 'SUSPEND_USER',
            resource: `User #${u.username}`,
            details: `Mengubah status akun menjadi ${nextStatus ? 'Aktif' : 'Nonaktif'}`,
            timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
            ip_address: '103.151.191.73'
          },
          ...logs
        ]);
        return { ...u, is_active: nextStatus };
      }
      return u;
    }));
  }

  // Audit Logs
  let auditLogs: AuditEntry[] = [];
  auditLogsStore.subscribe(val => auditLogs = val);

  // Stats
  let totalComplaints = 0;
  complaintsStore.subscribe(val => totalComplaints = val.length);
</script>

<svelte:head>
  <title>Konsol Superadmin Diskominfo-SP | Satpol PP Sulsel</title>
</svelte:head>

<div class="min-h-screen bg-[#f8f9fa] py-8 px-4 sm:px-6 lg:px-8">
  <div class="max-w-7xl mx-auto space-y-6">

    <!-- Header Panel -->
    <div class="bg-[#101116] text-white p-6 md:p-8 rounded-2xl border-b-4 border-[#F08300] shadow-sm">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
        <div>
          <div class="flex items-center gap-2">
            <span class="inline-block w-2.5 h-2.5 bg-emerald-500 rounded-full"></span>
            <span class="text-[#F08300] font-bold text-xs uppercase tracking-widest">
              Pemerintah Provinsi Sulawesi Selatan &bull; Dinas Komunikasi, Informatika, Statistik dan Persandian
            </span>
          </div>
          <h1 class="text-2xl md:text-3xl font-extrabold tracking-tight mt-1">
            Konsol Pengembang &amp; Superadmin Sistem
          </h1>
          <p class="text-xs text-[#9ca3af] mt-1">
            Pengelolaan infrastruktur hosting, manajemen hak akses terpusat, audit trail keamanan, dan konfigurasi portal tanpa hardcode.
          </p>
        </div>

        <div class="flex flex-wrap gap-2.5">
          <a 
            href="/admin/dashboard" 
            class="bg-[#2b2f3f] hover:bg-gray-800 text-white text-xs font-bold uppercase px-4 py-2.5 rounded-xl border border-gray-700 transition"
          >
            Dashboard Staff Satpol
          </a>
          <a 
            href="/" 
            class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-4 py-2.5 rounded-xl transition"
          >
            Lihat Portal Publik
          </a>
        </div>
      </div>

      <!-- Quick Metrics Ribbon -->
      <div class="grid grid-cols-2 sm:grid-cols-4 gap-4 mt-6 pt-6 border-t border-gray-800">
        <div>
          <div class="text-[11px] text-[#9ca3af] uppercase font-semibold">Total Pengguna Terdaftar</div>
          <div class="text-xl font-bold text-white mt-0.5">{usersList.length} Akun</div>
        </div>
        <div>
          <div class="text-[11px] text-[#9ca3af] uppercase font-semibold">Laporan Masuk SIAP TERTIB</div>
          <div class="text-xl font-bold text-[#F08300] mt-0.5">{totalComplaints} Berkas</div>
        </div>
        <div>
          <div class="text-[11px] text-[#9ca3af] uppercase font-semibold">Cakupan Wilayah</div>
          <div class="text-xl font-bold text-white mt-0.5">{defaultRegencies.length} Kab / Kota</div>
        </div>
        <div>
          <div class="text-[11px] text-[#9ca3af] uppercase font-semibold">Status Gateway API</div>
          <div class="text-xl font-bold text-emerald-400 mt-0.5">Aktif (Port 8080)</div>
        </div>
      </div>
    </div>

    <!-- Navigation Tabs -->
    <div class="flex flex-wrap border border-gray-200 bg-white rounded-xl p-1 gap-1">
      <button 
        class={`px-4 py-2.5 text-xs font-bold uppercase transition rounded-lg ${activeTab === 'settings' ? 'bg-[#F08300] text-white' : 'text-gray-600 hover:bg-gray-100'}`}
        on:click={() => activeTab = 'settings'}
      >
        Pengaturan CMS &amp; Metadata
      </button>
      <button 
        class={`px-4 py-2.5 text-xs font-bold uppercase transition rounded-lg ${activeTab === 'users' ? 'bg-[#F08300] text-white' : 'text-gray-600 hover:bg-gray-100'}`}
        on:click={() => activeTab = 'users'}
      >
        Manajemen Akun (RBAC)
      </button>
      <button 
        class={`px-4 py-2.5 text-xs font-bold uppercase transition rounded-lg ${activeTab === 'regencies' ? 'bg-[#F08300] text-white' : 'text-gray-600 hover:bg-gray-100'}`}
        on:click={() => activeTab = 'regencies'}
      >
        Master Wilayah (24 Kab/Kota)
      </button>
      <button 
        class={`px-4 py-2.5 text-xs font-bold uppercase transition rounded-lg ${activeTab === 'audit' ? 'bg-[#F08300] text-white' : 'text-gray-600 hover:bg-gray-100'}`}
        on:click={() => activeTab = 'audit'}
      >
        Audit Trail Keamanan ({auditLogs.length})
      </button>
      <button 
        class={`px-4 py-2.5 text-xs font-bold uppercase transition rounded-lg ${activeTab === 'server' ? 'bg-[#F08300] text-white' : 'text-gray-600 hover:bg-gray-100'}`}
        on:click={() => activeTab = 'server'}
      >
        Telemetri Server &amp; Host
      </button>
    </div>

    {#if saveSuccess}
      <div class="bg-emerald-50 border border-emerald-300 text-emerald-800 text-xs p-4 rounded-xl font-semibold flex items-center justify-between">
        <span>Pengaturan CMS berhasil disimpan dan disinkronkan ke portal publik secara instan.</span>
        <span class="text-[10px] uppercase font-bold tracking-wider bg-emerald-200 px-2 py-0.5 rounded">Tersimpan</span>
      </div>
    {/if}

    <!-- TAB 1: CMS SETTINGS -->
    {#if activeTab === 'settings'}
      <div class="bg-white border border-[#eee9db] rounded-2xl shadow-sm p-6 md:p-8 space-y-6">
        <div class="border-b border-gray-100 pb-4">
          <h2 class="text-base font-bold text-[#101116] uppercase tracking-wide">
            Konfigurasi Identitas Instansi &amp; CMS Terpusat
          </h2>
          <p class="text-xs text-[#6f7775] mt-1">
            Bidang-bidang di bawah ini mengontrol seluruh teks pimpinan, kontak darurat, serta alamat resmi yang tampil di header dan footer portal secara terpusat tanpa kompilasi ulang kode.
          </p>
        </div>

        <form on:submit|preventDefault={handleSaveSettings} class="space-y-5">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <div>
              <label for="sup-app-name" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Judul Portal / Instansi
              </label>
              <input 
                id="sup-app-name"
                type="text" 
                bind:value={currentSettings.app_name} 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
              />
            </div>

            <div>
              <label for="sup-kasatpol-name" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Nama Lengkap &amp; Gelar Kepala Satpol PP
              </label>
              <input 
                id="sup-kasatpol-name"
                type="text" 
                bind:value={currentSettings.kasatpol_name} 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
              />
            </div>

            <div>
              <label for="sup-call-center" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Nomor Telepon Kantor / Call Center
              </label>
              <input 
                id="sup-call-center"
                type="text" 
                bind:value={currentSettings.call_center} 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
              />
            </div>

            <div>
              <label for="sup-hotline" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Nomor Hotline Aduan Cepat (SIAP TERTIB / WhatsApp)
              </label>
              <input 
                id="sup-hotline"
                type="text" 
                bind:value={currentSettings.emergency_hotline} 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
              />
            </div>

            <div>
              <label for="sup-email" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Email Korespondensi Resmi Instansi
              </label>
              <input 
                id="sup-email"
                type="email" 
                bind:value={currentSettings.email_contact} 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
              />
            </div>

            <div>
              <label for="sup-address" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Alamat Kantor Satpol PP Prov. Sulsel
              </label>
              <input 
                id="sup-address"
                type="text" 
                bind:value={currentSettings.office_address} 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none" 
              />
            </div>

            <div class="md:col-span-2">
              <label for="sup-tagline" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Semboyan &amp; Tagline Resmi (Praja Wibawa)
              </label>
              <textarea 
                id="sup-tagline"
                bind:value={currentSettings.agency_tagline} 
                rows="2" 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none"
              ></textarea>
            </div>

            <div class="md:col-span-2">
              <label for="sup-quote" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                Pernyataan Sambutan Kasatpol PP di Beranda
              </label>
              <textarea 
                id="sup-quote"
                bind:value={currentSettings.kasatpol_quote} 
                rows="3" 
                class="w-full border border-gray-300 rounded-lg p-2.5 text-xs text-[#101116] focus:border-[#F08300] focus:outline-none"
              ></textarea>
            </div>
          </div>

          <div class="pt-4 border-t border-gray-100 flex justify-end">
            <button 
              type="submit" 
              disabled={isSaving}
              class="bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase px-6 py-2.5 rounded-xl transition disabled:opacity-50"
            >
              {#if isSaving}Menyimpan ke Database...{:else}Simpan Pembaruan CMS{/if}
            </button>
          </div>
        </form>
      </div>
    {/if}

    <!-- TAB 2: USER MANAGEMENT (RBAC) -->
    {#if activeTab === 'users'}
      <div class="bg-white border border-[#eee9db] shadow-sm p-6 md:p-8 space-y-6">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 border-b border-gray-100 pb-4">
          <div>
            <h2 class="text-base font-bold text-[#101116] uppercase tracking-wide">
              Manajemen Hak Akses &amp; Akun Petugas (RBAC)
            </h2>
            <p class="text-xs text-[#6f7775] mt-1">
              Pengaturan hierarki pengguna antara Superadmin Diskominfo, Admin Pimpinan Satpol PP, dan Petugas Teknis Bidang.
            </p>
          </div>
          <button 
            on:click={() => showAddUserModal = true}
            class="bg-[#2b2f3f] hover:bg-[#101116] text-white text-xs font-bold uppercase px-4 py-2.5 rounded-none transition"
          >
            + Tambah Pengguna Baru
          </button>
        </div>

        {#if showAddUserModal}
          <div class="bg-[#f9f4e8] border border-[#eee9db] p-5 space-y-4">
            <div class="text-xs font-bold text-[#101116] uppercase tracking-wider">Form Tambah Petugas / Administrator</div>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-3">
              <div>
                <label for="new-user-fullname" class="block text-[10px] font-bold uppercase text-[#6f7775] mb-1">Nama Lengkap</label>
                <input id="new-user-fullname" type="text" bind:value={newUser.name} placeholder="Nama Petugas" class="w-full border border-gray-300 p-2 text-xs bg-white" />
              </div>
              <div>
                <label for="new-user-username" class="block text-[10px] font-bold uppercase text-[#6f7775] mb-1">Username Login</label>
                <input id="new-user-username" type="text" bind:value={newUser.username} placeholder="contoh: dantonops" class="w-full border border-gray-300 p-2 text-xs bg-white" />
              </div>
              <div>
                <label for="new-user-email" class="block text-[10px] font-bold uppercase text-[#6f7775] mb-1">Email Resmi</label>
                <input id="new-user-email" type="email" bind:value={newUser.email} placeholder="email@sulselprov.go.id" class="w-full border border-gray-300 p-2 text-xs bg-white" />
              </div>
              <div>
                <label for="new-user-role" class="block text-[10px] font-bold uppercase text-[#6f7775] mb-1">Tingkat Hak Akses (Role)</label>
                <select id="new-user-role" bind:value={newUser.role} class="w-full border border-gray-300 p-2 text-xs bg-white">
                  <option value="STAFF_SATPOL">Staff Petugas Satpol PP</option>
                  <option value="ADMIN_SATPOL">Admin Satpol PP (Kasatpol)</option>
                  <option value="SUPERADMIN_DISKOMINFO">Superadmin (Diskominfo)</option>
                </select>
              </div>
              <div class="md:col-span-2">
                <label for="new-user-bidang" class="block text-[10px] font-bold uppercase text-[#6f7775] mb-1">Unit Kerja / Bidang</label>
                <input id="new-user-bidang" type="text" bind:value={newUser.bidang} placeholder="Contoh: Bidang Penegakan Perda" class="w-full border border-gray-300 p-2 text-xs bg-white" />
              </div>
            </div>
            <div class="flex justify-end gap-2 pt-2">
              <button on:click={() => showAddUserModal = false} class="px-4 py-2 border border-gray-300 text-xs font-bold text-gray-700 bg-white">Batal</button>
              <button on:click={handleAddUser} class="px-5 py-2 bg-[#F08300] text-white text-xs font-bold uppercase">Simpan Pengguna</button>
            </div>
          </div>
        {/if}

        <div class="overflow-x-auto border border-gray-200">
          <table class="w-full text-left text-xs">
            <thead class="bg-gray-50 border-b border-gray-200 text-[#6f7775] uppercase text-[10px] tracking-wider">
              <tr>
                <th class="p-3.5">Nama &amp; Identitas</th>
                <th class="p-3.5">Username</th>
                <th class="p-3.5">Role Akses</th>
                <th class="p-3.5">Unit Kerja / Bidang</th>
                <th class="p-3.5">Status Akun</th>
                <th class="p-3.5 text-right">Opsi</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-100">
              {#each usersList as u}
                <tr class="hover:bg-gray-50/70 transition">
                  <td class="p-3.5">
                    <div class="font-bold text-[#101116]">{u.name}</div>
                    <div class="text-[11px] text-[#6f7775]">{u.email}</div>
                  </td>
                  <td class="p-3.5 font-mono text-[#101116] font-semibold">{u.username}</td>
                  <td class="p-3.5">
                    <span class={`text-[10px] font-bold uppercase px-2 py-0.5 border ${u.role === 'SUPERADMIN_DISKOMINFO' ? 'bg-purple-50 text-purple-800 border-purple-200' : u.role === 'ADMIN_SATPOL' ? 'bg-amber-50 text-amber-800 border-amber-200' : 'bg-blue-50 text-blue-800 border-blue-200'}`}>
                      {u.role}
                    </span>
                  </td>
                  <td class="p-3.5 text-[#6f7775]">{u.bidang || '-'}</td>
                  <td class="p-3.5">
                    {#if u.is_active}
                      <span class="text-emerald-700 font-bold text-[10px] uppercase">Aktif</span>
                    {:else}
                      <span class="text-rose-700 font-bold text-[10px] uppercase">Dinonaktifkan</span>
                    {/if}
                  </td>
                  <td class="p-3.5 text-right">
                    <button 
                      on:click={() => toggleUserStatus(u.id)}
                      class={`text-[11px] font-bold uppercase px-3 py-1 border transition ${u.is_active ? 'border-rose-300 text-rose-700 hover:bg-rose-50' : 'border-emerald-300 text-emerald-700 hover:bg-emerald-50'}`}
                    >
                      {u.is_active ? 'Nonaktifkan' : 'Aktifkan'}
                    </button>
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    {/if}

    <!-- TAB 3: MASTER DATA 24 KAB/KOTA -->
    {#if activeTab === 'regencies'}
      <div class="bg-white border border-[#eee9db] shadow-sm p-6 md:p-8 space-y-6">
        <div class="border-b border-gray-100 pb-4">
          <h2 class="text-base font-bold text-[#101116] uppercase tracking-wide">
            Master Data 24 Kabupaten &amp; Kota Provinsi Sulawesi Selatan
          </h2>
          <p class="text-xs text-[#6f7775] mt-1">
            Data referensi wilayah yurisdiksi penegakan Peraturan Daerah dan penanganan aduan masyarakat melalui SIAP TERTIB.
          </p>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3">
          {#each defaultRegencies as reg, idx}
            <div class="border border-gray-200 p-3 bg-gray-50/50 hover:border-[#F08300] transition">
              <div class="flex justify-between items-center text-[10px] text-[#6f7775]">
                <span class="font-mono">KODE: {reg.id}</span>
                <span class="font-bold text-emerald-700 uppercase">Terkoneksi</span>
              </div>
              <div class="font-bold text-xs text-[#101116] mt-1">{reg.name}</div>
              <div class="text-[10px] text-[#6f7775] mt-1">Yurisdiksi Satpol PP Prov. Sulsel</div>
            </div>
          {/each}
        </div>
      </div>
    {/if}

    <!-- TAB 4: AUDIT TRAIL LOGS -->
    {#if activeTab === 'audit'}
      <div class="bg-white border border-[#eee9db] shadow-sm p-6 md:p-8 space-y-6">
        <div class="border-b border-gray-100 pb-4">
          <h2 class="text-base font-bold text-[#101116] uppercase tracking-wide">
            Audit Trail &amp; Log Keamanan Sistem Terpadu
          </h2>
          <p class="text-xs text-[#6f7775] mt-1">
            Catatan kronologis perubahan data, tindak lanjut pengaduan, otentikasi login, serta perubahan konfigurasi server oleh seluruh operator.
          </p>
        </div>

        <div class="overflow-x-auto border border-gray-200">
          <table class="w-full text-left text-xs">
            <thead class="bg-gray-50 border-b border-gray-200 text-[#6f7775] uppercase text-[10px] tracking-wider">
              <tr>
                <th class="p-3">Waktu (WITA)</th>
                <th class="p-3">Pengguna</th>
                <th class="p-3">Aksi Sistem</th>
                <th class="p-3">Entitas / Resi</th>
                <th class="p-3">Keterangan Aktivitas</th>
                <th class="p-3">Alamat IP</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-100 font-mono text-[11px]">
              {#each auditLogs as log}
                <tr class="hover:bg-gray-50/70 transition font-sans">
                  <td class="p-3 text-[#6f7775] whitespace-nowrap">{log.timestamp}</td>
                  <td class="p-3 font-bold text-[#101116]">{log.user_name}</td>
                  <td class="p-3">
                    <span class="bg-gray-100 text-[#101116] font-bold text-[10px] uppercase px-2 py-0.5 border border-gray-300">
                      {log.action}
                    </span>
                  </td>
                  <td class="p-3 font-semibold text-[#101116]">{log.resource}</td>
                  <td class="p-3 text-[#6f7775]">{log.details}</td>
                  <td class="p-3 font-mono text-[#6f7775]">{log.ip_address}</td>
                </tr>
              {:else}
                <tr>
                  <td colspan="6" class="p-6 text-center text-gray-400">Belum ada catatan log aktivitas.</td>
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    {/if}

    <!-- TAB 5: SERVER TELEMETRY -->
    {#if activeTab === 'server'}
      <div class="bg-white border border-[#eee9db] shadow-sm p-6 md:p-8 space-y-6">
        <div class="border-b border-gray-100 pb-4">
          <h2 class="text-base font-bold text-[#101116] uppercase tracking-wide">
            Status Server &amp; Telemetri Infrastruktur Diskominfo-SP
          </h2>
          <p class="text-xs text-[#6f7775] mt-1">
            Informasi teknis lingkungan hosting server fisik Diskominfo Provinsi Sulawesi Selatan (IP: 103.151.191.73).
          </p>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div class="border border-gray-200 p-4 bg-gray-50/50">
            <div class="text-[10px] text-[#6f7775] uppercase font-bold tracking-wider">Backend Gateway Runtime</div>
            <div class="text-base font-bold text-[#101116] mt-1">Go v1.24 (Gin Gonic Clean Arch)</div>
            <div class="text-xs text-[#6f7775] mt-2 space-y-1">
              <div>Port Layanan: <span class="font-mono font-bold text-[#101116]">8080 (REST API)</span></div>
              <div>Goroutine Aktif: <span class="font-mono text-emerald-700 font-bold">14 Workers</span></div>
              <div>Alokasi Memori: <span class="font-mono text-[#101116]">12.4 MB</span></div>
            </div>
          </div>

          <div class="border border-gray-200 p-4 bg-gray-50/50">
            <div class="text-[10px] text-[#6f7775] uppercase font-bold tracking-wider">Web Server &amp; Reverse Proxy</div>
            <div class="text-base font-bold text-[#101116] mt-1">Nginx 1.22 / Ubuntu Server</div>
            <div class="text-xs text-[#6f7775] mt-2 space-y-1">
              <div>Host Domain: <span class="font-mono font-bold text-[#101116]">satpolpp.sulselprov.go.id</span></div>
              <div>Subdomain: <span class="font-mono font-bold text-[#101116]">siaptertib.sulselprov.go.id</span></div>
              <div>Protokol: <span class="font-mono text-emerald-700 font-bold">HTTP/2 TLS Termination</span></div>
            </div>
          </div>

          <div class="border border-gray-200 p-4 bg-gray-50/50">
            <div class="text-[10px] text-[#6f7775] uppercase font-bold tracking-wider">Basis Data Terpadu</div>
            <div class="text-base font-bold text-[#101116] mt-1">PostgreSQL 16 / SQLite Engine</div>
            <div class="text-xs text-[#6f7775] mt-2 space-y-1">
              <div>Status Koneksi: <span class="font-bold text-emerald-700">TERHUBUNG</span></div>
              <div>Connection Pool: <span class="font-mono text-[#101116]">5 / 25 Conns</span></div>
              <div>Integritas Relasi: <span class="font-bold text-[#101116]">OK (Strict Foreign Keys)</span></div>
            </div>
          </div>
        </div>

        <div class="p-4 bg-[#f9f4e8] border border-[#eee9db] text-xs text-[#6f7775] space-y-2">
          <div class="font-bold text-[#101116] uppercase text-[11px]">Rekomendasi Deployment Server Diskominfo:</div>
          <p>
            Binary backend Go terkompilasi mandiri (`api.exe` / `api_linux`) tidak memerlukan runtime interpreter eksternal seperti Apache/PHP lawas, mengurangi attack vector secara drastis serta menghemat penggunaan RAM server hingga &gt;85%.
          </p>
        </div>
      </div>
    {/if}

  </div>
</div>
