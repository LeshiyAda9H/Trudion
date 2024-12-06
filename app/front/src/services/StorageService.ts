class StorageService {
  private tokenKey = "token";
  private usernameKey = "username";

  // Сохранение токена
  setToken(token: string): void {
    localStorage.setItem(this.tokenKey, token);
  }

  // Получение токена
  getToken(): string | null {
    return localStorage.getItem(this.tokenKey);
  }

  // Удаление токена
  removeToken(): void {
    localStorage.removeItem(this.tokenKey);
  }

  // Сохранение имени пользователя
  setUsername(username: string): void {
    localStorage.setItem(this.usernameKey, username);
  }

  // Получение имени пользователя
  getUsername(): string | null {
    return localStorage.getItem(this.usernameKey);
  }

  // Удаление имени пользователя
  removeUsername(): void {
    localStorage.removeItem(this.usernameKey);
  }

  // Очистка всех данных
  clear(): void {
    this.removeToken();
    this.removeUsername();
  }
}

export default new StorageService();
