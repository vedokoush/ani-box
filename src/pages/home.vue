<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getSeasonAnime } from '@/apis/handlers/getSeasonAnime'
import { getTopAnime } from '@/apis/handlers/getTopAnime'
import { useFeaturedSlider } from '@/composables/useFeaturedSlider'
import type { Anime } from '@/types/anime.ts'

const airingAnimes = ref<Anime[]>([])
const topAnimes = ref<Anime[]>([])
const featuredAnimes = ref<Anime[]>([])
const isLoading = ref(true)

const airingListRef = ref<HTMLElement | null>(null)
const topListRef = ref<HTMLElement | null>(null)

function scrollLeft(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: -400, behavior: 'smooth' })
}
function scrollRight(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: 400, behavior: 'smooth' })
}

const {
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
  currentAnime,
} = useFeaturedSlider(featuredAnimes)

onMounted(async () => {
  try {
    const [airingData, topData] = await Promise.all([getSeasonAnime(), getTopAnime()])

    airingAnimes.value = airingData.slice(0, 20)
    topAnimes.value = topData.slice(0, 20)
    featuredAnimes.value = topData.slice(0, 5)

    if (featuredAnimes.value.length > 0) {
      await loadFeaturedAnime(0)
      startAutoSlide()
    }
  } catch (err) {
    console.error('Error loading anime data:', err)
  } finally {
    isLoading.value = false
  }
})
</script>

<template>
  <div class="page-container">
    <section class="hero-banner" v-if="featuredAnimes.length > 0">
      <div class="hero-background">
        <img
          v-if="featuredBanner"
          :src="featuredBanner"
          :alt="displayTitle"
          class="banner-image"
          :class="{ transitioning: isTransitioning }"
        />
        <div class="hero-overlay"></div>
      </div>

      <button
        class="hero-nav-btn prev-btn"
        @click="prevSlide"
        @mouseenter="stopAutoSlide"
        @mouseleave="startAutoSlide"
      >
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path
            d="M15 18L9 12L15 6"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>

      <button
        class="hero-nav-btn next-btn"
        @click="nextSlide"
        @mouseenter="stopAutoSlide"
        @mouseleave="startAutoSlide"
      >
        <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
          <path
            d="M9 18L15 12L9 6"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </button>

      <div class="hero-content" :class="{ transitioning: isTransitioning }">
        <div class="hero-info">
          <h1 class="hero-title">{{ displayTitle }}</h1>

          <div class="hero-meta">
            <span class="rating-badge" v-if="currentAnime?.rating">{{ currentAnime.rating }}</span>
            <span class="score-badge" v-if="currentAnime?.score">⭐ {{ currentAnime.score }}</span>
            <span class="episode-badge" v-if="currentAnime?.episodes"
              >{{ currentAnime.episodes }} Episodes</span
            >
          </div>

          <p class="hero-description">{{ truncatedSynopsis }}</p>

          <div class="hero-actions">

            <router-link
              v-if="currentAnime"
              :to="`/anime/${currentAnime.mal_id}`"
              class="btn-primary"
            >
              Watch Now
            </router-link>


            <button class="btn-secondary">Add to List</button>
          </div>
        </div>
      </div>

      <div class="slide-indicators">
        <button
          v-for="(anime, index) in featuredAnimes"
          :key="anime.mal_id"
          class="indicator"
          :class="{ active: index === currentFeaturedIndex }"
          @click="goToSlide(index)"
          @mouseenter="stopAutoSlide"
          @mouseleave="startAutoSlide"
        >
          <span class="indicator-text">{{ anime.title_english || anime.title }}</span>
        </button>
      </div>
    </section>

    <section class="anime-section">
      <div class="section-header">
        <h2>Recent Releases</h2>
        <div class="navigation-arrows">
          <button class="nav-arrow" @click="scrollLeft(airingListRef)">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
              <path
                d="M15 18L9 12L15 6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
          <button class="nav-arrow" @click="scrollRight(airingListRef)">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
              <path
                d="M9 18L15 12L9 6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
        </div>
      </div>

      <div class="scroll-container" v-if="airingAnimes.length">
        <div class="anime-list" ref="airingListRef">
          <router-link
            v-for="anime in airingAnimes"
            :key="anime.mal_id"
            :to="`/anime/${anime.mal_id}`"
            class="anime-card"
          >
            <div class="card-image">
              <img
                :src="anime.images.jpg.large_image_url"
                :alt="anime.title_english || anime.title"
                loading="lazy"
              />
              <div class="card-overlay">
                <div class="play-icon">▶</div>
              </div>
            </div>
            <div class="card-info">
              <h3>{{ anime.title_english || anime.title }}</h3>
              <p class="anime-status">{{ anime.status }}</p>
            </div>
          </router-link>
        </div>
      </div>

      <div v-else-if="isLoading" class="loading-state">
        <div class="loading-spinner"></div>
        <p>Loading anime...</p>
      </div>
    </section>

    <section class="anime-section">
      <div class="section-header">
        <h2>Top Rated Anime</h2>
        <div class="navigation-arrows">
          <button class="nav-arrow" @click="scrollLeft(topListRef)">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
              <path
                d="M15 18L9 12L15 6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
          <button class="nav-arrow" @click="scrollRight(topListRef)">
            <svg width="24" height="24" viewBox="0 0 24 24" fill="none">
              <path
                d="M9 18L15 12L9 6"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>
          </button>
        </div>
      </div>

      <div class="scroll-container" v-if="topAnimes.length">
        <div class="anime-list" ref="topListRef">
          <router-link
            v-for="anime in topAnimes"
            :key="anime.mal_id"
            :to="`/anime/${anime.mal_id}`"
            class="anime-card"
          >
            <div class="card-image">
              <img
                :src="anime.images.jpg.large_image_url"
                :alt="anime.title_english || anime.title"
                loading="lazy"
              />
              <div class="card-overlay">
                <div class="play-icon">▶</div>
              </div>
              <div class="score-indicator">⭐ {{ anime.score }}</div>
            </div>
            <div class="card-info">
              <h3>{{ anime.title_english || anime.title }}</h3>
              <p class="anime-score">Score: {{ anime.score }}/10</p>
            </div>
          </router-link>
        </div>
      </div>

      <div v-else-if="isLoading" class="loading-state">
        <div class="loading-spinner"></div>
        <p>Loading top anime...</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.page-container {
  min-height: 100vh;
  color: white;
  max-width: 100vw;
  overflow-x: hidden;
}

.top-anime {
  background: #e0e0e0;
}

.hero-banner {
  position: relative;
  height: 70vh;
  min-height: 500px;
  overflow: hidden;
  margin-bottom: 60px;
}

.hero-background {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
}

.banner-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: opacity 0.3s ease;
}

.hero-overlay {
  position: absolute;
  z-index: 1;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: linear-gradient(
    to bottom,
    rgba(3, 14, 22, 0) 0%,
    rgba(3, 14, 22, 0.6) 40%,
    rgba(3, 14, 22, 0.9) 80%,
    rgba(3, 14, 22, 1) 100%
  );
}

.hero-content {
  position: relative;
  z-index: 5;
  height: 100%;
  display: flex;
  align-items: center;
  padding: 0 60px;
  transition: opacity 0.3s ease;
}

.hero-info {
  max-width: 1200px;
  padding-top: 400px;
  padding-left: 50px;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 700;
  margin-bottom: 20px;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
  line-height: 1.2;
}

.hero-meta {
  display: flex;
  gap: 15px;
  margin-bottom: 25px;
  flex-wrap: wrap;
}

.rating-badge,
.score-badge,
.episode-badge {
  padding: 6px 12px;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 600;
  backdrop-filter: blur(10px);
}

.rating-badge {
  background: rgba(220, 53, 69, 0.8);
}

.score-badge {
  background: rgba(255, 193, 7, 0.8);
  color: #000;
}

.episode-badge {
  background: rgba(13, 202, 240, 0.8);
}

.hero-description {
  font-size: 1.1rem;
  line-height: 1.6;
  margin-bottom: 30px;
  color: rgba(255, 255, 255, 0.9);
  text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.8);
}

.hero-actions {
  display: flex;
  gap: 15px;
  flex-wrap: wrap;
}

.btn-primary {
  padding: 12px 30px;
  background: linear-gradient(135deg, #e91e63, #f44336);
  color: white;
  text-decoration: none;
  border-radius: 25px;
  font-weight: 600;
  transition: all 0.3s ease;
  box-shadow: 0 4px 15px rgba(233, 30, 99, 0.3);
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 25px rgba(233, 30, 99, 0.4);
}

.btn-secondary {
  padding: 12px 30px;
  background: rgba(255, 255, 255, 0.1);
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-radius: 25px;
  font-weight: 600;
  transition: all 0.3s ease;
  cursor: pointer;
  backdrop-filter: blur(10px);
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: rgba(255, 255, 255, 0.5);
  transform: translateY(-2px);
}

.hero-nav-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 3;
  width: 50px;
  height: 50px;
  border: none;
  border-radius: 50%;
  background: rgba(0, 0, 0, 0.6);
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.hero-nav-btn:hover {
  background: rgba(0, 0, 0, 0.8);
  transform: translateY(-50%) scale(1.1);
}

.prev-btn {
  left: 30px;
}

.next-btn {
  right: 30px;
}

.slide-indicators {
  position: absolute;
  bottom: 30px;
  padding-top: 100px;
  left: 120px;
  z-index: 3;
  display: flex;
  gap: 12px;
}

.indicator {
  position: relative;
  padding: 0;
  border: none;
  background: none;
  cursor: pointer;
  transition: all 0.3s ease;
}

.indicator::before {
  content: '';
  display: block;
  width: 40px;
  height: 4px;
  background: rgba(255, 255, 255, 0.4);
  border-radius: 2px;
  transition: all 0.3s ease;
}

.indicator.active::before {
  background: #e91e63;
  width: 60px;
}

.indicator:hover::before {
  background: rgba(255, 255, 255, 0.7);
}

.indicator-text {
  position: absolute;
  bottom: 15px;
  left: 0;
  font-size: 0.8rem;
  color: white;
  white-space: nowrap;
  opacity: 0;
  transform: translateY(10px);
  transition: all 0.3s ease;
  pointer-events: none;
}

.indicator:hover .indicator-text {
  opacity: 1;
  transform: translateY(0);
}

.banner-image.transitioning,
.hero-content.transitioning {
  opacity: 0.7;
  transition: opacity 0.3s ease;
}

.anime-section {
  margin-bottom: 60px;
  padding: 0 60px;
  max-width: 100%;
  overflow: hidden;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.section-header h2 {
  font-size: 2rem;
  font-weight: 600;
  color: white;
}

.navigation-arrows {
  display: flex;
  gap: 10px;
}

.nav-arrow {
  width: 44px;
  height: 44px;
  border: none;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  color: white;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  backdrop-filter: blur(10px);
}

.nav-arrow:hover {
  background: rgba(255, 255, 255, 0.2);
  transform: scale(1.1);
}

.scroll-container {
  position: relative;
  max-width: 100%;
  overflow: hidden;
}

.anime-list {
  display: flex;
  overflow-x: auto;
  scroll-behavior: smooth;
  gap: 20px;
  padding: 20px 0;
  scrollbar-width: none;
  -ms-overflow-style: none;
  max-width: 100%;
  width: 100%;
}

.anime-list::-webkit-scrollbar {
  display: none;
}

.anime-card {
  flex: 0 0 auto;
  width: 280px;
  text-decoration: none;
  color: white;
  transition: all 0.4s ease;
}

.anime-card:hover {
  transform: translateY(-10px) scale(1.05);
}

.card-image {
  position: relative;
  //border-radius: 15px;
  overflow: hidden;
  margin-bottom: 15px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.3);
}

.card-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: all 0.4s ease;
}

.card-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: all 0.3s ease;
}

.anime-card:hover .card-overlay {
  opacity: 1;
}

.play-icon {
  font-size: 3rem;
  color: white;
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.8);
}

.score-indicator {
  position: absolute;
  top: 12px;
  right: 12px;
  background: rgba(0, 0, 0, 0.8);
  padding: 5px 10px;
  border-radius: 20px;
  font-size: 0.85rem;
  font-weight: 600;
  backdrop-filter: blur(5px);
}

.card-info h3 {
  font-size: 1.1rem;
  font-weight: 600;
  margin-bottom: 8px;
  line-height: 1.3;
  height: 50px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.card-info p {
  font-size: 0.9rem;
  color: rgba(255, 255, 255, 0.7);
  margin: 0;
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 200px;
  color: rgba(255, 255, 255, 0.7);
}

.loading-spinner {
  width: 40px;
  height: 40px;
  border: 3px solid rgba(255, 255, 255, 0.3);
  border-top: 3px solid #e91e63;
  border-radius: 50%;
  animation: spin 1s linear infinite;
  margin-bottom: 15px;
}

@keyframes spin {
  0% {
    transform: rotate(0deg);
  }
  100% {
    transform: rotate(360deg);
  }
}

@media (max-width: 1200px) {
  .hero-content,
  .anime-section {
    padding: 0 40px;
  }

  .hero-title {
    font-size: 3rem;
  }

  .anime-card {
    width: 240px;
  }

  .slide-indicators {
    left: 40px;
  }
}

@media (max-width: 768px) {
  .hero-content,
  .anime-section {
    padding: 0 20px;
  }

  .hero-title {
    font-size: 2.5rem;
  }

  .hero-meta {
    flex-direction: column;
    align-items: flex-start;
    gap: 10px;
  }

  .section-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 15px;
  }

  .anime-card {
    width: 200px;
  }

  .card-image img {
    height: 280px;
  }

  .hero-nav-btn {
    width: 44px;
    height: 44px;
  }

  .prev-btn {
    left: 20px;
  }

  .next-btn {
    right: 20px;
  }

  .slide-indicators {
    bottom: 20px;
    left: 20px;
  }

  .indicator::before {
    width: 30px;
  }

  .indicator.active::before {
    width: 45px;
  }
}

@media (max-width: 480px) {
  .hero-banner {
    height: 50vh;
    min-height: 400px;
  }

  .hero-title {
    font-size: 2rem;
  }

  .anime-card {
    width: 160px;
  }

  .card-image img {
    height: 220px;
  }
}
</style>

<!-- shouko -->
