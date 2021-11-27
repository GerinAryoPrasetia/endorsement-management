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
	Gender string
	Email string
	Password string
	Bio string
	Category []Category `gorm:"many2many:creator_category;"`
	SocialMedia []SocialMedia
}

type Category struct {
	gorm.Model
	CategoryName string
}

type SocialMedia struct {
	gorm.Model
	CreatorID int
	SocialMediaName string
	Url string
	Followers int
	VerifiedStatus bool
	RatePrice int

}

type FeaturedContent struct {
	gorm.Model
	Url string
}

func toCategoryRecords(category []creator.CategoryCore) []Category {
	convertedCategory := []Category{}
	for _, cat := range category {
		convertedCategory = append(convertedCategory, Category{
			CategoryName: cat.CategoryName,
		})
	}

	return convertedCategory
}

func toSocialMediaRecords(socmed []creator.SocialMediaCore) []SocialMedia {
	convertedSocmed := []SocialMedia{}
	for _, soc := range socmed {
		convertedSocmed = append(convertedSocmed, SocialMedia{
			SocialMediaName: soc.Name,
			Url: soc.Url,
			Followers: soc.Followers,
			VerifiedStatus: soc.VerifiedStatus,
			CreatorID: soc.CreatorID,
			RatePrice: soc.RatePrice,
		})
	}
	return convertedSocmed
}

func toFeaturedRecords(cont []creator.FeaturedContent) []FeaturedContent {
	convertedContent := []FeaturedContent{}
	for _,con := range cont {
		convertedContent = append(convertedContent, FeaturedContent{
			Url: con.Url,
		})
	}
	return convertedContent
}

//register creator
func toCreatorRecord(creator creator.Core) Creator {
	return Creator{
		Name: creator.Name,
		Location: creator.Location,
		Age: creator.Age,
		Gender: creator.Gender,
		Email: creator.Email,
		Password: creator.Password,
		Bio: creator.Bio,
		Category: toCategoryRecords(creator.Category),
		SocialMedia: toSocialMediaRecords(creator.SocialMedia),
	}
}
//get list
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