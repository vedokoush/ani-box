<script setup lang="ts">
import Logo from '@/assets/anibox.png'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import searchbar from './searchbar.vue'
const router = useRouter()
</script>

<template>
  <nav class="navbar">
    <div class="logo-section">
      <img :src="Logo" alt="logo" class="logo" />
      <a href="/">
        <span class="brand">AniBox</span>
      </a>
    </div>

    <div class="links">
      <searchbar/>
      <a href="/">Home</a>
      <a href="#">Favorites</a>
      <a href="#">My List</a>
    </div>
  </nav>
</template>

<style scoped>
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #030e16;
  color: #ff1b75;
  padding: 0.75rem 1.5rem;
  transition: opacity 0.8s ease;
}

.navbar:hover {
  opacity: 1;
}

.logo-section {
  display: flex;
  align-items: center;
  gap: 10px;
}

.logo {
  width: 40px;
  height: 40px;
}

.brand {
  font-size: 1.2rem;
  font-weight: bold;
}

.links {
  display: flex;
  align-items: center;
  gap: 1.5rem;
}

</style>
