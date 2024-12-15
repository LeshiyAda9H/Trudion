<template>
  <div class="user-card" @click="openModal">

    <div class="more-options">...</div>

    <div class="avatar">
      <!-- <img :src="user.avatar || '/default-avatar.png'" alt="avatar" /> -->
      <img :src="defaultAvatar" alt="avatar" />
    </div>

    <div class="user-info">

      <h3 class="nickname">{{ user.username }}</h3>

      <div class="labels">
        <span v-for="label in user.label" :key="label" class="label-tag">
          {{ label }}
        </span>
      </div>

      <button class="friend-button" @click="handleFriendRequest">
        Дружить
      </button>
    </div>

  </div>
  <UserModal
    :show="isModalOpen"
    :user="user"
    @close="closeModal"
  />
</template>

<script lang="ts">
import { defineComponent, type PropType, ref } from 'vue';
import { type ProfileUser } from '../classes';
import defaultAvatar from '../assets/default-avatar.png';
import UserModal from './UserModal.vue';

export default defineComponent({
  name: 'UserCard',
  components: {
    UserModal
  },
  props: {
    user: {
      type: Object as PropType<ProfileUser>,
      required: true
    }
  },
  setup() {
    const isModalOpen = ref(false);

    const openModal = () => {
      isModalOpen.value = true;
    };

    const closeModal = () => {
      isModalOpen.value = false;
    };

    const handleFriendRequest = () => {
      // Логика добавления в друзья
    };

    return {
      isModalOpen,
      openModal,
      closeModal,
      handleFriendRequest,
      defaultAvatar
    };
  }
});
</script>

<style scoped>

.user-card {
  background: #fff;
  border-radius: var(--border-radius);
  padding: 15px;
  width: 20em;
  height: auto;
  box-shadow: 0px 4px 20px 0px #00000040;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 0 auto;

  background-color: var(--secondary-color);
}

.nickname {
  font-size: 30px;
  font-weight: bold;
  margin: 1em 0;
}

.more-options {
  position: absolute;
  right: 10px;
  top: 10px;
}

.avatar img {
  width: 10em;
  height: 10em;
  border-radius: 50%;
  object-fit: cover;
}

.user-info {
  text-align: center;
}


.labels {
  display: flex;
  flex-wrap: wrap;
  gap: 5px;
  justify-content: center;
  margin: 10px 0;
}

.label-tag {
  background: var(--secondary-color);
  padding: 4px 0.7em;
  border-radius: var(--border-radius);
  border: 3px solid var(--primary-color);
  font-size: 24px;
  font-weight: bold;
}

.friend-button {
  background: #b08d57;
  margin-top: 1em;
  color: white;
  border: none;
  width: 10em;
  padding: 8px 0.7em;
  border-radius: var(--border-radius);
  cursor: pointer;
  font-size: 24px;

}
</style>
