<template>
  <div class="user-card" @click="openModal">
    <div class="more-options" @click.stop="toggleOptions">
      <span class="options-icon">...</span>
      <div v-if="showOptions" class="options-menu">
        <button @click="handleReport">Пожаловаться</button>
        <button @click="handleBlock">Заблокировать</button>
      </div>
    </div>

    <div class="avatar">
      <img :src="defaultAvatar" alt="avatar" />
      <div v-if="user.online_status" :class="['online-status', user.online_status]"
        :title="getStatusText(user.online_status)"></div>
    </div>

    <div class="user-info">
      <h3 class="nickname">{{ user.username }}</h3>
      <div class="labels">
        <span v-for="label in user.label" :key="label" class="label-tag" :title="getInterestName(label)">
          {{ getInterestName(label) }}
        </span>
      </div>

      <button class="friend-button" @click.stop="handleFriendRequest" :disabled="isSending">
        {{ friendButtonText }}
      </button>
    </div>
  </div>

  <UserModal :show="isModalOpen" :user="user" @close="closeModal" />
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import type { ProfileUser } from '../classes'
import defaultAvatar from '../assets/default-avatar.png'
import UserModal from './UserModal.vue'
import { interests } from '../config/interests'
import AuthService from '../services/AuthService'

const props = defineProps<{
  user: ProfileUser
}>()

const isModalOpen = ref(false)
const showOptions = ref(false)
const isSending = ref(false)

const friendButtonText = computed(() => {
  if (isSending.value) {
    return 'Обработка...'
  }
  return 'Дружить'
})

const openModal = () => {
  isModalOpen.value = true
}

const closeModal = () => {
  isModalOpen.value = false
}

const toggleOptions = () => {
  showOptions.value = !showOptions.value
}

const emit = defineEmits(['match-success']);

const handleFriendRequest = async () => {
  try {
    isSending.value = true;
    const response = await AuthService.sendFriendRequest(props.user.user_id);

    if (response.message === 'mutually') {
      emit('match-success', true);
    }
  } catch (error) {
    console.error('Ошибка при отправке запроса:', error);
  } finally {
    isSending.value = false;
  }
};

const handleReport = () => {
  console.log('Отправка жалобы')
  showOptions.value = false
}

const handleBlock = () => {
  console.log('Блокировка пользователя')
  showOptions.value = false
}

const getStatusText = (status: string) => {
  const statusMap: Record<string, string> = {
    online: 'В сети',
    offline: 'Не в сети',
    away: 'Отошел'
  }
  return statusMap[status] || status
}

const getInterestName = (value: string): string => {
  const interest = interests.find(i => i.value === value);
  return interest ? interest.name : value;
};

onMounted(() => {
  console.log('UserCard получил пользователя:', props.user)
  if (!props.user.user_id) {
    console.warn('Пользователь без ID:', props.user)
  }
})
</script>

<style scoped>
.user-card {
  background: #fff;
  border-radius: var(--border-radius);
  padding: 15px;
  width: 21em;
  height: 32em;
  box-shadow: 0px 4px 20px 0px #00000040;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
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

.avatar {
  position: relative;
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
  border: 2px solid var(--primary-color);
  font-size: 24px;
}

.friend-button {
  background: #b08d57;
  color: white;
  border: none;
  width: 10em;
  padding: 8px 0.7em;
  border-radius: var(--border-radius);
  cursor: pointer;
  font-size: 24px;
  position: absolute;
  bottom: 1em;
  left: 50%;
  transform: translateX(-50%);
}

.options-icon {
  cursor: pointer;
  padding: 5px;
}

.options-menu {
  position: absolute;
  top: 30px;
  right: 10px;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  padding: 8px;
  z-index: 100;
}

.options-menu button {
  display: block;
  width: 100%;
  padding: 8px 16px;
  border: none;
  background: none;
  text-align: left;
  cursor: pointer;
  border-radius: 4px;
}

.options-menu button:hover {
  background: #f5f5f5;
}

.online-status {
  position: absolute;
  bottom: 5px;
  right: 5px;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  border: 1px solid white;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
}

.online-status.online {
  background: #4CAF50;
}

.online-status.offline {
  background: #9e9e9e;
}

.online-status.away {
  background: #FFC107;
}

.friend-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}
</style>
