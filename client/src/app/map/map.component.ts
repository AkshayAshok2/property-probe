import { Component, OnInit } from '@angular/core';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import { fromLonLat } from 'ol/proj.js';
import Feature from 'ol/Feature';
import Point from 'ol/geom/Point';
import { Icon, Style } from 'ol/style';
import VectorLayer from 'ol/layer/Vector';
import VectorSource from 'ol/source/Vector';

@Component({
  selector: 'map-container',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css'],
})
export class MapMaker implements OnInit {
  
  map: Map = new Map;
  
  ngOnInit(): void {
    let centerCoordinates = [-82.324,29.654];
    let zipCode = '32601'; 

    // Check if zip code is in Gainesville and set center coordinates accordingly
    if (zipCode === '32601') {
      centerCoordinates = [-82.32146, 29.63964];
    } else if (zipCode === '32603') {
      centerCoordinates = [-82.35590, 29.65082];
    } else if (zipCode === '32605') {
      centerCoordinates = [-82.38564, 29.68129];
    } else if (zipCode === '32606') {
      centerCoordinates = [-82.44238, 29.68419];
    }

    this.map = new Map({
      view: new View({
        center: fromLonLat(centerCoordinates),
        zoom: 11,
      }),
      layers: [
        new TileLayer({
          source: new OSM(),
        }),
      ],
    target: 'map',
    });

    // Create a vector source for the markers
    const markerSource = new VectorSource();

    // create markers
    const markers = [
  {
    name: 'Marker 1',
    lon: -82.59850929999999,
    lat: 29.81119469999999,
  },
  {
    name: 'Marker 2',
    lon: -82.4178211,
    lat: 29.7390767,
  },
  {
    name: 'Marker 3',
    lon: -82.09298199999999,
    lat: 29.542478,
  },
  {
    name: 'Marker 4',
    lon: -82.4886009,
    lat: 29.793773,
  },
  {
    name: 'Marker 5',
    lon: -82.4037836,
    lat: 29.68089729999999,
  },
  {
    name: 'Marker 6',
    lon: -82.5502978,
    lat: 29.7281886,
  },
  {
    name: 'Marker 7',
    lon: -82.423918,
    lat: 29.605522,
  },
  {
    name: 'Marker 8',
    lon: -82.5614209,
    lat: 29.771643,
  },
  {
    name: 'Marker 9',
    lon: -82.30743460000001,
    lat: 29.6772913,
  },
  {
    name: 'Marker 10',
    lon: -82.4188873,
    lat: 29.63328079999999,
  },
  {
    name: 'Marker 11',
    lon: -82.29008739999999,
    lat: 29.5031695,
  },
  {
    name: 'Marker 12',
    lon: -82.478787,
    lat: 29.7571521,
  },
  {
    name: 'Marker 13',
    lon: -82.4360255,
    lat: 29.6614827,
  },
  {
    name: 'Marker 14',
    lon: -82.524773,
    lat: 29.608601,
  },
  {
    name: 'Marker 15',
    lon: -82.3347826,
    lat: 29.69196789999999,
  }
    ].map((m) => {
      const feature = new Feature({
        geometry: new Point(fromLonLat([m.lon, m.lat])),
        name: m.name,
      });

      // create marker style
      const iconStyle = new Style({
        image: new Icon({
          anchor: [0.5, 1],
          src: 'https://cdn.mapmarker.io/api/v1/pin?size=50&background=%23006cfc&icon=fa-home&color=%23FFFFFF&voffset=0',
        }),
      });

      feature.setStyle(iconStyle);

      // // add click event listener to marker
      // this.map.on('click', function (evt) {
      //   console.log('Marker clicked:', m.name);
      // });

      return feature;
    });

    // add markers to map
    const vectorLayer = new VectorLayer({
      source: new VectorSource({
        features: markers,
      }),
    });

    this.map.addLayer(vectorLayer);
    } 
}
