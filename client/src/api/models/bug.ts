export interface Bug {
  id: string
  title: string
  description: string
  status: string
  priority: string
  projectId: string
  asignee: string
  createdAt: Date
  updatedAt: Date
}

export interface CreateBugRequest {
  title: string
  description: string
  projectId: string
}
