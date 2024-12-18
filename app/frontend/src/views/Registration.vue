<template>
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

    <InputRegistration :writeEmail="writeEmail" :writePass="writePass" :writeConfirmPass="writeConfirmPass"
      :error="error" />

    <button class="auth-button" @click="sendData">Зарегистрироваться</button>

    <div class="footer-text">
      Уже есть аккаунт? <router-link to="/login" class="link">Войти</router-link>
    </div>

    <!-- Модальное окно -->
    <ModalWindow :message="modalMessage" :visible="showModal" @close="showModal = false" />

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


<style scoped>
.auth-container {
  border-radius: 15px;
  background-color: var(--secondary-color);
  box-shadow: 0px 4px 20px 0px #00000040;
  width: 30%;
  display: grid;
  place-items: center;
  margin: 6em auto;
  padding: 1em;
  text-align: center;
  position: relative;
  /* Добавлено для позиционирования */
}

.auth-title {
  font-family: "Press Start 2P", system-ui;
  font-weight: 400;
  font-style: normal;
  margin-bottom: 1em;
}

.image-container {
  position: relative;
  display: inline-block;
  width: 200px;
  height: 200px;
  margin: 0 auto;
}

.image-ava {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

/* Стили для позиционирования тултипа */
:deep(.tooltip) {
  position: absolute;
  top: 0;
  left: 100%;
  transform: translateX(10px);
  /* Отступ от изображения */
  z-index: 1000;
}

.auth-button {
  padding: 15px 20px;
  background-color: var(--primary-color);
  color: #fff;
  font-size: 16px;
  font-weight: bold;
  text-align: center;
  margin-top: 5px;
  min-width: 70%;
}

.footer-text {
  color: #8894b1;
  margin-top: 15px;
  margin-bottom: 15px;
  font-size: 14px;
}

.link {
  cursor: pointer;
  opacity: 0.9;
  color: #355299;
  text-decoration: none;
}

.link:hover {
  color: #a68136;
}
</style>
