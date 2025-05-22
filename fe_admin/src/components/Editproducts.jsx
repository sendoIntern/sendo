import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import { Table } from "antd";

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
  const [, setLoading] = useState(true);
  const [modalUpdateForm, setModalUpdateForm] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    setLoading(true);
    try {
      const res = await axiosInstance.get("/item/getAllItems", {
        withCredentials: true,
      });
      if (res.data !== null) setData(res.data);
    } catch (error) {
      console.error("Error fetching products:", error);
    } finally {
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

  const handleCreateItem = async () => {
    setLoading(true);
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

      window.alert("Product created successfully!");
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

      <Table dataSource={data}>
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
        <Table.Column
          title="Action"
          key="action"
          render={(_, item) => (
            <div>
              <button
                onClick={() => {
                  setSelectedItem(item);
                  setModalUpdateForm(true);
                }}
              >
                Update
              </button>

              <button onClick={() => handleDeleteByid(item.id)}>Delete</button>
            </div>
          )}
        />
      </Table>

      {modalUpdateForm && selectedItem && (
        <div style={{ marginTop: "20px" }}>
          <h3>Update Product</h3>
          <input
            type="text"
            placeholder="Name"
            value={selectedItem.name}
            onChange={(e) =>
              setSelectedItem({ ...selectedItem, name: e.target.value })
            }
          />
          <br />
          <input
            type="text"
            placeholder="Quantity"
            value={selectedItem.quantity}
            onChange={(e) =>
              setSelectedItem({ ...selectedItem, quantity: e.target.value })
            }
          />
          <br />
          <input
            type="text"
            placeholder="Price"
            value={selectedItem.price}
            onChange={(e) =>
              setSelectedItem({ ...selectedItem, price: e.target.value })
            }
          />
          <br />
          <input
            type="text"
            placeholder="Description"
            value={selectedItem.description}
            onChange={(e) =>
              setSelectedItem({ ...selectedItem, description: e.target.value })
            }
          />
          <br />
          <input
            type="file"
            accept="image/*"
            onChange={(e) => {
              const file = e.target.files[0];
              setSelectedItem({ ...selectedItem, picture: file });
            }}
          />
          <br />
          <button
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
          </button>
          <button onClick={() => setModalUpdateForm(false)}>Cancel</button>
        </div>
      )}
    </>
  );
}

export default Editproducts;
