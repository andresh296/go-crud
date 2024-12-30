package api

import (
	"github.com/andresh296/go-crud/config"
	domain "github.com/andresh296/go-crud/internal/domain/user"
	repo "github.com/andresh296/go-crud/internal/platform/user"
)

type Dependencies struct {
	user domain.Repository
}

func initDependencies() *Dependencies {
	cfg := config.Load()
	db, err := repo.GetDB(cfg.Database)
	if err != nil {
		panic("error get db")
	}
	userRepo := repo.NewRepository(db)

	return &Dependencies{
		user: userRepo,
	}
}