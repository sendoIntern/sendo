import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import Carousell from "./Carousell";
import CardItem from "./CardItem";
function Product() {
  return (
    <>
      <Carousell />
      <CardItem />
    </>
  );
}

export default Product;
