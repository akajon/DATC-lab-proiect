import React, { useState } from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from './components/Dashboard/Dashboard';
import Login from './components/Login/Login';

function App() {
  const [logged, setLogged] = useState(false);
  const [userType, setUserType] = React.useState(null);

  if(!logged) {
    return <Login logged={logged} setLogged={setLogged} setUserType={setUserType}/>
  }

  return (
    <div className="wrapper">
      <BrowserRouter>
        <Routes>
          <Route path="/dashboard" element={<Dashboard/>} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
