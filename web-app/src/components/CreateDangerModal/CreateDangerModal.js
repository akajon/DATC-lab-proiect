import React, { useState, Component } from "react";
import { createRoot } from 'react-dom/client';

import "./CreateDangerModal.css";
import { GoogleMap, useJsApiLoader, Marker} from "@react-google-maps/api"

export default function Modal() {
  const [modal, setModal] = useState(false);

  const ref = React.useRef(null);

  const toggleModal = () => {
    setModal(!modal);
    const container = document.getElementById("mapp");
    const root = createRoot(container); // createRoot(container!) if you use TypeScript
    root.render(<p>Hello</p>);
  };

  const  Map = async (lattitude, longitude) => {
    const key="AIzaSyDg2UhWgxYuS6dZgb7KO-8H_0yM6xEeQk8";
    const {isLoaded} = useJsApiLoader({
      googleMapsApiKey: key,
    })

    if (!isLoaded) {
      return <h2>Loading...</h2>
    }

    return (
      <div style={{ height: '100vh', width: '100%' }}>
        <GoogleMap
          zoom={15}
          center={{lat: lattitude , lng: longitude}}
          mapContainerClassName="map-container"
        ></GoogleMap>
      </div>
    )
  }

  const handleReportSubmit = async e => {
    // setLogged(true);
    // return;
    // e.preventDefault();
    // const token = await loginUser({
    //   username,
    //   password
    // });
    // setToken(token);
    return
  }

 

  if(modal) {
    document.body.classList.add('active-modal')
  } else {
    document.body.classList.remove('active-modal')
  }

  return (
    <>
      <button onClick={toggleModal} className="btn-modal">
        Report Danger
      </button>

      {modal && (
        <div className="modal">
          <div onClick={toggleModal} className="overlay"></div>
          <div className="modal-content">
            <h2>Report Danger </h2>

            <hr/>

            <form onSubmit={handleReportSubmit}>
              <label>
                <p>Category</p>
                <input type="text"/>
              </label>
              <label>
                <p>Name</p>
                <input type="text"/>
              </label>
              <label>
                <p>Description</p>
                <input type="text"/>
              </label>
              <label>
                <p>Grade</p>
                <input type="text"/>
              </label>
              <div>
                <br/>
                <button type="submit">Send report</button>
              </div>
            </form>
            <button className="close-modal" onClick={toggleModal}>
              Close
            </button>
          </div>
          <div id="mapp" style={{ height: '100vh', width: '100%' }} ref={ref}>
            {/* {Map(45.756745, 21.228737)} */}
            {/* <GoogleMap
              zoom={15}
              center={{lat: 45.756745 , lng: 21.228737}}
              mapContainerClassName="map-containerr"
            ></GoogleMap> */}
          </div>
        </div>
      )}
    </>
  );
}