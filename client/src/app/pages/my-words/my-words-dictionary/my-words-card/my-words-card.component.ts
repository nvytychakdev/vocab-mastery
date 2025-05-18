import { ChangeDetectionStrategy, Component, input, signal } from '@angular/core';
import { Word } from '../../../../core/interfaces/word.interface';

@Component({
  selector: 'app-my-words-card',
  imports: [],
  templateUrl: './my-words-card.component.html',
  styleUrl: './my-words-card.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MyWordsCardComponent {
  readonly word = input.required<Word>();
  readonly collapsed = signal(true);
}
