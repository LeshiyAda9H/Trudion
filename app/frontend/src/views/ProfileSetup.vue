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

        <div class="interests-section">
          <h3>Выберите ваши интересы</h3>
          <InterestSelector v-model="userStore.registrationData.label" placeholder="Выберите ваши интересы" />
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
import InterestSelector from '../components/InterestSelector.vue';


export default defineComponent({
  name: 'ProfileSetupPage',
  components: {
    InterestSelector
  },

  setup() {
    const userStore = useUserStore();
    const router = useRouter();
    const isLoading = ref(false); // Добавим индикатор загрузки

    // Инициализируем массив меток, если его нет
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
      if (!validateData()) {
        return;
      }

      isLoading.value = true;
      try {
        console.log('Отправка данных для завершения регистрации:', userStore.registrationData);

        // Завершение регистрации
        await AuthService.completeProfile({
          username: userStore.registrationData.username,
          email: userStore.registrationData.email,
          password: userStore.registrationData.password,
          gender: userStore.registrationData.gender,
          biography: userStore.registrationData.biography,
          label: userStore.registrationData.label,
        });
        

        // Автоматический вход после регистрации
        const authedUser = await AuthService.login({
          email: userStore.registrationData.email,
          password: userStore.registrationData.password,
        });

        console.log('Успешная авторизация после регистрации');

        // Сохранение токена
        StorageService.setToken(authedUser.token);

        // Проверка сохранения токена
        const savedToken = localStorage.getItem('token');
        if (!savedToken) {
          throw new Error('Ошибка сохранения токена');
        }

        // Очистка временных данных
        userStore.completeRegistration();

        alert("Регистрация завершена!");
        router.push("/home");
      }
      catch (error) {
        console.error("Ошибка при сохранении профиля:", error);
        alert("Не удалось завершить регистрацию.");
      }
      finally {
        isLoading.value = false;
      }
    };

    return {
      userStore,
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
