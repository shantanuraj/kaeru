import React from 'react';
import { useLoggedIn } from '../../store/auth/hooks';
import { Link } from 'react-router-dom';

import './Header.css';

export default function Header() {
  const isLoggedIn = useLoggedIn();
  return (
    <header className="App-header">
      {!isLoggedIn && (
        <nav>
          <Link to="/login">Login</Link>
          <Link to="/signup">Create Account</Link>
        </nav>
      )}
    </header>
  );
}
