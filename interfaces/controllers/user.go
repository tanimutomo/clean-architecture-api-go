package controllers

import (
	"net/http"

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
	tokenHandler token.TokenHandler,
) *UserController {
	return &UserController{
		Service: service.UserService{
			Repository: &database.UserRepository{
				DBHandler: dbHandler,
			},
			Tokenizer: &token.Tokenizer{
				TokenHandler: tokenHandler,
			},
		},
	}
}

func (controller *UserController) Signup(c Context) {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		BadRequestError(c, "Invalid request format.")
		return
	}
	user, err := controller.Service.Signup(user)
	if err != nil {
		switch e := err.(type) {
		case *domain.ErrorWithStatus:
			SendErrorResponse(c, e.Status, e.Message)
		}
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) Login(c Context) {
	var loginUser domain.LoginUser
	if err := c.Bind(&loginUser); err != nil {
		BadRequestError(c, "Invalid request format.")
		return
	}
	user, token, err := controller.Service.Login(loginUser)
	if err != nil {
		switch e := err.(type) {
		case *domain.ErrorWithStatus:
			SendErrorResponse(c, e.Status, e.Message)
		}
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"user": user, "token": token})
}

func (controller *UserController) Authenticate(c Context) {
	// Get token from request header
	var header domain.HeaderWithToken
	err := c.BindHeader(&header)
	if err != nil {
		BadRequestError(c, "Invalid request format.")
		return
	}

	// Verify token
	tokenString := domain.Token(header.Authorization)
	err = controller.Service.Authenticate(tokenString)
	if err != nil {
		switch e := err.(type) {
		case *domain.ErrorWithStatus:
			SendErrorResponse(c, e.Status, e.Message)
		}
		return
	}
}
