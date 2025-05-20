import { Route, Routes } from "react-router-dom";
import Login from "./page/Login";
import OAuthCallback from "./page/OAuthCallback";
import Welcome from "./page/Welcome";

function App() {
  const token = localStorage.getItem("token");

  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/auth/callback" element={<OAuthCallback />} />
      <Route path="/welcome" element={token ? <Welcome /> : <Login />} />
    </Routes>
  );
}

export default App;
