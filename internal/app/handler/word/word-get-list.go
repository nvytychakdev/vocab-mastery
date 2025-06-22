package word

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"
	"github.com/nvytychakdev/vocab-mastery/internal/app/middleware"
	"github.com/nvytychakdev/vocab-mastery/internal/app/model"
)

type WordGetListResponse struct {
	Items []*WordListItem `json:"items"`
	*model.PaginationResponse
}

type WordListItem struct {
	*model.Word
	Translations []*model.Translation `json:"translations,omitempty"`
}

func (u *WordGetListResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (wh *WordHandler) WordGetList(w http.ResponseWriter, r *http.Request) {
	include := middleware.GetIncludeContext(r)
	pagination := middleware.GetPaginationContext(r)
	dictionary := middleware.GetDictionaryContext(r)

	words, totalWords, err := wh.Deps.DB.Word().GetByDictionaryID(dictionary.ID, pagination)
	if err != nil {
		slog.Error("Not able to get words by dictionary id", "error", err)
		return
	}

	var wordsList []*WordListItem
	group := make(map[string][]*model.Translation)

	for _, w := range words {
		item := &WordListItem{Word: w}
		wordsList = append(wordsList, item)
	}

	if include != nil && include["translations"] {
		wordIDs := make([]string, 0, len(words))

		for _, word := range words {
			wordIDs = append(wordIDs, word.ID)
		}

		translations, err := wh.Deps.DB.Translation().ListByWordIDs(wordIDs)
		if err == nil {
			for _, t := range translations {
				group[t.WordId] = append(group[t.WordId], t)
			}

			for _, w := range wordsList {
				w.Translations = group[w.ID]
			}
		}
	}

	response := &WordGetListResponse{
		Items: wordsList,
		PaginationResponse: &model.PaginationResponse{
			Total:  totalWords,
			Offset: pagination.Offset,
			Limit:  pagination.Limit,
		},
	}

	render.Render(w, r, response)
}
