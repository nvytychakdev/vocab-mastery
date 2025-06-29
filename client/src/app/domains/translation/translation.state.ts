import { Injectable, signal } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { TranslationListItem } from './translation.interface';

@Injectable({ providedIn: 'root' })
export class TranslationState {
  private readonly _listItems = signal<TranslationListItem[]>([]);
  readonly listItems = this._listItems.asReadonly();
  readonly listItemsChanges$ = toObservable(this.listItems);

  private readonly _listItemsLoading = signal(false);
  readonly listItemsLoading = this._listItemsLoading.asReadonly();
  readonly listItemsLoadingChanges$ = toObservable(this._listItemsLoading);

  setItemsLoading(isLoading: boolean) {
    this._listItemsLoading.set(isLoading);
  }

  setItems(items: TranslationListItem[]) {
    this._listItems.set(items);
  }

  addItems(items: TranslationListItem[]) {
    this._listItems.update(value => [...value, ...items]);
  }

  removeItemById(id: string) {
    this._listItems.update(value => value.filter(v => v.id !== id));
  }

  removeItems() {
    this._listItems.set([]);
  }
}
