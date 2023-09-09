package helpers

import (
	"idp-go/app/practice/models"
	config "idp-go/configs"
)

var db = config.GetDb()

func Get(id int) (*models.ToDo, error) {
	todo := &models.ToDo{}
	if err := db.Where("ID = ?", id).Find(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func GetList() ([]*models.ToDo, error) {
	todos := []*models.ToDo{}
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func Save(todo models.ToDo) error {
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

func Update(id int, todo models.ToDo) error {
	if err := db.Where("ID = ?", id).Updates(todo).Error; err != nil {
		return err
	}
	return nil
}

func Delete(id int) error {
	if err := db.Delete(&models.ToDo{}, "ID = ?", id).Error; err != nil {
		return err
	}
	return nil
}
