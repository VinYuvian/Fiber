import React from 'react';
import HeaderMenu from './components/HeaderMenu';
import LoginForm from './components/Form';
import Login from './components/Login';
import Home from './components/Home';
import User from './components/User'
import { BrowserRouter as Router, Route } from 'react-router-dom';

//import './App.css';

function App() {
  return (
    <Router>
      <div>
      <HeaderMenu />
        <Route exact path="/" component={Home}/>
        <Route path="/signup" component={LoginForm}/>
        <Route path="/login" component={Login}/>
        <Route path="/user" component={User}/>
       </div>
    </Router>
  );
}

export default App;
