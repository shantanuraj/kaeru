import { auth } from '../index';

export async function validate(): Promise<boolean> {
  try {
    const response = await auth.get('/validate');
    return response.status === 200;
  } catch (err) {
    console.error({ context: 'auth.signup.error' }, err);
    throw err;
  }
}
