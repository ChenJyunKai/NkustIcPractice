import { TestBed } from '@angular/core/testing';

import { HamastarService } from './hamastar.service';

describe('HamastarService', () => {
  let service: HamastarService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(HamastarService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
