import { useSelector } from 'react-redux';
import { isLoggedIn } from './selectors';

export function useLoggedIn() {
  return useSelector(isLoggedIn);
}
