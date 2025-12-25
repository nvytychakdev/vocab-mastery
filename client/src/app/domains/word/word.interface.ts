import { Entity } from '@core/interfaces/dto.interface';

export type WordPartOfSpeech =
  | 'noun'
  | 'verb'
  | 'adverb'
  | 'adjective'
  | 'conjunction'
  | 'preposition'
  | 'pronoun'
  | 'determiner'
  | 'interjection';

export type WordBase = { word: string };
export type WordMeaningBase = { definition: string; partOfSpeech: WordPartOfSpeech };
export type WordExampleBase = { text: string };

export type WordExample = Entity<WordExampleBase>;
export type WordMeaning = Entity<WordMeaningBase & { examples: WordExample[]; synonyms?: Entity<WordBase>[] }>;
export type Word = Entity<WordBase & { meanings: WordMeaning[] }>;

export type WordListItem = Entity<WordBase>;
