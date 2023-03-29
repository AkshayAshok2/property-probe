import { ComponentFixture, TestBed } from '@angular/core/testing';
import { MapMaker } from './map.component';

describe('NavbarComponent', () => {
  let component: MapMaker;
  let fixture: ComponentFixture<MapMaker>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MapMaker ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MapMaker);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
