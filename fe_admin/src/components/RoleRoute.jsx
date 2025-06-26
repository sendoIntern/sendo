// src/components/RoleRoute.jsx
import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { auth } from "../lib/auth";

const RoleRoute = ({ allowRoles }) => {
  const role = auth.getRole();

  // Nếu chưa có role (chưa đăng nhập), hoặc không thuộc nhóm được phép
  if (!role || !allowRoles.includes(role)) {
    return <Navigate to="/login" replace />;
  }

  return <Outlet />;
};

export default RoleRoute;
