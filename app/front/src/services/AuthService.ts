import axios from 'axios' // Импортируем библиотеку Axios для выполнения HTTP-запросов
import type { AxiosInstance } from 'axios' // Импортируем тип AxiosInstance
import type { User, AuthedUser } from '../classes' // Импортируем типы данных User и AuthedUser
import apiClient from './ApiClient' // Импортируем экземпляр Axios для взаимодействия с API
import StorageService from './StorageService' // Импортируем сервис для работы с хранилищем

// Константы для URL-путей
const REGISTER_URL = '/register';
const LOGIN_URL = '/login';
const TOKEN_VERIFY_URL = '/token/verify';

// Класс AuthService для управления аутентификацией и регистрацией пользователей
class AuthService {
  private api: AxiosInstance // Приватное свойство для хранения экземпляра Axios

  // Конструктор класса, инициализирует экземпляр Axios
  constructor() {
    this.api = apiClient
  }

  // Метод для регистрации нового пользователя
  async register(data: User): Promise<void> {
    try {
      await this.api.post(REGISTER_URL, data) // Отправляем POST-запрос на регистрацию
      console.log("User registered successfully"); // Логирование успешной регистрации
    }
    catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 400) {
          throw new Error('Пользователь с таким email уже существует.') // Обрабатываем ошибку 400
        }
        console.error("Registration error:", error); // Логирование ошибки регистрации
      }
      throw new Error('Ошибка при регистрации.') // Обрабатываем другие ошибки
    }
  }

  // Метод для входа пользователя в систему
  async login(data: User): Promise<AuthedUser> {
    try {
      const response = await this.api.post(LOGIN_URL, data) // Отправляем POST-запрос на вход
      const token = response.data.access_token // Извлекаем токен из ответа

      const authedUser: AuthedUser = {
        email: data.email,
        access_token: token,
      }

      StorageService.setToken(token) // Сохраняем токен в хранилище
      StorageService.setEmail(authedUser.email) // Сохраняем email в хранилище

      console.log("User logged in successfully"); // Логирование успешного входа
      return authedUser // Возвращаем аутентифицированного пользователя
    }
    catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        if (error.response?.status === 401) {
          throw new Error('Неверный логин или пароль.') // Обрабатываем ошибку 401
        }
        console.error("Login error:", error); // Логирование ошибки входа
      }
      throw new Error('Ошибка при входе.') // Обрабатываем другие ошибки
    }
  }

  // Метод для проверки валидности токена
  async verifyToken(): Promise<boolean> {
    try {
      const token = StorageService.getToken() // Получаем токен из хранилища
      if (!token) return false // Если токена нет, возвращаем false

      await this.api.post(TOKEN_VERIFY_URL, { token }) // Отправляем POST-запрос на проверку токена
      console.log("Token verified successfully"); // Логирование успешной проверки токена
      return true // Если токен валиден, возвращаем true
    }
    catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        StorageService.clear() // Очищаем хранилище при ошибке
        console.error("Token verification error:", error); // Логирование ошибки проверки токена
      }
      return false // Возвращаем false при ошибке
    }
  }

  // Метод для выхода пользователя из системы
  logout(): void {
    StorageService.clear() // Очистка всех данных в хранилище
    console.log("User logged out successfully"); // Логирование успешного выхода
  }

  // Метод для проверки, аутентифицирован ли пользователь
  isAuth(): boolean {
    return !!StorageService.getToken() // Проверка наличия токена в хранилище
  }
}

// Экспортируем экземпляр класса AuthService
export default new AuthService()
