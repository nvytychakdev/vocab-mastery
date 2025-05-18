import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { provideIcons } from '@ng-icons/core';
import { lucidePlus } from '@ng-icons/lucide';
import { matTranslate } from '@ng-icons/material-icons/baseline';
import { Word } from '../../../core/interfaces/word.interface';
import { MyWordsCardComponent } from './my-words-card/my-words-card.component';
import { MyWordsInputComponent } from './my-words-input/my-words-input.component';

type Dictionary = Word[];

@Component({
  selector: 'app-my-words-dictionary',
  imports: [ReactiveFormsModule, MyWordsCardComponent, MyWordsInputComponent],
  providers: [provideIcons({ lucidePlus, matTranslate })],
  templateUrl: './my-words-dictionary.component.html',
  styleUrl: './my-words-dictionary.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MyWordsDictionaryComponent {
  readonly dictionary = signal<Dictionary>([]);

  addWord(word: Word) {
    this.dictionary.update(value => {
      return [word, ...value];
    });
  }
}
