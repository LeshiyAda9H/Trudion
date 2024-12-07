class StorageService {
  private tokenKey = "token";
  private emailKey = "email";

  // Сохранение токена
  setToken(token: string): void {
    try {
      localStorage.setItem(this.tokenKey, token);
    }
    catch (error) {
      console.error("Error setting token:", error);
    }
  }

  // Получение токена
  getToken(): string | null {
    try {
      return localStorage.getItem(this.tokenKey);
    }
    catch (error) {
      console.error("Error getting token:", error);
      return null;
    }
  }

  // Удаление токена
  removeToken(): void {
    try {
      localStorage.removeItem(this.tokenKey);
    }
    catch (error) {
      console.error("Error removing token:", error);
    }
  }

  // Сохранение имени пользователя
  setEmail(email: string): void {
    try {
      localStorage.setItem(this.emailKey, email);
    }
    catch (error) {
      console.error("Error setting email:", error);
    }
  }

  // Получение имени пользователя
  getEmail(): string | null {
    try {
      return localStorage.getItem(this.emailKey);
    }
    catch (error) {
      console.error("Error getting email:", error);
      return null;
    }
  }

  // Удаление имени пользователя
  removeEmail(): void {
    try {
      localStorage.removeItem(this.emailKey);
    }
    catch (error) {
      console.error("Error removing email:", error);
    }
  }

  // Очистка всех данных
  clear(): void {
    this.removeToken();
    this.removeEmail();
  }
}

export default new StorageService();
