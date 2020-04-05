import { Auth } from './types';
import { createSelector } from '@reduxjs/toolkit';

export const isSubmitting = createSelector(
  (state: Auth) => state.loginState,
  (state: Auth) => state.signupState,
  (loginState, signupState) => (
    loginState === 'inflight' || signupState === 'inflight'
  )
)

export const isLoggedIn = (state: Auth) => !!state.token
