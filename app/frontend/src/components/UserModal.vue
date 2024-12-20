<template>
  <!-- При клике на overlay вызывается closeModal -->
  <div v-if="show" class="modal-overlay" @click="closeModal">
    <!-- .stop предотвращает всплытие события клика к overlay -->
    <div class="modal-content" @click.stop>
      <div class="modal-header">
        <div class="avatar">
          <img :src="defaultAvatar" alt="avatar" />
        </div>
        <div class="user-info">
          <h2 class="nickname">
            {{ user.username }}
            <i class="fas" :class="genderIcon" :title="genderTitle"></i>
          </h2>
        </div>
      </div>

      <div class="user-details">
        <div v-if="user.biography" class="info-item">
          <p>{{ user.biography }}</p>
        </div>

        <div class="labels">
          <span v-for="label in user.label" :key="label" class="label-tag">
            {{ getInterestName(label) }}
          </span>
        </div>
      </div>

      <div class="modal-actions">
        <button class="friend-button" @click.stop="handleFriendRequest" :disabled="isSending">
          {{ buttonText }}
        </button>
        <button class="report-button" @click="handleReport">Жалоба</button>
      </div>

      <!-- Добавляем кнопку закрытия, которая также вызывает closeModal -->
      <button class="close-button" @click="closeModal">&times;</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import type { ProfileUser } from '../classes'
import defaultAvatar from '../assets/default-avatar.png'
import { interests } from '../config/interests'
import AuthService from '../services/AuthService'

const props = defineProps<{
  show: boolean
  user: ProfileUser
}>()

const emit = defineEmits<{
  close: []
  'match-success': []
}>()

const isSending = ref(false)

const buttonText = computed(() => {
  if (isSending.value) {
    return 'Обработка...'
  }
  return 'Дружить'
})

const handleFriendRequest = async () => {
  try {
    isSending.value = true
    await AuthService.sendFriendRequest(props.user.user_id)
    emit('match-success') // Вызываем анимацию духа
    closeModal()
  } catch (error) {
    console.error('Ошибка при отправке запроса:', error)
  } finally {
    isSending.value = false
  }
}

const handleReport = () => {
  // Логика отправки жалобы
  console.log('Отправка жалобы')
}

const getInterestName = (value: string): string => {
  const interest = interests.find(i => i.value === value);
  return interest ? interest.name : value;
};

// Этот метод вызывается при клике на overlay или кнопку закрытия
const closeModal = () => {
  emit('close')
}

const genderIcon = computed(() => ({
  'fa-mars': props.user.gender === 'male',
  'fa-venus': props.user.gender === 'female'
}))

const genderTitle = computed(() =>
  props.user.gender === 'male' ? 'Мужской' : 'Женский'
)
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: var(--secondary-color);
  padding: 2em;
  border-radius: var(--border-radius);
  width: 90%;
  max-width: 500px;
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.nickname {
  font-size: 30px;
  font-weight: bold;
  margin: 1em 0;
}

.modal-header {
  text-align: center;
}

.avatar img {
  width: 10em;
  height: 10em;
  border-radius: 50%;
  object-fit: cover;
}

.user-details {
  margin: 1em 0;
}

.info-item {
  border: 2px solid var(--primary-color);
  border-radius: var(--border-radius);
  outline: none;

  color: var(--primary-color);
  background-color: var(--secondary-color);

  padding: 15px 20px;
  font-size: 16px;
  font-weight: bold;

  width: 100%;
  text-align: center;
}

.labels {
  display: flex;
  justify-content: center;
  margin: 10px 0;
  gap: 0.5em;
  margin: 1em 0;
}

.modal-actions {
  display: flex;
  justify-content: space-between;
  gap: 1em;
  margin-top: 2em;
}

.friend-button,
.report-button {
  padding: 0.5em 2em;
  border-radius: var(--border-radius);
  border: none;
  cursor: pointer;
  font-size: 20px;
}

.friend-button {
  background: #b08d57;
  color: white;
  width: 10em;
}

.report-button {
  background: var(--primary-color);
  color: white;
  width: 10em;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  border: none;
  background: none;
  font-size: 1.5em;
  cursor: pointer;
}

.label-tag {
  background: var(--secondary-color);
  padding: 0.3em 0.7em;
  border-radius: var(--border-radius);
  border: 2px solid var(--primary-color);
}

.friend-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.user-info {
  display: flex;
  align-items: center;
  justify-content: center;
}

.nickname {
  display: flex;
  align-items: center;
  gap: 10px;
}

.nickname i {
  font-size: 0.8em;
  color: var(--primary-color);
}

.fa-mars {
  color: #4287f5;
}

.fa-venus {
  color: #f542aa;
}
</style>
