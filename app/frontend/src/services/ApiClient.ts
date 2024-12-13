  import axios from "axios";

  // Константы для ключей localStorage
  const TOKEN_KEY = "token";
  const EMAIL_KEY = "email";

  // Создаём экземпляр клиента
  const apiClient = axios.create({
    baseURL: "http://localhost:8080", // Укажите базовый URL вашего API
  });

  // Добавляем перехватчик запросов
  apiClient.interceptors.request.use(
    (config) => {
      // Получаем токен из localStorage
      const token = localStorage.getItem(TOKEN_KEY);

      if (token) {
        // Если токен существует, добавляем его в заголовок Authorization
        config.headers.Authorization = `Bearer ${token}`;
      }

      console.log("Request:", config); // Логирование запроса
      return config; // Возвращаем обновлённый config
    },
    (error) => {
      console.error("Request error:", error); // Логирование ошибки запроса
      return Promise.reject(error);
    }
  );

  // Добавляем перехватчик ответов
  apiClient.interceptors.response.use(
    (response) => {
      console.log("Response:", response); // Логирование ответа
      return response;
    },
    (error) => {
      console.error("Response error:", error); // Логирование ошибки ответа

      // Если ответ имеет ошибку
      if (error.response?.status === 401) {
        // Если статус ошибки 401 (Unauthorized), разлогиниваем пользователя
        localStorage.removeItem(TOKEN_KEY); // Удаляем токен
        localStorage.removeItem(EMAIL_KEY); // Удаляем email пользователя
        location.reload(); // Перезагружаем страницу (возврат на страницу логина)
      }

      // Пропускаем ошибку дальше для обработки в компонентах
      return Promise.reject(error);
    }
  );

  // // Метод для проверки email
  // export const verifyEmail = (email: string) => {
  //   return apiClient.post("/api/v1/verify/email", { email });
  // };

  // // Метод для отправки данных профиля
  // export const completeProfile = (data: { username: string; email: string; password: string;  gender: string }) => {
  //   return apiClient.post("/api/v1/complete/profile", data);
  // };

  export default apiClient;
