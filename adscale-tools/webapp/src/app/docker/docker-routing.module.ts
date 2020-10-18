import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CreateComponent } from './components/create/create.component';

const routes: Routes = [
  { path: '', redirectTo: 'create' },
  { path: 'create', component: CreateComponent, },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DockerRoutingModule { }
