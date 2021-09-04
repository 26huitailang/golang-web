import React from 'react';
import Login from './components/Login';
import TodoList from './components/TodoList';

function App() {
  const authed = localStorage.getItem('authed');

  return (
    <div className="App">
      {authed ? <TodoList /> : <Login />}
    </div>
  );
}

export default App;
