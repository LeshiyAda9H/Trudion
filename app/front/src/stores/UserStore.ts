import { defineStore } from 'pinia'; // Импортируем функцию для создания хранилища из Pinia
import type { AuthedUser } from '../classes'; // Импортируем тип данных аутентифицированного пользователя

// Определяем хранилище с именем 'user'
export const useUserStore = defineStore('user', {
  // Определяем начальное состояние хранилища
  state: () => ({
    // Инициализируем состояние 'user' данными из localStorage или null, если данных нет
    user: JSON.parse(localStorage.getItem('user') || 'null') as AuthedUser | null,
  }),

  // Определяем действия, которые могут изменять состояние хранилища
  actions: {
    // Действие для установки данных пользователя
    setUser(userData: AuthedUser) {
      this.user = userData; // Обновляем состояние 'user'
      localStorage.setItem('user', JSON.stringify(userData)); // Сохраняем данные пользователя в localStorage
    },

    // Действие для выхода пользователя
    logout() {
      this.user = null; // Устанавливаем состояние 'user' в null
      localStorage.removeItem('user'); // Удаляем данные пользователя из localStorage
    },
  },

  // Определяем геттеры, которые возвращают данные из состояния
  getters: {
    // Геттер, который возвращает true, если пользователь аутентифицирован, и false в противном случае
    isAuthenticated: (state) => !!state.user,

    // Геттер, который возвращает email пользователя (уверены, что state.user не null)
    email: (state) => state.user!.email,
  },
});
