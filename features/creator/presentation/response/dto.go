package response

import (
	"project/features/creator"
	"time"
)

type Creator struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Name string `json:"name"`
	Age int `json:"age"`
	Location string `json:"location"`
	SocialMedia []creator.SocialMediaCore `json:"social_media"`
	FeaturedContent []creator.FeaturedContentCore `json:"featured_content"`
}

func FromCore(core creator.Core) Creator {
	return Creator{
		ID: core.ID,
		Name: core.Name,
		UpdatedAt: core.UpdatedAt,
		CreatedAt: core.CreatedAt,
		Age: core.Age,
		Location: core.Location,
		SocialMedia: core.SocialMedia,
		FeaturedContent: core.FeaturedContent,
	}
}

func FromCoreSlice(core []creator.Core) []Creator {
	var crArray []Creator
	for key := range core{
		crArray=append(crArray, FromCore(core[key]))
	}
	return crArray
}