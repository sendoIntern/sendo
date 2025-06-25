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
import Link from "antd/es/typography/Link";

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

  const [showImportModal, setShowImportModal] = useState(false);

  const [alert, setAlert] = useState(null); // { type: "success" | "error", message: string }

  const [statusFillter, setStatusFillter] = useState(null);

  // Hàm showAlert tiện lợi, tự ẩn sau duration ms
  const showAlert = (type, messageText, duration = 3000) => {
    setAlert({ type, message: messageText });
    setTimeout(() => setAlert(null), duration);
  };

  useEffect(() => {
    fetchProducts();
  }, [currentPage, searchTerm, minPrice, maxPrice, statusFillter]);

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
      const res = await axiosInstance.get("/item/getAllItems", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        params: {
          page: currentPage,
          limit,
          search: searchTerm || undefined,
          minPrice: minPrice || undefined,
          maxPrice: maxPrice || undefined,
          status: statusFillter,
        },
        withCredentials: true,
      });
      if (res.data) {
        setData(res.data.data);
        setTotalPages(res.data.pagination?.total_pages || 1);
      }
      console.log(res.data);
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

      showAlert("success", "Product created successfully!");
      setShowCreateModal(false);
      formCreate.resetFields();
      fetchProducts();
    } catch (error) {
      console.error("Error creating product:", error);
      showAlert("error", "Error creating product!");
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

      showAlert("success", "Product updated successfully!");
      setShowUpdateModal(false);
      fetchProducts();
    } catch (error) {
      console.error("Error updating product:", error);
      showAlert("error", "Error updating product!");
    } finally {
      setLoading(false);
    }
  };

  const handleChangeStatus = async (record) => {
    setLoading(true);
    try {
      const res = await axiosInstance.patch(
        `/item/changeStatus/${record.id}`,
        {},
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
          },
          withCredentials: true,
        }
      );
      showAlert("success", `Trạng thái sản phẩm đã được đổi thành công!`);
      console.log(res.data);
      fetchProducts();
    } catch (error) {
      console.error("Error changing product status:", error);
      showAlert("error", "❌ Lỗi khi thay đổi trạng thái sản phẩm.");
    } finally {
      setLoading(false);
    }
  };

  const handleImportExcel = async (file) => {
    setLoading(true);
    try {
      const formData = new FormData();
      formData.append("file", file);

      const res = await axiosInstance.post("/item/import", formData, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });
      console.log(res.data);
    } catch (error) {
      console.error("Error importing Excel file:", error);
      showAlert("error", "❌ Lỗi khi import file.");
    } finally {
      fetchProducts();
      setLoading(false);
    }
  };

  const columnsDefault = [
    {
      title: "Image",
      dataIndex: "picture",
      render: (src) => <img src={src} alt="product" style={{ width: 50 }} />,
    },
    { title: "Name", dataIndex: "name" },
    {
      title: "Price",
      dataIndex: "price",
      render: (value) => `${value.toLocaleString()} $`,
    },
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
      dataIndex: "is_active", // cần thiết cho filter hoạt động đúng
      filters: [
        {
          text: "Active",
          value: true,
        },
        {
          text: "Inactive",
          value: false,
        },
      ],
      filterMultiple: false,
      render: (_, record) => (
        <Button onClick={() => handleChangeStatus(record)}>
          {record.is_active ? "Active" : "Inactive"}
        </Button>
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
          <Button
            type="primary"
            icon={<FileExcelOutlined />}
            onClick={() => setShowImportModal(true)}
          >
            Import File
          </Button>
          <Button
            type="primary"
            icon={<PlusOutlined />}
            onClick={() => setShowCreateModal(true)}
          >
            New
          </Button>
        </Space>

        <Table
          columns={columnsDefault}
          dataSource={data}
          loading={loading}
          rowKey="id"
          pagination={false}
          onChange={(pagination, filters) => {
            const activeStatus = filters?.is_active?.[0]; // true / false
            setStatusFillter(activeStatus);
            console.log("after set", activeStatus);
            fetchProducts();
            console.log("sau khi fetch");
          }}
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

      {/* Import Modal */}
      <Modal
        open={showImportModal}
        title="Import Products from Excel"
        onCancel={() => setShowImportModal(false)}
        footer={null}
      >
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
        <br />
        <Link href="\ItemIport.xlsx" download>
          Tải file mẫu ở đây
        </Link>
      </Modal>
    </>
  );
};

export default Dashboard;
