import React, { useState } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';

const Login = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/login', {
        username,
        password,
      });
      if (response.status === 200) {
        localStorage.setItem('token', response.data.token); // Save the token
        localStorage.setItem('user_id', response.data.user_id) // Save the user_id
        console.log("Token stored", response.data.token) // Log the token
        console.log("User ID stored", response.data.user_id) // Log the user_id
        navigate('/search');
      }
    } catch (error) {
      console.error('There was an error logging in!', error);
    }
  };

  return (
    <div style={{ backgroundColor: 'black', color: 'white', minHeight: '100vh', padding: '20px' }}>
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>
      <input
        type="text"
        placeholder="Username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <button type="submit">Login</button>
    </form>
    <p>Don't have an account? <Link to="/register" style={{ color: 'white' }}>Register here</Link></p>
    </div>
  );
};

export default Login;