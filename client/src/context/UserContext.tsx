import { LoginRequest } from '@/api/models/user'
import { postLogin } from '@/api/userService'
import React, { useState } from 'react'

export type UserContextType = {
  token: string
  loginUser: typeof postLogin
  logoutUser: () => void
}

export const UserContext = React.createContext<UserContextType | null>(null)

const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [token, setToken] = useState('')

  const loginUser = async (loginRequest: LoginRequest) => {
    const response = await postLogin(loginRequest)
    if (!response.error) {
      const currentToken = response.token
      console.log('setting token to', currentToken)
      setToken(currentToken)
      localStorage.setItem('token', currentToken)
    }
  }

  const logoutUser = () => {
    setToken('')
    localStorage.removeItem('token')
  }

  return (
    <UserContext.Provider value={{ loginUser, logoutUser, token }}>
      {children}
    </UserContext.Provider>
  )
}

export default UserProvider
