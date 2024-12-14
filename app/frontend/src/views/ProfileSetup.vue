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

        <select class="select" v-model="userStore.registrationData.label">
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
import { defineComponent } from 'vue';
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
          label: userStore.registrationData.label,
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
    };
  },
});
</script>
