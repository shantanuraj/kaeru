import { createSelector } from '@reduxjs/toolkit';
import { RootState } from '../configureStore';

const getAuth = (state: RootState) => state.auth;

export const isSubmitting = createSelector(
  getAuth,
  (auth) => auth.loginState === 'inflight' || auth.signupState === 'inflight',
);

export const isLoggedIn = (state: RootState) => !!getAuth(state).token;
