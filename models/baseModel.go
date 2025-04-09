package models

import (
	"Insignia-Backend/config"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SQLModel struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DBResponse `gorm:"-" json:"-"`
}

type DBResponse struct {
	DBRMessage string
	DBRStatus  int
}

func (s *SQLModel) HandleErr(result *gorm.DB, logger *log.Logger) {
	s.DBResponse.DBRMessage, s.DBResponse.DBRStatus = s.HandleDbError(result, logger)
	if s.DBRStatus != 0 {
		panic("Database Error")
	}
}

func (s *SQLModel) HandleDbError(result *gorm.DB, logger *log.Logger) (string, int) {
	//Handle common errors
	if result.Error != nil {
		logger.Println("Database Error: ", result.Error)
		switch {
		case errors.Is(result.Error, gorm.ErrRecordNotFound):
			return "Records not found", 404
		case errors.Is(result.Error, gorm.ErrDuplicatedKey):
			return "Record Exists", 400
		default:
			//Catch all other errors
			return "Internal Server Error, Kindly Contact Support", 500
		}
	}
	return "", 0
}

func (s *SQLModel) Insert(payload interface{}, logger *log.Logger) {
	// ToDO return iD??
	result := config.DBClient.Model(payload).Create(payload)
	s.HandleErr(result, logger)
}

func (s *SQLModel) FetchAll(payload, destination interface{}, logger *log.Logger) {
	result := config.DBClient.Model(payload).Find(destination)
	s.HandleErr(result, logger)
}

func (s *SQLModel) FetchByOrganizationID(target uuid.UUID, payload interface{}, logger *log.Logger) {
	result := config.DBClient.Where("organization_id = ?", target).Find(payload)
	s.HandleErr(result, logger)
}

func (s *SQLModel) IDUpsert(payload interface{}, targets []string, logger *log.Logger) {
	updateClause := clause.OnConflict{Columns: []clause.Column{{Name: "id"}}} // key column

	if len(targets) == 0 {
		updateClause.UpdateAll = true
	} else {
		updateClause.DoUpdates = clause.AssignmentColumns(targets)
	}
	result := config.DBClient.Clauses(updateClause).Create(payload)
	s.HandleErr(result, logger)
}
