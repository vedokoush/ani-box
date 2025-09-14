<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const recommendations = ref<any[]>([])
const recListRef = ref<HTMLElement | null>(null)

onMounted(async () => {
  const id = route.params.id
  const resRec = await fetch(`http://172.16.22.56:3000/anime/${id}/recommendations`)
  const dataRec = await resRec.json()
  recommendations.value = dataRec.data || []
})

function scrollLeft(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: -400, behavior: 'smooth' })
}

function scrollRight(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: 400, behavior: 'smooth' })
}
</script>

<template>
  <div class="more-like-this">
    <h2 v-if="recommendations.length">More Like This</h2>
    <div class="scroll-container" v-if="recommendations.length">
      <button class="scroll-btn left" @click="scrollLeft(recListRef)">‹</button>
      <div class="recommend-list" ref="recListRef">
        <router-link
          v-for="rec in recommendations"
          :key="rec.entry.mal_id"
          :to="`/anime/${rec.entry.mal_id}`"
          class="recommend-item"
        >
          <img
            :src="rec.entry.images.jpg.large_image_url"
            :alt="rec.entry.title"
            class="recommend-img"
          />
          <p class="recommend-title">{{ rec.entry.title }}</p>
        </router-link>
      </div>
      <button class="scroll-btn right" @click="scrollRight(recListRef)">›</button>
    </div>
  </div>
</template>

<style scoped>
.more-like-this {
  max-width: 1800px;
  margin: 40px auto;
  padding: 0 20px;
  color: white;
}

.more-like-this h2 {
  margin-bottom: 16px;
  font-weight: bold;
}

.scroll-container {
  position: relative;
  display: flex;
  align-items: center;
  width: 100%;
  overflow: hidden;
}

.recommend-list {
  display: flex;
  overflow-x: auto;
  scroll-behavior: smooth;
  gap: 16px;
  padding: 10px 0;
  width: 100%;
  scrollbar-width: none;
  -ms-overflow-style: none;
}
.recommend-list::-webkit-scrollbar {
  display: none;
}

.recommend-item {
  flex: 0 0 auto;
  width: 220px;
  color: white;
  padding: 10px;
  border-radius: 8px;
  transition: transform 0.25s ease, box-shadow 0.25s ease;
  text-decoration: none;
}
.recommend-item:hover {
  transform: translateY(-6px) scale(1.03);
  box-shadow: 0 8px 20px rgba(0,0,0,0.4);
}

.recommend-img {
  width: 100%;
  height: 300px;
  object-fit: cover;
  border-radius: 6px;
}

.recommend-title {
  font-size: 0.9rem;
  line-height: 1.3;
  margin-top: 8px;
  height: 40px;
  overflow: hidden;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}

.scroll-btn {
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  z-index: 10;
  background: rgba(0, 0, 0, 0.6);
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
  left: 0;
}
.scroll-btn.right {
  right: 0;
}
</style>
