import { Route, Routes } from "react-router-dom";
import Menuu from "./components/Menuu";
import Dashboard from "./components/Dashboard";
import Product from "./components/Product";
function App() {
  return (
    <Routes>
      <Route path="/" element={<Menuu />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/products" element={<Product />} />
    </Routes>
  );
}

export default App;
