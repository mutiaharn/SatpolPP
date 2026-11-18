<script lang="ts">
  import { goto } from '$app/navigation';
  import { loginAdmin } from '$lib/api/client';
  import { saveSession } from '$lib/utils/auth';
  import { 
    Shield, 
    ShieldCheck, 
    Lock, 
    User, 
    Eye, 
    EyeOff, 
    ArrowRight, 
    ArrowLeft, 
    Sparkles, 
    CheckCircle2, 
    Compass, 
    Radio, 
    Building2,
    KeyRound
  } from 'lucide-svelte';

  let username = '';
  let password = '';
  let errorMessage = '';
  let isLoading = false;
  let showPassword = false;

  // iOS Segmented Tab: 'staff' | 'kasatpol' | 'superadmin'
  let activeTab: 'staff' | 'kasatpol' | 'superadmin' = 'staff';

  const presets = {
    staff: {
      label: 'Staff Lapangan & Trantibum',
      username: 'stafftrantibum',
      pass: 'trantibum2026!',
      roleLabel: 'Aparatur Penegak Perda',
      hint: 'Akses penanganan aduan, validasi bukti lapangan, & update BAP tiket.'
    },
    kasatpol: {
      label: 'Pimpinan & Kasatpol',
      username: 'kasatpol',
      pass: 'satpolpp2026!',
      roleLabel: 'Komando Pimpinan Instansi',
      hint: 'Disposisi naskah, analitik 24 kab/kota, & persetujuan tindak lanjut.'
    },
    superadmin: {
      label: 'Superadmin Diskominfo-SP',
      username: 'superadmin',
      pass: 'sulselprov2026!',
      roleLabel: 'Pengelola Sistem Pemprov',
      hint: 'Manajemen pengguna, audit log keamanan SHA-256, & kendali struktur.'
    }
  };

  function selectTab(tab: 'staff' | 'kasatpol' | 'superadmin') {
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
  username = presets.staff.username;
  password = presets.staff.pass;
</script>

<svelte:head>
  <title>Portal Masuk Komando Petugas | Satpol PP Prov. Sulsel</title>
</svelte:head>

<!-- Full-Bleed Futuristic iOS-Style Split Viewport -->
<div class="min-h-screen w-full flex flex-col lg:flex-row bg-[#060d19] text-slate-800 font-sans selection:bg-[#F08300] selection:text-white overflow-hidden relative">

  <!-- =========================================================================
       LEFT PANEL: iOS-STYLE FROSTED GLASS AUTHENTICATION CONSOLE
       ========================================================================= -->
  <div class="w-full lg:w-[490px] xl:w-[540px] shrink-0 min-h-screen ds-glass-strong bg-white/90 backdrop-blur-3xl border-r border-white/60 flex flex-col justify-between p-6 sm:p-10 lg:p-12 relative z-20 shadow-2xl transition-all">
    
    <!-- Top Bar: Logo & Navigation Back -->
    <div>
      <div class="flex items-center justify-between gap-4 pb-6 sm:pb-8 border-b border-slate-200/50">
        <a href="/" class="flex items-center gap-3 group" title="Kembali ke Beranda Publik">
          <img
            src="/uploads/foto_website/1686928245logo_black.png"
            alt="Logo Satpol PP Sulsel"
            class="h-10 sm:h-11 w-auto object-contain transition group-hover:scale-105"
          />
        </a>

        <a
          href="/"
          class="inline-flex items-center gap-1.5 text-xs font-semibold text-slate-600 hover:text-[#F08300] ds-glass px-3.5 py-1.5 rounded-xl border border-white/80 hover:border-[#F08300]/40 transition ds-press"
        >
          <ArrowLeft class="w-3.5 h-3.5" />
          <span>Beranda</span>
        </a>
      </div>

      <!-- Header Greeting & Status -->
      <div class="pt-6 sm:pt-8 space-y-2">
        <div class="inline-flex items-center gap-2 px-3 py-1 rounded-full ds-glass-azure text-[#246bfd] font-bold text-[10px] uppercase tracking-widest border border-blue-200/60">
          <span>Sistem Komando Intra Sulsel</span>
        </div>
        
        <h1 class="text-2xl sm:text-3xl font-extrabold text-[#071936] tracking-tight leading-tight">
          Masuk. Satukan Wibawa.
        </h1>
        
        <p class="text-xs text-slate-500 leading-relaxed">
          Pilih profil akses Anda untuk membuka ruang tugas, pengawasan trantibum, dan tindak lanjut aduan warga.
        </p>
      </div>

      <!-- iOS Segmented Control (Smooth Sliding Tab Filter) -->
      <div class="mt-6 p-1 ds-glass rounded-2xl border border-white/80 grid grid-cols-3 gap-1 shadow-inner relative">
        <button
          type="button"
          on:click={() => selectTab('staff')}
          class="relative z-10 py-2.5 text-center text-xs font-bold rounded-xl transition-all duration-200 ds-press {activeTab === 'staff' ? 'ds-pill-active font-extrabold' : 'text-slate-600 hover:text-slate-900'}"
        >
          Staf Lapangan
        </button>
        <button
          type="button"
          on:click={() => selectTab('kasatpol')}
          class="relative z-10 py-2.5 text-center text-xs font-bold rounded-xl transition-all duration-200 ds-press {activeTab === 'kasatpol' ? 'ds-pill-active font-extrabold' : 'text-slate-600 hover:text-slate-900'}"
        >
          Pimpinan
        </button>
        <button
          type="button"
          on:click={() => selectTab('superadmin')}
          class="relative z-10 py-2.5 text-center text-xs font-bold rounded-xl transition-all duration-200 ds-press {activeTab === 'superadmin' ? 'ds-pill-active font-extrabold' : 'text-slate-600 hover:text-slate-900'}"
        >
          Superadmin
        </button>
      </div>

      <!-- Active Role Context Card (iOS Style) -->
      <div class="mt-3 px-3.5 py-2.5 rounded-xl ds-glass-azure border border-blue-200/50 flex items-center justify-between text-xs text-slate-700 animate-in fade-in duration-150">
        <div class="flex items-center gap-2">
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
        <div role="alert" class="mt-4 p-3.5 rounded-2xl bg-red-500/10 backdrop-blur-md border border-red-500/30 text-red-700 text-xs font-semibold flex items-start gap-2 animate-in shake duration-200">
          <span class="text-red-500 font-bold">&bull;</span>
          <span>{errorMessage}</span>
        </div>
      {/if}

      <!-- Main Login Form -->
      <form on:submit|preventDefault={handleLogin} class="mt-5 space-y-4">
        <!-- Username Input with iOS Squircle Style -->
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
              class="ds-glass-input w-full pl-10 pr-4 py-3 text-xs sm:text-sm text-slate-800 placeholder-slate-400 focus:outline-none transition shadow-2xs"
            />
          </div>
        </div>

        <!-- Password Input with Toggle Eye -->
        <div>
          <div class="flex items-center justify-between mb-1.5">
            <label for="admin-password" class="block text-xs font-bold text-slate-700">
              Kata Sandi (Password)
            </label>
            <span class="text-[11px] text-slate-400">Enkripsi SHA-256</span>
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
                class="ds-glass-input w-full pl-10 pr-11 py-3 text-xs sm:text-sm text-slate-800 placeholder-slate-400 focus:outline-none transition shadow-2xs font-mono"
              />
            {:else}
              <input
                id="admin-password"
                type="password"
                bind:value={password}
                required
                autocomplete="current-password"
                placeholder="••••••••"
                class="ds-glass-input w-full pl-10 pr-11 py-3 text-xs sm:text-sm text-slate-800 placeholder-slate-400 focus:outline-none transition shadow-2xs font-mono"
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

        <!-- Submit Button (iOS Fluid Physics) -->
        <button
          type="submit"
          disabled={isLoading}
          class="w-full mt-2 bg-gradient-to-r from-[#071936] to-[#0f2d5e] hover:from-[#0b244d] hover:to-[#173d7a] ds-hover-lift ds-press text-white font-bold text-xs uppercase tracking-wider py-3.5 rounded-2xl transition shadow-lg shadow-blue-950/20 disabled:opacity-50 flex items-center justify-center gap-2 group cursor-pointer ds-specular-top"
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
    <div class="pt-6 border-t border-slate-100 flex flex-col sm:flex-row items-center justify-between gap-3 text-[11px] text-slate-400">
      <div class="flex items-center gap-1.5">
        <ShieldCheck class="w-4 h-4 text-emerald-600" />
        <span>Jaringan Intra Pemerintah Prov. Sulsel</span>
      </div>
      <span class="font-mono text-slate-400">v2.4 Sovereign OS</span>
    </div>
  </div>

  <!-- =========================================================================
       RIGHT PANEL: FUTURISTIC SOVEREIGN SHOWCASE & FLOATING iOS GLASS CARDS
       ========================================================================= -->
  <div class="flex-1 hidden lg:flex flex-col justify-between p-12 xl:p-16 relative overflow-hidden bg-gradient-to-br from-[#060e1d] via-[#0b172a] to-[#030710] text-white">
    
    <!-- Ambient Futuristic Glow Orbs -->
    <div class="absolute -top-32 -right-32 w-[550px] h-[550px] rounded-full bg-[#246bfd]/20 blur-[120px] pointer-events-none"></div>
    <div class="absolute bottom-10 left-10 w-[450px] h-[450px] rounded-full bg-[#F08300]/15 blur-[100px] pointer-events-none"></div>

    <!-- Futuristic Top Headline (Matches reference "Karya kecil hari ini...") -->
    <div class="relative z-10 max-w-xl space-y-4">
      <div class="inline-flex items-center gap-2 px-4 py-1.5 rounded-full bg-white/10 backdrop-blur-xl border border-white/15 text-[#F08300] font-bold text-xs uppercase tracking-widest shadow-lg">
        <Sparkles class="w-3.5 h-3.5 text-[#F08300]" />
        <span>Semboyan Luhur - Praja Wibawa</span>
      </div>

      <h2 class="text-3xl xl:text-5xl font-black tracking-tight leading-[1.15] text-white">
        Ketenteraman Hari Ini.<br/>
        <span class="text-transparent bg-clip-text bg-gradient-to-r from-blue-400 via-sky-300 to-amber-300">
          Wibawa Sulsel Esok Hari.
        </span>
      </h2>

      <p class="text-sm xl:text-base text-slate-300 font-light leading-relaxed max-w-lg">
        Satu ruang digital komando terintegrasi untuk menyelenggarakan penegakan Perda, ketertiban umum, dan perlindungan masyarakat di 24 Kabupaten/Kota.
      </p>
    </div>

    <!-- Central Showcase with Kantor Gubernur Landmark & Floating Glass Badges -->
    <div class="relative z-10 my-8 flex items-center justify-center">
      <!-- Background Architectural Silhouette with Ambient Lighting -->
      <div class="relative w-full max-w-lg h-72 xl:h-80 rounded-3xl overflow-hidden border border-white/15 shadow-2xl group">
        <!-- Banner Photo with Ambient Gradient Overlay -->
        <img
          src="/uploads/foto_banner/1747198131welcome.png"
          alt="Markas Komando Satpol PP Sulsel"
          class="w-full h-full object-cover object-center group-hover:scale-105 transition duration-700 opacity-75"
        />
        <div class="absolute inset-0 bg-gradient-to-t from-[#060e1d] via-[#060e1d]/40 to-transparent"></div>

        <!-- Floating Dynamic iOS Glass Badge (Pinisi Sulsel Inspired by User Reference) -->
        <div class="absolute bottom-4 left-4 right-4 ds-glass-dark p-4 rounded-2xl border border-white/20 backdrop-blur-xl shadow-2xl flex items-center gap-3.5 transform transition-all hover:scale-[1.02]">
          <!-- Pinisi Emblem Icon -->
          <div class="w-12 h-12 rounded-xl bg-gradient-to-br from-amber-500/25 to-blue-500/25 border border-white/20 flex items-center justify-center shrink-0 shadow-md">
            <Compass class="w-6 h-6 text-[#F08300] animate-spin" style="animation-duration: 25s;" />
          </div>

          <div class="min-w-0 flex-1">
            <div class="flex items-center justify-between gap-2">
              <span class="font-extrabold text-xs text-white tracking-wide">Filosofi Layar Pinisi</span>
              <span class="text-[10px] text-amber-300 bg-amber-500/20 px-2 py-0.5 rounded-full border border-amber-400/30">Sulsel Maju</span>
            </div>
            <p class="text-[11px] text-slate-300 leading-snug mt-0.5 line-clamp-2">
              Keteguhan menegakkan aturan berkeadilan, ketulusan mengayomi dan melindungi seluruh warga.
            </p>
          </div>
        </div>

        <!-- Top-Right Live Operational Radar Chip -->
        <div class="absolute top-4 right-4 ds-glass-dark px-3 py-1.5 rounded-full border border-white/20 flex items-center gap-2 text-[10px] font-mono text-emerald-300 shadow-lg">
          <span class="w-2 h-2 rounded-full bg-emerald-400 animate-ping"></span>
          <span>PATROLI 24/7 AKTIF</span>
        </div>
      </div>
    </div>

    <!-- Bottom Feature Pills (Smooth iOS Capsule Badges) -->
    <div class="relative z-10 flex flex-wrap items-center gap-3 pt-4 border-t border-white/10 text-xs">
      <div class="inline-flex items-center gap-2 px-4 py-2 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-md text-slate-200 hover:bg-white/10 transition">
        <Building2 class="w-4 h-4 text-[#246bfd]" />
        <span>Penegakan Perda &amp; Perkada</span>
      </div>

      <div class="inline-flex items-center gap-2 px-4 py-2 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-md text-slate-200 hover:bg-white/10 transition">
        <Radio class="w-4 h-4 text-[#F08300]" />
        <span>Monitoring 24 Kabupaten/Kota</span>
      </div>

      <div class="inline-flex items-center gap-2 px-4 py-2 rounded-2xl bg-white/5 border border-white/10 backdrop-blur-md text-slate-200 hover:bg-white/10 transition">
        <CheckCircle2 class="w-4 h-4 text-emerald-400" />
        <span>Pelaporan SIAP TERTIB</span>
      </div>
    </div>
  </div>

</div>
