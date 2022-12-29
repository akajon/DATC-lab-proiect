import React from 'react';
import './Dashboard.css';
import { GoogleMap, useJsApiLoader, useLoadScript, Marker } from "@react-google-maps/api"

export default function Dashboard() {
  
  return <div>
      {Map()}
    </div>
}

function Map(){
  const {isLoaded} = useJsApiLoader({
    googleMapsApiKei: process.env.NEXT_PUBLIC_GOOGLE_MAPS_API_KEY,
  })

  if (!isLoaded) {
    return <h2>Loading...</h2>
  }

  return (
    <div style={{ height: '100vh', width: '100%' }}>
      <GoogleMap 
        zoom={15}
        center={{lat: 45.756745 , lng: 21.228737}}
        mapContainerClassName="map-container"
      ></GoogleMap>
    </div>
  )
}
