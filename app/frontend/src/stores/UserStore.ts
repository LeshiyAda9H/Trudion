import { defineStore } from 'pinia' // Импортируем функцию для создания хранилища из Pinia
import type { AuthedUser } from '../classes' // Импортируем тип данных аутентифицированного пользователя
import AuthService from '../services/AuthService'

// Определяем хранилище с именем 'user'
export const useUserStore = defineStore('user', {
  // Определяем начальное состояние хранилища
  state: () => ({
    // Инициализируем состояние 'user' данными из localStorage или null, если данных нет
    user: JSON.parse(localStorage.getItem('user') || 'null') as AuthedUser | null,

    // Временные данные
    registrationData: {
      username: '' as string,
      email: '' as string,
      password: '' as string,
      gender: '' as string,
      biography: '' as string,
      label: '' as string,
    },
  }),

  // Определяем действия, которые могут изменять состояние хранилища
  actions: {
    // Действие для установки данных пользователя
    setUser(userData: AuthedUser) {
      this.user = userData // Обновляем состояние 'user'
      localStorage.setItem('user', JSON.stringify(userData)) // Сохраняем данные пользователя в localStorage
    },

    // Устанавливаем временные данные для регистрации (например, email и password)
    setRegistrationData(partialData: Partial<typeof this.registrationData>) {
      this.registrationData = { ...this.registrationData, ...partialData }
    },

    // Отправляем данные регистрации на сервер и сбрасываем временное состояние
    async completeRegistration() {
      try {
        AuthService.register(this.registrationData)

        this.registrationData = {
          username: '',
          email: '',
          password: '',
          gender: '',
          biography: '',
          label: '',
        }
      } catch (error) {
        console.error('Ошибка регистрации:', error)
        throw new Error('Ошибка регистрации')
      }
    },
  },

  // Определяем геттеры, которые возвращают данные из состояния
  getters: {
    // Геттер, который возвращает true, если пользователь аутентифицирован, и false в противном случае
    isAuthenticated: (state) => !!state.user,

    // Геттер, который возвращает email пользователя (уверены, что state.user не null)
    email: (state) => (state.user ? state.user.email : null),

    userProfileFilled: (state) => {
      return state.registrationData.username && state.registrationData.gender
    },
  },
})
