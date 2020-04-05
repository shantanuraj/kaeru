import { auth } from '../index';
import { Kaeru } from '../../shared/types';
import { Response } from './types';
import { setToken } from '../headers';

export async function login(user: Pick<Kaeru.User, 'email' | 'password'>) {
  try {
    const response = await auth.post<Response.Token>('/login', user);
    const { token } = response.data;
    setToken(token);
    return token;
  } catch (err) {
    console.error({ context: 'auth.login.error' }, err);
    throw err;
  }
}
