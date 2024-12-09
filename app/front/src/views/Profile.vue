<template>
   <header>
      <Navbar/>
  </header>

  <body>
    <h1 class="title">МОЙ ПРОФИЛЬ</h1>

    <div class="auth-container">

      <div class="avatar-container">
        <div class="avatar">
          <i class="icon-user"></i>
          <div class="add-photo">+</div>
        </div>
      </div>

      <input type="text" placeholder="Никнейм" class="input-field" />
      <select class="form-select">
        <option value="" disabled selected>Выбери свой пол</option>
        <option value="skill1">Мужчина</option>
        <option value="skill2">Женщина</option>
        <option value="skill3">Абоба</option>
      </select>
      <textarea placeholder="О себе" class="form-textarea"></textarea>

      <select class="form-select">
        <option value="" disabled selected>Навыки</option>
        <option value="skill1">Навык 1</option>
        <option value="skill2">Навык 2</option>
        <option value="skill3">Навык 3</option>
      </select>

    </div>

    <button class="auth-button">Сохранить изменения</button>

  </body>

</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useRouter } from 'vue-router'; // Импортируем useRouter для навигации
import '../assets/profile.css'
import Navbar from '../components/Navbar.vue'

export default defineComponent({
  name: 'ProfilePage',
  components: {Navbar},
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
