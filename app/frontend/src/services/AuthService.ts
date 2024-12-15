// import type { AxiosInstance } from 'axios'
import type { User, LoginPayload, AuthedUser, ProfileUser } from '../classes'
import apiClient from './ApiClient'
//import axios from 'axios'

// Константы для URL-путей
const REGISTER_URL = '/api/v1/register'
const LOGIN_URL = '/api/v1/login'
const LOGOUT_URL = '/api/v1/logout'
const PROFILE_URL = 'api/v1/profile'

// const VERIFY_TOKEN_URL = '/api/v1/verify/token'
const VERIFY_EMAIL_URL = '/api/v1/verify/email'

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
    }
    catch (error) {
      console.error('Registration error:', error)
      throw new Error('Ошибка при регистрации.')
    }
  }

  // Метод для входа пользователя в систему
  async login(data: LoginPayload): Promise<AuthedUser> {
    try {
      const response = await this.api.post<AuthedUser, LoginPayload>(LOGIN_URL, data)
      const token = response.access_token

      const authedUser: AuthedUser = {
        email: data.email,
        access_token: token,
      }

      localStorage.setItem('token', token)
      console.log('User logged in successfully')
      return authedUser
    }
    catch (error) {
      console.error('Login error:', error)
      throw new Error('Ошибка при входе.')
    }
  }

  // Метод для обновления профиля пользователя
  async updateProfile(data: Partial<ProfileUser>): Promise<void> {
    try {
      await this.api.post<void, Partial<ProfileUser>>(PROFILE_URL, data)
      console.log('Профиль обновлен успешно')
    }
    catch (error) {
      console.error('Ошибка при обновлении профиля:', error)
      throw new Error('Не удалось обновить профиль')
    }
  }

  // Метод для получения профиля пользователя
  async getProfile(): Promise<ProfileUser> {
    try {
      return await this.api.get<ProfileUser>(PROFILE_URL)
    }
    catch (error) {
      throw error
    }
  }

  // Метод для проверки доступности email
  async verifyEmail(email: string): Promise<boolean> {
    try {
      return await this.api.post<boolean, { email: string }>(VERIFY_EMAIL_URL, { email })
    }
    catch (error) {
      console.error('Ошибка проверки email:', error)
      throw new Error('Такой email уже зарегистрирован.')
    }
  }

  // Метод для выхода пользователя из системы
  logout(): void {
    try {
      this.api.post<void>(LOGOUT_URL)
    }
    catch (error) {
      console.error('Error logging out on server:', error)
    }
    finally {
      localStorage.removeItem('token')
      console.log('User logged out successfully')
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
    }
    catch (error) {
      console.error('Error completing profile:', error)
      throw new Error('Ошибка завершения регистрации.')
    }
  }

  // Метод для получения списка пользователей 
  async getUsers(params: { usersNumber: number, labels?: string[] }) {
    return await this.api.get('/api/v1/users', { params });
  }
}

// Экспортируем экземпляр класса AuthService
export default new AuthService()
