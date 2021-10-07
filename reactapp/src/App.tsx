import React, {FC, ReactElement} from 'react';
import './App.css';
import {HashRouter, Link, Route, Switch} from 'react-router-dom';
import Login from './components/login';

function Home (props: any) {
  console.log('home', props)
  return (<div>Home</div>)
}

function Hello () {
  return (<div>Hello</div>)
}

function World (props: any) {
  console.log('world', props)
  return (<div>World</div>)
}

const App: FC = (): ReactElement => (
  <div className="App">
    <div>
      <Link to="/">Home</Link>
      <Link to="/login">Login</Link>
      <Link to="/hello">Hello</Link>
      <Link to="/world">World</Link>
      <Switch>
        <Route path="/" key='home' exact component={Home}/>
        <Route path="/login" key='login' exact component={Login}/>
        <Route path="/hello" key='hello' component={Hello}/>
        <Route path="/world" key='world' component={World}/>
      </Switch>
    </div>
  </div>
)

export default App;
