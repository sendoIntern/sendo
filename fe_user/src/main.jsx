import { createRoot } from "react-dom/client";
import App from "./App.jsx";
import { GoogleOAuthProvider } from "@react-oauth/google";
const clientId =
  "738724652425-poha3l9n6qs8hufr6d4ps0nqpukpb975.apps.googleusercontent.com";
createRoot(document.getElementById("root")).render(
  <GoogleOAuthProvider clientId={clientId}>
    <App />
  </GoogleOAuthProvider>
);
