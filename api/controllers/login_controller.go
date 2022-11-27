package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aulkiller/Deall_BE_Test/api/auth"
	"github.com/aulkiller/Deall_BE_Test/api/models"
	"github.com/aulkiller/Deall_BE_Test/api/responses"
	"github.com/aulkiller/Deall_BE_Test/api/utils/formaterror"
)

type LoginResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Token   string `json:"token"`
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userGotten, token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	// // If use session cookie - Unused ATM replaced by JWT
	// err = auth.SetSession(userGotten.Role, w)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	LoginResponse := LoginResponse{
		Message: "Succesfully logged in",
		Name:    userGotten.Name,
		Role:    userGotten.Role,
		Token:   token,
	}
	responses.JSON(w, http.StatusOK, LoginResponse)
}

func (server *Server) SignIn(username, password string) (*models.User, string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return nil, "", err
	}
	err = models.CheckPasswordHash(user.Password, password)
	if err != nil {
		return nil, "", err
	}
	token, err := auth.CreateToken(user.ID, user.Role)
	return &user, token, err
}
