package db

import (
	"classtime/models"
	u "classtime/utils"
	"fmt"
)

func CreateDiscipline(disciplineJSON *models.DisciplineJSON) map[string]interface{} {
	discipline := models.Discipline{}
	discipline.Name = disciplineJSON.Name
	discipline.Term = disciplineJSON.Term
	discipline.Classroom = disciplineJSON.Classroom
	discipline.WeekDays = models.ParseWeekDays(disciplineJSON.WeekDays)
	discipline.BeginAt = disciplineJSON.BeginAt
	discipline.EndAt = disciplineJSON.EndAt
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

	if len(discipline.Classroom) > 0 {
		disciplineModel.Classroom = discipline.Classroom
	}

	if len(discipline.WeekDays) > 0 {
		disciplineModel.WeekDays = models.ParseWeekDays(discipline.WeekDays)
	}

	if len(users) > 0 {
		GetDB().Model(&disciplineModel).Association("Users").Append(&users)
	}

	if len(usersRemove) > 0 {
		GetDB().Model(&disciplineModel).Association("Users").Delete(&usersRemove)
	}

	err := GetDB().Save(&disciplineModel).Error

	if err != nil {
		return u.Message(false, "Failed to update discipline, connection error.")
	}

	response := u.Message(true, "Discipline has been updated")
	GetDB().Preload("Users").Preload("Grid").First(&disciplineModel, "id = ?", id)
	disciplineModel.WeekDaysArray = models.GetDays(disciplineModel.WeekDays)
	response["discipline"] = &disciplineModel
	return response
}

func UpdateClassroom(id string, classroom string) map[string]interface{} {
	disciplineModel := &models.Discipline{}
	GetDB().First(&disciplineModel, "id = ?", id)
	disciplineModel.Classroom = classroom

	err := GetDB().Save(&disciplineModel).Error

	if err != nil {
		return u.Message(false, "Failed to update discipline, connection error.")
	}

	response := u.Message(true, "Discipline has been updated")
	response["discipline"] = &disciplineModel
	return response
}

func GetDisciplines(user *models.User) []*models.Discipline {
	disciplines := make([]*models.Discipline, 0)
	var err error
	if user.Role >= 2 {
		err = GetDB().Preload("Users").Preload("Grid").Table("disciplines").Find(&disciplines).Error
	} else {
		err = GetDB().Preload("Grid").Preload("Users").Model(&user).Related(&disciplines, "Disciplines").Error
	}

	for _, value := range disciplines {
		if value.WeekDays != "" {
			value.WeekDaysArray = models.GetDays(value.WeekDays)
		}
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return disciplines
}
