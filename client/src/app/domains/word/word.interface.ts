import { Entity } from '@core/interfaces/dto.interface';
import { Translation } from '@domain/translation/translation.interface';

export type WordBase = { word: string; language: string };
export type WithTranslations<T extends WordBase> = T & { translations: Translation[] };

export type Word = Entity<WordBase>;
export type WordWithTranslation = Entity<WordBase & { translation: Translation[] }>;

export type WordListItem = Entity<WordBase>;
export type WordListItemWithTranslation = Entity<WordBase & { translation: Translation[] }>;
