import { Routes, Route, Navigate } from "react-router-dom";

import AccountCheck from "./pages/AccountCheck";
import Login from "./pages/Login";
import Register from "./pages/Register";
import RegisterPassword from "./pages/RegisterPassword";
import Dashboard from "./pages/Dashboard";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/account-check" replace />} />

      <Route path="/account-check" element={<AccountCheck />} />

      <Route path="/login" element={<Login />} />

      <Route path="/register" element={<Register />} />
      <Route path="/register-password" element={<RegisterPassword />} />

      <Route
        path="/register/password"
        element={<RegisterPassword />}
      />

      <Route path="/dashboard" element={<Dashboard />} />
    </Routes>
  );
}

export default App;