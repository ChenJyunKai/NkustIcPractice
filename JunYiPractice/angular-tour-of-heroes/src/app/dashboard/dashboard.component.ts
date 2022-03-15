import { Component, OnInit } from '@angular/core';
import { Hamastar } from '../hamastar';
import { HamastarService } from '../hamastar.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css']
})
export class DashboardComponent implements OnInit {
  hamastars: Hamastar[] = [];
  constructor(private hamastarService: HamastarService) { }

  ngOnInit(): void {
    this.getHamastars();
  }
  getHamastars(): void {
    // console.log("123");
    this.hamastarService.getHamastarsInService()
      .subscribe(hamastars => this.hamastars = hamastars.slice(1, 5));
  }

}
