<template>
  <body>
    <h1 class="title">Завершите регистрацию</h1>
    <div class="prof-container">
      <div class="profile-header">
        <div class="gender-icon"
             :class="{ active: userStore.registrationData.gender === 'male' }"
             @click="userStore.registrationData.gender = 'male'"
        >
          <i class="fas fa-mars"></i>
        </div>

        <div class="avatar-container">
          <div class="avatar">
            <img :src="defaultAvatar" alt="avatar" />
            <div class="add-photo">+</div>
          </div>
        </div>

        <div class="gender-icon"
             :class="{ active: userStore.registrationData.gender === 'female' }"
             @click="userStore.registrationData.gender = 'female'"
        >
          <i class="fas fa-venus"></i>
        </div>
      </div>

      <input type="text" placeholder="Никнейм" class="input" v-model="userStore.registrationData.username" />
      <textarea placeholder="О себе" class="form-textarea" v-model="userStore.registrationData.biography"></textarea>

      <div class="interests-section">
        <h3>Ваши интересы</h3>
        <InterestSelector
          v-model="userStore.registrationData.label"
          placeholder="Выберите ваши интересы"
          direction="up"
        />
      </div>
    </div>

    <div class="save-button-container">
      <button class="save-button" @click="updateProfile">Сохранить изменения</button>
    </div>
  </body>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUserStore } from '../stores/UserStore';
import { useRouter } from 'vue-router';
import AuthService from '../services/AuthService';
import '../assets/css/profile.css'
import StorageService from '../services/StorageService'
import InterestSelector from '../components/InterestSelector.vue';
import defaultAvatar from '../assets/default-avatar.png'

const userStore = useUserStore();
const router = useRouter();
const isLoading = ref(false);

if (!userStore.registrationData.label) {
  userStore.registrationData.label = [];
}

const validateData = (): boolean => {
  const { username, gender } = userStore.registrationData;
  if (!username || !gender) {
    alert("Заполните никнейм и выберите пол.");
    return false;
  }
  return true;
};

const updateProfile = async () => {
  if (!validateData()) return;

  isLoading.value = true;
  try {
    await AuthService.completeProfile({
      username: userStore.registrationData.username,
      email: userStore.registrationData.email,
      password: userStore.registrationData.password,
      gender: userStore.registrationData.gender,
      biography: userStore.registrationData.biography,
      label: userStore.registrationData.label,
    });

    const authedUser = await AuthService.login({
      email: userStore.registrationData.email,
      password: userStore.registrationData.password,
    });

    StorageService.setToken(authedUser.token);
    userStore.completeRegistration();
    router.push("/home");
  } catch (error) {
    console.error("Ошибка при сохранении профиля:", error);
    alert("Не удалось завершить регистрацию.");
  } finally {
    isLoading.value = false;
  }
};
</script>

<style scoped>

.title {
  font-size: 47px;
  text-align: center;
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  margin-top: 0.5em;

}

.prof-container {
  position: fixed;
  left: 50%;
  transform: translateX(-50%);
  margin-top: 6em;

}

.profile-header {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 30px;
  margin-bottom: 20px;
}

.avatar img {
  width: 10em;
  height: 10em;
  border-radius: 50%;
  object-fit: cover;
}

.gender-icon {
  width: 50px;
  height: 50px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  border: 2px solid var(--primary-color);
  transition: all 0.3s ease;
  font-size: 24px;
  color: var(--secondary-color);
  margin-top: 5em;
  opacity: 0.1;
}

.gender-icon i {
  filter: invert(1);
}

.gender-icon:hover {
  background-color: var(--secondary-color);
  color: white;
}

.gender-icon.active {
  color: white;
  opacity: 1;
  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);
}

.save-button-container {
  position: fixed;
  margin: 41em auto;
  align-items: center;
  justify-content: center;
  display: flex;
  width: 100%;
}

.interests-section {
  margin: 20px 0;
}

.interests-section h3 {
  margin-bottom: 10px;
}
</style>
