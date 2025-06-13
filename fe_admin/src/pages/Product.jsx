import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import Banner from "../components/Banner";
import CardItem from "../components/CardItem";
import Nav from "../components/Nav";
function Product() {
  return (
    <>
      <Nav />
      <Banner />
      <CardItem />
    </>
  );
}

export default Product;
