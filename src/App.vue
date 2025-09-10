<template>
  <MainLayout>
    <router-view :key="$route.params.id" />
  </MainLayout>
</template>

<script setup lang="ts">
import MainLayout from './layouts/MainLayout.vue'
</script>
