import React, { useEffect, useState } from "react";
import Slider from "react-slick";
import { axiosInstance } from "../lib/axios";
import { Box, Typography, Card, CardContent, CardMedia } from "@mui/material";

function Carousell() {
  const [items, setItems] = useState([]);

  // useEffect(() => {
  //   const fetchItems = async () => {
  //     try {
  //       const response = await axiosInstance.get("/items/get3Items", {
  //         withCredentials: true,
  //       });
  //       setItems(response.data);
  //     } catch (error) {
  //       console.error("Error fetching items:", error);
  //     }
  //   };
  //   fetchItems();
  // }, []);

  const settings = {
    dots: true,
    infinite: true,
    speed: 500,
    autoplay: true,
    autoplaySpeed: 3000,
    slidesToShow: 1,
    slidesToScroll: 1,
  };

  return (
    <Box sx={{ maxWidth: 800, margin: "0 auto", mt: 4 }}>
      <Slider {...settings}>
        {items.map((item, index) => (
          <Box key={index} px={2}>
            <Card>
              <CardMedia
                component="img"
                height="220"
                image={item.picture}
                alt={item.name}
              />
              <CardContent sx={{ textAlign: "center" }}>
                <Typography variant="h6">{item.name}</Typography>
                <Typography variant="body2" color="text.secondary">
                  {item.description}
                </Typography>
              </CardContent>
            </Card>
          </Box>
        ))}
      </Slider>
    </Box>
  );
}

export default Carousell;
