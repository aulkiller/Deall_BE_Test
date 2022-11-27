package controllers

import (
	"net/http"

	"github.com/aulkiller/Deall_BE_Test/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "This is landing page, use /users and /login to interact with this application")
}
