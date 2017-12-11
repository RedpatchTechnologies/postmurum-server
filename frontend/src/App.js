import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Organization from './components/orgs';
import Auth from './components/auth';

export class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
        <div>
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Welcome to React</h1>
        </div>
        <div align="right">
            {<Auth />}
        </div>
        </header>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.asdf
        
        </p>
        see our orgs {<Organization />}
        
      </div>
    );
  }
}

export default App;
