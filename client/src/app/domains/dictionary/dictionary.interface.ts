import { Entity } from '@core/interfaces/dto.interface';

export type DictionaryBase = { name: string; description: string };
export type Dictionary = Entity<DictionaryBase>;
export type DictionaryListItem = Entity<DictionaryBase>;
