import { Component, OnInit ,Input} from '@angular/core';
import { Hamastar } from '../hamastar';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { HamastarService } from '../hamastar.service';

@Component({
  selector: 'app-hamastar-detail',
  templateUrl: './hamastar-detail.component.html',
  styleUrls: ['./hamastar-detail.component.css']
})
export class HamastarDetailComponent implements OnInit {
  @Input() hamastar?:Hamastar;
  

  constructor( private route: ActivatedRoute,
    private hamaService: HamastarService,
    private location: Location) 
    { }

  ngOnInit(): void {
    this.getHamastar();
  }
  getHamastar():void{
    const id = Number(this.route.snapshot.paramMap.get('id'));
    console.log(id);
    this.hamaService.getHamastar(id)
      .subscribe(hamastar => this.hamastar=hamastar)
  }
  goBack(): void {  
    this.location.back();
  }
  

}
