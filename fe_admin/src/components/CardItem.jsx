import React, { useEffect, useState } from "react";
import { Button, Card } from "antd";
import { axiosInstance } from "../lib/axios";
import Meta from "antd/es/card/Meta";

function CardItem() {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const res = await axiosInstance.get("/item/getAllItems", {
          withCredentials: true,
        });
        setData(res.data);
      } catch (error) {
        console.error("Error fetching products:", error);
      }
    };

    fetchProducts();
  }, []);

  return (
    <div
      style={{
        display: "flex",
        justifyItems: "space-between",
        flexWrap: "wrap",
      }}
    >
      {data.map((product) => (
        <Card
          key={product.ID}
          style={{ width: 300 }}
          cover={<img alt="Not Found" src={product.Picture} />}
        >
          <Meta title={product.Name} description={product.Description} />
          <p>Price: {product.Price}</p>
          <p>Recommend: {product.Recommend}</p>
          <Button style={{ marginTop: "10px", width: "100%" }} type="primary">
            Buy
          </Button>
        </Card>
      ))}
    </div>
  );
}

export default CardItem;
