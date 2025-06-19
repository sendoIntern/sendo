import { useGoogleLogin } from "@react-oauth/google";
import { useState } from "react";
import { axiosInstance } from "../lib/axios";
import axios from "axios";
import { Button, Card, Input, Typography, Form, message } from "antd";
import { GoogleOutlined } from "@ant-design/icons";
import { useNavigate } from "react-router-dom";

const { Title, Text } = Typography;

const Login = () => {
  const [, setUserInfor] = useState(null);
  const navigate = useNavigate();

  const login = useGoogleLogin({
    onSuccess: async (tokenResponse) => {
      try {
        const res = await axios.get(
          "https://www.googleapis.com/oauth2/v3/userinfo",
          {
            headers: {
              Authorization: `Bearer ${tokenResponse.access_token}`,
            },
          }
        );
        setUserInfor(res.data);

        const responce = await axiosInstance.post("/auth/login", {
          name: res.data.name,
          email: res.data.email,
          picture: res.data.picture,
        });

        localStorage.setItem("accessToken", responce.data.data.access_token);
        message.success("Đăng nhập thành công!");
        navigate("/product");
      } catch (err) {
        console.error("Lỗi đăng nhập", err);
        message.error("Đăng nhập thất bại!");
      }
    },
    onError: (error) => {
      console.error("Login lỗi:", error);
      message.error("Google login failed");
    },
  });

  return (
    <div
      style={{
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#f5f5f5",
      }}
    >
      <Card style={{ width: 350 }}>
        <div style={{ textAlign: "center", marginBottom: 24 }}>
          <Text type="secondary">Welcome</Text>
          <Title level={3}>Log in</Title>
        </div>

        <Form layout="vertical">
          <Form.Item label="Username" name="username">
            <Input placeholder="Enter your username" />
          </Form.Item>

          <Form.Item label="Password" name="password">
            <Input.Password placeholder="Enter your password" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" block>
              Đăng nhập
            </Button>
          </Form.Item>
        </Form>

        <Button
          icon={<GoogleOutlined />}
          block
          onClick={() => login()}
          style={{ marginTop: 8 }}
        >
          Đăng nhập bằng Google
        </Button>
      </Card>
    </div>
  );
};

export default Login;
