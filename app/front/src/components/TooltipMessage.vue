<template>
  <div v-if="visible" class="tooltip-container" :style="{ top: position.top + 'px', left: position.left + 'px' }">
    <div class="tooltip-content">
      {{ message }}
    </div>
    <div class="tooltip-arrow"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'TooltipMessage',
  props: {
    message: {
      type: String,
      required: true,
    },
    visible: {
      type: Boolean,
      required: true,
    },
    position: {
      type: Object,
      required: true,
      default: () => ({ top: 0, left: 0 }),
    },
  },
});
</script>

<style scoped>
.tooltip-container {
  position: absolute;
  z-index: 1000;
  animation: fadeIn 0.3s ease;
}

.tooltip-content {
  background-color: #333;
  color: white;
  padding: 8px 12px;
  border-radius: 4px;
  font-size: 14px;
  text-align: center;
  box-shadow: 0 2px 6px rgba(0, 0, 0, 0.2);
}

.tooltip-arrow {
  position: absolute;
  bottom: -8px;
  left: 50%;
  transform: translateX(-50%);
  width: 0;
  height: 0;
  border-left: 8px solid transparent;
  border-right: 8px solid transparent;
  border-top: 8px solid #333;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>
