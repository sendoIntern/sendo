import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { GoogleOAuthProvider } from "@react-oauth/google";

import App from "./App.jsx";

createRoot(document.getElementById("root")).render(
  <GoogleOAuthProvider clientId="738724652425-poha3l9n6qs8hufr6d4ps0nqpukpb975.apps.googleusercontent.com">
    <App />
  </GoogleOAuthProvider>
);