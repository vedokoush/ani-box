import { BASE_URL} from '@/urls/base_url.ts'

export async function rec(id: string | number) {
  const resRec = await fetch(`${BASE_URL}/anime/${id}/recommendations`)
  const dataRec = await resRec.json()
  return dataRec.data || []
}
