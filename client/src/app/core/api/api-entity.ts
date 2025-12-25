import { HttpClient, HttpParams } from '@angular/common/http';
import { inject } from '@angular/core';
import { IsAuthorizedContext } from '@core/models/authorized.model';
import { Observable } from 'rxjs';
import {
  ApiListQueryParams,
  ApiOptions,
  ApiUrlParams,
  ResponseCreate,
  ResponseDelete,
  ResponseList,
} from './api.interface';

export abstract class ApiEntity<
  Entity extends Record<string, unknown>,
  EntityListItem extends Record<string, unknown>,
  EntityParams extends ApiUrlParams = ApiUrlParams,
> {
  protected http = inject(HttpClient);

  protected abstract getEntityUrl(options?: ApiOptions<EntityParams>): string;

  protected getEntityHttpParams(options?: ApiOptions<EntityParams>): HttpParams {
    const params = new HttpParams();
    if (!options?.query) return params;
    return Object.entries(options.query).reduce((prm, [key, value]) => {
      if (!value) return prm;
      if (Array.isArray(value)) return value.reduce((p, v) => p.append(key, v), prm);
      return prm.set(key, value);
    }, params);
  }

  create(body: Partial<Entity>, options?: ApiOptions<EntityParams>): Observable<ResponseCreate> {
    const url = this.getEntityUrl(options);
    return this.http.post<ResponseCreate>(url, body, { context: IsAuthorizedContext });
  }

  getAll<EL extends Record<string, unknown> = EntityListItem>(
    options?: ApiOptions<EntityParams, ApiListQueryParams>
  ): Observable<ResponseList<EL>> {
    const url = this.getEntityUrl(options);
    const params = this.getEntityHttpParams(options);
    return this.http.get<ResponseList<EL>>(url, { params, context: IsAuthorizedContext });
  }

  getById<E extends Record<string, unknown> = Entity>(id: string, options?: ApiOptions<EntityParams>): Observable<E> {
    const url = this.getEntityUrl(options);
    const params = this.getEntityHttpParams(options);
    return this.http.get<E>(`${url}/${id}`, { params, context: IsAuthorizedContext });
  }

  deleteById(id: string, options?: ApiOptions<EntityParams>): Observable<ResponseDelete> {
    const url = this.getEntityUrl(options);
    return this.http.delete<ResponseDelete>(`${url}/${id}`, { context: IsAuthorizedContext });
  }
}
