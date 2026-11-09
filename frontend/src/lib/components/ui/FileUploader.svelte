<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let value: string = '';
  export let accept: string = 'image/png,image/jpeg,image/webp,image/jpg';
  export let maxSizeMB: number = 5;
  export let label: string = 'Unggah Gambar / Foto Thumbnail';
  export let description: string = 'Seret & lepas foto di sini, atau klik untuk memilih file (Maks. 5MB, format JPG/PNG/WEBP)';
  export let previewShape: 'banner' | 'square' = 'banner';
  export let required: boolean = false;
  export let fileName: string = '';

  const dispatch = createEventDispatcher<{
    change: { url: string; file: File; name: string };
    clear: void;
  }>();

  let fileInput: HTMLInputElement;
  let isDragging = false;
  let errorMessage = '';

  function handleDragOver(e: DragEvent) {
    e.preventDefault();
    isDragging = true;
  }

  function handleDragLeave(e: DragEvent) {
    e.preventDefault();
    isDragging = false;
  }

  function handleDrop(e: DragEvent) {
    e.preventDefault();
    isDragging = false;
    errorMessage = '';

    if (e.dataTransfer && e.dataTransfer.files && e.dataTransfer.files[0]) {
      processFile(e.dataTransfer.files[0]);
    }
  }

  function handleFileSelect(e: Event) {
    const input = e.target as HTMLInputElement;
    errorMessage = '';
    if (input.files && input.files[0]) {
      processFile(input.files[0]);
    }
  }

  function processFile(file: File) {
    const maxBytes = maxSizeMB * 1024 * 1024;
    if (file.size > maxBytes) {
      errorMessage = `Ukuran berkas (${(file.size / 1024 / 1024).toFixed(1)}MB) melebihi batas maksimum ${maxSizeMB}MB.`;
      return;
    }

    fileName = file.name;
    const reader = new FileReader();
    reader.onload = (event) => {
      if (event.target?.result) {
        value = event.target.result as string;
        dispatch('change', { url: value, file, name: file.name });
      }
    };
    reader.readAsDataURL(file);
  }

  function triggerBrowse() {
    fileInput?.click();
  }

  function handleRemove() {
    value = '';
    fileName = '';
    errorMessage = '';
    if (fileInput) fileInput.value = '';
    dispatch('clear');
  }
</script>

<div class="space-y-2">
  {#if label}
    <div class="flex items-center justify-between">
      <span class="text-xs font-bold text-[#1a120c] flex items-center gap-1">
        {label}
        {#if required}
          <span class="text-red-500">*</span>
        {/if}
      </span>
      {#if value}
        <button
          type="button"
          on:click={handleRemove}
          class="text-[11px] font-semibold text-red-600 hover:text-red-700 flex items-center gap-1 transition"
        >
          <svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
          </svg>
          Hapus Gambar
        </button>
      {/if}
    </div>
  {/if}

  <!-- Hidden native file input -->
  <input
    id="file-uploader-input"
    type="file"
    bind:this={fileInput}
    on:change={handleFileSelect}
    {accept}
    class="hidden"
  />

  {#if value}
    <!-- Image Preview Thumbnail State (Shadcn style) -->
    <div class="relative bg-[#faf7f2] border-2 border-[#e8dbc9] rounded-xl overflow-hidden group shadow-xs">
      <div class={`w-full overflow-hidden flex items-center justify-center bg-stone-900/5 ${previewShape === 'banner' ? 'h-52' : 'h-40'}`}>
        <img
          src={value}
          alt="Thumbnail Preview"
          class="w-full h-full object-cover group-hover:scale-[1.02] transition duration-300"
        />
      </div>

      <!-- Action Overlay Bar -->
      <div class="p-3 bg-white border-t border-[#eee9db] flex flex-wrap items-center justify-between gap-2 text-xs">
        <div class="flex items-center gap-2 truncate text-stone-600">
          <svg class="w-4 h-4 text-[#F08300] shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16l4.586-4.586a2 2 0 012.828 0L16 16m-2-2l1.586-1.586a2 2 0 012.828 0L20 14m-6-6h.01M6 20h12a2 2 0 002-2V6a2 2 0 00-2-2H6a2 2 0 00-2 2v12a2 2 0 002 2z" />
          </svg>
          <span class="truncate font-medium">{fileName || 'Foto terpilih'}</span>
        </div>

        <div class="flex items-center gap-1.5 shrink-0">
          <button
            type="button"
            on:click={triggerBrowse}
            class="px-2.5 py-1 text-[11px] font-semibold bg-stone-100 hover:bg-stone-200 text-stone-700 rounded-md transition"
          >
            Ganti Foto
          </button>
          <button
            type="button"
            on:click={handleRemove}
            class="p-1 text-stone-400 hover:text-red-600 hover:bg-red-50 rounded-md transition"
            title="Hapus"
            aria-label="Hapus foto"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  {:else}
    <!-- Dropzone state button -->
    <button
      type="button"
      on:dragover={handleDragOver}
      on:dragleave={handleDragLeave}
      on:drop={handleDrop}
      on:click={triggerBrowse}
      class={`w-full border-2 border-dashed rounded-xl p-6 sm:p-8 flex flex-col items-center justify-center text-center cursor-pointer transition-all duration-200 select-none ${
        isDragging
          ? 'border-[#F08300] bg-[#f9f4e8] scale-[1.01] shadow-inner'
          : 'border-stone-300 hover:border-[#F08300] bg-stone-50/60 hover:bg-[#faf7f2]/50'
      }`}
    >
      <div class={`w-12 h-12 rounded-full flex items-center justify-center mb-3 transition ${
        isDragging ? 'bg-[#F08300] text-white scale-110' : 'bg-stone-100 text-[#8c6b45]'
      }`}>
        <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12" />
        </svg>
      </div>

      <div class="text-xs font-semibold text-[#1a120c] mb-1">
        <span class="text-[#F08300] underline">Klik untuk mengunggah</span> atau drag & drop file ke sini
      </div>
      <p class="text-[11px] text-[#6f7775] max-w-sm">
        {description}
      </p>
    </button>
  {/if}

  {#if errorMessage}
    <div class="text-[11px] text-red-600 font-semibold flex items-center gap-1.5 mt-1">
      <svg class="w-3.5 h-3.5 shrink-0" fill="none" stroke="currentColor" viewBox="0 0 24 24">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
      </svg>
      <span>{errorMessage}</span>
    </div>
  {/if}
</div>
