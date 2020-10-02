import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { FilePathAutocompleteComponent } from './components/file-path-autocomplete/file-path-autocomplete.component';

const materialModules = [
  MatAutocompleteModule,
  MatButtonModule,
  MatCheckboxModule,
  MatFormFieldModule,
  MatInputModule,
  MatSortModule,
  MatTableModule,
];

const sharedModules = [
  CommonModule,
  FormsModule,
  ReactiveFormsModule,
  materialModules,
];

const sharedComponents = [
  FilePathAutocompleteComponent,
];

@NgModule({
  declarations: [
    sharedComponents,
  ],
  imports: [
    sharedModules,
  ],
  exports: [
    sharedModules,
    sharedComponents,
  ],
})
export class SharedModule { }
