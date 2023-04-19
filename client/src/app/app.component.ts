import { Component, OnInit, ɵsetAllowDuplicateNgModuleIdsForTest } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { provideProtractorTestingSupport } from '@angular/platform-browser';
import { interval, take, lastValueFrom } from 'rxjs';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  selectedProperty: any;
  searchTerm: any;

  onPropertySelected(property: any) {
    this.selectedProperty = property;
  }

  receiveSearchInfo($event: string) {
    this.searchTerm = $event;
    console.log(`Search term received at app! ${this.searchTerm}`);
  }

  ngOnInit() {}
}

