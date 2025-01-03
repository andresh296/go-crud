package api

import (
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service domain.Service
}

func New(service domain.Service) *handler {
	return &handler{
		service: service,
	}
}
func (h handler) GetAll() func(c *gin.Context) {
	return func(c *gin.Context) {
		users, err := h.service.GetAll()
		if err != nil {
			h.HandleError(c, err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func (h handler) Save() func(c *gin.Context) {
	return func(c *gin.Context) {
		var userRequest UserRequest
		err := c.BindJSON(&userRequest)
		if err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

		err = userRequest.Validate()
		if err != nil {
			h.HandleError(c, err)
			return
		}

		user, err := h.service.Save(userRequest.ToDomain())
		if err != nil {
			h.HandleError(c, err)
			return
	}

		response := UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Age:   user.Age,
			Email: user.Email,
		}
		c.JSON(http.StatusCreated, response)
	}
}

