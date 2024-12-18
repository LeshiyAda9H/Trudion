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

      <div class="interests-section">
        <h3>Ваши интересы</h3>
        <InterestSelector
          v-model="profileStore.temporaryData.label"
          placeholder="Выберите ваши интересы"
        />
      </div>
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
import { useRouter } from 'vue-router';
import InterestSelector from '../components/InterestSelector.vue';

export default defineComponent({
  name: 'ProfilePage',
  components: {
    Navbar,
    LoadingSpinner,
    InterestSelector
  },

  setup() {
    const profileStore = useProfileStore();
    const isLoading = ref(false);
    const router = useRouter();

    const fetchProfileData = async () => {
      isLoading.value = true;
      try {
        const data = await AuthService.getProfile();
        const labels = data.label ?
          (Array.isArray(data.label) ? data.label : [data.label]) :
          [];

        profileStore.temporaryData = {
          username: data.username,
          gender: data.gender,
          biography: data.biography,
          label: labels,
          online_status: data.online_status,
        };

        console.log('Загруженные интересы:', labels);
      }
      catch (error) {
        console.error("Ошибка при получении данных профиля:", error);
        alert("Не удалось загрузить данные профиля");
      }
      finally {
        isLoading.value = false;
      }
    };

    onMounted(async () => {
      const token = localStorage.getItem('token');
      if (!token) {
        router.push('/login');
        return;
      }
      await fetchProfileData();
    });

    const updateProfile = async () => {
      if (!validateData()) return;

      try {
        await AuthService.updateProfile({
          username: profileStore.temporaryData.username,
          gender: profileStore.temporaryData.gender,
          biography: profileStore.temporaryData.biography,
          label: profileStore.temporaryData.label,
          online_status: profileStore.temporaryData.online_status,
        });

        profileStore.completeUpdateProfile();
        alert("Данные обновлены!");
      }
      catch (error) {
        console.error("Ошибка при сохранении профиля:", error);
        alert("Не удалось обновить профиль :(");
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

    return {
      profileStore,
      updateProfile,
      isLoading
    };
  },
});
</script>

<style scoped>
.interests-section {
  margin: 20px 0;
}

.interests-section h3 {
  margin-bottom: 10px;
}
</style>
