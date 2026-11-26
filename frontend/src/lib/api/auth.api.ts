import { get } from 'svelte/store';
import type { User } from '$lib/types';
import { usersStore } from '$lib/stores/appState';
import { API_BASE, authHeaders, tryFetch } from './http';

const DEMO_PASSWORDS: Record<string, string> = {
  superadmin: 'sulselprov2026!',
  kasatpol: 'satpolpp2026!',
  stafftrantibum: 'trantibum2026!'
};
const DEFAULT_PASSWORD = 'satpolpp2026';

export async function loginAdmin(username: string, password: string): Promise<{ success?: boolean; token?: string; user?: User; error?: string }> {
  try {
    const res = await fetch(`${API_BASE}/auth/login`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    if (res.ok) return await res.json();
    if (res.status === 401 || res.status === 403) {
      const data = await res.json().catch(() => ({}));
      return { success: false, error: data.error || 'Username atau kata sandi tidak valid.' };
    }
  } catch {
    // Graceful offline fallback
  }

  const uname = username.trim().toLowerCase();
  const matched = get(usersStore).find(u => u.username.toLowerCase() === uname);
  const validPassword = password === DEFAULT_PASSWORD || (DEMO_PASSWORDS[uname] && password === DEMO_PASSWORDS[uname]);

  if (matched && validPassword) {
    if (!matched.is_active) {
      return { success: false, error: 'Akun ini sedang dinonaktifkan oleh Superadmin.' };
    }
    return { success: true, token: `jwt-mock-${matched.role}-${matched.id}`, user: matched };
  }
  return { success: false, error: 'Username atau kata sandi tidak valid.' };
}

export async function logoutAdmin(): Promise<void> {
  await tryFetch(`${API_BASE}/auth/logout`, {
    method: 'POST',
    headers: authHeaders()
  });
}

export async function saveUser(user: User, isNew: boolean, password?: string): Promise<User> {
  const url = isNew ? `${API_BASE}/superadmin/users` : `${API_BASE}/superadmin/users/${user.id}`;
  const remote = await tryFetch<{ data?: User }>(url, {
    method: isNew ? 'POST' : 'PUT',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify({ ...user, password })
  });
  const saved = remote?.data ?? user;
  usersStore.update(list => (isNew ? [...list, saved] : list.map(u => (u.id === saved.id ? saved : u))));
  return saved;
}

export async function setUserActive(id: number, isActive: boolean): Promise<void> {
  await tryFetch(`${API_BASE}/superadmin/users/${id}/status`, {
    method: 'PATCH',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify({ is_active: isActive })
  });
  usersStore.update(list => list.map(u => (u.id === id ? { ...u, is_active: isActive } : u)));
}

export async function deleteUser(id: number): Promise<void> {
  await tryFetch(`${API_BASE}/superadmin/users/${id}`, { method: 'DELETE', headers: authHeaders() });
  usersStore.update(list => list.filter(u => u.id !== id));
}
