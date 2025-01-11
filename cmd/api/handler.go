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


func (h handler) GetUserByEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		email:=c.Param("email")

		user, err := h.service.GetUserByEmail(email) // Usa una función del servicio para obtener el usuario por su email
		if err != nil {
			h.HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (h handler) GetByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

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
