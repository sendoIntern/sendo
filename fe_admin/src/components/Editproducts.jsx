import React, { useEffect, useState } from "react";
import {
  Table,
  Button,
  Modal,
  Form,
  Input,
  InputNumber,
  Upload,
  message,
} from "antd";
import { UploadOutlined } from "@ant-design/icons";
import { axiosInstance } from "../lib/axios";

function Editproducts() {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(false);

  const [isCreateOpen, setIsCreateOpen] = useState(false);
  const [isUpdateOpen, setIsUpdateOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);

  const [createForm] = Form.useForm();
  const [updateForm] = Form.useForm();

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    setLoading(true);
    try {
      const res = await axiosInstance.get("/item/getAllItems", {
        withCredentials: true,
      });
      setData(res.data.data || []);
    } catch (err) {
      console.error("Error fetching:", err);
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async (values) => {
    const formData = new FormData();
    for (let key in values) {
      if (key === "picture") {
        formData.append("picture", values.picture.file.originFileObj);
      } else {
        formData.append(key, values[key]);
      }
    }

    try {
      await axiosInstance.post("/item/createNewItem", formData, {
        withCredentials: true,
      });
      message.success("Created successfully");
      setIsCreateOpen(false);
      createForm.resetFields();
      fetchProducts();
    } catch (err) {
      message.error("Create failed");
    }
  };

  const handleUpdate = async (values) => {
    const formData = new FormData();
    for (let key in values) {
      if (key === "picture" && values.picture?.file) {
        formData.append("picture", values.picture.file.originFileObj);
      } else {
        formData.append(key, values[key]);
      }
    }

    try {
      await axiosInstance.put(`/item/${selectedItem.id}`, formData, {
        withCredentials: true,
      });
      message.success("Updated successfully");
      setIsUpdateOpen(false);
      updateForm.resetFields();
      fetchProducts();
    } catch (err) {
      message.error("Update failed");
    }
  };

  const handleDelete = async (id) => {
    Modal.confirm({
      title: "Confirm deletion",
      onOk: async () => {
        try {
          await axiosInstance.delete(`/item/${id}`, {
            withCredentials: true,
          });
          message.success("Deleted");
          fetchProducts();
        } catch (err) {
          message.error("Delete failed");
        }
      },
    });
  };

  const columns = [
    {
      title: "Image",
      dataIndex: "picture",
      key: "picture",
      render: (url) => (
        <img src={url} alt="item" style={{ width: 50, height: 50 }} />
      ),
    },
    { title: "Name", dataIndex: "name", key: "name" },
    { title: "Price", dataIndex: "price", key: "price" },
    { title: "Quantity", dataIndex: "quantity", key: "quantity" },
    { title: "Description", dataIndex: "description", key: "description" },
    {
      title: "Action",
      key: "action",
      render: (_, item) => (
        <>
          <Button
            type="link"
            onClick={() => {
              setSelectedItem(item);
              updateForm.setFieldsValue(item);
              setIsUpdateOpen(true);
            }}
          >
            Edit
          </Button>
          <Button type="link" danger onClick={() => handleDelete(item.id)}>
            Delete
          </Button>
        </>
      ),
    },
  ];

  return (
    <>
      <h1>Product List</h1>
      <Button type="primary" onClick={() => setIsCreateOpen(true)}>
        New Product
      </Button>
      <br />
      <br />
      <Table
        dataSource={data}
        columns={columns}
        rowKey="id"
        loading={loading}
        bordered
      />

      {/* Create Modal */}
      <Modal
        title="Create Product"
        open={isCreateOpen}
        onCancel={() => setIsCreateOpen(false)}
        onOk={() => createForm.submit()}
      >
        <Form form={createForm} layout="vertical" onFinish={handleCreate}>
          <Form.Item name="name" label="Name" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[{ required: true }]}
          >
            <InputNumber min={0} style={{ width: "100%" }} />
          </Form.Item>
          <Form.Item name="price" label="Price" rules={[{ required: true }]}>
            <InputNumber min={0} style={{ width: "100%" }} />
          </Form.Item>
          <Form.Item name="description" label="Description">
            <Input.TextArea />
          </Form.Item>
          <Form.Item
            name="picture"
            label="Picture"
            valuePropName="file"
            rules={[{ required: true }]}
          >
            <Upload beforeUpload={() => false} maxCount={1}>
              <Button icon={<UploadOutlined />}>Upload</Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>

      {/* Update Modal */}
      <Modal
        title="Update Product"
        open={isUpdateOpen}
        onCancel={() => setIsUpdateOpen(false)}
        onOk={() => updateForm.submit()}
      >
        <Form form={updateForm} layout="vertical" onFinish={handleUpdate}>
          <Form.Item name="name" label="Name" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[{ required: true }]}
          >
            <InputNumber min={0} style={{ width: "100%" }} />
          </Form.Item>
          <Form.Item name="price" label="Price" rules={[{ required: true }]}>
            <InputNumber min={0} style={{ width: "100%" }} />
          </Form.Item>
          <Form.Item name="description" label="Description">
            <Input.TextArea />
          </Form.Item>
          <Form.Item name="picture" label="Picture" valuePropName="file">
            <Upload beforeUpload={() => false} maxCount={1}>
              <Button icon={<UploadOutlined />}>Upload</Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
}

export default Editproducts;
