package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/utils"
	"github.com/anaard/simple-student-management/pkg/models"
	
)

func (s SystemController) RemoveStudentFromClass(w http.ResponseWriter, r * http.Request) {
	studentId := utils.GetId(r, "studentId")
	classId := utils.GetId(r, "classId")
	
	if err := models.RemoveStudentFromClass(uint(studentId), uint(classId)); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	
	student, _ := models.GetStudentbyId(studentId)
	
	utils.WriteJSONResponse(w, http.StatusOK, student)
	
}
