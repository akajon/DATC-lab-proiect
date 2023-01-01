import React from 'react';
import './Dashboard.css';
import { GoogleMap, useJsApiLoader, Marker } from "@react-google-maps/api"
import CreateDangerModal from './../CreateDangerModal/CreateDangerModal'

export default function Dashboard() {

  var marker = new Marker(45.756745, 21.228737);

  const handleReportSubmit = (event) => {
    // setLogged(true);
    // return;
    // e.preventDefault();
    // const token = await loginUser({
    //   username,
    //   password
    // });
    // setToken(token);
    event.preventDefault();
    console.log(event.target[0].value)
    return
  }
  
  return <div>
      <div>
        <div>
          <table  style={{align: "center", width: "90%", marginLeft: "5%"}}>
            <tbody>
              <tr>
                <td colSpan="2" style={{textAlign: "center"}}>
                  <h2>Report Danger </h2>
                </td>
              </tr>
              <tr>
                <td style={{textAlign: "center", width: "30%"}}>
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
                </td>
                <td>
                  <div id="map" style={{ height: '80vh', width: '100%' }}>
                    {Map(45.756745, 21.228737)}
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
      <hr/>
      <br/>

      <div style={{textAlign: "center"}}>
        <h2>Reported Dangers </h2>
      </div>

      <table style={{width: '50%'}}>
        <tbody>
          <tr>
            <td>
              Here is the map
            </td>
          </tr>
          <tr>
            <td>
              {/* {Map(45.756745, 21.228737)} */}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
}

function onClick(t, map, coord) {
  const { latLng } = coord;
  const lat = latLng.lat();
  const lng = latLng.lng();

  this.setState(previousState => {
    return {
      markers: [
        ...previousState.markers,
        {
          title: "",
          name: "",
          position: { lat, lng }
        }
      ]
    };
  });
}

function Map(lattitude, longitude){
  const key="AIzaSyDg2UhWgxYuS6dZgb7KO-8H_0yM6xEeQk8";
  const {isLoaded} = useJsApiLoader({
    googleMapsApiKey: key,
  })

  if (!isLoaded) {
    return <h2>Loading...</h2>
  }

  return (
    <div style={{ height: '80vh', width: '100%' }}>
      <GoogleMap
        zoom={15}
        center={{lat: lattitude , lng: longitude}}
        mapContainerClassName="map-container"
        // onClick={onClick}
      >
        <Marker
          key={lattitude}
          position={{lat: lattitude, lng: longitude}}
        />
      </GoogleMap>
    </div>
  )
}
