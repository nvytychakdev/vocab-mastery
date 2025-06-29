import { Component } from '@angular/core';
import { DictionaryDetails } from '../../../features/dictionary/dictionary-details/dictionary-details';

@Component({
  selector: 'app-my-words-dictionary',
  imports: [DictionaryDetails],
  templateUrl: './my-words-dictionary.html',
  styleUrl: './my-words-dictionary.css',
})
export class MyWordsDictionary {}
