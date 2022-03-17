import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HamastarDetailComponent } from './hamastar-detail.component';

describe('HamastarDetailComponent', () => {
  let component: HamastarDetailComponent;
  let fixture: ComponentFixture<HamastarDetailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ HamastarDetailComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(HamastarDetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
