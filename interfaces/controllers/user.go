package controllers

import (
	"net/http"
	"strconv"

	"github.com/tanimutomo/clean-architecture-api-go/domain"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/database"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/token"
	"github.com/tanimutomo/clean-architecture-api-go/service"
)

type UserController struct {
	Service service.UserService
}

func NewUserController(
	dbHandler database.DBHandler,
	tokenHandler token.TokenHandler
	) *UserController {
	return &UserController{
		Service: service.UserService{
			Repository: &database.UserRepository{
				DBHandler: dbHandler,
			},
			Token: &token.Token{
				TokenHandler: tokenHandler,
			},
		},
	}
}

func (controller *UserController) Signup(c Context) {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		// TODO
	}
	user, err := controller.Service.Signup(user)
	if err != nil {
		// TODO
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) Login(c Context) {
	loginUser := domain.LoginUser{}
	if err := c.Bind(&loginUser); err != nil {
		// TODO
	}
	user, token, err := controller.Service.Login(loginUser)
	if err != nil {
		// TODO
		return
	}
	c.JSON(http.StatusOK, {"user":, user, "token": token})
}

func (controller *UserController) Authenticate(c Context) {
	// Get token from request header
	token := domain.Token{}
	token, err = c.Request.Header.Get("token")
	if err != nil {
		// TODO
		return
	}

	// Verify token
	err := controller.Service.Authenticate(id)
	if err != nil {
		// TODO: Send error response and abort
		return
	}
}
