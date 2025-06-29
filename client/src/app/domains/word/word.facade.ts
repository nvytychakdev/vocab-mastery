import { DestroyRef, inject, Injectable } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';
import { WordApi } from './word.api';
import { Word, WordBase } from './word.interface';
import { WordState } from './word.state';

@Injectable({ providedIn: 'root' })
export class WordFacade {
  private readonly destroyRef = inject(DestroyRef);
  private readonly state = inject(WordState);
  private readonly api = inject(WordApi);

  readonly words = this.state.listItems;
  readonly wordsChanges$ = this.state.listItemsChanges$;
  readonly wordsLoading = this.state.listItemsLoading;
  readonly wordsLoadingChanges$ = this.state.listItemsLoadingChanges$;

  create(dictionaryId: string, word: WordBase) {
    this.state.setItemsLoading(true);
    return this.api.create(word, { params: { dictionaryId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        const newWord: Word = {
          ...word,
          id: data.id,
          craetedAt: new Date().toISOString(),
        };
        this.state.addItems([newWord]);
        this.state.setItemsLoading(false);
      })
    );
  }

  loadAll(dictionaryId: string) {
    this.state.setItemsLoading(true);
    return this.api.getAllWithTranslations({ params: { dictionaryId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.setItems(data.items);
        this.state.setItemsLoading(false);
      })
    );
  }

  deleteById(dictionaryId: string, wordId: string) {
    this.state.setItemsLoading(true);
    return this.api.deleteById(wordId, { params: { dictionaryId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(() => {
        this.state.removeItemById(wordId);
        this.state.setItemsLoading(false);
      })
    );
  }
}
