import {Component, OnInit} from '@angular/core';

@Component({
  selector: 'app-examples',
  templateUrl: './examples.component.html',
  styleUrls: ['./examples.component.scss']
})
export class ExamplesComponent implements OnInit {
  dashboard = [
    {cols: 2, rows: 2, y: 3, x: 5, minItemRows: 2, minItemCols: 2, label: 'Min rows & cols = 2'},
    {cols: 2, rows: 2, y: 2, x: 0, maxItemRows: 2, maxItemCols: 2, label: 'Max rows & cols = 2'},
    // {cols: 2, rows: 1, y: 2, x: 2, dragEnabled: true, resizeEnabled: true, label: 'Drag&Resize Enabled'},
    // {cols: 1, rows: 1, y: 2, x: 4, dragEnabled: false, resizeEnabled: false, label: 'Drag&Resize Disabled'},
    {
      cols: 4, rows: 4, y: 1, x: 0, type: 'table', attr: {
        rows: [
          {n: 1, name: 'Danil Akhtarov', gender: 'Male', result: '00:44:30'},
          {n: 2, name: 'Vity Korin', gender: 'Male', result: '00:44:45'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 3, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
          {n: 4, name: 'Molly Sveen', gender: 'Male', result: '00:44:53'},
        ], columns: [
          {prop: 'n'},
          {prop: 'name'},
          {name: 'Gender'},
          {name: 'Result'}
        ]
      }
    },
  ];

  constructor() { }

  ngOnInit() {
  }

}
