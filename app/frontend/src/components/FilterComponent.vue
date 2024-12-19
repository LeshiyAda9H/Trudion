<template>
  <div class="filter-container">
    <h2>Фильтры</h2>

    <div class="filter-section">
      <h3>Количество пользователей</h3>
      <input
        type="number"
        v-model="usersNumber"
        @change="applyFilters"
        min="1"
        max="100"
        class="number-input"
      />
    </div>

    <div class="filter-section">
      <h3>Интересы</h3>
      <InterestSelector
        v-model="selectedLabels"
        placeholder="Выберите интересы для фильтрации"
      />
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
      resetFilters
    };
  }
});
</script>

<style scoped>
.filter-container {
  position: fixed;
  top: 100px;
  padding: 20px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
  width: 100%;
  max-width: 350px;
  height: fit-content;
  max-height: calc(100vh - 120px);
  overflow-y: auto;

}

.filter-section {
  margin: 15px 0;
}

.number-input {
  width: 100%;
  padding: 8px;
  border: 1px solid #ddd;
  border-radius: 5px;
}

.apply-filters {
  width: 100%;
  padding: 10px;
  background: #b08d57;
  color: white;
  border: none;
  border-radius: 5px;
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
  border-radius: 5px;
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


