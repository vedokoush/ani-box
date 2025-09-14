import { BASE_URL} from '@/urls/base_url.ts'

export async function getAnimeById(id: string | number) {
  const res = await fetch(`${BASE_URL}/anime/${id}`)
  const data = await res.json()
  return data.data
}
