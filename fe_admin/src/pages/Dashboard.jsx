import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import { Table } from "antd";
import Nav from "../components/Nav";
import { Button, Modal, Box, TextField, Typography } from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";
import IconButton from "@mui/material/IconButton";

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
  const [showImportFile, setShowImportFile] = useState(false);
  const [loading, setLoading] = useState(true);

  const [newItem, setNewItem] = useState({
    name: "",
    price: "",
    quantity: "",
    description: "",
    picture: null,
  });

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    setLoading(true);
    try {
      const res = await axiosInstance.get("/item/getAllItems", {
        withCredentials: true,
      });
      if (res.data) setData(res.data);
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
        withCredentials: true,
      });
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
      if (window.confirm("Do you want delete")) {
        await axiosInstance.delete(`/item/${id}`, {
          withCredentials: true,
        });
      }
    } catch (error) {
      console.error("Error deleting product:", error);
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  const handleImmportExcel = async () => {
    setLoading(true);
    try {
      const formData = new FormData();
      formData.append("file", fileImport);

      await axiosInstance.post("/item/import", formData, {
        withCredentials: true,
      });

      const isErr = await axiosInstance.get("/item/getErrorItems", {
        withCredentials: true,
      });

      if (isErr.data.length === 0) {
        alert("Import successful!");
      } else {
        alert("Import failed");
      }
      console.log(formData);
    } catch (error) {
      console.error("Error importing Excel file:", error);
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  return (
    <>
      <Nav />
      <h1>Product List</h1>

      <Button
        variant="contained"
        onClick={() => setShowCreateForm(true)}
        sx={{ mr: 2 }}
      >
        New
      </Button>

      <Button variant="outlined" component="label" sx={{ mr: 2 }}>
        Import Excel
        <input
          type="file"
          accept=".xlsx"
          hidden
          onChange={(e) => {
            const file = e.target.files[0];
            if (!file) return;
            setFileImport(file);
            handleImmportExcel();
          }}
        />
      </Button>

      {/* CREATE MODAL */}
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
            onChange={(e) => setNewItem({ ...newItem, price: e.target.value })}
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
          <Button variant="outlined" component="label" fullWidth sx={{ mt: 2 }}>
            Upload Image
            <input
              type="file"
              hidden
              accept="image/*"
              onChange={(e) =>
                setNewItem({ ...newItem, picture: e.target.files[0] })
              }
            />
          </Button>
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

      {/* UPDATE MODAL */}
      <Modal open={modalUpdateForm} onClose={() => setModalUpdateForm(false)}>
        <Box sx={modalStyle}>
          <IconButton
            onClick={() => setShowCreateForm(false)}
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
          <Button variant="outlined" component="label" fullWidth sx={{ mt: 2 }}>
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
                alert("Updated successfully!");
                setModalUpdateForm(false);
              } catch (err) {
                console.error(err);
              }
            }}
          >
            Submit Update
          </Button>
        </Box>
      </Modal>

      <Table dataSource={data} rowKey="id" style={{ marginTop: 20 }}>
        <Table.Column
          title="Image"
          dataIndex="picture"
          key="picture"
          render={(imageUrl) => (
            <img
              src={imageUrl}
              alt="item"
              style={{ width: 50, height: 50, objectFit: "cover" }}
            />
          )}
        />
        <Table.Column title="Name" dataIndex="name" key="name" />
        <Table.Column title="Price" dataIndex="price" key="price" />
        <Table.Column title="Quantity" dataIndex="quantity" key="quantity" />
        <Table.Column
          title="Description"
          dataIndex="description"
          key="description"
        />
        <Table.Column title="View" dataIndex="view" key="view" />
        <Table.Column
          title="Action"
          key="action"
          render={(_, item) => (
            <>
              <Button
                onClick={() => {
                  setSelectedItem(item);
                  setModalUpdateForm(true);
                }}
                sx={{ mr: 1 }}
              >
                Update
              </Button>
              <Button color="error" onClick={() => handleDeleteByid(item.id)}>
                Delete
              </Button>
            </>
          )}
        />
      </Table>
    </>
  );
}

export default Dashboard;
