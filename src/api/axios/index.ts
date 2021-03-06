import axios from 'axios'

const instance = axios.create({
  baseURL: (import.meta.env.APP_API_URL as string) || 'http://localhost:8080/blog',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
})

// eslint-disable-next-line unicorn/prevent-abbreviations
export const req = instance
