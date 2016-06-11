package services

import (
	"github.com/evanh/fundmyworld/repositories"
)

type FundService struct {
	repo *repositories.FundRepository
}

func NewFundService(repo *repositories.FundRepository) (*FundService, error) {
	return &FundService{repo}, nil
}

func (fs *FundService) GetTestRows() ([]map[string]interface{}, error) {
	return fs.repo.GetTestRows()
}
