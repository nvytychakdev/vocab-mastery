import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { BadgeComponent } from '@vm/ui';
import { Word } from '../../core/interfaces/word.interface';

type Dictionary = {
  name: string;
  description: string;
  languages: {
    from: string;
    to: string;
  };
  words: Word[];
};

@Component({
  selector: 'app-my-words',
  imports: [ReactiveFormsModule, BadgeComponent],
  templateUrl: './my-words.component.html',
  styleUrl: './my-words.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MyWordsComponent {
  private readonly router = inject(Router);
  private readonly activatedRoute = inject(ActivatedRoute);

  readonly dictionaries = signal<Dictionary[]>([
    { name: 'Test 1', description: 'Test description', languages: { from: 'us', to: 'ua' }, words: [] },
    { name: 'Test 2', description: 'Test description', languages: { from: 'us', to: 'ua' }, words: [] },
    { name: 'Test 3', description: 'Test description', languages: { from: 'us', to: 'ua' }, words: [] },
    { name: 'Test 4', description: 'Test description', languages: { from: 'us', to: 'ua' }, words: [] },
  ]);

  openDetails(dicationaryId: number) {
    void this.router.navigate([dicationaryId + 1], { relativeTo: this.activatedRoute });
  }
}
