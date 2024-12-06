import axios from 'axios'
import type { AxiosInstance } from 'axios'
import type { User, AuthedUser } from '../classes'
import apiClient from './ApiClient'
import StorageService from './StorageService' // Импортируем наш сервис


class AuthService {
  private api: AxiosInstance

  constructor() {
    this.api = apiClient
  }

  async register(data: User): Promise<void> {
    try {
      await this.api.post('/register', data)
    } catch (error: unknown) {
      if (axios.isAxiosError(error) && error.response?.status === 400) {
        throw new Error('Пользователь с таким email уже существует.')
      }
      throw new Error('Ошибка при регистрации.')
    }
  }

  async login(data: User): Promise<AuthedUser> {
    try {
      const response = await this.api.post('/login', data)
      const token = response.data.access_token

      const authedUser: AuthedUser = {
        username: data.username,
        access_token: token,
      }

      StorageService.setToken(token)
      StorageService.setUsername(authedUser.username)

      return authedUser
    } catch (error: unknown) {
      if (axios.isAxiosError(error) && error.response?.status === 401) {
        throw new Error('Неверный логин или пароль.')
      }
      throw new Error('Ошибка при входе.')
    }
  }

  async verifyToken(): Promise<boolean> {
    try {
      const token = StorageService.getToken()
      if (!token) return false

      await this.api.post('/api/v1/token/verify', { token })
      return true
    } catch (error: unknown) {
      if (axios.isAxiosError(error)) {
        StorageService.clear()
      }
      return false
    }
  }

  logout(): void {
    StorageService.clear() // Очистка всех данных
  }

  isAuth(): boolean {
    return !!StorageService.getToken() // Проверка наличия токена
  }
}

export default new AuthService()
