import axios from 'axios';

import { attachToken } from './headers';

const AUTH_HOST = process.env.REACT_APP_AUTH_HOST;

export const auth = axios.create({
  baseURL: `${AUTH_HOST}/v1`,
});

auth.interceptors.request.use(attachToken);
