import { Project } from '@/api/models/project'
import { getProjectsByUsername } from '@/api/projectService'
import {
  createContext,
  useCallback,
  FC,
  type ReactNode,
  useContext,
} from 'react'
import { useState } from 'react'
import { useUserContext } from './UserContext'

export type ProjectContextType = {
  updateUserProjects: () => void
  selectedProject: Project | null
  getProjects: () => Project[]
  // eslint-disable-next-line no-unused-vars
  setCurrentProject: (project: Project) => void
}

interface ProjectContextValue {
  // eslint-disable-next-line no-unused-vars
  setCurrentProject: (project: Project) => void
  updateUserProjects: () => Promise<void>
  selectedProject: Project | null
  getProjects: () => Project[]
}

export const ProjectContext = createContext<ProjectContextType | null>(null)
const ProjectProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [selectedProject, setSelectedProject] = useState<Project | null>(null)
  const [projects, setProjects] = useState<Project[]>([])
  const { user, token, logoutUser } = useUserContext()

  const updateUserProjects = useCallback(async () => {
    if (!user || !user.id || !token) return
    const response = await getProjectsByUsername(user.username, token)
    if (response.code == 401) {
      logoutUser()
      return
    }
    if (response.error) {
      return
    }
    setProjects(response)
  }, [token, user, logoutUser])

  const setCurrentProject = useCallback((project: Project) => {
    console.log(project)
    setSelectedProject(project)
    localStorage.setItem('project', JSON.stringify(project))
  }, [])

  const getProjects = () => {
    return projects
  }

  const value: ProjectContextValue = {
    setCurrentProject,
    updateUserProjects,
    selectedProject,
    getProjects,
  }

  return (
    <ProjectContext.Provider value={value}>{children}</ProjectContext.Provider>
  )
}

export const useProjectContext = () => {
  return useContext(ProjectContext) as ProjectContextType
}

export default ProjectProvider
