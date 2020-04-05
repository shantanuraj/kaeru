import {
  combineReducers,
  configureStore,
  Action,
  ThunkAction,
} from '@reduxjs/toolkit';

import auth from './auth';

const rootReducer = combineReducers({
  auth,
});

export type RootState = ReturnType<typeof rootReducer>;

export type AppThunk = ThunkAction<void, RootState, unknown, Action<string>>;

export function getStore() {
  const store = configureStore({
    reducer: rootReducer,
  });
  return store;
}
