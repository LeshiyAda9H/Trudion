import axios from "axios";

// Создаём экземпляр клиента
const apiClient = axios.create({
  baseURL: "http://localhost:8080", // Укажите базовый URL вашего API
});

// Добавляем перехватчик запросов
apiClient.interceptors.request.use(
  (config) => {
    // Получаем токен из localStorage
    const token = localStorage.getItem("token");

    if (token) {
      // Если токен существует, добавляем его в заголовок Authorization
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config; // Возвращаем обновлённый config
  },
  (error) => {
    // Обрабатываем ошибку конфигурации запроса
    return Promise.reject(error);
  }
);

// Добавляем перехватчик ответов
apiClient.interceptors.response.use(
  (response) => {
    // Если запрос успешен, возвращаем ответ
    return response;
  },
  (error) => {
    // Если ответ имеет ошибку
    if (error.response?.status === 401) {
      // Если статус ошибки 401 (Unauthorized), разлогиниваем пользователя
      localStorage.removeItem("token"); // Удаляем токен
      localStorage.removeItem("username"); // Удаляем имя пользователя
      location.reload(); // Перезагружаем страницу (возврат на страницу логина)
    }

    // Пропускаем ошибку дальше для обработки в компонентах
    return Promise.reject(error);
  }
);

export default apiClient;
