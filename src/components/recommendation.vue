<script setup lang="ts">
  import { ref, onMounted } from 'vue'
  import { useRoute } from 'vue-router'

  const route = useRoute()
  const recommendations = ref<any[]>([])


  onMounted(async () => {
    const id = route.params.id
    const resRec = await fetch(`http://localhost:3000/anime/${id}/recommendations`)
    const dataRec = await resRec.json()
    recommendations.value = dataRec.data || []
  })

  // function scrollLeft(refEl: HTMLElement | null) {
  //   if (refEl) refEl.scrollBy({ left: -400, behavior: 'smooth' })
  // }
  //
  // function scrollRight(refEl: HTMLElement | null) {
  //   if (refEl) refEl.scrollBy({ left: 400, behavior: 'smooth' })
  // }

</script>

<template>
  <div class="more-like-this" v-if="recommendations.length">
    <h2>More Like This</h2>
    <div class="recommend-list">
      <div
        v-for="rec in recommendations"
        :key="rec.entry.mal_id"
        class="recommend-item"
      >
        <img
          :src="rec.entry.images.jpg.large_image_url"
          :alt="rec.entry.title"
          class="recommend-img"
        />
        <p class="recommend-title">{{ rec.entry.title }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.recommend-list {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.recommend-item {
  width: 160px;
  text-align: center;
  color: #ccc;
}

.recommend-img {
  width: 100%;
  border-radius: 8px;
  object-fit: cover;
  margin-bottom: 10px;
}

.recommend-title {
  font-size: 0.9rem;
  line-height: 1.3;
}

</style>
