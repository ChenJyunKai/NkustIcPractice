import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HamastarsComponent } from './hamastars.component';

describe('HamastarsComponent', () => {
  let component: HamastarsComponent;
  let fixture: ComponentFixture<HamastarsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HamastarsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HamastarsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
