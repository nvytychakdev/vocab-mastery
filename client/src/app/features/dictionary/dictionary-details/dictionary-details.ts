import { Component, input } from '@angular/core';
import { Dictionary } from '@domain/dictionary/dictionary.interface';

@Component({
  selector: 'app-dictionary-details',
  imports: [],
  templateUrl: './dictionary-details.html',
  styleUrl: './dictionary-details.css',
})
export class DictionaryDetails {
  readonly dictionary = input.required<Dictionary>();
}
