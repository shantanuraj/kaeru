import { PayloadAction, createSlice } from '@reduxjs/toolkit';
import { Auth } from './types';
import { Kaeru } from '../../shared/types';
import * as api from '../../api/endpoints';
import { AppThunk } from '../configureStore';

const auth = createSlice({
  name: 'auth',
  initialState: {
    token: '',
    error: null,
    loginState: 'idle',
    signupState: 'idle',
  } as Auth,
  reducers: {
    startLogin: (state) => {
      state.loginState = 'inflight';
    },
    finishLogin: (state, action: PayloadAction<null | string>) => {
      state.loginState = 'idle';
      state.error = action.payload;
    },

    startSignup: (state) => {
      state.signupState = 'inflight';
    },
    finishSignup: (state, action: PayloadAction<null | string>) => {
      state.signupState = 'idle';
      state.error = action.payload;
    },

    setToken: (state, action: PayloadAction<string>) => {
      state.token = action.payload;
    },
  },
});

export const {
  setToken,
  startLogin,
  startSignup,
  finishLogin,
  finishSignup,
} = auth.actions;

export default auth.reducer;

export const login = (user: Kaeru.User): AppThunk => async (dispatch) => {
  dispatch(startLogin());
  try {
    const token = await api.login(user);
    dispatch(setToken(token));
    dispatch(finishLogin(null));
  } catch (err) {
    dispatch(finishLogin((err as Error).message));
  }
};

export const signup = (user: Kaeru.User): AppThunk => async (dispatch) => {
  dispatch(startSignup());
  try {
    const token = await api.signup(user);
    dispatch(setToken(token));
    dispatch(finishLogin(null));
  } catch (err) {
    dispatch(finishSignup((err as Error).message));
  }
};
