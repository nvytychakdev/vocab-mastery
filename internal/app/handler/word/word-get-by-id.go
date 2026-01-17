package word

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordGetByIdResponse struct {
	ID        uuid.UUID                    `json:"id"`
	Word      string                       `json:"word"`
	CreatedAt time.Time                    `json:"createdAt"`
	Meanings  []WordGetByIdMeaningResponse `json:"meanings,omitempty"`
}

type WordGetByIdMeaningResponse struct {
	*model.WordMeaning
	Examples     []*model.WordExample     `json:"examples,omitempty"`
	Synonyms     []*model.WordSynonym     `json:"synonyms,omitempty"`
	Translations []*model.WordTranslation `json:"translations,omitempty"`
}

func (u *WordGetByIdResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordGetByID(w http.ResponseWriter, r *http.Request) {
	word := middleware.GetWordContext(r)

	response := &WordGetByIdResponse{
		ID:        word.ID,
		Word:      word.Word,
		CreatedAt: word.CreatedAt,
	}

	meanings, _, err := wh.Deps.DB.WordMeaning().ListAllByWordIDs(uuid.UUIDs{word.ID})
	if err != nil {
		slog.Error("Not able to get meanings by word id", "error", err)
		return
	}

	meaningIds := make(uuid.UUIDs, 0, len(meanings))
	for _, meaning := range meanings {
		meaningIds = append(meaningIds, meaning.ID)
	}

	examples, _, err := wh.Deps.DB.WordExample().ListAllByMeaningIDs(meaningIds)
	if err != nil {
		slog.Error("Not able to get examples by meanings ids", "error", err)
		return
	}

	synonyms, _, err := wh.Deps.DB.WordSynonym().ListAllByMeaningIDs(meaningIds)
	if err != nil {
		slog.Error("Not able to get synonyms by meanings ids", "error", err)
		return
	}

	translations, _, err := wh.Deps.DB.WordTranslation().ListAllByMeaningIDs(meaningIds)
	if err != nil {
		slog.Error("Not able to get synonyms by meanings ids", "error", err)
		return
	}

	var meaningsRes []WordGetByIdMeaningResponse
	for _, meaning := range meanings {
		meaningRes := WordGetByIdMeaningResponse{
			WordMeaning: meaning,
		}

		for _, example := range examples {
			if example.MeaningID == meaning.ID {
				meaningRes.Examples = append(meaningRes.Examples, example)
			}
		}

		for _, synonym := range synonyms {
			if synonym.MeaningID == meaning.ID {
				meaningRes.Synonyms = append(meaningRes.Synonyms, synonym)
			}
		}

		for _, translation := range translations {
			if translation.MeaningID == meaning.ID {
				meaningRes.Translations = append(meaningRes.Translations, translation)
			}
		}

		meaningsRes = append(meaningsRes, meaningRes)
	}

	response.Meanings = meaningsRes

	render.Render(w, r, response)
}
