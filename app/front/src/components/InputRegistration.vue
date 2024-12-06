<template>
  <input
    type="email"
    :class="
    error === 'not-valid-email-and-passwords' || error === 'not-valid-email'
    ? 'input-field-error'
    : 'input-field'
    "
    v-on:input="writeEmail(($event.target as HTMLInputElement).value)"
    :placeholder="
    error === 'not-valid-email-and-passwords' || error === 'not-valid-email'
    ? 'Введите электронную почту'
    : 'Электронная почта'"
    required
  />

  <!-- Поле ввода пароля -->
  <input
    type="password"
    :class="
    error === 'not-valid-email-and-passwords' || error === 'passwords-dont-match'
    ? 'input-field-error'
    : 'input-field'"

    v-on:input="writePass(($event.target as HTMLInputElement).value)"
    :placeholder="
    error === 'not-valid-email-and-passwords' || error === 'passwords-dont-match'
    ? 'Введите пароль'
    : 'Пароль'"
    required
  />

  <!-- Поле подтвержения пароля -->
  <input
  type="password"
    :class="
    error === 'not-valid-email-and-passwords' || error === 'passwords-dont-match'
    ? 'input-field-error'
    : 'input-field'"

    v-on:input="writeConfirmPass(($event.target as HTMLInputElement).value)"
    :placeholder="
    error === 'not-valid-email-and-passwords' || error === 'passwords-dont-match'
    ? 'Введите пароль'
    : 'Пароль'"
    required
  />
</template>

<script lang="ts">
import type { PropType } from 'vue';

export default {
  props: {
    writeEmail: {
      type: Function as PropType<(text_email: string) => void>,
      required: true,
    },
    writePass: {
      type: Function as PropType<(text_pass: string) => void>,
      required: true,
    },
    writeConfirmPass: {
      type: Function as PropType<(text_confirmPass: string) => void>,
      required: true,
    },
    error: {
      type: String,
      required: true,
    },
  },
};
</script>

<style scoped>
/* Базовый стиль для поля ввода */
.input-field {
  outline: none; /* Убирает стандартную обводку */
  width: 53%; /* Ширина поля */
  border: 2px solid #355299; /* Синяя рамка */
  border-radius: 25px; /* Закругленные углы */
  padding: 3% 8.5%; /* Внутренние отступы */
  margin: 5px 0; /* Внешние отступы */
  color: #355299; /* Цвет текста */
  font-size: 16px; /* Размер шрифта */
  font-weight: bold; /* Жирный текст */
  background-color: #bdcdf6; /* Фоновый цвет */
}

/* Стиль для текста-заполнителя */
.input-field::placeholder {
  color: #355299; /* Цвет текста-заполнителя */
  font-weight: normal; /* Обычный текст */
}

/* Стиль для поля с ошибкой */
.input-field-error {
  outline: none; /* Убирает стандартную обводку */
  width: 53%; /* Ширина поля */
  border: 2px solid red; /* Красная рамка */
  border-radius: 25px; /* Закругленные углы */
  padding: 3% 8.5%; /* Внутренние отступы */
  margin: 5px 0; /* Внешние отступы */
  color: #b18888; /* Цвет текста */
  font-size: 16px; /* Размер шрифта */
  font-weight: bold; /* Жирный текст */
  background-color: #bdcdf6; /* Фоновый цвет */
}

/* Стиль для текста-заполнителя поля с ошибкой */
.input-field-error::placeholder {
  color: #b18888; /* Цвет текста-заполнителя */
  font-weight: bold; /* Жирный текст */
}
</style>
