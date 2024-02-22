import { RequestInit } from 'next/dist/server/web/spec-extension/request'
const API_URL = process.env.API_URL

export async function get(endpoint: string): Promise<any> {
  try {
    const response = await fetch(`${API_URL}/${endpoint}`)
    if (!response.ok) {
      return {error: true, ...await response.json()}
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return {error: true, message: error}
  }
}

export async function post(endpoint: string, body: object): Promise<any> {
  try {
    const options: RequestInit = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    }
    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return {error: true, ...await response.json()}
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return {error: true, message: error}
  }
}

export async function put(endpoint: string, body: object): Promise<any> {
  try {
    const options: RequestInit = {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    }

    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return {error: true, ...await response.json()}
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return {error: true, message: error}
  }
}

export async function del(endpoint: string, body: object): Promise<any> {
  try {
    const options: RequestInit = {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(body),
    }

    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return {error: true, ...await response.json()}
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return {error: true, message: error}
  }
}
