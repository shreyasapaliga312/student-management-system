package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) EnrollStudentInClass(w http.ResponseWriter, r *http.Request) {
	classId := utils.GetId(r, "classId")
	studentId := utils.GetId(r, "studentId")

	if err := models.EnrollStudentInClass(uint(studentId), uint(classId)); err != nil {
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	student, _ := models.GetStudentbyId(studentId)

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
