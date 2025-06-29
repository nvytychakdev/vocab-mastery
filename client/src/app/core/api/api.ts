import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environment';
import { ApiOptions } from './api.interface';

@Injectable({
  providedIn: 'root',
})
export class Api {
  getUrl(endpoint: string, options?: ApiOptions<Record<string, string>>) {
    const url = `${environment.hostUrl}/${endpoint}`;
    if (!options) return url.trim();

    return this.getUrlWithParams(url, options.params).trim();
  }

  private getUrlWithParams(url: string, params?: Record<string, string>) {
    if (!params) {
      return url;
    }

    // replace formatted `:param` with actual `value`
    const urlWithParams = Object.entries(params).reduce(
      (u, [key, value]) => u.replaceAll(`:${key}`, encodeURIComponent(value)),
      url
    );

    return urlWithParams.trim();
  }
}
