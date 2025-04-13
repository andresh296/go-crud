package api

import (
	"encoding/json"
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
		email := c.Param("email")

		user, err := h.service.GetUserByEmail(email)
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
        data, exists := c.Get("validatedData")
        if !exists {
            h.HandleError(c, ErrInvalidJSONFormat)
            return
        }

		jsonBytes, err := json.Marshal(data)
        if err != nil {
            h.HandleError(c, ErrUnmarshalBody)
            return
        }

        var userRequest UserRequest 
        if err := json.Unmarshal(jsonBytes, &userRequest); err != nil {
            h.HandleError(c, ErrUnmarshalBody)
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

func (h handler) Login() func(c *gin.Context) {
	return func(c *gin.Context) {

		data, exists := c.Get("validatedData")
        if !exists {
            h.HandleError(c, ErrInvalidJSONFormat)
            return
        }

		var userLogin UserLogin
		jsonBytes, err := json.Marshal(data)
		if err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

		if err := json.Unmarshal(jsonBytes, &userLogin); err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

	

		user, token, err := h.service.Login(userLogin.ToDomain())
		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := LoginResponse{
			ID:    user.ID,
			Name:  user.Name,
			Age:   user.Age,
			Email: user.Email,
			Token: token,
		}

		c.JSON(http.StatusOK, response)
	}
}
