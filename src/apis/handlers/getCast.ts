import { BASE_URL} from '@/urls/base_url.ts'

export async function getCast(id: string | number) {
  const res = await fetch(`${BASE_URL}/anime/${id}/characters`)
  const data = await res.json()
  return data.data
}
