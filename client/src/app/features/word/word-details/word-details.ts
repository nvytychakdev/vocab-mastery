import { Component, input } from '@angular/core';
import { Word } from '@domain/word/word.interface';

@Component({
  selector: 'app-word-details',
  imports: [],
  templateUrl: './word-details.html',
  styleUrl: './word-details.css',
})
export class WordDetails {
  readonly word = input.required<Word>();
}
