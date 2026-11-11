<script lang="ts">
  import { onMount } from 'svelte';
  import { Building2, Edit3, Plus, RefreshCw, Trash2, Users } from 'lucide-svelte';
  import type { OrganizationNode } from '$lib/types';
  import { organizationStore, logAudit } from '$lib/stores/appState';
  import { fetchOrganization, saveOrganizationNode, deleteOrganizationNode } from '$lib/api/client';
  import DataTable from '$lib/components/ui/DataTable.svelte';
  import Modal from '$lib/components/ui/Modal.svelte';
  import FileUploader from '$lib/components/ui/FileUploader.svelte';

  export let actorName = 'admin';

  onMount(() => {
    fetchOrganization();
  });

  $: sortedNodes = [...$organizationStore].sort((a, b) =>
    a.level !== b.level ? a.level - b.level : a.sort_order - b.sort_order
  );

  // Modal state
  let modalOpen = false;
  let isNew = true;
  let editingId = 0;
  let formError = '';
  let isSaving = false;

  let orgCode = '';
  let parentCode = '';
  let positionTitle = '';
  let officialName = '';
  let nip = '';
  let rank = '';
  let description = '';
  let subUnitsText = '';
  let photoUrl = '';
  let level = 2;
  let sortOrder = 1;

  // Delete confirmation modal
  let deleteTarget: OrganizationNode | null = null;
  let deleteModalOpen = false;
  let deleteError = '';

  function generateOrgCode(existingList: OrganizationNode[]): string {
    let idx = existingList.length + 1;
    let candidate = `ORG-UNIT-${String(idx).padStart(2, '0')}`;
    const used = new Set(existingList.map(n => n.org_code.toUpperCase()));
    while (used.has(candidate)) {
      idx += 1;
      candidate = `ORG-UNIT-${String(idx).padStart(2, '0')}`;
    }
    return candidate;
  }

  function openCreateModal() {
    isNew = true;
    editingId = 0;
    formError = '';
    orgCode = generateOrgCode($organizationStore);
    parentCode = sortedNodes[0]?.org_code || '';
    positionTitle = '';
    officialName = '';
    nip = '';
    rank = 'Eselon IIIb';
    description = '';
    subUnitsText = '';
    photoUrl = '';
    level = 2;
    sortOrder = $organizationStore.length + 1;
    modalOpen = true;
  }

  function openEditModal(node: OrganizationNode) {
    isNew = false;
    editingId = node.id;
    formError = '';
    orgCode = node.org_code;
    parentCode = node.parent_code || '';
    positionTitle = node.position_title;
    officialName = node.official_name;
    nip = node.nip || '';
    rank = node.rank || '';
    description = node.description || '';
    subUnitsText = (node.sub_units || []).join('\n');
    photoUrl = node.photo_url || '';
    level = node.level;
    sortOrder = node.sort_order;
    modalOpen = true;
  }

  async function handleSave() {
    formError = '';
    const cleanCode = orgCode.trim().toUpperCase().replace(/\s+/g, '-');
    if (!cleanCode || cleanCode.length < 3) {
      formError = 'ID Khusus unit wajib diisi minimal 3 karakter (contoh: ORG-BID-04).';
      return;
    }
    if (!positionTitle.trim() || !officialName.trim()) {
      formError = 'Nama Jabatan/Unit dan Nama Pejabat wajib diisi.';
      return;
    }

    const duplicate = $organizationStore.find(
      n => n.org_code.toUpperCase() === cleanCode && (isNew || n.id !== editingId)
    );
    if (duplicate) {
      formError = `ID Khusus "${cleanCode}" sudah dipakai oleh unit "${duplicate.position_title}".`;
      return;
    }
    if (!isNew && parentCode === cleanCode) {
      formError = 'Unit tidak dapat menunjuk dirinya sendiri sebagai atasan langsung.';
      return;
    }

    isSaving = true;
    const subUnits = subUnitsText
      .split(/\r?\n/)
      .map(s => s.trim())
      .filter(Boolean);

    const payload: OrganizationNode = {
      id: isNew ? Date.now() : editingId,
      org_code: cleanCode,
      parent_code: parentCode.trim() ? parentCode.trim() : null,
      position_title: positionTitle.trim(),
      official_name: officialName.trim(),
      nip: nip.trim() || undefined,
      rank: rank.trim() || undefined,
      description: description.trim() || undefined,
      sub_units: subUnits,
      photo_url: photoUrl || undefined,
      level: Number(level) || 2,
      sort_order: Number(sortOrder) || 1
    };

    try {
      await saveOrganizationNode(payload, isNew);
      logAudit(
        actorName,
        isNew ? 'TAMBAH_STRUKTUR_ORG' : 'UBAH_STRUKTUR_ORG',
        `${payload.org_code} (${payload.position_title})`,
        `${isNew ? 'Menambahkan' : 'Memperbarui'} simpul struktur organisasi SOTK oleh ${actorName}`
      );
      modalOpen = false;
    } finally {
      isSaving = false;
    }
  }

  function askDelete(node: OrganizationNode) {
    deleteTarget = node;
    deleteError = '';
    const childrenCount = $organizationStore.filter(n => n.parent_code === node.org_code).length;
    if (childrenCount > 0) {
      deleteError = `Unit ini membawahi ${childrenCount} sub-unit aktif. Ubah atasan sub-unit tersebut terlebih dahulu atau hapus dari level terbawah.`;
    }
    deleteModalOpen = true;
  }

  async function confirmDelete() {
    if (!deleteTarget || deleteError) return;
    const target = deleteTarget;
    await deleteOrganizationNode(target.id);
    logAudit(
      actorName,
      'HAPUS_STRUKTUR_ORG',
      `${target.org_code} (${target.position_title})`,
      `Menghapus simpul struktur organisasi ${target.position_title}`
    );
    deleteModalOpen = false;
    deleteTarget = null;
  }

  function findParentTitle(code: string | null) {
    if (!code) return 'Pimpinan Tertinggi';
    const found = $organizationStore.find(n => n.org_code === code);
    return found ? `${found.org_code} — ${found.position_title}` : code;
  }
</script>

<div class="space-y-4">
  <div class="bg-white border border-[#eee9db] rounded-2xl p-5 flex flex-col sm:flex-row sm:items-center justify-between gap-4 shadow-xs">
    <div class="flex items-start gap-3">
      <div class="w-10 h-10 rounded-xl bg-[#f9f4e8] border border-[#eee9db] text-[#F08300] flex items-center justify-center shrink-0">
        <Building2 class="w-5 h-5" />
      </div>
      <div>
        <h2 class="text-sm md:text-base font-black text-[#1a120c]">Struktur Organisasi Dinamis (SOTK)</h2>
        <p class="text-xs text-[#3d261a]/75 mt-0.5">
          Kelola hierarki jabatan, ID khusus unit (<code class="font-mono bg-[#f9f4e8] px-1.5 py-0.5 rounded-md">ORG-*</code>), sub-unit seksi, dan data pejabat yang tampil di halaman Profil publik.
        </p>
      </div>
    </div>
    <button
      type="button"
      on:click={openCreateModal}
      class="inline-flex items-center justify-center gap-2 px-4 py-2.5 rounded-xl bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold shadow-sm transition shrink-0"
    >
      <Plus class="w-4 h-4" />
      Tambah Unit / Pejabat
    </button>
  </div>

  <DataTable
    items={sortedNodes}
    searchFields={['org_code', 'position_title', 'official_name', 'rank', 'parent_code']}
    searchPlaceholder="Cari ID Khusus (ORG-...), jabatan, atau nama pejabat..."
    pageSize={8}
    emptyMessage="Belum ada data unit struktur organisasi."
  >
    <tr slot="head">
      <th class="px-5 py-3.5 text-left">ID Khusus &amp; Level</th>
      <th class="px-5 py-3.5 text-left">Jabatan / Unit Kerja</th>
      <th class="px-5 py-3.5 text-left">Pejabat &amp; Eselon</th>
      <th class="px-5 py-3.5 text-left">Atasan Langsung</th>
      <th class="px-5 py-3.5 text-right">Aksi</th>
    </tr>

    <tr slot="row" let:item class="hover:bg-[#f9f4e8]/55 transition">
      <td class="px-5 py-4 align-top whitespace-nowrap">
        <span class="inline-flex items-center px-2.5 py-1 rounded-lg bg-[#1a120c] text-[#F08300] font-mono text-xs font-black tracking-wide">
          {item.org_code}
        </span>
        <div class="text-[10px] font-bold text-[#3d261a]/70 mt-1">
          Level {item.level} • Urutan {item.sort_order}
        </div>
      </td>

      <td class="px-5 py-4 align-top max-w-xs">
        <div class="text-xs font-black text-[#1a120c]">{item.position_title}</div>
        {#if item.description}
          <p class="text-[11px] text-[#3d261a]/75 line-clamp-2 mt-1">{item.description}</p>
        {/if}
        {#if item.sub_units && item.sub_units.length > 0}
          <div class="mt-2 flex flex-wrap gap-1">
            {#each item.sub_units as sub}
              <span class="inline-block px-2 py-0.5 rounded-lg bg-[#f9f4e8] border border-[#eee9db] text-[10px] font-medium text-[#3d261a]">
                {sub}
              </span>
            {/each}
          </div>
        {/if}
      </td>

      <td class="px-5 py-4 align-top">
        <div class="flex items-center gap-2.5">
          {#if item.photo_url}
            <img src={item.photo_url} alt={item.official_name} class="w-9 h-9 rounded-xl object-cover border border-[#eee9db] shrink-0" />
          {:else}
            <div class="w-9 h-9 rounded-xl bg-[#f9f4e8] border border-[#eee9db] text-[#3d261a] flex items-center justify-center shrink-0">
              <Users class="w-4 h-4" />
            </div>
          {/if}
          <div>
            <div class="text-xs font-bold text-[#1a120c]">{item.official_name}</div>
            <div class="text-[11px] text-[#3d261a]/70">
              {item.rank || 'Struktural'}{item.nip ? ` • NIP ${item.nip}` : ''}
            </div>
          </div>
        </div>
      </td>

      <td class="px-5 py-4 align-top text-xs text-[#3d261a]">
        <span class="inline-block px-2.5 py-1 rounded-xl bg-[#f9f4e8] border border-[#eee9db] font-medium">
          {findParentTitle(item.parent_code)}
        </span>
      </td>

      <td class="px-5 py-4 align-top text-right whitespace-nowrap">
        <div class="inline-flex items-center gap-1.5">
          <button
            type="button"
            on:click={() => openEditModal(item)}
            title="Edit unit {item.org_code}"
            class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] text-xs font-bold border border-[#eee9db] transition"
          >
            <Edit3 class="w-3.5 h-3.5" />
            Edit
          </button>
          <button
            type="button"
            on:click={() => askDelete(item)}
            title="Hapus unit {item.org_code}"
            class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-red-50 hover:bg-red-100 text-red-700 text-xs font-bold border border-red-200 transition"
          >
            <Trash2 class="w-3.5 h-3.5" />
          </button>
        </div>
      </td>
    </tr>
  </DataTable>
</div>

<!-- Modal Tambah / Edit Unit Organisasi -->
<Modal
  bind:open={modalOpen}
  title={isNew ? 'Tambah Simpul Struktur Organisasi' : `Edit Unit ${orgCode}`}
  subtitle="Perubahan langsung memperbarui tampilan Struktur & Pejabat di halaman Profil Publik."
  size="lg"
>
  <form id="org-node-form" on:submit|preventDefault={handleSave} class="space-y-5">
    {#if formError}
      <div role="alert" class="p-3.5 rounded-xl bg-red-50 border border-red-200 text-red-800 text-xs font-bold">
        {formError}
      </div>
    {/if}

    <fieldset class="space-y-4">
      <legend class="text-xs font-black uppercase tracking-wider text-[#3d261a] mb-2">Identitas Unit & Hierarki</legend>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div>
          <label for="org-code" class="block text-xs font-bold text-[#1a120c] mb-1.5">ID Khusus Unit *</label>
          <div class="flex gap-1.5">
            <input
              id="org-code"
              type="text"
              bind:value={orgCode}
              required
              minlength="3"
              placeholder="ORG-BID-04"
              class="w-full px-3 py-2 text-xs font-mono font-bold uppercase rounded-xl border border-[#eee9db] bg-[#f9f4e8]/50 text-[#1a120c] focus:outline-none focus:border-[#F08300]"
            />
            {#if isNew}
              <button
                type="button"
                on:click={() => (orgCode = generateOrgCode($organizationStore))}
                title="Buat ID Khusus otomatis"
                class="px-2.5 py-2 rounded-xl border border-[#eee9db] bg-[#f9f4e8] hover:bg-[#eee9db] text-[#3d261a] transition shrink-0"
              >
                <RefreshCw class="w-3.5 h-3.5" />
              </button>
            {/if}
          </div>
          <p class="field-error">ID Khusus minimal 3 karakter.</p>
        </div>

        <div class="md:col-span-2">
          <label for="org-parent" class="block text-xs font-bold text-[#1a120c] mb-1.5">Atasan Langsung</label>
          <select
            id="org-parent"
            bind:value={parentCode}
            class="w-full px-3 py-2 text-xs rounded-xl border border-[#eee9db] bg-white text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          >
            <option value="">-- Tidak Ada (Pimpinan Tertinggi / Root) --</option>
            {#each sortedNodes.filter(n => isNew || n.id !== editingId) as candidate}
              <option value={candidate.org_code}>{candidate.org_code} — {candidate.position_title}</option>
            {/each}
          </select>
        </div>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
        <div>
          <label for="org-level" class="block text-xs font-bold text-[#1a120c] mb-1.5">Level Hierarki</label>
          <select
            id="org-level"
            bind:value={level}
            class="w-full px-3 py-2 text-xs rounded-xl border border-[#eee9db] bg-white text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          >
            <option value={1}>Level 1 (Kepala Satuan)</option>
            <option value={2}>Level 2 (Sekretariat / Bidang)</option>
            <option value={3}>Level 3 (Seksi / Sub Bagian)</option>
            <option value={4}>Level 4 (Satuan Tugas / UPT)</option>
          </select>
        </div>
        <div>
          <label for="org-sort" class="block text-xs font-bold text-[#1a120c] mb-1.5">Urutan Tampil</label>
          <input
            id="org-sort"
            type="number"
            min="1"
            max="99"
            bind:value={sortOrder}
            class="w-full px-3 py-2 text-xs rounded-xl border border-[#eee9db] bg-white text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          />
        </div>
      </div>
    </fieldset>

    <fieldset class="space-y-4 pt-3 border-t border-[#eee9db]">
      <legend class="text-xs font-black uppercase tracking-wider text-[#3d261a] mb-2">Pejabat & Uraian Tugas</legend>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label for="org-title" class="block text-xs font-bold text-[#1a120c] mb-1.5">Nama Jabatan / Unit *</label>
          <input
            id="org-title"
            type="text"
            bind:value={positionTitle}
            required
            placeholder="Contoh: Bidang Penegakan Perda"
            class="w-full px-3.5 py-2 text-xs rounded-xl border border-[#eee9db] text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          />
          <p class="field-error">Nama jabatan/unit wajib diisi.</p>
        </div>
        <div>
          <label for="org-official" class="block text-xs font-bold text-[#1a120c] mb-1.5">Nama Pejabat Saat Ini *</label>
          <input
            id="org-official"
            type="text"
            bind:value={officialName}
            required
            placeholder="Contoh: Drs. Ahmad Hidayat, M.Si"
            class="w-full px-3.5 py-2 text-xs rounded-xl border border-[#eee9db] text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          />
          <p class="field-error">Nama pejabat wajib diisi.</p>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div>
          <label for="org-rank" class="block text-xs font-bold text-[#1a120c] mb-1.5">Eselon / Pangkat</label>
          <input
            id="org-rank"
            type="text"
            bind:value={rank}
            placeholder="Contoh: Eselon IIIb"
            class="w-full px-3.5 py-2 text-xs rounded-xl border border-[#eee9db] text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          />
        </div>
        <div>
          <label for="org-nip" class="block text-xs font-bold text-[#1a120c] mb-1.5">NIP (Opsional)</label>
          <input
            id="org-nip"
            type="text"
            bind:value={nip}
            placeholder="19780512 200312 1 004"
            class="w-full px-3.5 py-2 text-xs rounded-xl border border-[#eee9db] text-[#1a120c] focus:outline-none focus:border-[#F08300]"
          />
        </div>
      </div>

      <div>
        <label for="org-desc" class="block text-xs font-bold text-[#1a120c] mb-1.5">Ringkasan Tugas Pokok</label>
        <textarea
          id="org-desc"
          rows="2"
          bind:value={description}
          placeholder="Uraian singkat kewenangan dan fungsi unit kerja..."
          class="w-full px-3.5 py-2 text-xs rounded-xl border border-[#eee9db] text-[#1a120c] focus:outline-none focus:border-[#F08300]"
        ></textarea>
      </div>

      <div>
        <label for="org-subunits" class="block text-xs font-bold text-[#1a120c] mb-1.5">
          Sub Bagian / Seksi Bawahan <span class="font-normal text-[#3d261a]/60">(satu baris per seksi)</span>
        </label>
        <textarea
          id="org-subunits"
          rows="3"
          bind:value={subUnitsText}
          placeholder="Seksi Pengawasan dan Pembinaan&#10;Seksi Penindakan dan Penyidikan"
          class="w-full px-3.5 py-2 text-xs rounded-xl border border-[#eee9db] text-[#1a120c] focus:outline-none focus:border-[#F08300]"
        ></textarea>
      </div>

      <FileUploader
        label="Foto Profil Pejabat (Opsional)"
        description="Pilih foto formal pejabat atau biarkan kosong untuk ikon default."
        previewShape="square"
        value={photoUrl}
        on:change={e => (photoUrl = e.detail.url)}
        on:clear={() => (photoUrl = '')}
      />
    </fieldset>
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
      form="org-node-form"
      disabled={isSaving}
      class="px-5 py-2 rounded-xl bg-[#F08300] hover:bg-[#d97300] disabled:opacity-60 text-white text-xs font-bold transition shadow-sm"
    >
      {isSaving ? 'Menyimpan...' : isNew ? 'Tambah Unit' : 'Simpan Perubahan'}
    </button>
  </svelte:fragment>
</Modal>

<!-- Modal Konfirmasi Hapus -->
<Modal
  bind:open={deleteModalOpen}
  title="Konfirmasi Hapus Unit"
  subtitle={deleteTarget ? `${deleteTarget.org_code} — ${deleteTarget.position_title}` : ''}
  size="sm"
>
  {#if deleteError}
    <div role="alert" class="p-3.5 rounded-xl bg-amber-50 border border-amber-200 text-amber-900 text-xs font-medium leading-relaxed">
      {deleteError}
    </div>
  {:else if deleteTarget}
    <p class="text-xs text-[#3d261a] leading-relaxed">
      Apakah Anda yakin ingin menghapus unit <strong>{deleteTarget.position_title}</strong> (<code class="font-mono">{deleteTarget.org_code}</code>) dari struktur organisasi? Tindakan ini akan dicatat di log audit.
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
    {#if !deleteError}
      <button
        type="button"
        on:click={confirmDelete}
        class="px-4 py-2 rounded-xl bg-red-600 hover:bg-red-700 text-white text-xs font-bold transition"
      >
        Ya, Hapus Unit
      </button>
    {/if}
  </svelte:fragment>
</Modal>
