package controllers

import (
	"net/http"

	"github.com/anaard/simple-student-management/pkg/models"
	"github.com/anaard/simple-student-management/pkg/utils"
)

type ResponseOutput struct {
	User  models.User
	Token string
}

func (u UserController) Register(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	utils.ParseBody(r, &newUser)

	if err := utils.ValidateStruct(newUser); err != nil { // Validate the informations passed
		utils.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user := newUser.CreateUser()

	payload := utils.Payload{
		Username: user.Username,
		Email:    user.Email,
		Id:       user.ID,
	}

	token, err := utils.GenerateJWTToken(payload)

	if err != nil {
		utils.WriteErrorResponse(w, http.StatusInternalServerError, "Failed To Generate New JWT Token!")
		return
	}

	output := ResponseOutput{
		Token: token,
		User:  *user,
	}

	utils.WriteJSONResponse(w, http.StatusOK, output)
}
