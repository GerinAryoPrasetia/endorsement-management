package data

import (
	"project/features/creator"

	"gorm.io/gorm"
)
//record struct on data layer
type Creator struct {
	gorm.Model
	Name string
	Location string
	Age int
}

type SocialMedia struct {
	gorm.Model
	SocialMediaName string
	Url string
	Followers int
	VerifiedStatus bool
}

func (a *Creator) toCore() creator.Core {
	return creator.Core{
		ID:		int(a.ID),
		Name:	a.Name,
		Location: a.Location,
		Age:	a.Age,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []Creator) []creator.Core {
	a := []creator.Core{}
	for key := range resp {
		a = append(a, resp[key].toCore())
	}
	return a
}