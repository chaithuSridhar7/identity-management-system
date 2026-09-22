import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  getCurrentUser,
  updateCurrentUser,
} from "../services/authService";

function Dashboard() {
  const navigate = useNavigate();

  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  const [displayName, setDisplayName] = useState("");
  const [profileImage, setProfileImage] = useState(null);
  const [editingProfile, setEditingProfile] = useState(false);

  useEffect(() => {
    const loadUser = async () => {
      try {
        const currentUser = await getCurrentUser();

        setUser(currentUser);
        setDisplayName(
          currentUser.display_name || currentUser.username
        );
      } catch (error) {
        console.error("Failed to load user:", error);
        localStorage.removeItem("accessToken");
        navigate("/login");
      } finally {
        setLoading(false);
      }
    };

    loadUser();
  }, [navigate]);

  const handleSignOut = () => {
    localStorage.removeItem("accessToken");
    navigate("/login");
  };

  const handleProfileImageChange = (event) => {
    const file = event.target.files[0];

    if (!file) {
      return;
    }

    const imageURL = URL.createObjectURL(file);
    setProfileImage(imageURL);
  };

  const handleSaveProfile = async () => {
    try {
      const updatedUser = await updateCurrentUser(displayName);

      setUser(updatedUser);
      setDisplayName(updatedUser.display_name || updatedUser.username);
      setEditingProfile(false);
    } catch (error) {
      console.error("Failed to update profile:", error);
    }
  };

  if (loading) {
    return <p>Loading...</p>;
  }

  return (
    <main className="dashboard">
      <section className="dashboard-card">

        <h1>Welcome back, {displayName}! 👋</h1>

        <div className="profile-section">

          <label className="profile-picture">
            {profileImage ? (
              <img src={profileImage} alt="Profile" />
            ) : (
              <div className="profile-placeholder">
                {user.username.charAt(0).toUpperCase()}
              </div>
            )}

            <input
              type="file"
              accept="image/*"
              onChange={handleProfileImageChange}
              hidden
            />
          </label>

          <div className="profile-info">

            {editingProfile ? (
              <div className="profile-name-edit">
                <input
                  id="display-name"
                  type="text"
                  value={displayName}
                  onChange={(e) => setDisplayName(e.target.value)}
                  autoFocus
                />

                <button
                  className="dashboard-edit-button"
                  onClick={handleSaveProfile}
                >
                  Save
                </button>

                <button
                  className="dashboard-edit-button"
                  onClick={() => {
                    setDisplayName(
                      user.display_name || user.username
                    );
                    setEditingProfile(false);
                  }}
                >
                  Cancel
                </button>
              </div>
            ) : (
              <div className="profile-name-row">
                <h2>{displayName}</h2>

                <button
                  className="dashboard-edit-button"
                  onClick={() => setEditingProfile(true)}
                >
                  Edit
                </button>
              </div>
            )}

            <p>{user.email}</p>

          </div>

</div>

        <button
          onClick={handleSignOut}
          className="dashboard-button"
        >
          Sign out
        </button>

      </section>
    </main>
  );
}

export default Dashboard;