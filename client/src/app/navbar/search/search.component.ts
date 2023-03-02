import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { provideProtractorTestingSupport } from '@angular/platform-browser';
import { interval, take, lastValueFrom } from 'rxjs';

interface ISearchTerm {
  search_term: string
}

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {
  public search_term = ''
  public searchHistory: ISearchTerm[] = []

  constructor(
    private httpClient: HttpClient
  ){}

  async ngOnInit() {
    // await this.loadSearches()
  }

  async loadSearches() {
    this.searchHistory = await lastValueFrom(this.httpClient.get<ISearchTerm[]>('/api/search'))
  }

  async search() {
    await lastValueFrom(this.httpClient.post('/api/search', {
      search_term: this.search_term
    }))

    await this.loadSearches()
    this.search_term = ''
  }
}
