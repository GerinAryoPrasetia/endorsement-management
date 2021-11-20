package data

import (
	"project/features/creator"

	"gorm.io/gorm"
)

type mysqlCreatorRepository struct {
	Conn *gorm.DB
}

func NewCreatorRepository(conn *gorm.DB) creator.Data {
	return &mysqlCreatorRepository{
		Conn: conn,
	}
}
func (ar *mysqlCreatorRepository) SelectData(title string) (resp []creator.Core){
	var record []Creator
	
	return toCoreList(record)
}