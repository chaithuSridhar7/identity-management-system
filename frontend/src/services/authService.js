const API_URL = "http://localhost:8080";

export async function login(email, password) {
  const response = await fetch(`${API_URL}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email,
      password,
    }),
  });

  if (!response.ok) {
    const message = await response.text();
    throw new Error(message || "Login failed");
  }

  return response.json();
}

export async function register(username, email, password) {
  const response = await fetch(`${API_URL}/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      username,
      email,
      password,
    }),
  });

  if (!response.ok) {
    const message = await response.text();
    throw new Error(message || "Registration failed");
  }

  return response.json();
}

export async function getCurrentUser() {
  const token = localStorage.getItem("accessToken");

  const response = await fetch("http://localhost:8080/me", {
    method: "GET",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error("Unable to get current user");
  }

  return response.json();
}

export async function updateCurrentUser(displayName) {
  const token = localStorage.getItem("accessToken");

  const response = await fetch("http://localhost:8080/me", {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({
      display_name: displayName,
    }),
  });

  if (!response.ok) {
    throw new Error("Unable to update profile");
  }

  return response.json();
}