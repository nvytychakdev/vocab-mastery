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

func (wh *WordHandler) WordGetListByDictionary(w http.ResponseWriter, r *http.Request) {
	opts := middleware.GetQueryOptionsContext(r)
	dictionary := middleware.GetDictionaryContext(r)

	words, totalWords, err := wh.Deps.DB.Word().ListByDictionaryID(dictionary.ID, opts)
	if err != nil {
		slog.Error("Not able to get words by dictionary id", "error", err)
		return
	}

	var wordsList []*WordListItem

	for _, w := range words {
		item := &WordListItem{Word: w}
		wordsList = append(wordsList, item)
	}

	response := &WordGetListResponse{
		Items: wordsList,
		PaginationResponse: &model.PaginationResponse{
			Total:  totalWords,
			Offset: opts.Pagination.Offset,
			Limit:  opts.Pagination.Limit,
		},
	}

	render.Render(w, r, response)
}

func (wh *WordHandler) WordGetList(w http.ResponseWriter, r *http.Request) {
	opts := middleware.GetQueryOptionsContext(r)
	userId := middleware.GetAuthorizedUserId(r)

	words, totalWords, err := wh.Deps.DB.Word().ListAll(userId, opts)
	if err != nil {
		slog.Error("Not able to get words by dictionary id", "error", err)
		return
	}

	wordsList := make([]*WordListItem, 0, len(words))

	for _, w := range words {
		item := &WordListItem{Word: w}
		wordsList = append(wordsList, item)

	}

	response := &WordGetListResponse{
		Items: wordsList,
		PaginationResponse: &model.PaginationResponse{
			Total:  totalWords,
			Offset: opts.Pagination.Offset,
			Limit:  opts.Pagination.Limit,
		},
	}

	render.Render(w, r, response)
}
