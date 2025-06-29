import { DestroyRef, inject, Injectable } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';
import { TranslationApi } from './translation.api';
import { Translation, TranslationBase } from './translation.interface';
import { TranslationState } from './translation.state';

@Injectable({ providedIn: 'root' })
export class WordFacade {
  private readonly destroyRef = inject(DestroyRef);
  private readonly state = inject(TranslationState);
  private readonly api = inject(TranslationApi);

  readonly translations = this.state.listItems;
  readonly translationsChanges$ = this.state.listItemsChanges$;
  readonly translationsLoading = this.state.listItemsLoading;
  readonly translationsLoadingChanges$ = this.state.listItemsLoadingChanges$;

  create(dictionaryId: string, wordId: string, translation: TranslationBase) {
    this.state.setItemsLoading(true);
    return this.api.create(translation, { params: { dictionaryId, wordId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        const newTranslation: Translation = {
          ...translation,
          id: data.id,
          craetedAt: new Date().toISOString(),
        };
        this.state.addItems([newTranslation]);
        this.state.setItemsLoading(false);
      })
    );
  }

  loadAll(dictionaryId: string, wordId: string) {
    this.state.setItemsLoading(true);
    return this.api.getAll({ params: { dictionaryId, wordId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.setItems(data.items);
        this.state.setItemsLoading(false);
      })
    );
  }

  deleteById(dictionaryId: string, wordId: string, translationId: string) {
    this.state.setItemsLoading(true);
    return this.api.deleteById(translationId, { params: { dictionaryId, wordId } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(() => {
        this.state.removeItemById(wordId);
        this.state.setItemsLoading(false);
      })
    );
  }
}
