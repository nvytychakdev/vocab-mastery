import { Entity } from '@core/interfaces/dto.interface';

export type WordBase = { word: string };
export type WordMeaningBase = { definition: string; partOfSpeech: string };
export type WordExampleBase = { text: string };

export type Word = Entity<WordBase>;
export type WordExample = Entity<WordExampleBase>;
export type WordMeaning = Entity<WordMeaningBase & { examples: WordExample[]; synonyms: Word[] }>;
export type WordDetails = Entity<WordBase & { meanings: WordMeaning }>;

export type WordListItem = Entity<WordBase>;
