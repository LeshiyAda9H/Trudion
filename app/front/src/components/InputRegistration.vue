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
    ? 'Подтвердите пароль'
    : 'Подтвердите пароль'"
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



