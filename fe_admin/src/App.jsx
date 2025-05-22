import { Route, Routes } from "react-router-dom";
import Menuu from "./components/Menuu";
import Dashboard from "./components/Dashboard";
import Product from "./components/Product";
import Editproducts from "./components/Editproducts";
function App() {
  return (
    <Routes>
      <Route path="/" element={<Menuu />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/products" element={<Product />} />
      <Route path="/editProducts" element={<Editproducts />} />
    </Routes>
  );
}

export default App;
