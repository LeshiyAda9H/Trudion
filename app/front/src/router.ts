import { createRouter, createWebHistory } from 'vue-router';
import type { RouteLocationNormalized, NavigationGuardNext } from 'vue-router';
import AuthService from './services/AuthService';
import { useUserStore } from './stores/UserStore';
import { storeToRefs } from 'pinia';

// Функция для проверки аутентификации
const checkAuthentication = (next: NavigationGuardNext): void => {
  if (!AuthService.isAuth()) {
    next('/login'); // Если пользователь не аутентифицирован, перенаправляем на страницу входа
  } else {
    next(); // Иначе продолжаем переход
  }
};

// Функция для проверки заполненности профиля
const checkProfileFilled = async (next: NavigationGuardNext): Promise<void> => {
  const userStore = useUserStore();
  const { userProfileFilled } = storeToRefs(userStore);
  if (!(await userProfileFilled.value)) {
    next('/profile-setup'); // Если профиль не заполнен, перенаправляем на страницу настройки профиля
  } else {
    next(); // Иначе продолжаем переход
  }
};

// Определение маршрутов
const routes = [
  {
    path: '/',
    redirect: () => (AuthService.isAuth() ? '/home' : '/login'), // Перенаправление на главную или страницу входа
  },

  {
    path: '/register',
    name: 'RegistrationPage',
    component: () => import('./views/Registration.vue'),
    beforeEnter: (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
      if (AuthService.isAuth()) {
        next('/home'); // Если пользователь уже аутентифицирован, перенаправляем на домашнюю страницу
      } else {
        next(); // Иначе продолжаем переход
      }
    },
  },

  {
    path: '/login',
    name: 'AuthorizationPage',
    component: () => import('./views/Authorization.vue'),
    beforeEnter: (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
      if (AuthService.isAuth()) {
        next('/home'); // Если пользователь уже аутентифицирован, перенаправляем на домашнюю страницу
      } else {
        next(); // Иначе продолжаем переход
      }
    },
  },

  {
    path: '/home',
    name: 'HomePage',
    component: () => import('./views/Home.vue'),
    beforeEnter: (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
      checkAuthentication(next); // Проверяем аутентификацию
    },
  },

  {
    path: '/profile',
    name: 'ProfilePage',
    component: () => import('./views/Profile.vue'),
    beforeEnter: async (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
      checkAuthentication(next); // Проверяем аутентификацию
      await checkProfileFilled(next); // Проверяем заполненность профиля
    },
  },

  {
    path: '/profile-setup',
    name: 'ProfileSetupPage',
    component: () => import('./views/ProfileSetup.vue'),
    beforeEnter: async (to: RouteLocationNormalized, from: RouteLocationNormalized, next: NavigationGuardNext) => {
      next()
      await checkProfileFilled(next); // Проверяем заполненность профиля
    },
  },

  {
    path: '/:pathMatch(.*)*',
    redirect: '/', // Перенаправление на главную страницу для несуществующих маршрутов
  },
];

// Создание роутера
const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
