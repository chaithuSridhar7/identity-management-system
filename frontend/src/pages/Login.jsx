import { useState } from "react";
import { login } from "../services/authService";
import AuthLayout from "../components/AuthLayout";
import ChoiceCard from "../components/ChoiceCard";
import { useNavigate } from "react-router-dom";

function Login() {
  const navigate = useNavigate();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleLogin = async (e) => {
    e.preventDefault();

    setError("");
    setLoading(true);

    try {
      const response = await login(email, password);

      console.log("Login successful:", response);

      localStorage.setItem("accessToken", response.token);

      navigate("/dashboard");
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AuthLayout title="Welcome back! Sign in to continue">
      <div className="login-card">
        <form className="login-form" onSubmit={handleLogin}>
          <div className="login-field">
            <label htmlFor="email">Email</label>
            <input
              id="email"
              type="email"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div className="login-field">
            <label htmlFor="password">Password</label>
            <input
              id="password"
              type="password"
              placeholder="Enter your password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}

            />
          </div>

          <button type="submit" className="login-button" disabled={loading}>
            {loading ? "Signing in..." : "Sign me in"}
          </button>
        </form>
        {error && <p className="login-error">{error}</p>}
      </div>

      <ChoiceCard onClick={() => navigate("/register")}>
        Actually I don't have an account, oops!
      </ChoiceCard>
    </AuthLayout>
  );
}

export default Login;