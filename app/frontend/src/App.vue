<script setup lang="ts">
import { RouterView } from 'vue-router'
import PageTransition from './components/PageTransition.vue'
</script>

<template>
  <PageTransition>
    <RouterView />
  </PageTransition>
</template>

<style scoped>
/* Предотвращаем скролл во время анимации */
.page-enter-active,
.page-leave-active {
  position: absolute;
  width: 100%;
  transition: all 0.5s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateY(30px);
}

.page-leave-to {
  opacity: 0;
  transform: translateY(-30px);
}
</style>
