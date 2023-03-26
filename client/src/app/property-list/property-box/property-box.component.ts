import { Component, Input } from '@angular/core';

@Component({
  selector: 'property-box',
  templateUrl: './property-box.component.html',
  styleUrls: ['./property-box.component.css']
})
export class PropertyBoxComponent {
  // Import from backend housing data
  @Input() currentProperty = [];
}
