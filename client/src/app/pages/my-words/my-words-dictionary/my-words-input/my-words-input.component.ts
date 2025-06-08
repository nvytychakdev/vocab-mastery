import { ChangeDetectionStrategy, Component, output, signal } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucidePlus } from '@ng-icons/lucide';
import { matTranslate } from '@ng-icons/material-icons/baseline';
import { Button } from '@vm/ui';
import { Word } from '../../../../core/interfaces/word.interface';

@Component({
  selector: 'app-my-words-input',
  imports: [ReactiveFormsModule, NgIcon, Button],
  providers: [provideIcons({ lucidePlus, matTranslate })],
  templateUrl: './my-words-input.component.html',
  styleUrl: './my-words-input.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MyWordsInputComponent {
  readonly wordAdded = output<Word>();

  readonly word = new FormControl<string>('');
  readonly translations = signal(['one', 'two', 'three']);
  readonly translationsSelection = signal<string[]>([]);

  selectWord(word: string) {
    this.translationsSelection.update(value => {
      if (this.translationsSelection().includes(word)) {
        return [...value].filter(w => w !== word);
      }
      return [...value, word];
    });
  }

  isSelected(word: string) {
    return this.translationsSelection().includes(word);
  }

  addWord() {
    const text = this.word.value;
    if (!text) throw new Error('Word is required');

    const translation = this.translationsSelection();
    if (!translation.length) throw new Error('Translation was not provided');

    const word: Word = {
      text,
      translation,
      translationLang: 'en',
      definitions: {
        noun: ['A building for human habitation', 'A place where someone lives permanently'],
        verb: ['Provide (someone) with shelter or accommodation.'],
      },
    };

    this.wordAdded.emit(word);
    this.word.reset();
    this.translationsSelection.set([]);
  }
}
