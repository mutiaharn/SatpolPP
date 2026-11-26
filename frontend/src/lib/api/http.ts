export const API_BASE = '/api/v1';

export function authHeaders(): Record<string, string> {
  if (typeof sessionStorage === 'undefined') return {};
  try {
    const raw = sessionStorage.getItem('satpol_auth');
    const token = raw ? JSON.parse(raw).token : '';
    return token ? { Authorization: `Bearer ${token}` } : {};
  } catch {
    return {};
  }
}

export async function tryFetch<T>(url: string, init?: RequestInit): Promise<T | null> {
  try {
    const res = await fetch(url, init);
    if (res.ok) return (await res.json()) as T;
  } catch {
    // Graceful offline fallback when backend is not running
  }
  return null;
}
