import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { FormArray, FormBuilder, FormGroup } from '@angular/forms';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { finalize } from 'rxjs/operators';
import { ConfigService } from '../../services/congig.service';
import { IConfigurationProperty } from '../model/interfaces';

@Component({
  selector: 'adscale-properties',
  templateUrl: './properties.component.html',
  styleUrls: ['./properties.component.sass']
})
export class PropertiesComponent implements OnInit, AfterViewInit {

  @ViewChild(MatSort) sort: MatSort;

  loading: boolean;
  form: FormGroup;
  displayedColumns: string[] = ['name', 'value'];
  dataSource = new MatTableDataSource<IConfigurationProperty>();

  constructor(
    private configService: ConfigService,
    private fb: FormBuilder,
  ) {
    this.form = fb.group({
      rows: fb.array([]),
    });
  }

  ngOnInit() {
    this.loading = true;
    this.configService.getProperties()
      .pipe(
        finalize(() => this.loading = false),
      )
      .subscribe(e => {
        e.forEach(z => this.addRow(z));
        this.dataSource.data = e;
      });
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  submit() {
    if (this.form.invalid) {
      return;
    }

    const data = {};
    for (const control of this.rowsControl.controls) {
      if (control.dirty) {
        const r = control.value;
        data[r.name] = {
          line: r.line,
          value: r.type === 'boolean' ? r.value ? 'true' : 'false' : r.value,
        };
      }
    }

    this.loading = true;
    this.configService.saveProperties(data)
      .pipe(
        finalize(() => this.loading = false),
      )
      .subscribe();
  }

  get rowsControl(): FormArray {
    return this.form.get('rows') as FormArray;
  }

  private addRow(s: IConfigurationProperty) {
    const row = this.fb.group({
      name: [s.name],
      value: [s.value],
      line: [s.line],
      status: [s.status],
      type: [s.type],
    });

    this.rowsControl.push(row);
  }
}
