<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const results = ref<any[]>([])
const sortBy = ref<'default' | 'rank' | 'favorites' | 'members'>('default')

async function fetchSearch(q: string) {
  const res = await fetch(`http://localhost:3000/anime/search?q=${encodeURIComponent(q)}`)
  const data = await res.json()
  results.value = data.data || []
}

function sortedResults() {
  if (sortBy.value === 'rank') {
    return [...results.value].sort((a, b) => (a.rank || 999999) - (b.rank || 999999))
  }
  if (sortBy.value === 'favorites') {
    return [...results.value].sort((a, b) => (b.favorites || 0) - (a.favorites || 0))
  }
  if (sortBy.value === 'members') {
    return [...results.value].sort((a, b) => (b.members || 0) - (a.members || 0))
  }
  return results.value
}

onMounted(() => {
  if (route.query.q) {
    fetchSearch(route.query.q as string)
  }
})

watch(() => route.query.q, (newQ) => {
  if (newQ) fetchSearch(newQ as string)
})
</script>

<template>
  <div class="page-container">
    <h1>Search results for "{{ route.query.q }}"</h1>

    <div class="sort-bar" v-if="results.length">
      <label>Sort by: </label>
      <select v-model="sortBy">
        <option value="default">Default</option>
        <option value="rank">Rank</option>
        <option value="favorites">Favorites</option>
        <option value="members">Members</option>
      </select>
    </div>

    <div v-if="sortedResults().length">
      <div v-for="anime in sortedResults()" :key="anime.mal_id" class="anime-card">
        <router-link :to="`/anime/${anime.mal_id}`">
          <img :src="anime.images.jpg.image_url" :alt="anime.title" />
          <h3>{{ anime.title }}</h3>
          <p>Rank: {{ anime.rank ?? 'N/A' }} | Favorites: {{ anime.favorites ?? 0 }} | ID: {{ anime.mal_id ?? 0}}</p>
        </router-link>
      </div>
    </div>
    <p v-else>No results found</p>
  </div>
</template>

<style scoped>
.sort-bar {
  margin-bottom: 1rem;
  color: white;
}
.sort-bar select {
  padding: 4px 8px;
  border-radius: 6px;
}
</style>
