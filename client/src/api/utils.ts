import { RequestInit } from 'next/dist/server/web/spec-extension/request'
const API_URL = process.env.API_URL

export async function get(endpoint: string, token?: string): Promise<any> {
  try {
    const options: RequestInit = {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    }
    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return { error: true, ...(await response.json()) }
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return { error: true, message: error }
  }
}

export async function post(
  endpoint: string,
  body: object,
  token?: string,
): Promise<any> {
  try {
    const options: RequestInit = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(body),
    }
    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return { error: true, ...(await response.json()) }
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return { error: true, message: error }
  }
}

export async function put(
  endpoint: string,
  body: object,
  token?: string,
): Promise<any> {
  try {
    const options: RequestInit = {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(body),
    }

    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return { error: true, ...(await response.json()) }
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return { error: true, message: error }
  }
}

export async function del(
  endpoint: string,
  body: object,
  token?: string,
): Promise<any> {
  try {
    const options: RequestInit = {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(body),
    }

    const response = await fetch(`${API_URL}/${endpoint}`, options)
    if (!response.ok) {
      return { error: true, ...(await response.json()) }
    }
    return await response.json()
  } catch (error) {
    console.error(error)
    return { error: true, message: error }
  }
}
