import type { User } from '$lib/types';

const AUTH_KEY = 'satpol_auth';

export interface AuthSession {
  token: string;
  user: User;
}

export function getSession(): AuthSession | null {
  if (typeof sessionStorage === 'undefined') return null;
  try {
    const raw = sessionStorage.getItem(AUTH_KEY);
    if (!raw) return null;
    const parsed = JSON.parse(raw);
    return parsed?.user ? (parsed as AuthSession) : null;
  } catch {
    return null;
  }
}

export function saveSession(session: AuthSession) {
  sessionStorage.setItem(AUTH_KEY, JSON.stringify(session));
}

export function clearSession() {
  if (typeof sessionStorage !== 'undefined') sessionStorage.removeItem(AUTH_KEY);
}
