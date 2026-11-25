<script lang="ts">
  import { createEventDispatcher } from 'svelte';

  export let id: string = `input-${Math.random().toString(36).slice(2, 9)}`;
  export let name: string | undefined = undefined;
  export let label: string | undefined = undefined;
  export let hint: string | undefined = undefined;
  export let error: string | undefined = undefined;
  export let required: boolean = false;
  export let type: 'text' | 'email' | 'password' | 'number' | 'tel' | 'url' | 'search' | 'textarea' = 'text';
  export let value: string | number = '';
  export let placeholder: string | undefined = undefined;
  export let disabled: boolean = false;
  export let readonly: boolean = false;
  export let rows: number = 4;
  export let maxlength: number | undefined = undefined;
  export let pattern: string | undefined = undefined;
  export let autocomplete: string | undefined = undefined;
  export let className: string = '';

  const dispatch = createEventDispatcher<{
    input: Event;
    change: Event;
    focus: FocusEvent;
    blur: FocusEvent;
  }>();

  let showPassword = false;

  $: effectiveType = type === 'password' ? (showPassword ? 'text' : 'password') : type;
  $: errorId = error ? `${id}-error` : undefined;
  $: hintId = hint ? `${id}-hint` : undefined;
  $: describedBy = [errorId, hintId].filter(Boolean).join(' ') || undefined;

  function handleInput(e: Event) {
    const target = e.target as HTMLInputElement | HTMLTextAreaElement;
    value = target.value;
    dispatch('input', e);
  }

  function handleChange(e: Event) {
    dispatch('change', e);
  }
</script>

<div class="w-full flex flex-col gap-1.5 {className}">
  {#if label}
    <div class="flex items-center justify-between">
      <label for={id} class="text-xs font-semibold text-slate-700 tracking-wide select-none">
        {label}
        {#if required}
          <span class="text-rose-500 font-bold ml-0.5" aria-hidden="true">*</span>
        {/if}
      </label>
      <slot name="labelExtra" />
    </div>
  {/if}

  <div class="relative flex items-center w-full">
    {#if $$slots.iconLeading}
      <div class="absolute left-3.5 flex items-center pointer-events-none text-slate-400">
        <slot name="iconLeading" />
      </div>
    {/if}

    {#if type === 'textarea'}
      <textarea
        {id}
        {name}
        {rows}
        {placeholder}
        {disabled}
        {readonly}
        {required}
        {maxlength}
        aria-invalid={error ? 'true' : undefined}
        aria-describedby={describedBy}
        class="w-full px-3.5 py-2.5 text-sm bg-white/80 border rounded-2xl transition-all duration-150 backdrop-blur-md placeholder:text-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500/25 focus:border-blue-500 disabled:opacity-50 disabled:bg-slate-50 disabled:cursor-not-allowed resize-y {error
          ? 'border-rose-400 bg-rose-50/40 text-rose-900 focus:border-rose-500 focus:ring-rose-500/20'
          : 'border-slate-200/90 hover:border-slate-300'} {$$slots.iconLeading ? 'pl-10' : ''}"
        bind:value
        on:input={handleInput}
        on:change={handleChange}
        on:focus={(e) => dispatch('focus', e)}
        on:blur={(e) => dispatch('blur', e)}
        {...$$restProps}
      ></textarea>
    {:else}
      <input
        {id}
        {name}
        type={effectiveType}
        {placeholder}
        {disabled}
        {readonly}
        {required}
        {maxlength}
        {pattern}
        {autocomplete}
        aria-invalid={error ? 'true' : undefined}
        aria-describedby={describedBy}
        class="w-full px-3.5 py-2.5 text-sm bg-white/80 border rounded-2xl transition-all duration-150 backdrop-blur-md placeholder:text-slate-400 text-slate-900 focus:outline-none focus:ring-2 focus:ring-blue-500/25 focus:border-blue-500 disabled:opacity-50 disabled:bg-slate-50 disabled:cursor-not-allowed {error
          ? 'border-rose-400 bg-rose-50/40 text-rose-900 focus:border-rose-500 focus:ring-rose-500/20'
          : 'border-slate-200/90 hover:border-slate-300'} {$$slots.iconLeading ? 'pl-10' : ''} {type === 'password' || $$slots.iconTrailing ? 'pr-10' : ''}"
        value={value}
        on:input={handleInput}
        on:change={handleChange}
        on:focus={(e) => dispatch('focus', e)}
        on:blur={(e) => dispatch('blur', e)}
        {...$$restProps}
      />
    {/if}

    {#if type === 'password'}
      <button
        type="button"
        on:click={() => (showPassword = !showPassword)}
        aria-label={showPassword ? 'Sembunyikan password' : 'Tampilkan password'}
        class="absolute right-3 p-1 rounded-lg text-slate-400 hover:text-slate-700 transition"
      >
        {#if showPassword}
          <!-- Eye Off -->
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.875 18.825A10.05 10.05 0 0112 19c-4.478 0-8.268-2.943-9.543-7a9.97 9.97 0 011.563-3.029m5.858.908a3 3 0 114.243 4.243M9.878 9.878l4.242 4.242M9.88 9.88l-3.29-3.29m7.532 7.532l3.29 3.29M3 3l18 18" />
          </svg>
        {:else}
          <!-- Eye On -->
          <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z" />
          </svg>
        {/if}
      </button>
    {:else if $$slots.iconTrailing}
      <div class="absolute right-3.5 flex items-center text-slate-400">
        <slot name="iconTrailing" />
      </div>
    {/if}
  </div>

  {#if error}
    <p id={errorId} class="text-xs text-rose-600 font-medium flex items-center gap-1 mt-0.5" role="alert">
      <svg class="w-3.5 h-3.5 shrink-0" fill="currentColor" viewBox="0 0 20 20" aria-hidden="true">
        <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7 4a1 1 0 11-2 0 1 1 0 012 0zm-1-9a1 1 0 00-1 1v4a1 1 0 102 0V6a1 1 0 00-1-1z" clip-rule="evenodd" />
      </svg>
      {error}
    </p>
  {:else if hint}
    <p id={hintId} class="text-[11px] text-slate-500 mt-0.5">
      {hint}
    </p>
  {/if}
</div>
