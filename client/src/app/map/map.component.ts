import { Component, OnInit, Input } from '@angular/core';
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
import { HttpClient } from '@angular/common/http';
import { interval, take, lastValueFrom } from 'rxjs';
interface PropertyTerm {
	latlon          :string
}
@Component({
  selector: 'map-container',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css'],
})
export class MapMaker implements OnInit {
  @Input() zip!: string | null;
  public allProperties: PropertyTerm[] = []

  constructor(
    private httpClient: HttpClient
  ){}
  
  async loadProperties() {
    this.allProperties = await lastValueFrom(this.httpClient.get<PropertyTerm[]>('/api/properties/'));
  }
  async makeMap(){
    let centerCoordinates = [-82.324,29.654];

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

    // loop through properties and create features
    const features = this.allProperties.map((property) => {
      const [lon, lat] = property.latlon.replace('[','').replace(']','').split(',').map(Number);
      const feature = new Feature({
        geometry: new Point(fromLonLat([lon, lat])),
      });
      const iconStyle = new Style({
        image: new Icon({
          anchor: [0.5, 1],
          src: 'https://cdn.mapmarker.io/api/v1/pin?size=50&background=%23006cfc&icon=fa-home&color=%23FFFFFF&voffset=0',
        }),
      });
      feature.setStyle(iconStyle);
      return feature;
    });

    // create vector source and layer
    const markerSource = new VectorSource({
      features: features,
    });
    const markerLayer = new VectorLayer({
      source: markerSource,
    });

    // add layer to map
    this.map.addLayer(markerLayer);
  }

  CenterMap() {
    let centerCoordinates = [-82.324,29.654];

    // Check if zip code is in Gainesville and set center coordinates accordingly
    if (this.zip === '32601') {
      centerCoordinates = [-82.32146, 29.63964];
    } else if (this.zip === '32603') {
      centerCoordinates = [-82.35590, 29.65082];
    } else if (this.zip === '32605') {
      centerCoordinates = [-82.38564, 29.68129];
    } else if (this.zip === '32606') {
      centerCoordinates = [-82.44238, 29.68419];
    }else if (this.zip === '32669') {
      centerCoordinates = [-82.6267, 29.6534];
    }else if (this.zip === '33196') {
      centerCoordinates = [-80.4667, 25.6698];
    }else if (this.zip === '33186') {
      centerCoordinates = [-80.3993, 25.6654];
    }else if (this.zip === '33015') {
      centerCoordinates = [-80.3269, 25.9359];
    }else if (this.zip === '33165') {
      centerCoordinates = [-80.3547, 25.7328];
    }else if (this.zip === '33125') {
      centerCoordinates = [-80.2349, 25.7862];
    }else if (this.zip === '33147') {
      centerCoordinates = [-80.2377, 25.8479];
    }else if (this.zip === '33149') {
      centerCoordinates = [-80.1596, 25.7138];
    }else if (this.zip === '33169') {
      centerCoordinates = [-80.2154, 25.9401];
    }else if (this.zip === '33032') {
      centerCoordinates = [-80.3770, 25.5232];
    }else if (this.zip === '33055') {
      centerCoordinates = [-80.2823, 25.9446];
    }

    console.log(fromLonLat(centerCoordinates));
    this.map.getView().setCenter(fromLonLat(centerCoordinates));
    this.map.getView().setZoom(12);
  }

  map: Map = new Map;
  

  async ngOnInit() {
    await this.loadProperties()
    await this.makeMap()
  } 
  async ngOnChanges() {
    await this.loadProperties()
    this.CenterMap()
  }
}
