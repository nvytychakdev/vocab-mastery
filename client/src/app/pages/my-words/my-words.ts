import { Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { Button } from '@vm/ui';
import { DictionaryList } from '../../features/dictionary/dictionary-list/dictionary-list';

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
