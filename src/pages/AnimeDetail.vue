<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const anime = ref<any>(null)

onMounted(async () => {
  const id = route.params.id
  const res = await fetch(`http://localhost:3000/anime/${id}`)
  const data = await res.json()
  anime.value = data.data
})
</script>

<template>
  <div v-if="anime">
    <h1>{{ anime.title }}</h1>
    <img :src="anime.images.jpg.image_url" :alt="anime.title" />
    <p>{{ anime.synopsis }}</p>
  </div>
  <p v-else>Loading...</p>
</template>
