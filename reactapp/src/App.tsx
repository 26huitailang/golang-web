import React, {FC, ReactElement} from 'react';
import './App.css';
import {Link, Route, Router} from 'react-router-dom';
import MyRouter from './router';

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
    <MyRouter>
      <div>
        <Link to="/">Home</Link>
        <Link to="/hello">Hello</Link>
        <Link to="/world">World</Link>
      </div>
      <Route path="/" key='home' exact component={Home}/>
      <Route path="/hello" key='hello' component={Hello}/>
      <Route path="/world" key='world' component={World}/>
    </MyRouter>
  </div>
)

export default App;
