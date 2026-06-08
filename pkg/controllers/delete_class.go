package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) DeleteClass(w http.ResponseWriter, r *http.Request) { // TODO: Change students classes to None when a class is deleted
	id := utils.GetId(r, "classId")

	class, err := models.DeleteClass(id)

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if class.ID == 0 {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "The class 0 does not exist")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, class)
}
