<script lang="ts">
  import { goto } from '$app/navigation';
  import { loginAdmin } from '$lib/api/client';

  let username = '';
  let password = '';
  let errorMessage = '';
  let isLoading = false;

  async function handleLogin() {
    errorMessage = '';
    isLoading = true;
    try {
      const res = await loginAdmin(username, password);
      if (res.user) {
        sessionStorage.setItem('satpol_auth', JSON.stringify(res));
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

<div class="py-16 px-4 max-w-md mx-auto">
  <div class="bg-white border border-[#eee9db] shadow-md p-8 space-y-6">
    <div class="text-center space-y-2">
      <img src="/uploads/foto_website/1686928245logo_black.png" alt="Logo Satpol PP" class="h-16 w-auto mx-auto object-contain" />
      <h2 class="text-xl font-bold uppercase text-[#101116] tracking-tight">Portal Masuk Petugas</h2>
      <p class="text-xs text-[#6f7775]">Satuan Polisi Pamong Praja Prov. Sulawesi Selatan</p>
    </div>

    {#if errorMessage}
      <div class="bg-red-50 border border-red-200 text-red-800 text-xs p-3 rounded text-center">
        {errorMessage}
      </div>
    {/if}

    <form on:submit|preventDefault={handleLogin} class="space-y-4">
      <div>
        <label for="admin-username" class="block text-xs font-bold uppercase text-[#101116] mb-1">Nama Pengguna (Username)</label>
        <input 
          id="admin-username"
          type="text" 
          bind:value={username} 
          required 
          placeholder="Username akun petugas"
          class="w-full border border-gray-300 rounded p-2.5 text-xs focus:border-[#F08300] focus:outline-none"
        />
      </div>

      <div>
        <label for="admin-password" class="block text-xs font-bold uppercase text-[#101116] mb-1">Kata Sandi (Password)</label>
        <input 
          id="admin-password"
          type="password" 
          bind:value={password} 
          required 
          placeholder="••••••••"
          class="w-full border border-gray-300 rounded p-2.5 text-xs focus:border-[#F08300] focus:outline-none"
        />
      </div>

      <button 
        type="submit" 
        disabled={isLoading}
        class="w-full bg-[#F08300] hover:bg-[#d97400] text-white font-bold text-xs uppercase py-3 rounded transition disabled:opacity-50"
      >
        {#if isLoading}Memproses...{:else}Masuk ke Sistem{/if}
      </button>
    </form>

    <!-- Quick Credentials Helper for Development -->
    <div class="border-t border-gray-100 pt-4 text-xs space-y-2">
      <div class="font-bold text-[#101116] uppercase text-[10px] tracking-wider">Akun Pengujian (Testing):</div>
      <div class="flex flex-col gap-1.5">
        <button 
          type="button" 
          on:click={() => setCredentials('stafftrantibum', 'trantibum2026!')}
          class="text-left bg-[#f9f4e8] hover:bg-[#eee9db] p-2 border border-[#eee9db] text-xs text-[#101116] transition"
        >
          Staff Satpol PP: <strong>stafftrantibum</strong>
        </button>
        <button 
          type="button" 
          on:click={() => setCredentials('superadmin', 'sulselprov2026!')}
          class="text-left bg-gray-50 hover:bg-gray-100 p-2 border border-gray-200 text-xs text-[#101116] transition"
        >
          Superadmin Diskominfo: <strong>superadmin</strong>
        </button>
      </div>
    </div>
  </div>
</div>
