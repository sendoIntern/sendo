import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import Carousell from "./Carousell";
import CardItem from "./CardItem";
function Product() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const res = await axiosInstance.get("/item/getAllItems", {
          withCredentials: true,
        });
        setProducts(res.data);
      } catch (error) {
        console.error("Error fetching products:", error);
      }
    };

    fetchProducts();
  }, []);

  console.log(products);
  return (
    <>
      <Carousell />
      <CardItem />
    </>
  );
}

export default Product;
