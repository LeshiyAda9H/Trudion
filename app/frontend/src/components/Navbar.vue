<template>
  <nav class="navbar">
    <div class="navbar-left">
      <RouterLink to="/home" class="logo">
        <div class="logo-ellipse">Trudion</div>
      </RouterLink>
    </div>

    <div class="navbar-center">
      <h1 class="page-title">{{ currentPageTitle }}</h1>
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
        <div v-if="showSettings" class="dropdown-menu settings">
          <div class="dropdown-content">
            <button class="settings-button" @click="confirmLogout">
              <i class="fas fa-sign-out-alt"></i>
              Выйти
            </button>
          </div>
        </div>
      </div>

      <RouterLink to="/profile" class="nav-icon profile-icon">
        <img :src="userAvatar || defaultAvatar" alt="Profile" class="avatar" />
      </RouterLink>
    </div>
  </nav>

  <!-- Модальное окно подтверждения -->
  <div v-if="showLogoutModal" class="modal-overlay">
    <div class="modal-content">
      <h3 class="modal-title">Подтверждение выхода</h3>
      <p class="modal-text">Вы уверены, что хотите выйти?</p>
      <div class="modal-buttons">
        <button class="modal-button confirm" @click="handleLogout" title="Подтвердить">
          <i class="fas fa-check"></i>
        </button>
        <button class="modal-button cancel" @click="showLogoutModal = false" title="Отменить">
          <i class="fas fa-times"></i>
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import AuthService from '../services/AuthService';
import defaultAvatarImage from '../assets/default-avatar.png'; // Добавьте дефолтную аватарку

export default defineComponent({
  name: 'NavbarHeader',
  setup() {
    const route = useRoute();
    const router = useRouter();
    const showNotifications = ref(false);
    const showSettings = ref(false);
    const showLogoutModal = ref(false);
    const userAvatar = ref(''); // Здесь должен быть путь к аватарке пользователя
    const defaultAvatar = defaultAvatarImage;

    const currentPageTitle = computed(() => {
      switch (route.path) {
        case '/home':
          return 'Главная';
        case '/messenger':
        case '/messenger/chat':
          return 'Мессенджер';
        case '/messenger/match':
          return 'Match';
        case '/profile':
          return 'Мой профиль';
        default:
          return '';
      }
    });

    const toggleNotifications = () => {
      showNotifications.value = !showNotifications.value;
      showSettings.value = false; // Закрываем другое меню
    };

    const toggleSettings = () => {
      showSettings.value = !showSettings.value;
      showNotifications.value = false; // Закрываем другое меню
    };

    const confirmLogout = () => {
      showSettings.value = false;
      showLogoutModal.value = true;
    };

    const handleLogout = async () => {
      try {
        showLogoutModal.value = false;
        await AuthService.logout();
        router.push('/login');
      } catch (error) {
        console.error('Ошибка при выходе:', error);
      }
    };

    return {
      showNotifications,
      showSettings,
      toggleNotifications,
      toggleSettings,
      userAvatar,
      defaultAvatar,
      showLogoutModal,
      confirmLogout,
      handleLogout,
      currentPageTitle,
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
  padding: 20px 2em;
  border-bottom: 1px solid #e0e0e0;

  background-color: #fff;
  color: white;
  z-index: 1000;
}


.navbar-left {
  display: flex;
  align-items: center;
  margin-left: 5em;
}

.logo-ellipse {
  background-color: var(--primary-color);
  padding: 8px 30px;
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

.settings-button {
  width: 100%;
  padding: 10px;
  text-align: left;
  background: none;
  border: none;
  color: var(--color-text);
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 10px;

}

.settings-button:hover {
  background-color: var(--color-secondary);
  border-radius: var(--border-radius);
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;
}

.modal-content {
  background-color: white;
  padding: 2rem;
  border-radius: 10px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  min-width: 300px;
}

.modal-title {
  font-size: 1.2rem;
  margin-bottom: 1rem;
  color: var(--color-text);
}

.modal-text {
  color: var(--color-text);
  margin-bottom: 1.5rem;
}

.modal-buttons {
  display: flex;
  justify-content: center;
  gap: 1rem;
}

.modal-button {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  border: none;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.2s ease;
}

.modal-button i {
  font-size: 1.2rem;
}

.confirm {
  background-color: #4CAF50;
  color: white;
}

.cancel {
  background-color: #f44336;
  color: white;
}

.modal-button:hover {
  transform: scale(1.1);
  opacity: 0.9;
}

.navbar-center {
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
}

.page-title {
  color: var(--primary-color);
  font-size: 1.2rem;
  font-weight: bold;
  margin: 0;
  color: #A68136;
    font-size: 47px;
    text-align: center;
}
</style>
