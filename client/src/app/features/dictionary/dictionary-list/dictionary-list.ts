import { Component, input, output } from '@angular/core';
import { DictionaryListItem } from '@domain/dictionary/dictionary.interface';
import { Card } from '@vm/ui';

@Component({
  selector: 'app-dictionary-list',
  imports: [Card],
  templateUrl: './dictionary-list.html',
  styleUrl: './dictionary-list.css',
})
export class DictionaryList {
  readonly dictionaries = input.required<DictionaryListItem[]>();
  readonly dictionarySelect = output<string>();
}
