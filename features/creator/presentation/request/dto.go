package request

import "project/features/creator"

type CreatorRequest struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Age int `json:"age"`
	Bio string `json:"bio"`
}

func ToCreatorCore(cr CreatorRequest) creator.Core {
	return creator.Core{
		ID: cr.ID,
		Name: cr.Name,
		Location: cr.Location,
		Age: cr.Age,
	}
}