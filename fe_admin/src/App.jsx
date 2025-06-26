import { Routes, Route, Navigate } from "react-router-dom";
import Product from "./pages/Product";
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import RoleRoute from "./components/RoleRoute";

function App() {
  return (
    <Routes>
      <Route element={<RoleRoute allowRoles={["user", "admin"]} />}>
        <Route path="/product" element={<Product />} />
      </Route>

      <Route element={<RoleRoute allowRoles={["admin"]} />}>
        <Route path="/dashboard" element={<Dashboard />} />
      </Route>

      <Route path="/login" element={<Login />} />
      <Route path="*" element={<Navigate to="/product" replace />} />
    </Routes>
  );
}

export default App;
