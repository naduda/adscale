import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    children: [
      { path: '', redirectTo: 'config', pathMatch: 'full' },
      { path: 'config', loadChildren: () => import('./configuration/configuration.module').then(m => m.ConfigurationModule) },
      { path: 'docker', loadChildren: () => import('./docker/docker.module').then(m => m.DockerModule) },
    ],
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { useHash: true })],
  exports: [RouterModule]
})
export class AppRoutingModule { }
