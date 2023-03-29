import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-right-sidebar',
  template: `
    <div *ngIf="property">
      <h2>{{ property.name }}</h2>
      <p>{{ property.price }} - {{ property.sqFootage }}</p>
    </div>
  `,
})
export class RightSidebarComponent {
  @Input() property: any;
}
