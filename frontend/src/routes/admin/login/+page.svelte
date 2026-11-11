<script lang="ts">
  import { goto } from '$app/navigation';
  import { loginAdmin } from '$lib/api/client';
  import { saveSession } from '$lib/utils/auth';

  let username = '';
  let password = '';
  let errorMessage = '';
  let isLoading = false;

  async function handleLogin() {
    errorMessage = '';
    if (!username.trim() || !password) {
      errorMessage = 'Username dan kata sandi wajib diisi.';
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

  function setCredentials(u: string, p: string) {
    username = u;
    password = p;
  }
</script>

<svelte:head>
  <title>Portal Masuk Petugas | Satpol PP Prov. Sulsel</title>
</svelte:head>

<div class="py-16 px-4 max-w-md mx-auto">
  <div class="bg-white border border-[#eee9db] rounded-2xl shadow-md p-8 space-y-6">
    <div class="text-center space-y-2">
      <img src="/uploads/foto_website/1686928245logo_black.png" alt="Logo Satpol PP" class="h-16 w-auto mx-auto object-contain" />
      <h1 class="text-xl font-black uppercase text-[#1a120c] tracking-tight">Portal Masuk Petugas</h1>
      <p class="text-xs text-[#3d261a]/75">Satuan Polisi Pamong Praja Prov. Sulawesi Selatan</p>
    </div>

    {#if errorMessage}
      <div role="alert" class="bg-red-50 border border-red-200 text-red-800 text-xs font-semibold p-3 rounded-xl text-center">
        {errorMessage}
      </div>
    {/if}

    <form on:submit|preventDefault={handleLogin} class="space-y-4">
      <div>
        <label for="admin-username" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">Nama Pengguna (Username)</label>
        <input
          id="admin-username"
          type="text"
          bind:value={username}
          required
          autocomplete="username"
          placeholder="Username akun petugas"
          class="w-full border border-[#eee9db] rounded-xl p-3 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
        />
        <p class="field-error">Username wajib diisi.</p>
      </div>

      <div>
        <label for="admin-password" class="block text-xs font-bold uppercase text-[#1a120c] mb-1.5">Kata Sandi (Password)</label>
        <input
          id="admin-password"
          type="password"
          bind:value={password}
          required
          autocomplete="current-password"
          placeholder="••••••••"
          class="w-full border border-[#eee9db] rounded-xl p-3 text-xs text-[#1a120c] focus:border-[#F08300] focus:outline-none"
        />
        <p class="field-error">Kata sandi wajib diisi.</p>
      </div>

      <button
        type="submit"
        disabled={isLoading}
        class="w-full bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase py-3 rounded-xl transition disabled:opacity-50 shadow-sm"
      >
        {#if isLoading}Memproses...{:else}Masuk ke Sistem{/if}
      </button>
    </form>

    <div class="border-t border-[#eee9db] pt-4 text-xs space-y-2">
      <div class="font-bold text-[#1a120c] uppercase text-[10px] tracking-wider">Akun Pengujian Cepat:</div>
      <div class="flex flex-col gap-1.5">
        <button
          type="button"
          on:click={() => setCredentials('stafftrantibum', 'trantibum2026!')}
          class="text-left bg-[#f9f4e8] hover:bg-[#eee9db] px-3 py-2 rounded-xl border border-[#eee9db] text-xs text-[#1a120c] transition"
        >
          Staff Satpol PP: <strong>stafftrantibum</strong>
        </button>
        <button
          type="button"
          on:click={() => setCredentials('kasatpol', 'satpolpp2026!')}
          class="text-left bg-[#f9f4e8] hover:bg-[#eee9db] px-3 py-2 rounded-xl border border-[#eee9db] text-xs text-[#1a120c] transition"
        >
          Kasatpol / Admin: <strong>kasatpol</strong>
        </button>
        <button
          type="button"
          on:click={() => setCredentials('superadmin', 'sulselprov2026!')}
          class="text-left bg-[#1a120c]/5 hover:bg-[#1a120c]/10 px-3 py-2 rounded-xl border border-[#eee9db] text-xs text-[#1a120c] transition"
        >
          Superadmin Diskominfo: <strong>superadmin</strong>
        </button>
      </div>
    </div>
  </div>
</div>
