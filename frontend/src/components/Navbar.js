import React from "react";
import LogoutButton from "./Logout";
import CreateCar from "./CreateCar";

const Navbar = () => {
  return (
    <div>
      <nav className="navbar navbar-expand-xl navbar-dark bg-dark">
        <div className="container-fluid">
          <a className="navbar-brand" href="/">
            CarShop
          </a>
          <div className="collapse navbar-collapse" id="navbarNavAltMarkup">
            <div className="navbar-nav">
              <a className="nav-link active" aria-current="page" href="/carlist">
                CarList
              </a>
            </div>
          </div>
        </div>
        <CreateCar />
        <LogoutButton />
      </nav>
    </div>
  );
};

export default Navbar;
