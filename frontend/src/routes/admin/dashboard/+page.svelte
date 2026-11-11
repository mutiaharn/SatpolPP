<script lang="ts">
  import { onMount } from 'svelte';
  import {
    LayoutDashboard,
    ShieldAlert,
    Newspaper,
    FileText,
    Image,
    Building2,
    Plus,
    Trash2,
    Eye,
    EyeOff,
    Mail,
    CheckCircle2,
    MapPin,
    Download,
    Edit3
  } from 'lucide-svelte';
  import {
    complaintsStore,
    articlesStore,
    documentsStore,
    galleryStore,
    organizationStore,
    logAudit
  } from '$lib/stores/appState';
  import {
    updateComplaintStatus,
    uploadEvidence,
    fetchOrganization
  } from '$lib/api/client';
  import { maskComplaintIdentity } from '$lib/utils/masking';
  import { getSession } from '$lib/utils/auth';
  import type {
    Complaint,
    ComplaintStatus,
    Article,
    PublicDocument,
    GalleryItem,
    User
  } from '$lib/types';
  import DashboardShell, { type DashboardNavItem } from '$lib/components/layout/DashboardShell.svelte';
  import RichTextEditor from '$lib/components/RichTextEditor.svelte';
  import DataTable from '$lib/components/ui/DataTable.svelte';
  import FileUploader from '$lib/components/ui/FileUploader.svelte';
  import Modal from '$lib/components/ui/Modal.svelte';
  import StatusBadge from '$lib/components/ui/StatusBadge.svelte';
  import OrganizationManager from '$lib/components/cms/OrganizationManager.svelte';

  type AdminTab = 'overview' | 'complaints' | 'articles' | 'documents' | 'gallery' | 'organization';
  let currentTab: AdminTab = 'complaints';
  let currentUser: User | null = null;
  let toastMessage = '';
  let toastType: 'success' | 'info' = 'success';

  function showToast(msg: string, type: 'success' | 'info' = 'success') {
    toastMessage = msg;
    toastType = type;
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

  // ==========================================
  // 1. MANAJEMEN ADUAN SIAP TERTIB
  // ==========================================
  let statusFilter: 'ALL' | ComplaintStatus = 'ALL';
  let selectedComplaint: Complaint | null = null;
  let isEditingComplaint = false;
  let isSavingComplaint = false;

  let editStatus: ComplaintStatus = 'MENUNGGU_VERIFIKASI';
  let editDisposedTo = '';
  let editActionNotes = '';
  let unmaskRequested = false;

  $: filteredComplaints =
    statusFilter === 'ALL'
      ? $complaintsStore
      : $complaintsStore.filter(c => c.status === statusFilter);

  $: totalCount = $complaintsStore.length;
  $: pendingCount = $complaintsStore.filter(c => c.status === 'MENUNGGU_VERIFIKASI').length;
  $: disposedCount = $complaintsStore.filter(c => c.status === 'DISPOSISI_BIDANG').length;
  $: inProgressCount = $complaintsStore.filter(c => c.status === 'PENINDAKAN_LAPANGAN').length;
  $: completedCount = $complaintsStore.filter(c => c.status === 'SELESAI').length;

  $: dispositionOptions = $organizationStore
    .filter(n => n.level >= 2)
    .map(n => n.position_title);

  $: navItems = [
    { id: 'overview', label: 'Ikhtisar Operasional', icon: LayoutDashboard, badge: totalCount, section: 'Operasional' },
    { id: 'complaints', label: 'Aduan SIAP TERTIB', icon: ShieldAlert, badge: pendingCount || undefined },
    { id: 'articles', label: 'Warta & Berita', icon: Newspaper, badge: $articlesStore.length, section: 'Publikasi & CMS' },
    { id: 'documents', label: 'Dokumen Publik', icon: FileText, badge: $documentsStore.length },
    { id: 'gallery', label: 'Galeri Multimedia', icon: Image, badge: $galleryStore.length },
    { id: 'organization', label: 'Struktur Organisasi', icon: Building2, badge: $organizationStore.length, section: 'Kelembagaan' }
  ] satisfies DashboardNavItem[];

  const tabMeta: Record<AdminTab, { title: string; desc: string }> = {
    overview: {
      title: 'Ikhtisar Operasional & CMS',
      desc: 'Ringkasan pengaduan masyarakat SIAP TERTIB dan pembaruan konten portal publik.'
    },
    complaints: {
      title: 'Meja Verifikasi Aduan SIAP TERTIB',
      desc: 'Verifikasi laporan warga, buka masking teraudit, dan kirim notifikasi progres ke pelapor.'
    },
    articles: {
      title: 'Redaksi Warta & Berita Resmi',
      desc: 'Tulis, kelola, dan publikasikan laporan kegiatan patroli dan rilis Humas.'
    },
    documents: {
      title: 'Arsip Dokumen Transparansi Publik',
      desc: 'Unggah berkas Renja, Renstra, DPA, LAKIP, Perda, dan SOP resmi.'
    },
    gallery: {
      title: 'Galeri Multimedia Dokumentasi',
      desc: 'Kelola foto apel kesiapsiagaan, penertiban, dan koordinasi lapangan.'
    },
    organization: {
      title: 'Manajemen Struktur Organisasi (SOTK)',
      desc: 'Atur simpul unit kerja, ID khusus (ORG-*), sub-unit seksi, dan data pejabat secara dinamis.'
    }
  };

  function openComplaintDetail(comp: Complaint) {
    selectedComplaint = comp;
    editStatus = comp.status;
    editDisposedTo = comp.disposed_to || dispositionOptions[0] || 'Bidang Ketenteraman dan Ketertiban Umum';
    editActionNotes = comp.action_notes || '';
    unmaskRequested = false;
    isEditingComplaint = true;
  }

  function handleUnmaskIdentity() {
    if (!selectedComplaint) return;
    unmaskRequested = true;
    logAudit(
      currentUser?.username || 'petugas',
      'BUKA_MASKING_ANONIM',
      `Tiket ${selectedComplaint.ticket_code}`,
      'Membuka data identitas terlindung pelapor atas wewenang petugas verifikator'
    );
  }

  async function saveComplaintUpdate() {
    if (!selectedComplaint) return;
    isSavingComplaint = true;
    const ticket = selectedComplaint.ticket_code;
    try {
      const res = await updateComplaintStatus(ticket, {
        status: editStatus,
        disposed_to: editDisposedTo.trim(),
        action_notes: editActionNotes.trim(),
        officer_name: currentUser?.name || currentUser?.username || 'Petugas Verifikator'
      });

      logAudit(
        currentUser?.username || 'petugas',
        'UPDATE_STATUS_ADUAN',
        `Tiket ${ticket}`,
        `Status SOP: ${editStatus}, Disposisi: ${editDisposedTo}`
      );

      const emailInfo = res.email?.to
        ? res.email.sent
          ? ` Notifikasi email terkirim ke ${res.email.to}.`
          : ` Simulasi notifikasi email disiapkan untuk ${res.email.to}.`
        : '';
      showToast(`Status tiket ${ticket} diperbarui.${emailInfo}`);
      isEditingComplaint = false;
      selectedComplaint = null;
    } finally {
      isSavingComplaint = false;
    }
  }

  // ==========================================
  // 2. CMS BERITA (WITH FILEUPLOADER)
  // ==========================================
  let isWritingArticle = false;
  let editingArticleId: number | null = null;
  let articleError = '';

  let articleTitle = '';
  let articleSlug = '';
  let articleCategory = 'Kegiatan';
  let articleSummary = '';
  let articleContent = '';
  let articleAuthor = 'Humas Satpol PP';
  let articleThumbnail = '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg';
  let articleFileName = '';

  function openNewArticleEditor() {
    editingArticleId = null;
    articleError = '';
    articleTitle = '';
    articleSlug = '';
    articleCategory = 'Kegiatan';
    articleSummary = '';
    articleContent = '';
    articleAuthor = currentUser?.name || 'Humas Satpol PP';
    articleThumbnail = '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg';
    articleFileName = '';
    isWritingArticle = true;
  }

  function openEditArticle(art: Article) {
    editingArticleId = art.id;
    articleError = '';
    articleTitle = art.title;
    articleSlug = art.slug;
    articleCategory = art.category;
    articleSummary = art.summary;
    articleContent = art.content;
    articleAuthor = art.author_name;
    articleThumbnail = art.thumbnail;
    articleFileName = '';
    isWritingArticle = true;
  }

  function handleTitleInput(e: Event) {
    const target = e.target as HTMLInputElement;
    articleTitle = target.value;
    if (!editingArticleId) {
      articleSlug = articleTitle
        .toLowerCase()
        .replace(/[^a-z0-9]+/g, '-')
        .replace(/(^-|-$)+/g, '');
    }
  }

  async function handleArticleThumbChange(e: CustomEvent<{ url: string; file: File; name: string }>) {
    articleFileName = e.detail.name;
    articleThumbnail = e.detail.url;
    const remoteUrl = await uploadEvidence(e.detail.file);
    if (remoteUrl) articleThumbnail = remoteUrl;
  }

  function renderFormattedHtml(text: string): string {
    if (!text) return '';
    if (text.trim().startsWith('<p>') || text.trim().startsWith('<h')) return text;
    const lines = text.split('\n');
    let html = '';
    let inOl = false;
    let inUl = false;

    for (const line of lines) {
      if (line.startsWith('## ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<h2 class="text-xl font-extrabold text-[#1a120c] mt-5 mb-2">${formatInline(line.slice(3))}</h2>`;
        continue;
      }
      if (line.startsWith('### ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<h3 class="text-base font-bold text-[#1a120c] mt-4 mb-1.5">${formatInline(line.slice(4))}</h3>`;
        continue;
      }
      if (line.startsWith('> ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<blockquote class="border-l-4 border-[#F08300] pl-4 italic text-[#3d261a] my-3 bg-[#f9f4e8] py-2 rounded-r-xl">${formatInline(line.slice(2))}</blockquote>`;
        continue;
      }
      const numMatch = line.match(/^\s*(\d+)\.\s+(.*)$/);
      if (numMatch) {
        if (inUl) { html += '</ul>'; inUl = false; }
        if (!inOl) { html += '<ol class="list-decimal pl-5 space-y-1 my-2 text-xs text-[#3d261a]">'; inOl = true; }
        html += `<li>${formatInline(numMatch[2])}</li>`;
        continue;
      }
      const bulletMatch = line.match(/^\s*[-*]\s+(.*)$/);
      if (bulletMatch) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (!inUl) { html += '<ul class="list-disc pl-5 space-y-1 my-2 text-xs text-[#3d261a]">'; inUl = true; }
        html += `<li>${formatInline(bulletMatch[1])}</li>`;
        continue;
      }
      if (inOl) { html += '</ol>'; inOl = false; }
      if (inUl) { html += '</ul>'; inUl = false; }
      if (line.trim() !== '') {
        html += `<p class="text-xs text-[#3d261a] leading-relaxed my-2.5">${formatInline(line)}</p>`;
      }
    }
    if (inOl) html += '</ol>';
    if (inUl) html += '</ul>';
    return html;
  }

  function formatInline(str: string): string {
    return str
      .replace(/\*\*(.*?)\*\*/g, '<strong>$1</strong>')
      .replace(/\*(.*?)\*/g, '<em>$1</em>')
      .replace(/<u>(.*?)<\/u>/g, '<u>$1</u>')
      .replace(/\[(.*?)\]\((.*?)\)/g, '<a href="$2" target="_blank" class="text-[#F08300] underline font-semibold">$1</a>');
  }

  function handlePublishArticle() {
    articleError = '';
    if (!articleTitle.trim() || !articleSummary.trim()) {
      articleError = 'Judul berita dan ringkasan paragraf pembuka wajib diisi.';
      return;
    }
    const finalSlug = articleSlug || articleTitle.toLowerCase().replace(/[^a-z0-9]+/g, '-');
    const parsedHtml = renderFormattedHtml(articleContent);

    if (editingArticleId) {
      articlesStore.update(items =>
        items.map(a =>
          a.id === editingArticleId
            ? {
                ...a,
                title: articleTitle.trim(),
                slug: finalSlug,
                summary: articleSummary.trim(),
                content: parsedHtml || `<p>${articleSummary.trim()}</p>`,
                thumbnail: articleThumbnail,
                category: articleCategory,
                author_name: articleAuthor
              }
            : a
        )
      );
      logAudit(currentUser?.username || 'humas', 'EDIT_ARTICLE', `Berita: ${articleTitle}`, `Memperbarui konten berita`);
      showToast('Perubahan warta berita berhasil disimpan.');
    } else {
      const newArticle: Article = {
        id: Date.now(),
        title: articleTitle.trim(),
        slug: finalSlug,
        summary: articleSummary.trim(),
        content: parsedHtml || `<p>${articleSummary.trim()}</p>`,
        thumbnail: articleThumbnail,
        category: articleCategory,
        author_name: articleAuthor,
        view_count: 1,
        published_at: new Date().toISOString()
      };
      articlesStore.update(items => [newArticle, ...items]);
      logAudit(
        currentUser?.username || 'humas',
        'PUBLISH_ARTICLE',
        `Berita: ${newArticle.title}`,
        `Menerbitkan warta baru kategori ${newArticle.category}`
      );
      showToast('Warta berita berhasil dipublikasikan ke portal publik.');
    }

    isWritingArticle = false;
    editingArticleId = null;
  }

  // ==========================================
  // 3. CMS DOKUMEN PUBLIK (MODAL + FILEUPLOADER)
  // ==========================================
  let showDocModal = false;
  let docError = '';
  let docTitle = '';
  let docNumber = '';
  let docCategory: PublicDocument['category'] = 'renja';
  let docYear = 2026;
  let docSize = '1.8 MB';
  let docFileUrl = '/uploads/dokumen/dokumen_resmi.pdf';
  let uploadedFileName = '';

  function openDocModal() {
    docError = '';
    docTitle = '';
    docNumber = '';
    docCategory = 'renja';
    docYear = 2026;
    docSize = '1.8 MB';
    docFileUrl = '/uploads/dokumen/dokumen_resmi.pdf';
    uploadedFileName = '';
    showDocModal = true;
  }

  function handleDocFileChange(e: CustomEvent<{ url: string; file: File; name: string }>) {
    const file = e.detail.file;
    uploadedFileName = e.detail.name;
    if (!docTitle.trim()) {
      docTitle = file.name.replace(/\.[^/.]+$/, '').replace(/[-_]/g, ' ');
    }
    const sizeMb = (file.size / (1024 * 1024)).toFixed(1);
    docSize = sizeMb === '0.0' ? `${Math.max(1, Math.round(file.size / 1024))} KB` : `${sizeMb} MB`;
    docFileUrl = URL.createObjectURL(file);
  }

  function handleSaveDocument() {
    docError = '';
    if (!docTitle.trim()) {
      docError = 'Nama / Judul dokumen resmi wajib diisi.';
      return;
    }
    const newDoc: PublicDocument = {
      id: Date.now(),
      title: docTitle.trim(),
      doc_number: docNumber.trim() || `050/SATPOL/${docYear}`,
      category: docCategory,
      year: Number(docYear),
      file_url: docFileUrl,
      file_size: docSize
    };
    documentsStore.update(items => [newDoc, ...items]);
    logAudit(
      currentUser?.username || 'sekretariat',
      'UPLOAD_DOKUMEN_PUBLIK',
      `Dokumen: ${newDoc.title}`,
      `Mengunggah berkas transparansi publik kategori ${newDoc.category} TA ${newDoc.year}`
    );
    showToast(`Dokumen "${newDoc.title}" berhasil ditambahkan ke arsip publik.`);
    showDocModal = false;
  }

  // ==========================================
  // 4. CMS GALERI MULTIMEDIA (MODAL + FILEUPLOADER)
  // ==========================================
  let showGalleryModal = false;
  let galleryError = '';
  let galleryTitle = '';
  let galleryCategory = 'Penertiban';
  let galleryDate = new Date().toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' });
  let galleryImageUrl = '/uploads/foto_galeri/1687082076352158666_3553275628273619_3927930681691724075_n.jpg';
  let galleryFileName = '';

  function openGalleryModal() {
    galleryError = '';
    galleryTitle = '';
    galleryCategory = 'Penertiban';
    galleryImageUrl = '/uploads/foto_galeri/1687082076352158666_3553275628273619_3927930681691724075_n.jpg';
    galleryFileName = '';
    showGalleryModal = true;
  }

  async function handleGalleryFileChange(e: CustomEvent<{ url: string; file: File; name: string }>) {
    galleryFileName = e.detail.name;
    galleryImageUrl = e.detail.url;
    const remoteUrl = await uploadEvidence(e.detail.file);
    if (remoteUrl) galleryImageUrl = remoteUrl;
  }

  function handleSaveGallery() {
    galleryError = '';
    if (!galleryTitle.trim()) {
      galleryError = 'Judul atau keterangan dokumentasi kegiatan wajib diisi.';
      return;
    }
    const newItem: GalleryItem = {
      id: Date.now(),
      title: galleryTitle.trim(),
      category: galleryCategory,
      image_url: galleryImageUrl,
      date: galleryDate
    };
    galleryStore.update(items => [newItem, ...items]);
    logAudit(
      currentUser?.username || 'humas',
      'UPLOAD_GALERI_FOTO',
      `Foto: ${newItem.title}`,
      `Menambahkan foto dokumentasi kategori ${newItem.category}`
    );
    showToast('Foto dokumentasi kegiatan berhasil ditambahkan ke galeri.');
    showGalleryModal = false;
  }

  // ==========================================
  // 5. KONFIRMASI HAPUS REUSABLE MODAL
  // ==========================================
  let confirmDeleteOpen = false;
  let confirmDeleteType: 'article' | 'document' | 'gallery' = 'article';
  let confirmDeleteId = 0;
  let confirmDeleteLabel = '';

  function requestDelete(type: 'article' | 'document' | 'gallery', id: number, label: string) {
    confirmDeleteType = type;
    confirmDeleteId = id;
    confirmDeleteLabel = label;
    confirmDeleteOpen = true;
  }

  function executeDelete() {
    if (confirmDeleteType === 'article') {
      articlesStore.update(items => items.filter(a => a.id !== confirmDeleteId));
      logAudit(currentUser?.username || 'humas', 'HAPUS_ARTIKEL', confirmDeleteLabel, 'Menghapus warta berita');
    } else if (confirmDeleteType === 'document') {
      documentsStore.update(items => items.filter(d => d.id !== confirmDeleteId));
      logAudit(currentUser?.username || 'sekretariat', 'HAPUS_DOKUMEN', confirmDeleteLabel, 'Menghapus dokumen publik');
    } else {
      galleryStore.update(items => items.filter(g => g.id !== confirmDeleteId));
      logAudit(currentUser?.username || 'humas', 'HAPUS_GALERI', confirmDeleteLabel, 'Menghapus foto galeri');
    }
    showToast(`Item "${confirmDeleteLabel}" telah dihapus.`);
    confirmDeleteOpen = false;
  }
</script>

<svelte:head>
  <title>Panel Operasional Satpol PP &amp; CMS Terpadu</title>
</svelte:head>

<DashboardShell
  portalLabel="Satpol PP Sulsel"
  portalTag="SIAP TERTIB & CMS"
  {navItems}
  bind:activeTab={currentTab}
  pageTitle={tabMeta[currentTab].title}
  pageDescription={tabMeta[currentTab].desc}
  user={currentUser}
  quickSwitchHref={currentUser?.role === 'SUPERADMIN_DISKOMINFO' ? '/superadmin/dashboard' : ''}
  quickSwitchLabel="Panel Superadmin"
>
  <svelte:fragment slot="header-actions">
    <a
      href="/lapor"
      target="_blank"
      class="hidden sm:inline-flex items-center gap-1.5 px-3.5 py-2 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] text-xs font-bold border border-[#eee9db] transition"
    >
      Halaman Lapor Publik
    </a>
  </svelte:fragment>

  {#if toastMessage}
    <div
      role="status"
      class="flex items-center justify-between gap-3 px-4 py-3 rounded-2xl border shadow-xs text-xs font-bold {toastType === 'success' ? 'bg-emerald-50 border-emerald-200 text-emerald-900' : 'bg-[#f9f4e8] border-[#F08300]/40 text-[#1a120c]'}"
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

  <!-- TAB 1: IKHTISAR OPERASIONAL -->
  {#if currentTab === 'overview'}
    <div class="space-y-6">
      <div class="grid grid-cols-2 lg:grid-cols-5 gap-4">
        <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs">
          <div class="text-[11px] font-bold text-[#3d261a]/75 uppercase">Total Aduan</div>
          <div class="text-3xl font-black text-[#1a120c] mt-2">{totalCount}</div>
          <div class="text-[11px] text-[#3d261a]/60 mt-1">Registrasi SIAP TERTIB</div>
        </div>
        <div class="bg-white border border-amber-200 rounded-2xl p-5 shadow-xs">
          <div class="text-[11px] font-bold text-amber-800 uppercase">Menunggu Verifikasi</div>
          <div class="text-3xl font-black text-amber-700 mt-2">{pendingCount}</div>
          <div class="text-[11px] text-amber-700/80 mt-1">Perlu verifikasi awal</div>
        </div>
        <div class="bg-white border border-sky-200 rounded-2xl p-5 shadow-xs">
          <div class="text-[11px] font-bold text-sky-800 uppercase">Disposisi Bidang</div>
          <div class="text-3xl font-black text-sky-700 mt-2">{disposedCount}</div>
          <div class="text-[11px] text-sky-700/80 mt-1">Diteruskan ke unit</div>
        </div>
        <div class="bg-white border border-orange-200 rounded-2xl p-5 shadow-xs">
          <div class="text-[11px] font-bold text-[#F08300] uppercase">Penindakan Patroli</div>
          <div class="text-3xl font-black text-[#F08300] mt-2">{inProgressCount}</div>
          <div class="text-[11px] text-[#3d261a]/70 mt-1">Giat lapangan aktif</div>
        </div>
        <div class="bg-white border border-emerald-200 rounded-2xl p-5 shadow-xs col-span-2 lg:col-span-1">
          <div class="text-[11px] font-bold text-emerald-800 uppercase">Selesai (BAP)</div>
          <div class="text-3xl font-black text-emerald-700 mt-2">{completedCount}</div>
          <div class="text-[11px] text-emerald-700/80 mt-1">Tuntas terdokumentasi</div>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-5">
        <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs flex flex-col justify-between space-y-4">
          <div class="space-y-1.5">
            <div class="text-[11px] font-black text-[#F08300] uppercase tracking-wider">SOP SIAP TERTIB</div>
            <h2 class="text-sm font-black text-[#1a120c]">Meja Verifikasi Aduan</h2>
            <p class="text-xs text-[#3d261a]/75 leading-relaxed">
              Ada <strong>{pendingCount} laporan</strong> baru yang menunggu pengecekan dan disposisi.
            </p>
          </div>
          <button
            type="button"
            on:click={() => (currentTab = 'complaints')}
            class="w-full py-2.5 px-4 rounded-xl bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold transition"
          >
            Buka Meja Aduan
          </button>
        </div>

        <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs flex flex-col justify-between space-y-4">
          <div class="space-y-1.5">
            <div class="text-[11px] font-black text-[#3d261a] uppercase tracking-wider">Redaksi Humas</div>
            <h2 class="text-sm font-black text-[#1a120c]">Tulis Warta Berita</h2>
            <p class="text-xs text-[#3d261a]/75 leading-relaxed">
              Publikasikan kegiatan patroli penertiban dan pengumuman resmi Satpol PP.
            </p>
          </div>
          <button
            type="button"
            on:click={() => { currentTab = 'articles'; openNewArticleEditor(); }}
            class="w-full py-2.5 px-4 rounded-xl bg-[#3d261a] hover:bg-[#241710] text-white text-xs font-bold transition"
          >
            Tulis Berita Baru
          </button>
        </div>

        <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs flex flex-col justify-between space-y-4">
          <div class="space-y-1.5">
            <div class="text-[11px] font-black text-emerald-700 uppercase tracking-wider">Transparansi PPID</div>
            <h2 class="text-sm font-black text-[#1a120c]">Unggah Dokumen Publik</h2>
            <p class="text-xs text-[#3d261a]/75 leading-relaxed">
              Tambahkan Renja, Renstra, DPA, atau peraturan daerah ke katalog publik.
            </p>
          </div>
          <button
            type="button"
            on:click={() => { currentTab = 'documents'; openDocModal(); }}
            class="w-full py-2.5 px-4 rounded-xl bg-[#1a120c] hover:bg-[#3d261a] text-white text-xs font-bold transition"
          >
            Unggah Dokumen PDF
          </button>
        </div>

        <div class="bg-white border border-[#eee9db] rounded-2xl p-5 shadow-xs flex flex-col justify-between space-y-4">
          <div class="space-y-1.5">
            <div class="text-[11px] font-black text-[#F08300] uppercase tracking-wider">Kelembagaan Dinamis</div>
            <h2 class="text-sm font-black text-[#1a120c]">Struktur Organisasi</h2>
            <p class="text-xs text-[#3d261a]/75 leading-relaxed">
              Kelola { $organizationStore.length } simpul jabatan dan ID Khusus unit (<code class="font-mono">ORG-*</code>).
            </p>
          </div>
          <button
            type="button"
            on:click={() => (currentTab = 'organization')}
            class="w-full py-2.5 px-4 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] border border-[#eee9db] text-xs font-bold transition"
          >
            Kelola Struktur SOTK
          </button>
        </div>
      </div>
    </div>
  {/if}

  <!-- TAB 2: MANAJEMEN ADUAN SIAP TERTIB (FIXED: REAL COMPLAINTS DATATABLE) -->
  {#if currentTab === 'complaints'}
    <DataTable
      items={filteredComplaints}
      searchFields={['ticket_code', 'reporter_name', 'regency_name', 'district_name', 'violation_type', 'address_detail']}
      searchPlaceholder="Cari nomor tiket, kabupaten/kota, jenis pelanggaran (on-chain)..."
      pageSize={10}
      pageSizeOptions={[5, 10, 25, 50]}
      title="Daftar Pengaduan Masyarakat SIAP TERTIB"
      description="Identitas pelapor tetap disamarkan di tampilan daftar hingga dibuka pada log audit."
      emptyMessage="Tidak ada aduan yang cocok dengan filter status atau kata kunci."
    >
      <div slot="header-left" class="mt-3 flex flex-wrap gap-1.5">
        <button
          type="button"
          on:click={() => (statusFilter = 'ALL')}
          class="px-3 py-1.5 rounded-xl text-[11px] font-bold transition {statusFilter === 'ALL' ? 'bg-[#1a120c] text-white' : 'bg-[#f9f4e8] text-[#3d261a] hover:bg-[#eee9db]'}"
        >
          Semua ({totalCount})
        </button>
        <button
          type="button"
          on:click={() => (statusFilter = 'MENUNGGU_VERIFIKASI')}
          class="px-3 py-1.5 rounded-xl text-[11px] font-bold transition {statusFilter === 'MENUNGGU_VERIFIKASI' ? 'bg-[#F08300] text-white' : 'bg-amber-50 text-amber-900 hover:bg-amber-100'}"
        >
          Verifikasi ({pendingCount})
        </button>
        <button
          type="button"
          on:click={() => (statusFilter = 'DISPOSISI_BIDANG')}
          class="px-3 py-1.5 rounded-xl text-[11px] font-bold transition {statusFilter === 'DISPOSISI_BIDANG' ? 'bg-sky-700 text-white' : 'bg-sky-50 text-sky-900 hover:bg-sky-100'}"
        >
          Disposisi ({disposedCount})
        </button>
        <button
          type="button"
          on:click={() => (statusFilter = 'PENINDAKAN_LAPANGAN')}
          class="px-3 py-1.5 rounded-xl text-[11px] font-bold transition {statusFilter === 'PENINDAKAN_LAPANGAN' ? 'bg-[#3d261a] text-white' : 'bg-orange-50 text-orange-900 hover:bg-orange-100'}"
        >
          Penindakan ({inProgressCount})
        </button>
        <button
          type="button"
          on:click={() => (statusFilter = 'SELESAI')}
          class="px-3 py-1.5 rounded-xl text-[11px] font-bold transition {statusFilter === 'SELESAI' ? 'bg-emerald-700 text-white' : 'bg-emerald-50 text-emerald-900 hover:bg-emerald-100'}"
        >
          Selesai ({completedCount})
        </button>
      </div>

      <tr slot="head">
        <th class="px-4 py-3.5">Tiket &amp; Tanggal</th>
        <th class="px-4 py-3.5">Pelapor (Masked)</th>
        <th class="px-4 py-3.5">Lokasi &amp; Jenis Pelanggaran</th>
        <th class="px-4 py-3.5">Status SOP</th>
        <th class="px-4 py-3.5">Disposisi Unit</th>
        <th class="px-4 py-3.5 text-right">Aksi</th>
      </tr>

      <tr slot="row" let:item class="hover:bg-[#f9f4e8]/55 transition">
        <td class="px-4 py-3.5 align-top whitespace-nowrap">
          <div class="font-mono text-xs font-black text-[#1a120c]">{item.ticket_code}</div>
          <div class="text-[11px] text-[#3d261a]/65 mt-0.5">
            {new Date(item.created_at).toLocaleDateString('id-ID', { day: '2-digit', month: 'short', year: 'numeric' })}
          </div>
        </td>

        <td class="px-4 py-3.5 align-top">
          <div class="font-bold text-xs text-[#1a120c]">
            {maskComplaintIdentity(item).name}
          </div>
          <div class="text-[11px] font-mono text-[#3d261a]/70">
            {maskComplaintIdentity(item).phone}
          </div>
          {#if item.is_anonymous}
            <span class="inline-block mt-1 px-2 py-0.5 rounded-md bg-amber-100 text-amber-900 text-[10px] font-bold">
              Anonim Terlindungi
            </span>
          {/if}
        </td>

        <td class="px-4 py-3.5 align-top max-w-xs">
          <div class="font-bold text-xs text-[#1a120c]">{item.violation_type.replaceAll('_', ' ')}</div>
          <div class="text-[11px] text-[#3d261a]/80 mt-0.5 flex items-center gap-1">
            <MapPin class="w-3 h-3 text-[#F08300] shrink-0" />
            <span class="truncate">{item.regency_name}{item.district_name ? `, ${item.district_name}` : ''}</span>
          </div>
        </td>

        <td class="px-4 py-3.5 align-top whitespace-nowrap">
          <StatusBadge status={item.status} size="sm" />
        </td>

        <td class="px-4 py-3.5 align-top text-xs text-[#3d261a]">
          {item.disposed_to || 'Belum didisposisikan'}
        </td>

        <td class="px-4 py-3.5 align-top text-right whitespace-nowrap">
          <button
            type="button"
            on:click={() => openComplaintDetail(item)}
            class="inline-flex items-center gap-1.5 px-3 py-1.5 rounded-xl bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold transition shadow-xs"
          >
            Tindak Lanjut
          </button>
        </td>
      </tr>
    </DataTable>
  {/if}

  <!-- TAB 3: CMS BERITA -->
  {#if currentTab === 'articles'}
    {#if isWritingArticle}
      <div class="bg-white border border-[#eee9db] rounded-2xl p-6 md:p-8 shadow-xs space-y-6">
        <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 border-b border-[#eee9db] pb-4">
          <div>
            <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Redaksi &amp; Publikasi Warta</div>
            <h2 class="text-lg font-black text-[#1a120c] mt-0.5">
              {editingArticleId ? 'Edit Warta Berita' : 'Editor Penulisan Berita Resmi'}
            </h2>
          </div>
          <div class="flex items-center gap-2">
            <button
              type="button"
              on:click={() => (isWritingArticle = false)}
              class="px-4 py-2 border border-[#eee9db] text-[#3d261a] text-xs font-bold rounded-xl hover:bg-[#f9f4e8] transition"
            >
              Batal
            </button>
            <button
              type="button"
              on:click={handlePublishArticle}
              class="bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold px-5 py-2 rounded-xl transition shadow-sm"
            >
              {editingArticleId ? 'Simpan Perubahan' : 'Terbitkan Berita'}
            </button>
          </div>
        </div>

        {#if articleError}
          <div role="alert" class="p-3.5 rounded-xl bg-red-50 border border-red-200 text-red-800 text-xs font-bold">
            {articleError}
          </div>
        {/if}

        <div class="space-y-5">
          <div>
            <label for="editor-article-title" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">
              Judul Berita (Headline) *
            </label>
            <input
              id="editor-article-title"
              type="text"
              value={articleTitle}
              on:input={handleTitleInput}
              required
              placeholder="Masukkan judul berita yang informatif dan akurat..."
              class="w-full border border-[#eee9db] rounded-xl p-3 text-sm font-bold text-[#1a120c] focus:border-[#F08300] focus:outline-none"
            />
            <div class="flex justify-between items-center text-[11px] text-[#3d261a]/65 mt-1.5">
              <span class="font-mono">Slug: /berita/{articleSlug || 'judul-berita'}</span>
              <span>{articleTitle.length} karakter</span>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label for="editor-article-cat" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">Kategori Warta</label>
              <select id="editor-article-cat" bind:value={articleCategory} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white">
                <option value="Kegiatan">Kegiatan &amp; Patroli</option>
                <option value="Penertiban Aset">Penertiban Aset Daerah</option>
                <option value="Pengumuman">Pengumuman Resmi</option>
                <option value="Opini / Rilis">Rilis Humas Resmi</option>
              </select>
            </div>
            <div>
              <label for="editor-article-author" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">Penulis / Unit Peliput</label>
              <input id="editor-article-author" type="text" bind:value={articleAuthor} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs" />
            </div>
          </div>

          <FileUploader
            label="Foto Utama / Thumbnail Berita"
            description="Unggah gambar beresolusi baik (JPG, PNG, WebP) untuk kartu berita portal."
            accept="image/*"
            previewShape="banner"
            value={articleThumbnail}
            fileName={articleFileName}
            on:change={handleArticleThumbChange}
            on:clear={() => {
              articleThumbnail = '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg';
              articleFileName = '';
            }}
          />

          <div>
            <label for="editor-article-summary" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">
              Ringkasan / Lead Paragraph *
            </label>
            <textarea
              id="editor-article-summary"
              bind:value={articleSummary}
              rows="2"
              placeholder="Tuliskan 1-2 kalimat pengantar utama..."
              class="w-full border border-[#eee9db] rounded-xl p-3 text-xs focus:border-[#F08300] focus:outline-none"
            ></textarea>
          </div>

          <div>
            <label for="editor-article-content" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">
              Isi Berita Lengkap
            </label>
            <RichTextEditor
              bind:content={articleContent}
              placeholder="Mulai menulis naskah berita... Tekan Enter pada daftar nomor (1.) atau poin (-) untuk penomoran otomatis."
            />
          </div>
        </div>
      </div>
    {:else}
      <DataTable
        items={$articlesStore}
        searchFields={['title', 'category', 'author_name', 'summary']}
        searchPlaceholder="Cari judul warta berita, kategori, atau penulis (on-chain)..."
        pageSize={8}
        pageSizeOptions={[5, 8, 15, 30]}
        title="Daftar Warta Berita & Publikasi"
        description="Kelola rilis pers, laporan patroli, dan berita kegiatan resmi Satpol PP Sulsel."
        emptyMessage="Tidak ada warta berita yang sesuai dengan kata kunci pencarian."
      >
        <svelte:fragment slot="header-actions">
          <button
            type="button"
            on:click={openNewArticleEditor}
            class="inline-flex items-center gap-1.5 bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold px-4 py-2 rounded-xl transition shadow-xs"
          >
            <Plus class="w-4 h-4" />
            Tulis Berita Baru
          </button>
        </svelte:fragment>

        <tr slot="head">
          <th class="px-4 py-3.5">Thumbnail</th>
          <th class="px-4 py-3.5">Judul Berita</th>
          <th class="px-4 py-3.5">Kategori</th>
          <th class="px-4 py-3.5">Penulis</th>
          <th class="px-4 py-3.5">Dibaca</th>
          <th class="px-4 py-3.5 text-right">Aksi</th>
        </tr>

        <tr slot="row" let:item class="hover:bg-[#f9f4e8]/55 transition">
          <td class="px-4 py-3.5">
            <img src={item.thumbnail} alt={item.title} class="w-16 h-12 object-cover rounded-xl border border-[#eee9db]" />
          </td>
          <td class="px-4 py-3.5 font-bold text-[#1a120c] max-w-sm">
            <a href={`/berita/${item.slug}`} target="_blank" class="hover:text-[#F08300] transition line-clamp-2">
              {item.title}
            </a>
            <div class="text-[10px] text-[#3d261a]/60 font-normal mt-0.5 font-mono">/berita/{item.slug}</div>
          </td>
          <td class="px-4 py-3.5">
            <span class="bg-[#f9f4e8] text-[#3d261a] text-[10px] font-bold uppercase px-2.5 py-1 rounded-lg border border-[#eee9db]">
              {item.category}
            </span>
          </td>
          <td class="px-4 py-3.5 text-xs text-[#3d261a]/75">{item.author_name}</td>
          <td class="px-4 py-3.5 font-bold text-xs text-[#1a120c]">{item.view_count}x</td>
          <td class="px-4 py-3.5 text-right whitespace-nowrap">
            <div class="inline-flex items-center gap-1.5">
              <button
                type="button"
                on:click={() => openEditArticle(item)}
                class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] text-xs font-bold border border-[#eee9db] transition"
              >
                <Edit3 class="w-3.5 h-3.5" />
                Edit
              </button>
              <button
                type="button"
                on:click={() => requestDelete('article', item.id, item.title)}
                aria-label="Hapus berita"
                class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-red-50 hover:bg-red-100 text-red-700 text-xs font-bold border border-red-200 transition"
              >
                <Trash2 class="w-3.5 h-3.5" />
              </button>
            </div>
          </td>
        </tr>
      </DataTable>
    {/if}
  {/if}

  <!-- TAB 4: CMS DOKUMEN PUBLIK (DATATABLE + MODAL) -->
  {#if currentTab === 'documents'}
    <DataTable
      items={$documentsStore}
      searchFields={['title', 'doc_number', 'category', 'year']}
      searchPlaceholder="Cari judul berkas, nomor registrasi, kategori, atau tahun (on-chain)..."
      pageSize={10}
      pageSizeOptions={[5, 10, 25]}
      title="Katalog Dokumen Transparansi Publik"
      description="Kelola berkas Renja, Renstra, DPA, LAKIP, dan Peraturan Daerah."
      emptyMessage="Tidak ada berkas dokumen yang sesuai dengan kata kunci."
    >
      <svelte:fragment slot="header-actions">
        <button
          type="button"
          on:click={openDocModal}
          class="inline-flex items-center gap-1.5 bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold px-4 py-2 rounded-xl transition shadow-xs"
        >
          <Plus class="w-4 h-4" />
          Unggah Dokumen
        </button>
      </svelte:fragment>

      <tr slot="head">
        <th class="px-4 py-3.5">Nama Dokumen</th>
        <th class="px-4 py-3.5">No. Registrasi</th>
        <th class="px-4 py-3.5">Kategori</th>
        <th class="px-4 py-3.5">Tahun</th>
        <th class="px-4 py-3.5">Ukuran</th>
        <th class="px-4 py-3.5 text-right">Aksi</th>
      </tr>

      <tr slot="row" let:item class="hover:bg-[#f9f4e8]/55 transition">
        <td class="px-4 py-3.5 font-bold text-xs text-[#1a120c]">{item.title}</td>
        <td class="px-4 py-3.5 font-mono text-xs text-[#3d261a]/75">{item.doc_number || '-'}</td>
        <td class="px-4 py-3.5">
          <span class="bg-[#f9f4e8] text-[#3d261a] text-[10px] font-bold uppercase px-2.5 py-1 rounded-lg border border-[#eee9db]">
            {item.category}
          </span>
        </td>
        <td class="px-4 py-3.5 font-bold text-xs text-[#1a120c]">{item.year}</td>
        <td class="px-4 py-3.5 text-xs text-[#3d261a]/75">{item.file_size}</td>
        <td class="px-4 py-3.5 text-right whitespace-nowrap">
          <div class="inline-flex items-center gap-1.5">
            <a
              href={item.file_url}
              target="_blank"
              class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-[#f9f4e8] hover:bg-[#eee9db] text-[#1a120c] text-xs font-bold border border-[#eee9db] transition"
            >
              <Download class="w-3.5 h-3.5" />
              Unduh
            </a>
            <button
              type="button"
              on:click={() => requestDelete('document', item.id, item.title)}
              aria-label="Hapus dokumen"
              class="inline-flex items-center gap-1 px-2.5 py-1.5 rounded-xl bg-red-50 hover:bg-red-100 text-red-700 text-xs font-bold border border-red-200 transition"
            >
              <Trash2 class="w-3.5 h-3.5" />
            </button>
          </div>
        </td>
      </tr>
    </DataTable>
  {/if}

  <!-- TAB 5: CMS GALERI MULTIMEDIA -->
  {#if currentTab === 'gallery'}
    <div class="space-y-5">
      <div class="bg-white border border-[#eee9db] rounded-2xl p-5 flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 shadow-xs">
        <div>
          <h2 class="text-sm md:text-base font-black text-[#1a120c]">Galeri Multimedia Dokumentasi</h2>
          <p class="text-xs text-[#3d261a]/75 mt-0.5">Kelola foto kegiatan apel, operasi penertiban, dan koordinasi lapangan.</p>
        </div>
        <button
          type="button"
          on:click={openGalleryModal}
          class="inline-flex items-center gap-2 bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold px-4 py-2.5 rounded-xl transition shadow-xs shrink-0"
        >
          <Plus class="w-4 h-4" />
          Tambah Foto Galeri
        </button>
      </div>

      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
        {#each $galleryStore as item}
          <article class="bg-white border border-[#eee9db] rounded-2xl overflow-hidden shadow-xs group flex flex-col justify-between">
            <div>
              <div class="relative h-44 bg-[#f9f4e8] overflow-hidden">
                <img
                  src={item.image_url}
                  alt={item.title}
                  class="w-full h-full object-cover group-hover:scale-105 transition duration-300"
                />
                <span class="absolute top-3 left-3 bg-[#1a120c]/85 text-[#F08300] font-bold text-[10px] uppercase px-2.5 py-1 rounded-lg">
                  {item.category}
                </span>
              </div>
              <div class="p-4">
                <div class="text-[11px] text-[#3d261a]/65">{item.date}</div>
                <h3 class="text-xs font-black text-[#1a120c] line-clamp-2 min-h-[2.25rem] mt-1">{item.title}</h3>
              </div>
            </div>
            <div class="px-4 py-3 border-t border-[#eee9db] bg-[#f9f4e8]/40 flex justify-end">
              <button
                type="button"
                on:click={() => requestDelete('gallery', item.id, item.title)}
                class="inline-flex items-center gap-1 text-xs font-bold text-red-700 hover:underline"
              >
                <Trash2 class="w-3.5 h-3.5" />
                Hapus Foto
              </button>
            </div>
          </article>
        {/each}
      </div>
    </div>
  {/if}

  <!-- TAB 6: STRUKTUR ORGANISASI DINAMIS -->
  {#if currentTab === 'organization'}
    <OrganizationManager actorName={currentUser?.username || 'admin'} />
  {/if}
</DashboardShell>

<!-- MODAL 1: TINDAK LANJUT SOP PENGADUAN SIAP TERTIB -->
<Modal
  bind:open={isEditingComplaint}
  title={selectedComplaint ? `Tindak Lanjut Tiket #${selectedComplaint.ticket_code}` : 'Detail Tiket'}
  subtitle="Verifikasi bukti, disposisi bidang, dan kirim pembaruan email ke pelapor."
  size="xl"
>
  {#if selectedComplaint}
    {@const masked = maskComplaintIdentity(selectedComplaint)}
    {@const evidenceUrl = selectedComplaint.evidence_image_url || selectedComplaint.evidence_photos || ''}
    <div class="space-y-5 text-xs">
      <!-- Identitas & Lokasi -->
      <div class="bg-[#f9f4e8]/75 border border-[#eee9db] rounded-2xl p-4 space-y-4">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <span class="text-[#3d261a]/70 uppercase text-[10px] font-bold block mb-1.5">Identitas Pelapor</span>
            {#if selectedComplaint.is_anonymous}
              {#if unmaskRequested}
                <div class="bg-amber-50 border border-amber-200 rounded-xl p-3 space-y-1">
                  <div class="flex items-center justify-between gap-2">
                    <span class="text-[10px] font-black text-amber-900 uppercase">Data Terbuka (Audit Aktif)</span>
                    <span class="text-[10px] bg-amber-200 text-amber-900 px-2 py-0.5 rounded-lg font-bold flex items-center gap-1">
                      <EyeOff class="w-3 h-3" /> Tercatat
                    </span>
                  </div>
                  <div class="font-black text-[#1a120c]">{selectedComplaint.reporter_name}</div>
                  <div class="font-mono text-[11px] text-[#3d261a]">NIK: {selectedComplaint.reporter_nik || '-'}</div>
                  <div class="text-[11px] text-[#3d261a]">Telp: {selectedComplaint.reporter_phone || '-'}</div>
                  <div class="text-[11px] text-[#3d261a] flex items-center gap-1">
                    <Mail class="w-3 h-3 text-[#F08300]" /> {selectedComplaint.reporter_email || 'Tidak dilampirkan'}
                  </div>
                </div>
              {:else}
                <div class="bg-white border border-[#eee9db] rounded-xl p-3 space-y-1.5">
                  <div class="font-black text-[#1a120c]">{masked.name}</div>
                  <div class="text-[11px] font-mono text-[#3d261a]/75">NIK: {masked.nik}</div>
                  <div class="text-[11px] text-[#3d261a]/75">Telp: {masked.phone}</div>
                  <div class="text-[11px] text-[#3d261a]/75">Email: {masked.email || '-'}</div>
                  <button
                    type="button"
                    on:click={handleUnmaskIdentity}
                    class="mt-1 inline-flex items-center gap-1.5 px-3 py-1.5 bg-[#F08300] hover:bg-[#d97300] text-white text-[11px] font-bold rounded-xl transition"
                  >
                    <Eye class="w-3.5 h-3.5" />
                    Buka Masking (Audit Petugas)
                  </button>
                </div>
              {/if}
            {:else}
              <div class="bg-white border border-[#eee9db] rounded-xl p-3 space-y-1">
                <div class="font-black text-[#1a120c]">{selectedComplaint.reporter_name}</div>
                <div class="font-mono text-[11px] text-[#3d261a]/80">NIK: {selectedComplaint.reporter_nik || '-'}</div>
                <div class="text-[11px] text-[#3d261a]/80">Telp: {selectedComplaint.reporter_phone || '-'}</div>
                <div class="text-[11px] text-[#3d261a]/80 flex items-center gap-1">
                  <Mail class="w-3 h-3 text-[#F08300]" /> {selectedComplaint.reporter_email || 'Tidak dilampirkan'}
                </div>
              </div>
            {/if}
          </div>

          <div class="bg-white border border-[#eee9db] rounded-xl p-3 space-y-1">
            <span class="text-[#3d261a]/70 uppercase text-[10px] font-bold block mb-1">Wilayah &amp; Waktu Laporan</span>
            <div class="font-black text-[#1a120c]">{selectedComplaint.regency_name}</div>
            <div class="text-[11px] text-[#3d261a]/80">{selectedComplaint.district_name}, {selectedComplaint.village_name}</div>
            <div class="text-[11px] text-[#3d261a]/70">
              Dikirim: {new Date(selectedComplaint.created_at).toLocaleString('id-ID')}
            </div>
            <div class="pt-2">
              <StatusBadge status={selectedComplaint.status} size="sm" />
            </div>
          </div>
        </div>

        <div class="pt-2 border-t border-[#eee9db]">
          <span class="text-[#3d261a]/70 uppercase text-[10px] font-bold">Alamat Kejadian</span>
          <div class="text-[#1a120c] font-semibold mt-0.5">{selectedComplaint.address_detail}</div>
        </div>

        {#if evidenceUrl}
          <div class="pt-3 border-t border-[#eee9db]">
            <span class="text-[#3d261a]/70 uppercase text-[10px] font-bold block mb-1.5">Bukti Foto Lapangan</span>
            <a href={evidenceUrl} target="_blank" class="inline-block group">
              <img
                src={evidenceUrl}
                alt="Bukti Foto Aduan"
                class="w-60 h-36 object-cover rounded-xl border border-[#eee9db] group-hover:opacity-90 transition shadow-xs"
              />
              <span class="text-[11px] text-[#F08300] font-bold block mt-1 group-hover:underline">Buka Foto Ukuran Penuh &rarr;</span>
            </a>
          </div>
        {/if}
      </div>

      <div>
        <span class="text-[#3d261a]/70 uppercase text-[10px] font-bold">Kategori: {selectedComplaint.violation_type}</span>
        <div class="p-3.5 bg-white border border-[#eee9db] rounded-xl mt-1 text-[#3d261a] leading-relaxed">
          {selectedComplaint.description}
        </div>
      </div>

      <fieldset class="space-y-4 pt-3 border-t border-[#eee9db]">
        <legend class="text-xs font-black uppercase tracking-wider text-[#1a120c]">Pembaruan Status SOP &amp; Notifikasi Email</legend>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div>
            <label for="complaint-status-select" class="block text-[11px] font-bold text-[#3d261a] mb-1.5">Status Prosedur</label>
            <select id="complaint-status-select" bind:value={editStatus} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white font-bold">
              <option value="MENUNGGU_VERIFIKASI">1. Menunggu Verifikasi</option>
              <option value="DISPOSISI_BIDANG">2. Disposisi Bidang</option>
              <option value="PENINDAKAN_LAPANGAN">3. Penindakan Lapangan</option>
              <option value="SELESAI">4. Selesai &amp; BAP Terbit</option>
              <option value="DITOLAK">Ditolak (Tidak Sesuai Kewenangan)</option>
            </select>
          </div>

          <div>
            <label for="complaint-disposition-input" class="block text-[11px] font-bold text-[#3d261a] mb-1.5">Disposisi Unit Kerja</label>
            <input
              id="complaint-disposition-input"
              type="text"
              list="sop-org-units"
              bind:value={editDisposedTo}
              placeholder="Pilih atau ketik unit/pleton..."
              class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white"
            />
            <datalist id="sop-org-units">
              {#each dispositionOptions as opt}
                <option value={opt}></option>
              {/each}
            </datalist>
          </div>
        </div>

        <div>
          <label for="complaint-bap-notes" class="block text-[11px] font-bold text-[#3d261a] mb-1.5">Catatan Tindakan / Berita Acara (BAP)</label>
          <textarea
            id="complaint-bap-notes"
            bind:value={editActionNotes}
            rows="3"
            placeholder="Catat hasil verifikasi, penertiban lapangan, atau alasan penolakan..."
            class="w-full border border-[#eee9db] rounded-xl p-3 text-xs"
          ></textarea>
        </div>
      </fieldset>
    </div>
  {/if}

  <svelte:fragment slot="footer" let:handleClose>
    <button
      type="button"
      on:click={handleClose}
      class="px-4 py-2 rounded-xl border border-[#eee9db] bg-white hover:bg-[#f9f4e8] text-xs font-bold text-[#3d261a] transition"
    >
      Tutup
    </button>
    <button
      type="button"
      on:click={saveComplaintUpdate}
      disabled={isSavingComplaint}
      class="px-5 py-2 rounded-xl bg-[#F08300] hover:bg-[#d97300] disabled:opacity-60 text-white text-xs font-bold transition shadow-sm"
    >
      {isSavingComplaint ? 'Menyimpan...' : 'Simpan & Kirim Notifikasi'}
    </button>
  </svelte:fragment>
</Modal>

<!-- MODAL 2: UNGGAH DOKUMEN PUBLIK -->
<Modal
  bind:open={showDocModal}
  title="Unggah Dokumen Transparansi Publik"
  subtitle="Berkas yang ditambahkan langsung tersedia pada katalog Dokumen Publik."
  size="md"
>
  <form id="doc-upload-form" on:submit|preventDefault={handleSaveDocument} class="space-y-4">
    {#if docError}
      <div role="alert" class="p-3 rounded-xl bg-red-50 border border-red-200 text-red-800 text-xs font-bold">
        {docError}
      </div>
    {/if}

    <FileUploader
      label="Berkas Dokumen Resmi"
      description="Pilih file PDF, Word, atau Excel dari perangkat Anda."
      accept=".pdf,.doc,.docx,.xls,.xlsx"
      previewShape="banner"
      fileName={uploadedFileName}
      on:change={handleDocFileChange}
      on:clear={() => {
        uploadedFileName = '';
        docFileUrl = '/uploads/dokumen/dokumen_resmi.pdf';
      }}
    />

    <div>
      <label for="doc-title-input" class="block text-xs font-bold text-[#1a120c] mb-1">Nama / Judul Dokumen *</label>
      <input id="doc-title-input" type="text" required bind:value={docTitle} placeholder="Contoh: Rencana Kerja Satpol PP TA 2026" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
      <p class="field-error">Judul dokumen wajib diisi.</p>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
      <div>
        <label for="doc-number-input" class="block text-xs font-bold text-[#1a120c] mb-1">Nomor SK / Registrasi</label>
        <input id="doc-number-input" type="text" bind:value={docNumber} placeholder="050/112/SATPOL/2026" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
      </div>
      <div>
        <label for="doc-year-select" class="block text-xs font-bold text-[#1a120c] mb-1">Tahun Anggaran</label>
        <select id="doc-year-select" bind:value={docYear} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white">
          <option value={2026}>2026</option>
          <option value={2025}>2025</option>
          <option value={2024}>2024</option>
          <option value={2023}>2023</option>
          <option value={2022}>2022</option>
        </select>
      </div>
    </div>

    <div>
      <label for="doc-category-select" class="block text-xs font-bold text-[#1a120c] mb-1">Kategori Dokumen</label>
      <select id="doc-category-select" bind:value={docCategory} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white">
        <option value="renja">Rencana Kerja (Renja)</option>
        <option value="renstra">Rencana Strategis (Renstra)</option>
        <option value="dpa">Dokumen Pelaksanaan Anggaran (DPA)</option>
        <option value="lakip">LAKIP / LKjIP</option>
        <option value="keuangan">Laporan Keuangan</option>
        <option value="rpjmd">RPJMD</option>
        <option value="perda">Peraturan Daerah (Perda)</option>
        <option value="sop">Standar Operasional (SOP)</option>
      </select>
    </div>
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
      form="doc-upload-form"
      class="px-5 py-2 rounded-xl bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold transition shadow-sm"
    >
      Simpan Dokumen
    </button>
  </svelte:fragment>
</Modal>

<!-- MODAL 3: TAMBAH FOTO GALERI -->
<Modal
  bind:open={showGalleryModal}
  title="Tambah Foto Dokumentasi Galeri"
  subtitle="Foto baru akan langsung tampil pada halaman Galeri dan preview Beranda."
  size="md"
>
  <form id="gallery-upload-form" on:submit|preventDefault={handleSaveGallery} class="space-y-4">
    {#if galleryError}
      <div role="alert" class="p-3 rounded-xl bg-red-50 border border-red-200 text-red-800 text-xs font-bold">
        {galleryError}
      </div>
    {/if}

    <FileUploader
      label="Pilih Foto Dokumentasi"
      description="Format JPG, PNG, atau WebP."
      accept="image/*"
      previewShape="banner"
      value={galleryImageUrl}
      fileName={galleryFileName}
      on:change={handleGalleryFileChange}
      on:clear={() => {
        galleryFileName = '';
        galleryImageUrl = '/uploads/foto_galeri/1687082076352158666_3553275628273619_3927930681691724075_n.jpg';
      }}
    />

    <div>
      <label for="gallery-title-input" class="block text-xs font-bold text-[#1a120c] mb-1">Judul / Keterangan Kegiatan *</label>
      <input id="gallery-title-input" type="text" required bind:value={galleryTitle} placeholder="Contoh: Apel Kesiapsiagaan Anggota Satpol PP" class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white" />
      <p class="field-error">Judul kegiatan wajib diisi.</p>
    </div>

    <div>
      <label for="gallery-category-select" class="block text-xs font-bold text-[#1a120c] mb-1">Kategori Dokumentasi</label>
      <select id="gallery-category-select" bind:value={galleryCategory} class="w-full border border-[#eee9db] rounded-xl p-2.5 text-xs bg-white">
        <option value="Penertiban">Operasi &amp; Penertiban</option>
        <option value="Apel & Upacara">Apel &amp; Upacara Resmi</option>
        <option value="Koordinasi">Koordinasi &amp; Sosialisasi</option>
        <option value="Kegiatan">Giat Harian</option>
      </select>
    </div>
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
      form="gallery-upload-form"
      class="px-5 py-2 rounded-xl bg-[#F08300] hover:bg-[#d97300] text-white text-xs font-bold transition shadow-sm"
    >
      Simpan ke Galeri
    </button>
  </svelte:fragment>
</Modal>

<!-- MODAL 4: KONFIRMASI HAPUS ITEM CMS -->
<Modal
  bind:open={confirmDeleteOpen}
  title="Konfirmasi Hapus Data"
  subtitle="Tindakan ini akan menghapus item dari portal publik."
  size="sm"
>
  <p class="text-xs text-[#3d261a] leading-relaxed">
    Yakin ingin menghapus <strong>{confirmDeleteLabel}</strong>? Penghapusan ini akan dicatat di log audit sistem.
  </p>

  <svelte:fragment slot="footer" let:handleClose>
    <button
      type="button"
      on:click={handleClose}
      class="px-4 py-2 rounded-xl border border-[#eee9db] bg-white hover:bg-[#f9f4e8] text-xs font-bold text-[#3d261a] transition"
    >
      Batal
    </button>
    <button
      type="button"
      on:click={executeDelete}
      class="px-4 py-2 rounded-xl bg-red-600 hover:bg-red-700 text-white text-xs font-bold transition"
    >
      Ya, Hapus
    </button>
  </svelte:fragment>
</Modal>
