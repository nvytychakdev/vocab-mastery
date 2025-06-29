import { Entity } from '../../core/interfaces/dto.interface';

export type TranslationBase = { word: string; language: string };
export type Translation = Entity<TranslationBase>;
export type TranslationListItem = Entity<TranslationBase>;
