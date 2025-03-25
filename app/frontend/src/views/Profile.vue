<template>
  <header>
    <Navbar />
  </header>

  <LoadingSpinner v-if="isLoading" />
  <div v-else class = "screen">
  <div class="prof-container">
    <div class="profile-header">
      <div class="gender-icon" :class="{ active: profileStore.temporaryData.gender === 'male' }"
        @click="profileStore.temporaryData.gender = 'male'">
        <i class="fas fa-mars"></i>
      </div>

      <div class="avatar-container">
        <div class="avatar">
          <img :src="avatarSource" alt="avatar" />
          <div class="add-photo" @click="triggerFileInput">
            <i class="fas fa-camera"></i>
          </div>
        </div>
        <div v-if="isUploading" class="upload-indicator">
          Загрузка...
        </div>
      </div>

      <div class="gender-icon" :class="{ active: profileStore.temporaryData.gender === 'female' }"
        @click="profileStore.temporaryData.gender = 'female'">
        <i class="fas fa-venus"></i>
      </div>
    </div>

    <input type="text" placeholder="Никнейм" class="input" v-model="profileStore.temporaryData.username" />

    <textarea placeholder="О себе" class="form-textarea" v-model="profileStore.temporaryData.biography"></textarea>

    <div class="interests-section">
      <h3>Ваши интересы</h3>
      <InterestSelector v-model="profileStore.temporaryData.label" placeholder="Выберите ваши интересы"
        direction="up" />
    </div>

    <div class="save-button-container">
      <button class="save-button" @click="updateProfile">Сохранить изменения</button>
    </div>
  </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, computed } from 'vue';
import Navbar from '../components/Navbar.vue';
import LoadingSpinner from '../components/LoadingSpinner.vue';
import { useProfileStore } from '../stores/ProfileStore';
import AuthService from '../services/AuthService';
import '../assets/css/profile.css'
import { useRouter } from 'vue-router';
import InterestSelector from '../components/InterestSelector.vue';
import defaultAvatar from '../assets/default-avatar.png'

export default defineComponent({
  name: 'ProfilePage',
  components: {
    Navbar,
    LoadingSpinner,
    InterestSelector
  },

  setup() {
    const profileStore = useProfileStore();
    const isLoading = ref(true);
    const router = useRouter();

    const fetchProfileData = async () => {
      try {
        const data = await AuthService.getProfile();
        const labels = data.label ?
          (Array.isArray(data.label) ? data.label : [data.label]) :
          [];

        profileStore.temporaryData = {
          user_id: data.user_id,
          username: data.username,
          gender: data.gender,
          biography: data.biography,
          label: labels,
          online_status: data.online_status,
        };
      }
      catch (error) {
        console.error("Ошибка при получении данных профиля:", error);
        router.push('/login');
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
        const changedFields: Partial<typeof profileStore.temporaryData> = {};
        const originalData = profileStore.profileData;

        if (profileStore.temporaryData.username !== originalData.username) {
          changedFields.username = profileStore.temporaryData.username;
        }
        if (profileStore.temporaryData.gender !== originalData.gender) {
          changedFields.gender = profileStore.temporaryData.gender;
        }
        if (profileStore.temporaryData.biography !== originalData.biography) {
          changedFields.biography = profileStore.temporaryData.biography;
        }
        if (profileStore.temporaryData.label && originalData.label &&
          !arraysEqual(profileStore.temporaryData.label, originalData.label)) {
          changedFields.label = profileStore.temporaryData.label;
        }

        console.log('Измененные поля:', changedFields);

        if (Object.keys(changedFields).length > 0) {
          await AuthService.updateProfile(changedFields);
          profileStore.completeUpdateProfile();
          alert("Данные обновлены!");
        } else {
          alert("Нет изменений для сохранения");
        }
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

    const arraysEqual = (a: string[], b: string[]) => {
      if (a.length !== b.length) return false;
      return a.every((val, index) => val === b[index]);
    };

    const fileInputRef = ref<HTMLInputElement | null>(null);
    const isUploading = ref(false);

    const avatarSource = computed(() => {
      return profileStore.temporaryData.avatar || defaultAvatar;
    });

    const triggerFileInput = () => {
      const input = fileInputRef.value || document.createElement('input');
      input.type = 'file';
      input.accept = 'image/*';
      input.onchange = handleFileUpload;
      input.click();
    };

    const handleFileUpload = async (event: Event) => {
      const target = event.target as HTMLInputElement;
      const file = target.files?.[0];

      if (!file) return;

      console.log('File selected:', file);

      try {
        isUploading.value = true;
        const response = await AuthService.uploadAvatar(file);
        console.log('Upload response:', response);

        profileStore.temporaryData.avatar = response.path;
      } catch (error) {
        console.error('Ошибка при загрузке файла:', error);
        alert('Не удалось загрузить изображение');
      } finally {
        isUploading.value = false;
      }
    };

    return {
      profileStore,
      updateProfile,
      isLoading,
      defaultAvatar,
      avatarSource,
      triggerFileInput,
      handleFileUpload,
      isUploading
    };
  },
});
</script>
