import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn } from '@angular/router';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { Dictionary, DictionaryListItem } from '@domain/dictionary/dictionary.interface';
import { map, Observable, of } from 'rxjs';

export const dictionariesResolver: ResolveFn<DictionaryListItem[]> = () => {
  const dictionaries = inject(DictionaryFacade);
  return dictionaries.loadAll().pipe(map(data => data.items));
};

export const dictionaryByIdResolver: ResolveFn<Dictionary | null> = (
  route: ActivatedRouteSnapshot
): Observable<Dictionary | null> => {
  const dictionaryId = route.paramMap.get('dictionaryId');
  if (!dictionaryId) return of(null);

  const dictionaries = inject(DictionaryFacade);
  return dictionaries.loadActive(dictionaryId);
};
