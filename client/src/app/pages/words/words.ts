import {
  Component,
  computed,
  DestroyRef,
  effect,
  ElementRef,
  inject,
  OnInit,
  signal,
  viewChild,
  viewChildren,
} from '@angular/core';
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

  readonly letterControls = viewChild<unknown, ElementRef<HTMLDivElement>>('letterControls', {
    read: ElementRef<HTMLDivElement>,
  });
  readonly letterSections = viewChildren<unknown, ElementRef<HTMLDivElement>>('letterSection', {
    read: ElementRef<HTMLDivElement>,
  });

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

  constructor() {
    effect(() => {
      this.handleScrollLetters();
    });
  }

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

  private handleScrollLetters() {
    const observer = new IntersectionObserver(
      entries => {
        const visible = entries
          .filter(e => e.isIntersecting)
          .sort((a, b) => a.boundingClientRect.top - b.boundingClientRect.top);
        if (visible.length > 0) {
          const char = visible.at(0)?.target.getAttribute('data-letter');
          this.selectedChar.set(char || undefined);
        }
      },
      {
        threshold: [0, 1],
      }
    );

    this.letterSections().forEach(section => observer.observe(section.nativeElement));
  }

  onAlphabetSelection(char: string) {
    const isSelected = this.selectedChar() === char;
    this.selectedChar.set(isSelected ? undefined : char);
    if (!this.selectedChar()) return;
    this.scrollToLetterSection(char);
  }

  private scrollToLetterSection(char: string) {
    const controlsHeight = this.letterControls()?.nativeElement.getBoundingClientRect().height || 0;
    const block = this.letterSections().find(d => d.nativeElement.classList.contains(`letter-${char}`))?.nativeElement;
    if (!block) return;

    const blockSize = block.getBoundingClientRect();
    const blockY = blockSize.top + window.scrollY - (controlsHeight + blockSize.x);
    window.scrollTo({ behavior: 'smooth', top: blockY });
  }
}
