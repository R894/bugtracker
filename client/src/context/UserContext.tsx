import { LoginRequest, RegisterRequest, User } from '@/api/models/user'
import { createUser, postLogin } from '@/api/userService'
import React, {
  type FC,
  type ReactNode,
  createContext,
  useContext,
  useEffect,
  useState,
} from 'react'

export type UserContextType = {
  token: string
  user: User | null
  isLoggedIn: boolean
  loginUser: typeof postLogin
  registerUser: typeof createUser
  logoutUser: () => void
}

export const UserContext = createContext<UserContextType | null>(null)

const UserProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [token, setToken] = useState('')
  const [user, setUser] = useState<User | null>(null)
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  const initializeUser = () => {
    const storedToken = localStorage.getItem('token')
    const storedUser = localStorage.getItem('user')

    setToken(storedToken || '')
    setUser(storedUser ? JSON.parse(storedUser) : null)
    setIsLoggedIn(!!storedToken)
  }

  useEffect(() => {
    initializeUser()
  }, [])

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
        isLoggedIn,
        logoutUser,
        token,
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
