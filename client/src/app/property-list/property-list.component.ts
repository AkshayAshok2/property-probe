import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { PropertyBoxComponent } from './property-box/property-box.component';
import { HttpClient } from '@angular/common/http';
import { interval, take, lastValueFrom } from 'rxjs';

interface PropertyTerm {
  date            :string
	AuctionType     :string
	JudgementAmount :Float64Array
	address         :string
	assessedvalue   :Float64Array
	LatLon          :string
	Description     :string
	zip_code        :string
}

@Component({
  selector: 'property-container',
  templateUrl: './property-list.component.html',
  styleUrls: ['./property-list.component.css']
})
export class PropertyListComponent implements OnInit {
  public allProperties: PropertyTerm[] = []

  constructor(
    private httpClient: HttpClient
  ){}

  async loadProperties() {
    this.allProperties = await lastValueFrom(this.httpClient.get<PropertyTerm[]>('/api/properties'))
  }
  
  @Output() propertySelected = new EventEmitter<any>();
  selectedProperty: any;


  onSelect(property: any) {
    this.selectedProperty = property;
    this.propertySelected.emit(property);
  }
  async ngOnInit() {
    await this.loadProperties()
  }
}
