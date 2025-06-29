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

  readonly dictionaries = this.state.listItems;
  readonly dictionariesChanges$ = this.state.listItemsChanges$;
  readonly dictionariesLoading = this.state.listItemsLoading;
  readonly dictionariesLoadingChanges$ = this.state.listItemsLoadingChanges$;

  create(dictionaryId: string, word: WordBase) {
    this.state.setItemsLoading(true);
    return this.api.create(word, { params: { dictionaryId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        const newDictionary: Word = {
          ...word,
          id: data.id,
          craetedAt: new Date().toISOString(),
        };
        this.state.addItems([newDictionary]);
        this.state.setItemsLoading(false);
      })
    );
  }

  loadAll(dictionaryId: string) {
    this.state.setItemsLoading(true);
    return this.api.getAll({ params: { dictionaryId } }).pipe(
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
