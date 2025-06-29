import { Component, inject } from '@angular/core';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { DictionaryDetails } from '@feature/dictionary/dictionary-details/dictionary-details';

@Component({
  selector: 'app-my-words-dictionary',
  imports: [DictionaryDetails],
  templateUrl: './my-words-dictionary.html',
  styleUrl: './my-words-dictionary.css',
})
export class MyWordsDictionary {
  readonly facade = inject(DictionaryFacade);
}
