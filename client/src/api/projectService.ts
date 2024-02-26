import { get, post } from './utils'

export const createProject = async (body: object) => {
  await post('projects', body)
}

export const getProjectsByUsername = async (userId: string, token: string) => {
  const projects = await get(`projects/${userId}`, token)
  return projects
}
