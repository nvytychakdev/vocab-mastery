import { DestroyRef, inject, Injectable } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { tap } from 'rxjs';
import { FlashcardApi } from './flashcard.api';
import { FlashcardSessionAnswerRequest, FlashcardSessionRequest } from './flashcard.interface';
import { FlashcardState } from './flashcard.state';

/**
 * sessionStart() ->
 *    startLoading()
 *    setSessionState() -> stats, currentCard, nextCard, historyCards(null x cardsTotal)
 * sessionAnswerSubmit() ->
 *    startAnswerLoading()
 *    setSessionState() -> nextCard
 * nextQuestion() ->
 *    setSessionState() -> stats, nextCard
 */
@Injectable({ providedIn: 'root' })
export class FlashcardFacade {
  private readonly destroyRef = inject(DestroyRef);
  private readonly state = inject(FlashcardState);
  private readonly api = inject(FlashcardApi);

  readonly session = this.state.sessionState;

  sessionStart(body: Partial<FlashcardSessionRequest>) {
    this.state.setSessionLoading(true);
    return this.api.sessionStart(body).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.setSessionState({
          sessionId: data.sessionId,
          cardsTotal: data.cardsTotal,
          cardsAnswered: data.cardsAnswered,
          cardsCorrect: data.cardsCorrect,
          isCompleted: false,
          isAnswered: false,
          currentCard: data.nextCard,
          nextCard: null,
          answerResult: null,
          historyCards: [],
        });
        this.state.setSessionLoading(false);
      })
    );
  }

  sessionAnswerSubmit(sessionId: string, body: FlashcardSessionAnswerRequest) {
    this.state.setSessionAnswerLoading(true);
    return this.api.sessionAnswer(sessionId, body).pipe(
      takeUntilDestroyed(this.destroyRef),
      tap(data => {
        this.state.updateSessionState({
          cardsAnswered: data.cardsAnswered,
          cardsCorrect: data.cardsCorrect,
          nextCard: data.nextCard,
          isCompleted: data.isCompleted,
          answerResult: data.result,
          isAnswered: true,
        });
        this.state.setSessionAnswerLoading(false);
      })
    );
  }

  nextQuestion() {
    const state = this.session();
    if (!state?.nextCard) {
      this.state.updateSessionState({
        ...state,
        isCompleted: true,
      });
      return;
    }

    this.state.updateSessionState({
      ...state,
      currentCard: state.nextCard,
      isAnswered: false,
      answerResult: null,
      nextCard: null,
    });
  }
}
