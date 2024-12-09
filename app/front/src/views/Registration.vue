<template>
  <div class="container">
    <p>Регистрация</p>

    <img :src="imagePath" class="image-ava" />

    <InputRegistration
      :writeEmail="writeEmail"
      :writePass="writePass"
      :writeConfirmPass="writeConfirmPass"
      :error="error"
    />

    <button class="button" @click="sendData">Зарегистрироваться</button>

    <div class="footer-text">
      Уже есть аккаунт? <router-link to="/login" class="link">Войти</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import myImage from '../assets/trudion.png'
import InputRegistration from '../components/InputRegistration.vue'
import '../assets/css/authentication.css'
import { useUserStore } from '../stores/UserStore'; // Импортируем хранилище Pinia
import AuthService from '../services/AuthService'; // Импортируем AuthService для регистрации

export default defineComponent({
  name: 'RegistrationPage',

  components: { InputRegistration },

  computed: {
    imagePath(): string {
      return myImage // Путь к изображению
    },
  },

  data() {
    return {
      error: '' as string, // Ошибка
      userNickname: '' as string, // Никнейм
      userEmail: '' as string, // Электронная почта
      userPass: '' as string, // Пароль
      confirmPass: '' as string, // Подтверждение пароля
      userGender: '' as string, // Пол
    }
  },

  methods: {
    writeEmail(text_email: string): void {
      this.userEmail = text_email
    },
    writePass(text_pass: string): void {
      this.userPass = text_pass
    },
    writeConfirmPass(text_confirmPass: string): void {
      this.confirmPass = text_confirmPass
    },
  },
  setup() {
    const userStore = useUserStore();

    const sendData = async () => {
      // Проверяем обязательные поля
      if (!userStore.registrationData.email || !userStore.registrationData.password) {
        alert('Пожалуйста, заполните email и пароль');
        return;
      }

      // Сохраняем данные в хранилище
      userStore.setRegistrationData({
        email: userStore.registrationData.email,
        password: userStore.registrationData.password,
        nickname: userStore.registrationData.nickname,
        gender: userStore.registrationData.gender,
      });

      try {
        // Регистрируем пользователя через AuthService
        await AuthService.register({
          email: userStore.registrationData.email,
          password: userStore.registrationData.password,
          nickname: userStore.registrationData.nickname,
          gender: userStore.registrationData.gender,
        });

        // Переходим на страницу профиля, где пользователь сможет обновить информацию
        window.location.href = '/profile'; // Переход на страницу профиля для дальнейшего обновления данных
      }
      catch (error) {
        console.error('Ошибка при регистрации:', error);
        alert('Ошибка при регистрации');
      }
    };

    return {
      userStore,
      sendData,
    };
  },

})
</script>

