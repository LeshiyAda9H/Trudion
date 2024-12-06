export interface User {
  username: string;
  password: string;
};
export interface AuthedUser {
  access_token: string;
  username: string;
};

export interface FullLink {
  url: string;
}

export interface ShortLink {
  full_url: string;
  short_link: string;
}
