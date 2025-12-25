import { DestroyRef, inject, Injectable } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';
import { WordApi } from './word.api';
import { WORDS_LIMIT, WORDS_SORT_BY, WORDS_SORT_DIR } from './word.model';
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
  readonly activeWord = this.state.activeItem;
  readonly activeWordChanges$ = this.state.activeItemChanges$;

  loadAll() {
    this.state.setItemsLoading(true);
    return this.api.getAll({ query: { limit: WORDS_LIMIT, sortBy: WORDS_SORT_BY, dir: WORDS_SORT_DIR } }).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.setItems(data.items);
        this.state.setItemsLoading(false);
      })
    );
  }

  loadActive(id: string) {
    return this.api.getById(id).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.setActiveItem(data);
      })
    );
  }
}
