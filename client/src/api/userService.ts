import { get, post } from './utils'

export const getUsers = async () => {
  return get('/users')
}

export const createUser = async (registerData: object) => {
  return post('/users', registerData)
}

export const postLogin = async (loginData: object) => {
  return post('/users/login', loginData)
}

export const getUserById = async (id: string) => {
  return get(`/users/${id}`)
}
