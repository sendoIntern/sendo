import React, { useEffect, useState } from "react";
import {
  Row,
  Col,
  Card,
  Input,
  InputNumber,
  Button,
  Pagination,
  Modal,
  Typography,
} from "antd";
import { axiosInstance } from "../lib/axios";

const { Title, Paragraph } = Typography;

function CardItem() {
  const [data, setData] = useState([]);
  const [selectedItem, setSelectedItem] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const [searchTerm, setSearchTerm] = useState("");
  const [minPrice, setMinPrice] = useState(null);
  const [maxPrice, setMaxPrice] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const limit = 6;

  useEffect(() => {
    fetchProducts(currentPage);
  }, [currentPage]);

  const fetchProducts = async (page = 1) => {
    try {
      const res = await axiosInstance.get("/item/getAllItems", {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
        params: {
          page,
          limit,
          search: searchTerm || undefined,
          minPrice: minPrice !== null ? minPrice : undefined,
          maxPrice: maxPrice !== null ? maxPrice : undefined,
        },
      });
      setData(res.data.data || []);
      setTotalPages(res.data.pagination?.total_pages || 1);
    } catch (error) {
      console.error("Error fetching products:", error);
    }
  };

  const handleSearch = () => {
    setCurrentPage(1);
    fetchProducts(1);
  };

  const handleFilter = () => {
    setCurrentPage(1);
    fetchProducts(1);
  };

  const showModal = async (item) => {
    try {
      const res = await axiosInstance.get(`/item/getItemById/${item.id}`, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("accessToken")}`,
        },
        withCredentials: true,
      });
      setSelectedItem(res.data.data);
      setIsModalOpen(true);
    } catch (error) {
      console.error("Error fetching item details:", error);
    }
  };

  const handleCancel = () => {
    setIsModalOpen(false);
    setSelectedItem(null);
  };

  return (
    <div style={{ padding: "24px" }}>
      <Row gutter={[24, 24]}>
        {/* Left Filter */}
        <Col xs={24} md={6}>
          <Input
            placeholder="Search by name"
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            onPressEnter={handleSearch}
            style={{ marginBottom: 16 }}
          />
          <Title level={5}>Filter by Price</Title>
          <InputNumber
            placeholder="Min Price"
            value={minPrice}
            onChange={(value) => setMinPrice(value)}
            style={{ width: "100%", marginBottom: 12 }}
          />
          <InputNumber
            placeholder="Max Price"
            value={maxPrice}
            onChange={(value) => setMaxPrice(value)}
            style={{ width: "100%", marginBottom: 16 }}
          />
          <Button type="primary" block onClick={handleFilter}>
            Apply Filter
          </Button>
        </Col>

        {/* Product List */}
        <Col xs={24} md={18}>
          <Row gutter={[16, 16]}>
            {data?.map((item) => (
              <Col xs={24} sm={12} md={8} key={item.id}>
                <Card
                  hoverable
                  cover={
                    <img
                      alt={item.name}
                      src={item.picture}
                      style={{
                        height: 180,
                        objectFit: "cover",
                        cursor: "pointer",
                      }}
                      onClick={() => showModal(item)}
                    />
                  }
                >
                  <Card.Meta
                    title={item.name}
                    description={
                      <Paragraph
                        ellipsis={{ rows: 3 }}
                        onClick={() => showModal(item)}
                        style={{ cursor: "pointer" }}
                      >
                        {item.description}
                      </Paragraph>
                    }
                  />
                  <Title level={5} style={{ marginTop: 12 }}>
                    Price: {item.price}
                  </Title>
                  <Button type="link" block>
                    Buy
                  </Button>
                </Card>
              </Col>
            ))}
          </Row>
          <div style={{ textAlign: "center", marginTop: 24 }}>
            <Pagination
              current={currentPage}
              total={totalPages * limit}
              pageSize={limit}
              onChange={(page) => {
                setCurrentPage(page);
                window.scrollTo({ top: 0, behavior: "smooth" });
              }}
              showSizeChanger={false}
            />
          </div>
        </Col>
      </Row>

      {/* Modal */}
      <Modal
        open={isModalOpen}
        title={selectedItem?.name}
        onCancel={handleCancel}
        footer={null}
      >
        {selectedItem && (
          <>
            <img
              src={selectedItem.picture}
              alt={selectedItem.name}
              style={{
                width: "100%",
                borderRadius: 8,
                objectFit: "cover",
                marginBottom: 16,
              }}
            />
            <Paragraph>{selectedItem.description}</Paragraph>
            <Title level={5}>Price: {selectedItem.price}</Title>
          </>
        )}
      </Modal>
    </div>
  );
}

export default CardItem;
