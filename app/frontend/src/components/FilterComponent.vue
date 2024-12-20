<template>
  <div class="filter-container">
    <h2>Фильтры</h2>

    <div class="filter-section">
      <h3>Количество пользователей</h3>
      <div class="number-input-container">
        <button class="number-control" @click="decrementUsers">-</button>
        <input type="number" v-model="usersNumber" min="1" max="1000" class="number-input" />
        <button class="number-control" @click="incrementUsers">+</button>
      </div>
    </div>

    <div class="filter-section">
      <h3>Интересы</h3>
      <InterestSelector v-model="selectedLabels" placeholder="Выберите интересы" />
    </div>

    <div class="filter-buttons">
      <button class="apply-filters" @click="applyFilters">
        Применить фильтры
      </button>

      <button class="reset-filters" @click="resetFilters">
        Сбросить фильтры
      </button>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import InterestSelector from './InterestSelector.vue';

export default defineComponent({
  name: 'FilterComponent',
  components: {
    InterestSelector
  },
  emits: ['filter-changed'],

  setup(_, { emit }) {
    const usersNumber = ref(10);
    const selectedLabels = ref<string[]>([]);

    const incrementUsers = () => {
      if (usersNumber.value < 100) {
        usersNumber.value++;
      }
    };

    const decrementUsers = () => {
      if (usersNumber.value > 1) {
        usersNumber.value--;
      }
    };

    const applyFilters = () => {
      console.log('Применяем фильтры:', {
        usersNumber: usersNumber.value,
        labels: selectedLabels.value
      });

      emit('filter-changed', {
        usersNumber: usersNumber.value,
        labels: selectedLabels.value.length > 0 ? selectedLabels.value : undefined
      });
    };

    const resetFilters = () => {
      usersNumber.value = 10;
      selectedLabels.value = [];
      applyFilters();
    };

    return {
      usersNumber,
      selectedLabels,
      applyFilters,
      resetFilters,
      incrementUsers,
      decrementUsers
    };
  }
});
</script>

<style scoped>
.filter-container {
  position: fixed;
  top:5em;
  padding: 1.5em 20px;
  background: #fff;
  border-radius: 10px;
  width: 100%;
  max-width: 350px;
  height: fit-content;
  max-height: calc(100vh - 120px);
  min-height: calc(100vh - 10px);
  overflow-y: auto;

}

.filter-section {
  margin: 1em 0;
}

.number-input-container {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
  margin-top: 10px;

}

.number-input {
  width: 13em;
  text-align: center;
  padding: 8px 1em;
  border: 2px solid var(--primary-color);
  border-radius: var(--border-radius);
  font-size: 16px;
  color: black;
  background: var(--secondary-color);
  margin: 1em 0;
}

.number-input::-webkit-inner-spin-button,
.number-input::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;

}

.number-control {
  width: 36px;
  height: 36px;
  border: 2px solid var(--primary-color);
  background: var(--secondary-color);
  color: black;
  font-size: 20px;
  border-radius: var(--border-radius);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
}

.number-control:hover {
  background: var(--primary-color);
  color: white;
}

.number-control:active {
  transform: scale(0.95);
}

.apply-filters {
  width: 100%;
  padding: 10px;
  background: #b08d57;
  color: white;
  border: none;
  border-radius: var(--border-radius);
  margin-top: 20px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.apply-filters:hover {
  background: #97784a;
}

.reset-filters {
  width: 100%;
  padding: 10px;
  background: #f0f0f0;
  color: #333;
  border: 1px solid #ddd;
  border-radius: var(--border-radius);
  margin-top: 10px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.reset-filters:hover {
  background: #e0e0e0;
}

.filter-container::-webkit-scrollbar {
  width: 8px;
}

.filter-container::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

.filter-container::-webkit-scrollbar-thumb {
  background: #b08d57;
  border-radius: 4px;
}

.filter-container::-webkit-scrollbar-thumb:hover {
  background: #97784a;
}

.filter-buttons {
  display: flex;
  flex-direction: column;
  margin-top: 15em;
}
</style>
