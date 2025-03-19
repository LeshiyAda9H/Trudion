<template>
  <div class = "screen">
  <div class="auth-container">
    <p class="auth-title">Регистрация</p>
    <div class="image-container">
      <!-- Всплывающее облачко -->
      <TooltipMessage
      :message="modalMessage"
      :visible="showModal"
      :position="tooltipPosition"
      />
      <img :src="imagePath" class="image-ava" />
    </div>

    <InputRegistration
      :writeEmail="writeEmail"
      :writePass="writePass"
      :writeConfirmPass="writeConfirmPass"
      :error="error"
    />

    <button class="auth-button" @click="sendData">Зарегистрироваться</button>

    <div class="footer-text">
      Уже есть аккаунт? <router-link to="/login" class="link">Войти</router-link>
    </div>

    <!-- Модальное окно -->
    <ModalWindow :message="modalMessage" :visible="showModal" @close="showModal = false" />

  </div>

</div>
</template>

<script lang="ts">
import { defineComponent, ref, reactive } from 'vue'
import myImage from '../assets/trudion.svg'
import InputRegistration from '../components/InputRegistration.vue'
import '../assets/css/authentication.css'
import { useUserStore } from '../stores/UserStore'; // Импортируем хранилище Pinia
import AuthService from '../services/AuthService'; // Импортируем AuthService для регистрации
import { useRouter } from "vue-router"; // Используем Composition API для роутера
//import ModalWindow from '../components/ModalWindow.vue';

import TooltipMessage from '../components/TooltipMessage.vue';

export default defineComponent({
  name: 'RegistrationPage',

  components: { InputRegistration, TooltipMessage },

  setup() {
    const router = useRouter(); // Инициализация роутера
    const userStore = useUserStore(); // Инициализация хранилища Pinia

    const imagePath = myImage;
    const error = ref(''); // Ошибка
    const userEmail = ref(''); // Электронная почта
    const userPass = ref(''); // Пароль
    const confirmPass = ref(''); // Подтверждение пароля

    const showModal = ref(false); // Флаг видимости модального окна
    const modalMessage = ref(''); // Сообщение для модального окна

    // Позиция облачка
    const tooltipPosition = reactive({
    top: 0,
    left: 100
  });

    const writeEmail = (text_email: string): void => {
      userEmail.value = text_email;
    };

    const writePass = (text_pass: string): void => {
      userPass.value = text_pass;
    };

    const writeConfirmPass = (text_confirmPass: string): void => {
      confirmPass.value = text_confirmPass;
    };

    const showError = (errorCode: string, message: string) => {
      error.value = errorCode;
      modalMessage.value = message;
      showModal.value = true;
    };

    const validateData = (): boolean => {
      if (userEmail.value === '') {
        showError('not-valid-email', "Поле email обязательно для заполнения.");
        return false;
      }

      const re = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!re.test(userEmail.value)) {
        showError('not-valid-email', "Некорректный формат email.");
        return false;
      }

      if (userPass.value === '' || confirmPass.value === '') {
        showError('passwords-dont-match', "Поля с паролем обязательны для заполнения.");
        return false;
      }

      if (userPass.value !== confirmPass.value) {
        showError('passwords-dont-match', "Пароли не совпадают. Попробуйте ещё раз.");
        return false;
      }

      if (userPass.value.length < 6) {
        showError('passwords-too-short', "Пароль должен быть не короче 6 символов.");
        return false;
      }

      return true;
    };

    const sendData = async (): Promise<void> => {
      try {
        // Валидация данных
        if (!validateData()) {
          return;
        }

        const isEmailAvailable = await AuthService.verifyEmail(userEmail.value);
        if (!isEmailAvailable.available) {
          showError('email-already-exists', "Этот email уже зарегистрирован.");
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
        showError('server-error', "Этот email уже зарегистрирован.");
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
      sendData,
      showModal, // Флаг видимости модального окна
      modalMessage, // Сообщение для модального окна
      tooltipPosition,
    };
  },
});
</script>
