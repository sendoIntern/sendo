import { useGoogleLogin } from "@react-oauth/google";
import { useState } from "react";
import { axiosInstance } from "../lib/axios";
import axios from "axios";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import GoogleIcon from "@mui/icons-material/Google";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [, setUserInfor] = useState(null);
  const navigate = useNavigate();
  const bull = (
    <Box
      component="span"
      sx={{ display: "inline-block", mx: "2px", transform: "scale(0.8)" }}
    ></Box>
  );

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

        localStorage.setItem("accessToken", responce.data.accessToken);
        navigate("/product");
      } catch (err) {
        console.error("Lỗi đăng nhập", err);
      }
    },
    onError: (error) => console.error("Login lỗi:", error),
  });

  return (
    <Box
      sx={{
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#f5f5f5",
        textAlign: "center",
      }}
    >
      <Card sx={{ minWidth: 275, p: 2 }}>
        <CardContent>
          <Typography
            gutterBottom
            sx={{ color: "text.secondary", fontSize: 14 }}
          >
            Welcome
          </Typography>
          <Typography variant="h5" component="div">
            Log in
          </Typography>
        </CardContent>
        <Box
          component="form"
          sx={{ "& > :not(style)": { m: 1, width: "25ch" } }}
          noValidate
          autoComplete="off"
        >
          <TextField id="login-username" label="Username" variant="outlined" />
          <br />
          <TextField
            id="login-password"
            label="Password"
            type="password"
            variant="outlined"
          />
        </Box>
        <CardActions sx={{ justifyContent: "center" }}>
          <Button size="small" variant="contained" color="primary">
            Đăng nhập
          </Button>
        </CardActions>
        <CardActions sx={{ justifyContent: "center" }}>
          <GoogleIcon />
          <Button size="small" onClick={() => login()}>
            Đăng nhập bằng Google
          </Button>
        </CardActions>
      </Card>
    </Box>
  );
};

export default Login;
