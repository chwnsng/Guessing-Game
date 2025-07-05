import "./App.css";
import { useEffect } from "react";
import { Navigate, Routes, Route, useNavigate } from "react-router-dom";
import { AuthProvider, useAuth } from "./context/AuthContext";
import LoginPage from "./pages/LoginPage";
import GuessPage from "./pages/GuessPage";

const ProtectedRoute = ({ children }) => {
  const { isAuthenticated } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated) {
      navigate("/login", { replace: true });
    }
  }, [isAuthenticated, navigate]);

  // use effect will handle redirection above
  // return null here to prevent children from rendering before navigation
  if (!isAuthenticated) {
    return null;
  } else {
    return children;
  }
};

function App() {
  return (
    <AuthProvider>
      <div className="App">
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route
            path="/guess"
            element={
              <ProtectedRoute>
                <GuessPage />
              </ProtectedRoute>
            }
          />
          {/* Redirect any other pages to /guess if authenticated, else /login */}
          {/* <Route path="*" element={<Navigate to="/guess" replace />} /> */}
        </Routes>
      </div>
    </AuthProvider>
  );
}

export default App;
