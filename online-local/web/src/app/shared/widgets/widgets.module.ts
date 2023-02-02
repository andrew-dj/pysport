import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {GridsterModule} from 'angular-gridster2';
import {MatTableModule} from '@angular/material/table';
import {MatGridListModule} from '@angular/material/grid-list';
import {DragDropModule} from '@angular/cdk/drag-drop';
import { NgxDatatableModule } from '@swimlane/ngx-datatable';

import {WidgetsComponent} from './widgets.component';
import {TableComponent} from './components/table/table.component';
import {MapComponent} from './components/map/map.component';
import {InfoComponent} from './components/info/info.component';

@NgModule({
  declarations: [WidgetsComponent, TableComponent, MapComponent, InfoComponent],
  imports: [
    CommonModule,
    MatGridListModule,
    DragDropModule,
    GridsterModule,
    MatTableModule,
    NgxDatatableModule,
  ],
  exports: [
    WidgetsComponent
  ]
})
export class WidgetsModule {}
