package controllers

// TODO: Ver status de erros corretos
import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (s SystemController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := utils.GetId(r, "studentId")

	student, err := models.DeleteStudent(id)

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if student.ID == 0 {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Student does not exist")
		return
	}

	utils.WriteJSONResponse(w, http.StatusOK, student)
}
