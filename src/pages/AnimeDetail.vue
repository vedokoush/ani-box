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
  <div v-if="anime" class="anime-card">
    <div class="anime-card-left">
      <img :src="anime.images.jpg.large_image_url" :alt="anime.title" class="anime-poster"/>
    </div>

    <div class="anime-card-right">
      <h1 class="anime-title">{{ anime.title_english }}</h1>

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
        <button class="add-to-list">Add to List</button>
      </div>



      <div class="content-container">
        <div class="content-left">
          <p class="anime-synopsis">{{ anime.synopsis }}</p>
        </div>
      </div>
    </div>
  </div>

  <p v-else class="loading">Loading...</p>
</template>

<style scoped>
.anime-card {
  display: flex;
  gap: 30px;
  max-width: 4000px;
  margin: 50px auto;
  background: #030e16;
  border-radius: 20px;
  padding: 30px;
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

.anime-details {
  margin-top: 10px;
  margin-left: 100px;
  color: #c0c0c0;
  font-size: 0.9rem;
  line-height: 1.5;
}

.anime-details p {
  margin: 6px 0;
}

.anime-details strong {
  color: white;
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

@media (max-width: 768px) {
  .anime-card {
    flex-direction: column;
    align-items: center;
  }
  .anime-card-left {
    display: flex;
    justify-content: center;
  }
  .anime-poster {
    max-width: 250px;
  }
  .anime-title {
    font-size: 2rem;
    text-align: center;
  }
  .anime-meta, .anime-additional {
    justify-content: center;
  }
  .btn {
    justify-content: center;
  }
}
</style>
