<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import Recommendation from '@/components/recommendation.vue'

const route = useRoute()
const anime = ref<any>(null)
const cast = ref<any[]>([])
const banner = ref<string | null>(null)

const castListRef = ref<HTMLElement | null>(null)

function scrollLeft(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: -400, behavior: 'smooth' })
}
function scrollRight(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: 400, behavior: 'smooth' })
}

onMounted(async () => {
  const id = route.params.id

  const res = await fetch(`http://localhost:3000/anime/${id}`)
  const data = await res.json()
  anime.value = data.data

  const resCast = await fetch(`http://localhost:3000/anime/${id}/characters`)
  const dataCast = await resCast.json()
  cast.value = dataCast.data

  const resBanner = await fetch(`http://localhost:3000/anime/${id}/banner`)
  const dataBanner = await resBanner.json()
  banner.value = dataBanner.data?.Media?.bannerImage || null
})
</script>

<template>
  <div>
    <div class="anime-banner" v-if="banner">
      <img :src="banner" alt="banner" class="banner-img"/>
    </div>

    <div v-if="anime" class="anime-card">
      <div class="anime-card-left">
        <img :src="anime.images.jpg.large_image_url" :alt="anime.title" class="anime-poster"/>
      </div>

      <div class="anime-card-right">
        <h1 class="anime-title">{{ anime.title_english || anime.title }}</h1>

        <div class="anime-meta">
          <span>{{ anime.duration || 'Unknown duration' }}</span>
          <span>❤ {{ anime.score || 'N/A' }}%</span>
          <span>{{ anime.type || 'Animation' }}</span>
          <span>{{ anime.year || 'N/A' }}</span>
          <p>{{ anime.genres?.map(g => g.name).join(' · ') || 'N/A' }}</p>
        </div>

        <div class="anime-additional">
          <span><strong>Content Advisory: </strong>{{ anime.rating || 'N/A'}}</span>
          <span>Status: {{ anime.status || 'Unknown' }}</span>
          <span><strong>Aired: </strong>{{ anime.aired.string }}</span>
        </div>

        <div class="btn">
          <button class="watch-btn">▶ Watch Now</button>
          <button class="add-to-list">+ Add to List</button>
        </div>

        <div class="content-container">
          <div class="content-left">
            <p class="anime-synopsis">{{ anime.synopsis }}</p>

            <div class="cast">
              <div class="scroll-container" v-if="cast.length">
                <button class="scroll-btn left" @click="scrollLeft(castListRef)">‹</button>
                <div class="cast-list" ref="castListRef">
                  <div
                    v-for="c in cast"
                    :key="c.character.mal_id"
                    class="cast-item"
                  >
                    <div class="cast-img">
                      <img
                        :src="c.character.images.jpg.image_url"
                        :alt="c.character.name"
                        class="character-img"
                      />
                      <img
                        :src="c.voice_actors[0]?.person.images.jpg.image_url"
                        :alt="c.voice_actors[0]?.person.name"
                        class="voice-img"
                      />
                    </div>
                    <p class="cast-name">
                      <span class="character-name">{{ c.character.name }}</span>
                      <span class="voice-name">{{ c.voice_actors[0]?.person.name }}</span>
                    </p>
                  </div>
                </div>
                <button class="scroll-btn right" @click="scrollRight(castListRef)">›</button>
              </div>
              <p v-else>Loading cast...</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <p v-else class="loading">Loading...</p>
    <div class="anime-episode">

    </div>
    <Recommendation />
  </div>
</template>

<style scoped>
.anime-banner {
  width: 100%;
  height: 300px;
  overflow: hidden;
  position: relative;
}

.banner-img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.anime-card {
  display: flex;
  gap: 30px;
  max-width: 1200px;
  margin: -100px 400px 50px auto;
  padding: 30px;
  position: relative;
  z-index: 100;
  color: #e0e0e0;
}

.content-container {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}
.content-left {
  flex: 1 1 100%;
  max-width: 100%;
  overflow-x: hidden;
}
.content-right {
  flex: 1;
  max-width: 1600px;
}

.anime-card-left {
  flex-shrink: 0;
}

.anime-poster {
  width: 100%;
  max-width: 420px;
  border-radius: 4px;
  object-fit: cover;
}

.anime-card-right {
  flex: 1;
  display: flex;
  margin-top: 80px;
  flex-direction: column;
}

.anime-title {
  font-size: 2.5rem;
  font-weight: bold;
  margin-bottom: 10px;
}

.anime-meta, .anime-additional {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  font-size: 0.9rem;
  margin-bottom: 10px;
  color: #a0a0a0;
}

.anime-synopsis {
  margin: 10px 0 20px 0;
  line-height: 1.6;
  max-width: 1000px;
  color: white;
}

.btn {
  padding: 15px 0;
  display: flex;
  gap: 10px;
}

.watch-btn {
  background: #ff1b75;
  color: white;
  border: none;
  padding: 12px 30px;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.add-to-list {
  background: #f2d5cf;
  color: #030e16;
  border: none;
  padding: 12px 30px;
  border-radius: 4px;
  cursor: pointer;
}

.scroll-container {
  position: relative;
  display: flex;
  align-items: center;
  max-width: 100%;
  margin: 0 auto;
}

.cast-list {
  display: flex;
  overflow-x: auto;
  scroll-behavior: smooth;
  gap: 16px;
  padding: 10px 0;
  max-width: 50%;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.cast-list::-webkit-scrollbar {
  display: none;
}

.cast-item {
  flex: 0 0 auto;
  text-align: center;
  width: 120px;
}
.cast-img {
  position: relative;
  width: 100px;
  height: 100px;
  margin: 0 auto 10px auto;
}
.cast-img img {
  width: 100px;
  height: 100px;
  border-radius: 50%;
  object-fit: cover;
  position: absolute;
  top: 0;
  left: 0;
  transition: opacity 0.3s ease;
}
.cast-img .voice-img {
  opacity: 0;
}
.cast-item:hover .character-img {
  opacity: 0;
}
.cast-item:hover .voice-img {
  opacity: 1;
}
.cast-name {
  font-size: 0.85rem;
  color: #fff;
}
.cast-name .voice-name {
  display: none;
}
.cast-item:hover .character-name {
  display: none;
}
.cast-item:hover .voice-name {
  display: inline;
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
</style>
