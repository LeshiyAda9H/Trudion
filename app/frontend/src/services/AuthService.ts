// import type { AxiosInstance } from 'axios'
import type { AuthedUser, LoginPayload, ProfileUser, User } from '../classes'
import apiClient from './ApiClient' //import axios from 'axios'
//import axios from 'axios'

// Константы для URL-путей
const REGISTER_URL = '/api/v1/register'
const LOGIN_URL = '/api/v1/login'
const LOGOUT_URL = '/api/v1/logout'
const PROFILE_URL = 'api/v1/profile'

// const VERIFY_TOKEN_URL = '/api/v1/verify/token'
const VERIFY_EMAIL_URL = '/api/v1/verify/email'

interface verifyProps {
  available: boolean
}

// Класс AuthService для управления аутентификацией и регистрацией пользователей
class AuthService {
  private api: typeof apiClient

  // Конструктор класса, инициализирует экземпляр Axios
  constructor() {
    this.api = apiClient
  }

  // Метод для регистрации нового пользователя
  async register(data: Partial<User>): Promise<void> {
    try {
      await this.api.post<void, Partial<User>>(REGISTER_URL, data)
      console.log('User registered successfully')
    } catch (error) {
      console.error('Registration error:', error)
      throw new Error('Ошибка при регистрации.')
    }
  }

  // Метод для входа пользователя в систему
  async login(data: LoginPayload): Promise<AuthedUser> {
    try {
      const response = await this.api.post<{ token: string }>(LOGIN_URL, data)
      console.log('Ответ сервера при логине:', response)

      if (!response || !response.token) {
        console.error('Неверная структура ответа:', response)
        throw new Error('Неверный формат ответа от сервера')
      }

      return {
        email: data.email,
        token: response.token, // Используем token вместо access_token
      }
    } catch (error) {
      console.error('Login error:', error)
      throw new Error('Ошибка при входе.')
    }
  }

  // Метод для обновления профиля пользователя
  async updateProfile(data: Partial<ProfileUser>): Promise<void> {
    try {
      await this.api.post<void, Partial<ProfileUser>>(PROFILE_URL, data)
      console.log('Профиль обновлен успешно')
    } catch (error) {
      console.error('Ошибка при обновлении профиля:', error)
      throw new Error('Не удалось обновить профиль')
    }
  }

  // Метод для получения профиля пользователя
  async getProfile(): Promise<ProfileUser> {
    try {
      return await this.api.get<ProfileUser>(PROFILE_URL)
    } catch (error) {
      throw error
    }
  }

  // Метод для проверки доступности email
  async verifyEmail(email: string): Promise<verifyProps> {
    try {
      return await this.api.post<verifyProps, { email: string }>(VERIFY_EMAIL_URL, { email })
    } catch (error) {
      console.error('Ошибка проверки email:', error)
      throw new Error('Такой email уже зарегистрирован.')
    }
  }

  // Метод для выхода пользователя из системы
  async logout(): Promise<void> {
    try {
      await this.api.post<void>(LOGOUT_URL)
    } catch (error) {
      console.error('Error logging out:', error)
    } finally {
      localStorage.removeItem('token')
    }
  }

  // Метод для проверки, аутентифицирован ли пользователь
  isAuth(): boolean {
    return !!localStorage.getItem('token')
  }

  // Метод для завершения регистрации пользователя
  async completeProfile(data: Partial<User>): Promise<void> {
    try {
      await this.api.post<void, Partial<User>>(REGISTER_URL, data)
      console.log('Profile completed successfully')
    } catch (error) {
      console.error('Error completing profile:', error)
      throw new Error('Ошибка завершения регистрации.')
    }
  }

  // Метод для получения списка пользователей
  async getUsers(params: { usersNumber: number; labels?: string[] }) {
    try {
      console.log('Отправляем запрос с параметрами:', params)

      // Формируем правильные параметры запроса
      const queryParams = {
        usersnumber: params.usersNumber,
        // Преобразуем массив меток в строку, если они есть
        labels: params.labels ? params.labels.join(',') : undefined,
      }

      console.log('Подготовленные параметры запроса:', queryParams)

      const response = await this.api.get<{ result: ProfileUser[] }>('/api/v1/usersnumber', {
        params: queryParams,
      })

      console.log('Ответ от сервера:', response)

      return response
    } catch (error) {
      console.error('Error fetching users:', error)
      throw error
    }
  }
}

// Экспортируем экземпляр класса AuthService
export default new AuthService()
