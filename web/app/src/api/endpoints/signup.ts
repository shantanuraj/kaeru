import { auth } from '../index'
import { Kaeru } from '../../shared/types';
import { Response } from './types';
import { setToken } from '../headers';

export async function signup (user: Kaeru.User) {
  try {
    const response = await auth.post<Response.Token>('/signup', user)
    const { token } = response.data
    setToken(token)
    return token
  } catch (err) {
    console.error({ context: 'auth.signup.error' }, err)
    throw err
  }
}
