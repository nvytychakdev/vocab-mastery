import { ApiListQuerySortDir } from '@core/api/api.interface';
import { BadgeColor } from '../../../../projects/ui/src/lib/badge/badge';
import { WordPartOfSpeech } from './word.interface';

export const WORDS_LIMIT = 1000;
export const WORDS_SORT_BY = 'word';
export const WORDS_SORT_DIR: ApiListQuerySortDir = 'asc';
export const WORDS_PART_OF_SPEECH_COLOR_MAP: Record<WordPartOfSpeech, BadgeColor> = {
  noun: 'blue',
  verb: 'purple',
  adjective: 'lime',
  adverb: 'orange',
  conjunction: 'sky',
  preposition: 'green',
  pronoun: 'pink',
  determiner: 'yellow',
  interjection: 'red',
};
