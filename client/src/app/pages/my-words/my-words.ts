import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { DictionaryList } from '@feature/dictionary/dictionary-list/dictionary-list';
import { Button } from '@vm/ui';

@Component({
  selector: 'app-my-words',
  imports: [DictionaryList, Button],
  templateUrl: './my-words.html',
  styleUrl: './my-words.css',
})
export class MyWords {
  private readonly router = inject(Router);
  readonly facade = inject(DictionaryFacade);

  addDictionary() {
    void this.router.navigate(['/main/my-words/new']);
  }

  onDictionarySelect(id: string) {
    void this.router.navigate(['/main/my-words', id]);
  }
}
