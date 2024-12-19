import { defineStore } from 'pinia'
import type { ProfileUser } from '../classes'

// Используем существующий интерфейс ProfileUser
export const useProfileStore = defineStore('profile', {
  state: () => ({
    profileData: {} as ProfileUser,      // Оригинальные данные
    temporaryData: {} as ProfileUser,    // Временные данные для редактирования
  }),

  actions: {
    setProfileData(data: ProfileUser) {
      this.profileData = { ...data };
      this.temporaryData = { ...data };
    },

    completeUpdateProfile() {
      this.profileData = { ...this.temporaryData };
    }
  }
})
