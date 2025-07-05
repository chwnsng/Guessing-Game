import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext"; // import the custom hook

const LoginPage = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const { login, isAuthenticated } = useAuth();
  const navigate = useNavigate();

  // Redirect to guess page if authenticated
  useEffect(() => {
    if (isAuthenticated) {
      navigate("/guess", { replace: true });
    }
  }, [isAuthenticated, navigate]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(""); // clearing previous errors
    try {
      await login(username, password);
    } catch (err) {
      setError(err.message || "An unexpected error occured during login");
    }
  };

  return (
    <div>
      <h2>Guessing Game 🤔</h2>
      {/* <div>Login</div> */}
      <form onSubmit={handleSubmit} autoComplete="off" noValidate>
        <div className="input-field">
          <label htmlFor="username" className="input-label">
            Username
          </label>
          <input
            type="text"
            id="username"
            className="input-text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
        </div>
        <div className="input-field">
          <label htmlFor="password" className="input-label">
            Password
          </label>
          <input
            type="password"
            id="password"
            className="input-text"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Login</button>
        {error && <p className="error">{error}</p>}
      </form>
    </div>
  );
};

export default LoginPage;
