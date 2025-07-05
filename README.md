# Guessing Game
## Setup
**Prerequisites:**
This project was built using `go1.24.4`, `node v16.18.0`, and `npm 9.8.1`.

**1. clone the repository**
```
git clone https://github.com/chwnsng/Guessing-Game.git
cd Guessing-Game
```

*2. Set up the backend*
```
cd backend
go mod tidy
go run main.go
```
The backend should start on http://localhost:8080

*3. Set up the frontend*
```
cd ../frontend
npm install
npm start
```
The frontend should start on http://localhost:3000. Please use these sample login credentials
<br>
**Username**: `test` <br>
**Password**: `1234`

<br>

## Backend
**Overall flow**
1. The application starts in main.go, which configures a server with CORS enabled and waits for requests from the frontend.

2. Requests to **/login** is passed to the `LoginHandler` function defined in auth.go. If successful, LoginHandler calls `CreateToken` to create a jwt token, then calls `RespondJson` to send it back. Else, it calls `RespondError` so send back the error.

3. Requests to **/guess** first goes to `AuthMiddleware` (defined in authmdw.go). This middleware calls `VerifyToken` in jwt.go to validate the Authorization header. If the token is valid, AuthMiddleware passes the request to `GuessHandler`.

4. GuessHandler then gets the current secret from `GetSecretNumber` and checks it with the value from the request. If the guess is correct, it calls `GenerateSecretNumber` to regenerate a new number. Either RespondJson or RespondError is called to send back each response. 

<br>

## Frontend
**Overall flow**
1. App.js provides the `AuthContext` globally, and sets up routing.

2. Users visiting route **/** or **/login** will see the `LoginPage` if not yet authenticated. Else **/** or **/guess** will render the `GuessPage`. The **/guess** route is protected by `ProtectedRoute` which checks AuthContext for authentication status.

4. On LoginPage, credentials are submitted via a form which calls `handleSubmit` --> `login` --> `api.login`. the login function updates the auth status, token, and login/logout status in the AuthContext. The api.login function makes the actual api calls to the backend. The token is stored in localstorage to persist between reloads.

5. If login is successful, AuthContext is updated, and LoginPage redirects to GuessPage.

6. GuessPage checks AuthContext to get the token, then on user submit it calls `handleSubmit` --> `api.guessNumber` to send guesses to the backend.

7. GuessPage then displays results based on the backend's response. It also allows users to play again by calling `handlePlayAgain` to reset the state, or log out by calling `handleLogout` --> `logout` to remove the token from the AuthContext. This will redirect the user back to LoginPage.