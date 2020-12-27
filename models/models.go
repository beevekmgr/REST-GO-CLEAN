package models

import (
	"REST-GO-CLEAN/lib"

	"github.com/jinzhu/gorm"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMigrations),
)

// Migrations migration data type
type Migrations struct {
	db     lib.Database
	logger lib.Logger
	models []ModelMigration
}

// ModelMigration interface
type ModelMigration interface {
	Migrate(*gorm.DB) error
}

// NewMigrations creates new migrations instance
func NewMigrations(db lib.Database, logger lib.Logger) Migrations {
	return Migrations{
		db:     db,
		logger: logger,
		models: []ModelMigration{
			Product{},
		},
	}
}

// Migrate function to call when migrating
func (m Migrations) Migrate() {
	m.logger.Zap.Info("Automigrating schemas.")
	for _, model := range m.models {
		err := model.Migrate(m.db.DB)
		if err != nil {
			m.logger.Zap.Fatalf("Migration failure: ", err)
		}
	}
}
