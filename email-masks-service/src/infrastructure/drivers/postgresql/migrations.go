package postgresql

import (
	"email-masks-service/src/business/entities"
	"email-masks-service/src/infrastructure/configuration"
	"gorm.io/gorm"
)

type Migrations struct {
	logger configuration.Logger
	db     *gorm.DB
}

func NewMigrations() Migrations {
	config := configuration.LoadConfig()
	loggerLevel := config.GetString("logger.level")
	logger := configuration.NewLogger(loggerLevel)
	db := NewConnection(config.GetString("postgres.dsn"))
	return Migrations{logger, db}
}

func (p Migrations) Apply() {
	p.logger.Info("Starting migration")

	err := p.db.AutoMigrate(&entities.EmailMask{})
	if err != nil {
		p.logger.Fatal("Error applying migrations")
	}

	p.logger.Info("Done.")
}
