import { defineStore } from 'pinia'; // Импортируем функцию для создания хранилища из Pinia
import type { AuthedUser } from '../classes'; // Импортируем тип данных аутентифицированного пользователя

// Определяем хранилище с именем 'user'
export const useUserStore = defineStore('user', {
  // Определяем начальное состояние хранилища
  state: () => ({
    // Инициализируем состояние 'user' данными из localStorage или null, если данных нет
    user: JSON.parse(localStorage.getItem('user') || 'null') as AuthedUser | null,

    // Временные данные для многоэтапной регистрации
    registrationData: {
      email: '' as string,
      password: '' as string,
      nickname: '' as string,
      gender: '' as string,
      bio: '' as string,
      skills: '' as string,
    },
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

    // Устанавливаем временные данные для регистрации (например, email и password)
    setRegistrationData(partialData: Partial<typeof this.registrationData>) {
      this.registrationData = { ...this.registrationData, ...partialData };
    },


    // Отправляем данные регистрации на сервер и сбрасываем временное состояние
    async completeRegistration(){
      // Здесь вы можете вызвать API для завершения регистрации
      const response = await fetch('/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(this.registrationData),
      });

      if (!response.ok) {
        throw new Error('Ошибка регистрации');
      }

      // После успешной регистрации очищаем временные данные
      this.registrationData = {
        email: '',
        password: '',
        nickname: '',
        gender: '',
        bio: '',
        skills: '',
      };
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
