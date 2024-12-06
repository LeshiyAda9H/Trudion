<template>
  <div class="home">
    <h1>Добро пожаловать на главную страницу!</h1>
    <p>Здесь будет содержимое вашей главной страницы.</p>

    <!-- Кнопка для выхода из аккаунта -->
    <button class="logout-button" @click="logout">Выйти</button>
  </div>

  <button @click="goToUsersList" class="go-to-users-button">
      Перейти к списку пользователей
    </button>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router'; // Импортируем useRouter для навигации

export default defineComponent({
  name: 'HomePage',
  setup() {
    const router = useRouter();

    // Функция для перехода на страницу списка пользователей
    const goToUsersList = () => {
      router.push('/users-list');
    };

    return { goToUsersList };
  },

  methods: {
    // Метод для выхода из аккаунта
    logout() {
      // Удаляем токен из localStorage
      localStorage.removeItem('authToken');

      // Перенаправляем пользователя на страницу авторизации
      this.$router.push('/login');
    },
  },
});
</script>

<style scoped>
.home {
  text-align: center;
  margin-top: 50px;
}

.logout-button {
  margin-top: 20px;
  padding: 10px 20px;
  font-size: 16px;
  background-color: #ff4c4c;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

.logout-button:hover {
  background-color: #ff2a2a;
}

.go-to-users-button {
  margin-top: 20px;
  background-color: #355299;
  border: none;
  border-radius: 25px;
  color: white;
  padding: 10px 30px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: transform 500ms ease;
}

.go-to-users-button:hover {
  transform: scale(1.05) translateY(-0.5px);
  opacity: 0.9;
}

</style>
