export interface Anime {
  mal_id: number
  title: string
  title_english: string
  title_japanese: string
  images: {
    jpg: {
      large_image_url: string
    }
  }
  score: number
  synopsis: string
  rating: string
  episodes: number
  status: string
  aired: {
    from: string
  }
}
