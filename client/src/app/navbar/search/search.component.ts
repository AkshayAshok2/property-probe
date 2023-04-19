import { Component, HostListener, Output, EventEmitter, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { provideProtractorTestingSupport } from '@angular/platform-browser';
import { of, interval, take, lastValueFrom } from 'rxjs';
import { ViewChild } from '@angular/core';
import { ElementRef } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, ReactiveFormsModule } from '@angular/forms';
import { debounceTime, distinctUntilChanged, switchMap } from 'rxjs/operators';
import { DataService } from 'src/app/data.service';

interface ISearchTerm {
  searchTerm: string
}

@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent implements OnInit {
  @Output() searchInfo = new EventEmitter<any>();
  // public selected = ''
  // public searchTerm = ''
  // public searchHistory: ISearchTerm[] = []
  // public zipCodes: number[] = []

  // @ViewChild('zipCodeSearch') zipCodeSearchInput!: ElementRef;

  // public dummyZipcodes: number[] = [
  //   10000,
  //   20000,
  //   30000,
  //   40000,
  //   50000
  // ]

  // constructor(
  //   private httpClient: HttpClient
  // ){}

  // ngOnInit() {
  //   this
  // }

  // private getZipCodes() {
  //   return this.dummyZipcodes
  //   // implement to get zipcodes from backend
  // }

  // async loadSearches() {
  //   this.searchHistory = await lastValueFrom(this.httpClient.get<ISearchTerm[]>('/api/search'))
  // }

  // loadZipcodes() {
  //   console.log(this.zipCodeSearchInput.nativeElement.value)
  // }

  // async search() {
  //   await lastValueFrom(this.httpClient.post('/api/search', {
  //     searchTerm: this.searchTerm
  //   }))

  //   await this.loadSearches()
  //   this.searchTerm = ''
  // }

  clickedInside = false;
  validInput: boolean = true;
  zipcodeForm: FormGroup;
  searchHistory: ISearchTerm[] = [];
  zipcodes: string[] = [
    '32601',
    '32602',
    '32603',
    '32604',
    '32605',
    '32606',
    '32607',
    '32608',
    '32609',
    '32610',
    '32615',
    '32618',
    '32622',
    '32640',
    '32641',
    '32643',
    '32653',
    '32667',
    '32669',
    '33012',
    '33015',
    '33055',
    '33054',
    '33142',
    '33162',
    '33181',

  ]
  showZipcodeDropdown = false;

  constructor(private fb: FormBuilder) {
    this.zipcodeForm = this.fb.group({
      searchTerm: ['']
    });
  }

  ngOnInit(): void {}

  toggleZipcodeDropdown(): void {
    this.showZipcodeDropdown = !this.showZipcodeDropdown;
  }

  activateDropdown(): void {
    if (!this.showZipcodeDropdown)
      this.showZipcodeDropdown = true;
  }

  deactivateDropdown(): void {
    if (this.showZipcodeDropdown)
      this.showZipcodeDropdown = false;
  }

  selectZipcode(zipcode: string): void {
    this.zipcodeForm.get('searchTerm')?.setValue(zipcode);
    this.deactivateDropdown();
  }

  searchValidator(search: string | null): boolean {
    if (search && this.zipcodes.includes(search))
      return true;
    return false;
  }

  searchByZipcode() {
    this.deactivateDropdown()
    const searchTerm = this.zipcodeForm.get('searchTerm')?.value;
    // const searchTerm = this.selectedZipcode
    this.zipcodeForm.reset()

    if (!this.searchValidator(searchTerm)) {
      console.log("Input incorrect! You shall not pass.");
      this.validInput = false;
      return;
    }
    this.validInput = true;
    this.searchHistory.unshift({ searchTerm });
    console.log(`Search by zipcode: ${searchTerm}`);

    console.log(`Does zip move? ${searchTerm}`);
    this.searchInfo.emit(searchTerm);
  }

  // turn off dropdown if click outside of search box
  

  onElementClick() {
    this.clickedInside = true;
    this.activateDropdown()
  }

  @HostListener('document:click')
  onDocumentClick() {
    if (!this.clickedInside) {
      this.deactivateDropdown()
    }
    this.clickedInside = false;
  }
}
