package repository

import "sync"

type Repository struct {
	DailyWord string
	Mutex     sync.Mutex
}

func NewRepository() (*Repository, error) {
	var repo Repository
	repo.DailyWord = "apple"
	return &repo, nil
}
