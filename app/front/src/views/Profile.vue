<template>
   <header>
      <Navbar/>
  </header>

  <body>
    <h1 class="title">МОЙ ПРОФИЛЬ</h1>

    <div class="container">

      <div class="avatar-container">
        <div class="avatar">
          <i class="icon-user"></i>
          <div class="add-photo">+</div>
        </div>
      </div>

      <!-- Поле для ввода никнейма -->
      <input
        type="text"
        placeholder="Никнейм"
        class="input"
        v-model="userStore.registrationData.nickname"
      />

      <!-- Выбор пола -->
      <select
        class="select"
        v-model="userStore.registrationData.gender"
      >
        <option value="" disabled>Выбери свой пол</option>
        <option value="male">Мужчина</option>
        <option value="female">Женщина</option>
        <option value="other">Абоба</option>
      </select>

      <!-- Поле "О себе" -->
      <textarea
        placeholder="О себе"
        class="form-textarea"
        v-model="userStore.registrationData.bio"
      ></textarea>

      <!-- Выбор навыков -->
      <select
        class="select"
        v-model="userStore.registrationData.skills"
      >
        <option value="" disabled>Навыки</option>
        <option value="skill1">Навык 1</option>
        <option value="skill2">Навык 2</option>
        <option value="skill3">Навык 3</option>
      </select>

    </div>

    <button class="button"  @click="updateProfile">Сохранить изменения</button>

  </body>

</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Navbar from '../components/Navbar.vue';
import { useUserStore } from '../stores/UserStore';
import { useRouter } from 'vue-router';
import '../assets/css/profile.css'
import AuthService from '../services/AuthService';

export default defineComponent({
  name: 'ProfilePage',
  components: {Navbar},

  setup() {
    const userStore = useUserStore();
    const router = useRouter();

    const updateProfile = async () => {
      try {
        const { nickname, gender } = userStore.registrationData;

        if (!nickname || !gender) {
          alert('Пожалуйста, заполните никнейм и выберите пол');
          return;
        }

        // Отправляем данные профиля на сервер через AuthService
        await AuthService.updateProfile({ nickname, gender });

        // Сохраняем данные в Pinia-хранилище
        userStore.setRegistrationData({ nickname, gender });

        alert('Профиль успешно обновлен!');

        // Перенаправляем пользователя на главную страницу
        router.push('/home');

      } catch (error) {
        console.error('Ошибка при обновлении профиля:', error);
        alert('Ошибка при обновлении профиля');
      }
    };

    return {
      userStore,
      updateProfile,
    };



  },



});
</script>
