import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { DictionaryList } from '@feature/dictionary/dictionary-list/dictionary-list';
import { Button } from '@vm/ui';

@Component({
  selector: 'app-my-words',
  imports: [DictionaryList, Button],
  templateUrl: './my-words.html',
  styleUrl: './my-words.css',
})
export class MyWords {
  private readonly router = inject(Router);

  addDictionary() {
    void this.router.navigate(['/main/my-words/new']);
  }
}
