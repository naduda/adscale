import { Component } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'adscale-add-property-dialog',
  templateUrl: './add-property-dialog.component.html',
  styleUrls: ['./add-property-dialog.component.sass']
})
export class AddPropertyDialogComponent {

  form: FormGroup;

  constructor(
    fb: FormBuilder,
  ) {
    this.form = fb.group({
      name: [null, Validators.required],
      value: [null, Validators.required],
    });
  }

  get nameControl(): AbstractControl {
    return this.form.get('name');
  }

  get valueControl(): AbstractControl {
    return this.form.get('value');
  }
}
