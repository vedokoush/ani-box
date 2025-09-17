import { BASE_URL} from '@/urls/base_url.ts'

const cache = new Map<string, any>()

export async function getTopAnime() {
  const url = `${BASE_URL}/top/anime`
  if (cache.has(url)) return cache.get(url)

  const controller = new AbortController()
  const timeout = setTimeout(() => controller.abort(), 10000)
  try {
    const res = await fetch(url, { signal: controller.signal })
    const data = await res.json()
    cache.set(url, data.data)
    return data.data
  } finally {
    clearTimeout(timeout)
  }
}
