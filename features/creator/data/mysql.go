package data

import (
	"errors"
	"project/features/creator"

	"gorm.io/gorm"
)

type mysqlCreatorRepository struct {
	Conn *gorm.DB
}

func NewCreatorRepository(conn *gorm.DB) creator.Data {
	return &mysqlCreatorRepository{Conn: conn}
}

func (cr *mysqlCreatorRepository) SelectData(name string) (resp []creator.Core){
	var record []Creator
	if err := cr.Conn.Find(&record).Error; err != nil {
		return []creator.Core{}
	}
	return ToCoreList(record)
}

func (cr *mysqlCreatorRepository) SelectCreatorByName(name string) (resp creator.Core){
	var record Creator

	if err:= cr.Conn.Where("name = ?", name).Find(&record).Error; err != nil {
		return creator.Core{}
	}
	return ToCore(record)
}

func (cr *mysqlCreatorRepository) InsertData(data creator.Core) error {
	recordData := toCreatorRecord(data)
	err := cr.Conn.Create(&recordData)
	
	if err != nil {
		return err.Error
	}
	return nil
}

func (cr *mysqlCreatorRepository) SelectCreatorByID(data creator.Core) (resp creator.Core, err error) {
	var recordData Creator

	err = cr.Conn.Find(&recordData, data.ID).Error
	if recordData.Name == "" && recordData.ID == 0{
		return creator.Core{}, errors.New("Creator Not Exist")
	}
	if err != nil {
		return creator.Core{}, err
	}
	return ToCore(recordData), nil
}