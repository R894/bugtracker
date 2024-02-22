import { post } from "./utils"

export const createProject = async (body: object) => {
  await post('projects', body)
}
