import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";

function Editproducts() {
  const [data, setData] = useState([]);
  const [showCreateForm, setShowCreateForm] = useState(false);
  const [newItem, setNewItem] = useState({
    name: "",
    price: "",
    quantity: "",
    description: "",
    picture: null,
  });

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

  const handleUpdate = async (id, updatedData) => {
    try {
      await axiosInstance.put(`/item/${id}`, updatedData, {
        withCredentials: true,
      });
      console.log("Product updated successfully!");
    } catch (error) {
      console.error("Error updating product:", error);
    }
  };

  const handleDeleteByid = async (id) => {
    try {
      await axiosInstance.delete(`/item/${id}`, {
        withCredentials: true,
      });
      console.log("Product deleted successfully!");
    } catch (error) {
      console.error("Error deleting product:", error);
    }
  };

  const handleCreateItem = async () => {
    try {
      const formData = new FormData();
      formData.append("name", newItem.name);
      formData.append("price", newItem.price);
      formData.append("quantity", newItem.quantity);
      formData.append("description", newItem.description);
      formData.append("picture", newItem.picture); // file thực sự

      await axiosInstance.post("/item/createNewItem", formData, {
        withCredentials: true,
      });

      console.log("Product created successfully!");
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
    }
  };

  return (
    <>
      <h1>Product List</h1>
      <button onClick={() => setShowCreateForm(!showCreateForm)}>
        {showCreateForm ? "Close" : "New"}
      </button>

      {showCreateForm && (
        <div style={{ marginTop: "20px" }}>
          <h3>Create New Product</h3>
          <input
            type="text"
            placeholder="Name"
            value={newItem.name}
            onChange={(e) => setNewItem({ ...newItem, name: e.target.value })}
          />
          <br />
          <input
            type="text"
            placeholder="Quantity"
            value={newItem.quantity}
            onChange={(e) =>
              setNewItem({ ...newItem, quantity: e.target.value })
            }
          />
          <br />
          <input
            type="text"
            placeholder="Price"
            value={newItem.price}
            onChange={(e) => setNewItem({ ...newItem, price: e.target.value })}
          />
          <br />
          <input
            type="text"
            placeholder="Description"
            value={newItem.description}
            onChange={(e) =>
              setNewItem({ ...newItem, description: e.target.value })
            }
          />
          <br />
          <input
            type="file"
            accept="image/*"
            onChange={(e) => {
              const file = e.target.files[0];
              if (file) {
                console.log("Selected file:", file);
                setNewItem({ ...newItem, picture: file });
              } else {
                console.warn("No file selected");
              }
            }}
          />

          <br />
          <button onClick={handleCreateItem}>Submit</button>
        </div>
      )}
    </>
  );
}

export default Editproducts;
