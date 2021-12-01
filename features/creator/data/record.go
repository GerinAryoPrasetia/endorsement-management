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
	SocialMedia []SocialMedia `gorm:"foreignKey:CreatorID"`
	FeaturedContent []FeaturedContent
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
	CreatorID int
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

func toCategoryCore(cat Category) creator.CategoryCore {
	return creator.CategoryCore{
		ID: int(cat.ID),
		CategoryName: cat.CategoryName,
	}
}

func toCategoryCoreList(categoryList []Category) []creator.CategoryCore{
	catList := []creator.CategoryCore{}
	for _, category := range categoryList {
		catList = append(catList, toCategoryCore(category))
	}
	return catList
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

func toSocialMediaCore(socmed SocialMedia) creator.SocialMediaCore {
	return creator.SocialMediaCore{
		ID: int(socmed.ID),
		CreatorID: socmed.CreatorID,
		Name: socmed.SocialMediaName,
		Followers: socmed.Followers,
		VerifiedStatus: socmed.VerifiedStatus,
		Url: socmed.Url,
	}
}

func toSocialMediaCoreList(socmedList []SocialMedia) []creator.SocialMediaCore {
	sl := []creator.SocialMediaCore{}
	for _,socmed := range socmedList {
		sl = append(sl, toSocialMediaCore(socmed))
	}
	return sl
}

func toFeaturedRecords(cont []creator.FeaturedContentCore) []FeaturedContent {
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
		// FeaturedContent: toFeaturedRecords(creator.FeaturedContent),
	}
}
//get list
func (a *Creator) ToCore() creator.Core {
	return creator.Core{
		ID:		int(a.ID),
		Name:	a.Name,
		Location: a.Location,
		Age:	a.Age,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func ToCore(data *Creator) creator.Core {
	return creator.Core{
		ID: int(data.ID),
		Name: data.Name,
		Location: data.Location,
		Age: data.Age,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}

func ToCoreList(resp []Creator) []creator.Core {
	a := []creator.Core{}
	for key := range resp {
		a = append(a, resp[key].ToCore())
	}
	return a
}
