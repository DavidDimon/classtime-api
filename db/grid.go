package db

import (
	"classtime/models"
	u "classtime/utils"
)

func AddAlert(id string, user *models.User, alert *models.Alert) map[string]interface{} {
	grid := &models.Grid{}
	GetDB().Table("grids").First(&grid, "id = ?", id)
	if grid.ID == 0 { //Grid not found!
		return nil
	}
	alert.GridId = grid.ID
	alert.Username = user.Name
	alert.UserId = user.ID
	GetDB().Create(&alert)
	GetDB().Model(&grid).Association("Alerts").Append(&alert)

	response := u.Message(true, "Alert has been added")
	GetDB().Preload("Alerts").Table("grids").First(&grid, "id = ?", id)
	response["grid"] = grid
	return response
}

func RemoveAlert(id string, user *models.User) map[string]interface{} {
	grid := &models.Grid{}
	alert := &models.Alert{}
	GetDB().Table("alerts").First(&alert, "id = ?", id)
	GetDB().Table("grids").First(&grid, "id = ?", alert.GridId)
	if alert.ID == 0 { //alert not found!
		return nil
	}
	if (alert.UserId != user.ID && user.Role < 1) {
		return u.Message(false, "Permission denied")
	}

	GetDB().Model(&grid).Association("Alerts").Delete(&alert)
	GetDB().Delete(&alert, "id = ?", alert.ID)

	response := u.Message(true, "Alert has been removed")
	GetDB().Preload("Alerts").Table("grids").First(&grid, "id = ?", grid.ID)
	response["grid"] = grid
	return response
}

func GetGrid(id string) *models.Grid {
	grid := &models.Grid{}
	GetDB().Preload("Alerts").Table("grids").Find(&grid, "id = ?", id)
	return grid
}
