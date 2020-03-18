import { PayloadAction, createSlice } from '@reduxjs/toolkit'
import { Auth } from './types'

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
