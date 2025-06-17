import React, { useEffect, useState } from "react";
import {
  Box,
  Card,
  CardContent,
  CardMedia,
  Typography,
  Button,
  TextField,
  Select,
  MenuItem,
  InputLabel,
  FormControl,
} from "@mui/material";

import { axiosInstance } from "../lib/axios";

import { Modal, IconButton } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";

function CardItem() {
  const [data, setData] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [minPrice, setMinPrice] = useState("");
  const [maxPrice, setMaxPrice] = useState("");
  const [totalPages, setTotalPages] = useState(10);

  useEffect(() => {
    const fetchProducts = async () => {
      try {
        const res = await axiosInstance.get("/item/getAllItems", {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
          },
          withCredentials: true,
        });
        setData(res.data);
        setTotalPages(res.data.totalPages);
      } catch (error) {
        console.error("Error fetching products:", error);
      }
    };
    fetchProducts();
  }, []);

  const fetchDataByid = async (id) => {
    try {
      const res = await axiosInstance.get(`/item/getItemById/${id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
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
    <>
      {/* Layout chia 2 cột */}
      <Box
        sx={{
          display: "flex",
          alignItems: "flex-start",
          px: 2,
          mt: 4,
        }}
      >
        {/* Cột trái: Filter + Search */}
        <Box
          sx={{
            minWidth: 250,
            maxWidth: 300,
            mr: 3,
            display: "flex",
            flexDirection: "column",
            gap: 2,
          }}
        >
          <TextField
            label="Search by name"
            variant="outlined"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            fullWidth
          />
          <Typography variant="subtitle1">Filter by Price</Typography>
          <Box sx={{ display: "flex", gap: 2 }}>
            <TextField
              label="Min Price"
              type="number"
              value={minPrice}
              onChange={(e) => setMinPrice(e.target.value)}
              fullWidth
            />
            <TextField
              label="Max Price"
              type="number"
              value={maxPrice}
              onChange={(e) => setMaxPrice(e.target.value)}
              fullWidth
            />
          </Box>
        </Box>

        {/* Cột phải: Danh sách sản phẩm */}
        <Box
          sx={{
            flex: 1,
            display: "flex",
            flexWrap: "wrap",
            justifyContent: "center",
            gap: 3,
            overflowX: "hidden",
            width: "100%",
          }}
        >
          {data?.data?.map((item) => (
            <Card
              key={item.id}
              sx={{
                width: { xs: "100%", sm: 250, md: 300 },
                flexShrink: 0,
                transition: "transform 0.3s ease, box-shadow 0.3s ease",
                "&:hover": {
                  transform: "translateY(-8px)",
                  boxShadow: 6,
                },
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
                <Typography
                  variant="subtitle1"
                  color="text.primary"
                  sx={{ mt: 1 }}
                >
                  Price: {item.price}
                </Typography>
              </CardContent>
              <Button
                variant="outlined"
                fullWidth
                sx={{ color: "blue", mt: 1 }}
              >
                Buy
              </Button>
            </Card>
          ))}
        </Box>
      </Box>

      {/* Modal hiển thị chi tiết */}
      <Modal open={isModalOpen} onClose={handleCancel}>
        <Box
          sx={{
            position: "absolute",
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            width: 400,
            bgcolor: "background.paper",
            boxShadow: 24,
            borderRadius: 2,
            p: 3,
            maxHeight: "90vh",
            overflowY: "auto",
          }}
        >
          <Box sx={{ display: "flex", justifyContent: "flex-end" }}>
            <IconButton onClick={handleCancel}>
              <CloseIcon />
            </IconButton>
          </Box>
          {selectedItem && (
            <>
              <Typography variant="h6" gutterBottom>
                {selectedItem.name}
              </Typography>
              <img
                src={selectedItem.picture}
                alt={selectedItem.name}
                style={{
                  width: "100%",
                  borderRadius: "8px",
                  objectFit: "cover",
                }}
              />
              <Typography variant="body1" sx={{ mt: 2 }}>
                {selectedItem.description}
              </Typography>
              <Typography variant="subtitle1" sx={{ mt: 2 }}>
                Price: {selectedItem.price}
              </Typography>
            </>
          )}
        </Box>
      </Modal>
    </>
  );
}

export default CardItem;
