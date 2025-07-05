import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../context/AuthContext"; // import the custom hook
import * as api from "../api/api"; // import api call functions

const GuessPage = () => {
  const { token, isAuthenticated, logout } = useAuth();
  const navigate = useNavigate();

  const [guess, setGuess] = useState("");
  const [message, setMessage] = useState("Guess a number between 1 to 3");
  const [error, setError] = useState("");
  const [isCorrect, setIsCorrect] = useState(false);
  const [playAgain, setPlayAgain] = useState(false);

  useEffect(() => {
    if (!isAuthenticated) {
      navigate("/login"); // redirect to login page if not authenticated
    }
  }, [isAuthenticated, navigate]);

  // // Redirect to guess page if authenticated
  // if (isAuthenticated) {
  //   navigate("/guess");
  //   return null;
  // }

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(""); // clearing previous errors
    setMessage("");

    const guessNum = parseInt(guess, 10);

    // validate input
    if (isNaN(guessNum)) {
      setError("Please enter a valid number!");
      return;
    }

    if (guessNum < 1 || guessNum > 3) {
      setError(`Please enter a number between 1 and 3`);
      return;
    }

    try {
      const response = await api.guessNumber(token, guessNum);
      setMessage(response.message);
      setIsCorrect(response.correct);
      if (response.correct) {
        setPlayAgain(true); // show 'Play Again' button
      }
      setGuess(""); // clear input
    } catch (err) {
      setError(err.message || "An error occured during guessing");
    }
  };

  const handlePlayAgain = () => {
    // reset all states
    setMessage("Guess a number between 1 to 3");
    setGuess("");
    setError("");
    setIsCorrect(false);
    setPlayAgain(false);
  };

  const handleLogout = () => {
    logout();
    // navigate("/login");
  };

  if (!isAuthenticated) {
    return null;
  }

  return (
    <div>
      <h2>Make a Guess</h2>
      {message && <p>{message}</p>}
      {error && <p class="error">{error}</p>}
      {isCorrect ? (
        <div>
          <p class="correct">🎉</p>
          {playAgain && <button onClick={handlePlayAgain}>Play Again</button>}
        </div>
      ) : (
        <form onSubmit={handleSubmit} noValidate>
          <div>
            {/* <label htmlFor="guess">Your Guess:</label> */}
            <input
              type="number"
              id="guess"
              class="guess-input"
              value={guess}
              onChange={(e) => setGuess(e.target.value)}
              required
              min="1"
              max="3"
              autoFocus
            />
          </div>
          <button type="submit" class="submit-button">
            Submit
          </button>
        </form>
      )}
      <button onClick={handleLogout} class="logout-button">
        Logout
      </button>
    </div>
  );
};

export default GuessPage;
