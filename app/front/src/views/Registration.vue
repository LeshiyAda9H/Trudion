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
//import AuthService from '../services/AuthService'
import '../assets/css/authentication.css'
import { useUserStore } from '../stores/UserStore'; // Импортируем хранилище Pinia

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
    // async sendData(): Promise<void> {
    //   try {
    //     //Валидация данных
    //     if (!this.userEmail || !this.userPass || !this.confirmPass || !this.userNickname || !this.userGender) {
    //       this.error = 'Заполните все поля';
    //       return;
    //     }

    //     if (this.userEmail === '' && (this.userPass === '' || this.confirmPass === '')) {
    //       this.error = 'not-valid-email-and-passwords'
    //       return
    //     }

    //     if (this.userEmail === '') {
    //       this.error = 'not-valid-email'
    //       return
    //     }

    //     const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    //     if (!re.test(this.userEmail)) {
    //       this.error = 'not-valid-email'
    //       return
    //     }

    //     if (this.userPass === '' || this.confirmPass === '') {
    //       this.error = 'passwords-dont-match'
    //       return
    //     }

    //     // Проверка на совпадение паролей
    //     if (this.userPass !== this.confirmPass) {
    //       this.error = 'passwords-dont-match'
    //       return
    //     }

    //     // Отправка данных на сервер
    //     await AuthService.register({
    //       email: this.userEmail,
    //       password: this.userPass,
    //       nickname: this.userNickname,
    //       gender: this.userGender,
    //     })


    //     this.error = '' // Сброс ошибок
    //     alert('Регистрация прошла успешно!')

    //     this.$router.push('/profile') // Перенаправляем на профиль
    //   } catch (error) {
    //     console.error(error)
    //     this.error = 'ERROR-Registration'
    //   }
    // },
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
      });

      // Переходим на страницу профиля
      window.location.href = '/profile';
    };

    return {
      userStore,
      sendData,
    };
  },

})
</script>

