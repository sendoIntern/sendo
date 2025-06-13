import React, { useEffect, useState } from "react";
import {
  Box,
  Card,
  CardContent,
  CardMedia,
  Typography,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
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
        display: "flex",
        flexWrap: "wrap",
        justifyContent: "space-around",
        gap: "20px",
        padding: 2,
      }}
    >
      {data.map((item) => (
        <Card key={item.id} sx={{ width: 300 }}>
          <CardMedia
            component="img"
            height="180"
            image={item.picture}
            alt={item.name}
            onClick={() => showModal(item)}
            sx={{ cursor: "pointer" }}
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
              sx={{ height: "100px", mt: 1 }}
            >
              {item.description}
            </Typography>
            <Typography variant="subtitle1" color="text.primary" sx={{ mt: 1 }}>
              Price: {item.price}
            </Typography>
          </CardContent>
          <Button
            variant="outlined"
            fullWidth
            sx={{ color: "blue", mt: 1 }}
            onClick={() => console.log("Buy:", item)}
          >
            Buy
          </Button>
        </Card>
      ))}

      <Dialog open={isModalOpen} onClose={handleCancel} maxWidth="sm" fullWidth>
        <DialogTitle>Chi tiết sản phẩm</DialogTitle>
        <DialogContent dividers>
          {selectedItem && (
            <Card>
              <CardMedia
                component="img"
                height="200"
                image={selectedItem.picture}
                alt={selectedItem.name}
              />
              <CardContent>
                <Typography variant="h6" align="center">
                  {selectedItem.name}
                </Typography>
                <Typography
                  variant="body2"
                  color="text.secondary"
                  sx={{ height: "100px", mt: 1 }}
                >
                  {selectedItem.description}
                </Typography>
                <Typography
                  variant="subtitle1"
                  color="text.primary"
                  sx={{ mt: 1 }}
                >
                  Price: {selectedItem.price}
                </Typography>
                <Button
                  variant="outlined"
                  fullWidth
                  sx={{ color: "blue", mt: 2 }}
                >
                  Buy
                </Button>
              </CardContent>
            </Card>
          )}
        </DialogContent>
        <DialogActions>
          <Button onClick={handleOk}>Đóng</Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
}

export default CardItem;
