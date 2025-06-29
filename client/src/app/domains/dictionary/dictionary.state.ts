import { Injectable, signal } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { Dictionary, DictionaryListItem } from './dictionary.interface';

@Injectable({
  providedIn: 'root',
})
export class DictionaryState {
  private readonly _listItems = signal<DictionaryListItem[]>([]);
  readonly listItems = this._listItems.asReadonly();
  readonly listItemsChanges$ = toObservable(this.listItems);

  private readonly _listItemsLoading = signal(false);
  readonly listItemsLoading = this._listItemsLoading.asReadonly();
  readonly listItemsLoadingChanges$ = toObservable(this._listItemsLoading);

  private readonly _activeItem = signal<Dictionary | null>(null);
  readonly activeItem = this._activeItem.asReadonly();
  readonly activeItemChanges$ = toObservable(this.activeItem);

  setActiveItem(item: Dictionary) {
    this._activeItem.set(item);
  }

  removeActiveItem() {
    this._activeItem.set(null);
  }

  setItemsLoading(isLoading: boolean) {
    this._listItemsLoading.set(isLoading);
  }

  setItems(items: DictionaryListItem[]) {
    this._listItems.set(items);
  }

  addItems(items: DictionaryListItem[]) {
    this._listItems.update(value => [...value, ...items]);
  }

  removeItemById(id: string) {
    this._listItems.update(value => value.filter(v => v.id !== id));
  }

  removeItems() {
    this._listItems.set([]);
  }
}
