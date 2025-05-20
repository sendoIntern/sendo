function Login() {
  const handleSubmit = (e) => {
    e.preventDefault();
    window.location.href = "http://localhost:8080/auth/google/login";
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <button type="submit">Login with Google</button>
      </form>
    </div>
  );
}

export default Login;
