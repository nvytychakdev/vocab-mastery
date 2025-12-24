import { Entity } from '@core/interfaces/dto.interface';

export type DictionaryBase = { title: string; level: string; isDefault: boolean };
export type Dictionary = Entity<DictionaryBase>;
export type DictionaryListItem = Entity<DictionaryBase>;
