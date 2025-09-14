<script setup lang="ts">
import { onMounted } from 'vue'

onMounted(() => {
  const box = document.querySelector('.search-box') as HTMLElement | null
  const btn_search = document.querySelector('.btn-search') as HTMLElement | null
  const btn_close = document.querySelector('.btn-close') as HTMLElement | null

  btn_search?.addEventListener('click', (e) => {
    e.preventDefault()
    box?.classList.add('active')
  })

  btn_close?.addEventListener('click', (e) => {
    e.preventDefault()
    box?.classList.remove('active')
  })
})
</script>

<template>
  <div class="search-box">
    <a href="#" class="btn-search"><i class="uil uil-search"></i></a>
    <input type="text" placeholder="Search anime..." />
    <a href="#" class="btn-close"><i class="uil uil-times"></i></a>
  </div>
</template>

<style scoped>
.search-box {
  background-color: #121d26;
  padding: 6px 10px;
  border-radius: 25px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  box-shadow: 0 2px 8px rgba(0,0,0,0.4);
  transition: all 0.3s ease;
  overflow: hidden;
}

.search-box.active {
  box-shadow: 0 0 8px #ff1b75;
}

.search-box input {
  width: 0;
  height: 100%;
  outline: none;
  border: none;
  font-size: 15px;
  font-weight: 500;
  background: transparent;
  color: #e3e3e3;
  transition: width 0.3s ease;
}

.search-box .btn-close {
  display: none;
}

.search-box.active .btn-close {
  display: block;
}

.search-box.active input {
  width: 250px;
  padding-left: 10px;
}

.search-box input::placeholder {
  color: #777;
}
</style>
