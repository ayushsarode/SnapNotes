import { apiBaseUrl } from "../services/api";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useState } from "react";
import useAuthStore from "../store/authStore";

const Login: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const setToken = useAuthStore(state => state.setToken)

  const handleLogin = async () => {
    try {
      const response = await axios.post(`${apiBaseUrl}/login`, {
        email,
        password,
      });
      setToken(response.data.token)
      navigate("/dashboard"); 
    } catch (error) {
      console.error("Error during login:", error);
    }
  };

  return (
    <div>
      <div className="max-w-md mx-auto mt-8">
        <h1 className="text-2xl font-bold">Login</h1>
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
        <button onClick={handleLogin} className="btn btn-primary w-full mt-4">
          Login
        </button>
      </div>
    </div>
  );
};

export default Login;
