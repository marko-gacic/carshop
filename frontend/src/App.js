import React from "react";
import "./App.css";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";
import Login from "../src/components/Login";

import Navbar from "./components/Navbar";
import CarCard from "./components/CarCard";

function App() {
  return (
    <div className="vh-100 gradient-custom">
      <Navbar />
      <div className="container">
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Login />} />
            <Route path="/carlist" element={<CarCard />} />
          </Routes>
        </BrowserRouter>
      </div>
    </div>
  );
}

export default App;
