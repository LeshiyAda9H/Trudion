<template>
  <div class="work-spirit" :class="{ 'flying': isFlying, 'match-animation': isMatch }">
    <div class="spirit-container">
      <img src="../assets/trudion.svg" alt="Work Spirit" class="spirit-image" />
      <div class="sparkles">
        <i class="sparkle"></i>
        <i class="sparkle"></i>
        <i class="sparkle"></i>
      </div>
    </div>
    <div class="message-bubble" v-if="showMessage">
      {{ message }}
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const isFlying = ref(false);
const showMessage = ref(false);
const message = ref('');
const isMatch = ref(false);

const regularMessages = [
  'Вместе к успеху! 💪',
  'Труд объединяет! ⚡',
  'Найдём общее дело! 🌟',
  'Работа в радость! ✨'
];

const matchMessages = [
  'Отличное совпадение! 🤝',
  'Вы нашли друг друга! ✨',
  'Начало сотрудничества! 🌟',
  'Вместе к новым высотам! 🚀'
];

const flyIn = (matchSuccess = false) => {
  isFlying.value = true;
  isMatch.value = matchSuccess;

  setTimeout(() => {
    showMessage.value = true;
    message.value = matchSuccess
      ? 'Поздравляем с совпадением! 🎉 Начните общение прямо сейчас!'
      : regularMessages[Math.floor(Math.random() * regularMessages.length)];
  }, 1000);

  setTimeout(() => {
    showMessage.value = false;
    setTimeout(() => {
      isFlying.value = false;
      isMatch.value = false;
    }, 500);
  }, 4000);
};

defineExpose({ flyIn });
</script>

<style scoped>
.work-spirit {
  position: fixed;
  right: -100px;
  bottom: 50px;
  transition: all 0.5s ease;
  z-index: 1000;
}

.work-spirit.flying {
  right: 50px;
}

.work-spirit.match-animation {
  right: 50%;
  bottom: 50%;
  transform: translate(50%, 50%);
}

.spirit-container {
  position: relative;
  width: 100px;
  height: 100px;
}

.spirit-image {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.message-bubble {
  position: absolute;
  top: -60px;
  left: 50%;
  transform: translateX(-50%);
  background: white;
  padding: 10px 15px;
  border-radius: 15px;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1);
  white-space: nowrap;
  animation: popIn 0.3s ease-out;
}

.message-bubble::after {
  content: '';
  position: absolute;
  bottom: -10px;
  left: 50%;
  transform: translateX(-50%);
  border-left: 10px solid transparent;
  border-right: 10px solid transparent;
  border-top: 10px solid white;
}

.sparkles {
  position: absolute;
  width: 100%;
  height: 100%;
}

.sparkle {
  position: absolute;
  width: 8px;
  height: 8px;
  background: var(--primary-color);
  border-radius: 50%;
  opacity: 0;
}

.sparkle:nth-child(1) {
  top: 20%;
  left: 20%;
  animation: sparkle 2s ease-in-out infinite;
}

.sparkle:nth-child(2) {
  top: 60%;
  right: 20%;
  animation: sparkle 2s ease-in-out infinite 0.4s;
}

.sparkle:nth-child(3) {
  bottom: 20%;
  left: 50%;
  animation: sparkle 2s ease-in-out infinite 0.8s;
}

@keyframes float {

  0%,
  100% {
    transform: translateY(0);
  }

  50% {
    transform: translateY(-15px);
  }
}

@keyframes sparkle {

  0%,
  100% {
    opacity: 0;
    transform: scale(0);
  }

  50% {
    opacity: 1;
    transform: scale(1);
  }
}

@keyframes popIn {
  from {
    opacity: 0;
    transform: translateX(-50%) scale(0.8);
  }

  to {
    opacity: 1;
    transform: translateX(-50%) scale(1);
  }
}
</style>
