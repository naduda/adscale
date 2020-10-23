import { BreakpointObserver, Breakpoints, BreakpointState } from '@angular/cdk/layout';
import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Observable } from 'rxjs';
import { catchError, shareReplay } from 'rxjs/operators';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'adscale-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.sass']
})
export class NavigationComponent {

  isHandset$: Observable<BreakpointState>;

  constructor(
    breakpointObserver: BreakpointObserver,
    private http: HttpClient,
    private snackBar: MatSnackBar,
  ) {
    this.isHandset$ = breakpointObserver.observe(Breakpoints.HandsetPortrait)
      .pipe(shareReplay());
  }

  powerOff() {
    this.http.post<void>(`${environment.apiPreffix}/off`, null)
      .pipe(
        catchError(ex => {
          this.showMessage('Something wrong!');
          throw ex;
        })
      )
      .subscribe(_ => {
        this.showMessage('Application was stopped.');
      });
  }

  private showMessage(message: string) {
    this.snackBar.open(message, null, {
      duration: 3000,
    });
  }
}
