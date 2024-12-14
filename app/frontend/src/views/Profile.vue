<template>
  <header>
    <Navbar />
  </header>

  <body>
    <LoadingSpinner v-if="isLoading" />

    <h1 class="title">Мой профиль</h1>

    <div class="prof-container">
      <div class="avatar-container">
        <div class="avatar">
          <i class="icon-user"></i>
          <div class="add-photo">+</div>
        </div>
      </div>

      <input type="text" placeholder="Никнейм" class="input" v-model="profileStore.temporaryData.username" />

      <select class="select" v-model="profileStore.temporaryData.gender">
        <option value="" disabled>Выбери свой пол</option>
        <option value="male">Мужчина</option>
        <option value="female">Женщина</option>
        <option value="other">Абоба</option>
      </select>

      <textarea placeholder="О себе" class="form-textarea" v-model="profileStore.temporaryData.biography"></textarea>

      <select class="select" v-model="profileStore.temporaryData.label">
        <option value="" disabled>Навыки</option>
        <option value="skill1">Навык 1</option>
        <option value="skill2">Навык 2</option>
        <option value="skill3">Навык 3</option>
      </select>
    </div>

    <button class="save-button" @click="updateProfile">Сохранить изменения</button>
  </body>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import Navbar from '../components/Navbar.vue';
import LoadingSpinner from '../components/LoadingSpinner.vue';
import { useProfileStore } from '../stores/ProfileStore';
import AuthService from '../services/AuthService';
import '../assets/css/profile.css'

export default defineComponent({
  name: 'ProfilePage',
  components: {
    Navbar,
    LoadingSpinner
  },

  setup() {
    const profileStore = useProfileStore();
    const isLoading = ref(false);

    const fetchProfileData = async () => {
      isLoading.value = true;
      try {
        const data = await AuthService.getProfile();
        profileStore.temporaryData = {
          username: data.username,
          gender: data.gender,
          biography: data.biography,
          label: data.label,
        };
      } catch (error) {
        console.error("Ошибка при получении данных профиля:", error);
        alert("Не удалось загрузить данные профиля");
      } finally {
        isLoading.value = false;
      }
    };

    const validateData = (): boolean => {
      const { username, gender } = profileStore.temporaryData;
      if (!username || !gender) {
        alert("Заполните никнейм и выберите пол.");
        return false;
      }
      return true;
    };

    const updateProfile = async () => {
      if (!validateData()) {
        return;
      }

      try {
        await AuthService.updateProfile({
          username: profileStore.temporaryData.username,
          gender: profileStore.temporaryData.gender,
          biography: profileStore.temporaryData.biography,
          label: profileStore.temporaryData.label,
        });

        profileStore.completeUpdateProfile();
        alert("Данные обновлены!");
      }
      catch (error) {
        console.error("Ошибка при сохранении профиля:", error);
        alert("Не удалось обновить профиль :(");
      }
    };

    // Загружаем данные при монтировании компонента
    onMounted(() => {
      fetchProfileData();
    });

    return {
      profileStore,
      updateProfile,
      isLoading
    };
  },
});
</script>
