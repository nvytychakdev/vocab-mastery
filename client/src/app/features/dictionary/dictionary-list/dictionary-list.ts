import { Component, output } from '@angular/core';

@Component({
  selector: 'app-dictionary-list',
  imports: [],
  templateUrl: './dictionary-list.html',
  styleUrl: './dictionary-list.css',
})
export class DictionaryList {
  readonly dictonarySelect = output<string>();
}
