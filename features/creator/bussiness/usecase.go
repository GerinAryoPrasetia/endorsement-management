package bussiness

import (
	"project/features/creator"

	"github.com/go-playground/validator/v10"
	
)

type creatorUsecase struct {
	creatorData creator.Data
	validate *validator.Validate
}

func NewCreatorBussiness (crData creator.Data) creator.Bussiness {

	return &creatorUsecase{
		creatorData: crData,
		validate: validator.New(),
	}
}

func (cu *creatorUsecase) GetAllData(search string) (resp []creator.Core) {
	resp = cu.creatorData.SelectData(search)
	return
}