interface StorageKeys {
  token: string;
  email: string;
}

class StorageService {
  private keys: StorageKeys = {
    token: "token",
    email: "email",
  };

  // Сохранение токена
  setToken(token: string): void {
    try {
      console.log('StorageService: сохраняем токен:', token);
      localStorage.setItem(this.keys.token, token);
      console.log('StorageService: проверка сохраненного токена:', localStorage.getItem(this.keys.token));
    }
    catch (error) {
      console.error("Error setting token:", error);
    }
  }

  // Получение токена
  getToken(): string | null {
    try {
      return localStorage.getItem(this.keys.token);
    }
    catch (error) {
      console.error("Error getting token:", error);
      return null;
    }
  }

  // Удаление токена
  removeToken(): void {
    try {
      localStorage.removeItem(this.keys.token);
    }
    catch (error) {
      console.error("Error removing token:", error);
    }
  }

  // Сохранение email пользователя
  setEmail(email: string): void {
    try {
      localStorage.setItem(this.keys.email, email);
    }
    catch (error) {
      console.error("Error setting email:", error);
    }
  }

  // Получение email пользователя
  getEmail(): string | null {
    try {
      return localStorage.getItem(this.keys.email);
    }
    catch (error) {
      console.error("Error getting email:", error);
      return null;
    }
  }

  // Удаление email пользователя
  removeEmail(): void {
    try {
      localStorage.removeItem(this.keys.email);
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
