<template>
  <header>
    <Navbar />
  </header>

  <div class="messenger-container">
    <div class="messenger-sidebar">
      <div class="tab-buttons">
        <RouterLink to="/messenger/chat" class="tab-button" active-class="active">Чат</RouterLink>
        <RouterLink to="/messenger/match" class="tab-button" active-class="active">Match</RouterLink>
      </div>

      <template v-if="$route.path.includes('/messenger/chat')">
        <div class="search-bar">
          <input type="text" placeholder="Поиск..." v-model="searchQuery">
        </div>

        <div class="chat-list">
          <div v-for="chat in chats" :key="chat.user_id" class="chat-item"
            :class="{ 'active': activeChat?.user_id === chat.user_id }" @click="openDialog(chat)">
            <div class="chat-avatar" @click.stop="openProfile(chat)">
              <!-- <img :src="chat.avatar || defaultAvatar" alt="User avatar"> -->
              <img :src="defaultAvatar" alt="User avatar">
              <div class="status-indicator" :class="chat.online_status"></div>
            </div>
            <div class="chat-info">
              <div class="chat-name">{{ chat.username }}</div>
              <div class="chat-status">{{ chat.online_status }}</div>
            </div>
          </div>
        </div>
      </template>
    </div>

    <div class="messenger-content">
      <template v-if="$route.path.includes('/messenger/chat')">
        <div v-if="!activeChat" class="no-chat-selected">
          Выберите чат для начала общения
        </div>

        <div v-else class="chat-dialog">
          <div class="chat-header">
            <div class="user-info">
              <!-- <img :src="activeChat.avatar || defaultAvatar" alt="User avatar" class="user-avatar"> -->
              <img :src="defaultAvatar" alt="User avatar" class="user-avatar">
              <div class="user-details">
                <div class="user-name">{{ activeChat.username }}</div>
                <div class="user-status">{{ activeChat.online_status }}</div>
              </div>
            </div>
          </div>

          <div class="messages-container">
            <div v-for="(message, index) in messages" :key="index" class="message"
              :class="{ 'message-mine': message.isMine }">
              <div class="message-content">
                <div class="message-text">{{ message.text }}</div>
                <div class="message-time">{{ formatTime(message.timestamp) }}</div>
              </div>
            </div>
          </div>

          <div class="message-input-container">
            <div class="emoji-wrapper">
              <button class="emoji-button" @click.stop="toggleEmojiPicker">
                <i class="far fa-smile"></i>
              </button>

              <div v-if="showEmojiPicker" class="emoji-picker-container" @click.stop>
                <emoji-picker @emoji-click="handleEmojiSelect"></emoji-picker>
              </div>
            </div>

            <input type="text" v-model="newMessage" @keyup.enter="sendMessage" placeholder="Введите сообщение..."
              class="message-input">

            <button @click="sendMessage" class="send-button">
              <i class="fas fa-paper-plane"></i>
            </button>
          </div>
        </div>
      </template>

      <RouterView v-else />
    </div>

    <div v-if="selectedUser" class="user-profile-modal">
      <div class="modal-overlay" @click="closeProfile"></div>
      <div class="modal-content">
        <button class="close-btn" @click="closeProfile">
          <i class="fas fa-times"></i>
        </button>

        <div class="profile-content">
          <!-- <img :src="selectedUser.avatar || defaultAvatar" alt="User avatar" class="profile-avatar"> -->
          <img :src="defaultAvatar" alt="User avatar" class="profile-avatar">
          <div class="profile-info">
            <h2>{{ selectedUser.username }}</h2>
            <p class="bio">{{ selectedUser.biography }}</p>
            <div class="interests" v-if="selectedUser.label">
              <span v-for="interest in selectedUser.label" :key="interest" class="interest-tag">
                {{ interest }}
              </span>
            </div>
            <div class="status" :class="selectedUser.online_status">
              {{ selectedUser.online_status }}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
defineOptions({
  name: 'MessengerView'
});

import { ref, onMounted, onUnmounted } from 'vue';
import AuthService from '../services/AuthService';
import type { ChatUser } from '../classes';
import Navbar from '../components/Navbar.vue';
import defaultAvatar from '../assets/default-avatar.png';
import 'emoji-picker-element';

interface Message {
  text: string;
  timestamp: Date;
  isMine: boolean;
}

const chats = ref<ChatUser[]>([]);
const isLoading = ref(true);
const searchQuery = ref('');
const selectedUser = ref<ChatUser | null>(null);
const activeChat = ref<ChatUser | null>(null);
const newMessage = ref('');
const messages = ref<Message[]>([]);
const showEmojiPicker = ref(false);

onMounted(async () => {
  try {
    const response = await AuthService.getChats();
    chats.value = response.chats;
  } catch (error) {
    console.error('Ошибка при загрузке чатов:', error);
  } finally {
    isLoading.value = false;
  }
});

onMounted(() => {
  document.addEventListener('click', handleClickOutside);
});

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside);
});

const openProfile = (user: ChatUser) => {
  selectedUser.value = user;
};

const closeProfile = () => {
  selectedUser.value = null;
};

const openDialog = (user: ChatUser) => {
  activeChat.value = user;
};

const sendMessage = () => {
  if (newMessage.value.trim()) {
    const message: Message = {
      text: newMessage.value,
      timestamp: new Date(),
      isMine: true
    };
    messages.value.push(message);
    newMessage.value = '';

    // Прокрутка к последнему сообщению
    setTimeout(() => {
      const container = document.querySelector('.messages-container');
      if (container) {
        container.scrollTop = container.scrollHeight;
      }
    }, 0);
  }
};

const formatTime = (date: Date) => {
  return new Intl.DateTimeFormat('ru-RU', {
    hour: '2-digit',
    minute: '2-digit'
  }).format(date);
};

const toggleEmojiPicker = () => {
  showEmojiPicker.value = !showEmojiPicker.value;
};

interface EmojiEvent {
  detail: {
    emoji: {
      unicode: string;
    };
  };
}

const handleEmojiSelect = (e: EmojiEvent) => {
  const emoji = e.detail.emoji.unicode;
  newMessage.value += emoji;
};

const handleClickOutside = (e: MouseEvent) => {
  const picker = document.querySelector('.emoji-picker-container');
  const button = document.querySelector('.emoji-button');

  if (picker && button &&
    !picker.contains(e.target as Node) &&
    !button.contains(e.target as Node)) {
    showEmojiPicker.value = false;
  }
};
</script>

<style scoped>
.messenger-container {
  position: fixed;
  display: flex;
  width: 90%;
  height: 80vh;
  max-width: 1200px;
  border-radius: 30px;
  overflow: hidden;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
  /* border: 2px solid var(--primary-color); */
  background-color: #fff;
  left: 50%;
  top: 6em;
  transform: translateX(-50%);
}

.messenger-sidebar {
  width: 300px;
  background-color: var(--secondary-color);
  border-right: 1px solid #e0e0e0;
  display: flex;
  flex-direction: column;
  border-radius: 30px;
  margin: 20px 0 20px 20px;
  padding: 10px;
}

.messenger-content {
  flex: 1;
  background-color: #f5f5f5;
  border-top-right-radius: 30px;
  border-bottom-right-radius: 30px;
}

.tab-buttons {
  display: flex;
  padding: 10px;
  gap: 10px;
}

.tab-button {
  flex: 1;
  padding: 10px;
  text-align: center;
  background-color: #808080;
  color: white;
  border-radius: 8px;
  cursor: pointer;
  text-decoration: none;
  transition: background-color 0.3s;
}

.tab-button.active {
  background-color: var(--primary-color);
}

.search-bar {
  padding: 10px;
}

.search-bar input {
  width: 100%;
  padding: 8px;
  border: 1px solid #e0e0e0;
  border-radius: 4px;
}

.chat-list {
  flex: 1;
  overflow-y: auto;
}

.chat-item {
  display: flex;
  padding: 15px;
  gap: 10px;
  cursor: pointer;
  transition: background-color 0.3s;
  border-bottom: 1px solid #e0e0e0;
}

.chat-item:hover {
  background-color: #f0f0f0;
}

.chat-avatar {
  width: 50px;
  height: 50px;
}

.chat-avatar img {
  width: 100%;
  height: 100%;
  border-radius: 50%;
  object-fit: cover;
}

.chat-info {
  flex: 1;
}

.chat-name {
  font-weight: bold;
  margin-bottom: 2px;
}

.chat-status {
  font-size: 12px;
  color: #666;
  margin-bottom: 2px;
}

.chat-last-message {
  font-size: 14px;
  color: #666;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-profile-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
}

.modal-content {
  position: relative;
  background: white;
  padding: 30px;
  border-radius: 15px;
  max-width: 500px;
  width: 90%;
  z-index: 1;
}

.close-btn {
  position: absolute;
  top: 15px;
  right: 15px;
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
}

.profile-avatar {
  width: 150px;
  height: 150px;
  border-radius: 50%;
  object-fit: cover;
  margin-bottom: 20px;
}

.profile-info {
  text-align: center;
}

.profile-info h2 {
  margin-bottom: 10px;
}

.bio {
  margin: 15px 0;
  color: #666;
}

.interests {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  justify-content: center;
  margin: 15px 0;
}

.interest-tag {
  background: #f0f0f0;
  padding: 5px 12px;
  border-radius: 15px;
  font-size: 14px;
}

.status {
  display: inline-block;
  padding: 5px 12px;
  border-radius: 15px;
  font-size: 14px;
  margin-top: 10px;
}

.status.online {
  background: #4CAF50;
  color: white;
}

.status.offline {
  background: #9e9e9e;
  color: white;
}

.chat-dialog {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.chat-header {
  padding: 15px;
  border-bottom: 1px solid #e0e0e0;
  background: white;
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
  flex: 1;
}

.user-name {
  font-weight: bold;
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
  background: #f5f5f5;
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.message {
  display: flex;
  margin-bottom: 10px;
}

.message-content {
  max-width: 70%;
  padding: 10px 15px;
  border-radius: 15px;
  background: white;
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
  position: relative;
}

.message-mine {
  flex-direction: row-reverse;
}

.message-mine .message-content {
  background: var(--primary-color);
  color: white;
}

.message-text {
  margin-bottom: 4px;
  word-wrap: break-word;
}

.message-time {
  font-size: 12px;
  color: #666;
  text-align: right;
}

.message-mine .message-time {
  color: rgba(255, 255, 255, 0.8);
}

.message-input-container {
  padding: 15px;
  background: white;
  border-top: 1px solid #e0e0e0;
  display: flex;
  gap: 10px;
  align-items: center;
  position: relative;
}

.message-input {
  flex: 1;
  padding: 12px 20px;
  border: 1px solid #e0e0e0;
  border-radius: 20px;
  outline: none;
  font-size: 14px;
  transition: border-color 0.3s;
}

.message-input:focus {
  border-color: var(--primary-color);
}

.send-button {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: var(--primary-color);
  color: white;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: transform 0.2s;
}

.send-button:hover {
  transform: scale(1.1);
}

.send-button:active {
  transform: scale(0.95);
}

.no-chat-selected {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
  color: #666;
  font-size: 18px;
}

.chat-item.active {
  background: #f0f0f0;
}

.emoji-wrapper {
  position: relative;
  z-index: 1000;
}

.emoji-button {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: none;
  border: none;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: #666;
  transition: color 0.2s;
}

.emoji-button:hover {
  color: var(--primary-color);
}

.emoji-picker-container {
  position: absolute;
  bottom: calc(100% + 10px);
  left: 0;
  background: white;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

emoji-picker {
  width: 300px;
  height: 400px;
  --background: white;
  --category-emoji-size: 1.25rem;
  --emoji-size: 1.5rem;
  --indicator-color: var(--primary-color);
  --border-color: #e0e0e0;
  --category-font-size: 0.9rem;
}
</style>
