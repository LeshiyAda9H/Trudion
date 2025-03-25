<template>

<div class = "screen">
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
          <div class="add-photo">
            <i class="fas fa-camera"></i>
          </div>
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

    <div class="save-button-container">
    <button class="save-button" @click="updateProfile">Сохранить изменения</button>
    </div>
  </div>



</div>
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
.screen{
    height: 95vh;
    display: grid;
    justify-content: center;
    transform: translateY(0vh);
  }

  .title {
    padding: 2.5vh 0;
    color: #A68136;
    font-size: 47px;
    text-align: center;
    place-content: center;
  }
</style>
