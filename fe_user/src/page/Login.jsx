import { useEffect } from "react";

function Login() {
  // const [data, setData] = useState;
  // const handleLogin = () => {};

  useEffect(() => {}, []);

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
    >
      <a
        href="http://localhost:8080/auth/google/login"
        style={{
          padding: "10px 20px",
          backgroundColor: "#4285F4",
          color: "white",
          borderRadius: "4px",
          textDecoration: "none",
          fontWeight: "bold",
          fontSize: "16px",
          fontFamily: "Arial, sans-serif",
        }}
      >
        Đăng nhập bằng Google
      </a>
    </div>
  );
}

export default Login;
