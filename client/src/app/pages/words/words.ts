import { Component, computed, DestroyRef, inject, OnInit, signal } from '@angular/core';
import { takeUntilDestroyed, toObservable } from '@angular/core/rxjs-interop';
import { Field, form } from '@angular/forms/signals';
import { RouterLink } from '@angular/router';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { WordFacade } from '@domain/word/word.facade';
import { WordListItem } from '@domain/word/word.interface';
import { ALPHABET } from '@feature/dictionary/dictionary.model';
import { Divider, ToggleButton, ToggleButtonGroup } from '@vm/ui';
import { distinctUntilChanged, switchMap } from 'rxjs';

@Component({
  selector: 'app-words',
  imports: [ToggleButtonGroup, ToggleButton, RouterLink, RouterLink, Field, Divider],
  templateUrl: './words.html',
  styleUrl: './words.css',
})
export class Words implements OnInit {
  private readonly destroyRef = inject(DestroyRef);
  readonly words = inject(WordFacade);
  readonly dictionaries = inject(DictionaryFacade);

  readonly filtersModel = signal({ dictionaryId: '', letter: '' });
  readonly filtersForm = form(this.filtersModel);

  readonly dictionaryFilterChanges$ = toObservable(this.filtersForm.dictionaryId().value);

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
    this.dictionaries.loadAll().subscribe();
    this.dictionaryFilterChanges$
      .pipe(
        distinctUntilChanged(),
        switchMap(dictionaryId => this.words.loadAll(dictionaryId)),
        takeUntilDestroyed(this.destroyRef)
      )
      .subscribe();
  }

  onAlphabetSelection(char: string) {
    const isSelected = this.selectedChar() === char;
    this.selectedChar.set(isSelected ? undefined : char);
  }
}
