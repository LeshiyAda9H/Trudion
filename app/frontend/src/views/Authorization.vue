<template>


  <div class="auth-container">
    <p class="auth-title">Авторизация</p>

    <img :src="imagePath" class="image-ava" />

    <InputAuthorization :writeEmail="writeEmail" :writePass="writePass" :error="error" />

    <button class="auth-button" @click="sendData">Войти</button>

    <div class="footer-text">
      Нет аккаунта?
      <router-link to="/register" class="link">Зарегистрироваться</router-link>
      <br />
      <a class="link" href="#">Восстановить пароль</a>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import myImage from '../assets/trudion.svg'
import InputAuthorization from '../components/InputAuthorization.vue'
import AuthService from '../services/AuthService' // Подключаем AuthService
import StorageService from '../services/StorageService' // Подключаем StorageService
import { AxiosError } from 'axios'
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
        // Валидация данных
        if (!this.validateData()) {
          return;
        }

        // Авторизация пользователя через AuthService
        const authedUser = await AuthService.login({
          email: this.userEmail,
          password: this.userPass,
        })

        // Проверка токена через сервер
        const isTokenValid = await AuthService.verifyToken()
        if (!isTokenValid) {
          throw new Error('Недействительный токен.')
        }

        // Успешная авторизация
        StorageService.setToken(authedUser.access_token) // Сохранение токена
        this.$router.push('/home') // Перенаправляем на главную страницу

      }
      catch (error: unknown) {
        if (error instanceof AxiosError) {
          if (error.response?.status === 401) {
            this.error = 'Неверный логин или пароль.'
          }
          else {
            this.error = 'Ошибка на сервере. Пожалуйста, попробуйте позже.'
          }
        }
        else {
          console.error(error)
          this.error = 'Неизвестная ошибка.'
        }
      }
    },
  },
})
</script>


<style>


</style>
