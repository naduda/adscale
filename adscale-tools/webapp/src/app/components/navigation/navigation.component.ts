import { BreakpointObserver, Breakpoints, BreakpointState } from '@angular/cdk/layout';
import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { shareReplay } from 'rxjs/operators';

@Component({
  selector: 'adscale-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.sass']
})
export class NavigationComponent {

  isHandset$: Observable<BreakpointState>;

  constructor(
    breakpointObserver: BreakpointObserver,
  ) {
    this.isHandset$ = breakpointObserver.observe(Breakpoints.HandsetPortrait)
      .pipe(shareReplay());
  }

}
