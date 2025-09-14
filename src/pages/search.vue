<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const results = ref<any[]>([])
const sortBy = ref<'default' | 'rank' | 'favorites' | 'members'>('rank')

function matchesQuery(anime: any, q: string) {
  if (!q) return true
  q = q.toLowerCase()
  return (
    anime.title?.toLowerCase().includes(q) ||
    anime.title_english?.toLowerCase().includes(q) ||
    anime.title_japanese?.toLowerCase().includes(q) ||
    (anime.title_synonyms || []).some((syn: string) => syn?.toLowerCase().includes(q))
  )
}

async function fetchSearch(q: string) {
  if (!q) return
  results.value = []
  const API_BASE = "http://localhost:3000/anime/search"
  const rawResults: any[] = []

  try {

    const resFull = await fetch(`${API_BASE}?q=${encodeURIComponent(q)}`)
    const dataFull = await resFull.json()
    rawResults.push(...(dataFull?.data || []))

    if (!rawResults.length) {
      const words = q.split(/\s+/)
      for (const word of words) {
        const res = await fetch(`${API_BASE}?q=${encodeURIComponent(word)}`)
        const data = await res.json()
        rawResults.push(...(data?.data || []))
      }
    }

    const uniqueResults = Array.from(new Map(rawResults.map(a => [a.mal_id, a])).values())

    results.value = uniqueResults.filter((anime: any) => matchesQuery(anime, q))
  } catch (err) {
    console.error("Search error:", err)
    results.value = []
  }
}

function sortedResults() {
  switch (sortBy.value) {
    case 'rank':
      return [...results.value].sort((a, b) => (a.rank || 999999) - (b.rank || 999999))
    case 'favorites':
      return [...results.value].sort((a, b) => (b.favorites || 0) - (a.favorites || 0))
    case 'members':
      return [...results.value].sort((a, b) => (b.members || 0) - (a.members || 0))
    default:
      return results.value
  }
}

onMounted(() => {
  if (route.query.q) {
    fetchSearch(route.query.q as string)
  }
})

watch(
  () => route.query.q,
  (newQ) => {
    if (newQ) fetchSearch(newQ as string)
  }
)
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
      <div
        v-for="anime in sortedResults()"
        :key="anime.mal_id"
        class="anime-card"
      >
        <router-link :to="`/anime/${anime.mal_id}`">
          <img :src="anime.images.jpg.large_image_url" :alt="anime.title" />
          <h3>{{ anime.title_english || anime.title }}</h3>
          <p>
            Rank: {{ anime.rank ?? 'N/A' }} |
            Favorites: {{ anime.favorites ?? 0 }} |
            Members: {{ anime.members ?? 0 }}
          </p>
        </router-link>
      </div>
    </div>
    <p v-else>No results found</p>
  </div>
</template>


<style scoped>
.page-container {
  padding: 20px;
  color: white;
}
.anime-card {
  margin: 10px 0;
}
.sort-bar {
  margin: 10px 0;
}
</style>
