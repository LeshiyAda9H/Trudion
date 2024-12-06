import { defineStore } from 'pinia';
import type { AuthedUser } from '../classes'
export const useUserStore = defineStore('user', {
  state: () => ({
    user: JSON.parse(localStorage.getItem('user') || 'null') as AuthedUser | null,
  }
),
  actions: {
    setUser(userData: AuthedUser) {
      this.user = userData;
      localStorage.setItem('user', JSON.stringify(userData));
    },
    logout() {
      this.user = null;
      localStorage.removeItem('user');
    },
  },
  getters: {
    isAuthenticated: (state) => !!state.user,
    username: (state) => state.user!.username,
  },
});
