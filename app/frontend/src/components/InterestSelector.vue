<template>
  <div class="multiselect-wrapper">


    <div class="multiselect-container">
      <div class="multiselect-header" @click="toggleDropdown">
        <div class="selected-display">
          <span v-if="modelValue.length === 0">{{ placeholder }}</span>
          <span v-else>Выбрано: {{ modelValue.length }}</span>
        </div>
        <div class="arrow" :class="{ 'open': isOpen }">▼</div>
      </div>

      <div v-show="isOpen" class="options-container" :class="{ 'direction-up': direction === 'up' }">
        <div v-for="category in categories" :key="category" class="category-section">
          <div class="category-header">{{ category }}</div>
          <div v-for="interest in getInterestsByCategory(category)" :key="interest.id" class="option-item"
            :class="{ 'selected': isSelected(interest.value) }" @click="toggleInterest(interest.value)">
            <div class="checkbox-container">
              <input type="checkbox" :checked="isSelected(interest.value)" @click.stop readonly>
              <span>{{ interest.name }}</span>
            </div>

          </div>
        </div>
      </div>
    </div>
    <!-- Отображение выбранных меток -->
    <div class="selected-labels" v-if="modelValue.length > 0">
      <div v-for="value in modelValue" :key="value" class="label-tag">
        <span>{{ getInterestName(value) }}</span>
        <button class="remove-label" @click="removeInterest(value)" type="button">×</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, onUnmounted } from 'vue';
import { interests, getCategories, getInterestsByCategory } from '../config/interests';

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
    },
    direction: {
      type: String,
      default: 'down'
    }
  },
  emits: ['update:modelValue'],
  setup(props, { emit }) {
    const isOpen = ref(false);
    const categories = getCategories();

    const getInterestName = (value: string): string => {
      const interest = interests.find(i => i.value === value);
      return interest ? interest.name : value;
    };

    const isSelected = (value: string): boolean => {
      return props.modelValue.includes(value);
    };

    const toggleDropdown = () => {
      isOpen.value = !isOpen.value;
    };

    const toggleInterest = (value: string) => {
      const newInterests = [...props.modelValue];
      const index = newInterests.indexOf(value);

      if (index === -1) {
        if (newInterests.length >= 3) {
          alert('Можно выбрать максимум 3 интереса');
          return;
        }
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

    // Закрытие дропдауна при клике вне компонента
    const handleClickOutside = (event: MouseEvent) => {
      const target = event.target as HTMLElement;
      if (!target.closest('.multiselect-wrapper')) {
        isOpen.value = false;
      }
    };

    // Добавляем и удаляем обработчик при монтировании/размонтировании
    onMounted(() => {
      document.addEventListener('click', handleClickOutside);
    });

    onUnmounted(() => {
      document.removeEventListener('click', handleClickOutside);
    });

    return {
      categories,
      interests,
      isOpen,
      getInterestName,
      isSelected,
      toggleDropdown,
      toggleInterest,
      removeInterest,
      getInterestsByCategory
    };
  }
});
</script>

<style scoped>
.multiselect-wrapper {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 100%;
  padding: 0 10px;
}

.selected-labels {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  
}

.label-tag {
  display: flex;
  align-items: center;
  gap: 4px;

  background: var(--secondary-color);
  padding: 4px 8px;
  border-radius: var(--border-radius);
  border: 2px solid var(--primary-color);
  font-size: 14px;

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
  align-self: center;
}

.multiselect-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 15px;
  background: var(--secondary-color);
  cursor: pointer;
  user-select: none;
  border: 2px solid var(--primary-color);
  border-radius: var(--border-radius);
  width: 20em;
  margin: 1em 0;
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
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.options-container.direction-up {
  top: auto;
  bottom: 100%;
  border-top: 1px solid #ddd;
  border-bottom: none;
  border-radius: 4px 4px 0 0;
  box-shadow: 0 -2px 4px rgba(0, 0, 0, 0.1);
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
  cursor: pointer;
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

.category-section {
  border-bottom: 1px solid #eee;
  padding: 8px 0;
}

.category-header {
  padding: 8px 15px;
  font-weight: bold;
  color: #666;
  background-color: #f8f8f8;
}

.checkbox-container {
  display: flex;
  align-items: center;
  gap: 10px;
}
</style>
