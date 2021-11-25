package factory

import (
	"project/config"
	_creator_bussiness "project/features/creator/bussiness"
	_creator_data "project/features/creator/data"
	_creator_presentation "project/features/creator/presentation"
)

type Presenter struct {
	CreatorPresentation *_creator_presentation.CreatorHandler

}

func Init() Presenter {
	creatorData := _creator_data.NewCreatorRepository(config.DB)
	creatorBussiness := _creator_bussiness.NewCreatorBussiness(creatorData)
	creatorPresentation := _creator_presentation.NewCreatorHandler(creatorBussiness)

	return Presenter{
		CreatorPresentation: creatorPresentation,
	}
}