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
	SocialMedia SocialMedia
}

type SocialMedia struct {
	Name string
	Url string
	Followers int
	VerifiedStatus bool
}

//function abstraction
type Bussiness interface {
	GetAllData(search string) (resp []Core)
}

type Data interface {
	SelectData(name string) (resp []Core)
}