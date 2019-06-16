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

	grid := &models.Grid{}
	grid.DisciplineId = discipline.ID
	GetDB().Create(grid)

	response := u.Message(true, "Discipline has been created")
	response["discipline"] = discipline
	return response
}

func UpdateDiscipline(id string, discipline *models.DisciplineJSON) map[string]interface{} {
	disciplineModel := &models.Discipline{}
	GetDB().First(&disciplineModel, "id = ?", id)
	users := make([]*models.User, 0)
	usersRemove := make([]*models.User, 0)
	GetDB().Find(&users, "id in (?)", discipline.Users)
	GetDB().Find(&usersRemove, "id in (?)", discipline.UsersRemove)

	if len(discipline.Name) > 0 {
		disciplineModel.Name = discipline.Name
	}

	if len(discipline.Term) > 0 {
		disciplineModel.Term = discipline.Term
	}

	GetDB().Model(&disciplineModel).Association("Users").Append(&users)
	GetDB().Model(&disciplineModel).Association("Users").Delete(&usersRemove)

	err := GetDB().Save(&disciplineModel).Error

	if err != nil {
		return u.Message(false, "Failed to update discipline, connection error.")
	}

	response := u.Message(true, "Discipline has been updated")
	GetDB().Preload("Users").Preload("Grid").First(&disciplineModel, "id = ?", id)
	response["discipline"] = &disciplineModel
	return response
}

func GetDisciplines(user *models.User) []*models.Discipline {
	disciplines := make([]*models.Discipline, 0)
	var err error
	if user.Role >= 2 {
		err = GetDB().Preload("Users").Preload("Grid").Table("disciplines").Find(&disciplines).Error
	} else {
		err = GetDB().Preload("Grid").Model(&user).Related(&disciplines, "Disciplines").Error
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return disciplines
}
