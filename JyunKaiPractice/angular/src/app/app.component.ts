import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'todo';
  constructor(
    private http: HttpClient) { 
      this.http.get("http://localhost:8080/api/select", {observe: 'response', responseType: 'text'})
      .subscribe((res) => {
        console.log(res)
      });
    }
}
