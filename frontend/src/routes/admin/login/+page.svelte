<script lang="ts">
  import { goto } from '$app/navigation';
  import { loginAdmin } from '$lib/api/client';
  import { saveSession } from '$lib/utils/auth';
  import { Button, Input, SegmentedControl, GlassCard } from '$lib/components/ui';
  import { 
    ShieldCheck, 
    Lock, 
    User, 
    ArrowRight, 
    ArrowLeft, 
    KeyRound
  } from 'lucide-svelte';

  let username = '';
  let password = '';
  let errorMessage = '';
  let isLoading = false;

  // iOS Segmented Tab: 'admin' | 'superadmin' (2 Peran)
  let activeTab: 'admin' | 'superadmin' = 'admin';

  const roleOptions = [
    { value: 'admin', label: 'Admin Petugas' },
    { value: 'superadmin', label: 'Superadmin' }
  ];

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

  function handleTabChange(tab: string) {
    activeTab = tab as 'admin' | 'superadmin';
    username = presets[activeTab].username;
    password = presets[activeTab].pass;
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

          <Button
            href="/"
            variant="glass"
            size="sm"
            className="text-slate-600 hover:text-[#F08300]"
          >
            <ArrowLeft class="w-3.5 h-3.5" slot="iconLeading" />
            <span>Beranda</span>
          </Button>
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

        <!-- Segmented Control (Modular Apple Pill Control - 2 Peran) -->
        <div class="mt-4">
          <SegmentedControl
            options={roleOptions}
            value={activeTab}
            size="md"
            on:change={(e) => handleTabChange(e.detail)}
          />
        </div>

        <!-- Active Role Context Card -->
        <GlassCard variant="subtle" padding="sm" rounded="xl" className="mt-3 bg-amber-500/10 border-amber-500/20">
          <div class="flex items-center justify-between text-xs text-slate-700 animate-in fade-in duration-150">
            <div class="flex items-center gap-2.5">
              <div class="p-1.5 rounded-lg bg-amber-500/15 text-[#F08300] shrink-0">
                <KeyRound class="w-3.5 h-3.5" />
              </div>
              <div class="min-w-0">
                <span class="text-[10px] font-bold uppercase tracking-wider text-[#d97400] block truncate">{presets[activeTab].roleLabel}</span>
                <span class="text-[11px] text-slate-600 line-clamp-1">{presets[activeTab].hint}</span>
              </div>
            </div>
          </div>
        </GlassCard>

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
          <Input
            id="admin-username"
            name="username"
            label="Nama Pengguna (Username)"
            type="text"
            bind:value={username}
            required
            autocomplete="username"
            placeholder="Username akun petugas"
          >
            <User class="w-4 h-4" slot="iconLeading" />
          </Input>

          <!-- Password Input -->
          <Input
            id="admin-password"
            name="password"
            label="Kata Sandi (Password)"
            type="password"
            bind:value={password}
            required
            autocomplete="current-password"
            placeholder="••••••••"
          >
            <Lock class="w-4 h-4" slot="iconLeading" />
            <span slot="labelExtra" class="text-[10px] text-slate-400 font-normal">Enkripsi SHA-256</span>
          </Input>

          <!-- Submit Button (Modular Button with Apple Specular & Gradient) -->
          <Button
            type="submit"
            variant="primary"
            size="lg"
            fullWidth
            loading={isLoading}
            className="mt-2 text-xs uppercase tracking-wider font-bold py-3.5"
          >
            <span>{isLoading ? 'Mengautentikasi...' : 'Masuk ke Workspace Komando'}</span>
            <ArrowRight class="w-4 h-4 transform group-hover:translate-x-1 transition-transform" slot="iconTrailing" />
          </Button>
        </form>
      </div>

      <!-- Bottom Footnote & Security Badging -->
      <div class="pt-4 border-t border-slate-100 flex flex-col sm:flex-row items-center justify-between gap-2 text-[11px] text-slate-400">
        <div class="flex items-center gap-1.5">
          <ShieldCheck class="w-4 h-4 text-[#F08300]" />
          <span>Jaringan Intra Pemerintah Prov. Sulsel</span>
        </div>
        <span class="font-mono text-slate-400">v2.4 Sovereign OS</span>
      </div>
    </div>
  </div>

  <!-- =========================================================================
       RIGHT PANEL: SERENE AMBIENT ATMOSPHERE WITH OBSIDIAN/AMBER LIGHTING
       (Soft, calming atmosphere, zero distracting text/cards)
       ========================================================================= -->
  <div class="flex-1 hidden lg:block relative h-full overflow-hidden bg-slate-900">
    <!-- Background Architectural Banner Image with Soft Natural Blend -->
    <img
      src="/uploads/foto_banner/1747198131welcome.png"
      alt="Latar Belakang Markas Satpol PP Sulsel"
      class="w-full h-full object-cover object-center filter brightness-[0.75] contrast-[1.05] opacity-65"
    />

    <!-- Gentle Soft Dark & Warm Gradient Overlays -->
    <div class="absolute inset-0 bg-gradient-to-r from-slate-900/90 via-slate-900/50 to-transparent"></div>
    <div class="absolute inset-0 bg-gradient-to-t from-slate-900/80 via-transparent to-slate-900/40"></div>

    <!-- Soft Ambient Light Glow Orbs -->
    <div class="absolute -top-24 -right-24 w-[500px] h-[500px] rounded-full bg-[#F08300]/10 blur-[120px] pointer-events-none"></div>
    <div class="absolute bottom-10 right-10 w-[400px] h-[400px] rounded-full bg-white/5 blur-[100px] pointer-events-none"></div>
  </div>

</div>
