import { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";

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
        // api trả object có key là items
        setItems(response.data || []);
      } catch (error) {
        console.error("Error fetching items:", error);
      }
    };
    fetchItems();
  }, []);
  console.log(items);

  const settings = {
    dots: true,
    infinite: true,
    speed: 300,
    slidesToShow: 1,
    slidesToScroll: 1,
    autoplay: true,
    autoplaySpeed: 3000,
  };

  return (
    <Slider {...settings}>
      {items?.items?.map((item, index) => (
        <div key={index}>
          <img
            src={item.picture}
            alt={item.name || `Item ${index}`}
            style={{ width: "100%", height: "400px", objectFit: "cover" }}
          />
        </div>
      ))}
    </Slider>
  );
}

export default Banner;
