export const ALPHABET = Array.from('ABCDEFGHIJKLMNOPQRSTUVWXYZ');

export const ENGLISH_LEVEL = {
  A1: 'A1',
  A2: 'A2',
  B1: 'B1',
  B2: 'B2',
  C1: 'C1',
  C2: 'C2',
  PERSONAL: 'PERSONAL',
} as const;

export type EnglishLevel = (typeof ENGLISH_LEVEL)[keyof typeof ENGLISH_LEVEL];
