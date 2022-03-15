import { Injectable } from '@angular/core';
import { Hamastar } from './hamastar';
import { HAMASTARS } from './mock-hamastars';
import { Observable, of } from 'rxjs';
import { MessageService } from './message.service';

@Injectable({
  providedIn: 'root'
})
export class HamastarService {

  constructor(private messageService: MessageService) { }
  // getHamastarsInService(): Observable<Hamastar[]> {
  //   const hamastars =  of(HAMASTARS);
  //   return hamastars;
  // }
  getHamastarsInService(): Observable<Hamastar[]>{
    const all =of(HAMASTARS);
    this.messageService.addInService('HamastarService: fetched Hamastars');
    return all;
  }
  // getHeroes(): Observable<Hero[]> {
  //   const heroes = of(HEROES);
  //   return heroes;
  // }
  getHamastar(id:Number):Observable<Hamastar>{
    const oneselect=HAMASTARS.find(h => h.id === id)!;
    this.messageService.addInService(`HeroService: fetched hamastar id=${id}`);
    console.log("123");
    return of(oneselect);
  }
}
