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
	GetDB().First(disciplineModel, "id = ?", id)

	disciplineModel.Name = discipline.Name
	disciplineModel.Term = discipline.Term

	err := GetDB().Save(disciplineModel).Error

	if err != nil {
		return u.Message(false, "Failed to update discipline, connection error.")
	}

	response := u.Message(true, "Discipline has been updated")
	response["discipline"] = discipline
	return response
}

func GetDisciplines() []*models.Discipline {
	disciplines := make([]*models.Discipline, 0)
	err := GetDB().Table("disciplines").Find(&disciplines).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return disciplines
}