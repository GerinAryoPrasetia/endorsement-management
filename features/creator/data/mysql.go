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
	return ToCore(&record)
}

func (cr *mysqlCreatorRepository) InsertData(data creator.Core) (creator.Core, error) {
	newCreator := Creator{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
		Location: data.Location,
		Age: data.Age,
		Gender: data.Gender,
		Bio: data.Bio,
	}
	
	err := cr.Conn.Create(&newCreator).Error
	if err != nil {
		return creator.Core{}, err
	}
	return ToCore(&newCreator), nil
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
	return ToCore(&recordData), nil
}

func (cr *mysqlCreatorRepository) UpdateData(data creator.Core) (error) {
	err := cr.Conn.Debug().Model(&Creator{}).Where("id = ?", data.ID).Updates(Creator{
		Name: data.Name,
		Age: data.Age,
		Location: data.Location,
		Gender: data.Gender,
		Bio: data.Bio,
	}).Error
	if err != nil {
		return nil
	}
	return nil
}

func (cr *mysqlCreatorRepository) DeleteCreator(data creator.Core) error {
	err := cr.Conn.Debug().Delete(&Creator{}, data.ID).Error

	if err != nil {
		return err
	}

	return nil
}