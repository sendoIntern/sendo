import { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import { Carousel } from "antd";
import "antd/dist/reset.css";

function Banner() {
  const [items, setItems] = useState([]);

  useEffect(() => {
    const fetchItems = async () => {
      try {
        const response = await axiosInstance.get("/item/getItemDesc", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
          },
          withCredentials: true,
        });
        setItems(response.data.data || []);
      } catch (error) {
        console.error("Error fetching items:", error);
      }
    };
    fetchItems();
  }, []);

  return (
    <Carousel autoplay>
      {items?.map((item, index) => (
        <div key={index}>
          <img
            src={item.picture}
            alt={item.name || `Item ${index}`}
            style={{
              width: "100%",
              height: "400px",
              objectFit: "cover",
            }}
          />
        </div>
      ))}
    </Carousel>
  );
}

export default Banner;
