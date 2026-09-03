import { useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import AuthLayout from "../components/AuthLayout";
import ChoiceCard from "../components/ChoiceCard";
import { register } from "../services/authService";

function RegisterPassword() {
  const navigate = useNavigate();
  const location = useLocation();

  const { username, email } = location.state || {};

  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleRegister = async (e) => {
    e.preventDefault();

    setError("");

    if (password !== confirmPassword) {
      setError("Passwords do not match.");
      return;
    }

    setLoading(true);

    try {
      await register(username, email, password);

      navigate("/login");
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  if (!username || !email) {
    return (
      <AuthLayout title="Let's get you set up!">
        <div className="login-card">
          <p className="login-error">
            Please start registration from the beginning.
          </p>

          <button
            className="login-button"
            onClick={() => navigate("/register")}
          >
            Back to registration
          </button>
        </div>
      </AuthLayout>
    );
  }

  return (
    <AuthLayout title="Almost there! Create your password">
      <ChoiceCard onClick={() => navigate("/login")}>
        Never mind, I'll sign in!
      </ChoiceCard>
      <div className="login-card">
        <form className="login-form" onSubmit={handleRegister}>
          <div className="login-field">
            <label htmlFor="password">Password</label>

            <input
              id="password"
              type="password"
              placeholder="Create a password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          <div className="login-field">
            <label htmlFor="confirmPassword">Confirm Password</label>

            <input
              id="confirmPassword"
              type="password"
              placeholder="Re-enter your password"
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
            />
          </div>

          {error && <p className="login-error">{error}</p>}

          <button
            type="submit"
            className="login-button"
            disabled={loading}
          >
            {loading ? "Creating account..." : "Create my account"}
          </button>
        </form>
      </div>

    </AuthLayout>
  );
}

export default RegisterPassword;