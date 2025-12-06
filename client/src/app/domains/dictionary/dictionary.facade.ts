import { DestroyRef, inject, Injectable } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';
import { DictionaryApi } from './dictionary.api';
import { Dictionary, DictionaryBase } from './dictionary.interface';
import { DictionaryState } from './dictionary.state';

@Injectable({ providedIn: 'root' })
export class DictionaryFacade {
  private readonly destroyRef = inject(DestroyRef);
  private readonly state = inject(DictionaryState);
  private readonly api = inject(DictionaryApi);

  readonly dictionaries = this.state.listItems;
  readonly dictionariesChanges$ = this.state.listItemsChanges$;
  readonly dictionariesLoading = this.state.listItemsLoading;
  readonly dictionariesLoadingChanges$ = this.state.listItemsLoadingChanges$;
  readonly activeDictionary = this.state.activeItem;
  readonly activeDictionaryChanges$ = this.state.activeItemChanges$;

  create(dictionary: DictionaryBase) {
    this.state.setItemsLoading(true);
    return this.api.create(dictionary).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        const newDictionary: Dictionary = {
          ...dictionary,
          id: data.id,
          createdAt: new Date().toISOString(),
        };
        this.state.addItems([newDictionary]);
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

  loadAll() {
    this.state.setItemsLoading(true);
    return this.api.getAll().pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.setItems(data.items);
        this.state.setItemsLoading(false);
      })
    );
  }

  deleteById(id: string) {
    this.state.setItemsLoading(true);
    return this.api.deleteById(id).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(() => {
        this.state.removeItemById(id);
        this.state.setItemsLoading(false);
      })
    );
  }
}
