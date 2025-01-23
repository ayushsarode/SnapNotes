import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';
import { apiBaseUrl } from '../services/api';

const Register: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [userName, setUsername] = useState('');
  const navigate = useNavigate();

  const handleRegister = async () => {
    try {
      await axios.post(`${apiBaseUrl}/register`, { userName, email, password });
      navigate('/login');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-8">
      <h1 className="text-2xl font-bold">Register</h1>
      <input
        type="text"
        placeholder="Name"
        value={userName}
        onChange={(e) => setUsername(e.target.value)}
        className="input input-bordered w-full mt-4"
      />
      <input
        type="email"
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        className="input input-bordered w-full mt-4"
      />
      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        className="input input-bordered w-full mt-4"
      />
      <button onClick={handleRegister} className="btn btn-primary w-full mt-4">Register</button>
    </div>
  );
};

export default Register;
