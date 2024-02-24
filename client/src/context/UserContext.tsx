import { LoginRequest, RegisterRequest } from '@/api/models/user'
import { createUser, postLogin } from '@/api/userService'
import React, { useEffect, useState } from 'react'

export type UserContextType = {
  token: string
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
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  useEffect(() => {
    const t = localStorage.getItem('token')
    if (t && t != '') {
      setToken(t)
      setIsLoggedIn(true)
    }
  }, [])

  const loginUser = async (loginRequest: LoginRequest) => {
    setIsLoggedIn(false)
    const response = await postLogin(loginRequest)
    if (response.error) {
      return response
    }
    setIsLoggedIn(true)
    const currentToken = response.token
    setToken(currentToken)
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
      value={{ loginUser, registerUser, isLoggedIn, logoutUser, token }}
    >
      {children}
    </UserContext.Provider>
  )
}

export default UserProvider
