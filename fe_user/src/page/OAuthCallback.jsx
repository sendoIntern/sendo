import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const OAuthCallback = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const token = params.get("token");
    console.log(params);
    if (token) {
      localStorage.setItem("token", token);

      navigate("/welcome");
    } else {
      navigate("/login?error=missing_data");
    }
  }, [navigate]);

  return null;
};

export default OAuthCallback;
