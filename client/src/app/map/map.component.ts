import { Component, OnInit } from '@angular/core';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import { fromLonLat } from 'ol/proj.js';

@Component({
  selector: 'map-root',
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
        zoom: 13,
      }),
      layers: [
        new TileLayer({
          source: new OSM(),
        }),
      ],
    target: 'map',
    });
  }
}
