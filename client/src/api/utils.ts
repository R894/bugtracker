import { RequestInit } from 'next/dist/server/web/spec-extension/request'
const API_URL = process.env.API_URL

export async function get(endpoint: string): Promise<any> {
  fetch(`${API_URL}/${endpoint}`)
    .then((response: Response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }
      return response.json()
    })
    .catch((error: Error) => {
      console.error('There was a problem with the fetch operation:', error)
    })
}

export async function post(endpoint: string, body: object): Promise<any> {
  const options: RequestInit = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  }

  fetch(`${API_URL}/${endpoint}`, options)
    .then((response: Response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }
      return response.json()
    })
    .catch((error: Error) => {
      console.error('There was a problem with the fetch operation:', error)
    })
}

export async function put(endpoint: string, body: object): Promise<any> {
  const options: RequestInit = {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  }

  fetch(`${API_URL}/${endpoint}`, options)
    .then((response: Response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }
      return response.json()
    })
    .catch((error: Error) => {
      console.error('There was a problem with the fetch operation:', error)
    })
}

export async function del(endpoint: string, body: object): Promise<any> {
  const options: RequestInit = {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(body),
  }

  fetch(`${API_URL}/${endpoint}`, options)
    .then((response: Response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok')
      }
      return response.json()
    })
    .catch((error: Error) => {
      console.error('There was a problem with the fetch operation:', error)
    })
}
