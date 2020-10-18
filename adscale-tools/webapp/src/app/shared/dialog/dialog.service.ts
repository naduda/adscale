import { ComponentType } from '@angular/cdk/portal';
import { Injectable, TemplateRef } from '@angular/core';
import { MatDialog, MatDialogConfig, MatDialogRef } from '@angular/material/dialog';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { ConfirmComponent, IConfirm } from './confirm/confirm.component';

type CompTypeOrTemplRef<T> = ComponentType<T> | TemplateRef<T>;

@Injectable({
  providedIn: 'root'
})
export class SharedDialogService {

  constructor(private dialog: MatDialog) { }

  openDialog<T, R = any>(compOrTemplRef: CompTypeOrTemplRef<T>, data: any, settings: Partial<MatDialogConfig> = {}): MatDialogRef<T, R> {

    return this.dialog.open(compOrTemplRef, {
      disableClose: true,
      panelClass: ['dialog-helper'],
      data,
      ...settings
    });
  }

  openConfirm(data: IConfirm): Observable<boolean> {
    return this.openDialog(ConfirmComponent, data)
      .afterClosed()
      .pipe(
        map(e => !!e)
      );
  }
}
