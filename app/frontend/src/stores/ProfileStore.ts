import { defineStore } from 'pinia' // Импортируем функцию для создания хранилища из Pinia
import type { AuthedUser } from '../classes' // Импортируем тип данных аутентифицированного пользователя
import AuthService from '../services/AuthService'

// Определяем хранилище с именем 'user'
export const useProfileStore = defineStore('user', {
  // Определяем начальное состояние хранилища
  state: () => ({
    // Инициализируем состояние 'user' данными из localStorage или null, если данных нет
    user: JSON.parse(localStorage.getItem('user') || 'null') as AuthedUser | null,

    // Временные данные
    temporaryData: {
      username: '' as string,
      gender: '' as string,
      biography: '' as string,
      label: [] as string[],
      online_status: '' as string,
    },
  }),

  // Определяем действия, которые могут изменять состояние хранилища
  actions: {
    // Отправляем данные регистрации на сервер и сбрасываем временное состояние
    async completeUpdateProfile() {
      try {
        AuthService.updateProfile(this.temporaryData)

        this.temporaryData = {
          username: '',
          gender: '',
          biography: '',
          label: [],
          online_status: '',
        }
      } catch (error) {
        console.error('Profile update error:', error)
        throw new Error('Ошибка обновления профиля')
      }
    },
  },
})
