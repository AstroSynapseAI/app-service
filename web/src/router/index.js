import { createRouter, createWebHistory } from 'vue-router'
import AsaiChatView from '@/views/AsaiChatView.vue'
import HomeView from '@/views/Home/HomeView.vue'
import AboutView from '@/views/Home/AboutView.vue'
import LoginView from '@/views/LoginView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      children: [
        {
          path: 'about/:slug?',
          name: 'about',
          component: AboutView
        },
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: LoginView
    },
    {
      path: '/chat',
      name: 'asai',
      component: AsaiChatView
    },
    {
      path: '/avatar',
      name: 'home',
      component: () => import('../views/Admin/HomeView.vue'),
      children: [
        {
          path: ':name',
          name: 'chat',
          component: () => import('../views/Admin/HomeView.vue')
        },
        {
          path: '/:name/personality',
          name: 'personality',
          component: () => import('../views/Admin/PersonalityView.vue')
        },
        {
          path: '/:name/agents',
          name: 'agents',
          component: () => import('../views/Admin/AgentsView.vue')
        },
        {
          path: '/:name/agents/:agent/config',
          name: 'agent-config',
          component: () => import('../views/Admin/AgentConfigView.vue')
        },
        {
          path: '/:name/plugins',
          name: 'plugins',
          component: () => import('../views/Admin/PluginsView.vue')
        },
        {
          path: '/:name/models',
          name: 'models',
          component: () => import('../views/Admin/ModelsView.vue')
        },
        {
          path: '/account',
          name: 'account',
          component: () => import('../views/Admin/AccountView.vue')
        },

      ]
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    if (to.hash) {
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve({ el: to.hash, behavior: 'smooth' })
        }, 500)
      })
    }
    else if (savedPosition) {
      return savedPosition;
    }
    else {
      return { top: 0 };
    }
  }
})

export default router