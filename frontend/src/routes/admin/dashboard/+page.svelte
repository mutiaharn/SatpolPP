<script lang="ts">
  import { onMount } from 'svelte';
  import { 
    complaintsStore, 
    articlesStore, 
    documentsStore, 
    galleryStore,
    auditLogsStore,
    defaultRegencies 
  } from '$lib/stores/appState';
  import type { Complaint, Article, PublicDocument, GalleryItem } from '$lib/types';
  import RichTextEditor from '$lib/components/RichTextEditor.svelte';

  // Navigation tab state
  type AdminTab = 'overview' | 'complaints' | 'articles' | 'documents' | 'gallery';
  let currentTab: AdminTab = 'complaints';
  let sidebarOpen = false; // Mobile sidebar drawer state

  let currentUser: any = null;

  // ==========================================
  // 1. COMPLAINTS (SIAP TERTIB) MANAGEMENT
  // ==========================================
  let statusFilter = 'ALL';
  let searchQuery = '';
  let selectedComplaint: Complaint | null = null;
  let isEditingComplaint = false;

  let editStatus: Complaint['status'] = 'MENUNGGU_VERIFIKASI';
  let editDisposedTo = '';
  let editActionNotes = '';
  let unmaskRequested = false;

  $: filteredComplaints = $complaintsStore.filter(c => {
    const matchStatus = statusFilter === 'ALL' || c.status === statusFilter;
    const matchSearch = searchQuery === '' || 
      c.ticket_code.toLowerCase().includes(searchQuery.toLowerCase()) ||
      c.reporter_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      c.regency_name.toLowerCase().includes(searchQuery.toLowerCase()) ||
      c.violation_type.toLowerCase().includes(searchQuery.toLowerCase());
    return matchStatus && matchSearch;
  });

  $: totalCount = $complaintsStore.length;
  $: pendingCount = $complaintsStore.filter(c => c.status === 'MENUNGGU_VERIFIKASI').length;
  $: disposedCount = $complaintsStore.filter(c => c.status === 'DISPOSISI_BIDANG').length;
  $: inProgressCount = $complaintsStore.filter(c => c.status === 'PENINDAKAN_LAPANGAN').length;
  $: completedCount = $complaintsStore.filter(c => c.status === 'SELESAI').length;

  function openComplaintDetail(comp: Complaint) {
    selectedComplaint = comp;
    editStatus = comp.status;
    editDisposedTo = comp.disposed_to || 'Bidang Ketenteraman dan Ketertiban Umum';
    editActionNotes = comp.action_notes || '';
    unmaskRequested = false;
    isEditingComplaint = true;
  }

  function handleUnmaskIdentity() {
    if (!selectedComplaint) return;
    unmaskRequested = true;
    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: currentUser?.username || 'petugas',
        action: 'BUKA_MASKING_ANONIM',
        resource: `Tiket ${selectedComplaint.ticket_code}`,
        details: 'Membuka data identitas terlindung pelapor atas wewenang petugas verifikator',
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);
  }

  function saveComplaintUpdate() {
    if (!selectedComplaint) return;
    
    complaintsStore.update(list => {
      return list.map(item => {
        if (item.ticket_code === selectedComplaint?.ticket_code) {
          return {
            ...item,
            status: editStatus,
            disposed_to: editDisposedTo,
            action_notes: editActionNotes
          };
        }
        return item;
      });
    });

    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: currentUser?.username || 'petugas',
        action: 'UPDATE_STATUS_ADUAN',
        resource: `Tiket ${selectedComplaint.ticket_code}`,
        details: `Status SOP: ${editStatus}, Disposisi: ${editDisposedTo}`,
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);

    alert(`Status tiket ${selectedComplaint.ticket_code} berhasil diperbarui.`);
    isEditingComplaint = false;
    selectedComplaint = null;
  }

  // ==========================================
  // 2. CMS BERITA (KOMPASIANA-GRADE EDITOR)
  // ==========================================
  let isWritingArticle = false;
  let articleEditorTab: 'write' | 'preview' = 'write';

  let articleTitle = '';
  let articleSlug = '';
  let articleCategory = 'Kegiatan';
  let articleSummary = '';
  let articleContent = '';
  let articleAuthor = 'Humas Satpol PP';
  let articleThumbnail = '/uploads/foto_berita/1768182343WhatsApp Image 2026-01-12 at 08.08.43.jpeg';
  let articleThumbnailPreview = '';

  // Auto-generate slug when title changes
  function handleTitleInput(e: Event) {
    const target = e.target as HTMLInputElement;
    articleTitle = target.value;
    articleSlug = articleTitle
      .toLowerCase()
      .replace(/[^a-z0-9]+/g, '-')
      .replace(/(^-|-$)+/g, '');
  }

  // Handle article image upload
  function handleArticleImageUpload(e: Event) {
    const input = e.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      const reader = new FileReader();
      reader.onload = (event) => {
        if (event.target?.result) {
          articleThumbnailPreview = event.target.result as string;
          articleThumbnail = event.target.result as string;
        }
      };
      reader.readAsDataURL(file);
    }
  }

  // Helper format markdown/rich text ke clean HTML
  function renderFormattedHtml(text: string): string {
    if (!text) return '';
    const lines = text.split('\n');
    let html = '';
    let inOl = false;
    let inUl = false;

    for (let i = 0; i < lines.length; i++) {
      let line = lines[i];

      if (line.startsWith('## ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<h2 class="text-xl font-extrabold text-[#101116] mt-5 mb-2">${formatInline(line.slice(3))}</h2>`;
        continue;
      }
      if (line.startsWith('### ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<h3 class="text-base font-bold text-[#101116] mt-4 mb-1.5">${formatInline(line.slice(4))}</h3>`;
        continue;
      }
      if (line.startsWith('> ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<blockquote class="border-l-4 border-[#F08300] pl-4 italic text-[#2b2f3f] my-3 bg-[#fcfbf7] py-2 rounded-r-lg">${formatInline(line.slice(2))}</blockquote>`;
        continue;
      }
      if (line.startsWith('[Catatan Redaksi]') || line.startsWith(':::callout')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<div class="bg-[#f9f4e8] border border-[#eee9db] p-4 rounded-xl my-3 text-xs font-semibold text-[#101116]">${formatInline(line.replace(/:::callout|\[Catatan Redaksi\]/, ''))}</div>`;
        continue;
      }

      const numMatch = line.match(/^\s*(\d+)\.\s+(.*)$/);
      if (numMatch) {
        if (inUl) { html += '</ul>'; inUl = false; }
        if (!inOl) { html += '<ol class="list-decimal pl-5 space-y-1 my-2 text-xs text-[#2b2f3f]">'; inOl = true; }
        html += `<li>${formatInline(numMatch[2])}</li>`;
        continue;
      }

      const bulletMatch = line.match(/^\s*[-*]\s+(.*)$/);
      if (bulletMatch) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (!inUl) { html += '<ul class="list-disc pl-5 space-y-1 my-2 text-xs text-[#2b2f3f]">'; inUl = true; }
        html += `<li>${formatInline(bulletMatch[1])}</li>`;
        continue;
      }

      if (inOl) { html += '</ol>'; inOl = false; }
      if (inUl) { html += '</ul>'; inUl = false; }

      if (line.trim() !== '') {
        html += `<p class="text-xs text-[#2b2f3f] leading-relaxed my-2.5">${formatInline(line)}</p>`;
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
    if (!articleTitle || !articleSummary) {
      alert('Judul dan ringkasan warta berita wajib diisi.');
      return;
    }

    const finalSlug = articleSlug || articleTitle.toLowerCase().replace(/[^a-z0-9]+/g, '-');
    const parsedHtml = renderFormattedHtml(articleContent);

    const newArticle: Article = {
      id: Date.now(),
      title: articleTitle,
      slug: finalSlug,
      summary: articleSummary,
      content: parsedHtml ? parsedHtml : `<p>${articleSummary}</p>`,
      thumbnail: articleThumbnail,
      category: articleCategory,
      author_name: articleAuthor,
      view_count: 1,
      published_at: new Date().toISOString()
    };

    articlesStore.update(items => [newArticle, ...items]);

    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: currentUser?.username || 'humas',
        action: 'PUBLISH_ARTICLE',
        resource: `Berita: ${newArticle.title}`,
        details: `Menerbitkan warta baru kategori ${newArticle.category}`,
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);

    alert('Warta Berita berhasil dipublikasikan secara instan ke portal publik!');
    isWritingArticle = false;
    articleTitle = '';
    articleSlug = '';
    articleSummary = '';
    articleContent = '';
    articleThumbnailPreview = '';
  }

  function handleDeleteArticle(id: number) {
    if (confirm('Apakah Anda yakin ingin menghapus warta berita ini dari portal?')) {
      articlesStore.update(items => items.filter(a => a.id !== id));
    }
  }

  // ==========================================
  // 3. CMS DOKUMEN PUBLIK (REAL FILE UPLOADS)
  // ==========================================
  let showDocModal = false;
  let docTitle = '';
  let docNumber = '';
  let docCategory: PublicDocument['category'] = 'renja';
  let docYear = 2026;
  let docSize = '1.8 MB';
  let docFileUrl = '/uploads/dokumen/dokumen_resmi.pdf';
  let uploadedFileName = '';

  function handleDocFileUpload(e: Event) {
    const input = e.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      uploadedFileName = file.name;
      
      // Auto-populate Title if empty
      if (!docTitle) {
        docTitle = file.name.replace(/\.[^/.]+$/, '').replace(/[-_]/g, ' ');
      }

      // Compute actual file size
      const sizeMb = (file.size / (1024 * 1024)).toFixed(1);
      docSize = sizeMb === '0.0' ? `${Math.round(file.size / 1024)} KB` : `${sizeMb} MB`;

      // Create live browser blob URL
      docFileUrl = URL.createObjectURL(file);
    }
  }

  function handleSaveDocument() {
    if (!docTitle) {
      alert('Nama dokumen publik resmi wajib diisi.');
      return;
    }

    const newDoc: PublicDocument = {
      id: Date.now(),
      title: docTitle,
      doc_number: docNumber || `050/SATPOL/${docYear}`,
      category: docCategory,
      year: docYear,
      file_url: docFileUrl,
      file_size: docSize
    };

    documentsStore.update(items => [newDoc, ...items]);

    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: currentUser?.username || 'sekretariat',
        action: 'UPLOAD_DOKUMEN_PUBLIK',
        resource: `Dokumen: ${newDoc.title}`,
        details: `Mengunggah berkas transparansi publik kategori ${newDoc.category} TA ${newDoc.year}`,
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);

    alert('Dokumen publik berhasil diunggah dan langsung tampil di halaman Dokumen Publik!');
    showDocModal = false;
    docTitle = '';
    docNumber = '';
    uploadedFileName = '';
  }

  function handleDeleteDocument(id: number) {
    if (confirm('Apakah Anda yakin ingin menghapus dokumen publik ini?')) {
      documentsStore.update(items => items.filter(d => d.id !== id));
    }
  }

  // ==========================================
  // 4. CMS GALERI MULTIMEDIA
  // ==========================================
  let showGalleryModal = false;
  let galleryTitle = '';
  let galleryCategory = 'Penertiban';
  let galleryDate = new Date().toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' });
  let galleryImageUrl = '/uploads/foto_galeri/1687082076352158666_3553275628273619_3927930681691724075_n.jpg';
  let galleryPreview = '';

  function handleGalleryImageUpload(e: Event) {
    const input = e.target as HTMLInputElement;
    if (input.files && input.files[0]) {
      const file = input.files[0];
      const reader = new FileReader();
      reader.onload = (event) => {
        if (event.target?.result) {
          galleryPreview = event.target.result as string;
          galleryImageUrl = event.target.result as string;
        }
      };
      reader.readAsDataURL(file);
    }
  }

  function handleSaveGallery() {
    if (!galleryTitle) {
      alert('Judul dokumentasi kegiatan wajib diisi.');
      return;
    }

    const newItem: GalleryItem = {
      id: Date.now(),
      title: galleryTitle,
      category: galleryCategory,
      image_url: galleryImageUrl,
      date: galleryDate
    };

    galleryStore.update(items => [newItem, ...items]);

    auditLogsStore.update(logs => [
      {
        id: Date.now(),
        user_name: currentUser?.username || 'humas',
        action: 'UPLOAD_GALERI_FOTO',
        resource: `Foto: ${newItem.title}`,
        details: `Menambahkan foto dokumentasi kategori ${newItem.category}`,
        timestamp: new Date().toISOString().replace('T', ' ').substring(0, 19),
        ip_address: '103.151.191.73'
      },
      ...logs
    ]);

    alert('Foto dokumentasi kegiatan berhasil ditambahkan ke galeri portal!');
    showGalleryModal = false;
    galleryTitle = '';
    galleryPreview = '';
  }

  function handleDeleteGallery(id: number) {
    if (confirm('Apakah Anda yakin ingin menghapus foto kegiatan ini dari galeri?')) {
      galleryStore.update(items => items.filter(g => g.id !== id));
    }
  }

  onMount(() => {
    const authRaw = sessionStorage.getItem('satpol_auth');
    if (authRaw) {
      try {
        currentUser = JSON.parse(authRaw).user;
      } catch (e) {
        console.error(e);
      }
    }
  });

  function getStatusLabel(s: string) {
    switch (s) {
      case 'MENUNGGU_VERIFIKASI': return { text: 'Menunggu Verifikasi', bg: 'bg-amber-50 text-amber-900 border-amber-300' };
      case 'DISPOSISI_BIDANG': return { text: 'Disposisi Bidang', bg: 'bg-blue-50 text-blue-900 border-blue-300' };
      case 'PENINDAKAN_LAPANGAN': return { text: 'Penindakan Lapangan', bg: 'bg-purple-50 text-purple-900 border-purple-300' };
      case 'SELESAI': return { text: 'Selesai & BAP', bg: 'bg-emerald-50 text-emerald-900 border-emerald-300' };
      case 'DITOLAK': return { text: 'Ditolak', bg: 'bg-rose-50 text-rose-900 border-rose-300' };
      default: return { text: s, bg: 'bg-gray-100 text-gray-800 border-gray-300' };
    }
  }
</script>

<svelte:head>
  <title>Panel Operasional Satpol PP &amp; CMS Terpadu</title>
</svelte:head>

<div class="min-h-screen bg-[#f4f5f8] flex flex-col lg:flex-row text-[#101116] font-sans antialiased">

  <!-- ============================================================== -->
  <!-- RESPONSIVE SIDEBAR (Desktop: Fixed / Mobile: Off-canvas drawer) -->
  <!-- ============================================================== -->

  <!-- Mobile Backdrop Overlay -->
  {#if sidebarOpen}
    <div 
      class="fixed inset-0 z-40 bg-black/60 lg:hidden backdrop-blur-sm transition-opacity"
      on:click={() => sidebarOpen = false}
      role="button"
      tabindex="0"
      on:keydown={(e) => e.key === 'Escape' && (sidebarOpen = false)}
    ></div>
  {/if}

  <aside 
    class={`fixed lg:sticky top-0 left-0 z-50 h-screen w-72 bg-[#12151e] text-white flex flex-col justify-between transition-transform duration-300 ease-in-out border-r border-[#262a38] shadow-2xl lg:shadow-none lg:translate-x-0 ${sidebarOpen ? 'translate-x-0' : '-translate-x-full'}`}
  >
    <!-- Sidebar Header -->
    <div class="p-6 border-b border-[#262a38]">
      <div class="flex items-center gap-3">
        <img 
          src="/uploads/foto_website/1686928245logo_white.png" 
          alt="Emblem Satpol PP" 
          class="h-11 w-auto object-contain" 
        />
        <div>
          <div class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">Satpol PP Prov. Sulsel</div>
          <div class="text-sm font-extrabold text-white tracking-tight">Panel Pengendali</div>
        </div>
      </div>
      <div class="mt-4 pt-3 border-t border-[#262a38] text-[10px] text-[#9ca3af] uppercase tracking-widest font-semibold flex items-center justify-between">
        <span>SOP Terpadu &amp; CMS</span>
        <span class="inline-block w-2 h-2 rounded-full bg-emerald-500"></span>
      </div>
    </div>

    <!-- Navigation Menu Items -->
    <nav class="p-4 space-y-1.5 flex-1 overflow-y-auto">
      <button 
        class={`w-full flex items-center justify-between px-4 py-3 rounded-xl text-xs font-bold uppercase tracking-wider transition ${currentTab === 'overview' ? 'bg-[#F08300] text-white shadow-md' : 'text-gray-300 hover:bg-[#1d2232] hover:text-white'}`}
        on:click={() => { currentTab = 'overview'; sidebarOpen = false; }}
      >
        <span>Ikhtisar Operasional</span>
        <span class="text-[10px] bg-black/20 px-2 py-0.5 rounded-md font-mono">{totalCount}</span>
      </button>

      <button 
        class={`w-full flex items-center justify-between px-4 py-3 rounded-xl text-xs font-bold uppercase tracking-wider transition ${currentTab === 'complaints' ? 'bg-[#F08300] text-white shadow-md' : 'text-gray-300 hover:bg-[#1d2232] hover:text-white'}`}
        on:click={() => { currentTab = 'complaints'; sidebarOpen = false; }}
      >
        <span>Layanan SIAP TERTIB</span>
        <span class="text-[10px] bg-amber-500/30 text-amber-300 px-2 py-0.5 rounded-md font-mono">{pendingCount} Baru</span>
      </button>

      <button 
        class={`w-full flex items-center justify-between px-4 py-3 rounded-xl text-xs font-bold uppercase tracking-wider transition ${currentTab === 'articles' ? 'bg-[#F08300] text-white shadow-md' : 'text-gray-300 hover:bg-[#1d2232] hover:text-white'}`}
        on:click={() => { currentTab = 'articles'; sidebarOpen = false; }}
      >
        <span>CMS Berita &amp; Warta</span>
        <span class="text-[10px] bg-black/20 px-2 py-0.5 rounded-md font-mono">{$articlesStore.length}</span>
      </button>

      <button 
        class={`w-full flex items-center justify-between px-4 py-3 rounded-xl text-xs font-bold uppercase tracking-wider transition ${currentTab === 'documents' ? 'bg-[#F08300] text-white shadow-md' : 'text-gray-300 hover:bg-[#1d2232] hover:text-white'}`}
        on:click={() => { currentTab = 'documents'; sidebarOpen = false; }}
      >
        <span>CMS Dokumen Publik</span>
        <span class="text-[10px] bg-black/20 px-2 py-0.5 rounded-md font-mono">{$documentsStore.length}</span>
      </button>

      <button 
        class={`w-full flex items-center justify-between px-4 py-3 rounded-xl text-xs font-bold uppercase tracking-wider transition ${currentTab === 'gallery' ? 'bg-[#F08300] text-white shadow-md' : 'text-gray-300 hover:bg-[#1d2232] hover:text-white'}`}
        on:click={() => { currentTab = 'gallery'; sidebarOpen = false; }}
      >
        <span>CMS Galeri Kegiatan</span>
        <span class="text-[10px] bg-black/20 px-2 py-0.5 rounded-md font-mono">{$galleryStore.length}</span>
      </button>

      <!-- Divider -->
      <div class="pt-4 pb-2">
        <div class="text-[10px] font-bold text-[#6f7775] uppercase tracking-widest px-4">Tautan Eksternal</div>
      </div>

      <a 
        href="/superadmin/dashboard" 
        class="block px-4 py-2.5 rounded-xl text-xs font-bold text-gray-300 hover:bg-[#1d2232] hover:text-[#F08300] transition"
      >
        Konsol Diskominfo (Superadmin) &rarr;
      </a>

      <a 
        href="/" 
        target="_blank"
        class="block px-4 py-2.5 rounded-xl text-xs font-bold text-gray-300 hover:bg-[#1d2232] hover:text-[#F08300] transition"
      >
        Buka Portal Publik &rarr;
      </a>
    </nav>

    <!-- Sidebar User Footer -->
    <div class="p-4 bg-[#0d1017] border-t border-[#262a38]">
      <div class="flex items-center justify-between">
        <div class="flex-1 min-w-0 pr-2">
          <div class="text-xs font-bold text-white truncate">{currentUser?.name || 'Petugas Satpol PP'}</div>
          <div class="text-[10px] text-[#F08300] font-semibold truncate">{currentUser?.bidang || 'Penyelenggara Ketertiban'}</div>
        </div>
        <a 
          href="/admin/login" 
          class="bg-[#1f2433] hover:bg-rose-900/60 text-gray-300 hover:text-white text-[11px] font-bold px-2.5 py-1.5 rounded-lg border border-gray-700 transition"
          title="Keluar / Ganti Akun"
        >
          Logout
        </a>
      </div>
    </div>
  </aside>

  <!-- ============================================================== -->
  <!-- MAIN CONTENT AREA -->
  <!-- ============================================================== -->
  <main class="flex-1 flex flex-col min-w-0 overflow-y-auto">

    <!-- Top Bar Navigation (Responsive) -->
    <header class="bg-white border-b border-gray-200/80 px-4 sm:px-8 py-4 flex items-center justify-between sticky top-0 z-30 shadow-xs">
      <div class="flex items-center gap-3">
        <!-- Mobile Drawer Hamburger Toggle Button -->
        <button 
          class="lg:hidden p-2 rounded-lg bg-gray-100 hover:bg-gray-200 text-[#101116] transition"
          on:click={() => sidebarOpen = true}
          aria-label="Buka Menu Sidebar"
        >
          <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
          </svg>
        </button>

        <div>
          <span class="text-[11px] font-bold text-[#F08300] uppercase tracking-wider">
            Provinsi Sulawesi Selatan &bull; Praja Wibawa
          </span>
          <h2 class="text-lg md:text-xl font-extrabold text-[#101116] tracking-tight">
            {#if currentTab === 'overview'}Ikhtisar Sistem &amp; Ringkasan
            {:else if currentTab === 'complaints'}Manajemen Aduan Warga SIAP TERTIB
            {:else if currentTab === 'articles'}Pengelolaan Warta &amp; Publikasi
            {:else if currentTab === 'documents'}Tata Kelola Dokumen Publik
            {:else if currentTab === 'gallery'}Manajemen Galeri Kegiatan
            {/if}
          </h2>
        </div>
      </div>

      <div class="flex items-center gap-2.5">
        <a 
          href="/" 
          target="_blank"
          class="hidden sm:inline-flex items-center gap-1.5 bg-[#f9f4e8] hover:bg-[#eee9db] text-[#101116] font-bold text-xs uppercase px-4 py-2.5 rounded-lg border border-[#eee9db] transition"
        >
          <span>Portal Publik</span>
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" /></svg>
        </a>
        <div class="bg-gray-100 border border-gray-200 px-3 py-1.5 rounded-lg text-right hidden md:block">
          <div class="text-[10px] text-[#6f7775] font-semibold uppercase">Status Gateway</div>
          <div class="text-xs font-bold text-emerald-700">Aktif (Port 8080)</div>
        </div>
      </div>
    </header>

    <!-- Sub-page Content Wrapper -->
    <div class="p-4 sm:p-6 lg:p-8 space-y-6 max-w-7xl w-full mx-auto">

      <!-- ============================================================== -->
      <!-- TAB: OVERVIEW (IKHTISAR CEPAT) -->
      <!-- ============================================================== -->
      {#if currentTab === 'overview'}
        <div class="space-y-6">
          <!-- Metric Ribbon Cards -->
          <div class="grid grid-cols-2 lg:grid-cols-5 gap-4">
            <div class="bg-white border border-gray-200/90 rounded-2xl p-5 shadow-xs">
              <div class="text-[11px] font-bold text-[#6f7775] uppercase">Total Aduan Masuk</div>
              <div class="text-3xl font-extrabold text-[#101116] mt-2">{totalCount}</div>
              <div class="text-[10px] text-gray-400 mt-1">Registrasi SIAP TERTIB</div>
            </div>
            <div class="bg-white border border-gray-200/90 rounded-2xl p-5 shadow-xs">
              <div class="text-[11px] font-bold text-amber-700 uppercase">Menunggu Verifikasi</div>
              <div class="text-3xl font-extrabold text-amber-700 mt-2">{pendingCount}</div>
              <div class="text-[10px] text-amber-600 mt-1">Perlu dicek bukti foto</div>
            </div>
            <div class="bg-white border border-gray-200/90 rounded-2xl p-5 shadow-xs">
              <div class="text-[11px] font-bold text-blue-700 uppercase">Disposisi Bidang</div>
              <div class="text-3xl font-extrabold text-blue-700 mt-2">{disposedCount}</div>
              <div class="text-[10px] text-blue-600 mt-1">Diteruskan ke PPNS/Pleton</div>
            </div>
            <div class="bg-white border border-gray-200/90 rounded-2xl p-5 shadow-xs">
              <div class="text-[11px] font-bold text-purple-700 uppercase">Penindakan Patroli</div>
              <div class="text-3xl font-extrabold text-purple-700 mt-2">{inProgressCount}</div>
              <div class="text-[10px] text-purple-600 mt-1">Personel giat lapangan</div>
            </div>
            <div class="bg-white border border-gray-200/90 rounded-2xl p-5 shadow-xs col-span-2 lg:col-span-1">
              <div class="text-[11px] font-bold text-emerald-700 uppercase">Selesai (BAP Terbit)</div>
              <div class="text-3xl font-extrabold text-emerald-700 mt-2">{completedCount}</div>
              <div class="text-[10px] text-emerald-600 mt-1">Tiket tuntas &amp; terdokumentasi</div>
            </div>
          </div>

          <!-- Quick Action Panels -->
          <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
            <div class="bg-white border border-gray-200/90 rounded-2xl p-6 shadow-xs space-y-4">
              <div class="text-xs font-bold text-[#F08300] uppercase tracking-wider">Aduan Mendesak</div>
              <h3 class="text-base font-extrabold text-[#101116]">Tindak Lanjut SIAP TERTIB</h3>
              <p class="text-xs text-[#6f7775] leading-relaxed">
                Terdapat <strong>{pendingCount} laporan</strong> aduan masyarakat yang menunggu verifikasi keabsahan laporan.
              </p>
              <button 
                on:click={() => currentTab = 'complaints'}
                class="w-full bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase py-3 rounded-xl transition"
              >
                Buka Meja Aduan &rarr;
              </button>
            </div>

            <div class="bg-white border border-gray-200/90 rounded-2xl p-6 shadow-xs space-y-4">
              <div class="text-xs font-bold text-blue-700 uppercase tracking-wider">Publikasi Humas</div>
              <h3 class="text-base font-extrabold text-[#101116]">Tulis Artikel / Warta Baru</h3>
              <p class="text-xs text-[#6f7775] leading-relaxed">
                Publikasikan hasil giat patroli penertiban atau informasi penting dengan editor berita standar penerbitan resmi.
              </p>
              <button 
                on:click={() => { currentTab = 'articles'; isWritingArticle = true; }}
                class="w-full bg-[#2b2f3f] hover:bg-[#101116] text-white font-bold text-xs uppercase py-3 rounded-xl transition"
              >
                Mulai Menulis Berita &rarr;
              </button>
            </div>

            <div class="bg-white border border-gray-200/90 rounded-2xl p-6 shadow-xs space-y-4">
              <div class="text-xs font-bold text-emerald-700 uppercase tracking-wider">Transparansi Publik</div>
              <h3 class="text-base font-extrabold text-[#101116]">Unggah Dokumen Perencanaan</h3>
              <p class="text-xs text-[#6f7775] leading-relaxed">
                Unggah dokumen Renja, Renstra, DPA, atau Peraturan Daerah untuk diakses terbuka oleh masyarakat.
              </p>
              <button 
                on:click={() => { currentTab = 'documents'; showDocModal = true; }}
                class="w-full bg-emerald-700 hover:bg-emerald-800 text-white font-bold text-xs uppercase py-3 rounded-xl transition"
              >
                Unggah Dokumen PDF &rarr;
              </button>
            </div>
          </div>
        </div>
      {/if}

      <!-- ============================================================== -->
      <!-- TAB: COMPLAINTS (MANAJEMEN ADUAN SIAP TERTIB) -->
      <!-- ============================================================== -->
      {#if currentTab === 'complaints'}
        <div class="space-y-6">
          <!-- Filter & Search Controls -->
          <div class="bg-white border border-gray-200/90 rounded-2xl p-5 shadow-xs flex flex-col md:flex-row justify-between items-stretch md:items-center gap-4">
            <div class="flex flex-wrap gap-2">
              <button 
                class={`px-3.5 py-2 text-xs font-bold uppercase rounded-lg transition ${statusFilter === 'ALL' ? 'bg-[#101116] text-white' : 'bg-gray-100 text-gray-700 hover:bg-gray-200'}`}
                on:click={() => statusFilter = 'ALL'}
              >
                Semua ({totalCount})
              </button>
              <button 
                class={`px-3.5 py-2 text-xs font-bold uppercase rounded-lg transition ${statusFilter === 'MENUNGGU_VERIFIKASI' ? 'bg-amber-600 text-white' : 'bg-amber-50 text-amber-800 hover:bg-amber-100'}`}
                on:click={() => statusFilter = 'MENUNGGU_VERIFIKASI'}
              >
                Verifikasi ({pendingCount})
              </button>
              <button 
                class={`px-3.5 py-2 text-xs font-bold uppercase rounded-lg transition ${statusFilter === 'DISPOSISI_BIDANG' ? 'bg-blue-600 text-white' : 'bg-blue-50 text-blue-800 hover:bg-blue-100'}`}
                on:click={() => statusFilter = 'DISPOSISI_BIDANG'}
              >
                Disposisi ({disposedCount})
              </button>
              <button 
                class={`px-3.5 py-2 text-xs font-bold uppercase rounded-lg transition ${statusFilter === 'PENINDAKAN_LAPANGAN' ? 'bg-purple-600 text-white' : 'bg-purple-50 text-purple-800 hover:bg-purple-100'}`}
                on:click={() => statusFilter = 'PENINDAKAN_LAPANGAN'}
              >
                Penindakan ({inProgressCount})
              </button>
              <button 
                class={`px-3.5 py-2 text-xs font-bold uppercase rounded-lg transition ${statusFilter === 'SELESAI' ? 'bg-emerald-600 text-white' : 'bg-emerald-50 text-emerald-800 hover:bg-emerald-100'}`}
                on:click={() => statusFilter = 'SELESAI'}
              >
                Selesai ({completedCount})
              </button>
            </div>

            <div class="w-full md:w-72">
              <input 
                type="text" 
                bind:value={searchQuery}
                placeholder="Cari nomor tiket, nama, atau lokasi..."
                class="w-full border border-gray-300 rounded-lg px-3.5 py-2 text-xs focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] focus:outline-none"
              />
            </div>
          </div>

          <!-- Complaints Table -->
          <div class="bg-white border border-gray-200/90 rounded-2xl shadow-xs overflow-hidden">
            <div class="overflow-x-auto">
              <table class="w-full text-left text-xs">
                <thead class="bg-gray-50 border-b border-gray-200 text-[#6f7775] uppercase text-[10px] tracking-wider">
                  <tr>
                    <th class="p-4">No. Tiket Resi</th>
                    <th class="p-4">Pelapor</th>
                    <th class="p-4">Jenis Pelanggaran</th>
                    <th class="p-4">Wilayah &amp; Lokasi</th>
                    <th class="p-4">Status SOP</th>
                    <th class="p-4">Disposisi</th>
                    <th class="p-4 text-right">Tindakan</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-100">
                  {#each filteredComplaints as comp}
                    {@const st = getStatusLabel(comp.status)}
                    <tr class="hover:bg-gray-50/70 transition">
                      <td class="p-4 font-mono font-bold text-[#101116]">
                        {comp.ticket_code}
                        <div class="text-[10px] text-[#6f7775] font-sans font-normal mt-0.5">
                          {new Date(comp.created_at).toLocaleDateString('id-ID', { day: 'numeric', month: 'short', hour: '2-digit', minute: '2-digit' })} WITA
                        </div>
                      </td>
                      <td class="p-4">
                        <div class="font-bold text-[#101116]">
                          {comp.is_anonymous ? (comp.masked_name || 'Pelapor Anonim') : comp.reporter_name}
                        </div>
                        {#if comp.is_anonymous}
                          <span class="inline-block bg-amber-100 text-amber-900 border border-amber-300 text-[9px] font-bold px-1.5 py-0.5 rounded mt-0.5">
                            Anonim Terlindungi
                          </span>
                        {:else}
                          <div class="text-[11px] text-[#6f7775]">{comp.reporter_phone || 'Tanpa Kontak'}</div>
                        {/if}
                      </td>
                      <td class="p-4 font-semibold text-[#101116] max-w-xs">{comp.violation_type}</td>
                      <td class="p-4 text-[#6f7775]">
                        <div class="font-semibold text-[#101116]">{comp.regency_name}</div>
                        <div class="text-[11px] truncate max-w-[180px]">{comp.address_detail}</div>
                      </td>
                      <td class="p-4">
                        <span class={`font-bold text-[10px] uppercase px-2.5 py-1 rounded-md border ${st.bg}`}>
                          {st.text}
                        </span>
                      </td>
                      <td class="p-4 text-[11px] text-[#6f7775] font-medium max-w-[150px] truncate">
                        {comp.disposed_to || '-'}
                      </td>
                      <td class="p-4 text-right">
                        <button 
                          on:click={() => openComplaintDetail(comp)}
                          class="bg-[#2b2f3f] hover:bg-[#101116] text-white font-bold text-[11px] uppercase px-3 py-1.5 rounded-lg transition"
                        >
                          Tindak Lanjut &rarr;
                        </button>
                      </td>
                    </tr>
                  {:else}
                    <tr>
                      <td colspan="7" class="p-10 text-center text-gray-400 text-xs">
                        Tidak ditemukan aduan masyarakat dengan filter yang dipilih.
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      {/if}

      <!-- ============================================================== -->
      <!-- TAB: ARTICLES (KOMPASIANA-GRADE CMS PUBLISHER) -->
      <!-- ============================================================== -->
      {#if currentTab === 'articles'}
        <div class="space-y-6">
          {#if isWritingArticle}
            <!-- WRITING MODE (KOMPASIANA STYLE) -->
            <div class="bg-white border border-gray-200/90 rounded-2xl p-6 md:p-8 shadow-xs space-y-6">
              <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 border-b border-gray-100 pb-4">
                <div>
                  <div class="text-[#F08300] font-bold text-xs uppercase tracking-wider">Redaksi &amp; Publikasi Warta</div>
                  <h3 class="text-xl font-extrabold text-[#101116] mt-0.5">Editor Penulisan Berita Resmi</h3>
                </div>
                <div class="flex items-center gap-2">
                  <button 
                    on:click={() => isWritingArticle = false}
                    class="px-4 py-2 border border-gray-300 text-gray-700 text-xs font-bold uppercase rounded-lg hover:bg-gray-100 transition"
                  >
                    Batal
                  </button>
                  <button 
                    on:click={handlePublishArticle}
                    class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-6 py-2 rounded-lg transition shadow-sm"
                  >
                    Terbitkan Berita
                  </button>
                </div>
              </div>

              <!-- Article Editor Fields -->
              <div class="space-y-5">
                <!-- Title & Slug -->
                <div>
                  <label for="editor-article-title" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                    Judul Berita (Headline)
                  </label>
                  <input 
                    id="editor-article-title"
                    type="text" 
                    value={articleTitle}
                    on:input={handleTitleInput}
                    placeholder="Masukkan judul berita yang informatif dan akurat..."
                    class="w-full border border-gray-300 rounded-xl p-3 text-sm font-bold text-[#101116] focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] focus:outline-none"
                  />
                  <div class="flex justify-between items-center text-[11px] text-[#6f7775] mt-1.5">
                    <span class="font-mono">Slug: /berita/{articleSlug || 'judul-berita'}</span>
                    <span>{articleTitle.length} karakter</span>
                  </div>
                </div>

                <!-- Metadata Row -->
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label for="editor-article-cat" class="block text-xs font-bold uppercase text-[#101116] mb-1">Kategori Warta</label>
                    <select id="editor-article-cat" bind:value={articleCategory} class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white">
                      <option value="Kegiatan">Kegiatan &amp; Patroli</option>
                      <option value="Penertiban Aset">Penertiban Aset Daerah</option>
                      <option value="Pengumuman">Pengumuman Resmi</option>
                      <option value="Opini / Rilis">Rilis Humas Resmi</option>
                    </select>
                  </div>

                  <div>
                    <label for="editor-article-author" class="block text-xs font-bold uppercase text-[#101116] mb-1">Penulis / Unit Peliput</label>
                    <input id="editor-article-author" type="text" bind:value={articleAuthor} class="w-full border border-gray-300 rounded-lg p-2.5 text-xs" />
                  </div>
                </div>

                <!-- Featured Image File Upload -->
                <div>
                  <label for="editor-article-thumbnail-file" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                    Foto Utama / Thumbnail (Upload dari Perangkat)
                  </label>
                  <div class="border-2 border-dashed border-gray-300 rounded-xl p-4 flex flex-col md:flex-row items-center gap-4 bg-gray-50/50">
                    {#if articleThumbnailPreview}
                      <img src={articleThumbnailPreview} alt="Preview" class="w-32 h-24 object-cover rounded-lg border border-gray-300 flex-shrink-0" />
                    {:else}
                      <div class="w-32 h-24 bg-gray-200 rounded-lg flex items-center justify-center text-[10px] text-gray-500 font-bold uppercase text-center p-2">
                        Belum Ada Foto
                      </div>
                    {/if}
                    <div class="space-y-1 text-xs">
                      <input 
                        id="editor-article-thumbnail-file"
                        type="file" 
                        accept="image/*" 
                        on:change={handleArticleImageUpload} 
                        class="text-xs text-gray-600 file:mr-3 file:py-1.5 file:px-3 file:rounded-lg file:border-0 file:text-xs file:font-bold file:bg-[#2b2f3f] file:text-white hover:file:bg-[#101116]" 
                      />
                      <div class="text-[11px] text-[#6f7775]">Format JPG, PNG, atau WebP disarankan beresolusi minimal 800x500px.</div>
                    </div>
                  </div>
                </div>

                <!-- Lead Paragraph / Summary -->
                <div>
                  <label for="editor-article-summary" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                    Ringkasan / Lead Paragraph (Tampil di Kartu Indeks)
                  </label>
                  <textarea 
                    id="editor-article-summary"
                    bind:value={articleSummary} 
                    rows="2" 
                    placeholder="Tuliskan 1-2 kalimat pengantar utama..."
                    class="w-full border border-gray-300 rounded-xl p-3 text-xs focus:border-[#F08300] focus:ring-1 focus:ring-[#F08300] focus:outline-none"
                  ></textarea>
                </div>

                <!-- Rich Formatting Suite with Adaptive Lists -->
                <div>
                  <label for="editor-article-content" class="block text-xs font-bold uppercase text-[#101116] mb-1.5">
                    Isi Berita Lengkap (Editor Redaksi Kompasiana-Grade)
                  </label>
                  <RichTextEditor 
                    bind:content={articleContent} 
                    placeholder="Mulai menulis naskah berita... Tekan Enter pada daftar nomor (1.) atau poin (-) untuk penomoran otomatis adaptif."
                  />
                </div>
              </div>
            </div>
          {:else}
            <!-- ARTICLES LIST VIEW -->
            <div class="bg-white border border-gray-200/90 rounded-2xl p-6 shadow-xs space-y-5">
              <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 border-b border-gray-100 pb-4">
                <div>
                  <h3 class="text-base font-extrabold text-[#101116] uppercase tracking-wide">Daftar Warta Berita &amp; Publikasi</h3>
                  <p class="text-xs text-[#6f7775] mt-0.5">Kelola seluruh berita yang tampil pada beranda dan indeks warta portal.</p>
                </div>
                <button 
                  on:click={() => isWritingArticle = true}
                  class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition shadow-xs"
                >
                  + Tulis Artikel Baru
                </button>
              </div>

              <div class="overflow-x-auto">
                <table class="w-full text-left text-xs">
                  <thead class="bg-gray-50 border-b border-gray-200 text-[#6f7775] uppercase text-[10px] tracking-wider">
                    <tr>
                      <th class="p-3.5">Thumbnail</th>
                      <th class="p-3.5">Judul Berita</th>
                      <th class="p-3.5">Kategori</th>
                      <th class="p-3.5">Penulis</th>
                      <th class="p-3.5">Dibaca</th>
                      <th class="p-3.5 text-right">Opsi</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100">
                    {#each $articlesStore as art}
                      <tr class="hover:bg-gray-50/70 transition">
                        <td class="p-3.5">
                          <img src={art.thumbnail} alt={art.title} class="w-16 h-12 object-cover rounded-lg border border-gray-200" />
                        </td>
                        <td class="p-3.5 font-bold text-[#101116] max-w-sm">
                          <a href={`/berita/${art.slug}`} target="_blank" class="hover:text-[#F08300] transition">
                            {art.title}
                          </a>
                          <div class="text-[10px] text-[#6f7775] font-normal mt-0.5 font-mono">/berita/{art.slug}</div>
                        </td>
                        <td class="p-3.5">
                          <span class="bg-gray-100 text-[#101116] text-[10px] font-bold uppercase px-2 py-0.5 rounded border border-gray-300">
                            {art.category}
                          </span>
                        </td>
                        <td class="p-3.5 text-[#6f7775]">{art.author_name}</td>
                        <td class="p-3.5 font-semibold text-[#101116]">{art.view_count}x</td>
                        <td class="p-3.5 text-right space-x-2">
                          <a 
                            href={`/berita/${art.slug}`} 
                            target="_blank" 
                            class="text-xs font-bold text-blue-700 hover:underline"
                          >
                            Lihat
                          </a>
                          <button 
                            on:click={() => handleDeleteArticle(art.id)}
                            class="text-xs font-bold text-rose-700 hover:underline"
                          >
                            Hapus
                          </button>
                        </td>
                      </tr>
                    {/each}
                  </tbody>
                </table>
              </div>
            </div>
          {/if}
        </div>
      {/if}

      <!-- ============================================================== -->
      <!-- TAB: DOCUMENTS (REAL FILE UPLOADS CMS) -->
      <!-- ============================================================== -->
      {#if currentTab === 'documents'}
        <div class="space-y-6">
          <div class="bg-white border border-gray-200/90 rounded-2xl p-6 shadow-xs space-y-5">
            <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 border-b border-gray-100 pb-4">
              <div>
                <h3 class="text-base font-extrabold text-[#101116] uppercase tracking-wide">Transparansi Dokumen Publik</h3>
                <p class="text-xs text-[#6f7775] mt-0.5">Kelola berkas Renja, Renstra, DPA, LAKIP, dan Peraturan Daerah resmi.</p>
              </div>
              <button 
                on:click={() => showDocModal = true}
                class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition shadow-xs"
              >
                + Unggah Dokumen PDF
              </button>
            </div>

            <!-- Upload Modal -->
            {#if showDocModal}
              <div class="bg-[#f9f4e8] border border-[#eee9db] rounded-2xl p-6 space-y-4">
                <div class="text-xs font-bold text-[#101116] uppercase tracking-wider">Form Unggah Berkas Resmi</div>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <!-- File Selector Input -->
                  <div class="md:col-span-2">
                    <label for="doc-file-input" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                      Pilih Berkas Dokumen (PDF, DOCX, XLSX)
                    </label>
                    <input 
                      id="doc-file-input"
                      type="file" 
                      accept=".pdf,.doc,.docx,.xls,.xlsx"
                      on:change={handleDocFileUpload}
                      class="w-full text-xs text-gray-700 border border-gray-300 rounded-lg p-2 bg-white file:mr-3 file:py-1 file:px-3 file:rounded-md file:border-0 file:text-xs file:font-bold file:bg-[#2b2f3f] file:text-white"
                    />
                    {#if uploadedFileName}
                      <div class="text-[11px] text-emerald-700 font-semibold mt-1">
                        Berkas siap: {uploadedFileName} ({docSize})
                      </div>
                    {/if}
                  </div>

                  <div>
                    <label for="doc-title-input" class="block text-xs font-bold uppercase text-[#101116] mb-1">Nama / Judul Dokumen</label>
                    <input id="doc-title-input" type="text" bind:value={docTitle} placeholder="Contoh: Rencana Kerja Satpol PP TA 2026" class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white" />
                  </div>

                  <div>
                    <label for="doc-number-input" class="block text-xs font-bold uppercase text-[#101116] mb-1">Nomor Registrasi / SK</label>
                    <input id="doc-number-input" type="text" bind:value={docNumber} placeholder="Contoh: 050/112/SATPOL/2026" class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white" />
                  </div>

                  <div>
                    <label for="doc-category-select" class="block text-xs font-bold uppercase text-[#101116] mb-1">Kategori Dokumen</label>
                    <select id="doc-category-select" bind:value={docCategory} class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white">
                      <option value="renja">Rencana Kerja (Renja)</option>
                      <option value="renstra">Rencana Strategis (Renstra)</option>
                      <option value="dpa">Dokumen Pelaksanaan Anggaran (DPA)</option>
                      <option value="lakip">LAKIP / LKjIP</option>
                      <option value="keuangan">Laporan Keuangan</option>
                      <option value="perda">Peraturan Daerah (Perda)</option>
                      <option value="sop">Standar Operasional (SOP)</option>
                    </select>
                  </div>

                  <div>
                    <label for="doc-year-select" class="block text-xs font-bold uppercase text-[#101116] mb-1">Tahun Anggaran</label>
                    <select id="doc-year-select" bind:value={docYear} class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white">
                      <option value={2026}>2026</option>
                      <option value={2025}>2025</option>
                      <option value={2024}>2024</option>
                      <option value={2023}>2023</option>
                      <option value={2022}>2022</option>
                      <option value={2021}>2021</option>
                    </select>
                  </div>
                </div>

                <div class="flex justify-end gap-2 pt-2">
                  <button on:click={() => showDocModal = false} class="px-4 py-2 border border-gray-300 rounded-lg text-xs font-bold text-gray-700 bg-white">Batal</button>
                  <button on:click={handleSaveDocument} class="px-5 py-2 bg-[#F08300] text-white text-xs font-bold uppercase rounded-lg hover:bg-[#d97400]">Simpan Dokumen</button>
                </div>
              </div>
            {/if}

            <div class="overflow-x-auto">
              <table class="w-full text-left text-xs">
                <thead class="bg-gray-50 border-b border-gray-200 text-[#6f7775] uppercase text-[10px] tracking-wider">
                  <tr>
                    <th class="p-3.5">Nama Dokumen</th>
                    <th class="p-3.5">No. Registrasi</th>
                    <th class="p-3.5">Kategori</th>
                    <th class="p-3.5">Tahun</th>
                    <th class="p-3.5">Ukuran</th>
                    <th class="p-3.5 text-right">Opsi</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-gray-100">
                  {#each $documentsStore as doc}
                    <tr class="hover:bg-gray-50/70 transition">
                      <td class="p-3.5 font-bold text-[#101116]">{doc.title}</td>
                      <td class="p-3.5 font-mono text-[#6f7775]">{doc.doc_number || '-'}</td>
                      <td class="p-3.5">
                        <span class="bg-[#f9f4e8] text-[#101116] text-[10px] font-bold uppercase px-2 py-0.5 rounded border border-[#eee9db]">
                          {doc.category}
                        </span>
                      </td>
                      <td class="p-3.5 font-semibold text-[#101116]">{doc.year}</td>
                      <td class="p-3.5 text-[#6f7775]">{doc.file_size}</td>
                      <td class="p-3.5 text-right space-x-2">
                        <a href={doc.file_url} target="_blank" class="text-xs font-bold text-blue-700 hover:underline">Unduh</a>
                        <button on:click={() => handleDeleteDocument(doc.id)} class="text-xs font-bold text-rose-700 hover:underline">Hapus</button>
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            </div>
          </div>
        </div>
      {/if}

      <!-- ============================================================== -->
      <!-- TAB: GALLERY (CMS FOTO DOKUMENTASI KEGIATAN) -->
      <!-- ============================================================== -->
      {#if currentTab === 'gallery'}
        <div class="space-y-6">
          <div class="bg-white border border-gray-200/90 rounded-2xl p-6 shadow-xs space-y-5">
            <div class="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4 border-b border-gray-100 pb-4">
              <div>
                <h3 class="text-base font-extrabold text-[#101116] uppercase tracking-wide">CMS Galeri Multimedia Dokumentasi</h3>
                <p class="text-xs text-[#6f7775] mt-0.5">Kelola foto kegiatan apel, operasi penertiban, dan koordinasi yang tampil pada publik.</p>
              </div>
              <button 
                on:click={() => showGalleryModal = true}
                class="bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase px-5 py-2.5 rounded-xl transition shadow-xs"
              >
                + Tambah Foto Galeri
              </button>
            </div>

            <!-- Upload Modal -->
            {#if showGalleryModal}
              <div class="bg-[#f9f4e8] border border-[#eee9db] rounded-2xl p-6 space-y-4">
                <div class="text-xs font-bold text-[#101116] uppercase tracking-wider">Form Tambah Foto Kegiatan</div>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <!-- File Selector -->
                  <div class="md:col-span-2">
                    <label for="gallery-file-input" class="block text-xs font-bold uppercase text-[#101116] mb-1">
                      Pilih Foto dari Perangkat
                    </label>
                    <input 
                      id="gallery-file-input"
                      type="file" 
                      accept="image/*"
                      on:change={handleGalleryImageUpload}
                      class="w-full text-xs text-gray-700 border border-gray-300 rounded-lg p-2 bg-white file:mr-3 file:py-1 file:px-3 file:rounded-md file:border-0 file:text-xs file:font-bold file:bg-[#2b2f3f] file:text-white"
                    />
                    {#if galleryPreview}
                      <img src={galleryPreview} alt="Preview Foto" class="mt-2 w-48 h-32 object-cover rounded-xl border border-gray-300" />
                    {/if}
                  </div>

                  <div>
                    <label for="gallery-title-input" class="block text-xs font-bold uppercase text-[#101116] mb-1">Judul / Keterangan Giat</label>
                    <input id="gallery-title-input" type="text" bind:value={galleryTitle} placeholder="Contoh: Apel Kesiapsiagaan Anggota Satpol PP" class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white" />
                  </div>

                  <div>
                    <label for="gallery-category-select" class="block text-xs font-bold uppercase text-[#101116] mb-1">Kategori Dokumentasi</label>
                    <select id="gallery-category-select" bind:value={galleryCategory} class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white">
                      <option value="Penertiban">Operasi &amp; Penertiban</option>
                      <option value="Apel & Upacara">Apel &amp; Upacara Resmi</option>
                      <option value="Koordinasi">Koordinasi &amp; Sosialisasi</option>
                      <option value="Kegiatan">Giat Harian</option>
                    </select>
                  </div>
                </div>

                <div class="flex justify-end gap-2 pt-2">
                  <button on:click={() => showGalleryModal = false} class="px-4 py-2 border border-gray-300 rounded-lg text-xs font-bold text-gray-700 bg-white">Batal</button>
                  <button on:click={handleSaveGallery} class="px-5 py-2 bg-[#F08300] text-white text-xs font-bold uppercase rounded-lg hover:bg-[#d97400]">Simpan ke Galeri</button>
                </div>
              </div>
            {/if}

            <!-- Gallery Cards Grid -->
            <div class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
              {#each $galleryStore as item}
                <div class="border border-gray-200 rounded-xl overflow-hidden bg-white shadow-xs group flex flex-col justify-between">
                  <div>
                    <div class="relative h-40 bg-gray-100 overflow-hidden">
                      <img src={item.image_url} alt={item.title} class="w-full h-full object-cover group-hover:scale-105 transition duration-300" />
                      <span class="absolute top-2 left-2 bg-[#101116]/85 text-[#F08300] font-bold text-[9px] uppercase px-2 py-0.5 rounded">
                        {item.category}
                      </span>
                    </div>
                    <div class="p-3">
                      <div class="text-[10px] text-[#6f7775]">{item.date}</div>
                      <h4 class="text-xs font-bold text-[#101116] line-clamp-2 mt-0.5">{item.title}</h4>
                    </div>
                  </div>
                  <div class="p-3 pt-0 border-t border-gray-50 flex justify-end">
                    <button 
                      on:click={() => handleDeleteGallery(item.id)}
                      class="text-[11px] font-bold text-rose-700 hover:underline"
                    >
                      Hapus Foto
                    </button>
                  </div>
                </div>
              {/each}
            </div>
          </div>
        </div>
      {/if}

    </div>
  </main>
</div>

<!-- ============================================================== -->
<!-- SOP COMPLAINT HANDLING DRAWER / MODAL -->
<!-- ============================================================== -->
{#if isEditingComplaint && selectedComplaint}
  <div 
    class="fixed inset-0 z-50 bg-black/60 flex items-center justify-center p-4 backdrop-blur-sm"
    role="dialog"
    aria-modal="true"
  >
    <div class="bg-white rounded-2xl max-w-2xl w-full border border-gray-300 shadow-2xl overflow-hidden max-h-[90vh] flex flex-col">
      <!-- Modal Header -->
      <div class="bg-[#101116] text-white p-5 border-b-2 border-[#F08300] flex justify-between items-center">
        <div>
          <div class="text-[10px] text-[#F08300] font-bold uppercase tracking-wider">Tindak Lanjut SOP SIAP TERTIB</div>
          <h3 class="text-base font-bold mt-0.5">Tiket #{selectedComplaint.ticket_code}</h3>
        </div>
        <button 
          on:click={() => isEditingComplaint = false}
          class="text-gray-400 hover:text-white font-bold text-lg px-2 rounded-lg"
          aria-label="Tutup form"
        >
          &times;
        </button>
      </div>

      <!-- Modal Body -->
      <div class="p-6 overflow-y-auto space-y-5 text-xs">
        <!-- Reporter & Location Card -->
        <div class="bg-gray-50 border border-gray-200 rounded-xl p-4 space-y-3">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <span class="text-[#6f7775] uppercase text-[10px] font-bold block mb-1">Identitas Pelapor:</span>
              {#if selectedComplaint.is_anonymous}
                {#if unmaskRequested}
                  <!-- Unmasked officer view -->
                  <div class="bg-amber-50/70 border border-amber-200 rounded-lg p-2.5 space-y-1">
                    <div class="flex items-center justify-between">
                      <span class="text-[10px] font-bold text-amber-900 uppercase">Data Terbuka (Akses Petugas)</span>
                      <span class="text-[9px] bg-amber-200 text-amber-900 px-1.5 py-0.5 rounded font-bold">Audit Aktif</span>
                    </div>
                    <div class="font-bold text-[#101116] text-xs">{selectedComplaint.reporter_name}</div>
                    <div class="text-[11px] font-mono text-[#101116]">NIK: {selectedComplaint.reporter_nik || '-'}</div>
                    <div class="text-[11px] text-[#101116]">Telp: {selectedComplaint.reporter_phone || '-'}</div>
                    <div class="text-[10px] text-[#6f7775] italic pt-1 border-t border-amber-200/60">
                      Pelepasan masking telah tercatat dalam sistem audit forensik.
                    </div>
                  </div>
                {:else}
                  <!-- Masked default view -->
                  <div class="bg-gray-100 border border-gray-200 rounded-lg p-2.5 space-y-1">
                    <div class="font-bold text-[#101116]">{selectedComplaint.masked_name || 'Ah*** P*** (Anonim Terlindungi)'}</div>
                    <div class="text-[11px] font-mono text-[#6f7775]">NIK: {selectedComplaint.masked_nik || '7371**********12'}</div>
                    <div class="text-[11px] text-[#6f7775]">Telp: {selectedComplaint.masked_phone || '0812****8901'}</div>
                    <button 
                      type="button" 
                      on:click={handleUnmaskIdentity} 
                      class="mt-2 inline-flex items-center gap-1.5 px-3 py-1.5 bg-amber-600 hover:bg-amber-700 text-white text-[10px] font-bold uppercase rounded-lg transition"
                    >
                      Buka Masking Identitas (Hak Akses Petugas)
                    </button>
                  </div>
                {/if}
              {:else}
                <div class="font-bold text-[#101116]">{selectedComplaint.reporter_name}</div>
                <div class="text-[11px] text-[#6f7775]">NIK: {selectedComplaint.reporter_nik || '-'}</div>
                <div class="text-[11px] text-[#6f7775]">Telp: {selectedComplaint.reporter_phone || '-'}</div>
              {/if}
            </div>

            <div>
              <span class="text-[#6f7775] uppercase text-[10px] font-bold block mb-1">Wilayah &amp; Waktu:</span>
              <div class="font-bold text-[#101116]">{selectedComplaint.regency_name}</div>
              <div class="text-[11px] text-[#6f7775]">{selectedComplaint.district_name}, {selectedComplaint.village_name}</div>
              <div class="text-[11px] text-[#6f7775]">Waktu: {new Date(selectedComplaint.created_at).toLocaleDateString('id-ID')}</div>
            </div>
          </div>

          <div class="pt-2 border-t border-gray-200">
            <span class="text-[#6f7775] uppercase text-[10px] font-bold">Alamat Detail Kejadian:</span>
            <div class="text-[#101116] font-medium mt-0.5">{selectedComplaint.address_detail}</div>
          </div>

          <!-- Evidence Photo (Strictly Images Only) -->
          {#if selectedComplaint.evidence_image_url}
            <div class="pt-3 border-t border-gray-200">
              <span class="text-[#6f7775] uppercase text-[10px] font-bold block mb-1.5">Bukti Foto Pelanggaran Lapangan:</span>
              <div>
                <a href={selectedComplaint.evidence_image_url} target="_blank" class="inline-block group">
                  <img 
                    src={selectedComplaint.evidence_image_url} 
                    alt="Bukti Foto Aduan SIAP TERTIB" 
                    class="w-56 h-36 object-cover rounded-xl border border-gray-300 group-hover:opacity-90 transition shadow-xs" 
                  />
                  <span class="text-[10px] text-blue-700 font-semibold block mt-1 hover:underline">Lihat Foto Ukuran Penuh &rarr;</span>
                </a>
              </div>
            </div>
          {/if}
        </div>

        <!-- Violation & Description -->
        <div>
          <span class="text-[#6f7775] uppercase text-[10px] font-bold">Uraian Pelanggaran yang Dilaporkan:</span>
          <div class="font-bold text-sm text-[#101116] mt-0.5">{selectedComplaint.violation_type}</div>
          <div class="p-3 bg-white border border-gray-200 rounded-lg mt-1 text-[#2b2f3f] leading-relaxed">
            {selectedComplaint.description}
          </div>
        </div>

        <!-- SOP Status Transition Controls -->
        <div class="space-y-3 pt-3 border-t border-gray-200">
          <div class="font-bold text-xs uppercase text-[#101116]">Perbarui Progres SOP Penanganan:</div>

          <div>
            <label for="complaint-status-select" class="block text-[11px] font-bold uppercase text-[#6f7775] mb-1">Status Prosedur</label>
            <select id="complaint-status-select" bind:value={editStatus} class="w-full border border-gray-300 rounded-lg p-2.5 text-xs bg-white font-bold">
              <option value="MENUNGGU_VERIFIKASI">1. MENUNGGU VERIFIKASI (Pemeriksaan Awal)</option>
              <option value="DISPOSISI_BIDANG">2. DISPOSISI BIDANG (Diteruskan ke Unit Teknis)</option>
              <option value="PENINDAKAN_LAPANGAN">3. PENINDAKAN LAPANGAN (Patroli / Surat Peringatan)</option>
              <option value="SELESAI">4. SELESAI (Penyusunan BAP &amp; Penutupan Aduan)</option>
              <option value="DITOLAK">DITOLAK (Laporan Tidak Sesuai Kewenangan / Hoaks)</option>
            </select>
          </div>

          <div>
            <label for="complaint-disposition-input" class="block text-[11px] font-bold uppercase text-[#6f7775] mb-1">Disposisi Unit Kerja / Pleton</label>
            <input 
              id="complaint-disposition-input"
              type="text" 
              bind:value={editDisposedTo} 
              placeholder="Contoh: Bidang Ketenteraman dan Ketertiban Umum (Pleton 2 Ops)"
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs"
            />
          </div>

          <div>
            <label for="complaint-bap-notes" class="block text-[11px] font-bold uppercase text-[#6f7775] mb-1">Catatan Tindakan Lapangan / Berita Acara (BAP)</label>
            <textarea 
              id="complaint-bap-notes"
              bind:value={editActionNotes} 
              rows="3" 
              placeholder="Tuliskan nomor Berita Acara, hasil penertiban, barang bukti yang diamankan, atau alasan penolakan..."
              class="w-full border border-gray-300 rounded-lg p-2.5 text-xs"
            ></textarea>
          </div>
        </div>
      </div>

      <!-- Modal Footer -->
      <div class="p-4 bg-gray-50 border-t border-gray-200 flex justify-end gap-2">
        <button 
          on:click={() => isEditingComplaint = false}
          class="px-4 py-2 border border-gray-300 rounded-lg text-xs font-bold text-gray-700 bg-white hover:bg-gray-100"
        >
          Tutup
        </button>
        <button 
          on:click={saveComplaintUpdate}
          class="px-5 py-2 bg-[#F08300] hover:bg-[#d97400] text-white text-xs font-bold uppercase rounded-lg transition"
        >
          Simpan Pembaruan SOP
        </button>
      </div>
    </div>
  </div>
{/if}
