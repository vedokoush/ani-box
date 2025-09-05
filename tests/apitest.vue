<script setup>
import { ref, onMounted } from 'vue'

const anime = ref(null)

async function fetchAnime() {
  const res = await fetch("http://localhost:3000/anime/1") // gọi backend
  const data = await res.json()
  anime.value = data.data // Jikan API trả về { data: {...} }
}

onMounted(fetchAnime)
</script>

<template>
  <div>
    <h1>Anime Info</h1>
    <div v-if="anime">
      <h2>{{ anime.title }}</h2>
      <img :src="anime.images.jpg.image_url" :alt="anime.title" />
    </div>
    <p v-else>Loading...</p>
  </div>
</template>