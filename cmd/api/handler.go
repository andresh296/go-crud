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
        var userRequest UserRequest
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			h.HandleError(c, ErrInvalidJSONFormat)
			return
		}
       

        user, err := h.service.Save(userRequest.ToDomain())
        if err != nil {

            h.HandleError(c, domain.ErrUserCannotSave)
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
        var userLogin UserLogin
        if err := c.ShouldBindJSON(&userLogin); err != nil {
            h.HandleError(c, ErrInvalidJSONFormat)
            return
        }

        user, token, err := h.service.Login(userLogin.ToDomain())
        if err != nil {
            h.HandleError(c, ErrValidationUser)
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
