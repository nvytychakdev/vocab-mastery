package migrationsData

import (
	"log/slog"
)

type WordsDef struct {
	Word     string   `json:"word"`
	Synonyms []string `json:"synonyms"`
	Meaning  string   `json:"meaning"`
	Example  string   `json:"example"`
}

var InitialDataLoad = DataMigration{
	Version: 1,
	Name:    "Initial Data Load",
	Run: func() error {
		slog.Info("[DATA_MIGRATION] Successfully started initial data migration!!!")
		// contents, err := os.ReadFile("seeds/dictionaries/a1-list-v3.json")

		// if err != nil {
		// 	slog.Error("[DATA_MIGRATION] Error", "err", err)
		// 	return err
		// }

		// var words []WordsDef
		// if err = json.Unmarshal(contents, &words); err != nil {
		// 	slog.Error("[DATA_MIGRATION] Error", "err", err)
		// 	return err
		// }

		// slog.Info("Files loaded", "words", words)

		return nil
	},
}
