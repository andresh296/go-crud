package api
import (
    "github.com/andresh296/go-crud/config"
    "github.com/andresh296/go-crud/internal/domain/user"
    repo "github.com/andresh296/go-crud/internal/platform/user"
)

type Dependencies struct {
    user   user.Repository
    config *config.Config
    auth   *user.AuthService
}

func initDependencies() *Dependencies {
    cfg := config.Load()
    db, err := repo.GetDB(cfg.Database)
    if err != nil {
        panic("error get db")
    }
    userRepo := repo.NewRepository(db)
    userService := user.NewService(userRepo, user.NewAuthService(cfg.Auth)) // Crea un nuevo objeto Service

    return &Dependencies{
        user:   userService, // Usa el objeto Service en lugar de crear un nuevo objeto AuthService
        config: &cfg,
        auth:   userService.authService, // Accede al campo authService del objeto Service
    }
}