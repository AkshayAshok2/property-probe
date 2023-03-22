import { Component } from '@angular/core';

@Component({
  selector: 'app-property-list',
  templateUrl: './property-list.component.html',
  styleUrls: ['./property-list.component.css']
})
export class PropertyListComponent {
  properties = [
    { name: 'Property 1', sqFootage: '1000 sq. ft.', price: '$1000/month' },
    { name: 'Property 2', sqFootage: '1200 sq. ft.', price: '$1200/month' },
    { name: 'Property 3', sqFootage: '800 sq. ft.', price: '$800/month' }
  ];
}