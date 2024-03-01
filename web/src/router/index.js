import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'

import AsaiChatView from '@/views/AsaiChatView.vue'
import HomeView from '@/views/home/HomeView.vue'
import AboutView from '@/views/home/AboutView.vue'
import LoginView from '@/views/auth/LoginView.vue'
import AdminView from '@/views/admin/AdminView.vue'
import RegisterView from '@/views/auth/RegisterView.vue'
import PasswordRecoveryView from '@/views/auth/PasswordRecoveryView.vue'

import adminRoutes from './admin.router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: { requiresAuth: false },
      children: [
        {
          path: 'about/:slug?',
          name: 'about',
          meta: { requiresAuth: false },
          component: AboutView
        },
      ]
    },
    {
      path: '/register/:invite_token?',
      name: 'register',
      meta: { requiresAuth: false },
      component: RegisterView
    },
    {
      path: '/login',
      name: 'login',
      meta: { requiresAuth: false },
      component: LoginView
    },
    {
      path: '/password_recovery',
      meta: { requiresAuth: false },
      component: PasswordRecoveryView
    },
    // {
    //   path: '/chat',
    //   name: 'asai',
    //   meta: { requiresAuth: false },
    //   component: AsaiChatView
    // },
    { path: '/admin',
      component: AdminView,
      children: adminRoutes
    },
   { path: '/:pathMatch(.*)*', redirect: '/' }
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

router.beforeEach(async (to) => {
  const authStore = useAuthStore();

  if (to.meta.requiresAuth && !authStore.isLogedIn) {
      return '/login';
  }
})

export default router

