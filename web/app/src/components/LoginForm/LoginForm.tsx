import React from 'react';
import { useFormik } from 'formik';

import { Kaeru } from '../../shared/types';
import { useParams } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { login, signup } from '../../store/auth';

import './LoginForm.css';

interface Params {
  mode: 'login' | 'signup'
}

export default function LoginForm () {
  const { mode } = useParams<Params>()
  const dispatch = useDispatch()
  const formik = useFormik({
    initialValues: defaultValue,
    onSubmit: (user) => {
      const cb = mode === 'signup' ? signup : login
      dispatch(cb(user))
    }
  })
  return (
    <div>
      <form onSubmit={formik.handleSubmit}>
        {mode === 'signup' && (
          <React.Fragment>
            <label htmlFor="name">Name</label>
            <input
              type="name"
              name="name"
              onChange={formik.handleChange}
              value={formik.values.name}
            />
          </React.Fragment>
        )}
        <label htmlFor="email">Email</label>
        <input
          type="email"
          name="email"
          onChange={formik.handleChange}
          value={formik.values.email}
        />
        <label htmlFor="password">Password</label>
        <input
          type="password"
          name="password"
          onChange={formik.handleChange}
          value={formik.values.password}
        />
      </form>
    </div>
  );
}

const defaultValue: Kaeru.User = {
  name: '',
  email: '',
  password: ''
}
