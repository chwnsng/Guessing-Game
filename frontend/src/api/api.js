const API_BASE_URL = "http://localhost:8080"; // Go backend

// Login api call
export const login = async (username, password) => {
  try {
    // login request
    const response = await fetch(`${API_BASE_URL}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Login failed");
    }

    // return the jwt token once response is received
    const data = await response.json();
    return data.token;
  } catch (e) {
    console.error("Login API error:", e);
    throw e;
  }
};

// Number guessing API call
export const guessNumber = async (token, number) => {
  try {
    const response = await fetch(`${API_BASE_URL}/guess`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`, // including the jwt token in the request
      },
      body: JSON.stringify({ number }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Guess failed");
    }

    const data = await response.json();
    return data; // "message", "correct"
  } catch (error) {
    console.error("Guess API error:", error);
    throw error;
  }
};
