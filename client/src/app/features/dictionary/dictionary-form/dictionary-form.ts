import { Component, inject, output } from '@angular/core';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';
import { DictionaryBase } from '@domain/dictionary/dictionary.interface';
import { Button, Input } from '@vm/ui';

@Component({
  selector: 'app-dictionary-form',
  imports: [Input, Button, ReactiveFormsModule],
  templateUrl: './dictionary-form.html',
  styleUrl: './dictionary-form.css',
})
export class DictionaryForm {
  private readonly fb = inject(FormBuilder).nonNullable;
  readonly createDictonary = output<DictionaryBase>();

  readonly form = this.fb.group({
    name: [''],
    description: [''],
  });
  submit() {
    this.createDictonary.emit(this.form.value as DictionaryBase);
  }
}
