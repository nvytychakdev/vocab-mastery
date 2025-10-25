package language

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/bregydoc/gtranslate"
	"github.com/go-chi/render"
	httpError "github.com/nvytychakdev/vocab-mastery/internal/app/http-error"
)

// Request
type LanguageTranslateRequest struct {
	Word         string `json:"word"`
	LanguageFrom string `json:"languageFrom"`
	LanguageTo   string `json:"languageTo"`
}

func (s *LanguageTranslateRequest) Bind(r *http.Request) error {
	if s.Word == "" {
		return errors.New("name field is required")
	}
	if s.LanguageFrom == "" {
		return errors.New("original language is required")
	}
	if s.LanguageTo == "" {
		return errors.New("target language is required")
	}
	return nil
}

// Response
type LanguageTranslateResponse struct {
	Word        string               `json:"word"`
	Language    string               `json:"language"`
	Translation string               `json:"translation"`
	Dictionary  []DictionaryApiEntry `json:"dictionary"`
}

func (*LanguageTranslateResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type DictionaryApiEntry struct {
	Word      string `json:"word"`
	Phonetic  string `json:"photenic,omitempty"`
	Phonetics []struct {
		Audio string `json:"audio,omitempty"`
		Text  string `josn:"text,omitempty"`
	} `json:"phonetics"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string   `json:"definition"`
			Synonyms   []string `json:"synonyms"`
			Antonyms   []string `json:"antonyms"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms"`
		Antonyms []string `json:"antonyms"`
	} `json:"meanings"`
	License struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"license"`
	SourceUrls []string `json:"sourceUrls"`
}

func (wh *LanguageHandler) LanguageTranslate(w http.ResponseWriter, r *http.Request) {
	var data = &LanguageTranslateRequest{}

	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusBadRequest, httpError.ErrInvalidPayload, err))
		return
	}

	translatedWord, err := gtranslate.TranslateWithParams(
		data.Word,
		gtranslate.TranslationParams{From: data.LanguageFrom, To: data.LanguageTo},
	)
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	dictionaryRes, err := http.Get(fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/%s/%s", data.LanguageFrom, data.Word))
	if err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}
	defer dictionaryRes.Body.Close()

	var dictionaryEntries []DictionaryApiEntry
	if err := render.DecodeJSON(dictionaryRes.Body, &dictionaryEntries); err != nil {
		render.Render(w, r, httpError.NewErrorResponse(http.StatusInternalServerError, httpError.ErrInternalServer, err))
		return
	}

	response := &LanguageTranslateResponse{
		Word:        data.Word,
		Language:    data.LanguageTo,
		Translation: translatedWord,
		Dictionary:  dictionaryEntries,
	}

	render.Render(w, r, response)
}
