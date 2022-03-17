import { Component, OnInit ,Input} from '@angular/core';
import { Hamastar } from '../hamastar';

@Component({
  selector: 'app-hamastar-detail',
  templateUrl: './hamastar-detail.component.html',
  styleUrls: ['./hamastar-detail.component.css']
})
export class HamastarDetailComponent implements OnInit {
  @Input() hamastar?:Hamastar;
  

  constructor() { }

  ngOnInit(): void {
  }
  

}
