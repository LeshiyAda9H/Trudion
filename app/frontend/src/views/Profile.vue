<template>
  <header>
    <Navbar />
  </header>

  <LoadingSpinner v-if="isLoading" />

  <div v-else class="prof-container">
    <div class="profile-header">
      <div class="gender-icon" :class="{ active: profileStore.temporaryData.gender === 'male' }"
        @click="profileStore.temporaryData.gender = 'male'">
        <i class="fas fa-mars"></i>
      </div>

      <div class="avatar-container">
        <div class="avatar">
          <img :src="defaultAvatar" alt="avatar" />
          <!-- <i class="icon-user">  </i> -->
          <div class="add-photo">+</div>
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

    return {
      profileStore,
      updateProfile,
      isLoading,
      defaultAvatar
    };
  },
});
</script>

<style scoped>
.prof-container {
  margin-top: 5.5em;
  position: fixed;
  left: 50%;
  transform: translateX(-50%);

}

.interests-section {
  margin: 20px 0;
}

.interests-section h3 {
  margin-bottom: 10px;
}

.save-button-container {
  position: fixed;
  margin: 41em auto;
  align-items: center;
  justify-content: center;
  display: flex;
  width: 100%;
  margin-bottom: 1em;
}

.save-button {
  padding: 15px 20px;

  background-color: #A68136;

  font-size: 16px;
  font-weight: bold;

  text-align: center;
  color: #fff;


  width: 70%;

  display: grid;
  place-items: center;


}


.profile-header {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 30px;
  margin-bottom: 20px;
}

.avatar-container {
  /* Существующие стили для avatar-container */
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
}

.gender-icon i {
  filter: invert(1);
}

.gender-icon:hover {
  background-color: var(--secondary-color);
  color: white;
}

.gender-icon.active {
  /* background-color: var(--primary-color); */
  /* border-color: var(--primary-color); */
  color: white;
  opacity: 1;
  box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);
}

.gender-icon {
  /* background-color: var(--primary-color); */
  /* border-color: var(--primary-color); */
  opacity: 0.1;
  color: white;
}
</style>
