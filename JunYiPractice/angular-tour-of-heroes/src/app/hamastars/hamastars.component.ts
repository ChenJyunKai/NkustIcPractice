import { Component, OnInit } from '@angular/core';
import { Hamastar } from '../hamastar';
import { HAMASTARS } from '../mock-hamastars';
import { HamastarService } from '../hamastar.service';

@Component({
  selector: 'app-hamastars',
  templateUrl: './hamastars.component.html',
  styleUrls: ['./hamastars.component.css']
})
export class HamastarsComponent implements OnInit {
  hamastar :Hamastar={
    id:1,
    name:"Junyi",
  }
  hamastarss: Hamastar[] = [];
  selectedHamastar?: Hamastar;
  hamastars = HAMASTARS;

  constructor(private heroService: HeroService) { }

  ngOnInit(): void {
  }
  OnSelect(hamastar:Hamastar):void{
    this.selectedHamastar=hamastar;
  }

  getHeroes(): void {
    this.hamastarss = this.heroService.getHeroes();
  }

}
