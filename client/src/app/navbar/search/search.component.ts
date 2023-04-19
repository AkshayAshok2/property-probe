import { Component, HostListener, Output, EventEmitter, OnInit } from '@angular/core';
import { provideProtractorTestingSupport } from '@angular/platform-browser';
import { of, interval, take, lastValueFrom } from 'rxjs';
import { ViewChild } from '@angular/core';
import { ElementRef } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, ReactiveFormsModule } from '@angular/forms';
import { debounceTime, distinctUntilChanged, switchMap } from 'rxjs/operators';
import { HttpClient } from '@angular/common/http';


@Component({
  selector: 'app-search',
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})



export class SearchComponent implements OnInit {
  @Output() searchInfo = new EventEmitter<any>();


  async loadZips() {
    this.zipcodes = await lastValueFrom(this.httpClient.get<string[]>('/api/properties/zipcodes'));
  }

  clickedInside = false;
  validInput: boolean = true;
  zipcodeForm: FormGroup;

  zipcodes: string[] = [
    'All zipcodes',
  ]
  
  showZipcodeDropdown = false;

  constructor(private fb: FormBuilder, private httpClient: HttpClient) {
    this.zipcodeForm = this.fb.group({
      searchTerm: ['']
    });
  }

  async ngOnInit() {await this.loadZips}

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
    this.zipcodeForm.reset()

    if (!this.searchValidator(searchTerm)) {
      console.log("Input incorrect! You shall not pass.");
      this.validInput = false;
      return;
    }
    this.validInput = true;
    // this.searchHistory.unshift({ searchTerm });
    console.log(`Search by zipcode: ${searchTerm}`);
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
