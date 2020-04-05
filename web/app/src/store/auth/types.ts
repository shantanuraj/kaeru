export interface Auth {
  token: string
  error: string | null,
  signupState: 'idle' | 'inflight'
  loginState: 'idle' | 'inflight'
}
