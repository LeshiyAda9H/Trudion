export interface User {
  username: string;
  email: string;
  password: string;
  gender: string;
};

export interface LoginPayload {
  email: string;
  password: string;
}

export interface AuthedUser {
  access_token: string;
  email: string;
};

