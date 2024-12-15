<template>
  <header>
    <Navbar />
  </header>

  <div class="home-container">
    <aside class="filter-sidebar">
      <FilterComponent @filter-changed="handleFilterChange" />
    </aside>

    <main class="users-grid">
      <div v-if="error" class="error-message">
        {{ error }}
      </div>
      <LoadingSpinner v-else-if="isLoading" />

      <div v-else class="users-list">
        <UserCard v-for="user in users" :key="user.username" :user="user" />
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import Navbar from '../components/Navbar.vue';
import FilterComponent from '../components/FilterComponent.vue';
import UserCard from '../components/UserCard.vue';
import LoadingSpinner from '../components/LoadingSpinner.vue';
import AuthService from '../services/AuthService';
import { type ProfileUser } from '../classes';

// Сначала определим интерфейс для фильтров
interface FilterParams {
  usersNumber: number;
  labels?: string[];
}

interface ApiResponse {
  data: ProfileUser[];
}

export default defineComponent({
  name: 'HomePage',
  components: {
    Navbar,
    FilterComponent,
    UserCard,
    LoadingSpinner
  },
  setup() {
    const users = ref<ProfileUser[]>([]);
    const isLoading = ref(false);
    const error = ref<string | null>(null);

    const fetchUsers = async (params: FilterParams) => {
      isLoading.value = true;
      error.value = null;

      try {
        const response = await AuthService.getUsers(params) as ApiResponse;
        if (response && 'data' in response) {
          users.value = response.data;
        } else {
          error.value = 'Некорректный ответ от сервера';
          users.value = [];
        }
      } catch  {
        error.value = 'Ошибка при загрузке пользователей';
        users.value = [];
      } finally {
        isLoading.value = false;
      }
    };

    const handleFilterChange = (filters: FilterParams) => {
      fetchUsers(filters);
    };

    onMounted(() => {
      fetchUsers({ usersNumber: 10 });
    });

    return {
      users,
      isLoading,
      error,
      handleFilterChange
    };
  }
});
</script>

<style scoped>


.home-container {
  display: grid;
  grid-template-columns: 300px 1fr;
  gap: 20px;
  padding: 20px;
  max-width: 1400px;
  margin: 0 auto;
  margin-top: 5em;
}

.filter-sidebar {
  position: sticky;
  top: 20px;
  height: fit-content;
}

.users-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.users-list {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}

.error-message {
  color: #dc3545;
  padding: 1rem;
  text-align: center;
  background-color: #f8d7da;
  border: 1px solid #f5c6cb;
  border-radius: 4px;
  margin-bottom: 1rem;
}
</style>
