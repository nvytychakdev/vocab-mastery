import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { FormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { lucidePlus } from '@ng-icons/lucide';

type Dictionary = Array<{
  original?: string;
  translated?: string;
}>;

@Component({
  selector: 'app-my-words',
  imports: [ReactiveFormsModule, NgIcon],
  providers: [provideIcons({ lucidePlus })],
  templateUrl: './my-words.component.html',
  styleUrl: './my-words.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class MyWordsComponent {
  readonly fb = inject(FormBuilder);

  readonly form = this.fb.group({
    original: this.fb.control('', [Validators.required]),
    translated: this.fb.control('', [Validators.required]),
  });

  readonly translations = signal<string[]>(['One', 'Two', 'Three']);
  readonly selectedTranslations = signal<string[]>([]);

  readonly dictionary = signal<Dictionary>([]);

  selectWord(word: string) {
    this.selectedTranslations.update(value => {
      if (this.selectedTranslations().includes(word)) {
        return [...value].filter(w => w !== word);
      }
      return [...value, word];
    });

    console.log(this.selectedTranslations());
  }

  isSelected(word: string) {
    return this.selectedTranslations().includes(word);
  }

  addWord() {
    const original = this.form.value.original;
    const translated = this.selectedTranslations().join(', ');

    if (!original || !translated) {
      throw new Error('Words are required');
    }

    this.dictionary.update(value => {
      return [{ original, translated }, ...value];
    });

    this.form.controls.original.setValue('');
    this.selectedTranslations.set([]);
  }
}
