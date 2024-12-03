package handler

import (
	"github.com/AndreCDiniz/ApiRestFull-Go/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpenningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %s", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	id := context.Query("id")
	if id == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(context, http.StatusNotFound, "opening not found")
		return
	}

	//Update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	// Save opening
	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("error update opening: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error update opening")
		return
	}
	sendSuccess(context, "update-opening", opening)
}
