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

      <div class="search-bar">
        <input type="text" placeholder="Поиск..." v-model="searchQuery">
      </div>

      <div class="chat-list">
        <div v-for="chat in chats" :key="chat.id" class="chat-item">
          <div class="chat-avatar">
            <img :src="chat.avatar || defaultAvatar" alt="User avatar">
          </div>
          <div class="chat-info">
            <div class="chat-name">{{ chat.username }}</div>
            <div class="chat-status">{{ chat.status }}</div>
            <div class="chat-last-message">{{ chat.lastMessage }}</div>
          </div>
        </div>
      </div>
    </div>

    <div class="messenger-content">
      <RouterView />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import defaultAvatar from '../assets/default-avatar.png';
import Navbar from '../components/Navbar.vue';

export default defineComponent({
  name: 'MessengerView',
  components: {
    Navbar,
    //LoadingSpinner
  },

  setup() {
    const searchQuery = ref('');
    const chats = ref([
      {
        id: 1,
        username: 'Username',
        status: 'В сети',
        lastMessage: 'Последнее сообщение',
        avatar: '',
      },
      // Добавьте больше чатов для тестирования
    ]);

    return {
      searchQuery,
      chats,
      defaultAvatar,
    };
  },
});
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
</style>
