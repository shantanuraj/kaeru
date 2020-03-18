import axios from 'axios'

import { attachToken } from './headers'

const API_HOST = process.env.API_HOST

export const auth = axios.create({
  baseURL: `${API_HOST}/v1`,
})

auth.interceptors.request.use(attachToken)
