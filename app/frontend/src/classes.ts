export interface User {
  username: string
  email: string
  password: string
  gender: string
  biography: string
  label: string[]
}

export interface LoginPayload {
  email: string
  password: string
}

export interface AuthedUser {
  email: string
  token: string
}

export interface ProfileUser {
  user_id: number
  username: string
  gender: string
  biography: string
  label: string[]
  online_status: string
  avatar?: string
}

export interface Chat {
  chat_id: number
  first_id: number
  second_id: number
  match_date: string
}

export interface ChatUser {
  user_id: number
  username: string
  gender: string
  biography: string | null
  label: string[] | null
  online_status: string
}

export interface ChatsResponse {
  chats: ChatUser[]
}

export interface MatchUser {
  user_id: number
  username: string
  gender: 'male' | 'female'
  biography: string | null
  label: string[] | null
  online_status: string
}

export interface MatchResponse {
  matches: MatchUser[]
}

export interface Sender {
  user_id: number;
  username: string;
  email: string;
  gender: string;
  biography: string | null;
  online_status: string;
  image: string;
}

export interface Notification {
  notification_id: number;
  user_id: number;
  sender_id: number;
  message: string;
  is_read: boolean;
  notification_date: string;
  sender: Sender;
}

export interface NotificationsResponse {
  notifications: Notification[];
}
