import { Component, OnInit } from '@angular/core';
import Map from 'ol/Map';
import View from 'ol/View';
import TileLayer from 'ol/layer/Tile';
import OSM from 'ol/source/OSM';
import {fromLonLat} from 'ol/proj.js';

@Component({
  selector: 'map-root',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css'],
})

export class mapMaker implements OnInit {

  map: Map;

  ngOnInit(): void {
    this.map = new Map({
      view: new View({
        center: fromLonLat([-82.324,29.654]),
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
