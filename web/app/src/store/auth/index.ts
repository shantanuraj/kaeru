import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { Auth } from './types'
import { Kaeru } from '../../shared/types'
import * as api from '../../api/endpoints'
import { AppThunk } from '../configureStore'

const auth = createSlice({
  name: 'auth',
  initialState: { token: '' } as Auth,
  reducers: {
    setToken: (state, action: PayloadAction<string>) => {
      state.token = action.payload
    }
  }
})

export const { setToken } = auth.actions

export default auth.reducer

export const login =
  (user: Kaeru.User): AppThunk => async dispatch => {
    const token = await api.login(user)
    dispatch(setToken(token))
  }

export const signup =
  (user: Kaeru.User): AppThunk => async dispatch => {
    const token = await api.signup(user)
    dispatch(setToken(token))
  }
