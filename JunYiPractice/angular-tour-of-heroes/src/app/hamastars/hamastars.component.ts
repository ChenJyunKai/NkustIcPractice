import { Component, OnInit } from '@angular/core';
import { Hamastar } from '../hamastar';
import { HAMASTARS } from '../mock-hamastars';
import { HamastarService } from '../hamastar.service';
import { MessageService } from '../message.service';


@Component({
  selector: 'app-hamastars',
  templateUrl: './hamastars.component.html',
  styleUrls: ['./hamastars.component.css']
})
export class HamastarsComponent implements OnInit {
  // hamastar :Hamastar={
  //   id:1,
  //   name:"Junyi",
  // }


  // selectedHamastar?: Hamastar;


  hamastars: Hamastar[] = [];
  
  constructor(private hamastarService: HamastarService,private messageService: MessageService) { }

  ngOnInit(): void {
    // console.log("123");
    this.getHamastars();
  }
  // OnSelect(hamastar:Hamastar):void{
  //   this.selectedHamastar=hamastar;
  //   this.messageService.addInService(`HeroesComponent: Selected hero id=${hamastar.id}`);
  // }

  getHamastars(): void {

    this.hamastarService.getHamastarsInService().subscribe(hamastars => this.hamastars = hamastars)
  }

}
