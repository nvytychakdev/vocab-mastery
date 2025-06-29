import { inject, Injectable } from '@angular/core';
import { Api } from '@core/api/api';
import { ApiEntity } from '@core/api/api-entity';
import { ApiListQueryParams, ApiOptions, ApiUrlParams } from '@core/api/api.interface';
import {
  WithTranslations,
  Word,
  WordListItem,
  WordListItemWithTranslation,
  WordWithTranslation,
} from './word.interface';

enum WordEndpoint {
  Words = 'api/v1/dictionaries/:dictionaryId/words',
}

@Injectable({
  providedIn: 'root',
})
export class WordApi extends ApiEntity<Word, WordListItem> {
  private readonly api = inject(Api);

  private getApiUrl(endpoint: WordEndpoint, options?: ApiOptions<ApiUrlParams>) {
    return this.api.getUrl(endpoint, options);
  }

  protected override getEntityUrl(options?: ApiOptions<ApiUrlParams>): string {
    return this.getApiUrl(WordEndpoint.Words, options);
  }

  getByIdWithTranslations(id: string) {
    return this.getById<WithTranslations<WordWithTranslation>>(id, { query: { include: ['translations'] } });
  }

  getAllWithTranslations(options?: ApiOptions<ApiUrlParams, ApiListQueryParams>) {
    return this.getAll<WithTranslations<WordListItemWithTranslation>>({
      params: options?.params,
      query: { ...options?.query, include: ['translations'] },
    });
  }
}
