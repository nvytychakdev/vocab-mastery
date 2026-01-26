export type FlashcardType = 'recall' | 'recognition';

export type FlashcardCardChoice = {
  translationId: string;
  translation: string;
};

export type FlashcardCard = {
  meaningId: string;
  wordId: string;
  word: string;
  meaning: string;
  type: string;
  choices: FlashcardCardChoice[];
};

export type FlashcardSession = {
  sessionId: string;
  cardsTotal: number;
  cardsAnswered: number;
  cardsCorrect: number;
  nextCard: FlashcardCard;
};

export type FlashcardSessionRequest = {
  timezone: string;
  date: string;
};

export type FlashcardSessionAnswerRequest = {
  meaningId: string;
  selectedAnswer: string;
};

export type FlashcardSessionAnswerResult = {
  isCorrect: string;
  correctAnswer: string;
  selectedAnswer: string;
};

export type FlashcardSessionAnswer = {
  sessionId: string;
  cardsTotal: number;
  cardsAnswered: number;
  cardsCorrect: number;
  isCompleted: boolean;
  result: FlashcardSessionAnswerResult;
  nextCard?: FlashcardCard;
};

export type FlashcardSessionState = {
  sessionId: string;
  cardsTotal: number;
  cardsAnswered: number;
  cardsCorrect: number;
  isCompleted: boolean;
  isAnswered: boolean;
  currentCard: FlashcardCard;
  nextCard: FlashcardCard | null;
  answerResult: FlashcardSessionAnswerResult | null;
  historyCards: FlashcardCard[];
};
