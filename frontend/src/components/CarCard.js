import React, { useState, useEffect } from "react";
import axios from "axios";
import { toast, ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const CarCard = () => {
  const [carList, setCarList] = useState([]);
  const [metadata, setMetadata] = useState({ limit: 4, offset: 0 });
  const [totalCount, setTotalCount] = useState(0);
  const [editedData, setEditedData] = useState({});
  const [enlargedImage, setEnlargedImage] = useState(null);

  useEffect(() => {
    fetchCarList();
  }, [metadata]);

  const fetchCarList = async () => {
    try {
      const token = localStorage.getItem("token");
      const headers = {
        "Auth-Access-Token": token,
      };

      const response = await axios.post("http://localhost:8000/car/list", metadata, { headers });

      const responseData = response.data;
      const decodedData = JSON.parse(atob(responseData.data));

      if (decodedData && decodedData.length) {
        const slicedData = decodedData.slice(metadata.offset, metadata.offset + metadata.limit);
        setCarList(slicedData);
        setTotalCount(decodedData.length);
      } else {
        setCarList([]);
        setTotalCount(0);
      }
    } catch (error) {
      console.log(error);
    }
  };

  const handlePreviousPage = () => {
    if (metadata.offset - metadata.limit >= 0) {
      setMetadata((prevMetadata) => ({
        ...prevMetadata,
        offset: prevMetadata.offset - prevMetadata.limit,
      }));
    }
  };

  const handleNextPage = () => {
    setMetadata((prevMetadata) => ({
      ...prevMetadata,
      offset: prevMetadata.offset + prevMetadata.limit,
    }));
  };

  const currentPage = Math.floor(metadata.offset / metadata.limit) + 1;
  const totalPages = Math.ceil(totalCount / metadata.limit);

  const handleDeleteCar = async (carId) => {
    try {
      const token = localStorage.getItem("token");
      const headers = {
        "Auth-Access-Token": token,
      };

      const requestData = {
        data: {
          id: carId,
        },
      };

      const response = await axios.post("http://localhost:8000/car/delete", requestData, {
        headers,
      });

      if (response.data.status) {
        setCarList((prevCarList) => prevCarList.filter((car) => car.id !== carId));
        toast.success("Car deleted successfully!");
      } else {
        toast.error("Failed to delete car.");
      }
    } catch (error) {
      toast.error("An error occurred.");
      console.log(error);
    }
  };

  const handleEditCar = async (carId) => {
    try {
      const token = localStorage.getItem("token");
      const headers = {
        "Auth-Access-Token": token,
      };

      const requestData = {
        data: {
          id: carId,
          ...editedData,
        },
      };

      const response = await axios.post("http://localhost:8000/car/update", requestData, {
        headers,
      });

      if (response.data.status) {
        setCarList((prevCarList) =>
          prevCarList.map((car) => (car.id === carId ? { ...car, ...editedData } : car))
        );
        toast.success("Car updated successfully!");
      } else {
        toast.error("Failed to update car.");
      }
    } catch (error) {
      toast.error("An error occurred.");
      console.log(error);
    }
  };

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    setEditedData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleEnlargeImage = (imageSrc) => {
    setEnlargedImage(imageSrc);
  };

  const handleCloseEnlarged = () => {
    setEnlargedImage(null);
  };

  return (
    <div>
      <h1 className="list">Car List</h1>
      <div className="card-container">
        {carList.map((car) => (
          <div className="car-container" key={car.id}>
            <img
              src={car.picture}
              alt={car.make}
              className="card-img"
              onClick={() => handleEnlargeImage(car.picture)}
            />
            <div className="car-details">
              <h5 className="card-title">Manufacturer: {car.manufacturer}</h5>
              <p className="card-text">Model: {car.model}</p>
              <div className="car-info">
                <p>Fuel: {car.fuel}</p>
                <p>Type: {car.type}</p>
                <p>Price: {car.price}$</p>
              </div>
              <button className="btn-delete" onClick={() => handleDeleteCar(car.id)}>
                Delete
              </button>
              <button className="btn-edit" onClick={() => handleEditCar(car.id)}>
                Edit
              </button>
            </div>
          </div>
        ))}
      </div>
      <div className="pagination">
        <div className="pagination-info">
          Total: {totalCount}, Current Page: {currentPage}, Total Pages: {totalPages}
        </div>
        <button onClick={handlePreviousPage} disabled={metadata.offset === 0}>
          Previous
        </button>
        <button onClick={handleNextPage}>Next</button>
      </div>
      {enlargedImage && (
        <div className="enlarged-image-container" onClick={handleCloseEnlarged}>
          <div className="enlarged-image-overlay"></div>
          <img className="enlarged-image" src={enlargedImage} alt="Enlarged Car" />
        </div>
      )}
      <ToastContainer />
    </div>
  );
};

export default CarCard;
