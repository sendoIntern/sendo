import React, { useState } from "react";
import { Layout, Menu, Avatar, Dropdown, Input, Typography } from "antd";
import {
  LogoutOutlined,
  DashboardOutlined,
  AppstoreOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { useLocation, useNavigate } from "react-router-dom";

const { Header } = Layout;
const { Text } = Typography;

const Nav = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const location = useLocation();
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("accessToken");
    navigate("/login");
  };

  const isActive = (path) => location.pathname === path;

  const menuItems = [
    {
      key: "dashboard",
      icon: <DashboardOutlined />,
      label: "Dashboard",
      path: "/dashboard",
    },
    {
      key: "product",
      icon: <AppstoreOutlined />,
      label: "Product",
      path: "/product",
    },
  ];

  const dropdownMenu = (
    <Menu>
      <Menu.Item
        key="logout"
        icon={<LogoutOutlined />}
        onClick={handleLogout}
        style={{ fontWeight: "bold" }}
      >
        Logout
      </Menu.Item>
    </Menu>
  );

  return (
    <Header
      style={{
        background: "#fff",
        padding: "0 24px",
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        boxShadow: "0 2px 4px rgba(0,0,0,0.1)",
        position: "sticky",
        top: 0,
        zIndex: 1000,
      }}
    >
      {/* Menu trái */}
      <Menu
        mode="horizontal"
        selectedKeys={[location.pathname.slice(1)]}
        onClick={({ key }) => {
          const item = menuItems.find((i) => i.key === key);
          if (item) navigate(item.path);
        }}
        style={{ flex: 1, fontWeight: "bold" }}
      >
        {menuItems.map((item) => (
          <Menu.Item key={item.key} icon={item.icon}>
            {item.label}
          </Menu.Item>
        ))}
      </Menu>

      {/* Search + Avatar phải */}
      <div style={{ display: "flex", alignItems: "center", gap: 16 }}>
        <Dropdown overlay={dropdownMenu} placement="bottomRight" arrow>
          <Avatar
            style={{
              backgroundColor: "#1890ff",
              cursor: "pointer",
            }}
            icon={<UserOutlined />}
          />
        </Dropdown>
      </div>
    </Header>
  );
};

export default Nav;
