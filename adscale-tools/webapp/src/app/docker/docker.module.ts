import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';
import { CreateComponent } from './components/create/create.component';
import { DockerRoutingModule } from './docker-routing.module';
import { MainComponent } from './main.component';

@NgModule({
  declarations: [MainComponent, CreateComponent],
  imports: [
    DockerRoutingModule,
    SharedModule,
  ]
})
export class DockerModule { }
