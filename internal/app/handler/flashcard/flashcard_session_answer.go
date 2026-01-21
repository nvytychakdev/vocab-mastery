package flashcard

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardSessionAnswerRequest struct {
	MeaningID      uuid.UUID `json:"meaningId"`
	SelectedAnswer uuid.UUID `json:"selectedAnswer"`
	ResponseTimeMs int       `json:"responseTimeMs"`
}

func (s *FlashcardSessionAnswerRequest) Bind(r *http.Request) error {
	if s.MeaningID == uuid.Nil {
		return errors.New("meaning id field is required")
	}

	if s.SelectedAnswer == uuid.Nil {
		return errors.New("selected answer field is required")
	}

	return nil
}

type FlashcardSessionAnswerResponse struct {
	SessionID     uuid.UUID                   `json:"sessionId"`
	CardsTotal    int                         `json:"cardsTotal"`
	CardsAnswered int                         `json:"cardsAnswered"`
	CardsCorrect  int                         `json:"cardsCorrect"`
	IsCompleted   bool                        `json:"isCompleted"`
	Result        model.FlashcardAnswerResult `json:"result"`
	NextCard      model.FlashcardCard         `json:"nextCard"`
}

func (*FlashcardSessionAnswerResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (fc *FlashcardHandler) FlashcardSessionAnswer(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetAuthorizedUserId(r)
	session := middleware.GetFlashcardSessionContext(r)

	if session.IsCompleted() {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, errors.New("Session completed")))
		return
	}

	// 1. Validate payload
	var data = &FlashcardSessionAnswerRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	// 2. Make sure provided answer was one of the choices
	if !fc.Deps.FlashcardCardService.ValidateSessionChoice(session, data.SelectedAnswer) {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, errors.New("Invalid answer, out of choices range")))
		return
	}

	// 3. Verify answer from the choices list
	isCorrect, err := fc.Deps.FlashcardSessionService.CheckAnswer(data.MeaningID, data.SelectedAnswer)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 4. Reset current session meaning/answer/choices
	err = fc.Deps.DB.FlashcardSession().UpdateCurrentAnswer(session.ID, nil, nil, nil)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 5. Generate new card
	card, err := fc.Deps.FlashcardCardService.GenerateCard(userID, session.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 6. Update session counters
	completedCount := session.CardsCompleted + 1
	err = fc.Deps.DB.FlashcardSession().UpdateCompletedCounter(session.ID, completedCount)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 7. Create attempt of the flashcard submission
	_, err = fc.Deps.DB.FlashcardAttempt().Create(model.FlashcardAttemptCreate{
		SessionID:      session.ID,
		MeaningID:      data.MeaningID,
		PromptLanguage: "en",
		AnswerLanguage: "ru",
		IsCorrect:      isCorrect,
		ResponseTimeMs: data.ResponseTimeMs,
	})

	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 8. Count valid answers to provide stats back to user's response
	_, correctAttemptsCount, err := fc.Deps.DB.FlashcardAttempt().GetCorrectBySessionID(session.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	result := model.FlashcardAnswerResult{
		IsCorrect:      isCorrect,
		SelectedAnswer: data.SelectedAnswer,
		CorrectAnswer:  session.CurrentMeaningTranslationID,
	}

	// 9. Update stats on session in case session was completed
	cardsAnswered := session.CardsCompleted + 1
	isCompleted := session.CardsTotal == cardsAnswered

	if isCompleted {
		err = fc.Deps.DB.FlashcardSession().UpdateEndedAt(session.ID, time.Now())
		if err != nil {
			render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
			return
		}
	}

	response := &FlashcardSessionAnswerResponse{
		SessionID:     session.ID,
		CardsTotal:    session.CardsTotal,
		CardsAnswered: cardsAnswered,
		CardsCorrect:  correctAttemptsCount,
		IsCompleted:   isCompleted,
		Result:        result,
		NextCard:      *card,
	}
	render.Render(w, r, response)
}
