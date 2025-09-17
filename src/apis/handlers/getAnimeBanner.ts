import { BASE_URL} from '@/urls/base_url.ts'

const cache = new Map<string, string | null>()

export async function getAnimeBanner(id: string | number) {
  const url = `${BASE_URL}/anime/${id}/banner`
  if (cache.has(url)) return cache.get(url) ?? null

  const controller = new AbortController()
  const timeout = setTimeout(() => controller.abort(), 10000)
  try {
    const res = await fetch(url, { signal: controller.signal })
    const data = await res.json()
    const banner = data.data?.Media?.bannerImage || null
    cache.set(url, banner)
    return banner
  } finally {
    clearTimeout(timeout)
  }
}
