import { Bug, CreateBugRequest } from "./models/bug"
import { get, post } from "./utils"

export const getBugs = async (token: string): Promise<Bug[]> => {
    return get('bugs', token) as Promise<Bug[]>
}

export const createBug = async (bugData: CreateBugRequest, token: string) => {
    return post('bugs', bugData, token)
}

export const getBugById = async (bugId: string, token: string) => {
    return get(`bugs/${bugId}`, token) as Promise<Bug>
}

export const getBugsByProjectId = async (projectId: string, token: string) => {
    return get(`bugs/projects/${projectId}`, token) as Promise<Bug[]>
}