import { BrowserRouter, Route, Router, Routes } from "react-router-dom";
import Home from "./pages/Home";
import Register from "./pages/Register";
import Login from "./pages/Login";
import Notes from "./pages/Notes";

import "./App.css";
import ProtectedRoute from "./components/ProtectedRoute";
import Dashboard from "./pages/Dashboard";

const App: React.FC = () => {
  return (
    <div className="">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />

          <Route
            path="/dashboard"
            element={
              <ProtectedRoute>
                <Dashboard />
              </ProtectedRoute>
            }
          ></Route>

          <Route
            path="/dashboard/notes"
            element={
              <ProtectedRoute>
                <Notes />
              </ProtectedRoute>
            }
          ></Route>
          <Route path="/login" element={<Login />}>
            Login
          </Route>
          <Route path="/register" element={<Register />}>
            Register
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
};

export default App;
