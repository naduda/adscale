import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { catchError, debounceTime, distinctUntilChanged, finalize, switchMap, tap } from 'rxjs/operators';
import { ConfigService } from '../../services/congig.service';
import { IPath } from '../model/interfaces';

@Component({
  selector: 'adscale-settings',
  templateUrl: './settings.component.html',
  styleUrls: ['./settings.component.sass']
})
export class SettingsComponent implements OnInit {

  filteredOptions: Observable<string[]>;

  form: FormGroup;
  loading: boolean;
  errMessage: string;

  private dirPath: string;

  constructor(
    private configService: ConfigService,
    fb: FormBuilder,
  ) {
    this.form = fb.group({
      easyleads: [null, Validators.required],
      repo: [null, Validators.required],
      ui: [null, Validators.required],
      dbIP: [null, Validators.required],
      appPort: [null, Validators.required],
    });
  }

  ngOnInit(): void {
    this.loading = true;
    this.configService.getSettings()
      .pipe(
        finalize(() => this.loading = false),
      )
      .subscribe(e => this.form.setValue(e, { emitEvent: false }));

    this.filteredOptions = this.form.get('easyleads').valueChanges
      .pipe(
        distinctUntilChanged(),
        debounceTime(300),
        tap(e => this.dirPath = this.getDirPath(e)),
        switchMap(e => this.configService.autoCompleteFilePath(e)),
      );
  }

  onSelect(option: IPath, formControlName: string, e) {
    const control = this.form.get(formControlName);
    control.setValue(this.dirPath + '/' + option, { emitEvent: false });
  }

  submit() {
    if (this.form.invalid) {
      return;
    }

    const settings = this.form.value;
    this.loading = true;
    this.errMessage = null;
    this.configService.saveSettings(settings)
      .pipe(
        finalize(() => this.loading = false),
        catchError(err => {
          this.errMessage = err.error?.error;
          throw err;
        }),
      )
      .subscribe();
  }

  getDirPath(v: string): string {
    return v.substring(0, v.lastIndexOf('/'));
  }
}
