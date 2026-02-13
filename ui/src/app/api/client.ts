const DEFAULT_BASE_URL = "http://localhost:8080";

export function getApiBaseUrl(): string {
  return import.meta.env.VITE_API_BASE_URL || DEFAULT_BASE_URL;
}

export async function apiGet<T>(path: string): Promise<T> {
  const baseUrl = getApiBaseUrl();
  const res = await fetch(`${baseUrl}${path}`, {
    method: "GET",
    headers: { Accept: "application/json" },
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

  return res.json() as Promise<T>;
}