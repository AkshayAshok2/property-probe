import { Component, OnInit, Input, Output, EventEmitter, HostListener } from '@angular/core';
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
  @Input() zip!: string | null;
  propertyClicked: boolean = false;
  propertiesExist: boolean = true;
  public allProperties: PropertyTerm[] = []

  // if (this.zip != null) {
    
  // }
  // else {

  // }

  onElementClick(property: any) {
    this.propertyClicked = true;
    this.onSelect(property);
  }

  @HostListener('document:click')
  onDocumentClick() {
    if (!this.propertyClicked) {
      this.onSelect(null);
    }
    this.propertyClicked = false;
    console.log(`Zipcode received at property list: ${this.zip}`);
  }

  constructor(
    private httpClient: HttpClient
  ){}

  async loadProperties() {
    if (this.zip === 'All zipcodes')
      this.allProperties = await lastValueFrom(this.httpClient.get<PropertyTerm[]>('/api/properties/'));
    else
      this.allProperties = await lastValueFrom(this.httpClient.get<PropertyTerm[]>('/api/properties/' + this.zip));
    
    if (this.allProperties.length === 0)
      this.propertiesExist = false;
    else
      this.propertiesExist = true;
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

  async ngOnChanges() {
    await this.loadProperties()
  }
}
