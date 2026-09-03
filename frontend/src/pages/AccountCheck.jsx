import { useNavigate } from "react-router-dom";

import AuthLayout from "../components/AuthLayout";
import ChoiceCard from "../components/ChoiceCard";

function AccountCheck() {
  const navigate = useNavigate();

  return (
    <AuthLayout title="Do you have an account with us?">
      <ChoiceCard onClick={() => navigate("/login")}>
        Yes
      </ChoiceCard>

      <ChoiceCard onClick={() => navigate("/register")}>
        No
      </ChoiceCard>
    </AuthLayout>
  );
}

export default AccountCheck;