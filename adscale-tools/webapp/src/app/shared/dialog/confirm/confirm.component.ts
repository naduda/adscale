import { Component, Inject } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

export interface IConfirm {
  title: string;
  desc: string;
  acceptLabel?: string;
  rejecttLabel?: string;
}

@Component({
  selector: 'adscale-confirm',
  templateUrl: './confirm.component.html',
  styleUrls: ['./confirm.component.sass']
})
export class ConfirmComponent {

  form: FormGroup;

  constructor(
    private ref: MatDialogRef<ConfirmComponent>,
    @Inject(MAT_DIALOG_DATA) public data: IConfirm,
    fb: FormBuilder,
  ) {
    this.form = fb.group({});
  }

  reject() {
    this.ref.close(false);
  }
}
