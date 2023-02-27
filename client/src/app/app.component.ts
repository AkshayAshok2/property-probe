import { Component, OnInit, ɵsetAllowDuplicateNgModuleIdsForTest } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { provideProtractorTestingSupport } from '@angular/platform-browser';
import { interval, take, lastValueFrom } from 'rxjs';

interface ISearchTerm {
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
  public searchHistory: ISearchTerm[] = []

  constructor(
    private httpClient: HttpClient
  ){}

    async ngOnInit() {
      await this.loadSeaches()
    }

  async loadSeaches() {
    this.searchHistory = await lastValueFrom(this.httpClient.get<ISearchTerm[]>('/api/search'))
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
    await lastValueFrom(this.httpClient.post('/api/search', {
      search_term: this.search_term
    }))

    await this.loadSeaches()
    this.search_term = ''
  }
}
