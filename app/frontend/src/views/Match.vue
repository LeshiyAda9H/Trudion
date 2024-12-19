<template>
  <div class="match-container">
    <div v-if="isLoading" class="loading">
      Загрузка...
    </div>

    <div v-else-if="error" class="error">
      {{ error }}
    </div>

    <!-- Список пользователей -->
    <div v-else-if="!selectedUser" class="users-grid">
      <div v-if="pendingUsers.length === 0" class="no-requests">
        Нет новых запросов в друзья
      </div>

      <div v-else v-for="user in pendingUsers"
           :key="user.user_id"
           class="user-card"
           @click="selectUser(user)">
        <img :src="user.avatar || defaultAvatar" alt="User avatar" class="user-avatar">
        <div class="user-info">
          <h3>{{ user.username }}</h3>
        </div>
        <div class="action-buttons">
          <button class="accept-btn" @click.stop="acceptMatch(user)">
            <i class="fas fa-check"></i>
          </button>

          <button class="decline-btn" @click.stop="declineUser(user)">
            <i class="fas fa-times"></i>
          </button>
        </div>
      </div>
    </div>

    <!-- Профиль выбранного пользователя -->
    <div v-else class="user-profile">
      <button class="close-btn" @click="closeProfile">
        <i class="fas fa-times"></i>
      </button>

      <div class="profile-content">
        <img :src="selectedUser.avatar || defaultAvatar" alt="User avatar" class="profile-avatar">
        <div class="profile-info">
          <h2>{{ selectedUser.username }}</h2>
          <p class="bio">{{ selectedUser.biography }}</p>
          <div class="interests">
            <span v-for="interest in selectedUser.label"
                  :key="interest"
                  class="interest-tag">
              {{ interest }}
            </span>
          </div>
        </div>

        <div class="profile-actions">
          <button class="accept-btn large" @click="acceptMatch(selectedUser)">
            <i class="fas fa-check"></i>
          </button>
          <button class="decline-btn large" @click="declineUser(selectedUser)">
            <i class="fas fa-times"></i>
          </button>

        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue'
import defaultAvatar from '../assets/default-avatar.png'
import AuthService from '../services/AuthService'
import type { ProfileUser } from '../classes'

const selectedUser = ref<ProfileUser | null>(null)
const pendingUsers = ref<ProfileUser[]>([])
const isLoading = ref(false)
const error = ref<string | null>(null)

// Загрузка запросов в друзья при монтировании компонента
onMounted(async () => {
  try {
    isLoading.value = true
    const matches = await AuthService.getMatchRequests()
    pendingUsers.value = matches
  } catch (err) {
    console.error('Ошибка при загрузке запросов в друзья:', err)
    error.value = 'Не удалось загрузить запросы в друзья'
  } finally {
    isLoading.value = false
  }
})

const selectUser = (user: ProfileUser) => {
  selectedUser.value = user
}

const closeProfile = () => {
  selectedUser.value = null
}

const acceptMatch = async (user: ProfileUser) => {
  try {
    await AuthService.acceptMatch(user.user_id)
    pendingUsers.value = pendingUsers.value.filter(u => u.user_id !== user.user_id)
    selectedUser.value = null
  } catch (err) {
    console.error('Ошибка при принятии запроса:', err)
    // Можно добавить уведомление об ошибке
  }
}

const declineUser = async (user: ProfileUser) => {
  try {
    await AuthService.declineMatch(user.user_id)
    pendingUsers.value = pendingUsers.value.filter(u => u.user_id !== user.user_id)
    selectedUser.value = null
  } catch (err) {
    console.error('Ошибка при отклонении запроса:', err)
    // Можно добавить уведомление об ошибке
  }
}
</script>

<style scoped>
.match-container {

  padding: 20px;
  height: 100%;
  position: relative;
  background-color: var(--background-color);
}

.users-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
  padding: 20px;
}

.user-card {
  background: white;
  border-radius: 15px;
  padding: 15px;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  position: relative;
  cursor: pointer;
  transition: transform 0.2s;
}

.user-card:hover {
  transform: translateY(-5px);
}

.user-avatar {
  width: 100%;

  object-fit: cover;
  border-radius: 10px;
  margin-bottom: 10px;
}

.user-info {
  text-align: center;
}

.action-buttons {
  display: flex;
  justify-content: center;
  gap: 10px;
  margin-top: 10px;
}

.accept-btn, .decline-btn {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.2s;
}

.accept-btn {
  background-color: #4CAF50;
  color: white;
}

.decline-btn {
  background-color: #f44336;
  color: white;
}

.accept-btn:hover, .decline-btn:hover {
  transform: scale(1.1);
}

/* Стили для профиля */
.user-profile {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: white;
  z-index: 10;
  padding: 20px;
}

.close-btn {
  position: absolute;
  top: 20px;
  left: 20px;
  background: none;
  border: none;
  font-size: 24px;
  cursor: pointer;
  z-index: 11;
}

.profile-content {
  max-width: 800px;
  margin: 0 auto;
  padding-top: 40px;
  text-align: center;
}

.profile-avatar {
  width: 200px;
  height: 200px;
  border-radius: 20px;
  object-fit: cover;
  margin-bottom: 20px;
}

.profile-info {
  margin-bottom: 30px;
}

.profile-info h2 {
  font-size: 24px;
  margin-bottom: 10px;
}

.age {
  color: #666;
  margin-bottom: 15px;
}

.bio {
  margin-bottom: 20px;
  line-height: 1.6;
}

.interests {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  justify-content: center;
  margin-bottom: 30px;
}

.interest-tag {
  background: #f0f0f0;
  padding: 5px 15px;
  border-radius: 20px;
  font-size: 14px;
}

.profile-actions {
  display: flex;
  justify-content: center;
  gap: 20px;
}

.profile-actions .accept-btn.large,
.profile-actions .decline-btn.large {
  width: 60px;
  height: 60px;
  font-size: 24px;
}

.loading, .error, .no-requests {
  text-align: center;
  padding: 20px;
  font-size: 18px;
}

.error {
  color: #f44336;
}

.no-requests {
  color: #666;
}
</style>
