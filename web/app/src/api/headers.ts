import { AxiosRequestConfig } from 'axios'

let token: null | string = null

export function attachToken(config: AxiosRequestConfig) {
  if (token) {
    config.headers = {
      'Authorization': `Bearer:${token}`
    }
  }
  return config
}

export function setToken(newToken: string) {
  token = newToken
}
