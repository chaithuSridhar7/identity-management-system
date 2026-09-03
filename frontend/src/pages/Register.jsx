import { useState } from "react";
import AuthLayout from "../components/AuthLayout";
import ChoiceCard from "../components/ChoiceCard";
import { useNavigate } from "react-router-dom";

function Register() {
  const navigate = useNavigate();

  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");

  const handleContinue = (e) => {
    e.preventDefault();

    navigate("/register-password", {
      state: {
        username,
        email,
      },
    });
  };

  return (
    <AuthLayout title="Let's get you set up!">
      <ChoiceCard onClick={() => navigate("/login")}>
        Actually, I do have an account!
      </ChoiceCard>
      <div className="login-card">
        <form className="login-form" onSubmit={handleContinue}>
          <div className="login-field">
            <label htmlFor="username">Username</label>
            <input
              id="username"
              type="text"
              placeholder="Choose a username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>

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

          <button type="submit" className="login-button">
            Continue
          </button>
        </form>
      </div>

      
    </AuthLayout>
  );
}

export default Register;