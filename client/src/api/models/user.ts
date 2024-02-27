import { Project } from "./project"

export interface User {
  id: string
  username: string
  email: string
  projects: Project[]
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
