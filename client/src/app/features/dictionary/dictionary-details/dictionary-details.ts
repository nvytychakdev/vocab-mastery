import { Component, input } from '@angular/core';
import { Dictionary } from '@domain/dictionary/dictionary.interface';
import { WordListItem } from '@domain/word/word.interface';
import { WordPreview } from '@feature/word/word-preview/word-preview';

@Component({
  selector: 'app-dictionary-details',
  imports: [WordPreview],
  templateUrl: './dictionary-details.html',
  styleUrl: './dictionary-details.css',
})
export class DictionaryDetails {
  readonly dictionary = input.required<Dictionary>();
  readonly words = input.required<WordListItem[]>();
}
