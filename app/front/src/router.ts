import { createRouter, createWebHistory } from 'vue-router'
import type { RouteLocationNormalized, NavigationGuardNext } from 'vue-router'
import AuthService from './services/AuthService'
import { useUserStore } from './stores/UserStore'
import { storeToRefs } from 'pinia'

const isAuthenticated = (): boolean => AuthService.isAuth() ?? false

const routes = [
  {
    path: '/',
    redirect: () => (isAuthenticated() ? '/home' : '/login'),
  },

  {
    path: '/register',
    name: 'RegistrationPage',
    component: () => import('./views/Registration.vue'),
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext,
    ) => {
      if (isAuthenticated()) {
        next('/home') // Если пользователь уже аутентифицирован, перенаправляем на домашнюю страницу
      } else {
        next()
      }
    },
  },

  {
    path: '/login',
    name: 'AuthorizationPage',
    component: () => import('./views/Authorization.vue'),
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext,
    ) => {
      if (isAuthenticated()) {
        next('/home') // Если пользователь уже аутентифицирован, перенаправляем на домашнюю страницу
      } else {
        next()
      }
    },
  },

  {
    path: '/home',
    name: 'HomePage',
    component: () => import('./views/Home.vue'),
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext,
    ) => {
      if (!isAuthenticated()) {
        next('/login') // Если не аутентифицирован, перенаправляем на страницу входа
      } else {
        next()
      }
    },
  },

  {
    path: '/profile',
    name: 'ProfilePage',
    component: () => import('./views/Profile.vue'),
    beforeEnter: async (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext,
    ) => {
      if (!isAuthenticated()) {
        next('/login') // Если не аутентифицирован, перенаправляем на страницу входа
      } else {
        const userStore = useUserStore()
        const { userProfileFilled } = storeToRefs(userStore) // Проверяем, заполнен ли профиль
        if (!(await userProfileFilled.value)) {
          next('/profile-setup') // Если профиль не заполнен, перенаправляем на страницу настройки профиля
        } else {
          next() // Если профиль заполнен, продолжаем переход
        }
      }
    },
  },

  {
    path: '/profile-setup',
    name: 'ProfileSetupPage',
    component: () => import('./views/ProfileSetup.vue'),
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext,
    ) => {

      const userStore = useUserStore()
      const { userProfileFilled } = storeToRefs(userStore)
      if (userProfileFilled.value) {
        next('/profile') // Если профиль уже заполнен, перенаправляем на страницу профиля
      } else {
        next() // Иначе продолжаем на страницу настройки профиля
      }

    },
  },

  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
