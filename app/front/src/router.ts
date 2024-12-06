import { createRouter, createWebHistory } from 'vue-router';

// Импортируем типы с использованием type-only import
import type { RouteLocationNormalized, NavigationGuardNext } from 'vue-router';

// import Registration from './views/Registration.vue';
// import Authorization from './views/Authorization.vue';
// import Home from './views/Home.vue';

import AuthService from './services/AuthService';

const isAuthenticated = (): boolean => AuthService.isAuth();


const routes = [
  {
    path: '/',
    redirect: () => (isAuthenticated() ? '/home' : '/registration'),
  },

  {
    path: '/registration',
    name: 'RegistrationPage',
    component: () => import("./views/Registration.vue"),
    beforeEnter: (
      to: RouteLocationNormalized, // Указываем типы для параметров
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      if (isAuthenticated()) {
        next('/home');
      } else {
        next();
      }
    },
  },

  {
    path: '/login',
    name: 'AuthorizationPage',
    component: () => import("./views/Authorization.vue"),
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      if (isAuthenticated()) {
        next('/home');
      } else {
        next();
      }
    },
  },

  {
    path: '/home',
    name: 'HomePage',
    component: () => import("./views/Home.vue"),
    beforeEnter: (
      to: RouteLocationNormalized,
      from: RouteLocationNormalized,
      next: NavigationGuardNext
    ) => {
      if (!isAuthenticated()) {
        next('/registration');
      } else {
        next();
      }
    },
  },
  // {
  //   path: '/users-list',
  //   name: 'UsersList',
  //   component: UsersList,
  // },

  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
