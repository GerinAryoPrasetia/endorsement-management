package request

import "project/features/creator"

type Creator struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Age int `json:"age"`
	Bio int `json:"bio"`
}

func ToCreatorCore(cr Creator) creator.Core {
	return creator.Core{
		ID: cr.ID,
		Name: cr.Name,
		Location: cr.Location,
		Age: cr.Age,
	}
}