package response

import (
	"project/features/creator"
	"time"

	"github.com/labstack/echo/v4"
)

type Creator struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Location string `json:"location"`
	SocialMedia []SocialMediaResponse `json:"social_media"`
	Token string `json:"token"`
	// FeaturedContent []creator.FeaturedContentCore `json:"featured_content"`
}

type SocialMediaResponse struct {
	Username string `json:"username"`
	Followers int `json:"followers"`
}

type Response struct {
	Message string
	Data    interface{}
}

func FromCore(core creator.Core) Creator {
	return Creator{
		ID: core.ID,
		Name: core.Name,
		UpdatedAt: core.UpdatedAt,
		CreatedAt: core.CreatedAt,
		Age: core.Age,
		Location: core.Location,
		// SocialMedia: core.SocialMedia,
		// FeaturedContent: core.FeaturedContent,
	}
}

func toSocialMediaResponse(socmed creator.SocialMediaCore) SocialMediaResponse {
	return SocialMediaResponse{
		Username: socmed.Name,
		Followers: socmed.Followers,
	}
}

func toSocialMediaList(socmedList []creator.SocialMediaCore) []SocialMediaResponse {
	convertedSocialMedia := []SocialMediaResponse{}

	for _, socmed := range socmedList {
		convertedSocialMedia = append(convertedSocialMedia, toSocialMediaResponse(socmed))
	}
	return convertedSocialMedia
}

func NewErrorResponse(e echo.Context, msg string, code int) error {
	return e.JSON(code, Response{
		Message: msg,
	})
}

func ToCreatorResponse(creator creator.Core) Creator {
	return Creator{
		ID: creator.ID,
		Name: creator.Name,
		Age: creator.Age,
		Location: creator.Location,
		Gender: creator.Gender,
		SocialMedia: toSocialMediaList(creator.SocialMedia),
		// FeaturedContent: creator.FeaturedContent,
	}
}

func FromCoreSlice(core []creator.Core) []Creator {
	var crArray []Creator
	for key := range core{
		crArray=append(crArray, FromCore(core[key]))
	}
	return crArray
}