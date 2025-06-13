export const auth = {
  // phân rã token lấy role từ localStorage
  getRole: () => {
    const token = localStorage.getItem("token");
    if (!token) return null;

    try {
      const payload = JSON.parse(atob(token.split(".")[1]));
      return payload.role || null;
    } catch (error) {
      console.error("Error decoding token:", error);
      return null;
    }
  },

  getUser: () => {
    const token = localStorage.getItem("token");
    if (!token) return null;

    try {
      const payload = JSON.parse(atob(token.split(".")[1]));
      return payload.user || null;
    } catch (error) {
      console.error("Error decoding token:", error);
      return null;
    }
  },

  logout: () => {
    localStorage.removeItem("token");
  },
};
