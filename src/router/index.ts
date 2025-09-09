import { createRouter, createWebHistory } from 'vue-router'
import Home from '../pages/home.vue'
import AnimeDetail from '../pages/AnimeDetail.vue'
import Search from '../pages/Search.vue'

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
  },
  { path: '/search',
    name: 'Search',
    component: Search
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
