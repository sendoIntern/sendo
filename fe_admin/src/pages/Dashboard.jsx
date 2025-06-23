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
  Alert,
  InputNumber,
} from "antd";
import {
  UploadOutlined,
  PlusOutlined,
  FileExcelOutlined,
} from "@ant-design/icons";
import Nav from "../components/Nav";

import { inputValidate } from "../lib/inputValidate";

const Dashboard = () => {
  const [data, setData] = useState([]);
  const [loading, setLoading] = useState(true);
  const [_, setFileImport] = useState(null);
  const [searchTerm, setSearchTerm] = useState("");
  const [minPrice, setMinPrice] = useState("");
  const [maxPrice, setMaxPrice] = useState("");
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const limit = 6;

  const [showCreateModal, setShowCreateModal] = useState(false);
  const [showUpdateModal, setShowUpdateModal] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);
  const [formCreate] = Form.useForm();
  const [formUpdate] = Form.useForm();

  const [alert, setAlert] = useState(null); // { type: "success" | "error", message: string }

  // Hàm showAlert tiện lợi, tự ẩn sau duration ms
  const showAlert = (type, messageText, duration = 3000) => {
    setAlert({ type, message: messageText });
    setTimeout(() => setAlert(null), duration);
  };

  useEffect(() => {
    fetchProducts();
  }, [currentPage, searchTerm, minPrice, maxPrice]);

  useEffect(() => {
    if (selectedItem && showUpdateModal) {
      formUpdate.setFieldsValue({
        name: selectedItem.name,
        price: selectedItem.price,
        quantity: selectedItem.quantity,
        description: selectedItem.description,
        picture: [],
      });
    }
  }, [selectedItem, showUpdateModal]);

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
      Object.entries(values).forEach(([key, value]) => {
        if (key === "picture" && value?.[0]?.originFileObj) {
          formData.append("picture", value[0].originFileObj);
        } else {
          formData.append(key, value);
        }
      });

      await axiosInstance.post("/item/createNewItem", formData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });

      showAlert("success", "🎉 Product created successfully!");
      setShowCreateModal(false);
      formCreate.resetFields();
      fetchProducts();
    } catch (error) {
      console.error("Error creating product:", error);
      showAlert("error", "❌ Error creating product!");
    } finally {
      setLoading(false);
    }
  };

  const handleUpdate = async (values) => {
    setLoading(true);
    try {
      const formData = new FormData();
      Object.entries(values).forEach(([key, value]) => {
        if (key === "picture" && value?.[0]?.originFileObj) {
          formData.append("picture", value[0].originFileObj);
        } else {
          formData.append(key, value);
        }
      });

      await axiosInstance.put(`/item/${selectedItem.id}`, formData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });

      showAlert("success", "✅ Product updated successfully!");
      setShowUpdateModal(false);
      fetchProducts();
    } catch (error) {
      console.error("Error updating product:", error);
      showAlert("error", "❌ Error updating product!");
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
      showAlert("success", "🗑️ Xoá sản phẩm thành công!");
      fetchProducts();
    } catch (error) {
      console.error("Error deleting product:", error);
      showAlert("error", "❌ Không thể xoá sản phẩm.");
    } finally {
      setLoading(false);
    }
  };

  const handleImportExcel = async (file) => {
    setLoading(true);
    try {
      const formData = new FormData();
      formData.append("file", file);

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
        showAlert("success", "📥 Import Excel thành công!");
      } else {
        showAlert("error", "❌ Import thất bại, có dữ liệu lỗi.");
      }
    } catch (error) {
      console.error("Error importing Excel file:", error);
      showAlert("error", "❌ Lỗi khi import file.");
    } finally {
      fetchProducts();
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
              setShowUpdateModal(true);
            }}
          >
            Update
          </Button>
        </Space>
      ),
    },
    {
      title: "Status",
      render: (_, record) => (
        <Space>
          <Button danger onClick={() => handleDelete(record.id)}>
            active
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <>
      <Nav />

      {/* Alert */}
      {alert && (
        <Alert
          type={alert.type}
          message={alert.message}
          showIcon
          closable
          onClose={() => setAlert(null)}
          style={{ marginBottom: 16 }}
        />
      )}

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
          <Upload
            accept=".xlsx"
            showUploadList={false}
            beforeUpload={(file) => {
              setFileImport(file);
              handleImportExcel(file);
              return false;
            }}
          >
            <Button icon={<FileExcelOutlined />}>Import File</Button>
          </Upload>
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

      {/* Create Modal */}
      <Modal
        open={showCreateModal}
        title="Create Product"
        onCancel={() => setShowCreateModal(false)}
        onOk={() => formCreate.submit()}
      >
        <Form layout="vertical" form={formCreate} onFinish={handleCreate}>
          <Form.Item
            name="name"
            label="Name"
            rules={[
              { required: true, message: "Vui lòng nhập tên" },
              {
                validator: (_, value) => {
                  if (!value || inputValidate.isValidName(value)) {
                    return Promise.resolve();
                  }
                  return Promise.reject(
                    "Tên không hợp lệ. Chỉ cho phép chữ cái và khoảng trắng, từ 2–50 ký tự."
                  );
                },
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[
              { required: true, message: "Vui lòng nhập số lượng" },
              {
                validator: (_, value) => {
                  if (!value || inputValidate.isValidQuantity(value)) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error("Số lượng không hợp lệ"));
                },
              },
            ]}
          >
            <InputNumber
              onKeyDown={inputValidate.restrictNumberInputKeys}
              formatter={inputValidate.formatNumberWithCommas}
              parser={inputValidate.parseNumberFromString}
              style={{ width: "100%" }}
            />
          </Form.Item>
          <Form.Item
            name="price"
            label="Price"
            rules={[
              { required: true, message: "Vui lòng nhập giá" },
              {
                validator: (_, value) => {
                  if (!value || inputValidate.isValidPrice(value)) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error("Giá không hợp lệ"));
                },
              },
            ]}
          >
            <InputNumber
              formatter={inputValidate.formatNumberWithCommas}
              parser={inputValidate.parseNumberFromString}
              style={{ width: "100%" }}
            />
          </Form.Item>
          <Form.Item name="description" label="Description">
            <Input.TextArea />
          </Form.Item>
          <Form.Item
            name="picture"
            label="Upload Image"
            valuePropName="fileList"
            getValueFromEvent={(e) => (Array.isArray(e) ? e : e && e.fileList)}
          >
            <Upload beforeUpload={() => false} maxCount={1}>
              <Button icon={<UploadOutlined />}>Click to Upload</Button>
            </Upload>
          </Form.Item>
        </Form>
      </Modal>

      {/* Update Modal */}
      <Modal
        open={showUpdateModal}
        title="Update Product"
        onCancel={() => setShowUpdateModal(false)}
        onOk={() => formUpdate.submit()}
      >
        <Form layout="vertical" form={formUpdate} onFinish={handleUpdate}>
          <Form.Item
            name="name"
            label="Name"
            rules={[
              { required: true, message: "Vui lòng nhập tên" },
              {
                validator: (_, value) => {
                  if (!value || inputValidate.isValidName(value)) {
                    return Promise.resolve();
                  }
                  return Promise.reject(
                    new Error(
                      "Tên không hợp lệ (chỉ chứa chữ và dài 2-20 ký tự)"
                    )
                  );
                },
              },
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="quantity"
            label="Quantity"
            rules={[
              { required: true, message: "Vui lòng nhập số lượng" },
              {
                validator: (_, value) => {
                  if (!value || inputValidate.isValidQuantity(value)) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error("Số lượng không hợp lệ"));
                },
              },
            ]}
          >
            <InputNumber
              onKeyDown={inputValidate.restrictNumberInputKeys}
              formatter={inputValidate.formatNumberWithCommas}
              parser={inputValidate.parseNumberFromString}
              style={{ width: "100%" }}
            />
          </Form.Item>
          <Form.Item
            name="price"
            label="Price"
            rules={[
              { required: true, message: "Vui lòng nhập giá" },
              {
                validator: (_, value) => {
                  if (!value || inputValidate.isValidPrice(value)) {
                    return Promise.resolve();
                  }
                  return Promise.reject(new Error("Giá không hợp lệ"));
                },
              },
            ]}
          >
            <InputNumber
              onKeyDown={inputValidate.restrictNumberInputKeys}
              formatter={inputValidate.formatNumberWithCommas}
              parser={inputValidate.parseNumberFromString}
              style={{ width: "100%" }}
            />
          </Form.Item>
          <Form.Item name="description" label="Description">
            <Input.TextArea />
          </Form.Item>
          <Form.Item
            name="picture"
            label="Upload New Image"
            valuePropName="fileList"
            getValueFromEvent={(e) => (Array.isArray(e) ? e : e && e.fileList)}
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
