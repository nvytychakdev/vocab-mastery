import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { Api } from '@core/api/api';
import { ApiOptions, ApiUrlParams } from '@core/api/api.interface';
import { IsAuthorizedContext } from '@core/models/authorized.model';
import { Observable } from 'rxjs';
import {
  FlashcardSession,
  FlashcardSessionAnswer,
  FlashcardSessionAnswerRequest,
  FlashcardSessionRequest,
} from './flashcard.interface';

export enum FlashcardEndpoint {
  Start = 'api/v1/flashcards/sessions/start',
  Answer = 'api/v1/flashcards/sessions/:sessionId/answer',
}

@Injectable({ providedIn: 'root' })
export class FlashcardApi {
  private readonly api = inject(Api);
  protected http = inject(HttpClient);

  private getApiUrl(endpoint: FlashcardEndpoint, options?: ApiOptions<ApiUrlParams>) {
    return this.api.getUrl(endpoint, options);
  }

  sessionStart(body: Partial<FlashcardSessionRequest>): Observable<FlashcardSession> {
    const url = this.getApiUrl(FlashcardEndpoint.Start);
    return this.http.post<FlashcardSession>(url, body, { context: IsAuthorizedContext });
  }

  sessionAnswer(sessionId: string, body: FlashcardSessionAnswerRequest): Observable<FlashcardSessionAnswer> {
    const url = this.getApiUrl(FlashcardEndpoint.Answer, { params: { sessionId: sessionId } });
    return this.http.post<FlashcardSessionAnswer>(url, body, { context: IsAuthorizedContext });
  }
}
