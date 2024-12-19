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

      <div v-else-if="users.length === 0" class="no-results">
        <h3>Пользователи не найдены</h3>
        <p>Попробуйте изменить параметры фильтра</p>
      </div>

      <div v-else class="users-list">
        <UserCard v-for="user in users" :key="user.username" :user="user" />
      </div>

    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Navbar from '../components/Navbar.vue'
import FilterComponent from '../components/FilterComponent.vue'
import UserCard from '../components/UserCard.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
import AuthService from '../services/AuthService'
import type { ProfileUser } from '../classes'

interface FilterParams {
  usersNumber: number
  labels?: string[]
}

const users = ref<ProfileUser[]>([])
const isLoading = ref(false)
const error = ref<string | null>(null)

onMounted(async () => {
  console.log('Home компонент смонтирован')
  try {
    await handleFilterChange({ usersNumber: 10 })
    console.log('Начальная загрузка пользователей завершена')
  } catch (error) {
    console.error('Ошибка при начальной загрузке:', error)
  }
})

const handleFilterChange = async (filters: FilterParams) => {
  console.log('Начало handleFilterChange с фильтрами:', filters)
  isLoading.value = true
  error.value = null

  try {
    const response = await AuthService.getUsers(filters)
    console.log('Ответ от AuthService.getUsers:', response)

    if (response && response.result && Array.isArray(response.result)) {
      console.log('Полученные пользователи:', response.result)

      if (filters.labels && filters.labels.length > 0) {
        users.value = response.result.filter(user =>
          user.label && user.label.some(userLabel =>
            filters.labels!.includes(userLabel)
          )
        )
      } else {
        users.value = response.result
      }

      console.log('Установлены пользователи:', users.value)
    } else {
      error.value = 'Некорректный ответ от сервера'
      users.value = []
    }
  } catch (err) {
    console.error('Error in handleFilterChange:', err)
    error.value = 'Ошибка при загрузке пользователей'
    users.value = []
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.home-container {
  display: grid;
  grid-template-columns: auto 1fr;
  gap: 20px;
  padding: 20px;
  min-height: 100vh;
  margin-top: 80px;
}

.filter-sidebar {
  position: relative;
  height: calc(100vh - 100px);
  overflow: visible;
}

.users-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  margin-left: 22em;
}

.users-list {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  padding: 20px;

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

.no-results {
  text-align: center;
  padding: 2rem;
  background: #f8f9fa;
  border-radius: 8px;
  margin: 1rem;
  color: #6c757d;
}

.no-results h3 {
  margin-bottom: 0.5rem;
  font-size: 1.2rem;
}

.no-results p {
  margin: 0;
  font-size: 0.9rem;
}

@media (max-width: 1024px) {
  .users-list {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (max-width: 768px) {
  .users-list {
    grid-template-columns: 1fr;
  }
}
</style>
