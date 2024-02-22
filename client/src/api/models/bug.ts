export interface Bug {
    id: string,
    title: string,
    description: string,
    status: string,
    priority: string,
    projectId: string,
    createdAt: Date,
    updatedAt: Date
}