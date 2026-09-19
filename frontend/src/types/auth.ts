export interface User {
  id: number;
  username: string;
  role: string;
  tenant: string;
}

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}