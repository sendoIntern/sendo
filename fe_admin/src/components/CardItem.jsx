import React, { useEffect, useState } from "react";
import { Button, Card, Modal, Tooltip } from "antd";
import { axiosInstance } from "../lib/axios";
import Meta from "antd/es/card/Meta";
import { useNavigate } from "react-router-dom";

function CardItem() {
  const [data, setData] = useState([]);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const navigate = useNavigate();

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

  const showModal = () => {
    setIsModalOpen(true);
  };
  const handleOk = () => {
    setIsModalOpen(false);
  };
  const handleCancel = () => {
    setIsModalOpen(false);
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
          cover={<img alt="example" src={item.picture} onClick={showModal} />}
        >
          <Meta
            style={{ textAlign: "center" }}
            title={item.name}
            onClick={showModal}
          />
          <Meta
            style={{ height: "100px" }}
            description={`${item.description}`}
            onClick={showModal}
          />
          <Meta description={`Price: ${item.price}`} onClick={showModal} />
          <Button style={{ width: "100%", marginTop: "10px", color: "blue" }}>
            Buy
          </Button>
        </Card>
      ))}
      {/* <Modal
        title="Basic Modal"
        closable={{ "aria-label": "Custom Close Button" }}
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
      >
        <Card
          key={item.id}
          style={{ width: 300 }}
          cover={<img alt="example" src={item.picture} onClick={showModal} />}
        >
          <Meta
            style={{ textAlign: "center" }}
            title={item.name}
            onClick={showModal}
          />
          <Meta
            style={{ height: "100px" }}
            description={`${item.description}`}
            onClick={showModal}
          />
          <Meta description={`Price: ${item.price}`} onClick={showModal} />
          <Button style={{ width: "100%", marginTop: "10px", color: "blue" }}>
            Buy
          </Button>
        </Card>
      </Modal> */}
    </div>
  );
}

export default CardItem;
