package services

import (
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardCardService interface {
	GetSessionMeaning(userID uuid.UUID, session *model.FlashcardSession) (*model.WordMeaning, error)
	GetSessionAnswer(session *model.FlashcardSession, wordID uuid.UUID) (*model.WordTranslation, error)
	GetSessionChoices(session *model.FlashcardSession, word *model.Word, meaning *model.WordMeaning) (uuid.UUID, []model.FlashcardCardChoice, error)
	GenerateCard(userID uuid.UUID, sessionID uuid.UUID) (*model.FlashcardCard, error)
	ValidateSessionChoice(session *model.FlashcardSession, choiceID uuid.UUID) bool
}

type flashcardCardService struct {
	DB db.DB
}

func NewFlashcardCardService(db db.DB) FlashcardCardService {
	return &flashcardCardService{DB: db}
}

/*
Get random word for card:

SELECT wm.id AS meaning_id
FROM word_meanings wm
JOIN words w ON w.id = wm.word_id
LEFT JOIN flashcard_attempts fa
    ON fa.meaning_id = wm.id
    AND fa.session_id = current_session_id
WHERE dictionary_id = ?
  AND wm.level <= user_current_level
  AND fa.id IS NULL
ORDER BY RANDOM()
LIMIT 1

*/

func (fcs *flashcardCardService) GetSessionMeaning(userID uuid.UUID, session *model.FlashcardSession) (*model.WordMeaning, error) {
	if session.CurrentMeaningID != uuid.Nil {
		meaning, err := fcs.DB.WordMeaning().GetByID(session.CurrentMeaningID)
		if err != nil {
			return nil, err
		}

		return meaning, err
	}

	dictionaryID, err := uuid.Parse("fbdff10e-2d36-4387-be01-489b917a36e3")
	if err != nil {
		return nil, err
	}

	//  1. Retrieve random meaning based on status weights
	meaningID, _, err := fcs.DB.FlashcardSession().GetRandomMeaningToLearn(userID, dictionaryID, session.ID)
	if err != nil {
		return nil, err
	}

	meaning, err := fcs.DB.WordMeaning().GetByID(meaningID)
	if err != nil {
		return nil, err
	}

	return meaning, err
}

func (fcs *flashcardCardService) GetSessionAnswer(session *model.FlashcardSession, wordID uuid.UUID) (*model.WordTranslation, error) {
	if session.CurrentMeaningTranslationID != uuid.Nil {
		return fcs.DB.WordTranslation().GetByID(session.CurrentMeaningTranslationID)
	}

	return fcs.DB.FlashcardSession().GetRandomAnswerByWordID(wordID)
}

func (fcs *flashcardCardService) ValidateSessionChoice(session *model.FlashcardSession, choiceID uuid.UUID) bool {
	if session.CurrentMeaningChoicesIDs != nil {
		for _, cID := range session.CurrentMeaningChoicesIDs {
			if cID == choiceID {
				return true
			}
		}
	}

	return false
}

func (fcs *flashcardCardService) GetSessionChoices(session *model.FlashcardSession, word *model.Word, meaning *model.WordMeaning) (uuid.UUID, []model.FlashcardCardChoice, error) {
	// Check current session, if its ongoing, use stored options that was generated before
	if session.CurrentMeaningChoicesIDs != nil && session.CurrentMeaningTranslationID != uuid.Nil {
		answers, _, err := fcs.DB.WordTranslation().ListByIDs(session.CurrentMeaningChoicesIDs)
		if err != nil {
			return uuid.Nil, nil, err
		}

		var choices []model.FlashcardCardChoice
		for _, answer := range answers {
			choices = append(choices, model.FlashcardCardChoice{
				TranslationID: answer.ID,
				Translation:   answer.Translation,
			})
		}

		return session.CurrentMeaningTranslationID, choices, nil
	}

	// Generate new answer with options of session is clean
	answer, err := fcs.GetSessionAnswer(session, word.ID)
	if err != nil {
		return uuid.Nil, nil, err
	}

	answerChoices, _, err := fcs.DB.FlashcardSession().ListRandomAnswers(meaning.ID)
	if err != nil {
		return uuid.Nil, nil, err
	}

	// Ensure one of the translations is correct
	choices := []model.FlashcardCardChoice{
		{TranslationID: answer.ID, Translation: answer.Translation},
	}

	for _, t := range answerChoices {
		choices = append(choices, model.FlashcardCardChoice{TranslationID: t.ID, Translation: t.Translation})
	}

	// 4. Shuffle choices for the user
	rand.Shuffle(len(choices), func(i, j int) {
		choices[i], choices[j] = choices[j], choices[i]
	})

	return answer.ID, choices, nil
}

func (fcs *flashcardCardService) GenerateCard(userID uuid.UUID, sessionID uuid.UUID) (*model.FlashcardCard, error) {
	session, err := fcs.DB.FlashcardSession().GetByID(sessionID)
	if err != nil {
		return nil, err
	}

	if session.IsCompleted() {
		return nil, nil
	}

	meaning, err := fcs.GetSessionMeaning(userID, session)
	if err != nil {
		return nil, err
	}

	// 2. Gather word translations as Choices and return outcome card to the user
	word, err := fcs.DB.Word().GetByID(meaning.WordID)
	if err != nil {
		return nil, err
	}

	// 3. Generate list of choices based on session or meaning
	answerID, choices, err := fcs.GetSessionChoices(session, word, meaning)
	if err != nil {
		return nil, err
	}

	// 4. Update current session with latest card data
	var choicesIDs uuid.UUIDs
	for _, choice := range choices {
		choicesIDs = append(choicesIDs, choice.TranslationID)
	}
	err = fcs.DB.FlashcardSession().UpdateCurrentAnswer(session.ID, &meaning.ID, &answerID, choicesIDs)
	if err != nil {
		return nil, err
	}

	card := model.FlashcardCard{
		MeaningID: meaning.ID,
		WordID:    meaning.WordID,
		Meaning:   meaning.Definition,
		Word:      word.Word,
		Type:      "recall",
		Choices:   choices,
	}

	return &card, nil
}
