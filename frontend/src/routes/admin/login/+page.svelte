<script lang="ts">
  import { goto } from '$app/navigation';
  import { loginAdmin } from '$lib/api/client';
  import { saveSession } from '$lib/utils/auth';
  import { 
    ShieldCheck, 
    Lock, 
    User, 
    Eye, 
    EyeOff, 
    ArrowRight, 
    ArrowLeft, 
    KeyRound
  } from 'lucide-svelte';

  let username = '';
  let password = '';
  let errorMessage = '';
  let isLoading = false;
  let showPassword = false;

  // iOS Segmented Tab: 'admin' | 'superadmin' (2 Peran)
  let activeTab: 'admin' | 'superadmin' = 'admin';

  const presets = {
    admin: {
      label: 'Admin Satpol PP',
      username: 'stafftrantibum',
      pass: 'trantibum2026!',
      roleLabel: 'Admin Operasional & Penegak Perda',
      hint: 'Akses penanganan aduan SIAP TERTIB, validasi lapangan, BAP tiket, & CMS publik.'
    },
    superadmin: {
      label: 'Superadmin Diskominfo-SP',
      username: 'superadmin',
      pass: 'sulselprov2026!',
      roleLabel: 'Pengelola Sistem Komando Pemprov',
      hint: 'Manajemen akun pengguna, hak akses RBAC, log audit keamanan, & kendali portal.'
    }
  };

  function selectTab(tab: 'admin' | 'superadmin') {
    activeTab = tab;
    username = presets[tab].username;
    password = presets[tab].pass;
    errorMessage = '';
  }

  async function handleLogin() {
    errorMessage = '';
    if (!username.trim() || !password) {
      errorMessage = 'Nama pengguna dan kata sandi wajib diisi.';
      return;
    }
    isLoading = true;
    try {
      const res = await loginAdmin(username, password);
      if (res.user) {
        saveSession({
          token: res.token || `jwt-mock-${res.user.role}-${res.user.id}`,
          user: res.user
        });
        if (res.user.role === 'SUPERADMIN_DISKOMINFO') {
          goto('/superadmin/dashboard');
        } else {
          goto('/admin/dashboard');
        }
      } else {
        errorMessage = res.error || 'Nama pengguna atau kata sandi tidak sesuai.';
      }
    } catch (e: any) {
      errorMessage = 'Koneksi ke server gagal: ' + e.message;
    } finally {
      isLoading = false;
    }
  }

  // Set default initial credentials matching activeTab
  username = presets.admin.username;
  password = presets.admin.pass;
</script>

<svelte:head>
  <title>Portal Masuk Komando Petugas | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Full-Bleed Futuristic iOS-Style Split Viewport (Non-scrollable on Windows) -->
<div class="h-screen max-h-screen w-full flex flex-col lg:flex-row bg-[#f8fafc] text-slate-800 font-sans selection:bg-[#F08300] selection:text-white overflow-hidden relative">

  <!-- =========================================================================
       LEFT PANEL: AUTHENTICATION CONSOLE (EXPANDED ~48% COMFORTABLE RATIO)
       ========================================================================= -->
  <div class="w-full lg:w-[50%] xl:w-[48%] 2xl:w-[46%] lg:max-w-[680px] shrink-0 h-full max-h-screen bg-white border-r border-slate-200/70 flex flex-col justify-between p-6 sm:p-10 lg:p-12 xl:p-14 relative z-20 shadow-xl transition-all overflow-y-auto lg:overflow-hidden">
    
    <div class="w-full max-w-[460px] mx-auto flex flex-col justify-between h-full">
      <!-- Top Bar: Logo & Navigation Back -->
      <div>
        <div class="flex items-center justify-between gap-4 pb-4 border-b border-slate-100">
          <a href="/" class="flex items-center gap-3 group" title="Kembali ke Beranda Publik">
            <img
              src="/uploads/foto_website/1686928245logo_black.png"
              alt="Logo Satpol PP Sulsel"
              class="h-10 sm:h-11 w-auto object-contain transition group-hover:scale-105"
            />
          </a>

          <a
            href="/"
            class="inline-flex items-center gap-1.5 text-xs font-semibold text-slate-600 hover:text-[#246bfd] bg-slate-100/80 hover:bg-slate-200/80 px-3.5 py-1.5 rounded-xl border border-slate-200/60 transition ds-press"
          >
            <ArrowLeft class="w-3.5 h-3.5" />
            <span>Beranda</span>
          </a>
        </div>

        <!-- Header Greeting & Status (MAGANGKI' Clean Typography Style) -->
        <div class="pt-5 space-y-1.5">
          <div class="text-[11px] font-bold uppercase tracking-wider text-slate-400">
            Selamat Datang Kembali
          </div>
          
          <h1 class="text-2xl sm:text-3xl font-black text-slate-900 tracking-tight leading-tight">
            Masuk. Satukan Wibawa.
          </h1>
          
          <p class="text-xs text-slate-500 leading-relaxed">
            Pilih akun Anda untuk membuka ruang tugas, pengawasan trantibum, dan tindak lanjut aduan warga.
          </p>
        </div>

        <!-- Segmented Control (Clean MAGANGKI' Pill Container - 2 Peran) -->
        <div class="mt-4 p-1.5 bg-slate-100/90 rounded-2xl border border-slate-200/60 grid grid-cols-2 gap-1.5 shadow-inner relative">
          <button
            type="button"
            on:click={() => selectTab('admin')}
            class="relative z-10 py-2.5 text-center text-xs font-bold rounded-xl transition-all duration-200 ds-press {activeTab === 'admin' ? 'bg-white text-slate-900 shadow-sm font-extrabold' : 'text-slate-500 hover:text-slate-800 font-medium'}"
          >
            Admin Petugas
          </button>
          <button
            type="button"
            on:click={() => selectTab('superadmin')}
            class="relative z-10 py-2.5 text-center text-xs font-bold rounded-xl transition-all duration-200 ds-press {activeTab === 'superadmin' ? 'bg-white text-slate-900 shadow-sm font-extrabold' : 'text-slate-500 hover:text-slate-800 font-medium'}"
          >
            Superadmin
          </button>
        </div>

        <!-- Active Role Context Card -->
        <div class="mt-3 px-3.5 py-2.5 rounded-xl bg-blue-50/70 border border-blue-200/50 flex items-center justify-between text-xs text-slate-700 animate-in fade-in duration-150">
          <div class="flex items-center gap-2.5">
            <div class="p-1.5 rounded-lg bg-blue-100/80 text-[#246bfd]">
              <KeyRound class="w-3.5 h-3.5" />
            </div>
            <div>
              <span class="text-[10px] font-bold uppercase tracking-wider text-blue-700 block">{presets[activeTab].roleLabel}</span>
              <span class="text-[11px] text-slate-600 line-clamp-1">{presets[activeTab].hint}</span>
            </div>
          </div>
        </div>

        <!-- Error Alert -->
        {#if errorMessage}
          <div role="alert" class="mt-3 p-3 rounded-xl bg-red-500/10 border border-red-500/30 text-red-700 text-xs font-semibold flex items-start gap-2 animate-in shake duration-200">
            <span class="text-red-500 font-bold">&bull;</span>
            <span>{errorMessage}</span>
          </div>
        {/if}

        <!-- Main Login Form -->
        <form on:submit|preventDefault={handleLogin} class="mt-4 space-y-3.5">
          <!-- Username Input -->
          <div>
            <label for="admin-username" class="block text-xs font-bold text-slate-700 mb-1.5">
              Nama Pengguna (Username)
            </label>
            <div class="relative flex items-center">
              <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                <User class="w-4 h-4" />
              </div>
              <input
                id="admin-username"
                type="text"
                bind:value={username}
                required
                autocomplete="username"
                placeholder="Username akun petugas"
                class="w-full pl-10 pr-4 py-2.5 text-xs sm:text-sm text-slate-800 placeholder-slate-400 bg-slate-50/60 border border-slate-200 rounded-xl focus:outline-none focus:border-[#246bfd] focus:bg-white focus:ring-2 focus:ring-[#246bfd]/10 transition shadow-2xs"
              />
            </div>
          </div>

          <!-- Password Input with Toggle Eye -->
          <div>
            <div class="flex items-center justify-between mb-1.5">
              <label for="admin-password" class="block text-xs font-bold text-slate-700">
                Kata Sandi (Password)
              </label>
              <span class="text-[10px] text-slate-400">Enkripsi SHA-256</span>
            </div>
            <div class="relative flex items-center">
              <div class="absolute inset-y-0 left-0 pl-3.5 flex items-center pointer-events-none text-slate-400">
                <Lock class="w-4 h-4" />
              </div>
              {#if showPassword}
                <input
                  id="admin-password"
                  type="text"
                  bind:value={password}
                  required
                  autocomplete="current-password"
                  placeholder="••••••••"
                  class="w-full pl-10 pr-11 py-2.5 text-xs sm:text-sm text-slate-800 placeholder-slate-400 bg-slate-50/60 border border-slate-200 rounded-xl focus:outline-none focus:border-[#246bfd] focus:bg-white focus:ring-2 focus:ring-[#246bfd]/10 transition shadow-2xs font-mono"
                />
              {:else}
                <input
                  id="admin-password"
                  type="password"
                  bind:value={password}
                  required
                  autocomplete="current-password"
                  placeholder="••••••••"
                  class="w-full pl-10 pr-11 py-2.5 text-xs sm:text-sm text-slate-800 placeholder-slate-400 bg-slate-50/60 border border-slate-200 rounded-xl focus:outline-none focus:border-[#246bfd] focus:bg-white focus:ring-2 focus:ring-[#246bfd]/10 transition shadow-2xs font-mono"
                />
              {/if}
              <button
                type="button"
                on:click={() => showPassword = !showPassword}
                class="absolute inset-y-0 right-0 pr-3.5 flex items-center text-slate-400 hover:text-slate-600 transition focus:outline-none"
                aria-label={showPassword ? 'Sembunyikan password' : 'Tampilkan password'}
              >
                {#if showPassword}
                  <EyeOff class="w-4 h-4" />
                {:else}
                  <Eye class="w-4 h-4" />
                {/if}
              </button>
            </div>
          </div>

          <!-- Submit Button (Soft Vibrant Modern Blue) -->
          <button
            type="submit"
            disabled={isLoading}
            class="w-full mt-2 bg-gradient-to-r from-[#246bfd] to-[#1d58d8] hover:from-[#1d58d8] hover:to-[#1746b0] ds-hover-lift ds-press text-white font-bold text-xs uppercase tracking-wider py-3.5 rounded-xl transition shadow-lg shadow-blue-500/25 disabled:opacity-50 flex items-center justify-center gap-2 group cursor-pointer"
          >
            {#if isLoading}
              <svg class="animate-spin h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
              </svg>
              <span>Mengautentikasi...</span>
            {:else}
              <span>Masuk ke Workspace Komando</span>
              <ArrowRight class="w-4 h-4 transform group-hover:translate-x-1 transition-transform" />
            {/if}
          </button>
        </form>
      </div>

      <!-- Bottom Footnote & Security Badging -->
      <div class="pt-4 border-t border-slate-100 flex flex-col sm:flex-row items-center justify-between gap-2 text-[11px] text-slate-400">
        <div class="flex items-center gap-1.5">
          <ShieldCheck class="w-4 h-4 text-emerald-600" />
          <span>Jaringan Intra Pemerintah Prov. Sulsel</span>
        </div>
        <span class="font-mono text-slate-400">v2.4 Sovereign OS</span>
      </div>
    </div>
  </div>

  <!-- =========================================================================
       RIGHT PANEL: SOFT SKY-AZURE GRADIENT WITH HARMONIOUS AMBIENT LIGHTING
       (Soft, calming atmosphere replacing dark blue, zero distracting text/cards)
       ========================================================================= -->
  <div class="flex-1 hidden lg:block relative h-full overflow-hidden bg-gradient-to-br from-[#d9ebf5] via-[#e8f3f8] to-[#cbe5f2]">
    <!-- Background Architectural Banner Image with Soft Natural Blend -->
    <img
      src="/uploads/foto_banner/1747198131welcome.png"
      alt="Latar Belakang Markas Satpol PP Sulsel"
      class="w-full h-full object-cover object-center filter brightness-[1.03] contrast-[0.98] opacity-80 mix-blend-multiply"
    />

    <!-- Gentle Soft Pastel Gradient Overlays -->
    <div class="absolute inset-0 bg-gradient-to-r from-white/95 via-[#e2f1f8]/70 to-[#cde7f4]/40"></div>
    <div class="absolute inset-0 bg-gradient-to-t from-[#c5e3f2]/90 via-transparent to-white/60"></div>

    <!-- Soft Cloud-like Light Glow Orbs -->
    <div class="absolute -top-24 -right-24 w-[550px] h-[550px] rounded-full bg-white/70 blur-[130px] pointer-events-none"></div>
    <div class="absolute bottom-10 right-10 w-[450px] h-[450px] rounded-full bg-[#bde0f5]/50 blur-[110px] pointer-events-none"></div>
  </div>

</div>
