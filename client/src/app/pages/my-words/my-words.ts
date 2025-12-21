import { Component, signal } from '@angular/core';
import { ALPHABET, ENGLISH_LEVEL, EnglishLevel } from '@feature/dictionary/dictionary.model';
import { ToggleButton, ToggleButtonGroup } from '@vm/ui';

@Component({
  selector: 'app-my-words',
  imports: [ToggleButtonGroup, ToggleButton],
  templateUrl: './my-words.html',
  styleUrl: './my-words.css',
})
export class MyWords {
  readonly englishLevel = Object.values(ENGLISH_LEVEL);
  readonly selectedEnglishLevel = signal<EnglishLevel | undefined>(undefined);
  readonly alphabet = ALPHABET;
  readonly selectedChar = signal<string | undefined>(undefined);

  readonly words = ['word', 'word', 'word', 'word', 'word'];
  readonly groups = this.alphabet.map(letter => ({ letter, words: this.words }));

  onLevelSelection(level: string | number | undefined) {
    this.selectedEnglishLevel.set(level as EnglishLevel);
  }

  onAlphabetSelection(char: string) {
    const isSelected = this.selectedChar() === char;
    this.selectedChar.set(isSelected ? undefined : char);
  }
}
