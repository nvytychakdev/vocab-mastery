import { Component, input } from '@angular/core';
import { WordListItemWithTranslation } from '@domain/word/word.interface';

@Component({
  selector: 'app-word-preview',
  imports: [],
  templateUrl: './word-preview.html',
  styleUrl: './word-preview.css',
})
export class WordPreview {
  readonly word = input.required<WordListItemWithTranslation>();
}
