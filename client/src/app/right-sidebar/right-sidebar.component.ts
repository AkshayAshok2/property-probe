import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-right-sidebar',
  template: `

    <div *ngIf="property" >
      <h2 style="order: 1; width: 100%;flex-flow: row wrap;">{{ property.address }}</h2>
      <p style="order: 2; column-count: 2;">{{ property.assessedvalue }} - {{ property.date }}</p>
      <p> {{ property.description }}</p>
    </div>
  `,
})
export class RightSidebarComponent {
  @Input() property: any;
}
