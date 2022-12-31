import React, { useState } from 'react';
import './Login.css';
import PropTypes from 'prop-types';

async function loginUser(credentials) {
  return fetch('http://localhost:8080/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(credentials)
  })
    .then(data => data.json())
 }

export default function Login({ setToken, setLogged }) {
  const [username, setUserName] = useState();
  const [password, setPassword] = useState();

  const handleSubmit = async e => {
    setLogged(true);
    return;
    e.preventDefault();
    const token = await loginUser({
      username,
      password
    });
    setToken(token);
  }

  const handleUserCreate = async e => {
    console.log("clicked user create");
    console.log(username);
  }

  const handleUserDelete = async e => {
    console.log("clicked user delete");
    console.log(password);
  }

  return(
    <div className="login-wrapper">
      <h1>Please Log In</h1>
      <form onSubmit={handleSubmit}>
        <label>
          <p>Username</p>
          <input type="text" onChange={e => setUserName(e.target.value)}/>
        </label>
        <label>
          <p>Password</p>
          <input type="password" onChange={e => setPassword(e.target.value)}/>
        </label>
        <div>
        <br/>
          <button type="submit">Login</button>
        </div>
      </form>

      <h2>Other Actions</h2>
      <table>
        <tbody>
          <tr>
            <td>
              <form onSubmit={handleUserCreate}>
                <div>
                  <button type="button" onClick={handleUserCreate}>Create User</button>
                </div>
              </form>
            </td>
            <td>
              <form onSubmit={handleUserDelete}>
                <div>
                  <button type="button" onClick={handleUserDelete}>Delete User</button>
                </div>
              </form>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  )
}

Login.propTypes = {
  setToken: PropTypes.func.isRequired
}
