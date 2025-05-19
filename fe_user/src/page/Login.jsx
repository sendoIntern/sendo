// App.tsx
import { GoogleLogin } from "@react-oauth/google";
import { axiosInstance } from "../lib/axios";

function Login() {
  const handleLogin = async (credentialResponse) => {
    const { credential } = credentialResponse;
    // Gửi credential (id_token) về server để xác minh
    const res = await axiosInstance.post("/api/auth/google", {
      token: credential,
    });
    console.log("Login success:", res.data);
  };

  return (
    <div>
      <GoogleLogin
        onSuccess={handleLogin}
        onError={() => console.log("Login Failed")}
      />
    </div>
  );
}

export default Login;
