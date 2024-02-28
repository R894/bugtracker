export interface Project {
  id: string
  name: string
  description: string
  ownerId: string
  createdAt: Date
  updatedAt: Date
}

export interface CreateProjectRequest {
  name: string
  description: string
}
