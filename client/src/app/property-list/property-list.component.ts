import { Component } from '@angular/core';
import { PropertyBoxComponent } from './property-box/property-box.component';

@Component({
  selector: 'property-container',
  templateUrl: './property-list.component.html',
  styleUrls: ['./property-list.component.css']
})
export class PropertyListComponent {
  properties = [
    { name: 'Property 1', sqFootage: '1000 sq. ft.', price: '$1000/month' },
    { name: 'Property 2', sqFootage: '1200 sq. ft.', price: '$1200/month' },
    { name: 'Property 3', sqFootage: '800 sq. ft.', price: '$800/month' },
    { name: 'Property 4', sqFootage: '1100 sq. ft.', price: '$900/month'},
    { name: 'Property 5', sqFootage: '1300 sq. ft.', price: '$750/month'}
  ];

  // property = {};

  // reset() {
  //   this.property = {};
  // }
}
