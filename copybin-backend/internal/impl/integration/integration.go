package integration

import (
	"copybin/internal/model"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (rcv Store) CreateIfNotExists(
	userID uint,
	provider string,
	providerID string,
	accessToken string,
	refreshToken string,
) error {
	m := &model.Integration{
		Provider:     provider,
		ProviderID:   providerID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       userID,
	}
	cond := model.Integration{Provider: provider, UserID: userID}
	assign := model.Integration{AccessToken: accessToken, RefreshToken: refreshToken}

	if result := rcv.db.Where(cond).Assign(assign).FirstOrCreate(m); result.Error != nil {
		return result.Error
	}

	return nil
}
