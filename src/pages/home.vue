<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { getSeasonAnime } from '@/apis/handlers/getSeasonAnime.ts'
import { getTopAnime } from '@/apis/handlers/getTopAnime.ts'

const airingAnimes = ref<any[]>([])
const topAnimes = ref<any[]>([])

onMounted(async () => {
  airingAnimes.value = await getSeasonAnime()
  topAnimes.value = await getTopAnime()
})

const airingListRef = ref<HTMLElement | null>(null)
const topListRef = ref<HTMLElement | null>(null)

function scrollLeft(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: -400, behavior: 'smooth' })
}

function scrollRight(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: 400, behavior: 'smooth' })
}
</script>

<template>
  <div class="page-container">

    <section class="anime-section">
      <h1>Currently Airing Anime</h1>
      <div class="scroll-container" v-if="airingAnimes.length">
        <button class="scroll-btn left" @click="scrollLeft(airingListRef)">‹</button>
        <div class="anime-list" ref="airingListRef">
          <router-link
            v-for="anime in airingAnimes"
            :key="anime.mal_id"
            :to="`/anime/${anime.mal_id}`"
            class="anime-card"
          >
            <img :src="anime.images.jpg.large_image_url" :alt="anime.title_english" />
            <h3>{{ anime.title_english }}</h3>
          </router-link>
        </div>
        <button class="scroll-btn right" @click="scrollRight(airingListRef)">›</button>
      </div>
      <p v-else>Loading...</p>
    </section>

    <section class="anime-section">
      <h1>Top Anime</h1>
      <div class="scroll-container" v-if="topAnimes.length">
        <button class="scroll-btn left" @click="scrollLeft(topListRef)">‹</button>
        <div class="anime-list" ref="topListRef">
          <router-link
            v-for="anime in topAnimes"
            :key="anime.mal_id"
            :to="`/anime/${anime.mal_id}`"
            class="anime-card"
          >
            <img :src="anime.images.jpg.large_image_url" :alt="anime.title_english" />
            <h3>{{ anime.title_english }}</h3>
            <p>Score: {{ anime.score }}</p>
          </router-link>
        </div>
        <button class="scroll-btn right" @click="scrollRight(topListRef)">›</button>
      </div>
      <p v-else>Loading top anime...</p>
    </section>

  </div>
</template>

<style scoped>
.page-container {
  max-width: 100vw;
  overflow-x: hidden;
  padding: 20px;
}

.anime-section {
  margin-bottom: 40px;
}

.anime-section h1 {
  margin-bottom: 16px;
  color: white;
}

.scroll-container {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
  margin: 0 auto;
}

.anime-list {
  display: flex;
  overflow-x: auto;
  scroll-behavior: smooth;
  gap: 16px;
  padding: 10px 20px;
  width: 100%;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.anime-list::-webkit-scrollbar {
  display: none;
}

.anime-card {
  flex: 0 0 auto;
  width: 340px;
  color: white;
  padding: 10px;
  border-radius: 8px;
  transition: transform 0.25s ease, box-shadow 0.25s ease;
  text-decoration: none;
}

.anime-card:hover {
  transform: translateY(-6px) scale(1.03);
  box-shadow: 0 8px 20px rgba(0,0,0,0.4);
}

.anime-card img {
  width: 100%;
  height: auto;
  object-fit: contain;
}

.anime-card h3 {
  font-size: 14px;
  margin: 2px 0;
  line-height: 1.3;
  height: 36px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.anime-card p {
  font-size: 12px;
  color: #ccc;
  margin: 4px 0;
}

.scroll-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 10;
  background: rgba(0, 0, 0, 0.7);
  color: white;
  border: none;
  width: 40px;
  height: 40px;
  border-radius: 50%;
  font-size: 20px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: background-color 0.2s ease;
}

.scroll-btn:hover {
  background: rgba(0, 0, 0, 0.9);
}

.scroll-btn.left {
  left: -20px;
}

.scroll-btn.right {
  right: -20px;
}

@media (max-width: 768px) {
  .page-container {
    padding: 10px;
  }

  .anime-card {
    width: 150px;
  }

  .anime-card img {
    height: 200px;
  }

  .scroll-btn {
    width: 35px;
    height: 35px;
    font-size: 18px;
  }

  .scroll-btn.left {
    left: -15px;
  }

  .scroll-btn.right {
    right: -15px;
  }
}
</style>
