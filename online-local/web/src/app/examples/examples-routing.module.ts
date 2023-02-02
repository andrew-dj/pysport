import { NgModule } from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {ExamplesComponent} from './examples.component';

const routes: Routes = [
  {
    path: 'examples',
    component: ExamplesComponent
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers: []
})
export class ExamplesRoutingModule { }
