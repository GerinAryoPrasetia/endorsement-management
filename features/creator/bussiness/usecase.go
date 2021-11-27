package bussiness

import (
	"project/features/creator"
)

type creatorUsecase struct {
	creatorData creator.Data
	
}

func NewCreatorBussiness(crData creator.Data) creator.Bussiness {
	return &creatorUsecase{crData}
}

func (cu *creatorUsecase) GetAllData(data string) (resp []creator.Core) {
	resp = cu.creatorData.SelectData(data)
	return
}

func (cu *creatorUsecase) GetCreatorByName(data string) (resp creator.Core){
	resp = cu.creatorData.SelectCreatorByName(data)
	return resp
}

func (cu *creatorUsecase) GetCreatorByID(data creator.Core) (creator.Core, error) {
	resp, err := cu.creatorData.SelectCreatorByID(data)
	
	if err != nil {
		return creator.Core{}, nil
	}
	return resp, nil
}

func (cu *creatorUsecase) RegisterCreator(data creator.Core) error {
	err := cu.creatorData.InsertData(data)

	if err != nil {
		return err
	}

	return nil
}