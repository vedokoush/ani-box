import { BASE_URL} from '@/urls/base_url.ts'

export async function getSeasonAnime() {
  const res = await fetch(`${BASE_URL}/anime/now`)
  const data = await res.json()
  return data.data
}
