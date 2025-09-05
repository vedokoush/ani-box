<template>
  <div>
    <h1>Currently Airing Anime</h1>
    <div v-if="airingAnimes.length">
      <div v-for="anime in airingAnimes" :key="anime.mal_id" style="margin-bottom:20px;">
        <h2>{{ anime.title }}</h2>
        <img :src="anime.images.jpg.image_url" :alt="anime.title" width="200" />
        <p>Aired: {{ anime.aired.from }} → {{ anime.aired.to }}</p>
      </div>
    </div>
    <p v-else>Loading...</p>

    <h1>Top Anime</h1>
    <div v-if="topAnimes.length">
      <div v-for="anime in topAnimes" :key="anime.mal_id" style="margin-bottom:20px;">
        <h2>{{ anime.title }}</h2>
        <img :src="anime.images.jpg.image_url" :alt="anime.title" width="200" />
        <p>Score: {{ anime.score }}</p>
      </div>
    </div>
    <p v-else>Loading top anime...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const airingAnimes = ref<any[]>([])
const topAnimes = ref<any[]>([])

async function fetchAiring() {
  const res = await fetch("http://localhost:3000/anime/now")
  const data = await res.json()
  airingAnimes.value = data.data
}

async function fetchTop() {
  const res = await fetch("http://localhost:3000/top/anime")
  const data = await res.json()
  topAnimes.value = data.data
}

onMounted(() => {
  fetchAiring()
  fetchTop()
})
</script>
