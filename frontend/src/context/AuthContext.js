import { useState, useContext, createContext } from "react";
import * as api from "../api/api";
const AuthContext = createContext(null);

export const AuthProvider = ({ children }) => {
  const [token, setToken] = useState(localStorage.getItem("authToken")); // load jwt token from local storage

  const login = async (username, password) => {
    try {
      const newToken = await api.login(username, password);
      setToken(newToken);
      localStorage.setItem("authToken", newToken);
      return true;
    } catch (error) {
      console.error("Login failed in context:", error.message);
      setToken(null);
      localStorage.removeItem("authToken");
      throw error;
    }
  };

  const logout = () => {
    setToken(null);
    localStorage.removeItem("authToken");
  };

  const contextValue = {
    token,
    isAuthenticated: !!token,
    login,
    logout,
  };

  return (
    <AuthContext.Provider value={contextValue}>{children}</AuthContext.Provider>
  );
};

// using custom hook to simplify usage of auth context
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used in an AuthProvider");
  }
  return context;
};
