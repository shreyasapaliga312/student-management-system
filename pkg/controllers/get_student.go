package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) GetStudents(w http.ResponseWriter, r *http.Request) {
	students := models.GetAllStudents()
	utils.WriteJSONResponse(w, http.StatusOK, students)
}

func (s SystemController) GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(r, "studentId")

	student, _ := models.GetStudentbyId(id)

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
