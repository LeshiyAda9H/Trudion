<template>
  <!-- Поле ввода email -->
  <input
    type="email"
    :class="
      // Если ошибка связана с email или с обоими полями, применяем класс с ошибкой
      error === 'not-valid-email-and-password' || error === 'not-valid-email'
        ? 'input-error'
        : 'input'
    "
    v-on:input="writeEmail(($event.target as HTMLInputElement).value)"
    :placeholder="
      // Если ошибка связана с email, показываем текст ошибки
      error === 'not-valid-email-and-password' || error === 'not-valid-email'
        ? 'Введите электронную почту'
        : 'Электронная почта'
    "
    required
  />

  <!-- Поле ввода пароля -->
  <input
    type="password"
    :class="
      // Если ошибка связана с паролем или с обоими полями, применяем класс с ошибкой
      error === 'not-valid-email-and-password' || error === 'not-valid-password'
        ? 'input-error'
        : 'input'
    "
    v-on:input="writePass(($event.target as HTMLInputElement).value)"
    :placeholder="
      // Если ошибка связана с паролем, показываем текст ошибки
      error === 'not-valid-email-and-password' || error === 'not-valid-password'
        ? 'Введите пароль'
        : 'Пароль'
    "
    required
  />
</template>

<script lang="ts">
import type { PropType } from 'vue'; // Импортируем PropType как тип для проверки пропсов

export default {
  props: {
    // Функция для обработки ввода email
    writeEmail: {
      type: Function as PropType<(text_email: string) => void>,
      required: true, // Проп обязателен
    },
    // Функция для обработки ввода пароля
    writePass: {
      type: Function as PropType<(text_pass: string) => void>,
      required: true, // Проп обязателен
    },
    // Строка, описывающая текущую ошибку
    error: {
      type: String,
      required: true, // Проп обязателен
    },
  },
};
</script>


<style scoped>

.input {
  width: 20em;
  margin: 5px 0;
}
.input-error{
  width: 20em;
  margin: 5px 0;
  color: red;
}

</style>
