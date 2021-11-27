package bussiness

import (
	"project/features/creator"

	"github.com/go-playground/validator/v10"
)

type creatorUsecase struct {
	creatorData creator.Data
	validate *validator.Validate
}

func NewCreatorBussiness(crData creator.Data) creator.Bussiness {
	return &creatorUsecase{
		creatorData: crData,
		validate: validator.New(),
	}
}

func (cu *creatorUsecase) GetAllData(data string) (resp []creator.Core) {
	resp = cu.creatorData.SelectData(data)
	return
}

func (cu *creatorUsecase) GetCreatorByName(data string) (resp []creator.Core){
	resp = cu.creatorData.SelectCreatorByName(data)
	return
}

func (cu *creatorUsecase) RegisterCreator(data creator.Core) error {
	err := cu.creatorData.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}