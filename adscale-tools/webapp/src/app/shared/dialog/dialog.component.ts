import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';
import { FormGroup } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';

@Component({
  selector: 'adscale-dialog',
  templateUrl: './dialog.component.html',
  styleUrls: ['./dialog.component.sass']
})
export class DialogComponent implements OnChanges {

  @Input() form: FormGroup;
  @Input() closable = true;
  @Input() title: string;

  @Input() submitText = 'Apply';

  @Input() headerClass: string;
  @Input() footerClass = 'text-right';

  constructor(
    public ref: MatDialogRef<DialogComponent>,
  ) { }

  ngOnChanges(changes: SimpleChanges): void {
    if (changes.submitEvent && changes.submitEvent.currentValue) {
      this.submit();
    }
  }

  submit() {
    if (this.form.invalid) {
      return;
    }

    this.ref.close(this.form.value);
  }
}
