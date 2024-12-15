<template>
  <div class="chat-container">
    <!-- Заголовок чата с информацией о пользователе -->
    <div class="chat-header">
      <div class="user-info">
        <img :src="selectedUser.avatar || defaultAvatar" alt="User avatar" class="user-avatar">
        <div class="user-details">
          <div class="user-name">{{ selectedUser.name }}</div>
          <div class="user-status">{{ selectedUser.status }}</div>
        </div>
      </div>
    </div>

    <!-- Область сообщений -->
    <div class="messages-container" ref="messagesContainer">
      <div v-for="message in messages"
           :key="message.id"
           :class="['message', message.isOwn ? 'own-message' : 'other-message']">
        <div class="message-content">
          <div class="message-text">{{ message.text }}</div>
          <div class="message-time">{{ formatTime(message.timestamp) }}</div>
        </div>
      </div>
    </div>

    <!-- Форма отправки сообщения -->
    <div class="message-input-container">
      <input
        type="text"
        v-model="newMessage"
        @keyup.enter="sendMessage"
        placeholder="Введите сообщение..."
        class="message-input"
      >
      <button @click="sendMessage" class="send-button">
        <i class="fas fa-paper-plane"></i>
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, nextTick } from 'vue';
import defaultAvatar from '../assets/default-avatar.png';

export default defineComponent({
  name: 'ChatView',

  setup() {
    const messagesContainer = ref<HTMLElement | null>(null);
    const newMessage = ref('');
    const selectedUser = ref({
      name: 'Пользователь',
      status: 'В сети',
      avatar: '',
    });

    // Временные данные для демонстрации
    const messages = ref([
      {
        id: 1,
        text: 'Привет!',
        timestamp: new Date(),
        isOwn: false,
      },
      {
        id: 2,
        text: 'Здравствуйте!',
        timestamp: new Date(),
        isOwn: true,
      },
    ]);

    const scrollToBottom = async () => {
      await nextTick();
      if (messagesContainer.value) {
        messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight;
      }
    };

    const sendMessage = () => {
      if (newMessage.value.trim()) {
        messages.value.push({
          id: Date.now(),
          text: newMessage.value,
          timestamp: new Date(),
          isOwn: true,
        });
        newMessage.value = '';
        scrollToBottom();
      }
    };

    const formatTime = (date: Date) => {
      return new Date(date).toLocaleTimeString([], {
        hour: '2-digit',
        minute: '2-digit'
      });
    };

    onMounted(() => {
      scrollToBottom();
    });

    return {
      messages,
      newMessage,
      selectedUser,
      defaultAvatar,
      sendMessage,
      formatTime,
      messagesContainer,
    };
  },
});
</script>

<style scoped>
.chat-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  background-color: #f5f5f5;
}

.chat-header {
  padding: 15px;
  background-color: white;
  border-bottom: 1px solid #e0e0e0;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.user-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  object-fit: cover;
}

.user-details {
  display: flex;
  flex-direction: column;
}

.user-name {
  font-weight: bold;
}

.user-status {
  font-size: 12px;
  color: #666;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.message {
  display: flex;
  margin-bottom: 10px;
}

.message-content {
  max-width: 60%;
  padding: 10px 15px;
  border-radius: 15px;
  position: relative;
}

.own-message {
  justify-content: flex-end;
}

.own-message .message-content {
  background-color: var(--primary-color);
  color: white;
  border-bottom-right-radius: 5px;
}

.other-message .message-content {
  background-color: white;
  border-bottom-left-radius: 5px;
}

.message-text {
  margin-bottom: 5px;
  word-wrap: break-word;
}

.message-time {
  font-size: 11px;
  opacity: 0.7;
  text-align: right;
}

.message-input-container {
  padding: 15px;
  background-color: white;
  display: flex;
  gap: 10px;
  border-top: 1px solid #e0e0e0;
}

.message-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #e0e0e0;
  border-radius: 20px;
  outline: none;
}

.message-input:focus {
  border-color: var(--primary-color);
}

.send-button {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background-color: var(--primary-color);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.3s;
}

.send-button:hover {
  background-color: var(--secondary-color);
}
</style>
