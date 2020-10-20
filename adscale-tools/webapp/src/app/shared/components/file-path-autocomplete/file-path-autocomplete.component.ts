import { FocusMonitor } from '@angular/cdk/a11y';
import { coerceBooleanProperty } from '@angular/cdk/coercion';
import { HttpClient } from '@angular/common/http';
import { Component, ElementRef, HostBinding, Input, OnDestroy, OnInit, Optional, Self } from '@angular/core';
import { AbstractControl, FormBuilder, FormGroup, NgControl } from '@angular/forms';
import { MatFormFieldControl } from '@angular/material/form-field';
import { Subject } from 'rxjs';
import { Observable } from 'rxjs/internal/Observable';
import { debounceTime, distinctUntilChanged, switchMap, tap } from 'rxjs/operators';
import { environment } from 'src/environments/environment';
import { AValueAccessor } from '../../utils/abstract-value-accessor';

export interface IPath {
  path: string;
  name: string;
}

@Component({
  selector: 'adscale-file-path-autocomplete',
  templateUrl: './file-path-autocomplete.component.html',
  styleUrls: ['./file-path-autocomplete.component.sass'],
  providers: [
    { provide: MatFormFieldControl, useExisting: FilePathAutocompleteComponent },
  ],
})
export class FilePathAutocompleteComponent extends AValueAccessor implements MatFormFieldControl<string>, OnInit, OnDestroy {

  static nextId = 0;
  @HostBinding() id = `file-path-autocomplete-${FilePathAutocompleteComponent.nextId++}`;

  @Input()
  get value(): string {
    return this.pathControl.value;
  }
  set value(value: string | null) {
    this.pathControl.setValue(value);
    this.onUpdate(value);
  }

  @Input()
  get placeholder() {
    return this._placeholder;
  }
  set placeholder(plh) {
    this._placeholder = plh;
    this.stateChanges.next();
  }
  private _placeholder: string;

  @Input()
  get required() {
    return this._required;
  }
  set required(req) {
    this._required = coerceBooleanProperty(req);
    this.stateChanges.next();
  }
  private _required = false;

  @Input()
  get disabled(): boolean { return this._disabled; }
  set disabled(value: boolean) {
    this._disabled = coerceBooleanProperty(value);
    this._disabled ? this.form.disable() : this.form.enable();
    this.stateChanges.next();
  }
  private _disabled = false;

  get empty() {
    return !this.pathControl.value;
  }

  @HostBinding('class.floating')
  get shouldLabelFloat() {
    return this.focused || !this.empty;
  }

  controlType = 'file-path-autocomplete';
  focused: boolean;
  errorState: boolean;
  autofilled?: boolean;
  userAriaDescribedBy?: string;

  form: FormGroup;
  filteredOptions: Observable<string[]>;

  stateChanges = new Subject<void>();

  private dirPath: string;

  constructor(
    @Optional() @Self() public ngControl: NgControl,
    private http: HttpClient,
    private fm: FocusMonitor,
    private elRef: ElementRef<HTMLElement>,
    fb: FormBuilder,
  ) {
    super();

    if (this.ngControl != null) {
      this.ngControl.valueAccessor = this as any;
    }

    fm.monitor(elRef.nativeElement, true).subscribe(origin => {
      this.focused = !!origin;
      this.stateChanges.next();
    });

    this.form = fb.group({
      path: [null],
    });
  }

  ngOnInit(): void {
    this.filteredOptions = this.form.get('path').valueChanges
      .pipe(
        distinctUntilChanged(),
        debounceTime(300),
        tap(e => {
          this.value = e;
          this.dirPath = this.getDirPath(e);
        }),
        switchMap(path => this.http.post<string[]>(`${environment.apiPreffix}/file-path-autocomplete`, { path })),
      );
  }

  ngOnDestroy() {
    this.stateChanges.complete();
    this.fm.stopMonitoring(this.elRef.nativeElement);
  }

  writeValue(value: string): void {
    this.value = value;
  }

  onUpdate(value?: string): void {
    this.onChange(value);
    this.onTouched();
  }

  setDescribedByIds(ids: string[]) {
  }

  onContainerClick(event: MouseEvent): void {
  }

  onSelect(option: string) {
    this.value = this.dirPath + '/' + option;
  }

  get pathControl(): AbstractControl {
    return this.form.get('path');
  }

  private getDirPath(v: string): string {
    return v.substring(0, v.lastIndexOf('/'));
  }
}
