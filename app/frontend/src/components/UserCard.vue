<template>
  <div class="user-card">
    <div class="more-options">...</div>
    <div class="avatar">
      <img :src="user.avatar || '/default-avatar.png'" alt="avatar" />
    </div>
    <div class="user-info">
      <h3>{{ user.username }}</h3>
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
</template>

<script lang="ts">
import { defineComponent, type PropType } from 'vue';
import { type ProfileUser } from '../classes';

export default defineComponent({
  name: 'UserCard',
  props: {
    user: {
      type: Object as PropType<ProfileUser>,
      required: true
    }
  },
  setup() {
    const handleFriendRequest = () => {
      // Логика добавления в друзья
    };

    return {
      handleFriendRequest
    };
  }
});
</script>

<style scoped>
.user-card {
  background: #fff;
  border-radius: 10px;
  padding: 15px;
  margin: 10px;
  width: 250px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  position: relative;
}

.more-options {
  position: absolute;
  right: 10px;
  top: 10px;
}

.avatar img {
  width: 80px;
  height: 80px;
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
  background: #f0f0f0;
  padding: 4px 8px;
  border-radius: 15px;
  font-size: 12px;
}

.friend-button {
  background: #b08d57;
  color: white;
  border: none;
  padding: 8px 20px;
  border-radius: 20px;
  cursor: pointer;
}
</style>
