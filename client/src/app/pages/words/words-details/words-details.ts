import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { RouterLink } from '@angular/router';
import { WordFacade } from '@domain/word/word.facade';
import { WordPartOfSpeech } from '@domain/word/word.interface';
import { WORDS_PART_OF_SPEECH_COLOR_MAP } from '@domain/word/word.model';
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

  getPartOfSpeechColor(partOfSpeech: WordPartOfSpeech) {
    return WORDS_PART_OF_SPEECH_COLOR_MAP[partOfSpeech];
  }
}
