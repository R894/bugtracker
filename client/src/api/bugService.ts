import { Bug } from "./models/bug"
import { get, post } from "./utils"

export const getBugs = async (): Promise<Bug> => {
    return get('bugs')
}

export const createBug = async (bugData: object) => {
    return post('bugs', bugData)
}

export const getBugById = async (bugId: string) => {
    return get(`bugs/${bugId}`)
}

export const getBugsByProjectId = async (projectId: string) => {
    return get(`bugs/projects/${projectId}`)
}