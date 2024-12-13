<template>
  <header>
    <Navbar />
  </header>

  <body>
    <h1 class="title">Мой профиль</h1>

    <div class="prof-container">
      <div class="avatar-container">
        <div class="avatar">
          <i class="icon-user"></i>
          <div class="add-photo">+</div>
        </div>
      </div>

      <input type="text" placeholder="Никнейм" class="input" v-model="userStore.registrationData.username" />
      <select class="select" v-model="userStore.registrationData.gender">
        <option value="" disabled>Выбери свой пол</option>
        <option value="male">Мужчина</option>
        <option value="female">Женщина</option>
        <option value="other">Абоба</option>
      </select>
      <textarea placeholder="О себе" class="form-textarea" v-model="userStore.registrationData.bio"></textarea>
      <select class="select" v-model="userStore.registrationData.skills">
        <option value="" disabled>Навыки</option>
        <option value="skill1">Навык 1</option>
        <option value="skill2">Навык 2</option>
        <option value="skill3">Навык 3</option>
      </select>
    </div>

    <button class="save-button" @click="updateProfile">Сохранить изменения</button>
  </body>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import Navbar from '../components/Navbar.vue';
import { useUserStore } from '../stores/UserStore';
import '../assets/css/profile.css'

export default defineComponent({
  name: 'ProfilePage',
  components: { Navbar },

  setup() {
    const userStore = useUserStore();
    return {
      userStore,
    };
  },
});
</script>
