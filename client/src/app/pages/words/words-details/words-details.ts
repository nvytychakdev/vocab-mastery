import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { RouterLink } from '@angular/router';
import { WordFacade } from '@domain/word/word.facade';
import { BadgeComponent, Button } from '@vm/ui';

@Component({
  selector: 'app-words-details',
  imports: [BadgeComponent, RouterLink, Button],
  host: {
    class: 'flex justify-center',
  },
  templateUrl: './words-details.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class WordsDetails {
  readonly words = inject(WordFacade);
}
