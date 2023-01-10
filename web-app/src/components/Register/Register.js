import React from 'react';

export default function Register() {
    const handleUserCreate  = (event) =>  {
        console.log("clicked user create");
        console.log(event.target[0].value);
      }

    return <form onSubmit={handleUserCreate}>
    <div>
        <label>
            <p>First Name</p>
            <input type="text"/>
        </label>
        <label>
            <p>Last Name</p>
            <input type="text"/>
        </label>
        <label>
            <p>Email</p>
            <input type="text"/>
        </label>
        <label>
            <p>Password</p>
            <input type="password"/>
        </label>
        <button type="button" onClick={handleUserCreate}>Create User</button>
    </div>
  </form>
}