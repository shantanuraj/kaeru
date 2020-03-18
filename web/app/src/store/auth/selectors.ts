import { Auth } from './types';



export const isLoggedIn = (state: Auth) => !!state.token
