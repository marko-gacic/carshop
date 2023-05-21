import React, { useState } from "react";
import axios from "axios";
import { useNavigate, Link } from "react-router-dom";

export default function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();

  const handleUsernameChange = (e) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const logInUser = () => {
    if (username.length === 0) {
      alert("Username has left blank!");
    } else if (password.length === 0) {
      alert("Password has left blank!");
    } else {
      axios
        .post("http://localhost:8000/auth/login", {
          data: {
            username: username,
            password: password,
          },
        })
        .then(function (response) {
          const token = response.headers["auth-access-token"];
          localStorage.setItem("token", token);
          navigate("/carlist");
        })
        .catch(function (error) {
          if (error.response.status === 401) {
            alert("Invalid credentials");
          }
        });
    }
  };

  return (
    <div className="login-body">
      <div className="login-container">
        <div className="box">
          <div className="cover"></div>
          <div className="shadow"></div>
          <div className="content">
            <div className="form">
              <h3 className="logo">
                <i className="fa-solid fa-key"></i>
              </h3>
              <h2>Sign In</h2>
              <div className="inputBox">
                <input type="text" required value={username} onChange={handleUsernameChange} />
                <i className="fa-solid fa-user"></i>
                <span>Username</span>
              </div>
              <div className="inputBox">
                <input type="password" required value={password} onChange={handlePasswordChange} />
                <i className="fa-solid fa-lock"></i>
                <span>Password</span>
              </div>
              <div className="links">
                <Link to="#">
                  <i className="fa-solid fa-question"></i> Forgot Password
                </Link>
                <Link to="#">
                  <i className="fa-solid fa-user-plus"></i> Sign Up
                </Link>
              </div>
              <div className="inputBox">
                <input type="submit" value="Login" onClick={logInUser} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
