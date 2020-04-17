package controllers

import (
	"net/http"
	"strconv"

	"github.com/tanimutomo/clean-architecture-api-go/domain"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/database"
	"github.com/tanimutomo/clean-architecture-api-go/service"
)

type TagController struct {
	Service service.TagService
}

func NewTagController(dbHandler, database.DBHandler) *TagController {
	return &TagController{
		Service: service.TagService{
			Repository: &database.TagRepository{
				DBHandler: dbHandler,
			},
		},
	}
}
