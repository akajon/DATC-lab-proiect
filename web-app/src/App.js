import React, { useState } from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from './components/Dashboard/Dashboard';
import Login from './components/Login/Login';
import Register from './components/Register/Register';

function App() {
  const [logged, setLogged] = useState(false);
  const [user, setUser] = React.useState(null);

  if(!logged) {
    return <Login setLogged={setLogged} setUser={setUser}/>
  }

  return (
    <Dashboard user={user}/>
    // <div className="wrapper">
    //   <BrowserRouter>
    //     <Routes>
    //       <Route path="/dashboard" element={<Dashboard/>} />
    //       <Route path="/register" element={<Register/>} />
    //     </Routes>
    //   </BrowserRouter>
    // </div>
  );
}

export default App;
