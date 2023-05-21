import { useState, useEffect } from "react";

const useAuth = () => {
  const [token, setToken] = useState(null);
  const [userRole, setUserRole] = useState(null);

  useEffect(() => {
    // Check if the user is already authenticated (e.g., token stored in localStorage)
    const storedToken = localStorage.getItem("token");
    const storedUserRole = localStorage.getItem("userRole");

    if (storedToken && storedUserRole) {
      setToken(storedToken);
      setUserRole(storedUserRole);
    }
  }, []);

  const login = (token, role) => {
    // Store the token and user role in localStorage
    localStorage.setItem("token", token);
    localStorage.setItem("userRole", role);

    // Update the state
    setToken(token);
    setUserRole(role);
  };

  const logout = () => {
    // Remove the token and user role from localStorage
    localStorage.removeItem("token");
    localStorage.removeItem("userRole");

    // Update the state
    setToken(null);
    setUserRole(null);
  };

  return { token, userRole, login, logout };
};

export default useAuth;
