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
func (cr *mysqlCreatorRepository) SelectData(name string) (resp []creator.Core){
	var record []Creator
	if err := cr.Conn.Find(&record).Error; err != nil {
		return []creator.Core{}
	}
	return toCoreList(record)
}

func (cr *mysqlCreatorRepository) SelectCreatorByName(name string) (resp []creator.Core){
	var record []Creator

	if err:= cr.Conn.Where("name = ?", name).Find(&record).Error; err != nil {
		return []creator.Core{}
	}
	return toCoreList(record)
}
