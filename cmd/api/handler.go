package api

import (
	"net/http"

	"github.com/andresh296/go-crud/config"
	domain "github.com/andresh296/go-crud/internal/domain/user"
	"github.com/andresh296/go-crud/internal/platform/security"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service domain.Service
	cfg     *config.Config
}

func New(service domain.Service, cfg *config.Config) *handler {
	return &handler{
		service: service,
		cfg:     cfg,
	}
}

func (h handler) GetUserByEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")

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

func (h handler) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		var UserLogin UserLogin
		err := c.BindJSON(UserLogin)
		if err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

		err = UserLogin.Validate()
		if err != nil {
			h.HandleError(c, err)
			return
		}

		user, err := h.service.Login(UserLogin.ToDomain())
		if err != nil {
			h.HandleError(c, err)
			return
		}

		// Generar JWT
		token, err := security.GenerateJWT(
			user.ID,
			user.Email,
			h.cfg.JWT.SecretKey,
			h.cfg.JWT.ExpirationTime,
		)

		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := LoginResponse{
			ID:    user.ID,
			Email: user.Email,
			Token: token,
		}

		c.JSON(http.StatusOK, response)
	}
}
