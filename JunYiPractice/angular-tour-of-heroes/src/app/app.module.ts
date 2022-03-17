import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HamastarsComponent } from './hamastars/hamastars.component';

import { HamastarDetailComponent } from './hamastar-detail/hamastar-detail.component';

@NgModule({
  declarations: [
    AppComponent,
    HamastarsComponent,

    HamastarDetailComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
