package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

func (u UserController) Login(w http.ResponseWriter, r *http.Request) {
	credentials := models.Credentials{}
	utils.ParseBody(r, &credentials)

	if err := utils.ValidateStruct(credentials); err != nil { // Validate the informations passed
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	dbUser, err := models.GetUserById(credentials.Id)
	if err != nil {
		utils.WriteErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	if dbUser.Password != credentials.Password {
		utils.WriteErrorResponse(w, http.StatusNotFound, "Invalid Credentials!")
		return
	}

	payload := utils.Payload{
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Id:       dbUser.ID,
	}

	token, err := utils.GenerateJWTToken(payload)

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
		return
	}

	output := ResponseOutput{
		Token: token,
		User:  *dbUser,
	}
	utils.WriteJSONResponse(w, http.StatusOK, output)
}
