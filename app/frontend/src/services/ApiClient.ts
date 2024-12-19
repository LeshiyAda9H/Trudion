import axios, { type AxiosInstance, type AxiosRequestConfig } from 'axios'

export class ApiClient {
  private static instance: ApiClient
  private api: AxiosInstance

  private constructor() {
    this.api = axios.create({
      baseURL: 'http://localhost:8080',
      headers: {
        'Content-Type': 'application/json',
        // 'Accept': 'application/json',
        // 'Access-Control-Allow-Origin': '*',
        // 'Access-Control-Allow-Methods': 'GET, POST, PATCH, PUT, DELETE, OPTIONS',
        // 'Access-Control-Allow-Headers': 'Origin, Content-Type, X-Auth-Token'
      },
    })

    // Интерцептор для запросов
    this.api.interceptors.request.use(
      (config) => {
        const token = localStorage.getItem('token')
        console.log('Токен в интерцепторе:', token);
        if (token) {
          config.headers.Authorization = `Bearer ${token}`
        }
        return config
      },
      (error) => {
        return Promise.reject(error)
      },
    )

    // Интерцептор для ответов
    this.api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          // Если получаем 401, значит токен невалидный или истек
          localStorage.removeItem('token')
          // Можно добавить редирект на страницу логина
        }
        return Promise.reject(error)
      },
    )
  }

  // Метод для получения экземпляра ApiClient
  public static getInstance(): ApiClient {
    if (!ApiClient.instance) {
      ApiClient.instance = new ApiClient()
    }
    return ApiClient.instance
  }

  // Метод для отправки GET-запроса
  public async get<T>(url: string, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.api.get<T>(url, config)
    return response.data
  }

  // Метод для отправки POST-запроса
  public async post<T, D = unknown>(url: string, data?: D, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.api.post<T>(url, data, config)
    return response.data
  }

  public async patch<T, D = unknown>(url: string, data?: D, config?: AxiosRequestConfig): Promise<T> {
    const response = await this.api.patch<T>(url, data, config)
    return response.data
  }
  // Другие методы (put, delete и т.д.)
}

export default ApiClient.getInstance()
