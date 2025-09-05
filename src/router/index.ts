import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/home.vue'
import AnimeDetail from '../pages/AnimeDetail.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/anime/:id',
    name: 'AnimeDetail',
    component: AnimeDetail,
    props: true,
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
