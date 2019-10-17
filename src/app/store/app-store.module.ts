import { NgModule } from '@angular/core';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';
import { environment } from '../../environments/environment';
import { DefaultDataServiceConfig, NgrxDataModule } from 'ngrx-data';
import { entityConfig } from './entity-metadata';

const defaultDataServiceConfig: DefaultDataServiceConfig = {
  root: 'http://localhost:4201/api/',
  timeout: 3000, // request timeout
}

@NgModule({
  declarations: [],
  imports: [
    StoreModule.forRoot({}),
    EffectsModule.forRoot([]),
    NgrxDataModule.forRoot(entityConfig),
    environment.production ? [] : StoreDevtoolsModule.instrument()
  ],
  providers: [
    {
      provide: DefaultDataServiceConfig, useValue: defaultDataServiceConfig
    }
  ]
})
export class AppStoreModule { }