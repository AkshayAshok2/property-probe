import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms'
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule} from '@angular/common/http'
// import the MapsModule for the Maps component
import { MapMaker } from './map/map.component';
import { NavbarComponent } from './navbar/navbar.component';
import { SearchComponent } from './navbar/search/search.component';
import { PropertyListComponent } from './property-list/property-list.component';
import { PropertyBoxComponent } from './property-list/property-box/property-box.component';
// Change method of importing Material components using material module
import { MaterialModule } from './material/material.module';
import { RightSidebarComponent } from './right-sidebar/right-sidebar.component';

@NgModule({
  declarations: [
    AppComponent,
    MapMaker,
    NavbarComponent,
    SearchComponent,
    PropertyBoxComponent,
    PropertyListComponent,
    RightSidebarComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    MaterialModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
