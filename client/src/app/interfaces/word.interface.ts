export type Word = {
  text: string;
  translation: string[];
  translationLang: string;
  definitions: {
    noun?: string[];
    verb?: string[];
  };
};
