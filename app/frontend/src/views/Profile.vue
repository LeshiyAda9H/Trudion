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

      <div class="label-container">
          <h3>Выберите ваши интересы:</h3>
          <div class="label-grid">
            <div
              v-for="label in availablelabel"
              :key="label.id"
              :class="['label-item', { 'selected': selectedlabel.includes(label.value) }]"
              @click="toggleLabel(label.value)"
            >
              {{ label.name }}
            </div>
          </div>
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

export default defineComponent({
  name: 'ProfilePage',
  components: {
    Navbar,
    LoadingSpinner
  },

  setup() {
    const profileStore = useProfileStore();
    const isLoading = ref(false);

    const availablelabel = [
      { id: 0, name: 'ExampleLabel', value: 'ExampleLabel' },
      { id: 1, name: 'Программирование', value: 'programming' },
      { id: 2, name: 'Иностранные языки', value: 'languages' },
      { id: 3, name: 'Спорт', value: 'sports' },
      { id: 4, name: 'Кулинария', value: 'cooking' },
      { id: 5, name: 'Музыка', value: 'music' },
      { id: 6, name: 'Искусство', value: 'art' },
      { id: 7, name: 'Фотография', value: 'photography' },
      { id: 8, name: 'Путешествия', value: 'traveling' },
      { id: 9, name: 'Наука', value: 'science' },
      { id: 10, name: 'Литература', value: 'literature' }
    ];

    const selectedlabel = ref<string[]>([]);

    const toggleLabel = (value: string) => {
      const index = selectedlabel.value.indexOf(value);
      if (index === -1) {
        selectedlabel.value.push(value);
      } else {
        selectedlabel.value.splice(index, 1);
      }
      profileStore.temporaryData.label = selectedlabel.value;
    };

    const fetchProfileData = async () => {
      isLoading.value = true;
      try {
        const data = await AuthService.getProfile();
        profileStore.temporaryData = {
          username: data.username,
          gender: data.gender,
          biography: data.biography,
          label: data.label,
          online_status: data.online_status,
        };
        selectedlabel.value = data.label || [];
      }
      catch (error) {
        console.error("Ошибка при получении данных профиля:", error);
        alert("Не удалось загрузить данные профиля");
      }
      finally {
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

    // Загружаем данные при монтировании компонента
    onMounted(() => {
      fetchProfileData();
    });

    return {
      profileStore,
      updateProfile,
      isLoading,
      availablelabel,
      selectedlabel,
      toggleLabel
    };
  },
});
</script>

<style scoped>
.label-container {
  margin: 20px 0;
}

.label-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 10px;
  padding: 15px;
}

.label-item {
  padding: 10px 15px;
  border: 1px solid #ddd;
  border-radius: 20px;
  cursor: pointer;
  text-align: center;
  transition: all 0.3s ease;
}

.label-item:hover {
  background-color: #f0f0f0;
}

.label-item.selected {
  background-color: #007bff;
  color: white;
  border-color: #0056b3;
}
</style>
