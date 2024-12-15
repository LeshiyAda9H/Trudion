  <template>


    <body>
      <h1 class="title">Завершите регистрацию</h1>

      <div class="prof-container">
        <div class="avatar-container">
          <div class="avatar">
            <i class="icon-user"></i>
            <div class="add-photo">+</div>
          </div>
        </div>

        <input type="text" placeholder="Никнейм" class="input" v-model="userStore.registrationData.username" />

        <select class="select" v-model="userStore.registrationData.gender">
          <option value="" disabled>Выбери свой пол</option>
          <option value="male">Мужчина</option>
          <option value="female">Женщина</option>
          <option value="other">Абоба</option>
        </select>

        <textarea placeholder="О себе" class="form-textarea" v-model="userStore.registrationData.biography"></textarea>

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
import { defineComponent, ref } from 'vue';
import { useUserStore } from '../stores/UserStore';
import { useRouter } from 'vue-router';
import AuthService from '../services/AuthService';
import '../assets/css/profile.css'
import StorageService from '../services/StorageService' // Подключаем StorageService


export default defineComponent({
  name: 'ProfileSetupPage',
  //components: {  },

  setup() {
    const userStore = useUserStore();
    const router = useRouter(); // Инициализация роутера

    const selectedlabel = ref<string[]>([]);

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

    const toggleLabel = (value: string) => {
      const index = selectedlabel.value.indexOf(value);
      if (index === -1) {
        selectedlabel.value.push(value);
      } else {
        selectedlabel.value.splice(index, 1);
      }
      userStore.registrationData.label = selectedlabel.value;
    };

    const validateData = (): boolean => {
      const { username, gender } = userStore.registrationData;
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
        await AuthService.completeProfile({
          username: userStore.registrationData.username,
          email: userStore.registrationData.email,
          password: userStore.registrationData.password,
          gender: userStore.registrationData.gender,
          biography: userStore.registrationData.biography,
          label: selectedlabel.value,
        });

        const authedUser = await AuthService.login({
          email: userStore.registrationData.email,
          password: userStore.registrationData.password,
        })

        StorageService.setToken(authedUser.access_token) // Сохранение токена

        userStore.completeRegistration(); // Очищаем временные данные

        // Проверка токена через сервер
        // const isTokenValid = await AuthService.verifyToken()
        // if (!isTokenValid) {
        //   throw new Error('Недействительный токен.')
        // }


        alert("Регистрация завершена!");

        router.push("/home");
      }
      catch (error) {
        console.error("Ошибка при сохранении профиля:", error);
        alert("Не удалось завершить регистрацию.");
      }
    };

    return {
      userStore,
      updateProfile,
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
