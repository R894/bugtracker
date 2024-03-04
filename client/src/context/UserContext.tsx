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
  setSelectedProject: React.Dispatch<React.SetStateAction<Project | null>>
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
    const t = localStorage.getItem('token')
    if (t && t != '') {
      setToken(t)
      setIsLoggedIn(true)
    }
  }, [])

  useEffect(() => {
    const u = localStorage.getItem('user')
    if (u && u != '') {
      console.log(JSON.parse(u))
      setUser(JSON.parse(u))
    }
  }, [])

  const updateUserProjects = useCallback(async () => {
    if (!user || !user.id || !token) return
    const response = await getProjectsByUsername(user.username, token)
    console.log(response)
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
        setSelectedProject
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
