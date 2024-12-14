<template>
  <nav class="navbar">
    <div class="navbar-left">
      <RouterLink to="/home" class="logo">
        <div class="logo-ellipse">Trudion</div>
      </RouterLink>
    </div>

    <div class="navbar-right">
      <RouterLink to="/messenger" class="nav-icon">
        <i class="fas fa-envelope"></i>
      </RouterLink>

      <div class="nav-icon notification-icon" @click="toggleNotifications">
        <i class="fas fa-bell"></i>
        <!-- Выпадающее меню уведомлений (пока скрыто) -->
        <div v-if="showNotifications" class="dropdown-menu notifications">
          <div class="dropdown-content">
            <!-- Здесь будет контент уведомлений -->
            <p>Уведомления пока недоступны</p>
          </div>
        </div>
      </div>

      <div class="nav-icon settings-icon" @click="toggleSettings">
        <i class="fas fa-cog"></i>
        <!-- Выпадающее меню настроек (пока скрыто) -->
        <div v-if="showSettings" class="dropdown-menu settings">
          <div class="dropdown-content">
            <!-- Здесь будет контент настроек -->
            <p>Настройки пока недоступны</p>
          </div>
        </div>
      </div>

      <RouterLink to="/profile" class="nav-icon profile-icon">
        <img :src="userAvatar || defaultAvatar" alt="Profile" class="avatar" />
      </RouterLink>
    </div>
  </nav>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import defaultAvatarImage from '../assets/default-avatar.png'; // Добавьте дефолтную аватарку

export default defineComponent({
  name: 'NavbarHeader',
  setup() {
    const showNotifications = ref(false);
    const showSettings = ref(false);
    const userAvatar = ref(''); // Здесь должен быть путь к аватарке пользователя
    const defaultAvatar = defaultAvatarImage;

    const toggleNotifications = () => {
      showNotifications.value = !showNotifications.value;
      showSettings.value = false; // Закрываем другое меню
    };

    const toggleSettings = () => {
      showSettings.value = !showSettings.value;
      showNotifications.value = false; // Закрываем другое меню
    };

    return {
      showNotifications,
      showSettings,
      toggleNotifications,
      toggleSettings,
      userAvatar,
      defaultAvatar
    };
  },
});
</script>

<style scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 2em;
  border-bottom: 1px solid #e0e0e0;

  background-color: #fff;
  color: white;
  z-index: 1000;
}

.navbar-left {
  display: flex;
  align-items: center;
}

.logo-ellipse {
  background-color: var(--primary-color);
  padding: 8px 20px;
  border-radius: 20px;
  color: white;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
}

.logo-ellipse:hover {
  background-color: #666666;
}

.navbar-right {
  display: flex;
  align-items: center;
  gap: 20px;
}

.nav-icon {
  color: var(--primary-color);
  font-size: 30px;
  cursor: pointer;
  position: relative;
  transition: color 0.3s;
}

.nav-icon:hover {
  color: var(--secondary-color);
}

.avatar {
  width: 35px;
  height: 35px;
  border-radius: 50%;
  object-fit: cover;
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  right: 0;
  background-color: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  min-width: 200px;
  margin-top: 10px;
  animation: slideDown 0.3s ease;
}

.dropdown-content {
  padding: 10px;
  color: #333;
}

@keyframes slideDown {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
