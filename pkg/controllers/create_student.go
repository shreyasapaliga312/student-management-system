package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	createStudent := &models.Student{}
	utils.ParseBody(r, createStudent)

	if err := utils.ValidateStruct(createStudent); err != nil { // Validate the informations passed
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// Verify if a student with the same name does not exist
	if models.StudentExistName(createStudent.Name) {
		utils.WriteErrorResponse(w, http.StatusBadRequest, "student name already exists")
		return
	}

	// Not possible to enroll the student in a class from here
	if createStudent.ClassId != 0 {
		createStudent.ClassId = 0
	}

	student := createStudent.CreateStudent()

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
