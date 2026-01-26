import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { FlashcardFacade } from '@domain/flashcard/flashcard.facade';
import { BadgeComponent, Button } from '@vm/ui';

@Component({
  selector: 'app-play',
  imports: [Button, BadgeComponent],
  templateUrl: './play.component.html',
  styleUrl: './play.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PlayComponent {
  readonly flashcard = inject(FlashcardFacade);

  startSession() {
    this.flashcard.sessionStart({}).subscribe();
  }

  nextQuestion() {
    this.flashcard.nextQuestion();
  }

  submitSessionAnswer(answerId: string) {
    const session = this.flashcard.session();
    if (session?.isAnswered) return;

    const meaningId = session?.currentCard?.meaningId;
    if (!meaningId) throw new Error('Not able to submit answer without meaningId');

    this.flashcard
      .sessionAnswerSubmit(session.sessionId, { meaningId: meaningId, selectedAnswer: answerId })
      .subscribe();
  }
}
