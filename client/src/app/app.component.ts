import { Component, OnInit, ɵsetAllowDuplicateNgModuleIdsForTest } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { provideProtractorTestingSupport } from '@angular/platform-browser';
import { interval, take, lastValueFrom } from 'rxjs';

interface IPropertyListing {
  search_term: string
  //owner: string
  //address: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  public search_term = ''
  //public owner = ''
  //public address = ''
  public propertyListing: IPropertyListing[] = []

  constructor(
    private httpClient: HttpClient
  ){}

    async ngOnInit() {
      await this.loadProperties()
    }

  async loadProperties() {
    this.propertyListing = await lastValueFrom(this.httpClient.get<IPropertyListing[]>('/api/properties'))
  }

  // async addProperty() {
  //   await lastValueFrom(this.httpClient.post('/api/properties', {
  //     //owner: this.owner,
  //     //address: this.address
  //   }))
    
  //   await this.loadProperties()
  //   this.search_term = ''
  //   this.owner = ''
  //   this.address = ''
  // }

  async search() {
    await lastValueFrom(this.httpClient.post('/api/properties', {
      search_term: this.search_term
    }))

    await this.loadProperties()
    this.search_term = ''
  }
}
