import React from "react";
import axios from "axios";

const LogoutButton = () => {
  const logoutUser = () => {
    const token = localStorage.getItem("token"); // Get the token from local storage

    axios
      .post(
        "http://localhost:8000/auth/logout",
        {},
        {
          headers: {
            "Auth-Access-Token": token, // Pass the token in the request header
          },
        }
      )
      .then(function (response) {
        localStorage.removeItem("token"); // Remove the token from local storage upon successful logout
        console.log(response);

        // Redirect to login page upon successful logout
        window.location.href = "/";
      })
      .catch(function (error) {
        console.log(error, "error");
        // Handle logout error (e.g., display error message)
      });
  };

  const isLoggedIn = !!localStorage.getItem("token");

  return (
    isLoggedIn && (
      <button type="button" className="logout-btn" onClick={logoutUser}>
        Logout
      </button>
    )
  );
};

export default LogoutButton;
