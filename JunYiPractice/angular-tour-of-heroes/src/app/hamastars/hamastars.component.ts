import { Component, OnInit } from '@angular/core';
import { Hamastar } from '../hamastar';
import { HAMASTARS } from '../mock-hamastars';

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
  selectedHamastar?: Hamastar;
  hamastars = HAMASTARS;

  constructor() { }

  ngOnInit(): void {
  }
  OnSelect(hamastar:Hamastar):void{
    this.selectedHamastar=hamastar;
  }

}
