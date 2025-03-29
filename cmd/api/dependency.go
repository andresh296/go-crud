package api

import (
	"github.com/andresh296/go-crud/config"
	"github.com/andresh296/go-crud/internal/domain/user"
	"github.com/andresh296/go-crud/internal/platform/schema"
	repo "github.com/andresh296/go-crud/internal/platform/user"
)

type Dependencies struct {
    userRepo   user.Repository
    config     *config.Config
    validators *schema.Validators
}

func initDependencies() *Dependencies {
    cfg := config.Load()
    db, err := repo.GetDB(cfg.Database)
    if err != nil {
        panic("error get db")
    }
    
   
    userRepo := repo.NewRepository(db)

    validators, err := schema.NewValidator()
    if err != nil {
        panic("error creating validators")
    }

    return &Dependencies{
        userRepo:   userRepo,
        config:     &cfg,
        validators: validators,
    }
}