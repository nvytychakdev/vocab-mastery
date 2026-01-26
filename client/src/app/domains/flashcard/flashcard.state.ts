import { Injectable, signal } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { FlashcardSessionState } from './flashcard.interface';

@Injectable({ providedIn: 'root' })
export class FlashcardState {
  private readonly _sessionState = signal<FlashcardSessionState | null>(null);
  readonly sessionState = this._sessionState.asReadonly();
  readonly sessionStateChange$ = toObservable(this.sessionState);

  private readonly _sessionLoading = signal<boolean>(false);
  readonly sessionLoading = this._sessionLoading.asReadonly();
  readonly sessionLoadingChange$ = toObservable(this.sessionLoading);

  private readonly _sessionAnswerLoading = signal<boolean>(false);
  readonly sessionAnswerLoading = this._sessionAnswerLoading.asReadonly();
  readonly sessionAnswerLoadingChange$ = toObservable(this.sessionAnswerLoading);

  setSessionLoading(isLoading: boolean) {
    this._sessionLoading.set(isLoading);
  }

  setSessionAnswerLoading(isLoading: boolean) {
    this._sessionAnswerLoading.set(isLoading);
  }

  setSessionState(item: FlashcardSessionState) {
    this._sessionState.set(item);
  }

  updateSessionState(item: Partial<FlashcardSessionState>) {
    this._sessionState.update(value => {
      if (!value) throw new Error('Not able to update session state which was not defined');
      return { ...value, ...item };
    });
  }

  removeSessionState() {
    this._sessionState.set(null);
  }
}
