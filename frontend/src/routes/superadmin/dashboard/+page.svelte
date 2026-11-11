<script lang="ts">
  import { onMount } from 'svelte';
  import {
    Users,
    Building2,
    Settings,
    MapPin,
    History,
    ShieldCheck,
    Plus,
    Edit3,
    Trash2,
    CheckCircle2,
    UserCheck,
    UserX
  } from 'lucide-svelte';
  import {
    siteSettingsStore,
    usersStore,
    auditLogsStore,
    complaintsStore,
    organizationStore,
    defaultRegencies,
    logAudit
  } from '$lib/stores/appState';
  import {
    saveUser,
    setUserActive,
    deleteUser,
    fetchOrganization
  } from '$lib/api/client';
  import { getSession } from '$lib/utils/auth';
  import type { SiteSettings, User, UserRole } from '$lib/types';
  import DashboardShell, { type DashboardNavItem } from '$lib/components/layout/DashboardShell.svelte';
  import DataTable from '$lib/components/ui/DataTable.svelte';
  import Modal from '$lib/components/ui/Modal.svelte';
  import OrganizationManager from '$lib/components/cms/OrganizationManager.svelte';

  type SuperTab = 'users' | 'organization' | 'settings' | 'regencies' | 'audit' | 'server';
  let activeTab: SuperTab = 'users';
  let currentUser: User | null = null;

  let toastMessage = '';
  function showToast(msg: string) {
    toastMessage = msg;
    setTimeout(() => {
      if (toastMessage === msg) toastMessage = '';
    }, 5000);
  }

  onMount(() => {
    const session = getSession();
    if (session?.user) {
      currentUser = session.user;
    }
    fetchOrganization();
  });

  $: navItems = [
    { id: 'users', label: 'Manajemen Pengguna', icon: Users, badge: $usersStore.length, section: 'Akses & Organisasi' },
    { id: 'organization', label: 'Struktur Organisasi', icon: Building2, badge: $organizationStore.length },
    { id: 'settings', label: 'Konfigurasi Portal', icon: Settings, section: 'Sistem & Wilayah' },
    { id: 'regencies', label: '24 Kab/Kota', icon: MapPin, badge: defaultRegencies.length },
    { id: 'audit', label: 'Audit Trail', icon: History, badge: $auditLogsStore.length, section: 'Keamanan' },
    { id: 'server', label: 'Telemetri Sistem', icon: ShieldCheck }
  ] satisfies DashboardNavItem[];

  const tabMeta: Record<SuperTab, { title: string; desc: string }> = {
    users: {
      title: 'Manajemen Pengguna & Hak Akses (RBAC)',
      desc: 'Kelola akun Superadmin, Kasatpol, dan Staf Bidang termasuk aktivasi dan pengaturan peran.'
    },
    organization: {
      title: 'Manajemen Struktur Organisasi (SOTK)',
      desc: 'Tambah dan perbarui simpul jabatan, ID khusus (ORG-*), serta sub-unit kerja dinamis.'
    },
    settings: {
      title: 'Konfigurasi Identitas Portal & Kontak Resmi',
      desc: 'Ubah informasi instansi, nama pimpinan, sambutan, dan kontak SIAP TERTIB tanpa hardcode.'
    },
    regencies: {
      title: 'Referensi 24 Kabupaten / Kota Sulawesi Selatan',
      desc: 'Wilayah yurisdiksi penegakan Peraturan Daerah yang terhubung ke formulir aduan SIAP TERTIB.'
    },
    audit: {
      title: 'Audit Trail & Log Aktivitas Sistem',
      desc: 'Jejak kronologis pembukaan masking identitas, perubahan status SOP, serta administrasi akun.'
    },
    server: {
      title: 'Telemetri Sistem & Lingkungan Runtime',
      desc: 'Ringkasan arsitektur frontend SvelteKit, REST API Go, dan mode penyimpanan aktif.'
    }
  };

  // ==========================================
  // 1. MANAJEMEN PENGGUNA (MODAL CRUD + RBAC)
  // ==========================================
  let userModalOpen = false;
  let isNewUser = true;
  let userFormError = '';
  let isSavingUser = false;

  let formUserId = 0;
  let formName = '';
  let formUsername = '';
  let formEmail = '';
  let formRole: UserRole = 'STAFF_SATPOL';
  let formBidang = 'Bidang Ketenteraman dan Ketertiban Umum';
  let formPassword = '';
  let formActive = true;

  let deleteUserModalOpen = false;
  let deleteUserTarget: User | null = null;
  let deleteUserError = '';

  function openAddUserModal() {
    isNewUser = true;
    formUserId = 0;
    formName = '';
    formUsername = '';
    formEmail = '';
    formRole = 'STAFF_SATPOL';
    formBidang = $organizationStore[1]?.position_title || 'Bidang Ketenteraman dan Ketertiban Umum';
    formPassword = '';
    formActive = true;
    userFormError = '';
    userModalOpen = true;
  }

  function openEditUserModal(u: User) {
    isNewUser = false;
    formUserId = u.id;
    formName = u.name;
    formUsername = u.username;
    formEmail = u.email;
    formRole = u.role;
    formBidang = u.bidang || '';
    formPassword = '';
    formActive = u.is_active;
    userFormError = '';
    userModalOpen = true;
  }

  async function handleSaveUser() {
    userFormError = '';
    const cleanUsername = formUsername.trim().toLowerCase().replace(/\s+/g, '');
    if (!formName.trim() || !cleanUsername || !formEmail.trim()) {
      userFormError = 'Nama lengkap, username, dan email wajib diisi.';
      return;
    }
    const duplicate = $usersStore.find(
      u => u.username.toLowerCase() === cleanUsername && (isNewUser || u.id !== formUserId)
    );
    if (duplicate) {
      userFormError = `Username "${cleanUsername}" sudah terdaftar.`;
      return;
    }

    isSavingUser = true;
    const payload: User = {
      id: isNewUser ? Date.now() : formUserId,
      name: formName.trim(),
      username: cleanUsername,
      email: formEmail.trim(),
      role: formRole,
      bidang: formBidang.trim() || 'Sekretariat',
      is_active: formActive
    };

    try {
      await saveUser(payload, isNewUser, formPassword.trim() || undefined);
      logAudit(
        currentUser?.username || 'superadmin',
        isNewUser ? 'CREATE_USER' : 'UPDATE_USER',
        `User #${payload.username}`,
        `${isNewUser ? 'Menambahkan' : 'Memperbarui'} akun role ${payload.role} (${payload.name})`
      );
      showToast(`Akun "${payload.username}" berhasil ${isNewUser ? 'ditambahkan' : 'diperbarui'}.`);
      userModalOpen = false;
    } finally {
      isSavingUser = false;
    }
  }

  async function handleToggleUserStatus(u: User) {
    const nextStatus = !u.is_active;
    if (!nextStatus && currentUser && currentUser.id === u.id) {
      showToast('Tidak dapat menonaktifkan akun yang sedang Anda gunakan saat ini.');
      return;
    }
    await setUserActive(u.id, nextStatus);
    logAudit(
      currentUser?.username || 'superadmin',
      nextStatus ? 'ACTIVATE_USER' : 'SUSPEND_USER',
      `User #${u.username}`,
      `Mengubah status akun menjadi ${nextStatus ? 'Aktif' : 'Nonaktif'}`
    );
    showToast(`Status akun "${u.username}" diubah menjadi ${nextStatus ? 'Aktif' : 'Nonaktif'}.`);
  }

  function askDeleteUser(u: User) {
    deleteUserTarget = u;
    deleteUserError = '';
    const activeSuperadmins = $usersStore.filter(x => x.role === 'SUPERADMIN_DISKOMINFO' && x.id !== u.id);
    if (u.role === 'SUPERADMIN_DISKOMINFO' && activeSuperadmins.length === 0) {
      deleteUserError = 'Minimal harus tersisa satu akun Superadmin dalam sistem.';
    } else if (currentUser && currentUser.id === u.id) {
      deleteUserError = 'Tidak dapat menghapus akun yang sedang aktif masuk.';
    }
    deleteUserModalOpen = true;
  }

  async function confirmDeleteUser() {
    if (!deleteUserTarget || deleteUserError) return;
    const target = deleteUserTarget;
    await deleteUser(target.id);
    logAudit(
      currentUser?.username || 'superadmin',
      'DELETE_USER',
      `User #${target.username}`,
      `Menghapus akun pengguna ${target.name}`
    );
    showToast(`Akun "${target.username}" berhasil dihapus.`);
    deleteUserModalOpen = false;
    deleteUserTarget = null;
  }

  // ==========================================
  // 2. KONFIGURASI PORTAL (SITE SETTINGS)
  // ==========================================
  let currentSettings: SiteSettings = { ...$siteSettingsStore };
  let settingsInitialized = false;
  $: if ($siteSettingsStore && !settingsInitialized) {
    currentSettings = { ...$siteSettingsStore };
    settingsInitialized = true;
  }

  let isSavingSettings = false;
  function handleSaveSettings() {
    isSavingSettings = true;
    siteSettingsStore.set({ ...currentSettings });
    logAudit(
      currentUser?.username || 'superadmin',
      'UPDATE_CMS_SETTINGS',
      'SiteSettings',
      'Memperbarui konfigurasi identitas dan kontak resmi portal'
    );
    setTimeout(() => {
      isSavingSettings = false;
      showToast('Konfigurasi portal berhasil disimpan dan disinkronkan ke portal publik.');
    }, 250);
  }

  function roleBadgeStyle(role: UserRole) {
    if (role === 'SUPERADMIN_DISKOMINFO') return 'bg-[#1a120c] text-[#F08300] border-[#3d261a]';
    if (role === 'ADMIN_SATPOL') return 'bg-orange-50 text-orange-900 border-orange-200';
    return 'bg-[#f9f4e8] text-[#3d261a] border-[#eee9db]';
  }
</script>

<svelte:head>
  <title>Panel Superadmin &amp; Manajemen Akses | Satpol PP Sulsel</title>
</svelte:head>

<DashboardShell
  portalLabel="Superadmin Sistem"
  portalTag="DISKOMINFO & RBAC"
  {navItems}
  bind:activeTab
  pageTitle={tabMeta[activeTab].title}
  pageDescription={tabMeta[activeTab].desc}
  user={currentUser}
  quickSwitchHref="/admin/dashboard"
  quickSwitchLabel="Panel Admin Satpol"
>
  <svelte:fragment slot="header-actions">
    <a
      href="/admin/dashboard"
      class="hidden sm:inline-flex items-center gap-1.5 px-3.5 py-2 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] text-xs font-bold border border-[#eee9db] transition"
    >
      Meja Admin Satpol
    </a>
  </svelte:fragment>

  <!-- Ringkasan Metrik Superadmin -->
  <div class="grid grid-cols-2 lg:grid-cols-4 gap-4">
    <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs">
      <div class="text-[11px] font-bold text-[#3d261a]/75 uppercase">Total Akun RBAC</div>
      <div class="text-2xl md:text-3xl font-black text-[#1a120c] mt-1.5">{$usersStore.length}</div>
      <div class="text-[11px] text-[#3d261a]/65 mt-0.5">
        {$usersStore.filter(u => u.is_active).length} akun aktif
      </div>
    </div>
    <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs">
      <div class="text-[11px] font-bold text-[#F08300] uppercase">Simpul Struktur SOTK</div>
      <div class="text-2xl md:text-3xl font-black text-[#1a120c] mt-1.5">{$organizationStore.length}</div>
      <div class="text-[11px] text-[#3d261a]/65 mt-0.5">Unit kerja dinamis</div>
    </div>
    <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs">
      <div class="text-[11px] font-bold text-[#3d261a]/75 uppercase">Aduan SIAP TERTIB</div>
      <div class="text-2xl md:text-3xl font-black text-[#1a120c] mt-1.5">{$complaintsStore.length}</div>
      <div class="text-[11px] text-[#3d261a]/65 mt-0.5">24 kab/kota terhubung</div>
    </div>
    <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs">
      <div class="text-[11px] font-bold text-emerald-800 uppercase">Catatan Audit Trail</div>
      <div class="text-2xl md:text-3xl font-black text-emerald-700 mt-1.5">{$auditLogsStore.length}</div>
      <div class="text-[11px] text-emerald-700/80 mt-0.5">Log keamanan tercatat</div>
    </div>
  </div>

  {#if toastMessage}
    <div
      role="status"
      class="flex items-center justify-between gap-3 px-4 py-3 rounded-2xl bg-emerald-50 border border-emerald-200 text-emerald-900 text-xs font-bold shadow-xs"
    >
      <div class="flex items-center gap-2">
        <CheckCircle2 class="w-4 h-4 text-emerald-600 shrink-0" />
        <span>{toastMessage}</span>
      </div>
      <button type="button" on:click={() => (toastMessage = '')} class="text-xs underline opacity-75 hover:opacity-100">
        Tutup
      </button>
    </div>
  {/if}

  <!-- TAB 1: MANAJEMEN PENGGUNA -->
  {#if activeTab === 'users'}
    <DataTable
      items={$usersStore}
      searchFields={['name', 'username', 'email', 'role', 'bidang']}
      searchPlaceholder="Cari nama, username, bidang, atau role pengguna..."
      pageSize={10}
      pageSizeOptions={[5, 10, 25]}
      title="Daftar Akun Operator &amp; Administrator"
      description="Tambah, ubah peran, aktifkan/nonaktifkan, atau hapus akun pengelola portal."
      emptyMessage="Tidak ada akun pengguna yang sesuai dengan pencarian."
    >
      <svelte:fragment slot="header-actions">
        <button
          type="button"
          on:click={openAddUserModal}
          class="inline-flex items-center gap-1.5 bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold px-4 py-2 rounded-xl transition shadow-xs"
        >
          <Plus class="w-4 h-4" />
          Tambah Akun Baru
        </button>
      </svelte:fragment>

      <tr slot="head">
        <th class="px-4 py-3.5">Nama &amp; Email</th>
        <th class="px-4 py-3.5">Username</th>
        <th class="px-4 py-3.5">Peran (Role)</th>
        <th class="px-4 py-3.5">Unit / Bidang</th>
        <th class="px-4 py-3.5">Status</th>
        <th class="px-4 py-3.5 text-right">Aksi</th>
      </tr>

      <tr slot="row" let:item class="hover:bg-[#f9f4e8]/55 transition">
        <td class="px-4 py-3.5 align-top">
          <div class="font-bold text-xs text-[#1a120c]">{item.name}</div>
          <div class="text-[11px] text-[#3d261a]/70">{item.email}</div>
        </td>
        <td class="px-4 py-3.5 align-top font-mono text-xs font-bold text-[#3d261a]">@{item.username}</td>
        <td class="px-4 py-3.5 align-top">
          <span class="inline-block px-2.5 py-1 rounded-lg border text-[10px] font-bold uppercase {roleBadgeStyle(item.role)}">
            {item.role}
          </span>
        </td>
        <td class="px-4 py-3.5 align-top text-xs text-[#3d261a]">{item.bidang || '-'}</td>
        <td class="px-4 py-3.5 align-top">
          {#if item.is_active}
            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg bg-emerald-50 border border-emerald-200 text-emerald-800 text-[10px] font-bold">
              Aktif
            </span>
          {:else}
            <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg bg-red-50 border border-red-200 text-red-700 text-[10px] font-bold">
              Nonaktif
            </span>
          {/if}
        </td>
        <td class="px-4 py-3.5 align-top text-right whitespace-nowrap">
          <div class="inline-flex items-center gap-1.5">
            <button
              type="button"
              on:click={() => openEditUserModal(item)}
              class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] text-xs font-bold border border-[#eee9db] transition"
            >
              <Edit3 class="w-3.5 h-3.5" />
              Edit
            </button>
            <button
              type="button"
              on:click={() => handleToggleUserStatus(item)}
              title={item.is_active ? 'Nonaktifkan akun' : 'Aktifkan akun'}
              class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl border text-xs font-bold transition {item.is_active ? 'bg-amber-50 hover:bg-amber-100 text-amber-900 border-amber-200' : 'bg-emerald-50 hover:bg-emerald-100 text-emerald-800 border-emerald-200'}"
            >
              {#if item.is_active}
                <UserX class="w-3.5 h-3.5" />
                Nonaktifkan
              {:else}
                <UserCheck class="w-3.5 h-3.5" />
                Aktifkan
              {/if}
            </button>
            <button
              type="button"
              on:click={() => askDeleteUser(item)}
              aria-label="Hapus akun {item.username}"
              class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-red-50 hover:bg-red-100 text-red-700 text-xs font-bold border border-red-200 transition"
            >
              <Trash2 class="w-3.5 h-3.5" />
            </button>
          </div>
        </td>
      </tr>
    </DataTable>
  {/if}

  <!-- TAB 2: STRUKTUR ORGANISASI DINAMIS -->
  {#if activeTab === 'organization'}
    <OrganizationManager actorName={currentUser?.username || 'superadmin'} />
  {/if}

  <!-- TAB 3: KONFIGURASI PORTAL -->
  {#if activeTab === 'settings'}
    <div class="bg-white border border-[#eee9db] rounded-2xl shadow-xs p-6 md:p-8 space-y-6">
      <div class="border-b border-[#eee9db] pb-4">
        <h2 class="text-base font-black text-[#1a120c]">Identitas Instansi &amp; Kontak Portal</h2>
        <p class="text-xs text-[#3d261a]/75 mt-1">
          Perubahan langsung memperbarui informasi resmi di seluruh halaman publik.
        </p>
      </div>

      <form on:submit|preventDefault={handleSaveSettings} class="space-y-5">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
          <div>
            <label for="sup-app-name" class="block text-xs font-bold text-[#1a120c] mb-1.5">Judul Portal / Instansi</label>
            <input id="sup-app-name" type="text" bind:value={currentSettings.app_name} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" />
          </div>
          <div>
            <label for="sup-kasatpol-name" class="block text-xs font-bold text-[#1a120c] mb-1.5">Nama Kepala Satpol PP</label>
            <input id="sup-kasatpol-name" type="text" bind:value={currentSettings.kasatpol_name} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" />
          </div>
          <div>
            <label for="sup-call-center" class="block text-xs font-bold text-[#1a120c] mb-1.5">Telepon / Call Center</label>
            <input id="sup-call-center" type="text" bind:value={currentSettings.call_center} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" />
          </div>
          <div>
            <label for="sup-hotline" class="block text-xs font-bold text-[#1a120c] mb-1.5">Hotline Aduan Cepat</label>
            <input id="sup-hotline" type="text" bind:value={currentSettings.emergency_hotline} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" />
          </div>
          <div>
            <label for="sup-email" class="block text-xs font-bold text-[#1a120c] mb-1.5">Email Resmi Instansi</label>
            <input id="sup-email" type="email" bind:value={currentSettings.email_contact} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" />
          </div>
          <div>
            <label for="sup-address" class="block text-xs font-bold text-[#1a120c] mb-1.5">Alamat Markas Komando</label>
            <input id="sup-address" type="text" bind:value={currentSettings.office_address} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none" />
          </div>
          <div class="md:col-span-2">
            <label for="sup-tagline" class="block text-xs font-bold text-[#1a120c] mb-1.5">Semboyan / Tagline Resmi</label>
            <textarea id="sup-tagline" rows="2" bind:value={currentSettings.agency_tagline} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"></textarea>
          </div>
          <div class="md:col-span-2">
            <label for="sup-quote" class="block text-xs font-bold text-[#1a120c] mb-1.5">Sambutan Kepala Satpol PP di Beranda</label>
            <textarea id="sup-quote" rows="3" bind:value={currentSettings.kasatpol_quote} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"></textarea>
          </div>
        </div>

        <div class="pt-4 border-t border-[#eee9db] flex justify-end">
          <button
            type="submit"
            disabled={isSavingSettings}
            class="bg-[#F08300] hover:bg-[#d97300] text-white font-bold text-xs px-6 py-2.5 rounded-xl transition disabled:opacity-60 shadow-sm"
          >
            {isSavingSettings ? 'Menyimpan...' : 'Simpan Konfigurasi Portal'}
          </button>
        </div>
      </form>
    </div>
  {/if}

  <!-- TAB 4: 24 KABUPATEN/KOTA -->
  {#if activeTab === 'regencies'}
    <div class="bg-white border border-[#eee9db] rounded-2xl shadow-xs p-6 md:p-8 space-y-5">
      <div class="border-b border-[#eee9db] pb-4">
        <h2 class="text-base font-black text-[#1a120c]">Wilayah Administratif 24 Kabupaten / Kota</h2>
        <p class="text-xs text-[#3d261a]/75 mt-1">
          Daftar kode referensi Kemendagri Sulawesi Selatan yang digunakan pada formulir aduan SIAP TERTIB.
        </p>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-3.5">
        {#each defaultRegencies as reg}
          <div class="border border-[#eee9db] rounded-2xl p-4 bg-[#f9f4e8]/55 hover:border-[#F08300] transition">
            <div class="flex justify-between items-center text-[10px] text-[#3d261a]/70">
              <span class="font-mono font-bold">KODE: {reg.id}</span>
              <span class="font-bold text-emerald-700 uppercase">Aktif</span>
            </div>
            <div class="font-black text-xs text-[#1a120c] mt-1.5">{reg.name}</div>
            <div class="text-[11px] text-[#3d261a]/70 mt-0.5">Yurisdiksi Prov. Sulsel</div>
          </div>
        {/each}
      </div>
    </div>
  {/if}

  <!-- TAB 5: AUDIT TRAIL (DATATABLE REUSABLE) -->
  {#if activeTab === 'audit'}
    <DataTable
      items={$auditLogsStore}
      searchFields={['user_name', 'action', 'resource', 'details', 'ip_address']}
      searchPlaceholder="Cari operator, aksi, tiket/entitas, atau keterangan..."
      pageSize={10}
      pageSizeOptions={[10, 25, 50]}
      title="Audit Trail &amp; Log Keamanan"
      description="Catatan kronologis pembukaan masking, tindak lanjut SOP, dan administrasi akun."
      emptyMessage="Belum ada catatan aktivitas yang sesuai pencarian."
    >
      <tr slot="head">
        <th class="px-4 py-3.5">Waktu</th>
        <th class="px-4 py-3.5">Operator</th>
        <th class="px-4 py-3.5">Aksi Sistem</th>
        <th class="px-4 py-3.5">Entitas / Target</th>
        <th class="px-4 py-3.5">Rincian</th>
      </tr>

      <tr slot="row" let:item class="hover:bg-[#f9f4e8]/55 transition">
        <td class="px-4 py-3 text-[11px] font-mono text-[#3d261a]/75 whitespace-nowrap">{item.timestamp}</td>
        <td class="px-4 py-3 text-xs font-bold text-[#1a120c]">{item.user_name}</td>
        <td class="px-4 py-3">
          <span class="inline-block px-2.5 py-0.5 rounded-lg bg-[#f9f4e8] border border-[#eee9db] text-[#1a120c] font-mono text-[10px] font-bold">
            {item.action}
          </span>
        </td>
        <td class="px-4 py-3 text-xs font-semibold text-[#3d261a]">{item.resource}</td>
        <td class="px-4 py-3 text-xs text-[#3d261a]/80">{item.details}</td>
      </tr>
    </DataTable>
  {/if}

  <!-- TAB 6: TELEMETRI SISTEM -->
  {#if activeTab === 'server'}
    <div class="bg-white border border-[#eee9db] rounded-2xl shadow-xs p-6 md:p-8 space-y-6">
      <div class="border-b border-[#eee9db] pb-4">
        <h2 class="text-base font-black text-[#1a120c]">Status Runtime &amp; Arsitektur Sistem</h2>
        <p class="text-xs text-[#3d261a]/75 mt-1">
          Konfigurasi layanan frontend SvelteKit dan API Go (mode lokal / transisi basis data).
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div class="border border-[#eee9db] rounded-2xl p-5 bg-[#f9f4e8]/55 space-y-2">
          <div class="text-[10px] font-bold uppercase tracking-wider text-[#3d261a]/70">Frontend Portal</div>
          <div class="text-sm font-black text-[#1a120c]">SvelteKit + Tailwind CSS</div>
          <div class="text-xs text-[#3d261a]/80 space-y-1">
            <div>Dev Server: <span class="font-mono font-bold">Port 3000</span></div>
            <div>Proxy API: <span class="font-mono font-bold">/api &rarr; :8080</span></div>
          </div>
        </div>

        <div class="border border-[#eee9db] rounded-2xl p-5 bg-[#f9f4e8]/55 space-y-2">
          <div class="text-[10px] font-bold uppercase tracking-wider text-[#3d261a]/70">Backend &amp; Mailer</div>
          <div class="text-sm font-black text-[#1a120c]">Go net/http &amp; net/smtp</div>
          <div class="text-xs text-[#3d261a]/80 space-y-1">
            <div>Masking Publik: <span class="font-bold text-emerald-700">Selalu Aktif</span></div>
            <div>Email Pelapor: <span class="font-bold">SMTP / Simulasi Lokal</span></div>
          </div>
        </div>

        <div class="border border-[#eee9db] rounded-2xl p-5 bg-[#f9f4e8]/55 space-y-2">
          <div class="text-[10px] font-bold uppercase tracking-wider text-[#3d261a]/70">Penyimpanan Data</div>
          <div class="text-sm font-black text-[#1a120c]">In-Memory + LocalStorage</div>
          <div class="text-xs text-[#3d261a]/80 space-y-1">
            <div>Integrasi PostgreSQL: <span class="font-bold text-[#F08300]">Ditahan (Tahap Berikutnya)</span></div>
            <div>Fallback Offline: <span class="font-bold text-emerald-700">Aktif</span></div>
          </div>
        </div>
      </div>
    </div>
  {/if}
</DashboardShell>

<!-- MODAL: TAMBAH / EDIT PENGGUNA RBAC -->
<Modal
  bind:open={userModalOpen}
  title={isNewUser ? 'Tambah Akun Operator / Administrator' : `Edit Akun @${formUsername}`}
  subtitle="Atur hak akses dan unit kerja personel."
  size="md"
>
  <form id="user-rbac-form" on:submit|preventDefault={handleSaveUser} class="space-y-4">
    {#if userFormError}
      <div role="alert" class="p-3 rounded-xl bg-red-50 border border-red-200 text-red-800 text-xs font-bold">
        {userFormError}
      </div>
    {/if}

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
      <div>
        <label for="user-name" class="block text-xs font-bold text-[#1a120c] mb-1">Nama Lengkap &amp; Gelar *</label>
        <input id="user-name" type="text" required bind:value={formName} placeholder="Andi Baso, S.STP" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
        <p class="field-error">Nama wajib diisi.</p>
      </div>
      <div>
        <label for="user-username" class="block text-xs font-bold text-[#1a120c] mb-1">Username Login *</label>
        <input id="user-username" type="text" required bind:value={formUsername} disabled={!isNewUser} placeholder="staf_perda" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white disabled:bg-[#f9f4e8]" />
        <p class="field-error">Username wajib diisi.</p>
      </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
      <div>
        <label for="user-email" class="block text-xs font-bold text-[#1a120c] mb-1">Email Dinas *</label>
        <input id="user-email" type="email" required bind:value={formEmail} placeholder="nama@sulselprov.go.id" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
        <p class="field-error">Format email tidak valid.</p>
      </div>
      <div>
        <label for="user-role" class="block text-xs font-bold text-[#1a120c] mb-1">Tingkat Akses (Role)</label>
        <select id="user-role" bind:value={formRole} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white font-bold">
          <option value="STAFF_SATPOL">STAFF SATPOL PP (Verifikator &amp; Humas)</option>
          <option value="ADMIN_SATPOL">ADMIN SATPOL (Pimpinan / Kasatpol)</option>
          <option value="SUPERADMIN_DISKOMINFO">SUPERADMIN DISKOMINFO</option>
        </select>
      </div>
    </div>

    <div>
      <label for="user-bidang" class="block text-xs font-bold text-[#1a120c] mb-1">Unit Kerja / Bidang Penugasan</label>
      <input id="user-bidang" type="text" list="rbac-org-units" bind:value={formBidang} placeholder="Contoh: Bidang Penegakan Perda" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
      <datalist id="rbac-org-units">
        {#each $organizationStore as org}
          <option value={org.position_title}></option>
        {/each}
      </datalist>
    </div>

    <div>
      <label for="user-password" class="block text-xs font-bold text-[#1a120c] mb-1">
        {isNewUser ? 'Kata Sandi Awal' : 'Kata Sandi Baru (Kosongkan jika tidak diubah)'}
      </label>
      <input id="user-password" type="password" bind:value={formPassword} placeholder={isNewUser ? 'Default demo: satpolpp2026' : 'Biarkan kosong untuk tetap'} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
    </div>

    <label class="flex items-center gap-2.5 p-3 rounded-xl bg-[#f9f4e8] border border-[#eee9db] cursor-pointer">
      <input type="checkbox" bind:checked={formActive} class="rounded text-[#F08300] focus:ring-[#F08300]" />
      <span class="text-xs font-bold text-[#1a120c]">Akun aktif dan diizinkan masuk ke panel dashboard</span>
    </label>
  </form>

  <svelte:fragment slot="footer" let:handleClose>
    <button
      type="button"
      on:click={handleClose}
      class="px-4 py-2 rounded-xl border border-[#eee9db] bg-white hover:bg-[#f9f4e8] text-xs font-bold text-[#3d261a] transition"
    >
      Batal
    </button>
    <button
      type="submit"
      form="user-rbac-form"
      disabled={isSavingUser}
      class="px-5 py-2 rounded-xl bg-[#F08300] hover:bg-[#d97300] disabled:opacity-60 text-white text-xs font-bold transition shadow-sm"
    >
      {isSavingUser ? 'Menyimpan...' : isNewUser ? 'Buat Akun' : 'Simpan Perubahan'}
    </button>
  </svelte:fragment>
</Modal>

<!-- MODAL: KONFIRMASI HAPUS PENGGUNA -->
<Modal
  bind:open={deleteUserModalOpen}
  title="Konfirmasi Hapus Pengguna"
  subtitle={deleteUserTarget ? `@${deleteUserTarget.username} — ${deleteUserTarget.name}` : ''}
  size="sm"
>
  {#if deleteUserError}
    <div role="alert" class="p-3.5 rounded-xl bg-amber-50 border border-amber-200 text-amber-900 text-xs font-medium">
      {deleteUserError}
    </div>
  {:else if deleteUserTarget}
    <p class="text-xs text-[#3d261a] leading-relaxed">
      Apakah Anda yakin ingin menghapus akun <strong>@{deleteUserTarget.username}</strong> ({deleteUserTarget.role})?
    </p>
  {/if}

  <svelte:fragment slot="footer" let:handleClose>
    <button
      type="button"
      on:click={handleClose}
      class="px-4 py-2 rounded-xl border border-[#eee9db] bg-white hover:bg-[#f9f4e8] text-xs font-bold text-[#3d261a] transition"
    >
      Tutup
    </button>
    {#if !deleteUserError}
      <button
        type="button"
        on:click={confirmDeleteUser}
        class="px-4 py-2 rounded-xl bg-red-600 hover:bg-red-700 text-white text-xs font-bold transition"
      >
        Ya, Hapus Akun
      </button>
    {/if}
  </svelte:fragment>
</Modal>
