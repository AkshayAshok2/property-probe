import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms'
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule} from '@angular/common/http'
// import the MapsModule for the Maps component
import { mapMaker } from './map/map.component';
import { NavbarComponent } from './navbar/navbar.component';
import { SearchComponent } from './navbar/search/search.component'
// Change method of importing Material components using material module
import { MaterialModule } from './material/material.module';

@NgModule({
  declarations: [
    AppComponent,
    mapMaker,
    NavbarComponent,
    SearchComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
<<<<<<< Updated upstream
=======
    MaterialModule
>>>>>>> Stashed changes
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
