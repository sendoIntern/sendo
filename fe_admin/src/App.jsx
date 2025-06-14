import { Routes, Route } from "react-router-dom";
import Product from "./pages/Product";
import Dashboard from "./pages/Dashboard";

import Login from "./pages/Login";
import RoleRoute from "./components/RoleRoute";

function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      {/* Routes dành cho user */}
      <Route element={<RoleRoute allowRoles={["user"]} />}>
        <Route path="/product" element={<Product />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Route>

      <Route path="*" element={<Login />} />
    </Routes>
  );
}

export default App;
