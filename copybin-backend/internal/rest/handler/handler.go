package handler

import (
	"copybin/internal/config"
	"copybin/internal/impl/integration"
	"copybin/internal/impl/pasta"
	"copybin/internal/impl/user"
	"copybin/internal/storage"
)

type Handler struct {
	cfg              *config.Config
	userStore        *user.Store
	pastaStore       *pasta.Store
	integrationStore *integration.Store
}

func NewHandler(cfg *config.Config) *Handler {
	db := storage.New(cfg.DBURL)

	userRepo := user.NewStore(db)
	pastaRepo := pasta.NewStore(db)
	integrationStore := integration.NewStore(db)

	return &Handler{
		cfg:              cfg,
		userStore:        userRepo,
		pastaStore:       pastaRepo,
		integrationStore: integrationStore,
	}
}
