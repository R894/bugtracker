import { Bug } from "./models/bug"
import { get, post } from "./utils"

export const getBugs = async (token: string): Promise<Bug[]> => {
    return get('bugs', token)
}

export const createBug = async (bugData: object) => {
    return post('bugs', bugData)
}

export const getBugById = async (bugId: string) => {
    return get(`bugs/${bugId}`)
}

export const getBugsByProjectId = async (projectId: string, token: string) => {
    return get(`bugs/projects/${projectId}`, token)
}