package flashcard

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type FlashcardSessionStartRequest struct {
	Timezone string `json:"date"`
	Date     string `json:"timezone"`
}

func (s *FlashcardSessionStartRequest) Bind(r *http.Request) error {
	return nil
}

type FlashcardSessionStartResponse struct {
	SessionID     uuid.UUID           `json:"sessionId"`
	CardsTotal    int                 `json:"cardsTotal"`
	CardsAnswered int                 `json:"cardsAnswered"`
	CardsCorrect  int                 `json:"cardsCorrect"`
	NextCard      model.FlashcardCard `json:"nextCard"`
}

func (*FlashcardSessionStartResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (fc *FlashcardHandler) FlashcardSessionStart(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetAuthorizedUserId(r)
	data := &FlashcardSessionStartRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	// 1. Check existing `flashcard_engagement_state`, if does not exists create one
	engagement, err := fc.Deps.FlashcardSessionService.GetUserEngagement(userID)

	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 2. Update `flashcard_engagement_state` entry with last date
	err = fc.Deps.DB.FlashcardEngagement().UpdateDatesByUserID(userID, time.Now(), time.Now())
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 3. Check existing session `flashcard_session`, pull current session if it exists, otherwise create new one
	session, err := fc.Deps.FlashcardSessionService.StartOrResume(userID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	card, err := fc.Deps.FlashcardCardService.GenerateCard(userID, session.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	_, correctAttemptsCount, err := fc.Deps.DB.FlashcardAttempt().GetCorrectBySessionID(session.ID)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	// 4. Retrieve user's date + timezone for `flashcard_days`
	// 5. Check existing `flashcard_days`, if does not exists create new one
	slog.Info("Engagement created", "engagement", engagement)

	// 6. Return response with current state of the session
	response := &FlashcardSessionStartResponse{
		SessionID:     session.ID,
		CardsTotal:    session.CardsTotal,
		CardsAnswered: session.CardsCompleted,
		CardsCorrect:  correctAttemptsCount,
		NextCard:      *card,
	}

	render.Render(w, r, response)

}
