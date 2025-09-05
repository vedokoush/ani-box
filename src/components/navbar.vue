<script setup lang="ts">
import Logo from '@/assets/anibox.png'
</script>

<template>
  <nav class="navbar">
    <div class="logo-section">
      <img :src="Logo" alt="logo" class="logo" />
      <span class="brand">AniBox</span>
    </div>
    <div class="links">
      <a href="#">Home</a>
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
  background: #222;
  color: white;
  padding: 0.75rem 1.5rem;
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
  gap: 1.5rem;
}

.links a {
  color: white;
  text-decoration: none;
  transition: color 0.2s ease;
}

.links a:hover {
  color: #ff9800;
}
</style>
