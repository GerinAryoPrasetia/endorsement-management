package request

import "project/features/creator"

type CreatorRequest struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Age int `json:"age"`
	Bio string `json:"bio"`
	SocialMedia []SocialMediaRequest `json:"social_media"`
}

type SocialMediaRequest struct {
	Username string `json:"username"`
	Followers int `json:"followers"`
}

func (sr *SocialMediaRequest) toSocialMediaCore() creator.SocialMediaCore {
	return creator.SocialMediaCore{
		Name: sr.Username,
		Followers: sr.Followers,
	}
}

func ToSocialMediasCore(socials []SocialMediaRequest) []creator.SocialMediaCore {
	convertedSocialMedia := []creator.SocialMediaCore{}
	for _, socmed := range socials {
		convertedSocialMedia = append(convertedSocialMedia, socmed.toSocialMediaCore())
	}

	return convertedSocialMedia
}

func ToCreatorCore(cr CreatorRequest) creator.Core {
	return creator.Core{
		ID: cr.ID,
		Name: cr.Name,
		Location: cr.Location,
		Age: cr.Age,
		SocialMedia: ToSocialMediasCore(cr.SocialMedia),
	}
}