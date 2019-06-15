package db

import (
	"classtime/models"
	u "classtime/utils"
	"fmt"
)

func CreateDiscipline(discipline *models.Discipline) map[string]interface{} {
	GetDB().Create(discipline)

	if discipline.ID <= 0 {
		return u.Message(false, "Failed to create discipline, connection error.")
	}

	response := u.Message(true, "Discipline has been created")
	response["discipline"] = discipline
	return response
}

func UpdateDiscipline(id string, discipline *models.DisciplineJSON) map[string]interface{} {
	disciplineModel := &models.Discipline{}
	GetDB().First(&disciplineModel, "id = ?", id)
	users := make([]*models.User, 0)
	GetDB().Find(&users, "id in (?)", discipline.Users)

	if len(discipline.Name) > 0 {
		disciplineModel.Name = discipline.Name
	}

	if len(discipline.Term) > 0 {
		disciplineModel.Term = discipline.Term
	}

	GetDB().Model(&disciplineModel).Association("Users").Append(&users)

	err := GetDB().Save(&disciplineModel).Error

	if err != nil {
		return u.Message(false, "Failed to update discipline, connection error.")
	}

	response := u.Message(true, "Discipline has been updated")
	response["discipline"] = &disciplineModel
	return response
}

func GetDisciplines() []*models.Discipline {
	disciplines := make([]*models.Discipline, 0)
	err := GetDB().Preload("Users").Table("disciplines").Find(&disciplines).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return disciplines
}
