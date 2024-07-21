package user

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

func (s Store) CreateIfNotExists(username string) (*model.User, error) {
	m := &model.User{Username: username, Role: "normie"}
	cond := &model.User{Username: username}

	if result := s.db.Where(cond).FirstOrCreate(m); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}

func (s Store) Get(id uint) (*model.User, error) {
	m := &model.User{}
	cond := &model.User{Model: &gorm.Model{ID: id}}

	if result := s.db.First(m, cond); result.Error != nil {
		return nil, result.Error
	}

	return m, nil
}

func (s Store) GetByUsername(username string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}
