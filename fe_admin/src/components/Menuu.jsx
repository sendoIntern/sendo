import React from "react";
import {} from "@ant-design/icons";
import { Menu } from "antd";
import { useNavigate } from "react-router-dom";
const items = [
  {
    key: "Grp",
    type: "group",
    label: "Admin Page",
    children: [
      {
        key: "/dashboard",
        label: "Dashboard",
      },
      {
        key: "/products",
        label: "Products",
      },
      {
        key: "/editProducts",
        label: "EditProducts",
      },
    ],
  },
];
const Menuu = () => {
  const navigate = useNavigate();
  const onClick = (e) => {
    navigate(e.key);
  };

  return (
    <Menu
      onClick={onClick}
      style={{ width: 256 }}
      defaultSelectedKeys={["1"]}
      defaultOpenKeys={["sub1"]}
      items={items}
    />
  );
};
export default Menuu;
