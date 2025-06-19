// src/components/RoleRoute.jsx
import React from "react";
import { Navigate, Outlet } from "react-router-dom";
import { auth } from "../lib/auth";

const RoleRoute = ({ allowRoles }) => {
  const role = auth.getRole();
  console.log(role);

  if (!allowRoles.includes(role)) {
    return <Navigate to="/" replace />;
  }
  return <Outlet />; // Cho phép truy cập vào các <Route> con
};

export default RoleRoute;
