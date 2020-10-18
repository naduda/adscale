import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { FormArray, FormBuilder, FormGroup } from '@angular/forms';
import { MatSort } from '@angular/material/sort';
import { MatTableDataSource } from '@angular/material/table';
import { filter, finalize, switchMap, tap } from 'rxjs/operators';
import { DockerService } from 'src/app/docker/services/docker.service';
import { SharedDialogService } from 'src/app/shared/dialog/dialog.service';
import { ConfigService } from '../../services/congig.service';
import { IConfigurationProperty } from '../model/interfaces';
import { AddPropertyDialogComponent } from './add-property-dialog/add-property-dialog.component';

@Component({
  selector: 'adscale-properties',
  templateUrl: './properties.component.html',
  styleUrls: ['./properties.component.sass']
})
export class PropertiesComponent implements OnInit, AfterViewInit {

  @ViewChild(MatSort) sort: MatSort;

  loading: boolean;
  containerExists: boolean;
  form: FormGroup;
  displayedColumns: string[] = ['idx', 'ico', 'name', 'value'];
  dataSource = new MatTableDataSource<IConfigurationProperty>();

  constructor(
    private configService: ConfigService,
    private dockerService: DockerService,
    private dialogService: SharedDialogService,
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
        this.dataSource._updateChangeSubscription();
      });

    this.dockerService.state$.subscribe(e => this.containerExists = e.containerExists);
  }

  ngAfterViewInit() {
    this.dataSource.sort = this.sort;
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.dataSource.filter = filterValue.trim().toLowerCase();
  }

  addProperty() {
    this.dialogService.openDialog(AddPropertyDialogComponent, null)
      .afterClosed()
      .pipe(
        filter(e => !!e),
        tap(_ => this.loading = true),
        switchMap(e => this.configService.addProperty(e).pipe(finalize(() => this.loading = false))),
      )
      .subscribe(e => {
        this.rowsControl.clear();
        e.forEach(z => this.addRow(z));
        this.dataSource.data = e;
        this.dataSource._updateChangeSubscription();
      });
  }

  delete(idx: number, line: number) {
    const data = {
      title: 'Are you sure?',
      desc: 'You want to delete current property',
    };

    this.dialogService.openConfirm(data)
      .pipe(
        filter(Boolean),
        tap(_ => this.loading = true),
        switchMap(_ => this.configService.removeLine(line).pipe(finalize(() => this.loading = false))),
      )
      .subscribe(_ => {
        this.dataSource.data.splice(idx, 1);
        this.dataSource._updateChangeSubscription();
      });
  }

  removeExtraLines() {
    this.loading = true;
    this.configService.removeExtraLines()
      .pipe(finalize(() => this.loading = false))
      .subscribe(e => {
        this.dataSource.data = e;
        this.dataSource._updateChangeSubscription();
      });
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
          enabled: r.enabled,
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

  copyToContainer() {
    this.loading = true;
    this.configService.copyToContainer()
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
      enabled: [s.enabled],
      type: [s.type],
    });

    this.rowsControl.push(row);
  }
}
