package repositories

import (
	"github.com/evanh/fundmyworld/db"
)

type FundRepository struct {
	DB *db.AugmentedDB
}

func NewFundRepository(db *db.AugmentedDB) (*FundRepository, error) {
	return &FundRepository{db}, nil
}

func (fr *FundRepository) GetTestRows() ([]map[string]interface{}, error) {
	return fr.DB.MapQuery("SELECT * FROM test")
}
