import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { DictionaryFacade } from '@domain/dictionary/dictionary.facade';
import { DictionaryBase } from '@domain/dictionary/dictionary.interface';
import { DictionaryForm } from '@feature/dictionary/dictionary-form/dictionary-form';
import { ToastService } from '@vm/ui';

@Component({
  selector: 'app-my-words-dictionary-new',
  imports: [DictionaryForm],
  templateUrl: './my-words-dictionary-new.html',
  styleUrl: './my-words-dictionary-new.css',
})
export class MyWordsDictionaryNew {
  private readonly dictionary = inject(DictionaryFacade);
  private readonly toast = inject(ToastService);
  private readonly router = inject(Router);

  onCreateDictionary(dictionary: DictionaryBase) {
    return this.dictionary.create(dictionary).subscribe(({ id }) => {
      this.toast.success('Success', 'Dictionary created succesfully');
      void this.router.navigate(['/main/my-words', id]);
    });
  }
}
