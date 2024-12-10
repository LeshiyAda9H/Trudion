<template>
  <div class="auth-container">
    <p>Регистрация</p>

    <img :src="imagePath" class="image-ava" />

    <InputRegistration :writeEmail="writeEmail" :writePass="writePass" :writeConfirmPass="writeConfirmPass"
      :error="error" />

    <button class="auth-button" @click="sendData">Зарегистрироваться</button>

    <div class="footer-text">
      Уже есть аккаунт? <router-link to="/login" class="link">Войти</router-link>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import myImage from '../assets/trudion.png'
import InputRegistration from '../components/InputRegistration.vue'
import '../assets/css/authentication.css'
import { useUserStore } from '../stores/UserStore'; // Импортируем хранилище Pinia
import AuthService from '../services/AuthService'; // Импортируем AuthService для регистрации
import { useRouter } from "vue-router"; // Используем Composition API для роутера

export default defineComponent({
  name: 'RegistrationPage',

  components: { InputRegistration },

  setup() {
    const router = useRouter(); // Инициализация роутера
    const userStore = useUserStore(); // Инициализация хранилища Pinia

    const imagePath = myImage;
    const error = ref(''); // Ошибка
    const userEmail = ref(''); // Электронная почта
    const userPass = ref(''); // Пароль
    const confirmPass = ref(''); // Подтверждение пароля

    const writeEmail = (text_email: string): void => {
      userEmail.value = text_email;
    };

    const writePass = (text_pass: string): void => {
      userPass.value = text_pass;
    };

    const writeConfirmPass = (text_confirmPass: string): void => {
      confirmPass.value = text_confirmPass;
    };

    const sendData = async () => {
      try {
        // Валидация данных
        if (userEmail.value === '' && (userPass.value === '' || confirmPass.value === '')) {
          error.value = 'not-valid-email-and-passwords';
          return;
        }
        if (userEmail.value === '') {
          error.value = 'not-valid-email';
          return;
        }

        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!re.test(userEmail.value)) {
          error.value = 'not-valid-email';
          return;
        }
        if (userPass.value === '' || confirmPass.value === '') {
          error.value = 'passwords-dont-match';
          return;
        }

        if (userPass.value !== confirmPass.value) {
          error.value = 'passwords-dont-match';
          return;
        }

        const isEmailAvailable = await AuthService.verifyEmail(userEmail.value);
        if (!isEmailAvailable) {
          error.value = "Этот email уже зарегистрирован.";
          return;
        }

        userStore.setRegistrationData({
          email: userEmail.value,
          password: userPass.value,
        });

        // Перенаправление на страницу профиля
        router.push("/profile-setup");
      } catch (err) {
        console.error("Ошибка при проверке email:", err);
        error.value = "Не удалось проверить email. Попробуйте позже.";
      }
    };

    return {
      imagePath,
      error,
      userEmail,
      userPass,
      confirmPass,
      writeEmail,
      writePass,
      writeConfirmPass,
      sendData
    };
  },
});
</script>
