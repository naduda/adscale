import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListModule } from '@angular/material/list';
import { MatMenuModule } from '@angular/material/menu';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatSnackBarModule } from '@angular/material/snack-bar';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { MatToolbarModule } from '@angular/material/toolbar';
import { NgxMaskModule } from 'ngx-mask';
import { FilePathAutocompleteComponent } from './components/file-path-autocomplete/file-path-autocomplete.component';
import { ConfirmComponent } from './dialog/confirm/confirm.component';
import { DialogComponent } from './dialog/dialog.component';

const materialModules = [
  MatAutocompleteModule,
  MatButtonModule,
  MatCheckboxModule,
  MatDialogModule,
  MatFormFieldModule,
  MatIconModule,
  MatInputModule,
  MatListModule,
  MatMenuModule,
  MatSidenavModule,
  MatSlideToggleModule,
  MatSnackBarModule,
  MatSortModule,
  MatTableModule,
  MatToolbarModule,
];

const sharedModules = [
  CommonModule,
  FormsModule,
  ReactiveFormsModule,
  materialModules,
];

const sharedComponents = [
  FilePathAutocompleteComponent,
  DialogComponent,
];

@NgModule({
  declarations: [
    sharedComponents,
    ConfirmComponent,
  ],
  imports: [
    sharedModules,
    NgxMaskModule.forRoot(),
  ],
  exports: [
    sharedModules,
    sharedComponents,
    NgxMaskModule,
  ],
})
export class SharedModule { }
