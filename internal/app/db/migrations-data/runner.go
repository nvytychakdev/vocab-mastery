package migrationsData

import (
	"sort"

	"github.com/nvytychakdev/vocab-mastery/internal/app/db"
)

type DataMigration struct {
	Version int
	Name    string
	Run     func() error
}

var DataMigrationsList = []DataMigration{
	InitialDataLoad,
}

func RunLatest(mr db.MigrationRepo) error {
	if len(DataMigrationsList) == 0 {
		return nil
	}

	sort.Slice(DataMigrationsList, func(i, j int) bool {
		return DataMigrationsList[i].Version < DataMigrationsList[j].Version
	})

	latestMigration := DataMigrationsList[len(DataMigrationsList)-1]
	alreadyMigrated, err := mr.Check(latestMigration.Version)
	if err != nil {
		return err
	}

	if !alreadyMigrated {
		return mr.Run(latestMigration.Version, latestMigration.Run)
	}

	return nil
}
