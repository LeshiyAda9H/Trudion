<template>
  <div class="multiselect-wrapper">
    <!-- Отображение выбранных меток -->
    <div class="selected-labels" v-if="modelValue.length > 0">
      <div v-for="value in modelValue"
           :key="value"
           class="label-tag">
        <span>{{ getInterestName(value) }}</span>
        <button class="remove-label" @click.stop="removeInterest(value)">×</button>
      </div>
    </div>

    <div class="multiselect-container">
      <div class="multiselect-header" @click="isOpen = !isOpen">
        <div class="selected-display">
          <span v-if="modelValue.length === 0">{{ placeholder }}</span>
          <span v-else>Выбрано: {{ modelValue.length }}</span>
        </div>
        <div class="arrow" :class="{ 'open': isOpen }">▼</div>
      </div>

      <div v-show="isOpen" class="options-container">
        <div
          v-for="interest in interests"
          :key="interest.id"
          :class="['option-item', { 'selected': modelValue.includes(interest.value) }]"
          @click="toggleInterest(interest.value)"
        >
          <input
            type="checkbox"
            :checked="modelValue.includes(interest.value)"
            @click.stop
          >
          <span>{{ interest.name }}</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { interests } from '../config/interests';

export default defineComponent({
  name: 'InterestSelector',
  props: {
    modelValue: {
      type: Array as () => string[],
      required: true
    },
    placeholder: {
      type: String,
      default: 'Выберите интересы'
    }
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const isOpen = ref(false);

    const getInterestName = (value: string): string => {
      const interest = interests.find(i => i.value === value);
      return interest ? interest.name : value;
    };

    const toggleInterest = (value: string) => {
      const newInterests = [...props.modelValue];
      const index = newInterests.indexOf(value);

      if (index === -1) {
        newInterests.push(value);
      } else {
        newInterests.splice(index, 1);
      }

      emit('update:modelValue', newInterests);
    };

    const removeInterest = (value: string) => {
      const newInterests = props.modelValue.filter(v => v !== value);
      emit('update:modelValue', newInterests);
    };

    return {
      interests,
      isOpen,
      toggleInterest,
      removeInterest,
      getInterestName
    };
  }
});
</script>

<style scoped>
.multiselect-wrapper {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.selected-labels {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.label-tag {
  display: flex;
  align-items: center;
  background-color: #e3f2fd;
  padding: 4px 8px;
  border-radius: 16px;
  font-size: 14px;
  gap: 4px;
}

.remove-label {
  background: none;
  border: none;
  color: #666;
  font-size: 18px;
  cursor: pointer;
  padding: 0 2px;
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
}

.remove-label:hover {
  color: #333;
}

.multiselect-container {
  position: relative;
  width: 100%;
  max-width: 300px;
}

.multiselect-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
  cursor: pointer;
  user-select: none;
}

.arrow {
  transition: transform 0.2s ease;
}

.arrow.open {
  transform: rotate(180deg);
}

.options-container {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  max-height: 250px;
  overflow-y: auto;
  background: white;
  border: 1px solid #ddd;
  border-top: none;
  border-radius: 0 0 4px 4px;
  z-index: 1000;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.option-item {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 8px 15px;
  cursor: pointer;
  transition: background-color 0.2s ease;
}

.option-item:hover {
  background-color: #f5f5f5;
}

.option-item.selected {
  background-color: #e3f2fd;
}

.option-item input[type="checkbox"] {
  margin: 0;
}

/* Стилизация скролла */
.options-container::-webkit-scrollbar {
  width: 8px;
}

.options-container::-webkit-scrollbar-track {
  background: #f1f1f1;
}

.options-container::-webkit-scrollbar-thumb {
  background: #888;
  border-radius: 4px;
}

.options-container::-webkit-scrollbar-thumb:hover {
  background: #555;
}
</style>
