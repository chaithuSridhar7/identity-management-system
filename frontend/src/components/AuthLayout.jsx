function AuthLayout({ title, children }) {
  return (
    <main className="auth-layout">
      <h1 className="auth-layout__title">{title}</h1>

      <div className="auth-layout__content">
        {children}
      </div>
    </main>
  );
}

export default AuthLayout;