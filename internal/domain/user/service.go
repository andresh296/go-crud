package user

import (
    "github.com/andresh296/go-crud/config"
    "github.com/andresh296/go-crud/internal/platform/security"
    "golang.org/x/crypto/bcrypt"
)

type Repository interface {
    GetByID(id string) (*User, error)
    GetUserByEmail(email string) (*User, error)
    Save(user User) error
}

type Service interface {
    GetByID(id string) (*User, error)
    GetUserByEmail(email string) (*User, error)
    Save(user User) (User, error)
    Login(user User) (*User, string, error)
}

type service struct {
    repository Repository
    authService *AuthService
}

func NewService(repo Repository, authService *AuthService) Service {
    return &service{
        repository: repo,
        authService: authService,
    }
}

func (s service) GetByID(id string) (*User, error) {
    return s.repository.GetByID(id)
}

func (s service) GetUserByEmail(email string) (*User, error) {
    return s.repository.GetUserByEmail(email)
}

func (s service) Save(user User) (User, error) {
    user.setID()
    user.hashPassword()
    err := s.repository.Save(user)
    if err != nil {
        return User{}, err
    }

    return user, nil
}

func (u User) comparePassword(password string) error {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err
}

func (s service) Login(user User) (*User, string, error) {
    userFound, err := s.GetUserByEmail(user.Email)
    if err != nil {
        return nil, "", err
    }

    if err := userFound.comparePassword(user.Password); err != nil {
        return nil, "", err
    }

    token, err := s.authService.GenerateToken(userFound.ID, userFound.Email)
    if err != nil {
        return nil, "", err
    }

    return userFound, token, nil
}

// AuthService representa el servicio de autenticación
type AuthService struct {
    config config.AuthConfig
}

// NewAuthService inicializa y retorna una nueva instancia de AuthService
func NewAuthService(config config.AuthConfig) *AuthService {
    return &AuthService{
        config: config,
    }
}

// GenerateToken genera un token JWT para un usuario
func (a *AuthService) GenerateToken(userID, email string) (string, error) {
    cfg := config.Load()
    return security.GenerateJWT(
        userID,
        email,
        cfg.JWT.SecretKey,
        cfg.JWT.ExpirationTime,
    )
}