import React from 'react';
import { useFormik, FormikErrors } from 'formik';

import { Kaeru } from '../../shared/types';
import { useParams } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { login, signup } from '../../store/auth';
import { isSubmitting } from '../../store/auth/selectors';

import './LoginForm.css';

interface Params {
  mode: 'login' | 'signup'
}

export default function LoginForm () {
  const { mode } = useParams<Params>()
  const dispatch = useDispatch()
  const isSubmittingForm = useSelector(isSubmitting)
  const formik = useFormik({
    initialValues: defaultValue,
    validate: validateField(mode as Params['mode']),
    onSubmit: (user) => {
      const cb = mode === 'signup' ? signup : login
      dispatch(cb(user))
    }
  })
  const { errors, getFieldProps, handleSubmit, touched } = formik
  const disableSubmit = isSubmittingForm
  return (
    <div className="login-form">
      <form onSubmit={handleSubmit}>
        {mode === 'signup' && (
          <div className="field">
            <label htmlFor="name">Name</label>
            <input
              type="name"
              name="name"
              {...getFieldProps('name')}
            />
            {errors.name && touched.name && (<span className="error">{errors.name}</span>)}
          </div>
        )}

        <div className="field">
          <label htmlFor="email">Email</label>
          <input
            type="email"
            name="email"
            {...getFieldProps('email')}
          />
          {errors.email && touched.email && (<span className="error">{errors.email}</span>)}
        </div>

        <div className="field">
          <label htmlFor="password">Password</label>
          <input
            type="password"
            name="password"
            {...getFieldProps('password')}
          />
          {errors.password && touched.password && (<span className="error">{errors.password}</span>)}
        </div>

        <button
          type="submit"
          disabled={disableSubmit}
        >
          {mode === 'signup' ? 'Sign up' : 'Log in'}
        </button>
      </form>
    </div>
  );
}

const defaultValue: Kaeru.User = {
  name: '',
  email: '',
  password: ''
}

const validateField = (mode: Params['mode']) => (values: Kaeru.User) => {
  let errors: FormikErrors<Kaeru.User> = {}
  if (mode === 'signup' && !values.name) {
    errors.name = 'Name is required'
  }
  if (!values.email) {
    errors.email = 'Email is required'
  }
  if (!values.password) {
    errors.password = 'Password is required'
  }
  if (values.password.length < 8) {
    errors.password = 'Password should be atleast 8 characters'
  }
  return errors
}
