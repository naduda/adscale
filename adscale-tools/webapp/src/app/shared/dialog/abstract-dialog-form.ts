import { Directive, Inject, Injector, OnDestroy } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { Subject } from 'rxjs';

@Directive()
export abstract class ADialogFormDirective<R, D> implements OnDestroy {

  loading: boolean;
  form: FormGroup;

  protected destroy$ = new Subject<void>();

  constructor(
    public ref: MatDialogRef<R>,
    @Inject(MAT_DIALOG_DATA) public data: D,
    fb: FormBuilder,
    protected injector: Injector,
  ) {
    this.postConstructor(fb, data);
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }

  protected abstract postConstructor(fb: FormBuilder, data: D): void;

  abstract submit(): void;

  close(data = null) {
    this.ref.close(data);
  }
}
