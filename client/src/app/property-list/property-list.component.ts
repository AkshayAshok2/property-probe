import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { PropertyBoxComponent } from './property-box/property-box.component';

@Component({
  selector: 'property-container',
  templateUrl: './property-list.component.html',
  styleUrls: ['./property-list.component.css']
})
export class PropertyListComponent implements OnInit {
  properties = [
    { name: 'Property 1', sqFootage: '1000 sq. ft.', price: '$1000/month' },
    { name: 'Property 2', sqFootage: '1200 sq. ft.', price: '$1200/month' },
    { name: 'Property 3', sqFootage: '800 sq. ft.', price: '$800/month' },
    { name: 'Property 4', sqFootage: '1100 sq. ft.', price: '$900/month'},
    { name: 'Property 5', sqFootage: '1300 sq. ft.', price: '$750/month'},
    { name: 'Property 6', sqFootage: '1000 sq. ft.', price: '$1000/month'}, 
    { name: 'Property 7', sqFootage: '1200 sq. ft.', price: '$1200/month' },  
    { name: 'Property 8', sqFootage: '800 sq. ft.', price: '$800/month' },  
    { name: 'Property 9', sqFootage: '1100 sq. ft.', price: '$900/month'},  
    { name: 'Property 10', sqFootage: '1300 sq. ft.', price: '$750/month'},  
    { name: 'Property 11', sqFootage: '1000 sq. ft.', price: '$1000/month' },  
    { name: 'Property 12', sqFootage: '1200 sq. ft.', price: '$1200/month' },  
    { name: 'Property 13', sqFootage: '800 sq. ft.', price: '$800/month' },  
    { name: 'Property 14', sqFootage: '1100 sq. ft.', price: '$900/month'},  
    { name: 'Property 15', sqFootage: '1300 sq. ft.', price: '$750/month'},  
    { name: 'Property 16', sqFootage: '1000 sq. ft.', price: '$1000/month' },  
    { name: 'Property 17', sqFootage: '1200 sq. ft.', price: '$1200/month' },  
    { name: 'Property 18', sqFootage: '800 sq. ft.', price: '$800/month' },  
    { name: 'Property 19', sqFootage: '1100 sq. ft.', price: '$900/month'},  
    { name: 'Property 20', sqFootage: '1300 sq. ft.', price: '$750/month'}

  ];
  
  @Output() propertySelected = new EventEmitter<any>();
  selectedProperty: any;


  onSelect(property: any) {
    this.selectedProperty = property;
    this.propertySelected.emit(property);
  }
  ngOnInit() {}
}
