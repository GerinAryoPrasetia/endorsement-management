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
	FeaturedContent []FeaturedContentCore 
}

type SocialMediaCore struct {
	ID int
	CreatorID int
	Name string
	Url string
	Followers int
	VerifiedStatus bool
	RatePrice int
}

type CategoryCore struct {
	ID int
	CategoryName string
}

type FeaturedContentCore struct {
	ID int
	UserID int
	Url string
}

//function abstraction
type Bussiness interface {
	GetAllData(data string) (resp []Core)
	GetCreatorByName(data string) (resp Core)
	RegisterCreator(data Core) (creator Core,  err error)
	GetCreatorByID(data Core) (creator Core, err error)
}

type Data interface {
	SelectData(name string) (resp []Core)
	SelectCreatorByName(name string) (resp Core)
	InsertData(data Core) (creator Core, err error)
	SelectCreatorByID(data Core) (creator Core, err error)
}