import React, { useState } from 'react';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Dashboard from './components/Dashboard/Dashboard';
import Login from './components/Login/Login';

import { GoogleMap, useJsApiLoader, useLoadScript, Marker } from "@react-google-maps/api"

function App() {
  const [logged, setLogged] = useState(false);

  if(!logged) {
    return <Login logged={logged} setLogged={setLogged}/>
  }

  return (
    <div className="wrapper">
      <h1>City Dangers Alert</h1>
      <BrowserRouter>
        <Routes>
          <Route path="/dashboard" element={<Dashboard/>} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
