import { Project } from '@/api/models/project'
import { LoginRequest, RegisterRequest, User } from '@/api/models/user'
import { getProjectsByUsername } from '@/api/projectService'
import { createUser, postLogin } from '@/api/userService'
import React, { useCallback, useContext, useEffect, useState } from 'react'

export type UserContextType = {
  token: string
  user: User | null
  isLoggedIn: boolean
  loginUser: typeof postLogin
  registerUser: typeof createUser
  logoutUser: () => void
  updateUserProjects: () => {}
  selectedProject: Project | null
  // eslint-disable-next-line no-unused-vars
  setCurrentProject: (project: Project) => void
}

export const UserContext = React.createContext<UserContextType | null>(null)

const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [token, setToken] = useState('')
  const [user, setUser] = useState<User | null>(null)
  const [selectedProject, setSelectedProject] = useState<Project | null>(null)
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')
    const storedProject = localStorage.getItem('project')

    setToken(storedToken || '')
    setUser(storedUser ? JSON.parse(storedUser) : null)
    setSelectedProject(storedProject ? JSON.parse(storedProject) : null)
    setIsLoggedIn(!!storedToken)
  }, [])

  const updateUserProjects = useCallback(async () => {
    if (!user || !user.id || !token) return
    const response = await getProjectsByUsername(user.username, token)
    console.log(response)
    if (response.code == 401) {
      logoutUser()
      return
    }
    if (response.error) {
      return
    }
    let u = JSON.parse(localStorage.getItem('user') || '')
    u.projects = response
    localStorage.setItem('user', JSON.stringify(u))
    setUser((prevUser) => ({ ...prevUser, projects: response }) as User)
  }, [token, user])

  useEffect(() => {
    console.log(user?.projects)
  }, [user])

  const setCurrentProject = useCallback((project: Project) => {
    console.log(project)
    setSelectedProject(project)
    localStorage.setItem('project', JSON.stringify(project))
  },[])

  const loginUser = async (loginRequest: LoginRequest) => {
    setIsLoggedIn(false)
    const response = await postLogin(loginRequest)
    if (response.error) {
      return response
    }
    setIsLoggedIn(true)
    const currentToken = response.token
    const currentUser = response.user
    setToken(currentToken)
    setUser(currentUser)
    localStorage.setItem('user', JSON.stringify(currentUser))
    localStorage.setItem('token', currentToken)
    return response
  }

  const registerUser = async (registerRequest: RegisterRequest) => {
    const response = await createUser(registerRequest)
    if (response.error) {
      return response
    }

    return response
  }

  const logoutUser = () => {
    setIsLoggedIn(false)
    setToken('')
    setUser(null)
    setSelectedProject(null)
    localStorage.removeItem('token')
  }

  return (
    <UserContext.Provider
      value={{
        loginUser,
        user,
        registerUser,
        updateUserProjects,
        isLoggedIn,
        logoutUser,
        token,
        selectedProject,
        setCurrentProject,
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

export default UserProvider

export const useUserContext = () => {
  return useContext(UserContext) as UserContextType
}
