import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
//import App from './App';
//import reportWebVitals from './reportWebVitals';
import Welcome from './components/welcome'
import Tick from './components/tick'
import Tog from './components/tog'
import Ls from './components/ls'

interface User {
  firstName: string;
  lastName: string;
}
let user: User = {
  firstName: "Tao",
  lastName: "Wu"
}
function toString(u: User): string {
  return u.firstName + " " + u.lastName
}

function App() {
  return (
    <div>
      <h1>Hello, {toString(user)}!</h1>
      <Welcome name="welcome 1" />
      <Welcome name="welcome 2" />
      <Tick />
      <Tog />
      <Ls IsDouble={false}/>
      <Ls IsDouble={true}/>
    </div>
  );
}

ReactDOM.render(
  <App />,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
//reportWebVitals();
