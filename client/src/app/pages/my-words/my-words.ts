import { Component, computed, inject, OnInit, signal } from '@angular/core';
import { WordFacade } from '@domain/word/word.facade';
import { WordListItem } from '@domain/word/word.interface';
import { ALPHABET, ENGLISH_LEVEL, EnglishLevel } from '@feature/dictionary/dictionary.model';
import { ToggleButton, ToggleButtonGroup } from '@vm/ui';

@Component({
  selector: 'app-my-words',
  imports: [ToggleButtonGroup, ToggleButton],
  templateUrl: './my-words.html',
  styleUrl: './my-words.css',
})
export class MyWords implements OnInit {
  readonly words = inject(WordFacade);
  readonly englishLevel = Object.values(ENGLISH_LEVEL);
  readonly selectedEnglishLevel = signal<EnglishLevel | undefined>(undefined);
  readonly alphabet = ALPHABET;
  readonly selectedChar = signal<string | undefined>(undefined);

  readonly groups = computed(() => {
    const groupedWords = this.words.words().reduce<Record<string, WordListItem[]>>((acc, data) => {
      const firstLetter = data.word.trim().at(0)?.toUpperCase();
      if (!firstLetter) return acc;
      if (!acc[firstLetter]) {
        acc[firstLetter] = [data];
      } else {
        acc[firstLetter].push(data);
      }
      return acc;
    }, {});

    return Object.entries(groupedWords).map(([key, value]) => ({ letter: key, words: value }));
  });

  ngOnInit() {
    // TODO: move to resolver
    this.words.loadAll().subscribe();
  }

  onLevelSelection(level: string | number | undefined) {
    this.selectedEnglishLevel.set(level as EnglishLevel);
  }

  onAlphabetSelection(char: string) {
    const isSelected = this.selectedChar() === char;
    this.selectedChar.set(isSelected ? undefined : char);
  }
}
