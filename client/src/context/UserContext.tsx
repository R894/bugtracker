import { LoginRequest, RegisterRequest, User } from '@/api/models/user'
import { createUser, postLogin } from '@/api/userService'
import React, { useEffect, useState } from 'react'

export type UserContextType = {
  token: string
  user: User | null
  isLoggedIn: boolean
  loginUser: typeof postLogin
  registerUser: typeof createUser
  logoutUser: () => void
}

export const UserContext = React.createContext<UserContextType | null>(null)

const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [token, setToken] = useState('')
  const [user, setUser] = useState<User | null>(null)
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const t = localStorage.getItem('token')
    if (t && t != '') {
      setToken(t)
      setIsLoggedIn(true)
    }
  }, [])

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
    localStorage.removeItem('token')
  }

  return (
    <UserContext.Provider
      value={{ loginUser, user, registerUser, isLoggedIn, logoutUser, token }}
    >
      {children}
    </UserContext.Provider>
  )
}

export default UserProvider
