import { BASE_URL} from '@/urls/base_url.ts'

export async function getAnimeBanner(id: string | number) {
  const res = await fetch(`${BASE_URL}/anime/${id}/banner`)
  const data = await res.json()
  return data.data?.Media?.bannerImage || null
}
