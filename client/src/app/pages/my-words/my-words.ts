import { Component, inject, signal } from '@angular/core';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { Dictionary } from '@domain/dictionary/dictionary.interface';

@Component({
  selector: 'app-my-words',
  imports: [],
  templateUrl: './my-words.html',
  styleUrl: './my-words.css',
})
export class MyWords {
  readonly dictionaryFacade = inject(DictionaryFacade);
  readonly selectedDictionary = signal<Dictionary | null>(null);

  onDictionarySelection(dictionary: Dictionary) {
    if (this.selectedDictionary()) {
      this.selectedDictionary.set(null);
    } else {
      this.selectedDictionary.set(dictionary);
    }
  }
}
