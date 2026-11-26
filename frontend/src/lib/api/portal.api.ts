import { get } from 'svelte/store';
import type { Article, OrganizationNode, PublicDocument, Regency, SiteSettings } from '$lib/types';
import {
  articlesStore,
  documentsStore,
  organizationStore,
  siteSettingsStore
} from '$lib/stores/appState';
import { MASTER_REGENCIES } from '$lib/config/constants';
import { API_BASE, authHeaders, tryFetch } from './http';

export async function fetchSiteSettings(): Promise<SiteSettings> {
  const data = await tryFetch<{ settings?: SiteSettings }>(`${API_BASE}/portal/settings`);
  return data?.settings ?? get(siteSettingsStore);
}

export async function fetchArticles(): Promise<Article[]> {
  const data = await tryFetch<{ data?: Article[] }>(`${API_BASE}/portal/articles`);
  return data?.data ?? get(articlesStore);
}

export async function fetchArticleBySlug(slug: string): Promise<Article | undefined> {
  const data = await tryFetch<{ data?: Article }>(`${API_BASE}/portal/articles/${slug}`);
  return data?.data ?? get(articlesStore).find(a => a.slug === slug);
}

export async function fetchDocuments(category?: string): Promise<PublicDocument[]> {
  const url = category
    ? `${API_BASE}/portal/documents?category=${encodeURIComponent(category)}`
    : `${API_BASE}/portal/documents`;
  const data = await tryFetch<{ data?: PublicDocument[] }>(url);
  if (data?.data) return data.data;

  const all = get(documentsStore);
  if (!category) return all;
  return all.filter(d => d.category.toLowerCase() === category.toLowerCase());
}

export async function fetchRegencies(): Promise<Regency[]> {
  const data = await tryFetch<{ data?: Regency[] }>(`${API_BASE}/portal/regencies`);
  return data?.data ?? MASTER_REGENCIES;
}

export async function fetchOrganization(): Promise<OrganizationNode[]> {
  const data = await tryFetch<{ data?: OrganizationNode[] }>(`${API_BASE}/portal/organization`);
  if (data?.data && data.data.length) {
    organizationStore.set(data.data);
    return data.data;
  }
  return get(organizationStore);
}

export async function saveOrganizationNode(node: OrganizationNode, isNew: boolean): Promise<OrganizationNode> {
  const url = isNew ? `${API_BASE}/admin/organization` : `${API_BASE}/admin/organization/${node.id}`;
  const remote = await tryFetch<{ data?: OrganizationNode }>(url, {
    method: isNew ? 'POST' : 'PUT',
    headers: { 'Content-Type': 'application/json', ...authHeaders() },
    body: JSON.stringify(node)
  });
  const saved = remote?.data ?? node;
  organizationStore.update(list => (isNew ? [...list, saved] : list.map(n => (n.id === saved.id ? saved : n))));
  return saved;
}

export async function deleteOrganizationNode(id: number): Promise<void> {
  await tryFetch(`${API_BASE}/admin/organization/${id}`, { method: 'DELETE', headers: authHeaders() });
  organizationStore.update(list => list.filter(n => n.id !== id));
}
