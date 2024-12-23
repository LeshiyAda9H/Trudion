// import type { AxiosInstance } from 'axios'
import type {
  AuthedUser,
  LoginPayload,
  ProfileUser,
  User,
  ChatsResponse,
  MatchResponse,
  NotificationsResponse,
} from '../classes'
import apiClient from './ApiClient' //import axios from 'axios'
//import axios from 'axios'

// Константы для URL-путей
const VERIFY_EMAIL_URL = '/api/v1/verify/email'
const REGISTER_URL = '/api/v1/register'
const LOGIN_URL = '/api/v1/login'
const LOGOUT_URL = '/api/v1/logout'
const PROFILE_URL = '/api/v1/profile'

const HANDSHAKE_URL = '/api/v1/handshake'
const MATCHES_URL = '/api/v1/matches'
// const DIALOGS_URL = '/api/v1/dialogs'
const CHATS_URL = '/api/v1/chat'
const NOTIFICATIONS_URL = '/api/v1/notifications'

interface verifyProps {
  available: boolean
}

interface HandshakeResponse {
  message: 'mutually' | 'handshake'
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
      console.log('AuthService: отправляем PATCH запрос:', data) // добавим лог
      await this.api.patch<void, Partial<ProfileUser>>(PROFILE_URL, data)
      console.log('Профиль обновлен успешно')
    } catch (error) {
      console.error('Ошибка при обновлении профиля:', error)
      throw new Error('Не удалось обновить профиль')
    }
  }

  // Метод для получения профиля пользователя
  async getProfile(): Promise<ProfileUser> {
    try {
      const profile = await this.api.get<ProfileUser>(PROFILE_URL)
      if (!profile.user_id) {
        console.warn('Профиль получен без ID')
      }
      return profile
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

      const queryParams = {
        usersnumber: params.usersNumber,
        labels: params.labels ? params.labels.join(',') : undefined,
      }

      console.log('URL запроса:', '/api/v1/usersnumber')
      console.log('Подготовленные параметры запроса:', queryParams)

      const response = await this.api.get<{ result: ProfileUser[] }>('/api/v1/usersnumber', {
        params: queryParams,
      })

      // Добавляем подробное логирование
      console.log('Сырой ответ от сервера:', response)

      if (!response || !response.result) {
        console.error('Некорректный формат ответа:', response)
        throw new Error('Некорректный формат ответа от сервера')
      }

      return response
    } catch (error) {
      console.error('Error fetching users:', error)
      throw error
    }
  }

  async sendFriendRequest(targetUserId: number): Promise<HandshakeResponse> {
    try {
      const response = await this.api.post<HandshakeResponse>(HANDSHAKE_URL, { targetUserId })
      return response
    } catch (error) {
      console.error('Ошибка при отправке запроса в друзья:', error)
      throw error
    }
  }

  async getMatchRequests(): Promise<MatchResponse> {
    try {
      const { matches } = await this.api.get<MatchResponse>('/api/v1/matches')
      return { matches }
    } catch (error) {
      console.error('Ошибка при получении запросов в друзья:', error)
      throw error
    }
  }

  async acceptMatch(userId: number): Promise<void> {
    try {
      await this.api.post(`${MATCHES_URL}/accept`, { targetUserId: userId })
      console.log('Запрос в друзья принят')
    } catch (error) {
      console.error('Ошибка при принятии запроса в друзья:', error)
      throw error
    }
  }

  async declineMatch(userId: number): Promise<void> {
    try {
      await this.api.post(`${MATCHES_URL}/decline`, { targetUserId: userId })
      console.log('Запрос в друзья отклонен')
    } catch (error) {
      console.error('Ошибка при отклонении запроса в друзья:', error)
      throw error
    }
  }

  async getChats(): Promise<ChatsResponse> {
    try {
      const response = await this.api.get<ChatsResponse>(CHATS_URL)
      return response
    } catch (error) {
      console.error('Ошибка при получении чатов:', error)
      throw error
    }
  }

  // async getUserById(userId: number): Promise<ProfileUser> {
  //   try {
  //     const response = await this.api.get<ProfileUser>(`/api/v1/users/${userId}`);
  //     return response;
  //   } catch (error) {
  //     console.error('Ошибка при получении пользователя:', error);
  //     throw error;
  //   }
  // }

  async uploadAvatar(file: File): Promise<{ path: string }> {
    try {
      const formData = new FormData()
      formData.append('file', file)

      const response = await this.api.post<{ path: string }>('/api/v1/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      })
      return response
    } catch (error) {
      console.error('Ошибка при загрузке аватара:', error)
      throw error
    }
  }

  async getNotifications(): Promise<NotificationsResponse> {
    try {
      const response = await this.api.get<NotificationsResponse>(NOTIFICATIONS_URL)
      return response
    } catch (error) {
      console.error('Ошибка при получении уведомлений:', error)
      throw error
    }
  }

  // async markNotificationAsRead(notificationId: number): Promise<void> {
  //   try {
  //     await this.api.put(`/api/v1/notifications/${notificationId}/read`)
  //   } catch (error) {
  //     console.error('Ошибка при отметке уведомления как прочитанного:', error)
  //     throw error
  //   }
  // }
}

// Экспортируем экземпляр класса AuthService
export default new AuthService()
