import { inject, Injectable } from '@angular/core';
import { Api } from '../../core/api/api';
import { ApiEntity } from '../../core/api/api-entity';
import { ApiOptions, ApiUrlParams } from '../../core/api/api.interface';
import { Dictionary, DictionaryListItem } from './dictionary.interface';

export enum DictionaryEndpoint {
  Dictionaries = 'dictionaries',
}

@Injectable({
  providedIn: 'root',
})
export class DictionariesApi extends ApiEntity<Dictionary, DictionaryListItem> {
  private readonly api = inject(Api);

  private getApiUrl(endpoint: DictionaryEndpoint, options?: ApiOptions<ApiUrlParams>) {
    return this.api.getUrl(endpoint, options?.params);
  }

  protected override getEntityUrl(options?: ApiOptions<ApiUrlParams>): string {
    return this.getApiUrl(DictionaryEndpoint.Dictionaries, options);
  }
}
