import { combineReducers, configureStore } from '@reduxjs/toolkit'

import auth from './auth'

const rootReducer = combineReducers({
  auth
})

export type RootState = ReturnType<typeof rootReducer>

export function getStore () {
  const store = configureStore({
    reducer: rootReducer
  })
  return store
}
