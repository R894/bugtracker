import { CreateProjectRequest } from './models/project'
import { get, post } from './utils'

export const createProject = async (body: CreateProjectRequest, token: string) => {
  await post('projects', body, token)
}

export const getProjectsByUsername = async (userId: string, token: string) => {
  const projects = await get(`projects/${userId}`, token)
  return projects
}
