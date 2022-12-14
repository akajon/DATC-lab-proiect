import React, { useEffect, useState } from 'react';
import './Dashboard.css';
import { GoogleMap, useJsApiLoader, Marker } from "@react-google-maps/api"
import CreateDangerModal from './../CreateDangerModal/CreateDangerModal'
import Cookies from 'js-cookie';

export default function Dashboard(props) {
  var userId, userRole;
  props.user.then(data => {
    userId = data.Id;
    userRole = data.Role;
  })

  const handleDangerSubmit = (event) => {
    event.preventDefault();
    console.log(event.target[0].value);
    fetch('https://datcgoloverbackend.azurewebsites.net/danger/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      "Category": event.target[0].value,
      "Name": event.target[1].value,
      "Description": event.target[2].value,
      "Grade": event.target[3].value,
      "UserRole": userRole
    })
    });
    return
  }

  const handleAlertSubmit = (event) => {
    event.preventDefault();
    console.log(event.target[0].value);
    fetch('https://datcgoloverbackend.azurewebsites.net/alert/add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      "DangerId": event.target[0].value,
      "Latitude": event.target[1].value,
      "Longitude": event.target[2].value,
      "UserId": userId
    })
    });
    return
  }


  const [array, setArray] = useState([]);

  // async function fetchData() {
    fetch('https://datcgoloverbackend.azurewebsites.net/alert', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
    })
      .then(res => {
        res.json().then(data => {
          data.forEach(element => {
            var a = array;
            a.push(element);
            setArray(a);
            console.log(array)
          })
        })
      });
  //   }
  // await fetchData();

  return <div>
      <div>
        <div style={{textAlign: "center"}}>
          <h1 >City Dangers Alert Project</h1>
        </div>
        <hr/>

        <div style={{paddingTop: "2%"}}>
          <table style={{align: "center", width: "96%", marginLeft: "2%"}}>
            <tbody>
              <tr>
                <td colSpan="2" style={{textAlign: "center"}}>
                  <h2>Report Danger </h2>
                </td>
              </tr>
              <tr>
                <td style={{textAlign: "center", width: "30%"}}>
                  <form onSubmit={handleDangerSubmit}>
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
                      <button type="submit">Send danger</button>
                    </div>
                  </form>
                  <hr/>
                  <form onSubmit={handleAlertSubmit}>
                    <label>
                      <p>DangerId</p>
                      <input type="text"/>
                    </label>
                    <label>
                      <p>Latitude</p>
                      <input type="text"/>
                    </label>
                    <label>
                      <p>Longitude</p>
                      <input type="text"/>
                    </label>
                    <div>
                      <br/>
                      <button type="submit">Send alert</button>
                    </div>
                  </form>
                </td>
                <td>
                  <div id="map" style={{ height: '60vh', width: '100%' }}>
                    {Map(45.756745, 21.228737)}
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
          <br/>
        </div>
      </div>
      <hr/>
      <br/>

      <div style={{textAlign: "center"}}>
        <h2>Active Dangers </h2>
      </div>

      <table style={{align: "center", width: "96%", marginLeft: "2%"}}>
        <tbody>
          {array.map((alert) => {
            return <tr key={alert.Id}>
              <td style={{textAlign: "center", width: "30%"}}>
                <p>OwnerId: {alert.OwnerId}</p>
              </td>
              <td>
              <div style={{ height: '60vh', width: '100%' }}>
                <GoogleMap
                  zoom={15}
                  center={{lat: alert.Latitude , lng: alert.Longitude}}
                  mapContainerClassName="map-container"
                >
                  <Marker
                    key={alert.Id} 
                    position={{lat: alert.Latitude , lng: alert.Longitude}}
                  />
                </GoogleMap>
              </div>
              </td>
            </tr>
            })}
        </tbody>
      </table>
    </div>
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
    <div style={{ height: '60vh', width: '100%' }}>
      <GoogleMap
        zoom={15}
        center={{lat: lattitude , lng: longitude}}
        mapContainerClassName="map-container"
      >
        <Marker
          key={lattitude}
          position={{lat: lattitude, lng: longitude}}
        />
      </GoogleMap>
    </div>
  )
}
