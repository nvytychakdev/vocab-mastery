import { Component, signal } from '@angular/core';
import { ALPHABET, ENGLISH_LEVEL, EnglishLevel } from '@feature/dictionary/dictionary.model';

@Component({
  selector: 'app-my-words',
  imports: [],
  templateUrl: './my-words.html',
  styleUrl: './my-words.css',
})
export class MyWords {
  readonly englishLevel = Object.values(ENGLISH_LEVEL);
  readonly selectedEnglishLevel = signal<EnglishLevel | undefined>(undefined);
  readonly alphabet = ALPHABET;
  readonly selectedChar = signal<string | undefined>(undefined);

  readonly words = ['word', 'word', 'word'];
  readonly groups = [this.words, this.words, this.words];

  onLevelSelection(level: EnglishLevel) {
    this.selectedEnglishLevel.set(level);
  }

  onAlphabetSelection(char: string) {
    const isSelected = this.selectedChar() === char;
    this.selectedChar.set(isSelected ? undefined : char);
  }
}
