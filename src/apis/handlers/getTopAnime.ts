import { BASE_URL} from '@/urls/base_url.ts'

export async function getTopAnime() {
  const res = await fetch(`${BASE_URL}/top/anime`)
  const data = await res.json()
  return data.data
}
