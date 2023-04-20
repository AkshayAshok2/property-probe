import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent {
  @Input() searchTerm!: string;
  @Output() searchInfo = new EventEmitter<any>();

  receiveSearchInfo($event: string) {
    this.searchTerm = $event;
    console.log(`Search term received at navbar! ${this.searchTerm}`);
    this.searchInfo.emit(this.searchTerm);
  }

  async signIn() {}
}
