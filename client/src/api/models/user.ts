export interface User {
  id: string
  username: string
  email: string
  projects: string[]
  createdAt: Date
  updatedAt: Date
}

export interface RegisterRequest {
  username: string
  email: string
  password: string
}

export interface LoginRequest {
  email: string
  password: string
}
