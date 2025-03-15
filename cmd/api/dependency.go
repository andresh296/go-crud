package api

import (
	"github.com/andresh296/go-crud/config"
	"github.com/andresh296/go-crud/internal/domain/user"
	repo "github.com/andresh296/go-crud/internal/platform/user"
)

type Dependencies struct {
    userRepo   user.Repository
    config     *config.Config

}

func initDependencies() *Dependencies {
    cfg := config.Load()
    db, err := repo.GetDB(cfg.Database)
    if err != nil {
        panic("error get db")
    }
    
   
    userRepo := repo.NewRepository(db)

    

    return &Dependencies{
        userRepo:   userRepo,
        config:     &cfg,
    }
}