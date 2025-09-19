<template>
  <div class="app-layout">
    <Navbar />
    <div class="main">
      <div class="content">
        <slot />
      </div>
    </div>
    <Footer />
  </div>
</template>

<script setup lang="ts">
import Navbar from '../components/navbar.vue'
import Footer from '../components/footer.vue'
</script>

<style scoped>
.app-layout {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.main {
  display: flex;
  flex: 1;
  padding-top: 50px;
}

nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  opacity: 0.8;
  background: #030e16;
}

.content {
  flex: 1;
  padding: 1rem;
}
</style>
