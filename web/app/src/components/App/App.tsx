import React from 'react';
import './App.css';
import { Switch, Route, Redirect } from 'react-router-dom';
import { useLoggedIn } from '../../store/auth/hooks';
import Header from '../Header';
import LandingPage from '../LandingPage';
import LoginForm from '../LoginForm';

function App() {
  const isLoggedIn = useLoggedIn();
  return (
    <div className="App">
      <Header />
      <main id="App-content">
        <Switch>
          <Route path="/" exact component={LandingPage} />
          <Route path="/app" exact>
            {!isLoggedIn && <Redirect to="/login" />}
          </Route>
          <Route path="/:mode(login|signup)" exact component={LoginForm} />
        </Switch>
      </main>
    </div>
  );
}

export default App;
