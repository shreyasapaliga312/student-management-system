package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) GetClasses(w http.ResponseWriter, r *http.Request) {
	classes := models.GetAllClasses()
	utils.WriteJSONResponse(w, http.StatusOK, classes)
}

func (s SystemController) GetClassById(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(r, "classId")

	class, _, err := models.GetClassById(id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, class)
}
