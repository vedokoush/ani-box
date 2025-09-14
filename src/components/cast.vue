<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { getCast } from '@/apis/handlers/getCast.ts'
import { useRoute } from 'vue-router'
const route = useRoute()

const castListRef = ref<HTMLElement | null>(null)

const cast = ref<any[]>([])
function scrollLeft(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: -400, behavior: 'smooth' })
}
function scrollRight(refEl: HTMLElement | null) {
  if (refEl) refEl.scrollBy({ left: 400, behavior: 'smooth' })
}

onMounted(async () => {
  const id = route.params.id as string | number
  cast.value = await getCast(id)
})
</script>

<template>
  <div class="cast">
    <div class="scroll-container" v-if="cast.length">
      <button class="scroll-btn left" @click="scrollLeft(castListRef)">‹</button>
      <div class="cast-wrapper">
        <div class="cast-list" ref="castListRef">
          <div v-for="c in cast" :key="c.character.mal_id" class="cast-item">
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
      </div>
      <button class="scroll-btn right" @click="scrollRight(castListRef)">›</button>
    </div>
    <p v-else>Loading cast...</p>
  </div>
</template>

<style scoped></style>
