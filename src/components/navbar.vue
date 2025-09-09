<script setup lang="ts">
import Logo from '@/assets/anibox.png'
import { ref } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()

const query = ref('')
const showSearch = ref(false)

function toggleSearch() {
  showSearch.value = !showSearch.value
}

async function onSearch() {
  if (!query.value.trim()) return
  router.push({ path: '/search', query: { q: query.value } })
  showSearch.value = false
}
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
      <div class="search-container">
        <transition name="slide">
          <input
            v-if="showSearch"
            v-model="query"
            type="text"
            placeholder="Search..."
            @keyup.enter="onSearch"
            class="search-input"
          />
        </transition>
        <button class="search-btn" @click="toggleSearch">
          <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M784-120 532-372q-30 24-69 38t-83 14q-109 0-184.5-75.5T120-580q0-109 75.5-184.5T380-840q109 0 184.5 75.5T640-580q0 44-14 83t-38 69l252 252-56 56ZM380-400q75 0 127.5-52.5T560-580q0-75-52.5-127.5T380-760q-75 0-127.5 52.5T200-580q0 75 52.5 127.5T380-400Z"/></svg>
        </button>
      </div>
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

.search-container {
  display: flex;
  align-items: center;
  position: relative;
}

.search-btn {
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;
}

.search-btn .icon {
  width: 20px;
  height: 20px;
  stroke: #ff1b75;
  transition: stroke 0.2s;
}

.search-btn:hover .icon {
  stroke: #ff9800;
}

.search-input {
  position: absolute;
  right: 40px;
  background: #fff;
  border: none;
  border-radius: 20px;
  padding: 6px 12px;
  font-size: 14px;
  outline: none;
  box-shadow: 0 2px 6px rgba(0,0,0,0.15);
  width: 180px;
}

.slide-enter-active,
.slide-leave-active {
  transition: all 0.3s ease;
}
.slide-enter-from,
.slide-leave-to {
  opacity: 0;
  transform: translateX(20px);
}
</style>
