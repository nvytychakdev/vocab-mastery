import { Component, inject, OnInit } from '@angular/core';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { WordFacade } from '@domain/word/word.facade';
import { WordBase } from '@domain/word/word.interface';
import { DictionaryDetails } from '@feature/dictionary/dictionary-details/dictionary-details';
import { WordForm } from '@feature/word/word-form/word-form';
import { ToastService } from '@vm/ui';

@Component({
  selector: 'app-my-words-dictionary',
  imports: [DictionaryDetails, WordForm],
  templateUrl: './my-words-dictionary.html',
  styleUrl: './my-words-dictionary.css',
})
export class MyWordsDictionary implements OnInit {
  readonly dictionaryFacade = inject(DictionaryFacade);
  readonly wordFacade = inject(WordFacade);
  private readonly toast = inject(ToastService);

  ngOnInit() {
    const dictionary = this.dictionaryFacade.activeDictionary();
    if (!dictionary?.id) return;
    this.wordFacade.loadAll(dictionary.id).subscribe();
  }

  onCreateWord(word: WordBase) {
    const dictionary = this.dictionaryFacade.activeDictionary();
    if (!dictionary?.id) return;

    return this.wordFacade.create(dictionary.id, word).subscribe(() => {
      this.toast.success('Success', 'Word created succesfully');
    });
  }
}
