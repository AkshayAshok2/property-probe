import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PropertyBoxComponent } from './property-box.component';

describe('PropertyBoxComponent', () => {
  let component: PropertyBoxComponent;
  let fixture: ComponentFixture<PropertyBoxComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PropertyBoxComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PropertyBoxComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
