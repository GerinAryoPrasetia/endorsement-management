package creator

import (
	"time"
)

type Core struct {
	ID int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	Location string
	Age int
	Gender string
	Bio string
	Email string
	Password string
	SocialMedia []SocialMediaCore
	Category []CategoryCore 
}

type SocialMediaCore struct {
	ID int
	Name string
	Url string
	Followers int
	VerifiedStatus bool
	CreatorID int
	RatePrice int
}

type CategoryCore struct {
	ID int
	CategoryName string
}

type FeaturedContent struct {
	ID int
	Url string
}

//function abstraction
type Bussiness interface {
	GetAllData(search string) (resp []Core)
	GetCreatorByName(search string) (resp []Core)
}

type Data interface {
	SelectData(name string) (resp []Core)
	SelectCreatorByName(name string) (resp []Core)
}