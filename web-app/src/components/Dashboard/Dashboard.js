import React from 'react';
import './Dashboard.css';
import { GoogleMap, useJsApiLoader, Marker } from "@react-google-maps/api"
import CreateDangerModal from './../CreateDangerModal/CreateDangerModal'

export default function Dashboard() {

  const handleReportDanger = async e => {
    console.log("clicked danger create");
  }
  
  return <div>
      <>
        <CreateDangerModal />
      </>

      <hr/>
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

function Map(lattitude, longitude){
  const key="AIzaSyAXMm1QSVTI3wSLE5zRV66DcTYtFm_iFRY";
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
