import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { NotesComponent } from './notes/notes.component';

import { FormsModule } from '@angular/forms';
import { NotesService } from './notes.service';
import { HttpClient, HttpHandler, HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    NotesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [
    NotesService,
    HttpClient
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
