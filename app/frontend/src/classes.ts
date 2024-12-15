export interface User {
  username: string;
  email: string;
  password: string;
  gender: string;
  biography: string;
  label: string[];
};

export interface LoginPayload {
  email: string;
  password: string;
}

export interface AuthedUser {
  access_token: string;
  email: string;
};

export interface ProfileUser {
  username: string;
  gender: string;
  biography: string;
  label: string[];
  online_status : string;
};


