import { inject } from '@angular/core';
import { ActivatedRouteSnapshot, ResolveFn } from '@angular/router';
import { WordFacade } from '@domain/word/word.facade';
import { Word } from '@domain/word/word.interface';
import { Observable, of } from 'rxjs';

export const wordByIdResolver: ResolveFn<Word | null> = (route: ActivatedRouteSnapshot): Observable<Word | null> => {
  const wordId = route.paramMap.get('wordId');
  if (!wordId) return of(null);

  const words = inject(WordFacade);
  return words.loadActive(wordId);
};
