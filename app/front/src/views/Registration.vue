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
import { defineComponent } from 'vue'
import myImage from '../assets/trudion.png'
import InputRegistration from '../components/InputRegistration.vue'
import '../assets/css/authentication.css'
import { useUserStore } from '../stores/UserStore'; // Импортируем хранилище Pinia
import AuthService from '../services/AuthService'; // Импортируем AuthService для регистрации
import { useRouter } from "vue-router";

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
      userEmail: '' as string, // Электронная почта
      userPass: '' as string, // Пароль
      confirmPass: '' as string, // Подтверждение пароля
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


    async sendData() {
      try {
        //Валидация данных
        if (this.userEmail === '' && (this.userPass === '' || this.confirmPass === '')) {
          this.error = 'not-valid-email-and-passwords'
          return
        }
        if (this.userEmail === '') {
          this.error = 'not-valid-email'
          return
        }
        const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
        if (!re.test(this.userEmail)) {
          this.error = 'not-valid-email'
          return
        }
        if (this.userPass === '' || this.confirmPass === '') {
          this.error = 'passwords-dont-match'
          return
        }
        // Проверка на совпадение паролей
        if (this.userPass !== this.confirmPass) {
          this.error = 'passwords-dont-match'
          return
        }

        const isEmailAvailable = await AuthService.verifyEmail(this.userEmail);
        if (!isEmailAvailable) {
          this.error = "Этот email уже зарегистрирован.";
          return;
        }

        const userStore = useUserStore();
        userStore.setRegistrationData({
          email: this.userEmail,
          password: this.userPass,
        });

        // Перенаправляем пользователя на страницу профиля
        const router = useRouter();
        router.push("/profile");
      }
      catch (error) {
        console.error("Ошибка при проверке email:", error);
        this.error = "Не удалось проверить email. Попробуйте позже.";
      }
    },
  },


})
</script>
