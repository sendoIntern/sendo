import { Route, Routes } from "react-router-dom";
import { auth } from "./lib/auth";
import Product from "./pages/Product";
import Dashboard from "./pages/Dashboard";

function App() {
  const role = auth.getRole();
  return (
    <Routes>
      // check role để render các route khác nhau
      {/* <Route path="/" element={role === "admin" ? <Login /> : <Login />} /> */}
      {/* <Route
        path="/product"
        element={role === "admin" ? <Product /> : <Login />}
      /> */}
      {/* <Route
        path="/dashboard"
        element={role === "admin" ? <Dashboard /> : <Login />}
      /> */}
      {/* <Route path="/menuu" element={role === "admin" ? <Menuu /> : <Login />} /> */}
      <Route path="/product" element={<Product />} />
      <Route path="/" element={<Dashboard />} />
    </Routes>
  );
}

export default App;
