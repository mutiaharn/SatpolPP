<script lang="ts">
  import { onMount } from 'svelte';

  export let content = '';
  export let placeholder = 'Mulai menulis isi berita lengkap... Tekan Enter pada daftar nomor/poin untuk otomatis melanjutkan penomoran.';

  let activeTab: 'write' | 'preview' = 'write';
  let textarea: HTMLTextAreaElement;

  // Calculate word count & reading time
  $: wordCount = content ? content.trim().split(/\s+/).filter(Boolean).length : 0;
  $: readingTime = Math.max(1, Math.ceil(wordCount / 180));

  // Adaptive Enter Key handler (Auto Numbering & Auto Bullet Lists)
  function handleKeyDown(e: KeyboardEvent) {
    // Handle Keyboard Shortcuts: Ctrl+B, Ctrl+I, Ctrl+U
    if ((e.ctrlKey || e.metaKey) && !e.shiftKey && !e.altKey) {
      if (e.key === 'b' || e.key === 'B') {
        e.preventDefault();
        wrapSelection('**', '**');
        return;
      }
      if (e.key === 'i' || e.key === 'I') {
        e.preventDefault();
        wrapSelection('*', '*');
        return;
      }
      if (e.key === 'u' || e.key === 'U') {
        e.preventDefault();
        wrapSelection('<u>', '</u>');
        return;
      }
    }

    // Adaptive Enter Handler
    if (e.key === 'Enter') {
      const cursor = textarea.selectionStart;
      const textBefore = content.substring(0, cursor);
      const lastLine = textBefore.split('\n').pop() || '';

      // Check for Numbered List: e.g. "1. ", "2. ", "10. "
      const numMatch = lastLine.match(/^(\s*)(\d+)\.\s+(.*)$/);
      if (numMatch) {
        e.preventDefault();
        const indent = numMatch[1];
        const num = parseInt(numMatch[2], 10);
        const itemText = numMatch[3].trim();

        if (itemText === '') {
          // Empty item: User pressed enter twice, cancel list
          const lineStart = cursor - lastLine.length;
          content = content.substring(0, lineStart) + '\n' + content.substring(cursor);
          setTimeout(() => {
            textarea.selectionStart = textarea.selectionEnd = lineStart + 1;
          }, 0);
        } else {
          // Continue with next number
          const nextItem = `\n${indent}${num + 1}. `;
          insertAtCursor(nextItem);
        }
        return;
      }

      // Check for Bullet List: e.g. "- ", "* "
      const bulletMatch = lastLine.match(/^(\s*)([-*])\s+(.*)$/);
      if (bulletMatch) {
        e.preventDefault();
        const indent = bulletMatch[1];
        const bullet = bulletMatch[2];
        const itemText = bulletMatch[3].trim();

        if (itemText === '') {
          // Empty bullet: cancel list
          const lineStart = cursor - lastLine.length;
          content = content.substring(0, lineStart) + '\n' + content.substring(cursor);
          setTimeout(() => {
            textarea.selectionStart = textarea.selectionEnd = lineStart + 1;
          }, 0);
        } else {
          // Continue bullet list
          const nextBullet = `\n${indent}${bullet} `;
          insertAtCursor(nextBullet);
        }
        return;
      }
    }
  }

  function insertAtCursor(textToInsert: string) {
    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    content = content.substring(0, start) + textToInsert + content.substring(end);
    setTimeout(() => {
      textarea.selectionStart = textarea.selectionEnd = start + textToInsert.length;
      textarea.focus();
    }, 0);
  }

  function wrapSelection(prefix: string, suffix: string = '') {
    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const selected = content.substring(start, end);
    const replacement = prefix + selected + suffix;
    content = content.substring(0, start) + replacement + content.substring(end);
    setTimeout(() => {
      textarea.focus();
      textarea.setSelectionRange(start + prefix.length, end + prefix.length);
    }, 0);
  }

  function insertBlock(prefix: string) {
    const start = textarea.selectionStart;
    const textBefore = content.substring(0, start);
    const needsNewline = textBefore.length > 0 && !textBefore.endsWith('\n');
    const toInsert = (needsNewline ? '\n\n' : '') + prefix;
    insertAtCursor(toInsert);
  }

  // Convert raw formatted text to clean HTML preview
  function renderFormattedHtml(text: string): string {
    if (!text) return '<p class="text-gray-400 italic">Pratinjau isi berita akan tampil di sini saat Anda mengetik...</p>';

    const lines = text.split('\n');
    let html = '';
    let inOl = false;
    let inUl = false;

    for (let i = 0; i < lines.length; i++) {
      let line = lines[i];

      // Check Heading 2
      if (line.startsWith('## ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<h2 class="text-xl font-extrabold text-[#101116] mt-5 mb-2">${formatInline(line.slice(3))}</h2>`;
        continue;
      }

      // Check Heading 3
      if (line.startsWith('### ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<h3 class="text-base font-bold text-[#101116] mt-4 mb-1.5">${formatInline(line.slice(4))}</h3>`;
        continue;
      }

      // Check Blockquote
      if (line.startsWith('> ')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<blockquote class="border-l-4 border-[#F08300] pl-4 italic text-[#3d261a] my-3 bg-[#fcfbf7] py-2 rounded-r-lg">${formatInline(line.slice(2))}</blockquote>`;
        continue;
      }

      // Check Callout
      if (line.startsWith(':::callout') || line.startsWith('[Catatan Redaksi]')) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (inUl) { html += '</ul>'; inUl = false; }
        html += `<div class="bg-[#f9f4e8] border border-[#eee9db] p-4 rounded-xl my-3 text-xs font-semibold text-[#101116]">${formatInline(line.replace(/:::callout|\[Catatan Redaksi\]/, ''))}</div>`;
        continue;
      }

      // Check Numbered List Item
      const numMatch = line.match(/^\s*(\d+)\.\s+(.*)$/);
      if (numMatch) {
        if (inUl) { html += '</ul>'; inUl = false; }
        if (!inOl) { html += '<ol class="list-decimal pl-5 space-y-1 my-2 text-xs text-[#3d261a]">'; inOl = true; }
        html += `<li>${formatInline(numMatch[2])}</li>`;
        continue;
      }

      // Check Bullet List Item
      const bulletMatch = line.match(/^\s*[-*]\s+(.*)$/);
      if (bulletMatch) {
        if (inOl) { html += '</ol>'; inOl = false; }
        if (!inUl) { html += '<ul class="list-disc pl-5 space-y-1 my-2 text-xs text-[#3d261a]">'; inUl = true; }
        html += `<li>${formatInline(bulletMatch[1])}</li>`;
        continue;
      }

      // Close open lists
      if (inOl) { html += '</ol>'; inOl = false; }
      if (inUl) { html += '</ul>'; inUl = false; }

      // Paragraph
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
      .replace(/\[(.*?)\]\((.*?)\)/g, '<a href="$2" target="_blank" class="text-[#F08300] underline font-semibold">$1</a>');
  }
</script>

<div class="border border-gray-300 rounded-2xl overflow-hidden bg-white shadow-xs">
  <!-- Top Bar with Editor / Preview Toggle & Word Count -->
  <div class="bg-gray-50 border-b border-gray-200 p-2.5 flex flex-wrap justify-between items-center gap-2">
    <!-- View Switcher -->
    <div class="flex bg-gray-200/70 p-1 rounded-xl text-xs font-bold uppercase">
      <button 
        type="button"
        class={`px-3.5 py-1.5 rounded-lg transition ${activeTab === 'write' ? 'bg-white text-[#101116] shadow-xs' : 'text-gray-500 hover:text-[#101116]'}`}
        on:click={() => activeTab = 'write'}
      >
        Tulis (Editor)
      </button>
      <button 
        type="button"
        class={`px-3.5 py-1.5 rounded-lg transition ${activeTab === 'preview' ? 'bg-white text-[#101116] shadow-xs' : 'text-gray-500 hover:text-[#101116]'}`}
        on:click={() => activeTab = 'preview'}
      >
        Pratinjau Pembaca
      </button>
    </div>

    <!-- Metrics -->
    <div class="text-[11px] text-[#6f7775] font-mono flex items-center gap-3 pr-2">
      <span>{wordCount} kata</span>
      <span>&bull;</span>
      <span>~{readingTime} menit baca</span>
    </div>
  </div>

  <!-- Adaptive Formatting Toolbar (Only in Write Mode) -->
  {#if activeTab === 'write'}
    <div class="bg-white border-b border-gray-200 px-3 py-2 flex flex-wrap items-center gap-1.5 text-xs">
      <!-- Inline formatting -->
      <button type="button" on:click={() => wrapSelection('**', '**')} class="px-2.5 py-1 font-bold bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Tebal (Ctrl+B)">B</button>
      <button type="button" on:click={() => wrapSelection('*', '*')} class="px-2.5 py-1 italic bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Miring (Ctrl+I)">I</button>
      <button type="button" on:click={() => wrapSelection('<u>', '</u>')} class="px-2.5 py-1 underline bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Garis Bawah (Ctrl+U)">U</button>

      <span class="w-px h-5 bg-gray-200 mx-1"></span>

      <!-- Headings -->
      <button type="button" on:click={() => insertBlock('## ')} class="px-2.5 py-1 font-bold text-[11px] bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Subjudul H2">H2</button>
      <button type="button" on:click={() => insertBlock('### ')} class="px-2.5 py-1 font-bold text-[11px] bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Subjudul H3">H3</button>

      <span class="w-px h-5 bg-gray-200 mx-1"></span>

      <!-- Lists with adaptive enter feature hint -->
      <button type="button" on:click={() => insertBlock('1. ')} class="px-2.5 py-1 font-mono text-[11px] bg-amber-50 text-amber-900 border border-amber-200 hover:bg-amber-100 rounded-lg" title="Daftar Nomor Otomatis (Tekan Enter berlanjut otomatis)">1. Angka</button>
      <button type="button" on:click={() => insertBlock('- ')} class="px-2.5 py-1 font-mono text-[11px] bg-amber-50 text-amber-900 border border-amber-200 hover:bg-amber-100 rounded-lg" title="Daftar Poin Otomatis (Tekan Enter berlanjut otomatis)">&bull; Poin</button>

      <span class="w-px h-5 bg-gray-200 mx-1"></span>

      <!-- Blockquote & Callout -->
      <button type="button" on:click={() => insertBlock('> ')} class="px-2.5 py-1 text-[11px] bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Kutipan Pernyataan">Kutipan</button>
      <button type="button" on:click={() => insertBlock('[Catatan Redaksi] ')} class="px-2.5 py-1 text-[11px] bg-gray-100 hover:bg-gray-200 rounded-lg text-gray-800" title="Kotak Catatan Redaksi">Catatan</button>

      <span class="w-px h-5 bg-gray-200 mx-1"></span>

      <!-- Helper pill -->
      <span class="text-[10px] text-gray-400 italic hidden sm:inline ml-auto">
        Tips: Tekan Enter pada daftar nomor/poin akan otomatis menambahkan angka berikutnya
      </span>
    </div>
  {/if}

  <!-- Editor Textarea Area -->
  {#if activeTab === 'write'}
    <div class="p-4">
      <textarea 
        bind:this={textarea}
        bind:value={content}
        on:keydown={handleKeyDown}
        rows="12"
        {placeholder}
        class="w-full text-xs font-mono leading-relaxed text-[#101116] focus:outline-none resize-y bg-transparent"
      ></textarea>
    </div>
  {:else}
    <!-- Live Reader Preview Area -->
    <div class="p-6 bg-[#fcfbf7] min-h-[300px] prose max-w-none text-xs">
      {@html renderFormattedHtml(content)}
    </div>
  {/if}
</div>
