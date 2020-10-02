import { NgModule } from '@angular/core';
import { SharedModule } from '../shared/shared.module';
import { PropertiesComponent } from './components/properties/properties.component';
import { SettingsComponent } from './components/settings/settings.component';
import { ConfigurationRoutingModule } from './configuration-routing.module';
import { MainComponent } from './main.component';

@NgModule({
  declarations: [
    MainComponent,
    PropertiesComponent,
    SettingsComponent,
  ],
  imports: [
    ConfigurationRoutingModule,
    SharedModule,
  ]
})
export class ConfigurationModule { }
