import { useGoogleLogin } from "@react-oauth/google";
function Login() {
  const handleLogin = useGoogleLogin({
    onSuccess: (response) => console.log(response),

    onError: (error) => console.error(error),
  });
  // const handleSubmit = (e) => {
  //   e.preventDefault();
  //   window.location.href = "http://localhost:8080/auth/google/login";
  // };

  return (
    <div>
      // Button to trigger Google login when clicked
      <button onClick={() => handleLogin()}>Login with Google</button>
    </div>
  );
}

export default Login;
