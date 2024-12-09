export interface User {
  nickname: string;
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

export interface FullLink {
  url: string;
}

export interface ShortLink {
  full_url: string;
  short_link: string;
}
