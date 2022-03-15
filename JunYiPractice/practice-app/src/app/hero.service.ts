import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Hero } from './hero';
import { HEROES } from './mock-heroes';
import { HttpClient } from '@angular/common/http';
@Injectable({
  providedIn: 'root'
})

export class HeroService {
  constructor(
    private http: HttpClient){}

  getHeroes(): Observable<Hero[]> {
    this.http.get('http://localhost:8080/company/select')
    .subscribe(data => {
      console.log(data);}
    );
    const heroes = of(HEROES);
    return heroes;
  }

}
