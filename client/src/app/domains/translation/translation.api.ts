import { inject, Injectable } from '@angular/core';
import { Api } from '../../core/api/api';
import { ApiEntity } from '../../core/api/api-entity';
import { ApiOptions, ApiUrlParams } from '../../core/api/api.interface';
import { Translation, TranslationListItem } from './translation.interface';

enum TranslationEndpoint {
  Translations = 'translations',
}

@Injectable({
  providedIn: 'root',
})
export class TranslationApi extends ApiEntity<Translation, TranslationListItem> {
  private readonly api = inject(Api);

  private getApiUrl(endpoint: TranslationEndpoint, options?: ApiOptions<ApiUrlParams>) {
    return this.api.getUrl(endpoint, options);
  }

  protected override getEntityUrl(params?: ApiOptions<ApiUrlParams>): string {
    return this.getApiUrl(TranslationEndpoint.Translations, params);
  }
}
