import * as React from "react";
import {
  Box,
  Avatar,
  Menu,
  MenuItem,
  ListItemIcon,
  Divider,
  IconButton,
  Typography,
  Tooltip,
  TextField,
} from "@mui/material";
import Logout from "@mui/icons-material/Logout";
import { useLocation, useNavigate } from "react-router-dom";

const Nav = () => {
  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const navigate = useNavigate();
  const location = useLocation(); // để lấy route hiện tại
  const [searchTerm, setSearchTerm] = React.useState("");
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };
  const handleLogout = () => {
    localStorage.removeItem("accessToken");
    navigate("/login");
  };

  // Helper kiểm tra route hiện tại
  const isActive = (path) => location.pathname === path;

  return (
    <>
      <Box
        sx={{
          position: "fixed",
          top: 0,
          left: 0,
          right: 0,
          zIndex: 1100,
          bgcolor: "white",
          boxShadow: 1,
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          padding: "10px 20px",
        }}
      >
        {/* Menu bên trái */}
        <Box sx={{ display: "flex", gap: 4 }}>
          <Typography
            onClick={() => navigate("/dashboard")}
            sx={{
              cursor: "pointer",
              px: 2,
              py: 1,
              borderRadius: 1,
              fontWeight: isActive("/") ? "bold" : "normal",
              color: isActive("/") ? "#1976d2" : "inherit",
              backgroundColor: isActive("/") ? "#e3f2fd" : "transparent",
              "&:hover": {
                backgroundColor: "#e3f2fd",
                color: "#1976d2",
              },
            }}
          >
            Dashboard
          </Typography>

          <Typography
            onClick={() => navigate("/product")}
            sx={{
              cursor: "pointer",
              px: 2,
              py: 1,
              borderRadius: 1,
              fontWeight: isActive("/product") ? "bold" : "normal",
              color: isActive("/product") ? "#1976d2" : "inherit",
              backgroundColor: isActive("/product") ? "#e3f2fd" : "transparent",
              "&:hover": {
                backgroundColor: "#e3f2fd",
                color: "#1976d2",
              },
            }}
          >
            Product
          </Typography>
        </Box>

        {/* Search + Avatar */}
        <Box sx={{ display: "flex", alignItems: "center", gap: 2 }}>
          <Tooltip title="Account settings">
            <IconButton
              onClick={handleClick}
              size="small"
              sx={{
                border: "2px solid transparent",
                "&:hover": {
                  bgcolor: "#e3f2fd",
                  border: "2px solid #1976d2",
                },
              }}
            >
              <Avatar sx={{ width: 32, height: 32 }}>M</Avatar>
            </IconButton>
          </Tooltip>
        </Box>
      </Box>

      <Menu
        anchorEl={anchorEl}
        id="account-menu"
        open={open}
        onClose={handleClose}
        onClick={handleClose}
        slotProps={{
          paper: {
            elevation: 0,
            sx: {
              overflow: "visible",
              filter: "drop-shadow(0px 2px 8px rgba(0,0,0,0.32))",
              mt: 1.5,
              "& .MuiAvatar-root": {
                width: 32,
                height: 32,
                ml: -0.5,
                mr: 1,
              },
              "&::before": {
                content: '""',
                display: "block",
                position: "absolute",
                top: 0,
                right: 14,
                width: 10,
                height: 10,
                bgcolor: "background.paper",
                transform: "translateY(-50%) rotate(45deg)",
                zIndex: 0,
              },
            },
          },
        }}
        transformOrigin={{ horizontal: "right", vertical: "top" }}
        anchorOrigin={{ horizontal: "right", vertical: "bottom" }}
      >
        <Divider />
        <MenuItem onClick={handleLogout}>
          <ListItemIcon>
            <Logout fontSize="small" />
          </ListItemIcon>
          Logout
        </MenuItem>
      </Menu>
    </>
  );
};

export default Nav;
