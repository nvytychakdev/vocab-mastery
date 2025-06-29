import { Component, inject, output } from '@angular/core';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { WordBase } from '@domain/word/word.interface';
import { Button, Input } from '@vm/ui';

@Component({
  selector: 'app-word-form',
  imports: [Button, Input, ReactiveFormsModule],
  templateUrl: './word-form.html',
  styleUrl: './word-form.css',
})
export class WordForm {
  private readonly fb = inject(FormBuilder).nonNullable;
  readonly createWord = output<WordBase>();

  readonly form = this.fb.group({
    word: [''],
    language: [''],
  });

  submit() {
    this.createWord.emit(this.form.value as WordBase);
  }
}
