import { createRouter, createWebHashHistory } from 'vue-router'

import New from './components/Form.vue'
import Search from './components/Table.vue'

const routes = [
  {
    name: 'new',
    path: '/',
    component: New,
  },
  {
    name: 'search',
    path: '/search',
    component: Search,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
});

export default router;
