package api

import (
	"net/http"

	domain "github.com/andresh296/go-crud/internal/domain/user"

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

func (h handler) GetByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id") // Obtiene el ID de los parámetros de la URL

		user, err := h.service.GetByID(id)
		if err != nil {
			h.HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
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
