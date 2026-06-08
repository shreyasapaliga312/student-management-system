package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var updateStudent = &models.Student{}
	utils.ParseBody(r, updateStudent)

	if err := utils.ValidateStruct(updateStudent); err != nil { // Validate the informations passed
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id := utils.GetId(r, "studentId")

	student, db := models.GetStudentbyId(id)

	if student.ID == 0 { // Student doesn't exist on the db
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Student not found.")
		return
	}

	if student.ClassId != 0 && !models.ClassExist(int64(student.ClassId)) { // Class does not exist
		utils.WriteErrorResponse(w, http.StatusBadRequest, "Class does not exist.")
		return
	}

	// The classId can only be updated in the enroll_student_in_class

	if updateStudent.Name != "" && updateStudent.Name != student.Name {
		student.Name = updateStudent.Name
	}

	if updateStudent.GradeAverage > 0 && updateStudent.GradeAverage != student.GradeAverage {
		student.GradeAverage = updateStudent.GradeAverage
	}

	if updateStudent.TotalFaults != student.TotalFaults {
		student.TotalFaults = updateStudent.TotalFaults
	}

	if err := db.Save(&student).Error; err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
