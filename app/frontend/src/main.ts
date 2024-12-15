import './assets/css/variables.css' // Переменные
import './assets/css/base.css' // Глобальные стили
import './assets/css/main.css'
import '@fortawesome/fontawesome-free/css/all.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

createApp(App).use(router).use(createPinia()).mount('#app')
