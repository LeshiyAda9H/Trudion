<template>
  <div class="filter-container">
    <h2>Фильтры</h2>

    <div class="filter-section">
      <h3>Количество пользователей</h3>
      <input
        type="number"
        v-model="usersNumber"
        @change="updateFilters"
        min="1"
        max="100"
      />
    </div>

    <div class="filter-section">
      <h3>Интересы</h3>
      <div class="skills-list">
        <div
          v-for="label in availableLabels"
          :key="label.value"
          :class="['skill-item', { active: selectedLabels.includes(label.value) }]"
          @click="toggleLabel(label.value)"
        >
          {{ label.name }}
        </div>
      </div>
    </div>

    <button class="apply-filters" @click="applyFilters">
      Применить фильтры
    </button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
  name: 'FilterComponent',
  emits: ['filter-changed'],
  setup(_, { emit }) {
    const usersNumber = ref(10);
    const selectedLabels = ref<string[]>([]);

    const availableLabels = [
      { name: 'Программирование', value: 'programming' },
      { name: 'Иностранные языки', value: 'languages' },
      { name: 'Спорт', value: 'sports' },
      // ... другие метки
    ];

    const toggleLabel = (value: string) => {
      const index = selectedLabels.value.indexOf(value);
      if (index === -1) {
        selectedLabels.value.push(value);
      } else {
        selectedLabels.value.splice(index, 1);
      }
    };

    const applyFilters = () => {
      emit('filter-changed', {
        usersNumber: usersNumber.value,
        labels: selectedLabels.value
      });
    };

    return {
      usersNumber,
      availableLabels,
      selectedLabels,
      toggleLabel,
      applyFilters
    };
  }
});
</script>

<style scoped>
.filter-container {
  padding: 20px;
  background: #fff;
  border-radius: 10px;
  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
}

.filter-section {
  margin: 15px 0;
}

.skills-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.skill-item {
  padding: 8px;
  border-radius: 5px;
  cursor: pointer;
  border: 1px solid #ddd;
}

.skill-item.active {
  background: #b08d57;
  color: white;
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
}
</style>
