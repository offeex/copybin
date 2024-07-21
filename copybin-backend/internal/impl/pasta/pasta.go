package pasta

import (
	"copybin/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (rcv Store) Create(text string, userID uint) (*model.Pasta, error) {
	m := &model.Pasta{Text: text, UserID: userID}

	if result := rcv.db.Create(m); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}

func (rcv Store) Update(id uuid.UUID, text string) error {
	//TODO implement me
	panic("implement me")
}

func (rcv Store) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (rcv Store) Get(id uuid.UUID) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (rcv Store) List(username string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (rcv Store) ListAll(limit int) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
