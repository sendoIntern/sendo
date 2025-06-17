import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import {
  Button,
  Modal,
  Box,
  TextField,
  Typography,
  Stack,
  IconButton,
  Pagination,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
} from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";
import Nav from "../components/Nav";

const modalStyle = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  borderRadius: 2,
  boxShadow: 24,
  p: 4,
};

function Dashboard() {
  const [data, setData] = useState([]);
  const [showCreateForm, setShowCreateForm] = useState(false);
  const [modalUpdateForm, setModalUpdateForm] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const [fileImport, setFileImport] = useState(null);
  const [loading, setLoading] = useState(true);
  // Search, filter, and pagination states
  const [searchTerm, setSearchTerm] = useState("");
  const [minPrice, setMinPrice] = useState("");
  const [maxPrice, setMaxPrice] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const limit = 6;

  const [newItem, setNewItem] = useState({
    name: "",
    price: "",
    quantity: "",
    description: "",
    picture: null,
  });

  useEffect(() => {
    fetchProducts();
  }, [currentPage, searchTerm, minPrice, maxPrice]);

  const fetchProducts = async () => {
    setLoading(true);
    try {
      const params = {
        page: currentPage,
        limit,
        search: searchTerm || undefined,
        minPrice: minPrice || undefined,
        maxPrice: maxPrice || undefined,
      };
      const res = await axiosInstance.get("/item/getAllItems", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        params,
        withCredentials: true,
      });
      if (res.data) {
        setData(res.data);
        setTotalPages(res.data.pagination?.total_pages || 1);
      }
    } catch (error) {
      console.error("Error fetching products:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreateItem = async () => {
    setLoading(true);
    try {
      const formData = new FormData();
      formData.append("name", newItem.name);
      formData.append("price", newItem.price);
      formData.append("quantity", newItem.quantity);
      formData.append("description", newItem.description);
      formData.append("picture", newItem.picture);

      await axiosInstance.post("/item/createNewItem", formData, {
        withCredentials: true,
      });

      alert("Product created successfully!");
      setNewItem({
        name: "",
        price: "",
        description: "",
        picture: null,
        quantity: "",
      });
      setShowCreateForm(false);
    } catch (error) {
      console.error("Error creating product:", error);
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  const handleUpdate = async (id, updatedData) => {
    setLoading(true);
    try {
      await axiosInstance.put(`/item/${id}`, updatedData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });
      alert("Updated successfully!");
      setModalUpdateForm(false);
    } catch (error) {
      console.error("Error updating product:", error);
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  const handleDeleteByid = async (id) => {
    setLoading(true);
    try {
      if (window.confirm("Do you want to delete?")) {
        await axiosInstance.delete(`/item/${id}`, {
          withCredentials: true,
        });
        alert("Deleted successfully!");
      }
    } catch (error) {
      console.error("Error deleting product:", error);
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  const handleImportExcel = async (file) => {
    setLoading(true);
    try {
      const formData = new FormData();
      formData.append("file", file);

      console.log("Uploading file:", file.name);

      await axiosInstance.post("/item/import", formData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });

      const isErr = await axiosInstance.get("/item/getErrorItems", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });

      if (isErr.data.length === 0) {
        alert("Import successful!");
      } else {
        alert("Import failed");
      }
    } catch (error) {
      console.error("Error importing Excel file:", error);
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  const handlePageChange = (event, value) => {
    setCurrentPage(value);
  };

  return (
    <>
      <Nav />
      <Box sx={{ p: 3 }}>
        <Typography variant="h4" gutterBottom>
          Product List
        </Typography>

        {/* Search and Filter Section */}
        <Box
          sx={{
            display: "flex",
            flexWrap: "wrap",
            gap: 2,
            mb: 3,
            alignItems: "center",
          }}
        >
          <TextField
            label="Search by name"
            variant="outlined"
            value={searchTerm}
            onChange={(e) => {
              setSearchTerm(e.target.value);
              setCurrentPage(1); // Reset to page 1 on search
            }}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                fetchProducts();
              }
            }}
            sx={{ minWidth: 200 }}
          />
          <TextField
            label="Min Price"
            type="number"
            value={minPrice}
            onChange={(e) => {
              setMinPrice(e.target.value);
              setCurrentPage(1); // Reset to page 1 on filter
            }}
            sx={{ minWidth: 120 }}
          />
          <TextField
            label="Max Price"
            type="number"
            value={maxPrice}
            onChange={(e) => {
              setMaxPrice(e.target.value);
              setCurrentPage(1); // Reset to page 1 on filter
            }}
            sx={{ minWidth: 120 }}
          />
          <Button
            variant="contained"
            onClick={() => setShowCreateForm(true)}
            sx={{ ml: "auto" }}
          >
            New
          </Button>
          <Button variant="outlined" component="label">
            Import Excel
            <input
              type="file"
              accept=".xlsx"
              hidden
              onChange={(e) => {
                const file = e.target.files[0];
                if (!file) return;
                setFileImport(file);
                handleImportExcel(file);
              }}
            />
          </Button>
        </Box>

        {fileImport && (
          <Stack direction="row" alignItems="center" spacing={1} sx={{ mb: 2 }}>
            <Typography variant="body2">
              Selected file: {fileImport.name}
            </Typography>
            <IconButton
              size="small"
              onClick={() => setFileImport(null)}
              aria-label="remove selected file"
            >
              <CloseIcon fontSize="small" />
            </IconButton>
          </Stack>
        )}

        {/* Create Modal */}
        <Modal open={showCreateForm} onClose={() => setShowCreateForm(false)}>
          <Box sx={modalStyle}>
            <IconButton
              onClick={() => setShowCreateForm(false)}
              sx={{ position: "absolute", top: 8, right: 8 }}
            >
              <CloseIcon />
            </IconButton>
            <Typography variant="h6" mb={2}>
              Create New Product
            </Typography>
            <TextField
              label="Name"
              fullWidth
              value={newItem.name}
              onChange={(e) => setNewItem({ ...newItem, name: e.target.value })}
              margin="dense"
            />
            <TextField
              label="Quantity"
              fullWidth
              value={newItem.quantity}
              onChange={(e) =>
                setNewItem({ ...newItem, quantity: e.target.value })
              }
              margin="dense"
            />
            <TextField
              label="Price"
              fullWidth
              value={newItem.price}
              onChange={(e) =>
                setNewItem({ ...newItem, price: e.target.value })
              }
              margin="dense"
            />
            <TextField
              label="Description"
              fullWidth
              value={newItem.description}
              onChange={(e) =>
                setNewItem({ ...newItem, description: e.target.value })
              }
              margin="dense"
            />
            <Button
              variant="outlined"
              component="label"
              fullWidth
              sx={{ mt: 2 }}
            >
              Upload Image
              <input
                type="file"
                hidden
                accept="image/*"
                onChange={(e) => {
                  const file = e.target.files[0];
                  if (!file) return;
                  setNewItem({ ...newItem, picture: file });
                }}
              />
            </Button>
            {newItem.picture && (
              <Typography variant="body2" sx={{ mt: 1 }}>
                Selected file: {newItem.picture.name}
              </Typography>
            )}
            <Button
              variant="contained"
              fullWidth
              sx={{ mt: 2 }}
              onClick={handleCreateItem}
            >
              Submit
            </Button>
          </Box>
        </Modal>

        {/* Update Modal */}
        <Modal open={modalUpdateForm} onClose={() => setModalUpdateForm(false)}>
          <Box sx={modalStyle}>
            <IconButton
              onClick={() => setModalUpdateForm(false)}
              sx={{ position: "absolute", top: 8, right: 8 }}
            >
              <CloseIcon />
            </IconButton>
            <Typography variant="h6" mb={2}>
              Update Product
            </Typography>
            <TextField
              label="Name"
              fullWidth
              value={selectedItem?.name || ""}
              onChange={(e) =>
                setSelectedItem({ ...selectedItem, name: e.target.value })
              }
              margin="dense"
            />
            <TextField
              label="Quantity"
              fullWidth
              value={selectedItem?.quantity || ""}
              onChange={(e) =>
                setSelectedItem({ ...selectedItem, quantity: e.target.value })
              }
              margin="dense"
            />
            <TextField
              label="Price"
              fullWidth
              value={selectedItem?.price || ""}
              onChange={(e) =>
                setSelectedItem({ ...selectedItem, price: e.target.value })
              }
              margin="dense"
            />
            <TextField
              label="Description"
              fullWidth
              value={selectedItem?.description || ""}
              onChange={(e) =>
                setSelectedItem({
                  ...selectedItem,
                  description: e.target.value,
                })
              }
              margin="dense"
            />
            <Button
              variant="outlined"
              component="label"
              fullWidth
              sx={{ mt: 2 }}
            >
              Upload New Image
              <input
                type="file"
                hidden
                accept="image/*"
                onChange={(e) =>
                  setSelectedItem({
                    ...selectedItem,
                    picture: e.target.files[0],
                  })
                }
              />
            </Button>
            <Button
              variant="contained"
              fullWidth
              sx={{ mt: 2 }}
              onClick={async () => {
                try {
                  const formData = new FormData();
                  formData.append("name", selectedItem.name);
                  formData.append("price", selectedItem.price);
                  formData.append("quantity", selectedItem.quantity);
                  formData.append("description", selectedItem.description);
                  if (selectedItem.picture instanceof File) {
                    formData.append("picture", selectedItem.picture);
                  }
                  await handleUpdate(selectedItem.id, formData);
                } catch (err) {
                  console.error(err);
                }
              }}
            >
              Submit Update
            </Button>
          </Box>
        </Modal>

        {/* Product Table */}
        <TableContainer component={Paper} sx={{ mt: 3 }}>
          <Table sx={{ minWidth: 650 }} aria-label="product table">
            <TableHead>
              <TableRow>
                <TableCell>Image</TableCell>
                <TableCell>Name</TableCell>
                <TableCell>Price</TableCell>
                <TableCell>Quantity</TableCell>
                <TableCell>Description</TableCell>
                <TableCell>View</TableCell>
                <TableCell>Action</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {loading ? (
                <TableRow>
                  <TableCell colSpan={7} align="center">
                    Loading...
                  </TableCell>
                </TableRow>
              ) : (data.data || []).length === 0 ? (
                <TableRow>
                  <TableCell colSpan={7} align="center">
                    No data available
                  </TableCell>
                </TableRow>
              ) : (
                (data.data || []).map((item) => (
                  <TableRow key={item.id}>
                    <TableCell>
                      <img
                        src={item.picture}
                        alt={item.name}
                        style={{ width: 50, height: 50, objectFit: "cover" }}
                      />
                    </TableCell>
                    <TableCell>{item.name}</TableCell>
                    <TableCell>{item.price}</TableCell>
                    <TableCell>{item.quantity}</TableCell>
                    <TableCell>{item.description}</TableCell>
                    <TableCell>{item.view}</TableCell>
                    <TableCell>
                      <Button
                        onClick={() => {
                          setSelectedItem(item);
                          setModalUpdateForm(true);
                        }}
                        sx={{ mr: 1 }}
                      >
                        Update
                      </Button>
                      <Button
                        color="error"
                        onClick={() => handleDeleteByid(item.id)}
                      >
                        Delete
                      </Button>
                    </TableCell>
                  </TableRow>
                ))
              )}
            </TableBody>
          </Table>
        </TableContainer>

        {/* Pagination */}
        <Box sx={{ display: "flex", justifyContent: "center", mt: 3, mb: 2 }}>
          <Pagination
            count={totalPages}
            page={currentPage}
            onChange={handlePageChange}
            color="primary"
            variant="outlined"
            shape="rounded"
            sx={{
              "& .MuiPagination-ul": {
                justifyContent: "center",
              },
            }}
          />
        </Box>
      </Box>
    </>
  );
}

export default Dashboard;
