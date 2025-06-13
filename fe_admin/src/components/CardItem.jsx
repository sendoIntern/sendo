import React, { useEffect, useState } from "react";
import {
  Box,
  Card,
  CardContent,
  CardMedia,
  Typography,
  Button,
} from "@mui/material";
import { axiosInstance } from "../lib/axios";

function CardItem() {
  const [data, setData] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);

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

  const fetchDataByid = async (id) => {
    try {
      const res = await axiosInstance.patch(`/item/getItemById/${id}`, {
        withCredentials: true,
      });
      setSelectedItem(res.data);
    } catch (error) {
      console.error("Error fetching products:", error);
    }
  };

  const showModal = (item) => {
    setSelectedItem(item);
    fetchDataByid(item.id);
    setIsModalOpen(true);
  };

  const handleOk = () => {
    setIsModalOpen(false);
    setSelectedItem(null);
  };

  const handleCancel = () => {
    setIsModalOpen(false);
    setSelectedItem(null);
  };

  return (
    <Box
      sx={{
        mt: 5,
        display: "flex",
        flexWrap: "wrap", // tự động xuống dòng
        justifyContent: "center", // căn giữa
        gap: 3,
        px: 2,
        overflowX: "hidden", // ẩn thanh cuộn ngang
        width: "100%", // giới hạn trong viewport
      }}
    >
      {data.map((item) => (
        <Card
          key={item.id}
          sx={{
            width: { xs: "100%", sm: 250, md: 300 },
            flexShrink: 0,
          }}
        >
          <CardMedia
            component="img"
            height="180"
            image={item.picture}
            alt={item.name}
            onClick={() => showModal(item)}
            sx={{ cursor: "pointer", objectFit: "cover" }}
          />
          <CardContent
            onClick={() => showModal(item)}
            sx={{ cursor: "pointer" }}
          >
            <Typography variant="h6" align="center">
              {item.name}
            </Typography>
            <Typography
              variant="body2"
              color="text.secondary"
              sx={{
                height: "100px",
                overflow: "hidden",
                textOverflow: "ellipsis",
                display: "-webkit-box",
                WebkitLineClamp: 3,
                WebkitBoxOrient: "vertical",
                mt: 1,
              }}
            >
              {item.description}
            </Typography>
            <Typography variant="subtitle1" color="text.primary" sx={{ mt: 1 }}>
              Price: {item.price}
            </Typography>
          </CardContent>
          <Button variant="outlined" fullWidth sx={{ color: "blue", mt: 1 }}>
            Buy
          </Button>
        </Card>
      ))}
    </Box>
  );
}

export default CardItem;
