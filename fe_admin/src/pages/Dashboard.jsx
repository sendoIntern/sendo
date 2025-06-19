import React, { useEffect, useState } from "react";
import { axiosInstance } from "../lib/axios";
import {
  Table,
  Button,
  Input,
  Modal,
  Form,
  Upload,
  Pagination,
  Typography,
  Space,
  message,
} from "antd";
import { UploadOutlined, PlusOutlined } from "@ant-design/icons";
import Nav from "../components/Nav";

const Dashboard = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [fileImport, setFileImport] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [minPrice, setMinPrice] = useState("");
  const [maxPrice, setMaxPrice] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const limit = 6;

  const [showCreateModal, setShowCreateModal] = useState(false);
  const [showUpdateModal, setShowUpdateModal] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const [form] = Form.useForm();

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
        setData(res.data.data);
        setTotalPages(res.data.pagination?.total_pages || 1);
      }
    } catch (error) {
      console.error("Error fetching products:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async (values) => {
    setLoading(true);
    try {
      const formData = new FormData();
      Object.keys(values).forEach((key) => {
        formData.append(key, values[key]);
      });

      await axiosInstance.post("/item/createNewItem", formData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });

      message.success("Product created successfully!");
      setShowCreateModal(false);
      form.resetFields();
      fetchProducts();
    } catch (error) {
      console.error("Error creating product:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleUpdate = async (values) => {
    setLoading(true);
    try {
      const formData = new FormData();
      Object.keys(values).forEach((key) => {
        formData.append(key, values[key]);
      });

      await axiosInstance.put(`/item/${selectedItem.id}`, formData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });
      message.success("Updated successfully!");
      setShowUpdateModal(false);
      fetchProducts();
    } catch (error) {
      console.error("Error updating product:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleDelete = async (id) => {
    setLoading(true);
    try {
      await axiosInstance.delete(`/item/${id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });
      message.success("Deleted successfully!");
      fetchProducts();
    } catch (error) {
      console.error("Error deleting product:", error);
    } finally {
      setLoading(false);
    }
  };

  const columns = [
    {
      title: "Image",
      dataIndex: "picture",
      render: (src) => <img src={src} alt="product" style={{ width: 50 }} />,
    },
    { title: "Name", dataIndex: "name" },
    { title: "Price", dataIndex: "price" },
    { title: "Quantity", dataIndex: "quantity" },
    { title: "Description", dataIndex: "description" },
    { title: "View", dataIndex: "view" },
    {
      title: "Actions",
      render: (_, record) => (
        <Space>
          <Button
            onClick={() => {
              setSelectedItem(record);
              form.setFieldsValue(record);
              setShowUpdateModal(true);
            }}
          >
            Update
          </Button>
          <Button danger onClick={() => handleDelete(record.id)}>
            Delete
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <>
      <Nav />
      <div style={{ padding: 24 }}>
        <Typography.Title level={2}>Product List</Typography.Title>

        <Space style={{ marginBottom: 16, flexWrap: "wrap" }}>
          <Input
            placeholder="Search by name"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
          <Input
            placeholder="Min Price"
            type="number"
            value={minPrice}
            onChange={(e) => setMinPrice(e.target.value)}
          />
          <Input
            placeholder="Max Price"
            type="number"
            value={maxPrice}
            onChange={(e) => setMaxPrice(e.target.value)}
          />
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={() => setShowCreateModal(true)}
          >
            New
          </Button>
        </Space>

        <Table
          columns={columns}
          dataSource={data}
          loading={loading}
          rowKey="id"
          pagination={false}
        />

        <Pagination
          current={currentPage}
          total={totalPages * limit}
          pageSize={limit}
          onChange={setCurrentPage}
          style={{ marginTop: 16, textAlign: "center" }}
        />
      </div>

      <Modal
        open={showCreateModal}
        title="Create Product"
        onCancel={() => setShowCreateModal(false)}
        onOk={() => form.submit()}
      >
        <Form layout="vertical" form={form} onFinish={handleCreate}>
          <Form.Item name="name" label="Name" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[{ required: true }]}
          >
            <Input type="number" />
          </Form.Item>
          <Form.Item name="price" label="Price" rules={[{ required: true }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="description" label="Description">
            <Input.TextArea />
          </Form.Item>
          <Form.Item name="picture" label="Upload Image" valuePropName="file">
            <Upload beforeUpload={() => false} maxCount={1}>
              <Button icon={<UploadOutlined />}>Click to Upload</Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>

      <Modal
        open={showUpdateModal}
        title="Update Product"
        onCancel={() => setShowUpdateModal(false)}
        onOk={() => form.submit()}
      >
        <Form layout="vertical" form={form} onFinish={handleUpdate}>
          <Form.Item name="name" label="Name" rules={[{ required: true }]}>
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[{ required: true }]}
          >
            <Input type="number" />
          </Form.Item>
          <Form.Item name="price" label="Price" rules={[{ required: true }]}>
            <Input type="number" />
          </Form.Item>
          <Form.Item name="description" label="Description">
            <Input.TextArea />
          </Form.Item>
          <Form.Item
            name="picture"
            label="Upload New Image"
            valuePropName="file"
          >
            <Upload beforeUpload={() => false} maxCount={1}>
              <Button icon={<UploadOutlined />}>Click to Upload</Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default Dashboard;
