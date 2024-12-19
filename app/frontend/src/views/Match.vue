<template>
  <div class="match-container">
    <!-- Список пользователей (показывается, когда selectedUser === null) -->
    <div v-if="!selectedUser" class="users-grid">
      <div v-for="user in pendingUsers" :key="user.id" class="user-card" @click="selectUser(user)">
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
          <p class="bio">{{ selectedUser.bio }}</p>
          <div class="interests">
            <span v-for="interest in selectedUser.interests"
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

<script lang="ts">
import { defineComponent, ref } from 'vue';
import defaultAvatar from '../assets/default-avatar.png';

export default defineComponent({
  name: 'MatchView',

  setup() {
    const selectedUser = ref(null);
    const pendingUsers = ref([
      {
        id: 1,
        username: 'Турал',
        avatar: '',
        bio: 'Люблю программировать и бананы',
        interests: ['Программирование', 'Бананы', 'Спорт']
      },
      // Добавьте больше пользователей для тестирования
    ]);

    const selectUser = (user) => {
      selectedUser.value = user;
    };

    const closeProfile = () => {
      selectedUser.value = null;
    };

    const acceptMatch = (user) => {
      // Здесь логика для создания match
      // После успешного match пользователь должен появиться в чате
      pendingUsers.value = pendingUsers.value.filter(u => u.id !== user.id);
      selectedUser.value = null;
    };

    const declineUser = (user) => {
      pendingUsers.value = pendingUsers.value.filter(u => u.id !== user.id);
      selectedUser.value = null;
    };

    return {
      selectedUser,
      pendingUsers,
      defaultAvatar,
      selectUser,
      closeProfile,
      acceptMatch,
      declineUser
    };
  }
});
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
</style>
