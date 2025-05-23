// import { Route, Routes } from "react-router-dom";
// import Login from "./page/Login";
// import OAuthCallback from "./page/OAuthCallback";
// import Welcome from "./page/Welcome";

// function App() {
//   const token = localStorage.getItem("token");

//   return (
//     <Routes>
//       <Route path="/" element={<Login />} />
//       <Route path="/auth/callback" element={<OAuthCallback />} />
//       <Route path="/welcome" element={token ? <Welcome /> : <Login />} />
//     </Routes>
//   );
// }

// export default App;
import { useGoogleLogin } from "@react-oauth/google";
import { useState } from "react";
import axios from "axios";
const App = () => {
  const [setUserInfor] = useState(null);
  const login = useGoogleLogin({
    onSuccess: async (tokenResponse) => {
      try {
        const res = await axios.get(
          "https://www.googleapis.com/oauth2/v3/userinfo",
          {
            headers: {
              Authorization: `Bearer ${tokenResponse.access_token}`,
            },
          }
        );
        setUserInfor(res.data);
        console.log("Data của FE", res.data);

        // call api be
        const responce = await axios.post("", {
          name: res.data.name,
          email: res.data.email,
          picture: res.data.picture,
        });

        console.log("Data từ Be:", responce.data);
        localStorage.setItem("token", responce.data.access_token);
      } catch (err) {
        console.error("Lỗi đăng nhập", err);
      }
    },
    onError: (error) => console.error("Login lỗi:", error),
  });

  return (
    <div>
      <button onClick={() => login()}>Login with Google</button>
    </div>
  );
};

export default App;
