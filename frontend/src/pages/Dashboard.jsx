function Dashboard() {
  const token = localStorage.getItem("accessToken");

  return (
    <main className="dashboard">
      <section className="dashboard-card">
        <h1>Welcome! 👋</h1>
        <p>You are successfully signed in.</p>

        <p className="dashboard-status">
          Authentication token: {token ? "✓ Active" : "✗ Missing"}
        </p>
      </section>
    </main>
  );
}

export default Dashboard;