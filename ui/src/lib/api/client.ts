function normalizeBaseUrl(raw: string): string {
  return raw.replace(/\/+$/, '');
}

export function getApiBaseUrl(): string {
  const envUrl = import.meta.env.VITE_API_BASE_URL as string | undefined;
  const fallback = 'http://localhost:8080';
  return normalizeBaseUrl(envUrl?.trim() || fallback);
}

export async function apiGet<T>(path: string): Promise<T> {
  const baseUrl = getApiBaseUrl();
  const cleanPath = path.startsWith('/') ? path : `/${path}`;

  const res = await fetch(`${baseUrl}${cleanPath}`, {
    method: 'GET',
    headers: { Accept: 'application/json' }
  });

  if (!res.ok) {
    let message = `Request failed (${res.status})`;
    try {
      const body = await res.json();
      if (body?.error) message = body.error;
    } catch {
      // ignore
    }
    throw new Error(message);
  }

  return (await res.json()) as T;
}
