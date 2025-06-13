import { Route, Routes } from "react-router-dom";
import Menuu from "./components/Menuu";
import Dashboard from "./components/Dashboard";
import Product from "./components/Product";
import Login from "./components/Login";
import { auth } from "./lib/auth";
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
    </Routes>
  );
}

export default App;
