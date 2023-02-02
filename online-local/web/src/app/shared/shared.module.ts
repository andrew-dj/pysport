import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {ToolbarModule} from './toolbar/toolbar.module';
import {WidgetsModule} from './widgets/widgets.module';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    ToolbarModule,
    WidgetsModule
  ],
  exports: [
    ToolbarModule,
    WidgetsModule,
  ]
})
export class SharedModule {}
