import { ref, computed, onBeforeUnmount, type Ref } from 'vue'
import type { Anime } from '@/types/anime.ts'
import { getAnimeBanner } from '@/apis/handlers/getAnimeBanner'

export function useFeaturedSlider(featuredAnimes: Ref<Anime[]>) {
  const currentFeaturedIndex = ref(0)
  const featuredBanner = ref<string>('')
  const isTransitioning = ref(false)
  const autoSlideInterval = ref<ReturnType<typeof setInterval> | null>(null)

  async function loadFeaturedAnime(index: number) {
    if (featuredAnimes.value[index]) {
      isTransitioning.value = true
      currentFeaturedIndex.value = index
      featuredBanner.value = await getAnimeBanner(featuredAnimes.value[index].mal_id)
      setTimeout(() => {
        isTransitioning.value = false
      }, 300)
    }
  }

  function startAutoSlide() {
    if (autoSlideInterval.value) clearInterval(autoSlideInterval.value)
    autoSlideInterval.value = setInterval(() => {
      nextSlide()
    }, 5000)
  }

  function stopAutoSlide() {
    if (autoSlideInterval.value) {
      clearInterval(autoSlideInterval.value)
      autoSlideInterval.value = null
    }
  }

  function nextSlide() {
    const nextIndex = (currentFeaturedIndex.value + 1) % featuredAnimes.value.length
    loadFeaturedAnime(nextIndex)
  }

  function prevSlide() {
    const prevIndex =
      currentFeaturedIndex.value === 0
        ? featuredAnimes.value.length - 1
        : currentFeaturedIndex.value - 1
    loadFeaturedAnime(prevIndex)
  }

  function goToSlide(index: number) {
    stopAutoSlide()
    loadFeaturedAnime(index)
    startAutoSlide()
  }

  onBeforeUnmount(() => stopAutoSlide())

  const displayTitle = computed(() => {
    const anime = featuredAnimes.value[currentFeaturedIndex.value]
    return anime?.title_english || anime?.title || anime?.title_japanese || ''
  })

  const truncatedSynopsis = computed(() => {
    const anime = featuredAnimes.value[currentFeaturedIndex.value]
    if (!anime?.synopsis) return ''
    return anime.synopsis.length > 300
      ? anime.synopsis.substring(0, 300) + '...'
      : anime.synopsis
  })

  const currentAnime = computed(() => featuredAnimes.value[currentFeaturedIndex.value])

  return {
    currentFeaturedIndex,
    featuredBanner,
    isTransitioning,
    loadFeaturedAnime,
    startAutoSlide,
    stopAutoSlide,
    nextSlide,
    prevSlide,
    goToSlide,
    displayTitle,
    truncatedSynopsis,
    currentAnime
  }
}
