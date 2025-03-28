<template>
<div class = "screen">
  <div class="auth-container">
    <p class="auth-title">Авторизация</p>
    <div class="image-container">
      <img :src="imagePath" class="image-ava" />
    </div>
    <InputAuthorization
    :writeEmail="writeEmail"
    :writePass="writePass"
    :error="error"
    />

    <button class="auth-button" @click="sendData">Войти</button>

    <div class="footer-text">
      Нет аккаунта?
      <router-link to="/register" class="link">Зарегистрироваться</router-link>
      <br />
      <a class="link" href="#">Восстановить пароль</a>
    </div>
  </div>
</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import myImage from '../assets/trudion.svg'
import InputAuthorization from '../components/InputAuthorization.vue'
import AuthService from '../services/AuthService' // Подключаем AuthService
import StorageService from '../services/StorageService' // Подключаем StorageService
import '../assets/css/authentication.css'

export default defineComponent({
  name: 'AuthorizationPage',
  components: { InputAuthorization },
  computed: {
    imagePath(): string {
      return myImage // Путь к изображению
    },
  },
  data() {
    return {
      error: '' as string, // Строка ошибки
      userEmail: '' as string, // Электронная почта
      userPass: '' as string, // Пароль
    }
  },
  methods: {
    writeEmail(text_email: string): void {
      this.userEmail = text_email
      if (this.error === 'not-valid-email-and-password' || this.error === 'not-valid-email') {
        this.error = '' // Сброс ошибки, если пользователь начинает ввод
      }
    },
    writePass(text_pass: string): void {
      this.userPass = text_pass
      if (this.error === 'not-valid-email-and-password' || this.error === 'not-valid-password') {
        this.error = '' // Сброс ошибки, если пользователь начинает ввод
      }
    },

    validateData(): boolean {
      if (this.userEmail === '' && this.userPass === '') {
        this.error = 'not-valid-email-and-password';
        return false;
      }
      if (this.userEmail === '') {
        this.error = 'not-valid-email';
        return false;
      }
      if (this.userPass === '') {
        this.error = 'not-valid-password';
        return false;
      }
      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!re.test(this.userEmail)) {
        this.error = 'not-valid-email';
        return false;
      }
      return true;
    },

    async sendData(): Promise<void> {
      try {
        if (!this.validateData()) return;

        console.log('Отправляем данные:', {
          email: this.userEmail,
          password: this.userPass,
        });

        const authedUser = await AuthService.login({
          email: this.userEmail,
          password: this.userPass,
        })

        console.log('Ответ от AuthService:', authedUser);

        if (!authedUser || !authedUser.token) {
          throw new Error('Не получен токен авторизации');
        }

        StorageService.setToken(authedUser.token);

        const savedToken = localStorage.getItem('token');
        console.log('Проверка сохраненного токена:', savedToken);

        if (!savedToken) {
          throw new Error('Токен не был сохранен');
        }

        this.$router.push('/home')
      } catch (error) {
        console.error('Ошибка при авторизации:', error);
        this.error = 'Ошибка при авторизации';
      }
    },
  },
})
</script>

<style scoped>

.input {
  width: 20em;
  margin: 5px 0;
}
.input-error{
  width: 20em;
  margin: 5px 0;
  color: red;
}

</style>
