import React, { useEffect, useState } from "react";
import { Button, Card, Modal } from "antd";
import { axiosInstance } from "../lib/axios";
import Meta from "antd/es/card/Meta";

function CardItem() {
  const [data, setData] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedItem, setSelectedItem] = useState(null);

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

  const fetchDataByid = async (id) => {
    try {
      const res = await axiosInstance.patch(`/item/getItemById/${id}`, {
        withCredentials: true,
      });
      selectedItem(res.data);
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
    setSelectedItem(null); // clear sau khi đóng modal
  };

  const handleCancel = () => {
    setIsModalOpen(false);
    setSelectedItem(null);
  };

  return (
    <div
      style={{
        display: "flex",
        flexWrap: "wrap",
        justifyContent: "space-around",
        gap: "20px",
      }}
    >
      {data.map((item) => (
        <Card
          key={item.id}
          style={{ width: 300 }}
          cover={
            <img
              alt="example"
              src={item.picture}
              onClick={() => showModal(item)}
            />
          }
        >
          <Meta
            style={{ textAlign: "center" }}
            title={item.name}
            onClick={() => showModal(item)}
          />
          <Meta
            style={{ height: "100px" }}
            description={item.description}
            onClick={() => showModal(item)}
          />
          <Meta
            description={`Price: ${item.price}`}
            onClick={() => showModal(item)}
          />
          <Button style={{ width: "100%", marginTop: "10px", color: "blue" }}>
            Buy
          </Button>
        </Card>
      ))}

      <Modal
        title="Chi tiết sản phẩm"
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        {selectedItem && (
          <Card
            style={{ width: "100%" }}
            cover={<img alt="example" src={selectedItem.picture} />}
          >
            <Meta style={{ textAlign: "center" }} title={selectedItem.name} />
            <Meta
              style={{ height: "100px" }}
              description={selectedItem.description}
            />
            <Meta description={`Price: ${selectedItem.price}`} />
            <Button style={{ width: "100%", marginTop: "10px", color: "blue" }}>
              Buy
            </Button>
          </Card>
        )}
      </Modal>
    </div>
  );
}

export default CardItem;
