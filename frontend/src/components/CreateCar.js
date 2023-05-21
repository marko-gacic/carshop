import React, { useState } from "react";
import axios from "axios";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const initialState = {
  id: "",
  manufacturer: "",
  model: "",
  picture: "",
  transmission: "",
  fuel: "",
  type: "",
  price: 0,
};

const CreateCar = ({ onCreate }) => {
  const [showForm, setShowForm] = useState(false);
  const [formData, setFormData] = useState(initialState);

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    const parsedValue = name === "price" ? parseInt(value) : value;
    setFormData((prevState) => ({
      ...prevState,
      [name]: parsedValue,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const token = localStorage.getItem("token");
      const headers = {
        "Auth-Access-Token": token,
      };

      const response = await axios.post(
        "http://localhost:8000/car/create",
        { data: formData },
        {
          headers,
        }
      );

      if (response.status === 200) {
        setFormData(initialState);
        setShowForm(false);

        toast.success("Car created successfully!");
        onCreate();
      }
    } catch (error) {
      toast.error("An error occurred.");
      console.log(error);
    }
  };

  const toggleForm = () => {
    setShowForm((prevState) => !prevState);
  };

  const isLoggedIn = !!localStorage.getItem("token");

  return (
    <div className="car-create">
      {isLoggedIn && (
        <button className="create-button" onClick={toggleForm}>
          Create Car
        </button>
      )}

      {showForm && (
        <div className="modal-overlay">
          <form className="create-form" onSubmit={handleSubmit}>
            <h2>Create Car</h2>
            {Object.entries(formData).map(([key, value]) => (
              <label key={key}>
                {key.charAt(0).toUpperCase() + key.slice(1)}:
                <input type="text" name={key} value={value} onChange={handleInputChange} required />
              </label>
            ))}
            <div className="buttons">
              <button type="submit">Submit</button>
              <button type="button" onClick={toggleForm}>
                Cancel
              </button>
            </div>
          </form>
          <ToastContainer />
        </div>
      )}
    </div>
  );
};

export default CreateCar;
